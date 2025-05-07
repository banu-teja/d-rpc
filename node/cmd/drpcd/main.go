package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"bytes"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"

	"github.com/banu-teja/d-rpc/node/pkg/contracts"
	"github.com/banu-teja/d-rpc/node/pkg/loadbalancer"
	"github.com/banu-teja/d-rpc/node/pkg/qos"
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

type ProviderInfo struct {
	Stake      *big.Int
	QosScore   *big.Int
	Registered bool
}

type PaymentChannelService interface {
	CloseChannel(auth *bind.TransactOpts, channelID [32]byte, amount *big.Int, signature []byte) (*types.Transaction, error)
}

type RPCRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	ID      interface{}   `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Payment struct {
		ChannelID  string `json:"channelId"`
		Amount     string `json:"amount"`
		Signature  string `json:"signature"`
		From       string `json:"from"`
		ValidUntil int64  `json:"validUntil"`
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
	mu             sync.Mutex
	usedNonces     map[[32]byte]*big.Int
	qosMonitor     *qos.QoSMonitor
	loadBalancer   *loadbalancer.LoadBalancer
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

	// Initialize QoS monitor
	qosMonitor := qos.NewMonitor(providerReg, client, cfg.PrivateKey)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go qosMonitor.Start(ctx)

	// Initialize load balancer
	lb, err := loadbalancer.New(cfg.ContractAddrs.ProviderRegistry, client)
	if err != nil {
		log.Fatalf("Failed to initialize load balancer: %v", err)
	}

	rpcServer := &RPCServer{
		registry:       providerReg,
		paymentChannel: paymentChannel,
		providerReg:    providerReg,
		privateKey:     privateKey,
		config:         cfg,
		usedNonces:     make(map[[32]byte]*big.Int),
		qosMonitor:     qosMonitor,
		loadBalancer:   lb,
	}

	if os.Getenv("SKIP_PROVIDER_REGISTRATION") != "" {
		log.Println("Skipping provider registration")
	} else {
		if err := rpcServer.registerProvider(); err != nil {
			if strings.Contains(err.Error(), "no contract code") {
				log.Fatalf("No registry contract found at %s", rpcServer.config.ContractAddrs.ProviderRegistry.String())
			}
			log.Fatalf("Failed to register provider: %v", err)
		}
	}

	rpcServer.startHTTPServer()
}

func loadConfig() Config {
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

	stakeAmount := big.NewInt(1000000000000000000)
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
		minStake, err := s.registry.MinStake(&bind.CallOpts{})
		if err != nil {
			return err
		}
		if provider.Stake.Cmp(minStake) < 0 {
			depositAmount := new(big.Int).Sub(minStake, provider.Stake)
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
	r.Use(loggingMiddleware)
	r.Use(corsMiddleware)
	r.HandleFunc("/", s.handleRPCRequest).Methods("POST")

	// Add discovery endpoint
	r.HandleFunc("/discovery", s.handleDiscovery).Methods("GET", "OPTIONS")

	// Health check endpoint
	r.HandleFunc("/health", s.handleHealthCheck).Methods("GET", "OPTIONS")

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

// CORS middleware to handle cross-origin requests
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *RPCServer) validatePayment(ctx context.Context, req *RPCRequest) error {
	if req.Payment.ChannelID == "" || req.Payment.Amount == "" ||
		req.Payment.Signature == "" || req.Payment.From == "" ||
		req.Payment.ValidUntil == 0 {
		return fmt.Errorf("missing payment parameters")
	}

	// Check signature expiration
	if time.Now().Unix() > req.Payment.ValidUntil {
		return fmt.Errorf("signature expired")
	}

	// Decode channel ID
	chStr := req.Payment.ChannelID
	if !strings.HasPrefix(chStr, "0x") {
		chStr = "0x" + chStr
	}
	cid, err := hexutil.Decode(chStr)
	if err != nil || len(cid) != 32 {
		return fmt.Errorf("invalid channel ID")
	}
	var channelID [32]byte
	copy(channelID[:], cid)

	// Parse amount
	amount := new(big.Int)
	if _, ok := amount.SetString(req.Payment.Amount, 10); !ok {
		return fmt.Errorf("invalid amount")
	}

	// Decode signature
	sig, err := hexutil.Decode(req.Payment.Signature)
	if err != nil {
		return fmt.Errorf("invalid signature")
	}

	// Verify signature
	msg := crypto.Keccak256Hash(
		channelID[:],
		common.LeftPadBytes(amount.Bytes(), 32),
		common.LeftPadBytes(big.NewInt(req.Payment.ValidUntil).Bytes(), 32),
	)
	digest := crypto.Keccak256Hash(append([]byte("\x19Ethereum Signed Message:\n32"), msg.Bytes()...))
	pub, err := crypto.SigToPub(digest.Bytes(), sig)
	if err != nil {
		return fmt.Errorf("signature verification failed")
	}
	signer := crypto.PubkeyToAddress(*pub)
	if signer != common.HexToAddress(req.Payment.From) {
		return fmt.Errorf("signature mismatch")
	}

	// Check nonce replay protection
	s.mu.Lock()
	defer s.mu.Unlock()
	if prevAmount, exists := s.usedNonces[channelID]; exists && amount.Cmp(prevAmount) <= 0 {
		return fmt.Errorf("nonce replay detected")
	}
	s.usedNonces[channelID] = amount

	return nil
}

func (s *RPCServer) handleRPCRequest(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var req RPCRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		respondWithError(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	// Record metrics for payment validation
	var paymentSuccess bool
	defer func() {
		if req.Payment.From != "" {
			provider := common.HexToAddress(req.Payment.From)
			responseTime := time.Since(start)
			s.qosMonitor.RecordMetric(provider, responseTime, paymentSuccess)
		}
	}()

	// Validate payment if required
	if req.Payment.ChannelID != "" {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
		defer cancel()

		err := s.validatePayment(ctx, &req)
		if err != nil {
			respondWithError(w, fmt.Sprintf("Payment validation failed: %v", err), http.StatusPaymentRequired)
			return
		}
		paymentSuccess = true
	}

	// Forward the request to the blockchain node
	proxyReq, err := http.NewRequest("POST", s.config.EthRPCURL, bytes.NewBuffer([]byte(buildRPCRequest(req))))
	if err != nil {
		respondWithError(w, "Failed to create proxy request", http.StatusInternalServerError)
		return
	}
	proxyReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		respondWithError(w, "Failed to proxy request", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// Copy response headers
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Copy response status code
	w.WriteHeader(resp.StatusCode)

	// Copy response body
	var respData json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		respondWithError(w, "Failed to parse response", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(respData)
}

func buildRPCRequest(req RPCRequest) string {
	// Strip payment info from request
	cleanReq := RPCRequest{
		JSONRPC: req.JSONRPC,
		ID:      req.ID,
		Method:  req.Method,
		Params:  req.Params,
	}
	jsonData, _ := json.Marshal(cleanReq)
	return string(jsonData)
}

func respondWithError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": map[string]interface{}{
			"code":    code,
			"message": message,
		},
	})
}

// handleDiscovery serves provider discovery information
func (s *RPCServer) handleDiscovery(w http.ResponseWriter, r *http.Request) {
	providers := s.loadBalancer.GetAllProviders()

	// Convert to response format
	response := struct {
		Providers []struct {
			Address  string `json:"address"`
			QoSScore string `json:"qosScore"`
			Stake    string `json:"stake"`
		} `json:"providers"`
		RecommendedProvider string `json:"recommendedProvider"`
	}{
		Providers: make([]struct {
			Address  string `json:"address"`
			QoSScore string `json:"qosScore"`
			Stake    string `json:"stake"`
		}, 0, len(providers)),
	}

	for _, p := range providers {
		response.Providers = append(response.Providers, struct {
			Address  string `json:"address"`
			QoSScore string `json:"qosScore"`
			Stake    string `json:"stake"`
		}{
			Address:  p.Address.Hex(),
			QoSScore: p.QoSScore.String(),
			Stake:    p.Stake.String(),
		})
	}

	// Get recommended provider
	recommended, err := s.loadBalancer.GetProvider()
	if err == nil {
		response.RecommendedProvider = recommended.Address.Hex()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handleHealthCheck responds with service health status
func (s *RPCServer) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Status    string `json:"status"`
		Timestamp int64  `json:"timestamp"`
	}{
		Status:    "OK",
		Timestamp: time.Now().Unix(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
