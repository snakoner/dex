// This setup uses Hardhat Ignition to manage smart contract deployments.
// Learn more about it at https://hardhat.org/ignition

import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const TokensModule = buildModule("TokensModule", (m) => {

  const grt = m.contract("GRT");
  const pepe = m.contract("PEPE");

  return { grt, pepe };
});

export default TokensModule;