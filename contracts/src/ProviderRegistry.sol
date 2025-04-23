// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";
import "openzeppelin-contracts/contracts/utils/ReentrancyGuard.sol";
import "openzeppelin-contracts/contracts/access/Ownable.sol";

/// @title Provider Registry with staking and slashing
/// @notice Allows nodes to deposit stake, register/deregister as RPC providers, and be slashed for misbehavior
contract ProviderRegistry is Ownable, ReentrancyGuard {
    IERC20 public immutable stakeToken;
    uint256 public minStake;

    struct Provider {
        uint256 stake;
        uint256 qosScore;
        bool registered;
    }

    mapping(address => Provider) public providers;

    event ProviderRegistered(address indexed provider, uint256 stake);
    event ProviderDeregistered(address indexed provider);
    event StakeDeposited(address indexed provider, uint256 amount);
    event StakeWithdrawn(address indexed provider, uint256 amount);
    event ProviderSlashed(address indexed provider, uint256 amount);
    event QoSUpdated(address indexed provider, uint256 newScore);

    /// @param _stakeToken ERC20 token used for staking
    /// @param _minStake Minimum stake required to register
    constructor(IERC20 _stakeToken, uint256 _minStake) Ownable(msg.sender) {
        stakeToken = _stakeToken;
        minStake = _minStake;
    }

    /// @notice Owner can update minimum stake
    function setMinStake(uint256 _minStake) external onlyOwner {
        minStake = _minStake;
    }

    /// @notice Deposit stake tokens
    function depositStake(uint256 amount) external nonReentrant {
        require(amount > 0, "Amount must be > 0");
        stakeToken.transferFrom(msg.sender, address(this), amount);
        providers[msg.sender].stake += amount;
        emit StakeDeposited(msg.sender, amount);
    }

    /// @notice Register as provider (must have >= minStake deposited)
    function register() external nonReentrant {
        Provider storage p = providers[msg.sender];
        require(!p.registered, "Already registered");
        require(p.stake >= minStake, "Insufficient stake");
        p.registered = true;
        emit ProviderRegistered(msg.sender, p.stake);
    }

    /// @notice Deregister without withdrawing stake
    function deregister() external nonReentrant {
        Provider storage p = providers[msg.sender];
        require(p.registered, "Not registered");
        p.registered = false;
        emit ProviderDeregistered(msg.sender);
    }

    /// @notice Withdraw stake after deregistration
    function withdrawStake(uint256 amount) external nonReentrant {
        Provider storage p = providers[msg.sender];
        require(!p.registered, "Deregister first");
        require(amount > 0 && amount <= p.stake, "Invalid amount");
        p.stake -= amount;
        stakeToken.transfer(msg.sender, amount);
        emit StakeWithdrawn(msg.sender, amount);
    }

    /// @notice Slash a provider for misbehavior, only owner
    function slashProvider(address provider_, uint256 amount) external onlyOwner nonReentrant {
        Provider storage p = providers[provider_];
        require(p.stake >= amount, "Insufficient stake to slash");
        p.stake -= amount;
        // transferred stake to owner (governance)
        stakeToken.transfer(owner(), amount);
        emit ProviderSlashed(provider_, amount);
    }

    /// @notice Update QoS score for a provider
    function updateQoS(address provider_, uint256 score) external onlyOwner nonReentrant {
        Provider storage p = providers[provider_];
        require(p.registered, "Provider not registered");
        p.qosScore = score;
        emit QoSUpdated(provider_, score);
    }
}
