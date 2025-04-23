// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {StakeToken} from "../src/StakeToken.sol";
import {ProviderRegistry} from "../src/ProviderRegistry.sol";

/// @title Script to stake and register a provider
contract RegisterProvider is Script {
    function run() external {
        vm.startBroadcast();

        // Deployed addresses from previous run
        StakeToken stakeToken = StakeToken(0x5FbDB2315678afecb367f032d93F642f64180aa3);
        ProviderRegistry registry = ProviderRegistry(0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512);

        uint256 minStake = registry.minStake();
        console.log("Minimum stake:", minStake);

        // Approve and deposit stake
        stakeToken.approve(address(registry), minStake);
        registry.depositStake(minStake);
        console.log("Staked", minStake, "tokens");

        // Register as provider
        registry.register();
        console.log("Provider registered at address:", address(this));

        vm.stopBroadcast();
    }
}
