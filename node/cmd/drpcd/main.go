package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"bytes"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"

	"github.com/banu-teja/d-rpc/node/pkg/contracts"
)

type Config struct {
	EthRPCURL     string
	PrivateKey    string
	ContractAddrs ContractAddresses
	StakeAmount   *big.Int
	Port          string
}

type ContractAddresses struct {
	PaymentChannel   common.Address
	ProviderRegistry common.Address
	StakeToken       common.Address
}

// ProviderInfo represents provider state in the registry.
type ProviderInfo struct {
	Stake      *big.Int
	QosScore   *big.Int
	Registered bool
}

// PaymentChannelService defines methods for payment channel interactions.
type PaymentChannelService interface {
	CloseChannel(auth *bind.TransactOpts, channelID [32]byte, amount *big.Int, signature []byte) (*types.Transaction, error)
}

// QoSService defines methods for provider QoS updates.
type RPCRequest struct {
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Payment struct {
		ChannelID string `json:"channelId"`
		Amount    string `json:"amount"`
		Signature string `json:"signature"`
	} `json:"payment"`
}

type QoSService interface {
	UpdateQoS(auth *bind.TransactOpts, provider common.Address, score *big.Int) (*types.Transaction, error)
}

type RPCServer struct {
	registry       *contracts.ProviderRegistry
	paymentChannel PaymentChannelService
	providerReg    QoSService
	privateKey     *ecdsa.PrivateKey
	config         Config
}

