// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

interface IDEXFactory {
    error PoolAlreadyExist();

    error TokenAddressIsZero();

    event PoolCreated(
        address indexed token0,
        address indexed token1,
        uint24 indexed fee,
        address pool
    );

    function createPool(
        address token0,
        address token1,
        string memory name,
        string memory symbol,
        uint24 fee) external returns (address);

    function getPool(address token0, address token1) external view returns (address);

    function getFee(address pool) external view returns (uint24);
}