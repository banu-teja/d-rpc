package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"strconv"
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

	"github.com/go-redis/redis/v8"
)

type Config struct {
	EthRPCURL     string
	PrivateKey    string
	ContractAddrs ContractAddresses
	StakeAmount   *big.Int
	Port          string
	ChainID       *big.Int // Added: Chain ID for transactions
	RedisAddr     string   // Added: Redis address for usage tracking
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

	// Rate Limiting
	rateLimitStore map[common.Address][]time.Time
	rateLimitMu    sync.Mutex

	// Usage Tracking
	redisClient    *redis.Client // Added: Redis client for usage tracking
}

const (
	rateLimitRequests = 100 // Max requests
	rateLimitDuration = 1 * time.Minute // Per duration
)

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
	qosMonitor := qos.NewMonitor(providerReg, client, cfg.PrivateKey, cfg.ChainID)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go qosMonitor.Start(ctx)

	// Initialize load balancer
	lb, err := loadbalancer.New(cfg.ContractAddrs.ProviderRegistry, client)
	if err != nil {
		log.Fatalf("Failed to initialize load balancer: %v", err)
	}

	// Start health checker
	healthCheckInterval := 1 * time.Minute // How often to check providers
	go startHealthChecker(ctx, lb, qosMonitor, healthCheckInterval)

	// Initialize Redis Client
	var redisClient *redis.Client
	if cfg.RedisAddr != "" {
		redisClient = redis.NewClient(&redis.Options{
			Addr: cfg.RedisAddr,
		})
		// Check Redis connection
		if _, err := redisClient.Ping(ctx).Result(); err != nil {
			log.Printf("Warning: Failed to connect to Redis at %s: %v. Usage tracking disabled.", cfg.RedisAddr, err)
			redisClient = nil // Disable Redis if connection fails
		} else {
		    log.Printf("Connected to Redis at %s for usage tracking.", cfg.RedisAddr)
		}
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
		redisClient:    redisClient, // Add redis client
		rateLimitStore: make(map[common.Address][]time.Time),
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

	redisAddr := os.Getenv("REDIS_ADDR") // e.g., "localhost:6379"

	chainID := big.NewInt(31337) // Default to Anvil/Foundry local
	if chainIDStr := os.Getenv("CHAIN_ID"); chainIDStr != "" {
		var ok bool
		chainID, ok = new(big.Int).SetString(chainIDStr, 10)
		if !ok {
			log.Fatal("Invalid CHAIN_ID format")
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
		ChainID:     chainID, // Add chain ID to config
		RedisAddr:   redisAddr, // Add Redis address to config
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


// responseWriter is a wrapper around http.ResponseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	// Default status code is 200
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := newResponseWriter(w)

		// Log start of request
		log.Printf(
			"Request Start: Method=%s URI=%s RemoteAddr=%s UserAgent=%s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			r.UserAgent(),
		)

		// Serve the request
		next.ServeHTTP(rw, r)

		// Log end of request
		duration := time.Since(start)
		log.Printf(
			"Request End: Method=%s URI=%s StatusCode=%d Duration=%s",
			r.Method,
			r.RequestURI,
			rw.statusCode,
			duration,
		)
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

	// --- Client Authentication --- 
	signatureHeader := r.Header.Get("X-DRPC-Signature")
	timestampHeader := r.Header.Get("X-DRPC-Timestamp")
	if signatureHeader == "" || timestampHeader == "" {
		respondWithError(w, "Authentication headers missing (X-DRPC-Signature, X-DRPC-Timestamp)", http.StatusUnauthorized)
		return
	}

	signerAddr, err := s.verifyClientSignature(signatureHeader, timestampHeader)
	if err != nil {
		respondWithError(w, fmt.Sprintf("Authentication failed: %v", err), http.StatusUnauthorized)
		return
	}
	log.Printf("Request authenticated for client: %s", signerAddr.Hex())

	// --- Rate Limiting Check ---
	if !s.checkRateLimit(signerAddr) {
		log.Printf("Rate limit exceeded for client: %s", signerAddr.Hex())
		// Add specific header for rate limiting
		w.Header().Set("Retry-After", "60") // Suggest retrying after 60 seconds
		respondWithError(w, "Rate limit exceeded", http.StatusTooManyRequests)
		return
	}
	// ---------------------------

	// TODO: Add other authorization logic here (check balance, etc.)

	var req RPCRequest
	// Use MaxBytesReader to prevent large request bodies (adjust size as needed)
	r.Body = http.MaxBytesReader(w, r.Body, 1*1024*1024) // 1MB limit
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		if err.Error() == "http: request body too large" {
			respondWithError(w, "Request body too large", http.StatusRequestEntityTooLarge)
		} else {
			respondWithError(w, "Invalid request format", http.StatusBadRequest)
		}
		return
	}

	// ----- Remove old payment validation section -----
	// The old payment validation logic embedded within the request body 
	// is superseded by the header-based signature authentication.
	// It could potentially be integrated with this new auth later if needed.
	// -------------------------------------------------

	// ---- Core Logic Change: Use Load Balancer ----

	// Validate payment if required
	if req.Payment.ChannelID != "" {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
		defer cancel()

		err := s.validatePayment(ctx, &req)
		if err != nil {
			respondWithError(w, fmt.Sprintf("Payment validation failed: %v", err), http.StatusPaymentRequired)
			return
		}
		paymentSuccess = true // Assume payment validation passed if we got here
	}

	// ---- Core Logic Change: Use Load Balancer ----

	// 1. Select a provider using the load balancer
	selectedProvider, err := s.loadBalancer.GetProvider()
	if err != nil {
		respondWithError(w, fmt.Sprintf("Failed to select provider: %v", err), http.StatusServiceUnavailable)
		return
	}
	if selectedProvider.EndpointURL == "" {
		respondWithError(w, fmt.Sprintf("Selected provider %s has no endpoint URL configured", selectedProvider.Address.Hex()), http.StatusInternalServerError)
		return // Or maybe try selecting another provider?
	}

	log.Printf("Routing request ID %v (Method: %s) to provider %s (%s)", req.ID, req.Method, selectedProvider.Address.Hex(), selectedProvider.EndpointURL)

	// 2. Forward the request to the selected provider's endpoint
	proxyStart := time.Now()
	var resp *http.Response
	var proxyErr error

	proxyReq, err := http.NewRequestWithContext(r.Context(), "POST", selectedProvider.EndpointURL, bytes.NewBuffer([]byte(buildRPCRequest(req))))
	if err != nil {
		proxyErr = fmt.Errorf("failed to create proxy request: %w", err)
	} else {
		proxyReq.Header.Set("Content-Type", "application/json")
		// Consider adding timeout to the client
		// client := &http.Client{ Timeout: 10 * time.Second } 
		client := &http.Client{}
		resp, proxyErr = client.Do(proxyReq)
	}

	proxyDuration := time.Since(proxyStart)
	var success bool
	if proxyErr != nil || resp == nil || resp.StatusCode >= 400 {
		success = false
		log.Printf("Error proxying request ID %v to %s: Err: %v, Status Code: %d", req.ID, selectedProvider.Address.Hex(), proxyErr, resp.StatusCode)
	} else {
		success = true
	}


	// 3. Record QoS metric for the selected provider
	s.qosMonitor.RecordMetric(selectedProvider.Address, proxyDuration, success)

	// 3b. Record successful request for billing/usage tracking in Redis
	if success && s.redisClient != nil {
		// Use a key like "usage:<provider_address>"
		// Consider adding a timestamp component or using a specific key format for reward periods
		usageKey := "usage:" + selectedProvider.Address.Hex()
		// Increment the counter for this provider
		if err := s.redisClient.Incr(r.Context(), usageKey).Err(); err != nil {
			log.Printf("Error incrementing usage counter in Redis for %s: %v", selectedProvider.Address.Hex(), err)
			// Don't fail the request, just log the error
		}
	}

	// Handle proxy errors after recording metric
		respondWithError(w, fmt.Sprintf("Failed to proxy request to provider %s: %v", selectedProvider.Address.Hex(), proxyErr), http.StatusBadGateway)
		return
	}
	if resp == nil { // Should not happen if proxyErr is nil, but defensively check
		respondWithError(w, fmt.Sprintf("No response received from provider %s", selectedProvider.Address.Hex()), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// 4. Relay the response back to the client

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


// checkRateLimit checks if a client has exceeded the request limit using Redis.
// It uses a simple fixed window counter for simplicity.
A sliding window in Redis is more complex (requires sorted sets or LUA scripts).
func (s *RPCServer) checkRateLimit(clientAddr common.Address) bool {
	if s.redisClient == nil {
		log.Printf("Warning: Redis client not available, skipping rate limit check for %s", clientAddr.Hex())
		return true // Allow if Redis is down
	}

	ctx := context.Background() // Use background context for rate limiting check

	// Use a key like "ratelimit:<client_address>"
	key := "ratelimit:" + clientAddr.Hex()

	// Increment the counter for this client within the window
	count, err := s.redisClient.Incr(ctx, key).Result()
	if err != nil {
		log.Printf("Error incrementing rate limit counter in Redis for %s: %v", clientAddr.Hex(), err)
		return true // Allow if Redis error occurs
	}

	// If this is the first request in the window, set an expiration
	if count == 1 {
		if err := s.redisClient.Expire(ctx, key, rateLimitDuration).Err(); err != nil {
			log.Printf("Error setting expiration for rate limit key %s: %v", key, err)
			// Proceed even if expiration fails, might lead to longer blocking if Redis persists
		}
	}

	// Check if the count exceeds the limit
	if count > rateLimitRequests {
		// Optional: Log only once when limit is first exceeded
		// if count == rateLimitRequests+1 {
		// 	log.Printf("Rate limit exceeded for client: %s", clientAddr.Hex())
		// }
		return false // Limit exceeded
	}

	return true // Limit not exceeded
}
// verifyClientSignature checks the signature provided by the client in headers.
func (s *RPCServer) verifyClientSignature(signatureHex string, timestampStr string) (common.Address, error) {
	// 1. Decode signature
	signature, err := hexutil.Decode(signatureHex)
	if err != nil {
		return common.Address{}, fmt.Errorf("invalid signature format: %w", err)
	}
	// Ethereum signatures are 65 bytes (r, s, v), v can be 27/28 or 0/1.
	if len(signature) != 65 {
		return common.Address{}, fmt.Errorf("invalid signature length")
	}
	// Normalize V component (0/1 for EIP-155 compatibility)
	if signature[64] == 27 || signature[64] == 28 {
		signature[64] -= 27
	}

	// 2. Parse and validate timestamp
	timestampInt, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		return common.Address{}, fmt.Errorf("invalid timestamp format: %w", err)
	}
	timestamp := time.Unix(timestampInt, 0)
	currentTime := time.Now()
	allowedSkew := 60 * time.Second // Allow 60 second clock skew

	if timestamp.After(currentTime.Add(allowedSkew)) {
		return common.Address{}, fmt.Errorf("timestamp is in the future")
	}
	if timestamp.Before(currentTime.Add(-allowedSkew)) {
		return common.Address{}, fmt.Errorf("timestamp is too old (replay?) %v < %v", timestamp, currentTime.Add(-allowedSkew))
	}

	// 3. Construct the message that should have been signed
	// Using a simple fixed string with the timestamp prevents pre-signing generic messages.
	message := fmt.Sprintf("DRPC Request Timestamp: %d", timestampInt)
	// Use EIP-191 compliant hashing (standard \x19Ethereum Signed Message)
	messageHash := crypto.Keccak256Hash(
		[]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(message))),
		[]byte(message),
	)

	// 4. Recover the public key (and address) from the signature and hash
	pubKey, err := crypto.SigToPub(messageHash.Bytes(), signature)
	if err != nil {
		return common.Address{}, fmt.Errorf("signature recovery failed: %w", err)
	}
	signerAddr := crypto.PubkeyToAddress(*pubKey)

	return signerAddr, nil
}

