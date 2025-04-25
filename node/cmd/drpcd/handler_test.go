package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// fakeChannel mocks the payment channel CloseChannel method.
type fakeChannel struct {
	called    bool
	channelID [32]byte
	amount    *big.Int
	sig       []byte
}

func (f *fakeChannel) CloseChannel(auth *bind.TransactOpts, cid [32]byte, amt *big.Int, sig []byte) (*types.Transaction, error) {
	f.called, f.channelID, f.amount, f.sig = true, cid, amt, sig
	return nil, nil
}

// fakeQoS mocks the QoSService UpdateQoS method.
type fakeQoS struct {
	called bool
	from   common.Address
	score  *big.Int
}

func (f *fakeQoS) UpdateQoS(auth *bind.TransactOpts, p common.Address, s *big.Int) (*types.Transaction, error) {
	f.called, f.from, f.score = true, p, s
	return nil, nil
}

// fakeChannelErr mocks CloseChannel returning an error.
type fakeChannelErr struct{}

func (f *fakeChannelErr) CloseChannel(auth *bind.TransactOpts, cid [32]byte, amt *big.Int, sig []byte) (*types.Transaction, error) {
	return nil, errors.New("payment failed")
}

func TestHandleRPCFailureModes(t *testing.T) {
	srv := &RPCServer{}
	cases := []struct {
		body   string
		status int
	}{
		{"not json", http.StatusBadRequest},
		{"{\"method\":\"x\",\"params\":[],\"payment\":{\"channelId\":\"0xzz\",\"amount\":\"1\",\"signature\":\"0x00\"}}", http.StatusBadRequest},
		{"{\"method\":\"x\",\"params\":[],\"payment\":{\"channelId\":\"0x" + strings.Repeat("a", 64) + "\",\"amount\":\"abc\",\"signature\":\"0x00\"}}", http.StatusBadRequest},
		{"{\"method\":\"x\",\"params\":[],\"payment\":{\"channelId\":\"0x" + strings.Repeat("a", 64) + "\",\"amount\":\"1\",\"signature\":\"0xzz\"}}", http.StatusBadRequest},
	}
	for _, c := range cases {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/", bytes.NewBufferString(c.body))
		srv.handleRPCRequest(w, r)
		if w.Code != c.status {
			t.Errorf("body=%q expected %d, got %d", c.body, c.status, w.Code)
		}
	}
}

func TestHandleRPCSuccess(t *testing.T) {
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"jsonrpc": "2.0", "result": "0xdead", "id": 1})
	}))
	defer backend.Close()

	fc := &fakeChannel{}
	fq := &fakeQoS{}
	key, _ := crypto.GenerateKey()
	srv := &RPCServer{
		paymentChannel: fc,
		providerReg:    fq,
		privateKey:     key,
		config:         Config{EthRPCURL: backend.URL},
	}

	zeroID := [32]byte{}
	payload := map[string]interface{}{
		"method": "eth_blockNumber",
		"params": []interface{}{},
		"payment": map[string]string{
			"channelId": hexutil.Encode(zeroID[:]),
			"amount":    "1",
			"signature": hexutil.Encode([]byte{0x01}),
		},
	}
	b, _ := json.Marshal(payload)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", bytes.NewBuffer(b))

	srv.handleRPCRequest(w, r)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var resp map[string]interface{}
	json.NewDecoder(w.Body).Decode(&resp)
	if resp["result"] != "0xdead" {
		t.Errorf("expected forwarded ‘0xdead’, got %v", resp["result"])
	}
	if !fc.called {
		t.Error("CloseChannel was not invoked")
	}
	// Verify CloseChannel parameters
	if fc.channelID != zeroID {
		t.Errorf("expected channelID %v, got %v", zeroID, fc.channelID)
	}
	if fc.amount.Cmp(big.NewInt(1)) != 0 {
		t.Errorf("expected amount 1, got %v", fc.amount)
	}
	if !bytes.Equal(fc.sig, []byte{0x01}) {
		t.Errorf("expected sig [0x01], got %v", fc.sig)
	}
	if !fq.called {
		t.Error("UpdateQoS was not invoked")
	}
}

