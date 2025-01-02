// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import { DEXPool } from "./DEXPool.sol";
import { IDEXPoolDeployer } from "./interfaces/IDEXPoolDeployer.sol";
import { LiquidityProviderERC20 } from "./LiquidityProviderERC20.sol";

contract DEXPoolDeployer is IDEXPoolDeployer {
    struct Parameters {
        address factory; // The factory address deploying the pool.
        address token0;  // The address of the first token in the pool.
        address token1;  // The address of the second token in the pool.
        address lpToken; // The address of the LP token contract.
        uint24 fee;      // The fee rate for the pool (in basis points).
    }

    // Parameters used during the pool deployment process.
    Parameters public override parameters;

    /**
     * @dev Deploys a new liquidity pool and its associated LP token contract.
     * @param factory The address of the factory initiating the deployment.
     * @param token0 The address of the first token in the pool.
     * @param token1 The address of the second token in the pool.
     * @param name The name of the LP token.
     * @param symbol The symbol of the LP token.
     * @param fee The fee rate applied to swaps in the pool.
     * @return pool The address of the deployed pool contract.
     */
    function deploy(
        address factory,
        address token0,
        address token1,
        string memory name,
        string memory symbol,
        uint24 fee
    ) internal returns (address) {
        // Create a new LP token with the specified name and symbol.
        address lpToken = address(new LiquidityProviderERC20(name, symbol, address(this)));

        // Initialize parameters for the new pool.
        parameters = Parameters({
            factory: factory,
            token0: token0,
            token1: token1,
            lpToken: lpToken,
            fee: fee
        });

        // Deploy the liquidity pool using a unique salt derived from token addresses and the fee.
        address pool = address(new DEXPool{salt: keccak256(abi.encode(token0, token1, fee))}());
        
        // Transfer ownership of the LP token contract to the newly deployed pool.
        LiquidityProviderERC20(lpToken).transferOwnership(pool);

        // Clear parameters after deployment to optimize gas usage.
        delete parameters;

        return pool;
    } 
}