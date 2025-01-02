import { loadFixture, ethers, expect } from "./setup";
import { HardhatEthersSigner } from "@nomicfoundation/hardhat-ethers/src/signers";

describe("DEX test", function() {
        async function deploy() {
        const owner = (await ethers.getSigners())[0];
        const parts = (await ethers.getSigners()).slice(1, 10);
        let Factory = await ethers.getContractFactory("AToken");
        const aToken = await Factory.deploy();
        await aToken.waitForDeployment();

        Factory = await ethers.getContractFactory("BToken");
        const bToken = await Factory.deploy();
        await bToken.waitForDeployment();


        return {aToken, bToken, owner, parts};
    }

    it ("should be finished", async function() {
        const {aToken, bToken, owner, parts} = await loadFixture(deploy);        
    });
})
