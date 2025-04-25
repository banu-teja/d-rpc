package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleRPCRequest_InvalidJSON(t *testing.T) {
	srv := &RPCServer{}
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer([]byte("invalid json")))
	w := httptest.NewRecorder()
	srv.handleRPCRequest(w, req)
	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, res.StatusCode)
	}
}

func TestHandleRPCRequest_InvalidChannelID(t *testing.T) {
	payload := "{\"method\":\"eth_blockNumber\",\"params\":[],\"payment\":{\"channelId\":\"0xzz\",\"amount\":\"1\",\"signature\":\"0x00\"}}"
	srv := &RPCServer{}
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer([]byte(payload)))
	w := httptest.NewRecorder()
	srv.handleRPCRequest(w, req)
	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, res.StatusCode)
	}
}

func TestHandleRPCRequest_InvalidAmount(t *testing.T) {
	validChannelID := "0x" + strings.Repeat("a", 64)
	payload := "{\"method\":\"eth_blockNumber\",\"params\":[],\"payment\":{\"channelId\":\"" + validChannelID + "\",\"amount\":\"abc\",\"signature\":\"0x00\"}}"
	srv := &RPCServer{}
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer([]byte(payload)))
	w := httptest.NewRecorder()
	srv.handleRPCRequest(w, req)
	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, res.StatusCode)
	}
}

func TestHandleRPCRequest_InvalidSignature(t *testing.T) {
	validChannelID := "0x" + strings.Repeat("a", 64)
	payload := "{\"method\":\"eth_blockNumber\",\"params\":[],\"payment\":{\"channelId\":\"" + validChannelID + "\",\"amount\":\"1\",\"signature\":\"0xzz\"}}"
	srv := &RPCServer{}
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer([]byte(payload)))
	w := httptest.NewRecorder()
	srv.handleRPCRequest(w, req)
	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, res.StatusCode)
	}
}
