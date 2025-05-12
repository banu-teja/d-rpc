package loadbalancer

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/banu-teja/d-rpc/node/pkg/contracts"
)

// Provider represents an RPC provider with its quality metrics
type Provider struct {
	Address    common.Address
	QoSScore   *big.Int
	Stake      *big.Int
	LastUpdate time.Time
	Latency    time.Duration
	UpSince    time.Time
	ConsecutiveFailures int // Track consecutive health check failures
}

// LoadBalancer manages RPC providers and selects them based on QoS
type LoadBalancer struct {
	registry    *contracts.ProviderRegistry
	client      *ethclient.Client
	providers   map[common.Address]*Provider
	mu          sync.RWMutex
	updateEvery time.Duration
	lastUpdate  time.Time
	failureThreshold int // Consecutive health check failures threshold
}

// New creates a new LoadBalancer instance
func New(registryAddr common.Address, client *ethclient.Client) (*LoadBalancer, error) {
	registry, err := contracts.NewProviderRegistry(registryAddr, client)
	if err != nil {
		return nil, err
	}

	lb := &LoadBalancer{
		registry:    registry,
		client:      client,
		providers:   make(map[common.Address]*Provider),
		updateEvery: 1 * time.Minute, // More frequent updates for testing
		failureThreshold: 3, // Default consecutive failure threshold
	}

	// Initial providers load
	if err := lb.updateProviders(context.Background()); err != nil {
		log.Printf("Warning: initial provider update failed: %v", err)
		// Continue anyway with test providers
		lb.addTestProviders()
	}

	// Start background refresh
	go lb.refreshLoop(context.Background())

	return lb, nil
}

// refreshLoop periodically updates the provider list
func (lb *LoadBalancer) refreshLoop(ctx context.Context) {
	ticker := time.NewTicker(lb.updateEvery)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := lb.updateProviders(ctx); err != nil {
				log.Printf("Error updating providers: %v", err)
				// Make sure we always have test providers if real ones fail
				lb.addTestProviders()
			}

			// Implement ping/health check system here.
			// Iterate through providers and check their RPC endpoint health and latency.
			lb.mu.Lock()
			// Note: In a real implementation, you would fetch RPC URLs for providers
			// from a reliable source (e.g., a database or the ProviderRegistry if it stored URLs).
			// Implement ping/health check system here.
			// Iterate through providers and check their RPC endpoint health and latency.
			lb.mu.Lock()
			// Note: In a real implementation, the list of providers should be fetched
			// from the ProviderRegistry contract's registered providers.
			
			// For now, iterating through the providers currently in the load balancer's map.
			for addr, provider := range lb.providers {
				
				latency, healthy := lb.checkProviderHealth(ctx, addr)
				
				if healthy {
					provider.Latency = latency
					provider.ConsecutiveFailures = 0 // Reset failures on success
					if provider.UpSince.IsZero() {
						provider.UpSince = time.Now()
					}
				} else {
					provider.ConsecutiveFailures++ // Increment failures on failure
					// Mark provider as down after a configurable threshold of failures
					if provider.ConsecutiveFailures > lb.failureThreshold {
						provider.UpSince = time.Time{} // Indicate not currently up
						log.Printf("Provider %s marked as down after %d failures", addr.Hex(), provider.ConsecutiveFailures)
					}
				}
			}
			lb.mu.Unlock()

		case <-ctx.Done():
			return
		}
	}
}

// addTestProviders adds test providers for development and testing
func (lb *LoadBalancer) addTestProviders() {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	// Add test providers
	testAddresses := []string{
		"0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266", // Anvil's first test address
		"0x70997970C51812dc3A010C7d01b50e0d17dc79C8", // Anvil's second test address
		"0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC", // Anvil's third test address
	}

	// Generate some random QoS scores and stakes
	for i, addrStr := range testAddresses {
		addr := common.HexToAddress(addrStr)

		// Gradually decreasing scores for variety
		qosScore := big.NewInt(int64(95 - i*10))
		if qosScore.Cmp(big.NewInt(40)) < 0 {
			qosScore = big.NewInt(40) // Minimum QoS score
		}

		// Different stake amounts
		stake := new(big.Int).Mul(
			big.NewInt(1+int64(i)),
			big.NewInt(1000000000000000000), // 1 ETH and up
		)

		lb.providers[addr] = &Provider{
			Address:    addr,
			QoSScore:   qosScore,
			Stake:      stake,
			LastUpdate: time.Now(),
			Latency:    time.Duration(50+i*25) * time.Millisecond,
			UpSince:    time.Now().Add(-time.Duration(24*(i+1)) * time.Hour),
		}
	}

	lb.lastUpdate = time.Now()
}

