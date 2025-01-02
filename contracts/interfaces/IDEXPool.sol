// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

interface IDEXPool {
    error InvalidInputAmount();

    error InsufficientBalance();

    error InsufficientSlipperage();

    event SwapTokens(
        address indexed account,
        uint256 amountIn,
        uint256 amountOut,
        bool zeroToOne
    );

    function swap(
        uint256 amountIn, 
        uint256 amountOutMin, 
        bool zeroToOne) 
    external;

    function getReserve0() external view returns (uint256);

    function getReserve1() external view returns (uint256);
}