// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

interface IDEXSwapPoolDeployer {
    function parameters() 
    external view
    returns (
        address factory,
        address token0,
        address token1,
        address lpToken,
        uint24 fee
    );
}