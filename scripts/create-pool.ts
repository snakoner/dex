const { ethers } = require("ethers");
import { readPools, Pools } from "./read-pools";
require('dotenv').config();

const ALCHEMY_RPC_URL = process.env.RPC_URL;
const PRIVATE_KEY = process.env.PRIVATE_KEY;
const pools: Pools = readPools("./config/pools.json");
const factoryABI = [
    "function createPool(address token0, address token1, string memory name, string memory symbol, uint24 fee) external returns (address)",
    "function getPool(address token0, address token1) external view returns (address)"
];
const pair = {
    nameA: pools.pairs[0].nameA,
    nameB: pools.pairs[0].nameB,
    tokenA: pools.pairs[0].tokenA,
    tokenB: pools.pairs[0].tokenB,
    fee: 3,
    lpTokenName: "DexSwap PEPE/GRT LP Token",
    lpTokenSymbol: "DS-PEPE-GRT-LP",
};
const factoryAddress = pools.factoryAddress;

async function createPool() {
    if (ALCHEMY_RPC_URL !== undefined && PRIVATE_KEY != undefined) {
        const provider = new ethers.JsonRpcProvider(ALCHEMY_RPC_URL);
        const wallet = new ethers.Wallet(PRIVATE_KEY, provider);
        const contract = new ethers.Contract(factoryAddress, factoryABI, wallet);
        const result = await contract.createPool(
            pair.tokenA, 
            pair.tokenB,
            pair.lpTokenName,
            pair.lpTokenSymbol,
            pair.fee
        );  
    
        console.log(result);    
    }
}

createPool().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});