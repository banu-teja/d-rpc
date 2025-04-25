//go:build integration
// +build integration

package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

// doPost sends a POST to the local RPC endpoint, skips if server isn't running.
func doPost(t *testing.T, payload string) *http.Response {
	url := "http://localhost:8080/"
	resp, err := http.Post(url, "application/json", strings.NewReader(payload))
	if err != nil {
		if strings.Contains(err.Error(), "connection refused") {
			t.Skipf("Skipping integration: server not running at %s", url)
		}
		t.Fatalf("http.Post error: %v", err)
	}
	return resp
}

// TestE2E_RequestValidation covers validation errors on the live HTTP endpoint.
func TestE2E_RequestValidation(t *testing.T) {
	zeroID := "0x0000000000000000000000000000000000000000000000000000000000000000"
	tests := []struct {
		name       string
		payload    string
		wantStatus int
	}{
		{"empty body", "", http.StatusBadRequest},
		{"invalid JSON", "not json", http.StatusBadRequest},
		{"missing payment", `{"method":"eth_blockNumber","params":[]}`, http.StatusBadRequest},
		{"invalid channelId", `{"method":"x","params":[],"payment":{"channelId":"0xzz","amount":"1","signature":"0x00"}}`, http.StatusBadRequest},
		{"invalid amount", `{"method":"x","params":[],"payment":{"channelId":"` + zeroID + `","amount":"abc","signature":"0x00"}}`, http.StatusBadRequest},
		{"invalid signature", `{"method":"x","params":[],"payment":{"channelId":"` + zeroID + `","amount":"1","signature":"0xzz"}}`, http.StatusBadRequest},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp := doPost(t, tc.payload)
			defer resp.Body.Close()
			if resp.StatusCode != tc.wantStatus {
				body, _ := ioutil.ReadAll(resp.Body)
				t.Errorf("%s: expected status %d, got %d, body=%s", tc.name, tc.wantStatus, resp.StatusCode, string(body))
			}
		})
	}
}

// TestE2E_PaymentProcessing covers a valid payment payload (hex decode) and expects a 402 due to no on-chain channel closure.
func TestE2E_PaymentProcessing(t *testing.T) {
	zeroID := "0x0000000000000000000000000000000000000000000000000000000000000000"
	payload := `{"method":"eth_blockNumber","params":[],"payment":{"channelId":"` + zeroID + `","amount":"1","signature":"0x01"}}`
	resp := doPost(t, payload)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusPaymentRequired {
		body, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("expected 402 PaymentRequired, got %d, body=%s", resp.StatusCode, string(body))
	}
	body, _ := ioutil.ReadAll(resp.Body)
	if !strings.Contains(string(body), "Payment processing failed") {
		t.Errorf("expected Payment processing failed error, got body=%s", string(body))
	}
}
