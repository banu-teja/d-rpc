package main

import (
	"bytes"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// fakePaymentChannel mocks CloseChannel
type fakePaymentChannel struct {
	called    bool
	channelID [32]byte
	amount    *big.Int
	sig       []byte
}
func (f *fakePaymentChannel) CloseChannel(auth *bind.TransactOpts, channelID [32]byte, amount *big.Int, sig []byte) (*types.Transaction, error) {
	f.called = true
	f.channelID = channelID
	f.amount = amount
	f.sig = sig
	return nil, nil
}

// fakeProviderReg mocks UpdateQoS
type fakeProviderReg struct {
	called bool
	from   common.Address
	score  *big.Int
}
func (f *fakeProviderReg) UpdateQoS(auth *bind.TransactOpts, provider common.Address, score *big.Int) (*types.Transaction, error) {
	f.called = true
	f.from = provider
	f.score = score
	return nil, nil
}

func TestHandleRPCRequest_Success(t *testing.T) {
	// Stub backend RPC server
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{"jsonrpc": "2.0", "result": "0x1", "id": 1})
	}))
	defer backend.Close()

	// Generate a random private key
	priv, _ := crypto.GenerateKey()

	// Prepare dummy payment data
	var zeroID [32]byte
	hexID := hexutil.Encode(zeroID[:])
	hexSig := hexutil.Encode([]byte{0x1, 0x2})

	// Set up RPCServer with fakes
	srv := &RPCServer{
		config:         Config{EthRPCURL: backend.URL},
		paymentChannel: &fakePaymentChannel{},
		providerReg:    &fakeProviderReg{},
		privateKey:     priv,
	}

	// Build request payload
	payload := map[string]interface{}{
		"method": "eth_blockNumber",
		"params": []interface{}{},
		"payment": map[string]string{
			"channelId": hexID,
			"amount":    "1",
			"signature": hexSig,
		},
	}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
	w := httptest.NewRecorder()

	srv.handleRPCRequest(w, req)
	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("Expected 200, got %d", res.StatusCode)
	}

	// Verify backend result forwarded
	var resp map[string]interface{}
	json.NewDecoder(res.Body).Decode(&resp)
	if resp["result"] != "0x1" {
		t.Errorf("Expected result 0x1, got %v", resp["result"])
	}

	// Verify payment processing
	fpc := srv.paymentChannel.(*fakePaymentChannel)
	if !fpc.called {
		t.Error("CloseChannel not called")
	}

	// Verify QoS update
	fpr := srv.providerReg.(*fakeProviderReg)
	if !fpr.called {
		t.Error("UpdateQoS not called")
	}
}
