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
    "function mint(address account, uint256 amount) external",
    "function totalSupply() public view returns (uint256)"
];

const lpTokenName = "DexSwap ATK/BTK LP Token";
const lpTokenSymbol = "DS-ATK-BTK-LP";
const fee = 2;      // 0.2, max = 10e3
const mint0Amount = 1000000000;
const mint1Amount = 20000000;

describe("DEX test", function() {
    async function deploy() {
        const owner = (await ethers.getSigners())[0];
        const account = (await ethers.getSigners())[1];
        
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

        await factory.createPool(
            await aToken.getAddress(),
            await bToken.getAddress(),
            lpTokenName,
            lpTokenSymbol,
            fee
        );

        const poolAddress = await factory.getPool(await aToken.getAddress(), await bToken.getAddress());
        const pool = new ethers.Contract(poolAddress, poolABI, owner);

        const lpTokenAddress = await pool.lpToken();
        const lpToken = new ethers.Contract(lpTokenAddress, erc20ABI, ethers.provider);

        // for future use mint
        await aToken.mint(owner.address, mint0Amount);
        await bToken.mint(owner.address, mint1Amount);

        // for future use approve all
        await aToken.approve(await pool.getAddress(), mint0Amount);
        await bToken.approve(await pool.getAddress(), mint1Amount);

        return {factory, aToken, bToken, pool, lpToken, owner, account};
    }

    it ("Should not allow to mint/burn lp tokens by any except pool", async function() {
        const {lpToken, owner} = await loadFixture(deploy);        

        try {
            lpToken.mint(owner.address, 10000);
            expect(false).to.be.eq(true);
        } catch (error) {}
    });

    it ("Should be non null address contracts and A,B tokens shooud be minted", async function() {
        const {factory, aToken, bToken, pool, lpToken, owner} = await loadFixture(deploy);        

        expect(await factory.getAddress()).not.to.be.eq(ethers.ZeroAddress);
        expect(await aToken.getAddress()).not.to.be.eq(ethers.ZeroAddress);
        expect(await bToken.getAddress()).not.to.be.eq(ethers.ZeroAddress);
        expect(await pool.getAddress()).not.to.be.eq(ethers.ZeroAddress);
        expect(await lpToken.getAddress()).not.to.be.eq(ethers.ZeroAddress);

        expect(await aToken.totalSupply()).to.be.eq(mint0Amount);
        expect(await bToken.totalSupply()).to.be.eq(mint1Amount);

        expect(await aToken.balanceOf(owner.address)).to.be.eq(mint0Amount);
        expect(await bToken.balanceOf(owner.address)).to.be.eq(mint1Amount);
    });

    it ("Should not allow to mint/burn lp tokens by any except pool", async function() {
        const {lpToken, owner} = await loadFixture(deploy);        

        try {
            lpToken.mint(owner.address, 10000);
            expect(false).to.be.eq(true);
        } catch (error) {

        }
    });

    it ("Should be possible to add liquidity in any proportions", async function() {
        const {pool} = await loadFixture(deploy);        

        const addLiquidityData = {
            amount0In: 10000000,
            amount1In: 1000000,
        };

        await pool.addLiquidity(addLiquidityData.amount0In, addLiquidityData.amount1In);

        expect(await pool.getReserve0()).to.be.eq(addLiquidityData.amount0In);
        expect(await pool.getReserve1()).to.be.eq(addLiquidityData.amount1In);
    });

    it ("Should create lp tokens with same amount as tokenA", async function() {
        const {pool, lpToken, owner} = await loadFixture(deploy);        

        expect(await lpToken.balanceOf(owner.address)).to.be.eq(0);

        const addLiquidityData = {
            amount0In: 10000000,
            amount1In: 1000000,
        };

        await pool.addLiquidity(addLiquidityData.amount0In, addLiquidityData.amount1In);

        expect(await lpToken.totalSupply()).to.be.eq(addLiquidityData.amount0In);
    });


    it ("Should be not possible to add liquidity in any proportions if pool not empty", async function() {
        const {pool} = await loadFixture(deploy);        

        // 1+. add some more liquidity
        const addLiquidityData = {
            amount0In: 1000,
            amount1In: 1000, // calculate later
        };

        try {
            await pool.addLiquidity(addLiquidityData.amount0In, addLiquidityData.amount1In);
            expect(false).to.be.eq(true);
        } catch (error) {

        }
    });

    it ("Should be possible to add liquidity in right proportions if pool not empty", async function() {
        const {pool} = await loadFixture(deploy);        

        const addLiquidity1Data = {
            amount0In: 10000000,
            amount1In: 1000000,
        };

        const addLiquidity2Data = {
            amount0In: 10000,
            amount1In: 0, // calc later
        };

        await pool.addLiquidity(addLiquidity1Data.amount0In, addLiquidity1Data.amount1In);

        addLiquidity2Data.amount1In = await pool.getAmountToAdd(addLiquidity2Data.amount0In, true);
        await pool.addLiquidity(addLiquidity2Data.amount0In, addLiquidity2Data.amount1In);

        expect(await pool.getReserve0()).to.be.eq(
            ethers.toBigInt(addLiquidity1Data.amount0In) + 
            ethers.toBigInt(addLiquidity2Data.amount0In));
        expect(await pool.getReserve1()).to.be.eq(
            ethers.toBigInt(addLiquidity1Data.amount1In) + 
            ethers.toBigInt(addLiquidity2Data.amount1In));
    });

    it ("Should swap small amount with small slippage", async function() {
        const {pool} = await loadFixture(deploy);        

        // 1. swap A -> B
        const swapData = {
            amountIn: 10000,    // 10k
            amountOutMin: 0,    // calc later
            slippage: 10,       // 1%
        };

        const addLiquidity1Data = {
            amount0In: 10000000,    // 10kk is 1000 bigger than swapData.amountIn
            amount1In: 1000000,
        };

        await pool.addLiquidity(
            addLiquidity1Data.amount0In, 
            addLiquidity1Data.amount1In
        );

        const price = Number(await pool.getReserve0()) / Number(await pool.getReserve1());
        swapData.amountOutMin = Math.floor((swapData.amountIn / price * (1 - (swapData.slippage + fee) / 1000.)));

        await pool.swap(
            ethers.toBigInt(swapData.amountIn), 
            ethers.toBigInt(swapData.amountOutMin), 
            true
        );
    });
    

    it ("Should not swap large amount with small slippage", async function() {
        const {pool} = await loadFixture(deploy);        

        // 1. swap A -> B
        const swapData = {
            amountIn: 1000000,    // 1kk
            amountOutMin: 0,    // calc later
            slippage: 10,       // 1%
        };

        const addLiquidity1Data = {
            amount0In: 10000000,    // 10kk is only 10 times bigger than swapData.amountIn
            amount1In: 1000000,
        };

        await pool.addLiquidity(
            addLiquidity1Data.amount0In, 
            addLiquidity1Data.amount1In
        );

        const price = Number(await pool.getReserve0()) / Number(await pool.getReserve1());
        swapData.amountOutMin = Math.floor((swapData.amountIn / price * (1 - (swapData.slippage + fee) / 1000.)));

        try {
            await pool.swap(
                ethers.toBigInt(swapData.amountIn), 
                ethers.toBigInt(swapData.amountOutMin), 
                true
            );
            expect(false).to.be.eq(true);
        } catch (error) {}
    });
    
    it ("Should swap large amount with large slippage", async function() {
        const {pool} = await loadFixture(deploy);        

        // 1. swap A -> B
        const swapData = {
            amountIn: 1000000,    // 1kk
            amountOutMin: 0,    // calc later
            slippage: 120,       // 12%
        };

        const addLiquidity1Data = {
            amount0In: 10000000,    // 10kk is 10 bigger than swapData.amountIn
            amount1In: 1000000,
        };

        await pool.addLiquidity(
            addLiquidity1Data.amount0In, 
            addLiquidity1Data.amount1In
        );

        const price = Number(await pool.getReserve0()) / Number(await pool.getReserve1());
        swapData.amountOutMin = Math.floor((swapData.amountIn / price * (1 - (swapData.slippage + fee) / 1000.)));

        await pool.swap(
            ethers.toBigInt(swapData.amountIn), 
            ethers.toBigInt(swapData.amountOutMin), 
            true
        );
    });

    it ("Should be possible to remove liquidity", async function() {
        const {aToken, bToken, pool, lpToken, owner} = await loadFixture(deploy);

        const addLiquidityData = {
            amount0In: 10000000, 
            amount1In: 1000000,
        };

        await pool.addLiquidity(
            addLiquidityData.amount0In, 
            addLiquidityData.amount1In
        );

        const ownerATokenBalanceBefore = await aToken.balanceOf(owner.address);
        const ownerBTokenBalanceBefore = await bToken.balanceOf(owner.address);

        await pool.removeLiquidity(await lpToken.balanceOf(owner.address));

        // make sure that there is not liquidity in pool
        expect(await lpToken.balanceOf(owner.address)).to.be.eq(0);
        expect(await lpToken.totalSupply()).to.be.eq(0);
        expect(await pool.getReserve0()).to.be.eq(0);
        expect(await pool.getReserve1()).to.be.eq(0);

        const ownerATokenBalanceAfter = await aToken.balanceOf(owner.address);
        const ownerBTokenBalanceAfter = await bToken.balanceOf(owner.address);

        expect(ownerATokenBalanceAfter - ownerATokenBalanceBefore).to.be.eq(addLiquidityData.amount0In);
        expect(ownerBTokenBalanceAfter - ownerBTokenBalanceBefore).to.be.eq(addLiquidityData.amount1In);

    });
})
