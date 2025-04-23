// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {StakeToken} from "../src/StakeToken.sol";
import {ProviderRegistry} from "../src/ProviderRegistry.sol";
import {PaymentChannel} from "../src/PaymentChannel.sol";
import {IERC20} from "openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";

/// @title Deployment script for ProviderRegistry and PaymentChannel
contract DeployContracts is Script {
    function run() external {
        // Start broadcasting via private key from environment: FORGE_PRIVATE_KEY
        vm.startBroadcast();

        // Deploy the StakeToken ERC20 for staking/payments
        StakeToken stakeToken = new StakeToken();
        console.log("Deployed StakeToken at:", address(stakeToken));

        // Minimum stake for providers
        uint256 minStake = 100 ether;
        ProviderRegistry registry = new ProviderRegistry(IERC20(stakeToken), minStake);
        console.log("Deployed ProviderRegistry at:", address(registry));

        // Deploy payment channel contract
        PaymentChannel channel = new PaymentChannel();
        console.log("Deployed PaymentChannel at:", address(channel));

        vm.stopBroadcast();
    }
}
