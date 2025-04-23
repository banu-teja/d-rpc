// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {PaymentChannel} from "../src/PaymentChannel.sol";
import {IERC20} from "openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";

/// @title Script to close a payment channel using a signed voucher
contract CloseChannel is Script {
    function run() external {
        // Read keys from env: USER_KEY for signing, for broadcast use --private-key (provider key)
        uint256 userKey = vm.envUint("USER_KEY");

        // Load deployed addresses and channel ID from env
        PaymentChannel channel = PaymentChannel(vm.envAddress("CHANNEL_CONTRACT"));
        IERC20 token = IERC20(vm.envAddress("STK_CONTRACT"));
        // Read channel ID from environment
        bytes32 channelId = vm.envBytes32("CHANNEL_ID");
        uint256 payout = 5 ether;

        // Off-chain signing by user
        bytes32 message = keccak256(abi.encodePacked(channelId, payout));
        // Ethereum Signed Message prefix to match personal_sign
        bytes32 digest = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", message));
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(userKey, digest);
        bytes memory sig = abi.encodePacked(r, s, v);

        // Start broadcast using provider key passed via CLI flag
        vm.startBroadcast();
        // Debug: print broadcasting address
        console.log("Broadcasting as:", vm.addr(vm.envUint("FORGE_PRIVATE_KEY")));

        // Provider closes the channel
        channel.closeChannel(channelId, payout, sig);
        console.log("Channel closed for payout:", payout);

        // Log final balances
        address provider = msg.sender;
        // Destructure channels mapping return values to get user
        (, address user, , , , ) = channel.channels(channelId);
        console.log("Provider balance:", token.balanceOf(provider));
        console.log("User balance:", token.balanceOf(user));

        vm.stopBroadcast();
    }
}
