// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";
import "../src/ProviderRegistry.sol";

contract MockToken is ERC20 {
    constructor() ERC20("MockToken", "MTK") {}
    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }
}

contract ProviderRegistryTest is Test {
    MockToken token;
    ProviderRegistry reg;
    address owner = address(0x1);
    address providerAddr = address(0x2);
    address other = address(0x3);
    uint256 minStake = 100 ether;

    function setUp() public {
        vm.prank(owner);
        token = new MockToken();
        vm.prank(owner);
        reg = new ProviderRegistry(IERC20(token), minStake);
        token.mint(providerAddr, 500 ether);
        token.mint(other, 500 ether);
    }

    function testDepositAndRegister() public {
        vm.startPrank(providerAddr);
        token.approve(address(reg), 200 ether);
        reg.depositStake(200 ether);
        (uint256 stake, , , , ) = reg.providers(providerAddr);
        assertEq(stake, 200 ether);
        reg.register();
        (, , bool isRegistered, , ) = reg.providers(providerAddr);
        assertTrue(isRegistered);
        vm.stopPrank();
    }

    function testRevertRegisterInsufficientStake() public {
        vm.prank(providerAddr);
        vm.expectRevert("Insufficient stake");
        reg.register();
    }

    function testDeregisterAndWithdraw() public {
        vm.startPrank(providerAddr);
        token.approve(address(reg), minStake);
        reg.depositStake(minStake);
        reg.register();
        reg.deregister();
        reg.withdrawStake(minStake);
        assertEq(token.balanceOf(providerAddr), 500 ether);
        vm.stopPrank();
    }

    function testSlashProvider() public {
        vm.startPrank(providerAddr);
        token.approve(address(reg), minStake);
        reg.depositStake(minStake);
        reg.register();
        vm.stopPrank();

        vm.prank(owner);
        reg.slashProvider(providerAddr, 50 ether);
        (uint256 newStake, , , , ) = reg.providers(providerAddr);
        assertEq(newStake, minStake - 50 ether);
        assertEq(token.balanceOf(owner), 50 ether);
    }

    function testUpdateQoS() public {
        vm.startPrank(providerAddr);
        token.approve(address(reg), minStake);
        reg.depositStake(minStake);
        reg.register();
        vm.stopPrank();

        vm.prank(owner);
        reg.updateQoS(providerAddr, 42);
        (, uint256 score, , , ) = reg.providers(providerAddr);
        assertEq(score, 42);
    }
}