// startHealthChecker runs a loop to periodically check provider health.
func startHealthChecker(ctx context.Context, lb *loadbalancer.LoadBalancer, qm *qos.QoSMonitor, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	log.Println("Starting provider health checker...")

	for {
		select {
		case <-ticker.C:
			providers := lb.GetAllProviders() // Get current list from load balancer
			if len(providers) == 0 {
				log.Println("Health Checker: No providers to check.")
				continue
			}
			log.Printf("Health Checker: Checking %d providers...", len(providers))
			var wg sync.WaitGroup
			for _, p := range providers {
				wg.Add(1)
				go func(provider *loadbalancer.Provider) {
					defer wg.Done()
					checkProviderHealth(ctx, provider, qm)
				}(p)
			}
			wg.Wait()
			log.Printf("Health Checker: Finished checking %d providers.", len(providers))
		case <-ctx.Done():
			log.Println("Stopping provider health checker...")
			return
		}
	}
}

// checkProviderHealth performs a single health check on a provider and records the metric.
func checkProviderHealth(ctx context.Context, provider *loadbalancer.Provider, qm *qos.QoSMonitor) {
	if provider.EndpointURL == "" {
		// Should not happen if LB filters correctly, but check anyway
		qm.RecordMetric(provider.Address, 0, false) // Record as failure if no URL
		return
	}

	// Use a short timeout specific to health checks
	checkCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Construct a simple JSON-RPC request (e.g., eth_blockNumber)
	// Using a fixed ID or method helps distinguish health checks in node logs
	healthCheckPayload := `{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":"dRPCHealthCheck"}`

	proxyStart := time.Now()
	var resp *http.Response
	var checkErr error

	proxyReq, err := http.NewRequestWithContext(checkCtx, "POST", provider.EndpointURL, bytes.NewBuffer([]byte(healthCheckPayload)))
	if err != nil {
		checkErr = fmt.Errorf("failed to create health check request: %w", err)
	} else {
		proxyReq.Header.Set("Content-Type", "application/json")
		client := &http.Client{} // Create a new client for each check for simplicity
		resp, checkErr = client.Do(proxyReq)
	}

	proxyDuration := time.Since(proxyStart)
	var success bool
	if checkErr != nil {
		// Network error, timeout, etc.
		success = false
		// log.Printf("Health Check FAIL (Network): %s - %v", provider.Address.Hex(), checkErr)
	} else if resp == nil {
		// Should not happen if checkErr is nil
		success = false
		// log.Printf("Health Check FAIL (No Response): %s", provider.Address.Hex())
	} else {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			// Check if the body is a valid JSON-RPC response (basic check)
			var rpcResp struct {
				Jsonrpc string          `json:"jsonrpc"`
				ID      interface{}     `json:"id"`
				Result  json.RawMessage `json:"result"`
				Error   *struct {
					Code    int    `json:"code"`
					Message string `json:"message"`
				} `json:"error"`
			}
			bodyBytes, readErr := io.ReadAll(resp.Body)
			if readErr == nil && json.Unmarshal(bodyBytes, &rpcResp) == nil && rpcResp.Jsonrpc == "2.0" && rpcResp.Error == nil {
				success = true // Got HTTP 200 and valid JSON-RPC response without error field
				// log.Printf("Health Check OK: %s - %v", provider.Address.Hex(), proxyDuration)
			} else {
				success = false // HTTP 200 but invalid body or RPC error
				// log.Printf("Health Check FAIL (Bad Response Body): %s - Status %d", provider.Address.Hex(), resp.StatusCode)
			}
		} else {
			success = false // Non-200 HTTP status
			// log.Printf("Health Check FAIL (Status %d): %s", resp.StatusCode, provider.Address.Hex())
		}
	}

	// Record the metric
	qm.RecordMetric(provider.Address, proxyDuration, success)
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
