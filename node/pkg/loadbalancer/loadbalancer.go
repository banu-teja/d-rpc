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
		updateEvery: 1 * time.Minute, // More frequent updates for testing
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
	// Try to fetch real providers from the blockchain
	// This is a simplified implementation that would be enhanced
	// in a production system with event filtering and pagination

	// Add our test providers for development
	lb.addTestProviders()

	lb.mu.Lock()
	defer lb.mu.Unlock()

	// In a production implementation, we would iterate through
	// "ProviderRegistered" events to discover all providers
	// Then query their QoS scores and other metrics

	// For now we're using the test providers added in addTestProviders

	// Simulate periodic updates to QoS scores to make UI more dynamic
	for _, provider := range lb.providers {
		// Randomly adjust QoS score slightly up or down
		adjustment := rand.Intn(5)
		if rand.Intn(2) == 0 {
			adjustment = -adjustment
		}

		newScore := new(big.Int).Add(provider.QoSScore, big.NewInt(int64(adjustment)))

		// Keep scores in a reasonable range
		if newScore.Cmp(big.NewInt(100)) > 0 {
			newScore = big.NewInt(100)
		}
		if newScore.Cmp(big.NewInt(40)) < 0 {
			newScore = big.NewInt(40)
		}

		provider.QoSScore = newScore
		provider.LastUpdate = time.Now()

		// Also update latency - simulate some variation
		latencyAdjustment := time.Duration(rand.Intn(20)) * time.Millisecond
		if rand.Intn(2) == 0 {
			provider.Latency += latencyAdjustment
		} else {
			if provider.Latency > latencyAdjustment {
				provider.Latency -= latencyAdjustment
			}
		}

		// Keep latency in reasonable range
		if provider.Latency < 30*time.Millisecond {
			provider.Latency = 30 * time.Millisecond
		}
		if provider.Latency > 500*time.Millisecond {
			provider.Latency = 500 * time.Millisecond
		}
	}

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
