// This setup uses Hardhat Ignition to manage smart contract deployments.
// Learn more about it at https://hardhat.org/ignition

import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const FactoryModule = buildModule("FactoryModule", (m) => {
  const factory = m.contract("DEXFactory");

  return { factory };
});

export default FactoryModule;
