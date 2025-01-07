import React, { useState } from "react";
import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { ArrowDownUp } from "lucide-react";
import TokenField from "./TokenField";
import SwapDetails from "./SwapDetails";
import { ethers } from "ethers";
import { ALCHEMY_RPC_URL, ERC20_ABI } from "../constants";

interface InputToken {
  symbol: string;
  icon: string;
  balance: string;
};

interface OutputToken {
  symbol: string;
  icon: string;
  balance: string;
};


interface SwapCardProps {
  input?: InputToken;
  output?: OutputToken;
  inputAmount?: string;
  outputAmount?: string;
  onSwap?: () => void;
  onInputAmountChange?: (value: string) => void;
  onInputTokenSelect?: (token: string) => void;
  onOutputTokenSelect?: (token: string) => void;
}

const defaultInputToken0: InputToken = {
  symbol: "PEPE",
  icon: "https://cryptologos.cc/logos/pepe-pepe-logo.png?v=040",
  balance: "1.5",
};

const defaultInputToken1 : InputToken = {
  symbol: "GRT",
  icon: "https://cryptologos.cc/logos/the-graph-grt-logo.png?v=040",
  balance: "1000.0",
};

const SwapCard = ({
  inputAmount = "",
  outputAmount = "",
  onSwap = () => {},
  onInputAmountChange = () => {},
  onInputTokenSelect = () => {},
  onOutputTokenSelect = () => {},
}: SwapCardProps) => {
  const [token0, setToken0] = useState<InputToken>(defaultInputToken0);
  const [token1, setToken1] = useState<InputToken>(defaultInputToken1);

  const providerRpc = new ethers.JsonRpcProvider(ALCHEMY_RPC_URL);
  // const erc20InteractProvider = new ethers.Contract();

  // const getToken0Balance = async () => {  
  //   try {
  //     // await providerRpc.
  //   }
  // };

  return (
    <Card className="w-[460px] p-6 space-y-4 bg-background border-border backdrop-blur-sm bg-opacity-95 shadow-lg hover:shadow-xl transition-all duration-300 animate-slide-in">
      <div className="space-y-2">
        <TokenField
          label="From"
          tokenSymbol={token0.symbol}
          tokenIcon={token0.icon}
          balance={token0.balance}
          amount={inputAmount}
          onAmountChange={onInputAmountChange}
          onTokenSelect={onInputTokenSelect}
        />

        <div className="flex justify-center">
          <Button
            variant="ghost"
            size="icon"
            className="rounded-full hover:bg-primary/20 transition-colors duration-300 hover:scale-110 transform"
            onClick={() => {
              setToken0(token1);
              setToken1(token0);
              onInputTokenSelect(token1.symbol);
              onOutputTokenSelect(token0.symbol);
            }}
          >
            <ArrowDownUp className="h-4 w-4" />
          </Button>
        </div>

        <TokenField
          label="To"
          tokenSymbol={token1.symbol}
          tokenIcon={token1.icon}
          balance={token1.balance}
          amount={outputAmount}
          onTokenSelect={onOutputTokenSelect}
        />
      </div>

      <SwapDetails />

      <Button
        className="w-full hover:scale-[1.02] transition-transform duration-300 bg-gradient-to-r from-purple-600 to-blue-600 hover:from-purple-700 hover:to-blue-700"
        size="lg"
        onClick={onSwap}
      >
        Swap
      </Button>
    </Card>
  );
};

export default SwapCard;
