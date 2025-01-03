import { loadFixture, ethers, expect } from "./setup";

const poolABI = [
    "function getReserve0() public view returns (uint256)",
    "function getReserve1() public view returns (uint256)",
    "function getFactory() public view returns (address)",
    "function getAmountToAdd(uint256 amountIn, bool zeroToOne) public view returns (uint256)",
    "function getAmountsFromLp(uint amount) public view returns (uint256, uint256)",
    "function swap(uint256 amountIn, uint256 amountOutMin, bool zeroToOne) external",
    "function addLiquidity(uint256 amount0In, uint256 amount1In) external",
    "function removeLiquidity(uint256 amount) external",
    "function getOutputAmount(uint256 amount, uint256 inReserve, uint256 outReserve) public view returns (uint256)",
    "function lpToken() public view returns (address)",
    
];

const erc20ABI = [
    "function balanceOf(address account) external view returns (uint256)",
    "function owner() public view returns (address)",
    "function mint(address account, uint256 amount) external"
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
            "AToken-BToken-LP",
            "LP-ATKBTK",
            fee
        );
        const poolAddress = await factory.getPool(await aToken.getAddress(),
                                await bToken.getAddress());
        const poolContract = new ethers.Contract(poolAddress, poolABI, owner);
        const lpTokenAddress = await poolContract.lpToken();

        const lpToken = new ethers.Contract(lpTokenAddress, erc20ABI, ethers.provider);

        // make sure that owner of lp token is factory and nobody cant mint and burn tokens
        expect(await lpToken.owner()).to.be.eq(poolAddress);
        try {
            lpToken.mint(owner.address, 10000);
            expect(false).to.be.eq(true);
        } catch (error) {}


        expect(poolAddress).to.be.not.eq(ethers.ZeroAddress);

        console.log(await lpToken.balanceOf(owner.address));

        console.log("aToken address = ", await aToken.getAddress());
        console.log("bToken address = ", await bToken.getAddress());
        console.log("owner address = ", await owner.getAddress());
        console.log("pool address = ", await poolContract.getAddress());

        // 1. add some liquidity
        const addLiquidityData = {
            mint0Value: 1000000000,
            mint1Value: 20000000,
            amount0In: 10000000,
            amount1In: 10000000,
        };

        await aToken.mint(owner.address, addLiquidityData.mint0Value);
        await bToken.mint(owner.address, addLiquidityData.mint1Value);

        expect(await aToken.totalSupply()).to.be.eq(addLiquidityData.mint0Value);
        expect(await bToken.totalSupply()).to.be.eq(addLiquidityData.mint1Value);

        expect(await aToken.balanceOf(owner.address)).to.be.eq(addLiquidityData.mint0Value);
        expect(await bToken.balanceOf(owner.address)).to.be.eq(addLiquidityData.mint1Value);

        // for future use approve all
        await aToken.approve(await poolContract.getAddress(), addLiquidityData.mint0Value);
        await bToken.approve(await poolContract.getAddress(), addLiquidityData.mint1Value);

        // add
        await poolContract.addLiquidity(addLiquidityData.amount0In, addLiquidityData.amount1In);

        console.log('lpToken: ', await lpToken.balanceOf(owner.address));

        // 1+. add some more liquidity
        const addLiquidity2Data = {
            amount0In: 1000,
            amount1In: 0, // calculate later
        };

        addLiquidity2Data.amount1In = await poolContract.getAmountToAdd(addLiquidity2Data.amount0In, true);

        console.log(addLiquidity2Data);

        await poolContract.addLiquidity(addLiquidity2Data.amount0In, addLiquidity2Data.amount1In);

        console.log('lpToken: ', await lpToken.balanceOf(owner.address));

        // 1. swap A -> B
        // add some tokens to owner
        const swapData = {
            amountIn: 1000,
            amountOutMin: 0,    // calc later
            slippage: 10,
        };

        await aToken.mint(owner.address, swapData.amountIn);

        expect(await poolContract.getReserve0()).to.be.eq(ethers.toBigInt(addLiquidityData.amount0In + addLiquidity2Data.amount0In));
        expect(await poolContract.getReserve1()).to.be.eq(ethers.toBigInt(addLiquidityData.amount1In) + ethers.toBigInt(addLiquidity2Data.amount1In));

        const price = Number(await poolContract.getReserve0()) / Number(await poolContract.getReserve1());
        const minOut = Math.floor((swapData.amountIn / price * (1 - (swapData.slippage + fee) / 1000.)));
        swapData.amountOutMin = minOut;

        console.log(price);
        console.log("minOut = ", minOut);
        console.log("A/B price: ", (await poolContract.getReserve0()) / (await poolContract.getReserve1()));
        console.log("B/A price: ", Number(await poolContract.getReserve1()) / Number(await poolContract.getReserve0()));

        const inReserve = await poolContract.getReserve0();
        const outReserve = await poolContract.getReserve1();
        let swappedBAmount = await poolContract.getOutputAmount(swapData.amountIn, inReserve, outReserve);
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

        console.log("reserve A: ", await poolContract.getReserve0());
        console.log("reserve B: ", await poolContract.getReserve1());


        const lpTokensAmount = await lpToken.balanceOf(owner.address);
        console.log(await poolContract.getAmountsFromLp(lpTokensAmount));

        // remove full liq
        console.log("!!!! remove liquidity");

        console.log("lpTokensAmount = ", lpTokensAmount);
        await poolContract.removeLiquidity(lpTokensAmount - ethers.toBigInt(100000));


        console.log("reserve A: ", await poolContract.getReserve0());
        console.log("reserve B: ", await poolContract.getReserve1());
    });
})
