import React, { useEffect, useState } from "react";
import Header from "./dex/Header";
import SwapCard from "./dex/SwapCard";
import Features from "./dex/Features";
import Footer from "./footer";
import { ethers } from "ethers";

import { ALCHEMY_RPC_URL, ERC20_ABI, POOL_ABI, localStorageWalletConnectHandler, supportedChains, icons } from "./constants";
import LiquidityCard from "./dex/LiquidityCard";

const convertSupportedChains = (): Map<bigint, ethers.Network> => {
    let map = new Map<bigint, ethers.Network>();

    for (const suppChains of supportedChains) {
        map.set(
            ethers.toBigInt(suppChains.chainId), 
            new ethers.Network(suppChains.network, suppChains.chainId)
        );
    }

    return map;
}

interface HomeProps {
  isWalletConnected?: boolean;
  walletAddress?: string;
  networkName?: string;
  onWalletConnect?: () => void;
  onWalletDisconnect?: () => void;
  onNetworkChange?: (network: string) => void;
  onSwap?: () => void;
}

export interface Token {
	name: string;
	address: string;
	totalSupply: bigint;
	decimals: bigint;
	icon?: string;
	balance?: bigint;
};

interface Pair {
  nameA: string;
  nameB: string;
  pool: string;
  tokenA: string;
  tokenB: string;
  tokenLP: string;
};