func TestHandleRPCPaymentProcessingFailure(t *testing.T) {
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"jsonrpc": "2.0", "result": "0x1", "id": 1})
	}))
	defer backend.Close()
	fcErr := &fakeChannelErr{}
	fq := &fakeQoS{}
	key, _ := crypto.GenerateKey()
	srv := &RPCServer{paymentChannel: fcErr, providerReg: fq, privateKey: key, config: Config{EthRPCURL: backend.URL}}
	zeroID := [32]byte{}
	payload := map[string]interface{}{"method": "eth_blockNumber", "params": []interface{}{}, "payment": map[string]string{"channelId": hexutil.Encode(zeroID[:]), "amount": "1", "signature": hexutil.Encode([]byte{0x01})}}
	b, _ := json.Marshal(payload)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", bytes.NewBuffer(b))
	srv.handleRPCRequest(w, r)
	if w.Code != http.StatusPaymentRequired {
		t.Fatalf("expected %d, got %d", http.StatusPaymentRequired, w.Code)
	}
}

func TestHandleRPCBackendError(t *testing.T) {
	fc := &fakeChannel{}
	fq := &fakeQoS{}
	key, _ := crypto.GenerateKey()
	srv := &RPCServer{paymentChannel: fc, providerReg: fq, privateKey: key, config: Config{EthRPCURL: "http://127.0.0.1:0"}}
	zeroID := [32]byte{}
	payload := map[string]interface{}{"method": "eth_blockNumber", "params": []interface{}{}, "payment": map[string]string{"channelId": hexutil.Encode(zeroID[:]), "amount": "1", "signature": hexutil.Encode([]byte{0x01})}}
	b, _ := json.Marshal(payload)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", bytes.NewBuffer(b))
	srv.handleRPCRequest(w, r)
	if w.Code != http.StatusBadGateway {
		t.Fatalf("expected %d, got %d", http.StatusBadGateway, w.Code)
	}
}

func TestHandleRPCBadRPCResponse(t *testing.T) {
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("not json"))
	}))
	defer backend.Close()
	fc := &fakeChannel{}
	fq := &fakeQoS{}
	key, _ := crypto.GenerateKey()
	srv := &RPCServer{paymentChannel: fc, providerReg: fq, privateKey: key, config: Config{EthRPCURL: backend.URL}}
	zeroID := [32]byte{}
	payload := map[string]interface{}{"method": "eth_blockNumber", "params": []interface{}{}, "payment": map[string]string{"channelId": hexutil.Encode(zeroID[:]), "amount": "1", "signature": hexutil.Encode([]byte{0x01})}}
	b, _ := json.Marshal(payload)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", bytes.NewBuffer(b))
	srv.handleRPCRequest(w, r)
	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected %d, got %d", http.StatusInternalServerError, w.Code)
	}
}

// Test RPC returns error field: HTTP 200, error echoed, QoS updated with success=0
func TestHandleRPCWithErrorResponse(t *testing.T) {
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"jsonrpc": "2.0", "error": map[string]string{"message": "fail"}, "id": 1})
	}))
	defer backend.Close()
	fc := &fakeChannel{}
	fq := &fakeQoS{}
	key, _ := crypto.GenerateKey()
	srv := &RPCServer{paymentChannel: fc, providerReg: fq, privateKey: key, config: Config{EthRPCURL: backend.URL}}
	zeroID := [32]byte{}
	payload := map[string]interface{}{"method": "eth_blockNumber", "params": []interface{}{}, "payment": map[string]string{"channelId": hexutil.Encode(zeroID[:]), "amount": "1", "signature": hexutil.Encode([]byte{0x01})}}
	b, _ := json.Marshal(payload)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", bytes.NewBuffer(b))
	srv.handleRPCRequest(w, r)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var resp map[string]interface{}
	json.NewDecoder(w.Body).Decode(&resp)
	if resp["error"] == nil {
		t.Error("expected error field in response")
	}
	if !fq.called {
		t.Error("UpdateQoS not invoked on error response")
	}
}

// Test missing payment field yields BadRequest
func TestHandleRPCMissingPayment(t *testing.T) {
	backend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"jsonrpc": "2.0", "result": "0x1", "id": 1})
	}))
	defer backend.Close()
	fc := &fakeChannel{}
	fq := &fakeQoS{}
	key, _ := crypto.GenerateKey()
	srv := &RPCServer{paymentChannel: fc, providerReg: fq, privateKey: key, config: Config{EthRPCURL: backend.URL}}
	// Request without payment
	payload := map[string]interface{}{"method": "eth_blockNumber", "params": []interface{}{}}
	b, _ := json.Marshal(payload)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", bytes.NewBuffer(b))
	srv.handleRPCRequest(w, r)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected %d for missing payment, got %d", http.StatusBadRequest, w.Code)
	}
}

