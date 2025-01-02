// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

interface IDEXPool {
    error InvalidInputAmount();

    error InsufficientBalance();

    error InsufficientSlipperage();

    error WrongLiquidityAmout(
        uint256 required, 
        uint256 value
    );

    event SwapTokens(
        address indexed account,
        uint256 amountIn,
        uint256 amountOut,
        bool zeroToOne
    );

    event AddLiquidity(
        address indexed account,
        uint256 amount0In,
        uint256 amount1In
    );

    function swap(
        uint256 amountIn, 
        uint256 amountOutMin, 
        bool zeroToOne) 
    external;

    function addLiquidity(
        uint256 amount0In,
        uint256 amount1In
    ) external;

    function getReserve0() external view returns (uint256);

    function getReserve1() external view returns (uint256);
}