const Home = ({
  networkName = "Sepolia Testnet",
  onNetworkChange = () => {},
  onSwap = () => {},
}: HomeProps) => {
  	const [pairs, setPairs] = useState<Map<string, Pair>>();
  	const [factoryAddress, setFactoryAddress] = useState<string|null>();
	const [walletError, setWalletError] = useState<string|null>(null);
	const [connected, setConnected] = useState<boolean>(false);
  	const [account, setAccount] = useState<string|null>(null);
	const [network, setNetwork] = useState<ethers.Network|null>(null);
	const [outputAmount, setOutputAmount] = useState<number>(0);
	// swap table
	const [token0, setToken0] = useState<Token>(null);
	const [token1, setToken1] = useState<Token>(null);
	const [isSwapTokensCard, setIsSwapTokensCard] = useState<boolean>(true);

	const provider = new ethers.JsonRpcProvider(ALCHEMY_RPC_URL);

	const onInputAmountChange = async(name: string, amount: number) => {
		let tokenFrom: Token = null;
		let tokenTo: Token = null;

		if (amount == 0) {
			return;
		}

		if (name == token0.name) {
			tokenFrom = token0;
			tokenTo = token1;
		} else {
			tokenFrom = token1;
			tokenTo = token0;
		}

		// get pair name
		let pairName: string;
		if (tokenFrom.address < tokenTo.address)
			pairName = token0.name + '-' + token1.name;
		else
			pairName = token1.name + '-' + token0.name;

		try {
			const poolContract = new ethers.Contract(pairs.get(pairName).pool, POOL_ABI, provider);
			const reserve0 = await poolContract.getReserve0();
			const reserve1 = await poolContract.getReserve0();
			const outputAmount = await poolContract.getOutputAmount(amount, reserve0, reserve1);
			
			setOutputAmount(Number(outputAmount));
		} catch (error) {
			console.log(error);
		}
	}

	const setTokens = async(pairs: Pair[]) => {
		// if no pairs received
		if (pairs && pairs.length == 0) {
			return;
		}

		// if wallet is connected than can set balance of wallet
		const _token0: Token = {
			name: pairs[0].nameA,
			address: pairs[0].tokenA,
			icon: icons.has(pairs[0].nameA) ? icons.get(pairs[0].nameA) : "",
		};

		const _token1: Token = {
			name: pairs[0].nameB,
			address: pairs[0].tokenB,
			icon: icons.has(pairs[0].nameB) ? icons.get(pairs[0].nameB) : "",
		};

		const contract0 = new ethers.Contract(_token0.address, ERC20_ABI, provider);
		const contract1 = new ethers.Contract(_token1.address, ERC20_ABI, provider);

		try {
			_token0.totalSupply = await contract0.totalSupply();
			_token0.decimals = await contract0.decimals();
			_token1.totalSupply = await contract1.totalSupply();
			_token1.decimals = await contract1.decimals();
		} catch (error) {
			console.log(error);
			return;
		}

		if (localStorageWalletConnectHandler()) {
			try {
				console.log("account: ", account);
				_token0.balance = await contract0.balanceOf(account);
				_token1.balance = await contract0.balanceOf(account);
			} catch (error) {
				console.log(error);
				return;
			}
		}

		setToken0(_token0);
		setToken1(_token1);

		console.log(_token0);
		console.log(_token1);
	}

	const onDirectionSwap =() => {
		setToken0(token1);
		setToken1(token0);
	}

  	const getPools = async() => {
		try {
			const response = await fetch("http://localhost:8000/get-pools", {
				method: 'GET',
				headers: {
					'Content-Type': 'application/json',
				},     
			});


		const data = await response.json();
		setFactoryAddress(data?.factoryAddress);

		// set pairs
		let _pairs: Pair[] = [];
		let pairsHm = new Map<string, Pair>();
		for (const p of data?.pairs) {
			const _pair: Pair = {
				nameA: p.nameA,
				nameB: p.nameB,
				tokenA: p.tokenA,
				tokenB: p.tokenB,
				tokenLP: p.tokenLP,
				pool: p.pool,
			};

			let pairName: string;
			if (p.tokenA < p.tokenB) {
				pairName = p.nameA + '-' + p.nameB;
			} else {
				pairName = p.nameB + '-' + p.nameA;
			}

			pairsHm.set(pairName, _pair);
			_pairs.push(_pair);
		}

		setPairs(pairsHm);

		for (const [k, v] of pairsHm) {
			console.log(k, v);
		}

		setTokens(_pairs);
		} catch(error) {
			console.log(error);
		}
	};

	const connectWallet = async () => {
		if (typeof window.ethereum === undefined) {
			setWalletError("MetaMask is not installed. Please install it to use this feature.");
			console.log(walletError);
			return;
		}

		try {
			const accounts = await window.ethereum.request({
				method: "eth_requestAccounts",
			});

			console.log(accounts);

			if (accounts.length === 0) {
				setConnected(false);
				localStorage.setItem('walletConnected', 'false');
				return;
			}

			setAccount(accounts[0]);
			setConnected(true);
			localStorage.setItem('walletConnected', 'true');

			const provider = new ethers.BrowserProvider(window.ethereum);

			const chainID =  (await provider.getNetwork()).chainId;
			if (convertSupportedChains().get(chainID) === undefined) {
				await window.ethereum.request({
					method: 'wallet_switchEthereumChain',
					params: [{ chainId: '0x' + supportedChains[0].chainId.toString(16) }], // chainId в формате hex (например, '0x1' для Ethereum Mainnet)
				});
			}

			setNetwork((await provider.getNetwork()));
		} catch (err) {
			console.log(err);
			setWalletError("Failed to connect wallet. Please try again.");
			console.log(walletError);
		}
	}

	const disconnectWallet = () => {
        setConnected(false);
        localStorage.setItem('walletConnected', 'false');
        setAccount(null);
    }

	const switchCard = () => {
		setIsSwapTokensCard(!isSwapTokensCard);
	}

  	useEffect(() => {
		const walletConnected = localStorageWalletConnectHandler();
		
		setConnected(walletConnected);
		if (walletConnected) {
			connectWallet();
		}
		if (account)
			getPools();
	}, [account]);

  return (
    <div className="min-h-screen bg-background flex flex-col">
      <Header
        isConnected={connected}
        walletAddress={account}
        networkName={networkName}
        onConnect={connectWallet}
        onDisconnect={disconnectWallet}
        onNetworkChange={onNetworkChange}
      />

      <main className="flex-1 container mx-auto flex flex-col items-center py-16 px-4 relative overflow-hidden">
        <Features />

        <div className="mt-16">
			<div style={{cursor:"pointer"}} onClick={switchCard}>
				<h1 className="text-4xl font-bold mb-8 text-center text-foreground bg-clip-text text-transparent bg-gradient-to-r from-purple-600 to-blue-600 animate-slide-in">
					{isSwapTokensCard ? "Swap Tokens" : "Add Liquidity"}
				</h1>
		  	</div>
          <div className="w-full max-w-[460px]">
			{isSwapTokensCard 
			? 
			<SwapCard 
				input={token0} 
				output={token1}
				onDirectionSwap={onDirectionSwap}
				onSwap={onSwap}
				onInputAmountChange={onInputAmountChange}
				outputAmount={outputAmount}
			/>:			
			<LiquidityCard 
			input={token0} 
			output={token1}
			onDirectionSwap={onDirectionSwap}
			onSwap={onSwap}
			onInputAmountChange={onInputAmountChange}
			outputAmount={outputAmount}
			/>
			}
            


          </div>

          <div className="mt-8 text-center text-sm text-muted-foreground">
            <p>Trade tokens in an instant</p>
            <p className="mt-2">Powered by automated liquidity pools</p>
          </div>
        </div>

        {/* Background decorative elements */}
        <div className="absolute top-0 left-0 w-full h-full -z-10 overflow-hidden">
          <div className="absolute top-20 left-20 w-64 h-64 bg-purple-500/10 dark:bg-purple-500/5 rounded-full filter blur-3xl animate-float"></div>
          <div
            className="absolute bottom-20 right-20 w-64 h-64 bg-blue-500/10 dark:bg-blue-500/5 rounded-full filter blur-3xl animate-float"
            style={{ animationDelay: "1s" }}
          ></div>
          <div
            className="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-[800px] h-[800px] rounded-full"
            style={{
              background:
                "radial-gradient(circle, rgba(147,51,234,0.07) 0%, rgba(79,70,229,0.03) 50%, transparent 70%)",
            }}
          ></div>
        </div>
      </main>

      <Footer />
    </div>
  );
};

export default Home;