func main() {
	cfg := loadConfig()

	client, err := ethclient.Dial(cfg.EthRPCURL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	paymentChannel, err := contracts.NewPaymentChannel(cfg.ContractAddrs.PaymentChannel, client)
	if err != nil {
		log.Fatalf("Failed to instantiate PaymentChannel: %v", err)
	}

	providerReg, err := contracts.NewProviderRegistry(cfg.ContractAddrs.ProviderRegistry, client)
	if err != nil {
		log.Fatalf("Failed to instantiate ProviderRegistry: %v", err)
	}

	rpcServer := &RPCServer{
		registry:       providerReg,
		paymentChannel: paymentChannel,
		providerReg:    providerReg,
		privateKey:     privateKey,
		config:         cfg,
	}

	// attempt on-chain provider registration (skip if env var set)
	if os.Getenv("SKIP_PROVIDER_REGISTRATION") != "" {
		log.Println("Skipping provider registration")
	} else {
		if err := rpcServer.registerProvider(); err != nil {
			if strings.Contains(err.Error(), "no contract code") {
				log.Fatalf("No registry contract found at %s. Please deploy contracts using .commands/.forge.zsh (e.g., 'source .commands/.forge.zsh && deploy_contracts') and set REGISTRY_CONTRACT accordingly.", rpcServer.config.ContractAddrs.ProviderRegistry.String())
			}
			log.Fatalf("Failed to register provider: %v", err)
		}
	}

	rpcServer.startHTTPServer()
}

func loadConfig() Config {
	// Load from environment variables
	ethURL := os.Getenv("RPC_URL")
	if ethURL == "" {
		ethURL = "http://localhost:8545"
	}

	privateKey := strings.TrimPrefix(os.Getenv("FORGE_PRIVATE_KEY"), "0x")
	if privateKey == "" {
		log.Fatal("FORGE_PRIVATE_KEY environment variable required")
	}

	paymentAddr := os.Getenv("CHANNEL_CONTRACT")
	registryAddr := os.Getenv("REGISTRY_CONTRACT")
	stakeAddr := os.Getenv("STK_CONTRACT")
	if paymentAddr == "" || registryAddr == "" || stakeAddr == "" {
		log.Fatal("Contract addresses (CHANNEL_CONTRACT, REGISTRY_CONTRACT, STK_CONTRACT) required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	stakeAmount := big.NewInt(1000000000000000000) // Default 1 token
	if stakeStr := os.Getenv("STAKE_AMOUNT"); stakeStr != "" {
		var ok bool
		stakeAmount, ok = new(big.Int).SetString(stakeStr, 10)
		if !ok {
			log.Fatal("Invalid STAKE_AMOUNT format")
		}
	}

	return Config{
		EthRPCURL:  ethURL,
		PrivateKey: privateKey,
		ContractAddrs: ContractAddresses{
			PaymentChannel:   common.HexToAddress(paymentAddr),
			ProviderRegistry: common.HexToAddress(registryAddr),
			StakeToken:       common.HexToAddress(stakeAddr),
		},
		StakeAmount: stakeAmount,
		Port:        port,
	}
}

func (s *RPCServer) registerProvider() error {
	auth, err := bind.NewKeyedTransactorWithChainID(s.privateKey, big.NewInt(31337))
	if err != nil {
		return err
	}

	provider, err := s.registry.Providers(&bind.CallOpts{}, auth.From)
	if err != nil {
		return err
	}

	if !provider.Registered {
		if provider.Stake.Cmp(s.config.StakeAmount) < 0 {
			depositAmount := new(big.Int).Sub(s.config.StakeAmount, provider.Stake)
			_, err = s.registry.DepositStake(auth, depositAmount)
			if err != nil {
				return err
			}
		}
		_, err = s.registry.Register(auth)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *RPCServer) startHTTPServer() {
	r := mux.NewRouter()
	// Attach logging middleware
	r.Use(loggingMiddleware)
	r.HandleFunc("/", s.handleRPCRequest).Methods("POST")

	srv := &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: r,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	log.Printf("Server started on port %s", s.config.Port)
	<-done
	log.Println("Server stopped")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Incoming request: %s %s from %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func (s *RPCServer) validatePayment(ctx context.Context, req *RPCRequest) error {
	// Implement payment validation logic here
	// 1. Verify signature matches channel ID and amount
	// 2. Check channel has sufficient balance
	// 3. Verify non-reuse of payment authorization
	return nil
}

func (s *RPCServer) handleRPCRequest(w http.ResponseWriter, r *http.Request) {
	log.Printf("handleRPCRequest called: %s %s", r.Method, r.URL.Path)
	start := time.Now()
	var req RPCRequest

	// Decode request payload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Invalid request body: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	log.Printf("Parsed RPCRequest: method=%s, params=%v, payment=%+v", req.Method, req.Params, req.Payment)

	// Payment parameter validation
	if err := s.validatePayment(r.Context(), &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Decode payment channel ID
	// Auto-prefix hex if missing
	chStr := req.Payment.ChannelID
	if !strings.HasPrefix(chStr, "0x") {
		chStr = "0x" + chStr
	}
	cidBytes, err := hexutil.Decode(chStr)
	if err != nil {
		http.Error(w, "Invalid channel ID", http.StatusBadRequest)
		return
	}
	var channelID [32]byte
	copy(channelID[:], cidBytes)
	// Parse amount
	amount := new(big.Int)
	amount, ok := amount.SetString(req.Payment.Amount, 10)
	if !ok {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}
	// Decode signature
	sigBytes, err := hexutil.Decode(req.Payment.Signature)
	if err != nil {
		http.Error(w, "Invalid signature", http.StatusBadRequest)
		return
	}
	// Close payment channel on-chain
	auth, err := bind.NewKeyedTransactorWithChainID(s.privateKey, big.NewInt(31337))
	if err != nil {
		http.Error(w, "Payment verification failed", http.StatusInternalServerError)
		return
	}
	_, err = s.paymentChannel.CloseChannel(auth, channelID, amount, sigBytes)
	if err != nil {
		http.Error(w, "Payment processing failed", http.StatusPaymentRequired)
		return
	}
	// Forward RPC request to Ethereum node
	backendReq := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  req.Method,
		"params":  req.Params,
		"id":      1,
	}
	reqBody, err := json.Marshal(backendReq)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	httpResp, err := http.Post(s.config.EthRPCURL, "application/json", bytes.NewReader(reqBody))
	if err != nil {
		http.Error(w, "Backend RPC error", http.StatusBadGateway)
		return
	}
	defer httpResp.Body.Close()
	var rpcResp map[string]interface{}
	if err := json.NewDecoder(httpResp.Body).Decode(&rpcResp); err != nil {
		http.Error(w, "Invalid RPC response", http.StatusInternalServerError)
		return
	}
	// Update QoS metrics with comprehensive measurements
	latency := time.Since(start).Milliseconds()
	success := 1 // Assume success since we reached here
	if rpcResp["error"] != nil {
		success = 0
	}

	qosAuth, err := bind.NewKeyedTransactorWithChainID(s.privateKey, big.NewInt(31337))
	if err == nil && s.providerReg != nil {
		// Calculate composite QoS score (latency * success)
		qosScore := big.NewInt(latency)
		qosScore.Mul(qosScore, big.NewInt(int64(success)))
		_, _ = s.providerReg.UpdateQoS(qosAuth, auth.From, qosScore)
	}
	// Return RPC response to client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rpcResp)
}
