// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";
import "openzeppelin-contracts/contracts/utils/ReentrancyGuard.sol";
import "openzeppelin-contracts/contracts/access/Ownable.sol";

contract ProviderRegistry is Ownable, ReentrancyGuard {
    IERC20 public immutable stakeToken;
    uint256 public minStake;
    uint256 public constant DEREGISTRATION_COOLDOWN = 1 days;

    struct Provider {
        uint256 stake;
        uint256 qosScore;
        bool registered;
        uint40 registrationTime;
        uint40 deregistrationTime;
        string rpcUrl;

    // TODO: Implement reward distribution logic based on usage and/or QoS
    // This function would calculate and distribute rewards to providers
    // based on collected metrics over a period.
    // function distributeRewards() external onlyOwner {
    //     // Distribution logic here
    // }
}

    mapping(address => Provider) public providers;
    address[] public registeredProviderAddresses; // Array to keep track of registered provider addresses

    event ProviderRegistered(address indexed provider, uint256 stake);
    event ProviderDeregistered(address indexed provider);
    event StakeDeposited(address indexed provider, uint256 amount);
    event StakeWithdrawn(address indexed provider, uint256 amount);
    event ProviderSlashed(address indexed provider, uint256 amount);
    event QoSUpdated(address indexed provider, uint256 newScore);

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

    function register(string calldata rpcUrl) external nonReentrant {
        Provider storage p = providers[msg.sender];
        require(!p.registered, "Already registered");
        require(p.stake >= minStake, "Insufficient stake");
        require(block.timestamp > p.deregistrationTime + DEREGISTRATION_COOLDOWN, "Cooldown active");
        p.registered = true;
        p.registrationTime = uint40(block.timestamp);
        p.rpcUrl = rpcUrl;
        emit ProviderRegistered(msg.sender, p.stake);

        // Add provider address to the list if not already present (should not be if not registered)
        bool found = false;
        for (uint i = 0; i < registeredProviderAddresses.length; i++) {
            if (registeredProviderAddresses[i] == msg.sender) {
                found = true;
                break;
            }
        }
        if (!found) {
            registeredProviderAddresses.push(msg.sender);
        }
    }

    function deregister() external nonReentrant {
        Provider storage p = providers[msg.sender];
        require(p.registered, "Not registered");

        // Remove provider address from the list
        for (uint i = 0; i < registeredProviderAddresses.length; i++) {
            if (registeredProviderAddresses[i] == msg.sender) {
                registeredProviderAddresses[i] = registeredProviderAddresses[registeredProviderAddresses.length - 1];
                registeredProviderAddresses.pop();
                break;
            }
        }

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

    // Function to get the list of registered provider addresses
    function getRegisteredProviders() external view returns (address[] memory) {
        return registeredProviderAddresses;
    }
}

```
        p.qosScore = score;
        emit QoSUpdated(provider_, score);
    }
}
