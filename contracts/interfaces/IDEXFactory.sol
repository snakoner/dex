// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

interface IDEXFactory {
    error OnlyOwnerAllowed();

    error PoolAlreadyExist();

    error TokenAddressIsZero();

    event OwnerChanged(address indexed oldOwner, address indexed newOwner);

    event PoolCreated(
        address indexed token0,
        address indexed token1,
        uint24 indexed fee,
        address pool
    );

    function owner() external view returns (address);

    function setOwner(address newOwner) external;

    function getPool(address token0, address token1) external view returns (address);

    function createPool(address token0, address token1, uint24 fee) external returns (address);

    function getFee(address pool) external view returns (uint24);
}