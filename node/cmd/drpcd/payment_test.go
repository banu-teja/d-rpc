package main

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	return args.Get(0).(*types.Transaction), args.Error(1)
}

type MockQoSService struct {
	mock.Mock
}

func (m *MockQoSService) UpdateQoS(auth *bind.TransactOpts, provider common.Address, score *big.Int) (*types.Transaction, error) {
	args := m.Called(auth, provider, score)
	return args.Get(0).(*types.Transaction), args.Error(1)
}

func createSignedRequest(t *testing.T, pk *ecdsa.PrivateKey, channelID common.Hash, amount *big.Int, validUntil int64) *bytes.Buffer {
	from := crypto.PubkeyToAddress(pk.PublicKey)

	msg := crypto.Keccak256Hash(
		channelID.Bytes(),
		common.LeftPadBytes(amount.Bytes(), 32),
		common.LeftPadBytes(big.NewInt(validUntil).Bytes(), 32),
	)

	digest := crypto.Keccak256Hash(
		[]byte("\x19Ethereum Signed Message:\n32"),
		msg.Bytes(),
	)

	sig, err := crypto.Sign(digest.Bytes(), pk)
	assert.NoError(t, err)

	req := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_blockNumber",
		"params":  []interface{}{},
		"id":      1,
		"payment": map[string]interface{}{
			"channelId":  channelID.Hex(),
			"amount":     amount.String(),
			"signature":  hexutil.Encode(sig),
			"from":       from.Hex(),
			"validUntil": validUntil,
		},
	}

	body, _ := json.Marshal(req)
	return bytes.NewBuffer(body)
}

func TestPaymentHandler_Success(t *testing.T) {
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"jsonrpc": "2.0", "result": "0x1", "id": 1})
	}))
	defer backend.Close()

	pk, _ := crypto.GenerateKey()
	channelID := common.HexToHash("0x1234")
	amount := big.NewInt(1e18)
	validUntil := time.Now().Add(1 * time.Hour).Unix()

	mockChannel := new(MockPaymentChannel)
	mockQoS := new(MockQoSService)

	mockChannel.On("CloseChannel", mock.Anything, channelID, amount, mock.Anything).Return(
		types.NewTransaction(0, common.Address{}, nil, 0, nil, nil), nil)
	mockQoS.On("UpdateQoS", mock.Anything, mock.Anything, mock.Anything).Return(
		types.NewTransaction(1, common.Address{}, nil, 0, nil, nil), nil)

	srv := &RPCServer{
		privateKey:     pk,
		paymentChannel: mockChannel,
		providerReg:    mockQoS,
		config:         Config{EthRPCURL: backend.URL},
	}

	ts := httptest.NewServer(http.HandlerFunc(srv.handleRPCRequest))
	defer ts.Close()

	resp, err := http.Post(ts.URL, "application/json",
		createSignedRequest(t, pk, channelID, amount, validUntil))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	assert.Equal(t, "0x1", result["result"])
	assert.Contains(t, result["context"], "qos")
}

func TestPaymentHandler_ExpiredSignature(t *testing.T) {
	pk, _ := crypto.GenerateKey()
	channelID := common.HexToHash("0x1234")
	expiredValidUntil := time.Now().Add(-1 * time.Hour).Unix()

	srv := &RPCServer{
		privateKey: pk,
		config:     Config{EthRPCURL: "http://localhost:8545"},
	}

	ts := httptest.NewServer(http.HandlerFunc(srv.handleRPCRequest))
	defer ts.Close()

	resp, err := http.Post(ts.URL, "application/json",
		createSignedRequest(t, pk, channelID, big.NewInt(1e18), expiredValidUntil))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	assert.Contains(t, result["error"].(map[string]interface{})["message"], "expired")
}

func TestPaymentHandler_ReplayAttack(t *testing.T) {
	pk, _ := crypto.GenerateKey()
	channelID := common.HexToHash("0x1234")
	amount := big.NewInt(1e18)
	validUntil := time.Now().Add(1 * time.Hour).Unix()

	mockChannel := new(MockPaymentChannel)
	mockChannel.On("CloseChannel", mock.Anything, channelID, amount, mock.Anything).Return(
		types.NewTransaction(0, common.Address{}, nil, 0, nil, nil), nil).Once()

	srv := &RPCServer{
		privateKey:     pk,
		paymentChannel: mockChannel,
		config:         Config{EthRPCURL: "http://localhost:8545"},
	}

	ts := httptest.NewServer(http.HandlerFunc(srv.handleRPCRequest))
	defer ts.Close()

	// First request should succeed
	reqBody := createSignedRequest(t, pk, channelID, amount, validUntil)
	resp1, _ := http.Post(ts.URL, "application/json", reqBody)
	assert.Equal(t, http.StatusOK, resp1.StatusCode)

	// Second request with same params should fail
	resp2, _ := http.Post(ts.URL, "application/json", reqBody)
	assert.Equal(t, http.StatusBadRequest, resp2.StatusCode)

	var result map[string]interface{}
	json.NewDecoder(resp2.Body).Decode(&result)
	assert.Contains(t, result["error"].(map[string]interface{})["message"], "replay")
}

func TestPaymentHandler_InvalidSender(t *testing.T) {
	pk, _ := crypto.GenerateKey()
	channelID := common.HexToHash("0x1234")
	amount := big.NewInt(1e18)
	validUntil := time.Now().Add(1 * time.Hour).Unix()

	// Create valid request then tamper with sender
	reqBody := createSignedRequest(t, pk, channelID, amount, validUntil)
	var req map[string]interface{}
	json.Unmarshal(reqBody.Bytes(), &req)
	req["payment"].(map[string]interface{})["from"] = common.HexToAddress("0x0000").Hex()
	tamperedBody, _ := json.Marshal(req)

	srv := &RPCServer{
		privateKey: pk,
		config:     Config{EthRPCURL: "http://localhost:8545"},
	}

	ts := httptest.NewServer(http.HandlerFunc(srv.handleRPCRequest))
	defer ts.Close()

	resp, _ := http.Post(ts.URL, "application/json", bytes.NewBuffer(tamperedBody))
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
