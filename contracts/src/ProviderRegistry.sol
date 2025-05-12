// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";
import "openzeppelin-contracts/contracts/utils/ReentrancyGuard.sol";
import "openzeppelin-contracts/contracts/access/Ownable.sol";
import "openzeppelin-contracts/contracts/utils/math/Math.sol"; // For safe math operations

contract ProviderRegistry is Ownable, ReentrancyGuard {
    using Math for uint256;

    IERC20 public immutable stakeToken;
    uint256 public minStake;
    uint256 public constant DEREGISTRATION_COOLDOWN = 1 days;

    // Reward Pool
    uint256 public rewardPoolBalance; // Total STK tokens available for rewards
    mapping(address => uint256) public providerRewards; // Amount of STK claimable by each provider

    struct Provider {
        uint256 stake;
        uint256 qosScore;
        bool registered;
        uint40 registrationTime;
        uint40 deregistrationTime;
        string endpointURL; // Added: RPC endpoint URL
    }

    mapping(address => Provider) public providers;

    event ProviderRegistered(address indexed provider, uint256 stake, string endpointURL); // Added URL
    event ProviderDeregistered(address indexed provider);
    event StakeDeposited(address indexed provider, uint256 amount);
    event StakeWithdrawn(address indexed provider, uint256 amount);
    event ProviderSlashed(address indexed provider, uint256 amount);
    event QoSUpdated(address indexed provider, uint256 newScore);
    event ProviderURLUpdated(address indexed provider, string newURL); // Added event for URL changes
    event RewardPoolFunded(address indexed funder, uint256 amount);
    event RewardsAllocated(address indexed allocator, uint256 totalAmount);
    event RewardClaimed(address indexed provider, uint256 amount);

    constructor(IERC20 _stakeToken, uint256 _minStake) Ownable(msg.sender) {
        stakeToken = _stakeToken;
        minStake = _minStake;
    }

    function setMinStake(uint256 _minStake) external onlyOwner {
        require(_minStake >= 0.1 ether, "Min stake too low");
        require(_minStake <= 1000 ether, "Min stake too high");
        minStake = _minStake;
    }

    function depositStake(uint256 amount) external nonReentrant {
        require(amount > 0, "Amount must be > 0");
        stakeToken.transferFrom(msg.sender, address(this), amount);
        providers[msg.sender].stake += amount;
        emit StakeDeposited(msg.sender, amount);
    }

    function register() external nonReentrant {
        Provider storage p = providers[msg.sender];
        require(!p.registered, "Already registered");
        require(p.stake >= minStake, "Insufficient stake");
        require(block.timestamp > p.deregistrationTime + DEREGISTRATION_COOLDOWN, "Cooldown active");
        p.registered = true;
        p.registrationTime = uint40(block.timestamp);
        // Provider must set URL before registering fully
        require(bytes(p.endpointURL).length > 0, "Endpoint URL not set"); 
        emit ProviderRegistered(msg.sender, p.stake, p.endpointURL);
    }

    /// @notice Allows a provider to set or update their RPC endpoint URL.
    /// @param _url The new endpoint URL (e.g., "https://node.example.com:8545").
    function setEndpointURL(string memory _url) external {
        Provider storage p = providers[msg.sender];
        // Basic validation: Ensure URL is not excessively long (e.g., > 256 chars)
        require(bytes(_url).length > 0 && bytes(_url).length <= 256, "Invalid URL length");
        // Optional: Could add more sophisticated URL format checks off-chain or basic on-chain
        string memory oldUrl = p.endpointURL;
        p.endpointURL = _url;
        // Emit event if URL actually changed
        if (keccak256(bytes(oldUrl)) != keccak256(bytes(_url))) { 
             emit ProviderURLUpdated(msg.sender, _url);
        }
    }

    function deregister() external nonReentrant {
        Provider storage p = providers[msg.sender];
        require(p.registered, "Not registered");
        p.registered = false;
        p.deregistrationTime = uint40(block.timestamp);
        emit ProviderDeregistered(msg.sender);
    }

    function withdrawStake(uint256 amount) external nonReentrant {
        Provider storage p = providers[msg.sender];
        require(!p.registered, "Deregister first");
        require(block.timestamp > p.deregistrationTime + DEREGISTRATION_COOLDOWN, "Cooldown active");
        require(amount > 0 && amount <= p.stake, "Invalid amount");
        p.stake -= amount;
        stakeToken.transfer(msg.sender, amount);
        emit StakeWithdrawn(msg.sender, amount);
    }

    function slashProvider(address provider_, uint256 amount) external onlyOwner nonReentrant {
        Provider storage p = providers[provider_];
        require(p.stake >= amount, "Insufficient stake to slash");
        p.stake -= amount;
        stakeToken.transfer(owner(), amount);
        emit ProviderSlashed(provider_, amount);
    }

    function updateQoS(address provider_, uint256 score) external onlyOwner nonReentrant {
        Provider storage p = providers[provider_];
        require(p.registered, "Provider not registered");
        p.qosScore = score;
        emit QoSUpdated(provider_, score);
    }

    // --- Reward Functions ---

    /// @notice Fund the reward pool with STK tokens. Only callable by the owner.
    /// @param amount The amount of STK tokens to add to the pool.
    function fundRewardPool(uint256 amount) external onlyOwner nonReentrant {
        require(amount > 0, "Amount must be > 0");
        stakeToken.transferFrom(msg.sender, address(this), amount);
        rewardPoolBalance += amount;
        emit RewardPoolFunded(msg.sender, amount);
    }

    /// @notice Calculate and allocate rewards to providers based on stake and QoS score.
    /// @dev Calculates rewards proportionally based on `stake * qosScore`.
    ///      Intended to be called by the owner with off-chain calculated inputs.
    /// @param providersToReward List of provider addresses to reward.
    /// @param totalRewardAmount Total amount of STK to distribute from the reward pool.
    function calculateAndAllocateRewards(address[] calldata providersToReward, uint256 totalRewardAmount)
        external
        onlyOwner
        nonReentrant
    {
        require(providersToReward.length > 0, "No providers specified");
        require(totalRewardAmount > 0, "Total reward must be > 0");
        require(rewardPoolBalance >= totalRewardAmount, "Insufficient pool balance");

        uint256 totalWeight = 0;
        uint256[] memory weights = new uint256[](providersToReward.length);

        // Calculate total weight (stake * qosScore)
        for (uint i = 0; i < providersToReward.length; i++) {
            address providerAddr = providersToReward[i];
            Provider storage p = providers[providerAddr];
            require(p.registered, "Provider not registered");
            // Basic weight: stake * qosScore (ensure qosScore > 0 to avoid division by zero later)
            // If qosScore is 0, weight is 0. Add 1 to qosScore to prevent weight=0 if stake>0? No, 0 score = 0 reward.
            uint256 weight = p.stake * p.qosScore; 
            weights[i] = weight;
            totalWeight += weight;
        }

        require(totalWeight > 0, "Total weight must be > 0");

        uint256 allocatedAmount = 0;
        // Allocate rewards proportionally
        for (uint i = 0; i < providersToReward.length; i++) {
            address providerAddr = providersToReward[i];
            uint256 weight = weights[i];
            if (weight > 0) {
                // reward = (weight / totalWeight) * totalRewardAmount
                uint256 reward = (weight * totalRewardAmount) / totalWeight;
                providerRewards[providerAddr] += reward;
                allocatedAmount += reward;
            }
        }

        // Ensure allocated amount does not exceed total reward (handles potential rounding dust)
        // This check should ideally not fail due to the way allocation is done, but good for safety.
        require(allocatedAmount <= totalRewardAmount, "Allocation exceeded total");

        rewardPoolBalance -= allocatedAmount; // Decrease pool balance by allocated amount
        emit RewardsAllocated(msg.sender, allocatedAmount);
    }

    /// @notice Allows a provider to claim their accumulated rewards.
    function claimRewards() external nonReentrant {
        uint256 rewardAmount = providerRewards[msg.sender];
        require(rewardAmount > 0, "No rewards to claim");

        providerRewards[msg.sender] = 0; // Reset rewards before transfer (prevents reentrancy)
        stakeToken.transfer(msg.sender, rewardAmount);

        emit RewardClaimed(msg.sender, rewardAmount);
    }
}
