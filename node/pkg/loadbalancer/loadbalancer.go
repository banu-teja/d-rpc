package loadbalancer

import (
	"context"
	"errors"
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
	EndpointURL string // Added: RPC Endpoint URL
}

// LoadBalancer manages RPC providers and selects them based on QoS
type LoadBalancer struct {
	registry    *contracts.ProviderRegistry
	client      *ethclient.Client
	providers   map[common.Address]*Provider
	mu          sync.RWMutex
	updateEvery time.Duration
	lastUpdate  time.Time
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
		updateEvery: 5 * time.Minute, // Check for updates less frequently now
	}

	// Initial providers load
	// Use a separate context for initial load, might take longer
	ctxInitial, cancelInitial := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelInitial()
	if err := lb.updateProviders(ctxInitial); err != nil {
		log.Printf("Warning: initial provider update failed: %v. Falling back to test providers.", err)
		// Fallback to test providers if initial load fails
		lb.addTestProviders()
	} else if len(lb.providers) == 0 {
		log.Printf("Warning: No registered providers found. Adding test providers.")
		lb.addTestProviders() // Also add test providers if none found
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
			// Add placeholder URL for test providers, assuming they run on sequential ports
			EndpointURL: fmt.Sprintf("http://localhost:854%d", 5+i), 
		}
	}

	lb.lastUpdate = time.Now()
}

// updateProviders refreshes the list of active providers from the registry.
// NOTE: This implementation re-queries known providers. A more scalable approach
// would use event watching (e.g., WatchProviderRegistered) to incrementally update.
func (lb *LoadBalancer) updateProviders(ctx context.Context) error {
	log.Println("Updating provider list...")
	lb.mu.Lock()
	defer lb.mu.Unlock()

	// Create a temporary list of current provider addresses
	currentAddrs := make([]common.Address, 0, len(lb.providers))
	for addr := range lb.providers {
		currentAddrs = append(currentAddrs, addr)
	}

	// In a real system, we'd need a discovery mechanism for *new* providers.
	// For now, we just refresh the state of known providers.
	// A simple discovery could involve querying logs for ProviderRegistered events
	// within a recent block range, but this is omitted for brevity here.

	// Refresh state for existing providers
	updatedProviders := make(map[common.Address]*Provider)
	for _, addr := range currentAddrs {
		providerData, err := lb.registry.Providers(&bind.CallOpts{Context: ctx}, addr)
		if err != nil {
			// If provider data can't be fetched, assume it's deregistered or contract issue
			log.Printf("Error fetching provider data for %s: %v. Removing from active list.", addr.Hex(), err)
			continue // Skip this provider
		}

		// Check if provider is still registered
		if !providerData.Registered {
			log.Printf("Provider %s is no longer registered. Removing from active list.", addr.Hex())
			continue // Skip unregistered provider
		}
		if providerData.EndpointURL == "" {
			log.Printf("Provider %s has no endpoint URL set. Skipping.", addr.Hex())
			continue // Skip provider without URL
		}

		// Update or keep existing provider entry
		existingProvider, exists := lb.providers[addr]
		if exists {
			// Update fields from contract
			existingProvider.QoSScore = providerData.QosScore
			existingProvider.Stake = providerData.Stake
			existingProvider.EndpointURL = providerData.EndpointURL
			existingProvider.LastUpdate = time.Now() // Mark as updated
			// Keep existing Latency/UpSince unless we implement active health checks
			updatedProviders[addr] = existingProvider
		} else {
			// This case shouldn't be hit with the current logic, but handles discovery if added
			updatedProviders[addr] = &Provider{
				Address:    addr,
				QoSScore:   providerData.QosScore,
				Stake:      providerData.Stake,
				LastUpdate: time.Now(),
				Latency:    100 * time.Millisecond, // Default latency for new provider
				UpSince:    time.Now(),           // Assume up since discovery
				EndpointURL: providerData.EndpointURL,
			}
		}
	}

	// Replace the old map with the updated one
	lb.providers = updatedProviders
	lb.lastUpdate = time.Now()

	log.Printf("Provider list update complete. Active providers: %d", len(lb.providers))

	// TODO: Implement event watching for ProviderRegistered, ProviderDeregistered, 
	//       ProviderURLUpdated, QoSUpdated for more efficient updates.
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
			EndpointURL: provider.EndpointURL, // Store the fetched URL (Requires updated Go bindings)
		}

	return nil
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
