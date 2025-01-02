// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import { DEXPool } from "./DEXPool.sol";
import { IDEXPoolDeployer } from "./interfaces/IDEXPoolDeployer.sol";
import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import { ERC20Lp } from "./ERC20Lp.sol";

contract DEXPoolDeployer is IDEXPoolDeployer {
    struct Parameters {
        address factory;
        address token0;
        address token1;
        address lpToken;
        uint24 fee;
    }

    Parameters public override parameters;

    function deploy(
        address factory,
        address token0,
        address token1,
        string memory name,
        string memory symbol,
        uint24 fee
    ) internal returns (address) {
        address lpToken = address(new ERC20Lp(name, symbol, address(this)));

        parameters = Parameters({
            factory: factory,
            token0: token0,
            token1: token1,
            lpToken: lpToken,
            fee: fee
        });

        address pool = address(new DEXPool{salt: keccak256(abi.encode(token0, token1, fee))}());
        ERC20Lp(lpToken).transferOwnership(pool);

        delete parameters;

        return pool;
    } 
}