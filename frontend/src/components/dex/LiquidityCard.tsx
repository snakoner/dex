import React, { useState } from "react";
import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { ArrowDownUp } from "lucide-react";
import TokenField from "./TokenField";
import SwapDetails from "./SwapDetails";
import { ethers } from "ethers";
import { ALCHEMY_RPC_URL, ERC20_ABI } from "../constants";
import { Token } from "../home";
import LiquidityDetails from "./LiquidityDetails";

interface InputToken {
  symbol: string;
  icon: string;
  balance: string;
};

interface InputToken {
  symbol: string;
  icon: string;
  balance: string;
};


interface SwapCardProps {
  input?: Token;
  output?: Token;
  inputAmount?: number;
  outputAmount?: number;
  onSwap?: () => void;
  onDirectionSwap?: () => void;
  onInputAmountChange?: (name: string, amount: number) => void;
  onInputTokenSelect?: (token: string) => void;
  onOutputTokenSelect?: (token: string) => void;
}

const toStringBalance = (balance: bigint, decimals: bigint) =>  {
  try {
    if (balance === ethers.toBigInt(0)) {
      return "0.0";
    } else {
      const num = balance;
      const divisor = Math.pow(10, Number(decimals));
      const b = num / BigInt(divisor);
      return b.toString();
    }  
  } catch (error) {
    return "0.0";
  }
}

const LiquidityCard = ({
  input, 
  output,
  inputAmount,
  outputAmount,
  onDirectionSwap = () => {},
  onSwap = () => {},
  onInputAmountChange = (name: string, amount: number) => {},
  onInputTokenSelect = () => {},
  onOutputTokenSelect = () => {},
}: SwapCardProps) => {

  return (
    <Card className="w-[460px] p-6 space-y-4 bg-background border-border backdrop-blur-sm bg-opacity-95 shadow-lg hover:shadow-xl transition-all duration-300 animate-slide-in">
      <div className="space-y-2">
        <TokenField
          label="Max"
          tokenSymbol={input?.name}
          tokenIcon={input?.icon}
          balance={toStringBalance(input?.balance, input?.decimals)}
          amount={inputAmount}
          useInputField={true}
          onAmountChange={onInputAmountChange}
          onTokenSelect={onInputTokenSelect}
        />


        <TokenField
          label="Max"
          tokenSymbol={output?.name}
          tokenIcon={output?.icon}
          balance={toStringBalance(output?.balance, output?.decimals)}
          amount={outputAmount}
          useInputField={false}
          onTokenSelect={onOutputTokenSelect}
        />
      </div>

        <LiquidityDetails 
            token0Name={input?.name}
            token1Name={output?.name}
            reserves0={0}
            reserves1={0}
        />

      <Button
        className="w-full hover:scale-[1.02] transition-transform duration-300 bg-gradient-to-r from-purple-600 to-blue-600 hover:from-purple-700 hover:to-blue-700"
        size="lg"
        onClick={onSwap}
      >
        Add liquidity
      </Button>
    </Card>
  );
};

export default LiquidityCard;
