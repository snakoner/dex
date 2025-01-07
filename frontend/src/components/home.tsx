import React, { useEffect, useState } from "react";
import Header from "./dex/Header";
import SwapCard from "./dex/SwapCard";
import Features from "./dex/Features";
import Footer from "./footer";
import { readPools } from "@/lib/read-json";

interface HomeProps {
  isWalletConnected?: boolean;
  walletAddress?: string;
  networkName?: string;
  onWalletConnect?: () => void;
  onWalletDisconnect?: () => void;
  onNetworkChange?: (network: string) => void;
  onSwap?: () => void;
}

interface Pair {
  nameA: string;
  nameB: string;
  pool: string;
  tokenA: string;
  tokenB: string;
  tokenLP: string;
};

const Home = ({
  isWalletConnected = false,
  walletAddress = "0x1234...5678",
  networkName = "Ethereum Mainnet",
  onWalletConnect = () => {},
  onWalletDisconnect = () => {},
  onNetworkChange = () => {},
  onSwap = () => {},
}: HomeProps) => {
  const [pairs, setPairs] = useState<Pair[] | null>();
  const [factoryAddress, setFactoryAddress] = useState<string|null>();

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
      for (const p of data?.pairs) {
        const _pair: Pair = {
          nameA: p.nameA,
          nameB: p.nameB,
          tokenA: p.tokenA,
          tokenB: p.tokenB,
          tokenLP: p.tokenLP,
          pool: p.pool,
        };
        _pairs.push(_pair);
      }

      console.log(_pairs);

      setPairs(_pairs);
		} catch(error) {
			console.log(error);
		}
	};


  useEffect(() => {
    getPools();
  }, []);

  return (
    <div className="min-h-screen bg-background flex flex-col">
      <Header
        isConnected={isWalletConnected}
        walletAddress={walletAddress}
        networkName={networkName}
        onConnect={onWalletConnect}
        onDisconnect={onWalletDisconnect}
        onNetworkChange={onNetworkChange}
      />

      <main className="flex-1 container mx-auto flex flex-col items-center py-16 px-4 relative overflow-hidden">
        <Features />

        <div className="mt-16">
          <h1 className="text-4xl font-bold mb-8 text-center text-foreground bg-clip-text text-transparent bg-gradient-to-r from-purple-600 to-blue-600 animate-slide-in">
            Swap Tokens
          </h1>

          <div className="w-full max-w-[460px]">
            <SwapCard onSwap={onSwap} />
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
