// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import { IDEXFactory } from "./interfaces/IDEXFactory.sol";
import { DEXPoolDeployer } from "./DEXPoolDeployer.sol";

contract DEXFactory is IDEXFactory, DEXPoolDeployer {
    address private _owner;

    mapping (address pool => uint24 fee) public _fees;
    mapping (address token0 => mapping (address token1 => address)) _pools;

    constructor() {
        _owner = msg.sender;
        emit OwnerChanged(address(0), msg.sender);
    }

    modifier onlyOwner() {
        require(msg.sender == _owner, OnlyOwnerAllowed());
        _;
    }

    function owner() external view returns (address) {
        return _owner;
    }

    function setOwner(address newOwner) external onlyOwner {
        emit OwnerChanged(_owner, newOwner);
        _owner = newOwner;
    }

    function getPool(address token0, address token1) external view returns (address) {
        return _pools[token0][token1];
    }

    function createPool(
        address token0, 
        address token1, 
        uint24 fee) 
    external returns (address) {
        require(_pools[token0][token1] == address(0), PoolAlreadyExist());
        require(token0 != address(0) && token1 != address(0), TokenAddressIsZero());
        (address tokenA, address tokenB) = token0 < token1 ? (token0, token1) : (token1, token0);
        address pool = deploy(address(this), tokenA, tokenB, fee);

        _pools[tokenA][tokenB] = pool;
        _pools[tokenB][tokenA] = pool;

        emit PoolCreated(tokenA, tokenB, fee, pool);

        return pool;
    }

    function getFee(address pool) external view returns (uint24) {
        return _fees[pool];
    }
}