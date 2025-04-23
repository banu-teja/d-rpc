// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script, console} from "forge-std/Script.sol";
import {StakeToken} from "../src/StakeToken.sol";

/// @title Script to fund a user account with STK tokens
contract FundUser is Script {
    function run() external {
        vm.startBroadcast();

        // Read token, user, and amount from env
        StakeToken token = StakeToken(vm.envAddress("STK_CONTRACT"));
        address user = vm.envAddress("USER_ADDRESS");
        uint256 amount = vm.envUint("CHANNEL_DEPOSIT");

        // Transfer from deployer (msg.sender) to user
        token.transfer(user, amount);
        console.log("Funded user");
        console.log(user);
        console.log("with");
        console.log(amount);
        console.log("STK");

        vm.stopBroadcast();
    }
}
