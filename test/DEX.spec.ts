import { loadFixture, ethers, expect } from "./setup";

const poolABI = [
    "function getReserve0() public view returns (uint256)",
    "function getReserve1() public view returns (uint256)",
    "function getOutputAmount(uint256 amount, bool zeroToOne) public view returns (uint256)",
    "function getFactory() public view returns (address)",
    "function getAmountToAdd(uint256 amountIn, bool zeroToOne) public view returns (uint256)",
    "function swap(uint256 amountIn, uint256 amountOutMin, bool zeroToOne) external",
    "function addLiquidity(uint256 amount0In, uint256 amount1In) external",
];

describe("DEX test", function() {
    async function deploy() {
        const owner = (await ethers.getSigners())[0];
        const parts = (await ethers.getSigners()).slice(1, 10);
        
        // aToken
        const atokenFactory = await ethers.getContractFactory("AToken");
        const aToken = await atokenFactory.deploy();
        await aToken.waitForDeployment();
        
        // bToken
        const btokenFactory = await ethers.getContractFactory("BToken");
        const bToken = await btokenFactory.deploy();
        await bToken.waitForDeployment();

        // factory
        const factoryFactory = await ethers.getContractFactory("DEXFactory");
        const factory = await factoryFactory.deploy();
        await factory.waitForDeployment();

        return {factory, aToken, bToken, owner, parts};
    }

    it ("should be create pool", async function() {
        const {factory, aToken, bToken, owner, parts} = await loadFixture(deploy);        
        const fee = 2;

        // 0. deploy pool
        await factory.createPool(
            await aToken.getAddress(),
            await bToken.getAddress(),
            fee
        );
        const poolAddress = await factory.getPool(await aToken.getAddress(),
                                await bToken.getAddress());
        const poolContract = new ethers.Contract(poolAddress, poolABI, owner);

        expect(poolAddress).to.be.not.eq(ethers.ZeroAddress);

        // 1. add some liquidity
        const addLiquidityData = {
            amount0In: 10000000,
            amount1In: 200000,
        };

        await aToken.mint(owner.address, addLiquidityData.amount0In);
        await bToken.mint(owner.address, addLiquidityData.amount1In);

        expect(await aToken.totalSupply()).to.be.eq(addLiquidityData.amount0In);
        expect(await bToken.totalSupply()).to.be.eq(addLiquidityData.amount1In);

        expect(await aToken.balanceOf(owner.address)).to.be.eq(addLiquidityData.amount0In);
        expect(await bToken.balanceOf(owner.address)).to.be.eq(addLiquidityData.amount1In);

        await aToken.approve(await poolContract.getAddress(), addLiquidityData.amount0In);
        await bToken.approve(await poolContract.getAddress(), addLiquidityData.amount1In);
        await poolContract.addLiquidity(addLiquidityData.amount0In, addLiquidityData.amount1In);

        // 1. swap A -> B
        // add some tokens to owner
        const swapData = {
            amountIn: 40000,
            amountOutMin: 0,    // calc later
            slipperage: 1,
        };

        await aToken.mint(owner.address, swapData.amountIn);

        expect(await poolContract.getReserve0()).to.be.eq(addLiquidityData.amount0In);
        expect(await poolContract.getReserve1()).to.be.eq(addLiquidityData.amount1In);

        const price = Number(await poolContract.getReserve0()) / Number(await poolContract.getReserve1());
        const minOut = Math.floor(swapData.amountIn / price * (1 - swapData.slipperage / 100.));
        swapData.amountOutMin = minOut;

        console.log(price);
        console.log(minOut);
        console.log("A/B price: ", (await poolContract.getReserve0()) / (await poolContract.getReserve1()));
        console.log("B/A price: ", Number(await poolContract.getReserve1()) / Number(await poolContract.getReserve0()));

        let swappedBAmount = await poolContract.getOutputAmount(swapData.amountIn, true);
        console.log("will receive about A->B: ", swappedBAmount);

        await aToken.approve(poolAddress, swapData.amountIn);

        // 2. swap
        await poolContract.swap(
            ethers.toBigInt(swapData.amountIn), 
            ethers.toBigInt(swapData.amountOutMin), 
            true
        );

        console.log("owner balance A: ", await aToken.balanceOf(owner.address));
        console.log("owner balance B: ", await bToken.balanceOf(owner.address));

        expect(await aToken.balanceOf(owner.address)).to.be.eq(0);
        expect(await bToken.balanceOf(owner.address)).to.be.eq(swappedBAmount);
    });
})
