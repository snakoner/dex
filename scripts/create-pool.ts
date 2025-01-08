const { ethers } = require("ethers");
import { readPools, Pools } from "./read-pools";
require('dotenv').config();

const ALCHEMY_RPC_URL = process.env.RPC_URL;
const PRIVATE_KEY = process.env.PRIVATE_KEY;

const { abi } = require("../artifacts/contracts/DEXFactory.sol/DEXFactory.json");
const pairNumber = 0;
const pools: Pools = readPools("./config/pools.json");
const factoryAddress = pools.factoryAddress;
const pair = {
    nameA: pools.pairs[pairNumber].nameA,
    nameB: pools.pairs[pairNumber].nameB,
    tokenA: pools.pairs[pairNumber].tokenA,
    tokenB: pools.pairs[pairNumber].tokenB,
    fee: 3,
    lpTokenName: "DexSwap PEPE/GRT LP Token",
    lpTokenSymbol: "DS-PEPE-GRT-LP",
};

async function createPool() {
    if (ALCHEMY_RPC_URL !== undefined && PRIVATE_KEY != undefined) {
        const provider = new ethers.JsonRpcProvider(ALCHEMY_RPC_URL);
        const wallet = new ethers.Wallet(PRIVATE_KEY, provider);
        const contract = new ethers.Contract(factoryAddress, abi, wallet);
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