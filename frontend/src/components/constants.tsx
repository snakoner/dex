export const ALCHEMY_RPC_URL = 'https://eth-sepolia.g.alchemy.com/v2/UtjFRzFoEQd533NSskUCCCKEpW7z93t2';
export const CONTRACT_ADDRESS = '0x87Ec2437F2A3451385dbC31FdcfE40a1B1814E7F';
export const CONTRACT_ABI = [
    "event Bid(address indexed account, uint amount, uint timestamp, uint indexed round)",
    "event WinnerSelected(address indexed account, uint amount, uint indexed round)",
    "event Withdraw(address indexed account, uint amount)",
    "event Deposit(address indexed user, uint amount, uint timestamp)",
    "function balances(address) public view returns (uint256)",
    "function weights(uint256, address) public view returns (uint256)",
    "function totalWeight(uint256) public view returns (uint256)",
    "function participants(uint256) public view returns (address)",
    "function bid(uint amount) external payable",
    "function withdraw(uint amount) external",
    "function getParticipantsNumber() external view returns (uint)",
    "function getTimeLeft() external view returns (uint)",
    "function bidFromBalance(uint amount) external",
    "function ticketPrice() public view returns (uint256)",
    "function round() public view returns (uint256)",
    "function deposit() external payable"
];

export const POOL_ABI = [
    "function getReserve0() public view returns (uint256)",
    "function getReserve1() public view returns (uint256)",
    "function getOutputAmount(uint256 amount, uint256 inReserve,uint256 outReserve) public view returns (uint256)",
    "function getAmountToAdd(uint256 amountIn, bool zeroToOne) public view returns (uint256)"
];

export const ERC20_ABI = [
    "function balanceOf(address account) external view returns (uint256)",
    "function transfer(address to, uint256 value) external returns (bool)",
    "function allowance(address owner, address spender) external view returns (uint256)",
    "function approve(address spender, uint256 value) external returns (bool)",
    "function transferFrom(address from, address to, uint256 value) external returns (bool)",
    "function totalSupply() external view returns (uint256)",
    "function decimals() external view returns (uint8)",
];

export const icons = new Map<string, string>([
    ["PEPE", "https://cryptologos.cc/logos/pepe-pepe-logo.png?v=040"],
    ["GRT", "https://cryptologos.cc/logos/the-graph-grt-logo.png?v=040"],
]);

export const localStorageWalletConnectHandler = () => {
    if (localStorage.getItem('walletConnected') === undefined) {
        localStorage.setItem('walletConnected', 'false');
    }

    return localStorage.getItem('walletConnected') === 'true' ? true : false;
}

export const supportedChains = [
    {
        network: "sepolia",
        chainId: 11155111,
    },
];