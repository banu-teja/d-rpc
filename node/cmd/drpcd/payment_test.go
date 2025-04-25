package main

import (
	"bytes"
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPaymentChannel struct {
	mock.Mock
}

func (m *MockPaymentChannel) CloseChannel(auth *bind.TransactOpts, channelID [32]byte, amount *big.Int, signature []byte) (*types.Transaction, error) {
	args := m.Called(auth, channelID, amount, signature)
	first := args.Get(0)
	if first == nil {
		return nil, args.Error(1)
	}
	return first.(*types.Transaction), args.Error(1)
}

type MockQoSService struct {
	mock.Mock
}

func (m *MockQoSService) UpdateQoS(auth *bind.TransactOpts, provider common.Address, score *big.Int) (*types.Transaction, error) {
	args := m.Called(auth, provider, score)
	first := args.Get(0)
	if first == nil {
		return nil, args.Error(1)
	}
	return first.(*types.Transaction), args.Error(1)
}

func TestPaymentHandler_Success(t *testing.T) {
	// Stub backend RPC server for forwarding
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{"jsonrpc": "2.0", "result": "0x1", "id": 1})
	}))
	defer backend.Close()

	mockChannel := new(MockPaymentChannel)
	mockQoS := new(MockQoSService)

	// Mock successful payment channel closure
	mockChannel.On("CloseChannel", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
		types.NewTransaction(0, common.Address{}, nil, 0, nil, nil), nil)

	// Mock QoS update
	mockQoS.On("UpdateQoS", mock.Anything, mock.Anything, mock.Anything).Return(
		types.NewTransaction(1, common.Address{}, nil, 0, nil, nil), nil)

	// Generate a test private key
	pk, _ := crypto.GenerateKey()

	srv := &RPCServer{
		paymentChannel: mockChannel,
		providerReg:    mockQoS,
		privateKey:     pk,
		config: Config{
			EthRPCURL: backend.URL,
			ContractAddrs: ContractAddresses{
				PaymentChannel:   common.HexToAddress("0x123"),
				ProviderRegistry: common.HexToAddress("0x456"),
				StakeToken:       common.HexToAddress("0x789"),
			},
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(srv.handleRPCRequest))
	defer ts.Close()

	reqBody := bytes.NewBufferString(`{
		"method": "eth_blockNumber",
		"params": [],
		"payment": {
			"channelId": "0x1234",
			"amount": "1000",
			"signature": "0xabcd"
		}
	}`)

	resp, err := http.Post(ts.URL, "application/json", reqBody)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// verify response body contains forwarded result
	var respData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&respData)
	assert.NoError(t, err)
	assert.Equal(t, "0x1", respData["result"])

	mockChannel.AssertExpectations(t)
	mockQoS.AssertExpectations(t)
}

func TestPaymentHandler_Failure(t *testing.T) {
	// Generate a test private key
	pk, _ := crypto.GenerateKey()
	mockChannel := new(MockPaymentChannel)
	mockQoS := new(MockQoSService)

	// Mock failed payment channel closure to return error
	mockChannel.On("CloseChannel", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
		nil, assert.AnError)

	srv := &RPCServer{
		privateKey:     pk,
		paymentChannel: mockChannel,
		providerReg:    mockQoS,
		config: Config{
			EthRPCURL: "http://localhost:8545",
			ContractAddrs: ContractAddresses{
				PaymentChannel:   common.HexToAddress("0x123"),
				ProviderRegistry: common.HexToAddress("0x456"),
				StakeToken:       common.HexToAddress("0x789"),
			},
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(srv.handleRPCRequest))
	defer ts.Close()

	reqBody := bytes.NewBufferString(`{
		"method": "eth_blockNumber",
		"params": [],
		"payment": {
			"channelId": "0x1234",
			"amount": "1000",
			"signature": "0xabcd"
		}
	}`)

	resp, err := http.Post(ts.URL, "application/json", reqBody)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusPaymentRequired, resp.StatusCode)

	// verify error message in body
	data, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(data), "Payment processing failed")

	// Verify QoS wasn't updated on failure
	mockQoS.AssertNotCalled(t, "UpdateQoS")
}

// New test cases for extended coverage
func TestPaymentHandler_InvalidPayment(t *testing.T) {
	pk, _ := crypto.GenerateKey()
	srv := &RPCServer{
		privateKey: pk,
		config: Config{
			EthRPCURL: "http://localhost:8545",
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(srv.handleRPCRequest))
	defer ts.Close()

	// Test missing payment field
	reqBody := bytes.NewBufferString(`{
		"method": "eth_blockNumber",
		"params": []
	}`)

	resp, err := http.Post(ts.URL, "application/json", reqBody)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	// Test invalid payment format
	reqBody = bytes.NewBufferString(`{
		"method": "eth_blockNumber",
		"params": [],
		"payment": "invalid"
	}`)

	resp, err = http.Post(ts.URL, "application/json", reqBody)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestPaymentHandler_InvalidSignature(t *testing.T) {
	pk, _ := crypto.GenerateKey()
	mockChannel := new(MockPaymentChannel)

	srv := &RPCServer{
		privateKey:     pk,
		paymentChannel: mockChannel,
		config: Config{
			EthRPCURL: "http://localhost:8545",
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(srv.handleRPCRequest))
	defer ts.Close()

	// Test with invalid signature format
	reqBody := bytes.NewBufferString(`{
		"method": "eth_blockNumber",
		"params": [],
		"payment": {
			"channelId": "0x1234",
			"amount": "1000",
			"signature": "invalid"
		}
	}`)

	resp, err := http.Post(ts.URL, "application/json", reqBody)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestPaymentHandler_ForwardingError(t *testing.T) {
	// Create backend that returns error
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer backend.Close()

	pk, _ := crypto.GenerateKey()
	mockChannel := new(MockPaymentChannel)
	mockQoS := new(MockQoSService)

	mockChannel.On("CloseChannel", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(
		types.NewTransaction(0, common.Address{}, nil, 0, nil, nil), nil)

	srv := &RPCServer{
		privateKey:     pk,
		paymentChannel: mockChannel,
		providerReg:    mockQoS,
		config: Config{
			EthRPCURL: backend.URL,
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(srv.handleRPCRequest))
	defer ts.Close()

	reqBody := bytes.NewBufferString(`{
		"method": "eth_blockNumber",
		"params": [],
		"payment": {
			"channelId": "0x1234",
			"amount": "1000",
			"signature": "0xabcd"
		}
	}`)

	resp, err := http.Post(ts.URL, "application/json", reqBody)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}