// updateProviders refreshes the list of active providers from the registry
func (lb *LoadBalancer) updateProviders(ctx context.Context) error {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	// Clear current providers
	lb.providers = make(map[common.Address]*Provider)

	// Fetch the list of registered provider addresses from the contract
	registeredAddrs, err := lb.registry.GetRegisteredProviders(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("error fetching registered providers from contract: %w", err)
	}

	// Fetch details for each registered provider
	for _, addr := range registeredAddrs {
		providerInfo, err := lb.registry.Providers(&bind.CallOpts{Context: ctx}, addr)
		if err != nil {
			log.Printf("Error fetching provider info for %s from contract: %v", addr.Hex(), err)
			continue // Skip this provider if we can't fetch details
		}

		// Only add if actually registered (double check)
		if providerInfo.Registered {
			// Check if provider already exists in map to preserve state like ConsecutiveFailures
			existingProvider, exists := lb.providers[addr]
			if exists {
				// Update existing provider info
				existingProvider.QoSScore = providerInfo.QosScore
				existingProvider.Stake = providerInfo.Stake
				existingProvider.LastUpdate = time.Now()
				log.Printf("Updated registered provider: %s", addr.Hex())
			} else {
				// Add new provider
				lb.providers[addr] = &Provider{
					Address:    addr,
					QoSScore:   providerInfo.QosScore,
					Stake:      providerInfo.Stake,
					LastUpdate: time.Now(),
					Latency:    0, // Will be updated by health check
					UpSince:    time.Time{}, // Will be updated by health check
					ConsecutiveFailures: 0,
				}
				log.Printf("Added new registered provider: %s with RPC URL: %s", addr.Hex(), providerInfo.rpcUrl)
			}
		}
	}

	// Note: This implementation does not handle providers that have deregistered
	// since the last update. A more robust solution would compare the list from the
	// contract with the current providers in the map and remove deregistered ones.

	lb.lastUpdate = time.Now()
	return nil
}

// AddProvider manually adds a provider to the balancer
func (lb *LoadBalancer) AddProvider(addr common.Address) error {
	provider, err := lb.registry.Providers(&bind.CallOpts{}, addr)
	if err != nil {
		return err
	}

	if !provider.Registered {
		return errors.New("provider not registered")
	}

	lb.mu.Lock()
	defer lb.mu.Unlock()

	lb.providers[addr] = &Provider{
		Address:    addr,
		QoSScore:   provider.QosScore,
		Stake:      provider.Stake,
		LastUpdate: time.Now(),
		Latency:    100 * time.Millisecond, // Default latency
		UpSince:    time.Now(),
	}

	return nil
}

// checkProviderHealth performs a health check on a provider's RPC endpoint
func (lb *LoadBalancer) checkProviderHealth(ctx context.Context, providerAddr common.Address) (time.Duration, bool) {
	// Fetch the provider's RPC URL from the registry
	providerInfo, err := lb.registry.Providers(&bind.CallOpts{Context: ctx}, providerAddr)
	if err != nil || !providerInfo.Registered {
		log.Printf("Could not fetch registered provider info for %s or provider not registered: %v", providerAddr.Hex(), err)
		return 0, false
	}

	rpcURL := providerInfo.rpcUrl
	if rpcURL == "" {
		log.Printf("RPC URL is empty for provider %s", providerAddr.Hex())
		return 0, false
	}

	start := time.Now()
	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		log.Printf("Health check failed for %s (%s): %v", providerAddr.Hex(), rpcURL, err)
		return 0, false
	}
	defer client.Close()

	// Perform a simple RPC call, e.g., get the latest block number
	_, err = client.BlockNumber(ctx)
	if err != nil {
		log.Printf("Health check failed for %s (%s): %v", providerAddr.Hex(), rpcURL, err)
		return 0, false
	}

	latency := time.Since(start)
	return latency, true
}

// GetProvider selects a provider using weighted random selection based on QoS
func (lb *LoadBalancer) GetProvider() (*Provider, error) {
	lb.mu.RLock()
	defer lb.mu.RUnlock()

	if len(lb.providers) == 0 {
		return nil, errors.New("no providers available")
	}

	// Convert map to slice for sorting
	providersList := make([]*Provider, 0, len(lb.providers))
	for _, p := range lb.providers {
		providersList = append(providersList, p)
	}

	// Sort by QoS score (highest first)
	sort.Slice(providersList, func(i, j int) bool {
		return providersList[i].QoSScore.Cmp(providersList[j].QoSScore) > 0
	})

	// Weighted selection based on QoS score
	totalWeight := big.NewInt(0)
	weights := make([]*big.Int, len(providersList))

	for i, p := range providersList {
		// Weight = QoS score ^ 2 to give higher preference to high-quality providers
		weight := new(big.Int).Mul(p.QoSScore, p.QoSScore)
		weights[i] = weight
		totalWeight = new(big.Int).Add(totalWeight, weight)
	}

	// Get a random number between 0 and totalWeight
	if totalWeight.Cmp(big.NewInt(0)) <= 0 {
		// Fallback to simple random selection if weights are invalid
		return providersList[rand.Intn(len(providersList))], nil
	}

	// Generate a random number between 0 and totalWeight-1
	randomValue := new(big.Int).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), totalWeight)

	// Find the provider whose weight range contains the random value
	cumulativeWeight := big.NewInt(0)
	for i, weight := range weights {
		cumulativeWeight = new(big.Int).Add(cumulativeWeight, weight)
		if randomValue.Cmp(cumulativeWeight) < 0 {
			return providersList[i], nil
		}
	}

	// Fallback to the first provider if something goes wrong with the weighted selection
	return providersList[0], nil
}

// GetAllProviders returns a list of all active providers
func (lb *LoadBalancer) GetAllProviders() []*Provider {
	lb.mu.RLock()
	defer lb.mu.RUnlock()

	providers := make([]*Provider, 0, len(lb.providers))
	for _, p := range lb.providers {
		providers = append(providers, p)
	}

	return providers
}
