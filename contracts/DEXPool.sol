// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { IDEXPoolDeployer } from "./interfaces/IDEXPoolDeployer.sol";
import { IDEXPool } from "./interfaces/IDEXPool.sol";

/*
 * explanation:
 * we want to exchange dx of token0:
 ************************************
 1. (x + dx)(y - dy) = xy
 2. y - dy = xy / (x + dx)
 3. dy = y - xy / (x + dx) = (xy + y * dx - xy) / (x + dx) = y * dx / (x + dx)
*/

contract DEXPool is IDEXPool {
    address immutable public factory;
    address immutable public token0;
    address immutable public token1;
    uint24 immutable public fee;

    constructor() {
        (factory, token0, token1, fee) = IDEXPoolDeployer(msg.sender).parameters();
    }

    /**
     * @dev Returns the reserve balance of token0 in the pool.
     * @return The balance of token0 held by the contract.
    */
    function getReserve0() public view returns (uint256) {
        return IERC20(token0).balanceOf(address(this));
    }

    /**
     * @dev Returns the reserve balance of token1 in the pool.
     * @return The balance of token1 held by the contract.
    */
    function getReserve1() public view returns (uint256) {
        return IERC20(token1).balanceOf(address(this));
    }

    /**
     * @dev Returns the address of the factory that deployed this pool.
     * @return The factory address.
    */
    function getFactory() public view returns (address) {
        return factory;
    }

    /**
     * @dev Calculates the output amount of the opposite token for a given input amount.
     * @param amount The input amount of the token being swapped.
     * @param zeroToOne A boolean indicating the direction of the swap. True if swapping from token0 to token1, false otherwise.
     * @return The calculated output amount of the opposite token.
    */
    function getOutputAmount(
        uint256 amount, 
        bool zeroToOne
    ) public view returns (uint256) {
        require(amount > 0, InvalidInputAmount());
        (uint256 _reserve0, uint256 _reserve1) = zeroToOne 
            ? (getReserve0(), getReserve1()) 
            : (getReserve1(), getReserve0());

        // probably unused
        require(_reserve0 > 0 && _reserve1 > 0, InsufficientBalance());

        return _reserve1 * amount / (_reserve0 + amount);
    }

    /**
     * @dev Calculates the required amount of the opposite token to add liquidity in proportion.
     * @param amountIn The amount of the token being added.
     * @param zeroToOne A boolean indicating the direction. True if adding token0, false if adding token1.
     * @return The required amount of the opposite token to maintain the pool's balance.
    */
    function getAmountToAdd(
        uint256 amountIn, 
        bool zeroToOne
    ) public view returns (uint256) {
        (uint256 _reserve0, uint256 _reserve1) = zeroToOne 
            ? (getReserve0(), getReserve1())
            : (getReserve1(), getReserve0());

        uint256 requiredAmount = amountIn * _reserve1 / _reserve0;
        return requiredAmount;
    }

    // @example: slipperage: 100 tokenOut за 1 tokenIn, slip: 1%, amountOutMin = 99 tokenOut

    /**
     * @dev Executes a token swap between token0 and token1.
     * @param amountIn The amount of the input token to swap.
     * @param amountOutMin The minimum acceptable amount of the output token.
     * @param zeroToOne A boolean indicating the direction of the swap. True if swapping from token0 to token1, false otherwise.
    */
    function swap(
        uint256 amountIn, 
        uint256 amountOutMin, 
        bool zeroToOne
    ) external {
        uint256 amountOut = getOutputAmount(amountIn, zeroToOne);
        require(amountOut >= amountOutMin, InsufficientSlipperage());

        (address _token0, address _token1) = zeroToOne 
            ? (token0, token1) 
            : (token1, token0);

        IERC20(_token0).transferFrom(msg.sender, address(this), amountIn);
        IERC20(_token1).transfer(msg.sender, amountOut);

        emit SwapTokens(msg.sender, amountIn, amountOut, zeroToOne);
    }

    // @example: reserve0 = 100, reserve1 = 10
    // * wanna add: amount0In = 200
    // * requiredTokens = amount0In * reserve1 / reserve0 = 200 * 10 / 100 = 20 == amount1In

    /**
     * @dev Allows a user to add liquidity to the pool by providing both tokens.
     * @param amount0In The amount of token0 to add.
     * @param amount1In The amount of token1 to add.
    */
    function addLiquidity(
        uint256 amount0In,
        uint256 amount1In
    ) external {
        // we are sure that if token0 reserve is zero than token1 reserve is also zero
        if (getReserve0() == 0) {
            IERC20(token0).transferFrom(msg.sender, address(this), amount0In);
            IERC20(token1).transferFrom(msg.sender, address(this), amount1In);
        } else {
            // always considering that token0 is first token
            uint256 requiredAmount = getAmountToAdd(amount0In, true);        
            require(amount1In >= requiredAmount, WrongLiquidityAmout(requiredAmount, amount1In));

            IERC20(token0).transferFrom(msg.sender, address(this), amount0In);
            IERC20(token1).transferFrom(msg.sender, address(this), requiredAmount);
        }

        emit AddLiquidity(msg.sender, amount0In, amount1In);
    }
}