// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

interface IDEXPool {
    error InvalidInputAmount();

    error InsufficientBalance();

    error InvalidFeeValue(
        uint24 oldFee,
        uint24 newFee
    );

    error InsufficientSlippageSet(
        uint256 minAmount, 
        uint256 requiredAmount
    );

    error WrongLiquidityAmout(
        uint256 required, 
        uint256 value
    );

    error OnlyFactoryAllowed();

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

    event RemoveLiquidity(
        address indexed account,
        uint256 amount
    );

    event FeeUpdated(
        uint24 oldFee,
        uint24 newFee
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

    function removeLiquidity(
        uint256 amount
    ) external;

    function getReserves() external view returns (uint256, uint256);

    function updateFee(uint24 _fee) external;
}