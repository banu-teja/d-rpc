package qos

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/banu-teja/d-rpc/node/pkg/contracts"
)

// Metric represents a performance metric for an RPC provider
type Metric struct {
	Provider     common.Address
	ResponseTime time.Duration
	Success      bool
	Timestamp    time.Time
}

// QoSMonitor tracks and updates provider performance metrics
type QoSMonitor struct {
	registry   *contracts.ProviderRegistry
	client     *ethclient.Client
	privateKey string
	metrics    map[common.Address][]Metric
	mu         sync.RWMutex

	// Configuration
	updateInterval   time.Duration
	metricsWindow    time.Duration
	maxMetricsStored int
}

// NewMonitor creates a new QoS monitor
func NewMonitor(registry *contracts.ProviderRegistry, client *ethclient.Client, privateKey string) *QoSMonitor {
	return &QoSMonitor{
		registry:         registry,
		client:           client,
		privateKey:       privateKey,
		metrics:          make(map[common.Address][]Metric),
		updateInterval:   30 * time.Minute,
		metricsWindow:    24 * time.Hour,
		maxMetricsStored: 1000,
	}
}

// Start begins the QoS monitoring process
func (q *QoSMonitor) Start(ctx context.Context) {
	ticker := time.NewTicker(q.updateInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			q.updateScores(ctx)
		case <-ctx.Done():
			return
		}
	}
}

// RecordMetric records a performance metric for a provider
func (q *QoSMonitor) RecordMetric(provider common.Address, responseTime time.Duration, success bool) {
	metric := Metric{
		Provider:     provider,
		ResponseTime: responseTime,
		Success:      success,
		Timestamp:    time.Now(),
	}

	q.mu.Lock()
	defer q.mu.Unlock()

	// Initialize if needed
	if _, exists := q.metrics[provider]; !exists {
		q.metrics[provider] = make([]Metric, 0, q.maxMetricsStored)
	}

	// Add new metric
	q.metrics[provider] = append(q.metrics[provider], metric)

	// Trim old metrics
	if len(q.metrics[provider]) > q.maxMetricsStored {
		q.metrics[provider] = q.metrics[provider][len(q.metrics[provider])-q.maxMetricsStored:]
	}

	// Remove metrics older than the window
	cutoff := time.Now().Add(-q.metricsWindow)
	var newMetrics []Metric
	for _, m := range q.metrics[provider] {
		if m.Timestamp.After(cutoff) {
			newMetrics = append(newMetrics, m)
		}
	}
	q.metrics[provider] = newMetrics
}

// GetScoreForProvider calculates a QoS score for a specific provider
func (q *QoSMonitor) GetScoreForProvider(provider common.Address) *big.Int {
	q.mu.RLock()
	defer q.mu.RUnlock()

	metrics, exists := q.metrics[provider]
	if !exists || len(metrics) == 0 {
		return big.NewInt(0)
	}

	// Calculate the QoS score (0-100)
	var successCount, totalCount int
	var totalResponseTime time.Duration

	cutoff := time.Now().Add(-q.metricsWindow)
	for _, m := range metrics {
		if m.Timestamp.After(cutoff) {
			totalCount++
			if m.Success {
				successCount++
			}
			totalResponseTime += m.ResponseTime
		}
	}

	if totalCount == 0 {
		return big.NewInt(0)
	}

	// Calculate success rate (0-100)
	successRate := float64(successCount) / float64(totalCount) * 100

	// Calculate average response time in milliseconds
	avgResponseTime := float64(totalResponseTime.Milliseconds()) / float64(totalCount)

	// Response time score: 0-100, where lower is better
	// 100ms or less is perfect (100), 1000ms or more is terrible (0)
	responseTimeScore := 100 - min(100, max(0, avgResponseTime-100)*100/900)

	// Combined score (weighted 70% success rate, 30% response time)
	score := 0.7*successRate + 0.3*responseTimeScore

	return big.NewInt(int64(score))
}

// updateScores calculates and updates the QoS scores for all providers
func (q *QoSMonitor) updateScores(ctx context.Context) {
	q.mu.RLock()
	providers := make([]common.Address, 0, len(q.metrics))
	for provider := range q.metrics {
		providers = append(providers, provider)
	}
	q.mu.RUnlock()

	for _, provider := range providers {
		score := q.GetScoreForProvider(provider)

		// Create an authorized transactor
		auth, err := bind.NewKeyedTransactorWithChainID(
			parsePrivateKey(q.privateKey),
			big.NewInt(31337), // Replace with actual chain ID
		)
		if err != nil {
			continue // Skip this update if auth fails
		}

		// Update the QoS score on-chain
		_, err = q.registry.UpdateQoS(auth, provider, score)
		if err != nil {
			continue // Skip this update on error
		}
	}
}

// Helper function to convert string private key to ECDSA private key
func parsePrivateKey(privateKeyHex string) *ecdsa.PrivateKey {
	// Remove 0x prefix if present
	if len(privateKeyHex) > 2 && privateKeyHex[:2] == "0x" {
		privateKeyHex = privateKeyHex[2:]
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil
	}
	return privateKey
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
