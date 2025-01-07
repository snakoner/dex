export const poolABI = [
    "function getReserve0() public view returns (uint256)",
    "function getReserve1() public view returns (uint256)",
    "function getFactory() public view returns (address)",
    "function getOutputAmount(uint256 amount, uint256 inReserve,uint256 outReserve) public view returns (uint256)",
    "function getAmountToAdd(uint256 amountIn, bool zeroToOne) public view returns (uint256)",
    "function getAmountsFromLp(uint amount) public view returns (uint256, uint256)",
    "function swap(uint256 amountIn, uint256 amountOutMin, bool zeroToOne) external",
    "function addLiquidity(uint256 amount0In,uint256 amount1In) external",
    "function removeLiquidity(uint256 amount) external",
];

export const ercMintable20ABI = [
    "event Transfer(address indexed from, address indexed to, uint256 value)",
    "event Approval(address indexed owner, address indexed spender, uint256 value)",
    "function totalSupply() external view returns (uint256)",
    "function balanceOf(address account) external view returns (uint256)",
    "function transfer(address to, uint256 value) external returns (bool)",
    "function allowance(address owner, address spender) external view returns (uint256)",
    "function approve(address spender, uint256 value) external returns (bool)",
    "function transferFrom(address from, address to, uint256 value) external returns (bool)",
    "function mint(address account, uint256 value) external"
];

export const factoryABI = [
    "function createPool(address token0, address token1, string memory name, string memory symbol, uint24 fee) external returns (address)",
    "function getPool(address token0, address token1) external view returns (address)",
    "function getFee(address pool) external view returns (uint24)"
];