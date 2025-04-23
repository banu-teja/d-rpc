// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";

/// @title Stake Token for ProviderRegistry
/// @notice ERC20 token used for staking and payments
contract StakeToken is ERC20 {
    constructor() ERC20("StakeToken", "STK") {
        // initial mint to deployer
        _mint(msg.sender, 1000000 * 10 ** decimals());
    }
}
