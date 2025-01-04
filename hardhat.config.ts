import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";

require('dotenv').config()

const ALCHEMY_RPC_URL = process.env.RPC_URL;
const PRIVATE_KEY = process.env.PRIVATE_KEY;

module.exports = {
  defaultNetwork: "sepolia",
  networks: {
    hardhat: {
      
    },
    sepolia: {
      url: ALCHEMY_RPC_URL, 
      accounts: [PRIVATE_KEY],
      saveDeployments: true,
    },
    // etherscan: {
    //   apiKey: {
    //     sepolia: ETHERSCAN_API_KEY,
    //   }
    // }
  },
  solidity: "0.8.28",
  paths: {
    sources: "./contracts",
    artifacts: "./artifacts",
  },
};

const config: HardhatUserConfig = {
  solidity: "0.8.28",
};

export default config;