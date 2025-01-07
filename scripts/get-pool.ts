const { ethers } = require("ethers");
import { readPools, Pools } from "./read-pools";
require('dotenv').config();
import { factoryABI } from "./abi";

const ALCHEMY_RPC_URL = process.env.RPC_URL;

async function getPool() {
    const pools: Pools = readPools("./config/pools.json");
    const factoryAddress: string = pools.factoryAddress;

    if (ALCHEMY_RPC_URL !== undefined && factoryAddress !== undefined) {
        const provider = new ethers.JsonRpcProvider(ALCHEMY_RPC_URL);
        const contract = new ethers.Contract(factoryAddress, factoryABI, provider);
        const result = await contract.getPool(
            pools.pairs[0].tokenA, 
            pools.pairs[0].tokenB,
        );  
    
        console.log("pool address = ", result);    
    }
}

getPool().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});