// TestHandleRPCRequest_VariousPaths covers all code paths with a table-driven approach.
func TestHandleRPCRequest_VariousPaths(t *testing.T) {
	zeroID := [32]byte{}
	goodPay := map[string]string{
		"channelId": hexutil.Encode(zeroID[:]),
		"amount":    "1",
		"signature": hexutil.Encode([]byte{0x01}),
	}

	// stub backends
	okBackend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"jsonrpc": "2.0", "result": "0x1", "id": 1})
	}))
	defer okBackend.Close()
	badJSONBackend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("not json"))
	}))
	defer badJSONBackend.Close()
	errFieldBackend := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{"jsonrpc": "2.0", "error": map[string]string{"message": "fail"}, "id": 1})
	}))
	defer errFieldBackend.Close()

	tests := []struct {
		name             string
		payload          interface{}
		backendURL       string
		channel          PaymentChannelService
		qos              QoSService
		wantStatus       int
		expectErrorField bool
		expectClose      bool
		expectQoS        bool
	}{
		{"invalid JSON", "not json", okBackend.URL, &fakeChannel{}, &fakeQoS{}, http.StatusBadRequest, false, false, false},
		{"missing payment", map[string]interface{}{"method": "x"}, okBackend.URL, &fakeChannel{}, &fakeQoS{}, http.StatusBadRequest, false, false, false},
		{"bad channelId", map[string]interface{}{"method": "x", "payment": map[string]string{"channelId": "0xzz", "amount": "1", "signature": "0x00"}}, okBackend.URL, &fakeChannel{}, &fakeQoS{}, http.StatusBadRequest, false, false, false},
		{"CloseChannel fails", map[string]interface{}{"method": "x", "payment": goodPay}, okBackend.URL, &fakeChannelErr{}, &fakeQoS{}, http.StatusPaymentRequired, false, false, false},
		{"backend unreachable", map[string]interface{}{"method": "x", "payment": goodPay}, "http://127.0.0.1:0", &fakeChannel{}, &fakeQoS{}, http.StatusBadGateway, false, false, false},
		{"backend bad JSON", map[string]interface{}{"method": "x", "payment": goodPay}, badJSONBackend.URL, &fakeChannel{}, &fakeQoS{}, http.StatusInternalServerError, false, false, false},
		{"error field", map[string]interface{}{"method": "x", "payment": goodPay}, errFieldBackend.URL, &fakeChannel{}, &fakeQoS{}, http.StatusOK, true, false, true},
		{"success", map[string]interface{}{"method": "x", "payment": goodPay}, okBackend.URL, &fakeChannel{}, &fakeQoS{}, http.StatusOK, false, true, true},
	}

	// generate a single private key for all test cases
	privKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			srv := &RPCServer{paymentChannel: tc.channel, providerReg: tc.qos, privateKey: privKey, config: Config{EthRPCURL: tc.backendURL}}
			body, _ := json.Marshal(tc.payload)
			req := httptest.NewRequest("POST", "/", bytes.NewReader(body))
			w := httptest.NewRecorder()

			srv.handleRPCRequest(w, req)
			if w.Code != tc.wantStatus {
				t.Fatalf("%s: expected status %d, got %d", tc.name, tc.wantStatus, w.Code)
			}
			if tc.expectErrorField {
				var resp map[string]interface{}
				json.NewDecoder(w.Body).Decode(&resp)
				if resp["error"] == nil {
					t.Errorf("%s: expected error field in response", tc.name)
				}
			}
			if tc.expectClose {
				fc := tc.channel.(*fakeChannel)
				if !fc.called {
					t.Errorf("%s: expected CloseChannel call", tc.name)
				}
			}
			if tc.expectQoS {
				fqo := tc.qos.(*fakeQoS)
				if !fqo.called {
					t.Errorf("%s: expected UpdateQoS call", tc.name)
				}
			}
		})
	}
}
