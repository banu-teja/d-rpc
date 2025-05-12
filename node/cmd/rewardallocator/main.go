package main


import (
	"context"
	"fmt" // Added for Sscan and error formatting
	"strings" // Added for ABI parsing

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types" // Added for receipt status check
	ethereum "github.com/ethereum/go-ethereum"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v8"

	"github.com/banu-teja/d-rpc/node/pkg/contracts"
	// We might need to fetch all providers efficiently. 
	// Currently loadbalancer doesn't expose a way to get *all* registered providers easily.
	// We might need to add a function to ProviderRegistry contract or query events.
	// For now, we'll assume a way to get the list.
)

var (
	// Keccak256 hash of "ProviderRegistered(address,uint256)"
	providerRegisteredEventSig = crypto.Keccak256Hash([]byte("ProviderRegistered(address,uint256)"))
)

type Config struct {
	EthRPCURL        string
	PrivateKey       string // Owner's private key
	RegistryAddr     common.Address
	RedisAddr        string
	ChainID          *big.Int
	TotalReward      *big.Int // Total reward amount per period
	MinUsage         int64    // Minimum usage count to be eligible
	RewardPeriod     time.Duration
	ResetCounters    bool     // Whether to reset Redis counters after allocation
}

func main() {
	cfg := loadConfig()
	ctx := context.Background()

	// Connect to Ethereum client
	client, err := ethclient.Dial(cfg.EthRPCURL)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum client: %v", err)
	}

	// Connect to Redis client
	var redisClient *redis.Client
	if cfg.RedisAddr != "" {
		redisClient = redis.NewClient(&redis.Options{Addr: cfg.RedisAddr})
		if _, err := redisClient.Ping(ctx).Result(); err != nil {
			log.Fatalf("Failed to connect to Redis: %v", err)
		}
		log.Printf("Connected to Redis at %s", cfg.RedisAddr)
	} else {
		log.Fatalf("Redis address is required for usage tracking.")
	}

	// Instantiate ProviderRegistry contract
	registry, err := contracts.NewProviderRegistry(cfg.RegistryAddr, client)
	if err != nil {
		log.Fatalf("Failed to instantiate ProviderRegistry: %v", err)
	}

	log.Println("Fetching registered providers via event logs...")

	// --- Fetching Provider List via Events ---
	query := ethereum.FilterQuery{
		Addresses: []common.Address{cfg.RegistryAddr},
		FromBlock: big.NewInt(0), // Query from the beginning
		ToBlock:   nil,           // Query up to the latest block
		Topics:    [][]common.Hash{{providerRegisteredEventSig}},
	}

	logs, err := client.FilterLogs(ctx, query)
	if err != nil {
		log.Fatalf("Failed to filter ProviderRegistered logs: %v", err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(contracts.ProviderRegistryMetaData.ABI))
	if err != nil {
		log.Fatalf("Failed to parse ProviderRegistry ABI: %v", err)
	}

	// Use a map to automatically handle duplicates if a provider registers multiple times
	providerMap := make(map[common.Address]bool) 
	for _, vLog := range logs {
		if len(vLog.Topics) > 1 { // Ensure the indexed provider topic is present
			// The first topic (Topics[0]) is the event signature.
			// For indexed event parameters, subsequent topics hold their values.
			// ProviderRegistered(address indexed provider, ...)
			// So, Topics[1] will be the provider's address.
			providerAddr := common.BytesToAddress(vLog.Topics[1].Bytes())
			providerMap[providerAddr] = true
		} else {
			log.Printf("Warning: Found ProviderRegistered log with insufficient topics: Tx %s", vLog.TxHash.Hex())
		}
	}

	// Convert map keys to slice
	providerList := make([]common.Address, 0, len(providerMap))
	for addr := range providerMap {
		providerList = append(providerList, addr)
	}
	// --- End Fetching Provider List ---

	log.Printf("Found %d unique provider addresses from event logs.", len(providerList))

	eligibleProviders := make([]common.Address, 0)
	usageCounts := make(map[common.Address]int64)

	log.Println("Checking eligibility and usage...")
	for _, addr := range providerList {
		// Check if provider is still registered (important!)
		providerData, err := registry.Providers(&bind.CallOpts{Context: ctx}, addr)
		if err != nil {
			log.Printf("Error fetching data for provider %s: %v. Skipping.", addr.Hex(), err)
			continue
		}
		if !providerData.Registered {
			log.Printf("Provider %s is not registered. Skipping.", addr.Hex())
			continue
		}

		// Get usage count from Redis
		usageKey := "usage:" + addr.Hex()
		usageStr, err := redisClient.Get(ctx, usageKey).Result()
		var usage int64 = 0
		if err == nil {
			// Parse usage string to int64
			_, scanErr := Sscan(usageStr, &usage) // Simple string to int conversion
			if scanErr != nil {
				log.Printf("Error parsing usage count '%s' for provider %s: %v. Assuming 0.", usageStr, addr.Hex(), scanErr)
				usage = 0
			}
		} else if err != redis.Nil {
			// Log error only if it's not "key not found"
			log.Printf("Error getting usage from Redis for provider %s: %v. Assuming 0.", addr.Hex(), err)
		}

		log.Printf("Provider: %s, Registered: %t, QoS: %s, Stake: %s, Usage: %d", addr.Hex(), providerData.Registered, providerData.QosScore.String(), providerData.Stake.String(), usage)

		// Check eligibility (e.g., minimum usage)
		if usage >= cfg.MinUsage {
			eligibleProviders = append(eligibleProviders, addr)
			usageCounts[addr] = usage // Store usage for potential later use (e.g. logging)
		} else {
			log.Printf("Provider %s is ineligible (Usage: %d < MinUsage: %d)", addr.Hex(), usage, cfg.MinUsage)
		}
	}

	if len(eligibleProviders) == 0 {
		log.Println("No eligible providers found for reward allocation.")
		return
	}

	log.Printf("Allocating total reward of %s STK to %d eligible providers.", cfg.TotalReward.String(), len(eligibleProviders))

	// Prepare transaction options
	ownerPrivateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		log.Fatalf("Failed to parse owner private key: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(ownerPrivateKey, cfg.ChainID)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	// Call the allocation function
	tx, err := registry.CalculateAndAllocateRewards(auth, eligibleProviders, cfg.TotalReward)
	if err != nil {
		log.Fatalf("Failed to send CalculateAndAllocateRewards transaction: %v", err)
	}

	log.Printf("Reward allocation transaction sent: %s", tx.Hash().Hex())
	log.Println("Waiting for transaction receipt...")

	receipt, err := bind.WaitMined(ctx, client, tx)
	if err != nil {
		log.Fatalf("Error waiting for transaction receipt: %v", err)
	}

	if receipt.Status == types.ReceiptStatusSuccessful {
		log.Printf("Reward allocation successful! Block: %d", receipt.BlockNumber)
		
		// Optionally reset counters
		if cfg.ResetCounters {
			log.Println("Resetting Redis usage counters...")
			keysToReset := make([]string, len(eligibleProviders))
			for i, addr := range eligibleProviders {
				keysToReset[i] = "usage:" + addr.Hex()
			}
			if len(keysToReset) > 0 {
				_, err := redisClient.Del(ctx, keysToReset...).Result()
				if err != nil {
					log.Printf("Warning: Failed to reset some Redis counters: %v", err)
				} else {
					log.Println("Redis counters reset.")
				}
			}
		}
	} else {
		log.Fatalf("Reward allocation transaction failed! Block: %d, Status: %d", receipt.BlockNumber, receipt.Status)
	}

	log.Println("Reward allocation process finished.")
}

func loadConfig() Config {
	ethURL := os.Getenv("RPC_URL")
	if ethURL == "" {
		ethURL = "http://localhost:8545"
	}

	// Owner's private key - MUST be set
	privateKey := strings.TrimPrefix(os.Getenv("ALLOCATOR_PRIVATE_KEY"), "0x")
	if privateKey == "" {
		log.Fatal("ALLOCATOR_PRIVATE_KEY environment variable required")
	}

	registryAddrStr := os.Getenv("REGISTRY_CONTRACT")
	if registryAddrStr == "" {
		log.Fatal("REGISTRY_CONTRACT environment variable required")
	}

	redisAddr := os.Getenv("REDIS_ADDR") // e.g., "localhost:6379"
	if redisAddr == "" {
		log.Fatal("REDIS_ADDR environment variable required")
	}

	chainID := big.NewInt(31337) // Default chain ID
	if chainIDStr := os.Getenv("CHAIN_ID"); chainIDStr != "" {
		var ok bool
		chainID, ok = new(big.Int).SetString(chainIDStr, 10)
		if !ok {
			log.Fatal("Invalid CHAIN_ID format")
		}
	}

	totalReward := big.NewInt(100000000000000000000) // Default: 100 STK (assuming 18 decimals)
	if rewardStr := os.Getenv("TOTAL_REWARD_AMOUNT"); rewardStr != "" {
		var ok bool
		totalReward, ok = new(big.Int).SetString(rewardStr, 10)
		if !ok {
			log.Fatal("Invalid TOTAL_REWARD_AMOUNT format")
		}
	}
	
	var minUsage int64 = 1 // Default: Minimum 1 request
	if minUsageStr := os.Getenv("MIN_USAGE_COUNT"); minUsageStr != "" {
		_, scanErr := Sscan(minUsageStr, &minUsage)
		if scanErr != nil || minUsage < 0 {
			log.Fatal("Invalid MIN_USAGE_COUNT format")
		}
	}

	rewardPeriod := 24 * time.Hour // Default: 24 hours
	if periodStr := os.Getenv("REWARD_PERIOD_HOURS"); periodStr != "" {
		var hours int
		_, scanErr := Sscan(periodStr, &hours)
		if scanErr != nil || hours <= 0 {
			log.Fatal("Invalid REWARD_PERIOD_HOURS format")
		}
		rewardPeriod = time.Duration(hours) * time.Hour
	}

    resetCounters := false // Default: false
    if resetStr := os.Getenv("RESET_COUNTERS"); resetStr == "true" || resetStr == "1" {
        resetCounters = true
    }

	return Config{
		EthRPCURL:     ethURL,
		PrivateKey:    privateKey,
		RegistryAddr:  common.HexToAddress(registryAddrStr),
		RedisAddr:     redisAddr,
		ChainID:       chainID,
		TotalReward:   totalReward,
		MinUsage:      minUsage,
		RewardPeriod:  rewardPeriod, // Note: Not directly used in this simple script run, but good for context
		ResetCounters: resetCounters,
	}
}

// Sscan mimics fmt.Sscan behaviour for simpler cases needed here
func Sscan(str string, a ...interface{}) (int, error) {
    r := strings.NewReader(str)
    return fmt.Fscan(r, a...)
}

