// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { IDEXPoolDeployer } from "./interfaces/IDEXPoolDeployer.sol";

// Explanation:
// * we want to exchange dx of token0:
// ************************************
// (x + dx)(y - dy) = xy
// y - dy = xy / (x + dx)
// dy = y - xy / (x + dx) = (xy + y*dx - xy) / (x + dx) = y*dx / (x + dx)

contract DEXPool {
    error InvalidInputAmount();

    error InsufficientBalance();

    address immutable public factory;
    address immutable public token0;
    address immutable public token1;
    uint24 immutable public fee;

    constructor() {
        (factory, token0, token1, fee) = IDEXPoolDeployer(msg.sender).parameters();
    }

    /*
        read functions
    */
    function getReserve0() public view returns (uint256) {
        return IERC20(token0).balanceOf(address(this));
    }

    function getReserve1() public view returns (uint256) {
        return IERC20(token1).balanceOf(address(this));
    }

    function getOutputAmount(uint256 amount, bool zeroToOne) public view returns (uint256) {
        require(amount > 0, InvalidInputAmount());
        (uint256 _reserve0, uint256 _reserve1) = zeroToOne 
            ? (getReserve0(), getReserve1()) 
            : (getReserve1(), getReserve0());

        // probably unused
        require(_reserve0 > 0 && _reserve1 > 0, InsufficientBalance());

        return _reserve1 * amount / (_reserve0 + amount);
    }

    function swap(uint256 amount, bool zeroToOne) external {
        uint256 ouputAmount = getOutputAmount(amount, zeroToOne);
        (address _token0, address _token1) = zeroToOne 
            ? (token0, token1) 
            : (token1, token0);

        IERC20(_token0).transferFrom(msg.sender, address(this), amount);
        IERC20(_token1).transfer(msg.sender, ouputAmount);
    }

}