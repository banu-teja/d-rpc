// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";
import "openzeppelin-contracts/contracts/utils/cryptography/MessageHashUtils.sol";
import "../src/PaymentChannel.sol";

contract MockToken2 is ERC20 {
    constructor() ERC20("MockToken2", "MTK2") {}
    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }
}

contract PaymentChannelTest is Test {
    using MessageHashUtils for bytes32;
    MockToken2 token;
    PaymentChannel chan;
    uint256 userKey;
    address user;
    address provider = address(0x2);
    uint256 deposit = 100 ether;
    uint256 duration = 1 days;
    bytes32 channelId;

    function setUp() public {
        userKey = 1;
        user = vm.addr(userKey);
        token = new MockToken2();
        chan = new PaymentChannel();
        token.mint(user, deposit);

        vm.startPrank(user);
        token.approve(address(chan), deposit);
        chan.openChannel(provider, token, deposit, duration);
        // calculate channelId deterministically
        channelId = keccak256(abi.encodePacked(user, provider, address(token), deposit, block.timestamp));
        vm.stopPrank();
    }

    function testCloseChannel() public {
        // user signs prefixed message off-chain
        bytes32 message = keccak256(abi.encodePacked(channelId, deposit/2));
        bytes32 digest = message.toEthSignedMessageHash();
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(userKey, digest);
        bytes memory sig = abi.encodePacked(r, s, v);

        vm.prank(provider);
        chan.closeChannel(channelId, deposit/2, sig);

        assertEq(token.balanceOf(provider), deposit/2);
        assertEq(token.balanceOf(user), deposit/2);
    }

    function testClaimTimeout() public {
        // move past expiration
        skip(duration + 1);
        vm.prank(user);
        chan.claimTimeout(channelId);
        assertEq(token.balanceOf(user), deposit);
    }
}
