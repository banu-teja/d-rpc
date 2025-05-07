// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "forge-std/Script.sol";
import "../src/StakeToken.sol";
import "../src/ProviderRegistry.sol";
import "../src/PaymentChannel.sol";

contract DeployScript is Script {
    function setUp() public {}

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        // Deploy StakeToken
        StakeToken token = new StakeToken();
        console.log("StakeToken deployed at:", address(token));

        // Deploy ProviderRegistry with minimum stake of 1 ETH
        uint256 minStake = 1 ether;
        ProviderRegistry registry = new ProviderRegistry(token, minStake);
        console.log("ProviderRegistry deployed at:", address(registry));

        // Deploy PaymentChannel
        PaymentChannel channel = new PaymentChannel();
        console.log("PaymentChannel deployed at:", address(channel));

        // Write deployment addresses to file
        string memory contractAddresses = string(
            abi.encodePacked(
                "export STK_CONTRACT=", vm.toString(address(token)), "\n",
                "export REGISTRY_CONTRACT=", vm.toString(address(registry)), "\n",
                "export CHANNEL_CONTRACT=", vm.toString(address(channel)), "\n"
            )
        );
        vm.writeFile(".env", contractAddresses);
        console.log("Contract addresses written to .env file");

        vm.stopBroadcast();
    }
} 