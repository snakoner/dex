import React from "react";
import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { ArrowDownUp } from "lucide-react";
import TokenField from "./TokenField";
import SwapDetails from "./SwapDetails";

interface SwapCardProps {
  inputToken?: {
    symbol: string;
    icon: string;
    balance: string;
  };
  outputToken?: {
    symbol: string;
    icon: string;
    balance: string;
  };
  inputAmount?: string;
  outputAmount?: string;
  onSwap?: () => void;
  onInputAmountChange?: (value: string) => void;
  onInputTokenSelect?: (token: string) => void;
  onOutputTokenSelect?: (token: string) => void;
}

const SwapCard = ({
  inputToken = {
    symbol: "ETH",
    icon: "https://api.dicebear.com/7.x/avataaars/svg?seed=ETH",
    balance: "1.5",
  },
  outputToken = {
    symbol: "USDC",
    icon: "https://api.dicebear.com/7.x/avataaars/svg?seed=USDC",
    balance: "1000.0",
  },
  inputAmount = "",
  outputAmount = "",
  onSwap = () => {},
  onInputAmountChange = () => {},
  onInputTokenSelect = () => {},
  onOutputTokenSelect = () => {},
}: SwapCardProps) => {
  return (
    <Card className="w-[460px] p-6 space-y-4 bg-background border-border backdrop-blur-sm bg-opacity-95 shadow-lg hover:shadow-xl transition-all duration-300 animate-slide-in">
      <div className="space-y-2">
        <TokenField
          label="From"
          tokenSymbol={inputToken.symbol}
          tokenIcon={inputToken.icon}
          balance={inputToken.balance}
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
              onInputTokenSelect(outputToken.symbol);
              onOutputTokenSelect(inputToken.symbol);
            }}
          >
            <ArrowDownUp className="h-4 w-4" />
          </Button>
        </div>

        <TokenField
          label="To"
          tokenSymbol={outputToken.symbol}
          tokenIcon={outputToken.icon}
          balance={outputToken.balance}
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
