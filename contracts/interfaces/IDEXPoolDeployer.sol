// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

interface IDEXPoolDeployer {
        function parameters()
        external
        view
        returns (
            address factory,
            address token0,
            address token1,
            uint24 fee
        );

}