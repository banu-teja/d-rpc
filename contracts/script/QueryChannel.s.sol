// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {PaymentChannel} from "../src/PaymentChannel.sol";

/// @title Script to query channel info
contract QueryChannel is Script {
    function run() external {
        // Load channel contract and ID from env
        PaymentChannel channel = PaymentChannel(vm.envAddress("CHANNEL_CONTRACT"));
        bytes32 channelId = vm.envBytes32("CHANNEL_ID");

        // Fetch struct via public getter
        (address user, address provider, , uint256 deposit, uint256 expiration, bool open) = channel.channels(channelId);

        console.log("Channel ID:");
        console.logBytes32(channelId);
        console.log("User:", user);
        console.log("Provider:", provider);
        console.log("Deposit:", deposit);
        console.log("Expiration:", expiration);
        console.log("Open:", open);
    }
}
