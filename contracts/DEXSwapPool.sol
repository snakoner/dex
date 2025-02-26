// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/utils/math/Math.sol";
import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { IDEXSwapPoolDeployer } from "./interfaces/IDEXSwapPoolDeployer.sol";
import { IDEXSwapPool } from "./interfaces/IDEXSwapPool.sol";
import { DEXSwapLiquidityProviderERC20 } from "./DEXSwapLiquidityProviderERC20.sol";

/*
 * explanation:
 * we want to exchange dx of token0:
 ************************************
 * (x + dx)(y - dy) = xy        (1) 
 * y - dy = xy / (x + dx)       (2)
 * dy = y - xy / (x + dx) = (xy + y * dx - xy) / (x + dx) = y * dx / (x + dx)       (3)
*/

contract DEXSwapPool is IDEXSwapPool {
    address immutable public factory;
    address immutable public token0;
    address immutable public token1;
    address immutable public lpToken;
    uint24 public fee;    // 10e4
    uint48 public constant PRECISION = 10**4;

    constructor() {
        (factory, token0, token1, lpToken, fee) = IDEXSwapPoolDeployer(msg.sender).parameters();

        emit FeeUpdated(0, fee);
    }

    modifier onlyFactory() {
        require(msg.sender == factory, OnlyFactoryAllowed());
        _;
    }


    function getReserves() public view returns (uint256, uint256) {
        return (
            IERC20(token0).balanceOf(address(this)),
            IERC20(token1).balanceOf(address(this))
        );
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
     * @param inReserve The input reserve of the token being swapped.
     * @param outReserve The output reserve of the token being swapped.
     * @return The calculated output amount of the opposite token.
    */
    function getOutputAmount(
        uint256 amount, 
        uint256 inReserve,
        uint256 outReserve
    ) public view returns (uint256) {
        uint256 amountWithFee = amount * (PRECISION - fee);
        uint256 outAmount = Math.mulDiv(amountWithFee, outReserve, (inReserve * PRECISION) + amountWithFee);

        return outAmount;
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
        (uint256 _reserve0, uint256 _reserve1) = getReserves();
        if (!zeroToOne) {
            (_reserve0, _reserve1) = (_reserve1, _reserve0);
        }

        return Math.mulDiv(amountIn, _reserve1, _reserve0 + amountIn, Math.Rounding.Ceil);
    }

    /**
     * @dev Calculates the amounts of token0 and token1 corresponding to a given amount of LP tokens.
     * @param amount The amount of LP tokens.
     * @return token0Amount The amount of token0 corresponding to the LP tokens.
     * @return token1Amount The amount of token1 corresponding to the LP tokens.
     */
    function getAmountsFromLp(uint amount) public view returns (uint256, uint256) {
        uint256 _totalSupply = DEXSwapLiquidityProviderERC20(lpToken).totalSupply();
        (uint256 _reserve0, uint256 _reserve1) = getReserves();

        uint256 token0Amount = Math.mulDiv(_reserve0, amount, _totalSupply);
        uint256 token1Amount = Math.mulDiv(_reserve1, amount, _totalSupply);

        return (token0Amount, token1Amount);
    }

    // @example: Slippage: 100 tokenOut за 1 tokenIn, slip: 1%, amountOutMin = 99 tokenOut

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
        require(amountIn > 0, InvalidInputAmount());

        (uint256 _reserve0, uint256 _reserve1) = getReserves();
        if (!zeroToOne) {
            (_reserve0, _reserve1) = (_reserve1, _reserve0);
        }

        uint256 amountOut = getOutputAmount(amountIn, _reserve0, _reserve1);
        require(amountOut >= amountOutMin, InsufficientSlippageSet(amountOutMin, amountOut));

        address account = msg.sender;
        (address _token0, address _token1) = zeroToOne 
            ? (token0, token1) 
            : (token1, token0);

        IERC20(_token0).transferFrom(account, address(this), amountIn);
        IERC20(_token1).transfer(account, amountOut);

        emit SwapTokens(account, amountIn, amountOut, zeroToOne);
    }

    // @example: reserve0 = 100, reserve1 = 10
    // * wanna add: amount0In = 200
    // * requiredTokens = amount0In * reserve1 / reserve0 = 200 * 10 / 100 = 20 == amount1In
    // * lp tokens: LP_issued = Total_LP * (ΔL / Total_L)
    /**
     * @dev Allows a user to add liquidity to the pool by providing both tokens.
     * @param amount0In The amount of token0 to add.
     * @param amount1In The amount of token1 to add.
    */
    function addLiquidity(
        uint256 amount0In,
        uint256 amount1In
    ) external {
        address account = msg.sender;
        (uint256 _reserve0, ) = getReserves();
        // we are sure that if token0 reserve is zero than token1 reserve is also zero
        if (_reserve0 == 0) {
            IERC20(token0).transferFrom(account, address(this), amount0In);
            IERC20(token1).transferFrom(account, address(this), amount1In);

            DEXSwapLiquidityProviderERC20(lpToken).mint(account, amount0In);
        } else {
            // always considering that token0 is first token
            uint256 requiredAmount = getAmountToAdd(amount0In, true);        
            require(amount1In >= requiredAmount, WrongLiquidityAmout(requiredAmount, amount1In));

            IERC20(token0).transferFrom(account, address(this), amount0In);
            IERC20(token1).transferFrom(account, address(this), requiredAmount);
            
            // mint LP tokens 
            DEXSwapLiquidityProviderERC20 lp = DEXSwapLiquidityProviderERC20(lpToken);
            (_reserve0, ) = getReserves();
            uint256 lpIssued = (lp.totalSupply() * amount0In) / _reserve0;

            lp.mint(account, lpIssued);
        }

        emit AddLiquidity(account, amount0In, amount1In);
    }

    /**
     * @dev Allows a user to remove liquidity from the pool by burning LP tokens.
     * @param amount The amount of LP tokens to burn.
     */
    function removeLiquidity(
        uint256 amount // lp token amount
    ) external {
        require(amount > 0, InvalidInputAmount());

        address account = msg.sender;

        (uint256 token0Amount, uint256 token1Amount) = getAmountsFromLp(amount);
        DEXSwapLiquidityProviderERC20 _lpToken = DEXSwapLiquidityProviderERC20(lpToken);

        _lpToken.burn(account, amount);
        IERC20(token0).transfer(account, token0Amount);
        IERC20(token1).transfer(account, token1Amount);

        emit RemoveLiquidity(account, amount);
    }

    function updateFee(uint24 _fee) external onlyFactory {
        emit FeeUpdated(fee, _fee);
        fee = _fee;
    }
}