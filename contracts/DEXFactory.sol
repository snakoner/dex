// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { IDEXFactory } from "./interfaces/IDEXFactory.sol";
import { DEXPoolDeployer } from "./DEXPoolDeployer.sol";

contract DEXFactory is IDEXFactory, DEXPoolDeployer, Ownable {
    // Mapping to store the fee associated with each pool
    mapping (address pool => uint24 fee) public _fees;

    // Nested mapping to store the addresses of pools for token pairs
    mapping (address token0 => mapping (address token1 => address)) _pools;

    /**
     * @dev Initializes the contract by setting the deployer as the owner.
     */
    constructor() Ownable(msg.sender) {}

    /**
     * @dev Retrieves the address of the pool for a given pair of tokens.
     * @param token0 The address of the first token.
     * @param token1 The address of the second token.
     * @return The address of the corresponding pool.
     */
    function getPool(address token0, address token1) external view returns (address) {
        return _pools[token0][token1];
    }

    /**
     * @dev Creates a new liquidity pool for the given token pair.
     * @param token0 The address of the first token.
     * @param token1 The address of the second token.
     * @param name The name of the LP token.
     * @param symbol The symbol of the LP token.
     * @param fee The fee rate for the pool in basis points.
     * @return The address of the newly created pool.
     */
    function createPool(
        address token0, 
        address token1, 
        string memory name,
        string memory symbol,
        uint24 fee
    ) external onlyOwner returns (address) {
        require(_pools[token0][token1] == address(0), PoolAlreadyExist());
        require(token0 != address(0) && token1 != address(0), TokenAddressIsZero());

        // Ensure consistent ordering of token addresses
        (address tokenA, address tokenB) = token0 < token1 ? (token0, token1) : (token1, token0);

        // Deploy the new pool
        address pool = deploy(address(this), tokenA, tokenB, name, symbol, fee);

        // Update the pool mappings
        _pools[tokenA][tokenB] = pool;
        _pools[tokenB][tokenA] = pool;

        emit PoolCreated(tokenA, tokenB, fee, pool);

        return pool;
    }

    /**
     * @dev Retrieves the fee associated with a specific pool.
     * @param pool The address of the pool.
     * @return The fee rate for the pool in basis points.
     */
    function getFee(address pool) external view returns (uint24) {
        return _fees[pool];
    }
}