// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {Vm} from "forge-std/Vm.sol";
import {StakeToken} from "../src/StakeToken.sol";
import {PaymentChannel} from "../src/PaymentChannel.sol";

/// @title Script to open a payment channel
contract OpenChannel is Script {
    function run() external {
        vm.startBroadcast();

        // Load deployed contracts from env
        StakeToken stakeToken = StakeToken(vm.envAddress("STK_CONTRACT"));
        PaymentChannel channel = PaymentChannel(vm.envAddress("CHANNEL_CONTRACT"));

        // Channel parameters
        address provider = vm.envAddress("PROVIDER_ADDRESS");
        uint256 deposit = 10 ether;
        uint256 duration = 1 days;

        // Approve token transfer
        stakeToken.approve(address(channel), deposit);

        // Open channel and capture event
        vm.recordLogs();
        channel.openChannel(provider, stakeToken, deposit, duration);
        Vm.Log[] memory logs = vm.getRecordedLogs();
        
        // Extract channel ID from event
        bytes32 channelId;
        for (uint i = 0; i < logs.length; i++) {
            if (logs[i].topics[0] == keccak256("ChannelOpened(bytes32,address,address,uint256,uint256)")) {
                channelId = bytes32(logs[i].topics[1]);
                break;
            }
        }
        
        console.log("Channel opened with ID:");
        console.logBytes32(channelId);

        vm.stopBroadcast();
    }
}
