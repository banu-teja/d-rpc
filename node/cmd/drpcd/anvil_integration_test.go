//go:build integration
// +build integration

package main

import (
	"encoding/hex"
	"encoding/json"
	"math/big"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

const (
	anvilURL  = "http://localhost:8545"
	serverURL = "http://localhost:8080"
	// Using account[0] from Anvil

	testPrivKey = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	testAddress = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
)

func requireAnvil(t *testing.T) *ethclient.Client {
	client, err := ethclient.Dial(anvilURL)
	if err != nil {
		t.Skipf("Skipping anvil integration: %v", err)
	}
	return client
}

func TestAnvil_HealthCheck(t *testing.T) {
	resp, err := http.Get(serverURL + "/")
	require.NoError(t, err)
	// GET should not be allowed (404 or 405)
	require.NotEqual(t, http.StatusOK, resp.StatusCode)
}

func TestAnvil_PaymentRequiredWithoutChannel(t *testing.T) {
	_ = requireAnvil(t)
	// sign voucher for non-existent channel (zeroID)
	key, err := crypto.HexToECDSA(testPrivKey)
	require.NoError(t, err)
	dummyID := common.Hash{}
	amount := big.NewInt(1000000000000000000) // 1 ETH in wei
	// hash = keccak256(channelId || amount)
	buf := append(dummyID.Bytes(), amount.Bytes()...)
	hash := crypto.Keccak256Hash(buf)
	sig, err := crypto.Sign(hash.Bytes(), key)
	require.NoError(t, err)

	payload := map[string]interface{}{
		"method": "eth_blockNumber",
		"params": []interface{}{},
		"payment": map[string]string{
			"channelId": hex.EncodeToString(dummyID.Bytes()),
			"amount":    amount.String(),
			"signature": "0x" + hex.EncodeToString(sig),
			"from":      testAddress,
		},
	}
	body, _ := json.Marshal(payload)
	resp, err := http.Post(serverURL+"/", "application/json", strings.NewReader(string(body)))
	require.NoError(t, err)
	require.Equal(t, http.StatusPaymentRequired, resp.StatusCode)
}

func TestAnvil_ForwardRPC_InvalidJSON(t *testing.T) {
	_ = requireAnvil(t)
	resp, err := http.Post(serverURL+"/", "application/json", strings.NewReader("not json"))
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestAnvil_ForwardRPC_ValidRequest(t *testing.T) {
	requireAnvil(t)
	// dummy payment fields to reach forwarding
	payload := `{"method":"eth_blockNumber","params":[],"payment":{"channelId":"0x` + strings.Repeat("0", 64) + `","amount":"0","signature":"0x"}}`
	// allow some processing time
	time.Sleep(2 * time.Second)
	resp, err := http.Post(serverURL+"/", "application/json", strings.NewReader(payload))
	require.NoError(t, err)
	// accept either forwarded error or payment required
	require.Contains(t, []int{http.StatusOK, http.StatusPaymentRequired}, resp.StatusCode)
}
