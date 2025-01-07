const { ethers } = require("ethers");
require('dotenv').config();
import { ercMintable20ABI } from "./abi";

const ALCHEMY_RPC_URL = process.env.RPC_URL;
const PRIVATE_KEY = process.env.PRIVATE_KEY;


async function mintTokens() {
    const tokenAddress = "0xc40D76aD1130fD9ec53a58E54F12aA6e088c1527"; // 0x4104993a0a0F0F81eC76077E8CAaC5137b2556e6 - PEPE
    const to = "0x986D35f14A58d4A23787c7C1F17ebd4c398f0f75";
    const amount = ethers.toBigInt(1000000) * ethers.parseEther("1.0");

    if (ALCHEMY_RPC_URL !== undefined && PRIVATE_KEY != undefined) {
        const provider = new ethers.JsonRpcProvider(ALCHEMY_RPC_URL);
        const wallet = new ethers.Wallet(PRIVATE_KEY, provider);
        const contract = new ethers.Contract(tokenAddress, ercMintable20ABI, wallet);
        const result = await contract.mint(
            to,
            amount
        );  
    
        console.log(result);    
    }
}

mintTokens().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});