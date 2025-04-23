// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "openzeppelin-contracts/contracts/token/ERC20/IERC20.sol";
import "openzeppelin-contracts/contracts/utils/cryptography/ECDSA.sol";
import "openzeppelin-contracts/contracts/utils/cryptography/MessageHashUtils.sol";
import "openzeppelin-contracts/contracts/utils/ReentrancyGuard.sol";

/// @title Micropayment Channel for off-chain payments
/// @notice Enables users to open unidirectional payment channels with providers using ERC20 tokens
contract PaymentChannel is ReentrancyGuard {
    using ECDSA for bytes32;
    using MessageHashUtils for bytes32;

    struct Channel {
        address user;
        address provider;
        IERC20 token;
        uint256 deposit;
        uint256 expiration;
        bool open;
    }

    mapping(bytes32 => Channel) public channels;

    event ChannelOpened(bytes32 indexed channelId, address indexed user, address indexed provider, uint256 deposit, uint256 expiration);
    event ChannelClosed(bytes32 indexed channelId, uint256 amountReceived);
    event ChannelExpired(bytes32 indexed channelId);

    /// @notice Open a new payment channel and return its ID
    /// @param provider Recipient of payments
    /// @param token ERC20 token to use
    /// @param deposit Amount of tokens to deposit
    /// @param duration Channel duration in seconds
    function openChannel(address provider, IERC20 token, uint256 deposit, uint256 duration)
        external
        nonReentrant
        returns (bytes32 channelId)
    {
        require(deposit > 0, "Deposit must be > 0");
        require(provider != address(0), "Invalid provider");
        token.transferFrom(msg.sender, address(this), deposit);
        channelId = keccak256(abi.encodePacked(msg.sender, provider, address(token), deposit, block.timestamp));
        channels[channelId] = Channel(msg.sender, provider, token, deposit, block.timestamp + duration, true);
        emit ChannelOpened(channelId, msg.sender, provider, deposit, block.timestamp + duration);
    }

    /// @notice Close channel by provider presenting a signed voucher from user
    /// @param channelId Channel identifier
    /// @param amount Amount to send to provider
    /// @param signature Signature of voucher (channelId and amount) from user
    function closeChannel(bytes32 channelId, uint256 amount, bytes memory signature) external nonReentrant {
        Channel storage ch = channels[channelId];
        require(ch.open, "Channel is closed");
        require(block.timestamp <= ch.expiration, "Channel expired");
        require(msg.sender == ch.provider, "Only provider can close");
        require(amount <= ch.deposit, "Amount exceeds deposit");

        bytes32 message = keccak256(abi.encodePacked(channelId, amount));
        bytes32 digest = message.toEthSignedMessageHash();
        address signer = digest.recover(signature);
        require(signer == ch.user, "Invalid signature");

        ch.open = false;
        uint256 remaining = ch.deposit - amount;
        ch.token.transfer(ch.provider, amount);
        if (remaining > 0) {
            ch.token.transfer(ch.user, remaining);
        }

        emit ChannelClosed(channelId, amount);
    }

    /// @notice Claim refund if channel expired and not closed
    /// @param channelId Channel identifier
    function claimTimeout(bytes32 channelId) external nonReentrant {
        Channel storage ch = channels[channelId];
        require(ch.open, "Channel is closed");
        require(block.timestamp > ch.expiration, "Channel not yet expired");
        require(msg.sender == ch.user, "Only user can claim");

        ch.open = false;
        uint256 amount = ch.deposit;
        ch.token.transfer(ch.user, amount);

        emit ChannelExpired(channelId);
    }
}
