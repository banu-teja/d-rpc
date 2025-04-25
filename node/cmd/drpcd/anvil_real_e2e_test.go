//go:build integration
// +build integration

package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

const (
	// projectRoot is relative path from this test to repo root
	projectRoot   = "../../.."
	realAnvilURL  = "http://localhost:8545"
	realServerURL = "http://localhost:8080"
	// Provider (RPC server) account: Anvil account[0]
	providerPrivKey = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	providerAddress = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	// User account: Anvil account[1]
	userPrivKey = "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	userAddress = "0x70997970C51812dc3A010C7d01b50e0d17dc79C8"
)

// setupContracts runs deployment scripts and opens a payment channel. Returns channel ID.
func setupContracts(t *testing.T) string {
	t.Helper()
	wd, err := os.Getwd()
	require.NoError(t, err)
	root := filepath.Join(wd, projectRoot)
	contractsDir := filepath.Join(root, "contracts")

	// 1. Deploy contracts using provider key
	deployCmd := exec.Command("bash", "-lc", fmt.Sprintf(
		"forge script script/DeployContracts.s.sol --broadcast --rpc-url %s --private-key %s",
		realAnvilURL, providerPrivKey))
	deployCmd.Dir = contractsDir
	deployOut, err := deployCmd.CombinedOutput()
	require.NoError(t, err, "deploy contracts failed: %s", string(deployOut))

	// Parse deployed addresses
	var stakeTokenAddr, registryAddr, channelContractAddr string
	for _, line := range strings.Split(string(deployOut), "\n") {
		if strings.Contains(line, "Deployed StakeToken at:") {
			parts := strings.Split(line, ":")
			stakeTokenAddr = strings.TrimSpace(parts[len(parts)-1])
		} else if strings.Contains(line, "Deployed ProviderRegistry at:") {
			parts := strings.Split(line, ":")
			registryAddr = strings.TrimSpace(parts[len(parts)-1])
		} else if strings.Contains(line, "Deployed PaymentChannel at:") {
			parts := strings.Split(line, ":")
			channelContractAddr = strings.TrimSpace(parts[len(parts)-1])
		}
	}
	require.NotEmpty(t, stakeTokenAddr)
	require.NotEmpty(t, registryAddr)
	require.NotEmpty(t, channelContractAddr)

	// Set environment for RPC server
	os.Setenv("RPC_URL", realAnvilURL)
	os.Setenv("FORGE_PRIVATE_KEY", providerPrivKey)
	os.Setenv("STK_CONTRACT", stakeTokenAddr)
	os.Setenv("REGISTRY_CONTRACT", registryAddr)
	os.Setenv("CHANNEL_CONTRACT", channelContractAddr)
	// Skip on-chain provider registration in RPC server
	os.Setenv("SKIP_PROVIDER_REGISTRATION", "true")
	os.Setenv("PORT", "8081")
	os.Setenv("STAKE_AMOUNT", "100000000000000000000") // 100 STK

	// 2. Fund user (use provider key)
	fundCmdStr := fmt.Sprintf(
		"STK_CONTRACT=%s USER_ADDRESS=%s CHANNEL_DEPOSIT=%s forge script script/FundUser.s.sol --broadcast --rpc-url %s --private-key %s",
		stakeTokenAddr, userAddress, "10000000000000000000", realAnvilURL, providerPrivKey)
	fundCmd := exec.Command("bash", "-lc", fundCmdStr)
	fundCmd.Dir = contractsDir
	fundOut, err := fundCmd.CombinedOutput()
	require.NoError(t, err, "fund user failed: %s", string(fundOut))

	// 3. Open channel (user funds channel to provider)
	openCmdStr := fmt.Sprintf(
		"CHANNEL_CONTRACT=%s PROVIDER_ADDRESS=%s forge script script/OpenChannel.s.sol --broadcast --rpc-url %s --private-key %s",
		channelContractAddr, providerAddress, realAnvilURL, userPrivKey)
	openCmd := exec.Command("bash", "-lc", openCmdStr)
	openCmd.Dir = contractsDir
	openOut, err := openCmd.CombinedOutput()
	require.NoError(t, err, "open channel failed: %s", string(openOut))

	// 4. Extract channel ID from logs
	jsonPath := filepath.Join(contractsDir, "broadcast", "OpenChannel.s.sol", "31337", "run-latest.json")
	data, err := os.ReadFile(jsonPath)
	require.NoError(t, err, "read OpenChannel logs: %v", err)
	type logEntry struct {
		Topics []string `json:"topics"`
	}
	type receipt struct {
		Logs []logEntry `json:"logs"`
	}
	type result struct {
		Receipts []receipt `json:"receipts"`
	}
	var res result
	require.NoError(t, json.Unmarshal(data, &res))
	var channelID string
	const eventSig = "0x506f81b7a67b45bfbc6167fd087b3dd9b65b4531a2380ec406aab5b57ac62152"
	logs := res.Receipts[len(res.Receipts)-1].Logs
	for _, log := range logs {
		if log.Topics[0] == eventSig && len(log.Topics) > 1 {
			channelID = log.Topics[1]
			break
		}
	}
	require.NotEmpty(t, channelID, "channel ID not found in logs")
	return channelID
}

func TestRealE2E_PaymentFlow(t *testing.T) {
	// ensure Anvil is running
	_, err := http.Get(realAnvilURL)
	require.NoError(t, err, "anvil not running at %s", realAnvilURL)

	// setup contracts and open channel
	channelID := setupContracts(t)

	// start RPC server
	go main()
	// allow server to listen
	time.Sleep(2 * time.Second)

	// derive server URL from PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	serverURL := fmt.Sprintf("http://localhost:%s", port)

	// create payment voucher
	amount := big.NewInt(1000000000000000) // 0.001 token units
	chBytes, err := hex.DecodeString(strings.TrimPrefix(channelID, "0x"))
	require.NoError(t, err)
	// Compute message hash and Ethereum signed message digest
	// Pad amount to 32 bytes for abi.encodePacked(uint256)
	paddedAmt := common.LeftPadBytes(amount.Bytes(), 32)
	messageHash := crypto.Keccak256Hash(append(chBytes, paddedAmt...))
	digest := crypto.Keccak256Hash(append([]byte("\x19Ethereum Signed Message:\n32"), messageHash.Bytes()...))
	key, err := crypto.HexToECDSA(strings.TrimPrefix(userPrivKey, "0x"))
	require.NoError(t, err)
	sig, err := crypto.Sign(digest.Bytes(), key)
	require.NoError(t, err)
	// Adjust recovery ID (v) to 27/28 for ECDSA
	sig[64] += 27

	// send RPC request
	payload := map[string]interface{}{
		"method": "eth_blockNumber",
		"params": []interface{}{},
		"payment": map[string]string{
			"channelId": channelID,
			"amount":    amount.String(),
			"signature": "0x" + hex.EncodeToString(sig),
			"from":      userAddress,
		},
	}
	body, err := json.Marshal(payload)
	require.NoError(t, err)
	resp, err := http.Post(serverURL+"/", "application/json", bytes.NewReader(body))
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
}
