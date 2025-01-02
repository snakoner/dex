// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import { DEXPool } from "./DEXPool.sol";
import { IDEXPoolDeployer } from "./interfaces/IDEXPoolDeployer.sol";

contract DEXPoolDeployer is IDEXPoolDeployer {
    struct Parameters {
        address factory;
        address token0;
        address token1;
        uint24 fee;
    }

    Parameters public parameters;

    function deploy(
        address factory,
        address token0,
        address token1,
        uint24 fee
    ) internal returns (address) {
        parameters = Parameters({
            factory: factory,
            token0: token0,
            token1: token1,
            fee: fee
        });

        address pool = address(new DEXPool{salt: keccak256(abi.encode(token0, token1, fee))}());
        delete parameters;
        return pool;
    } 
}