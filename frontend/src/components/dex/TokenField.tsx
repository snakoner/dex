import React from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { ChevronDown } from "lucide-react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";

interface TokenFieldProps {
  label?: string;
  tokenSymbol?: string;
  tokenIcon?: string;
  balance?: string;
  amount?: number;
  useInputField?: boolean;
  onAmountChange?: (name: string, amount: number) => void;
  onTokenSelect?: (token: string) => void;
}

const defaultTokens = [
  {
    symbol: "ETH",
    icon: "https://api.dicebear.com/7.x/avataaars/svg?seed=ETH",
    balance: "1.5",
  },
  {
    symbol: "USDC",
    icon: "https://api.dicebear.com/7.x/avataaars/svg?seed=USDC",
    balance: "1000.0",
  },
  {
    symbol: "DAI",
    icon: "https://api.dicebear.com/7.x/avataaars/svg?seed=DAI",
    balance: "1000.0",
  },
];

const TokenField = ({
  label = "From",
  tokenSymbol = "ETH",
  tokenIcon = "https://api.dicebear.com/7.x/avataaars/svg?seed=ETH",
  balance = "1.5",
  amount = 0,
  useInputField = false,
  onAmountChange = (num: string, amount: number) => {},
  onTokenSelect = () => {},
}: TokenFieldProps) => {
  return (
    <div className="w-full p-4 rounded-lg bg-card border border-border hover:border-primary/50 transition-colors duration-300">
      <div className="flex justify-between items-center mb-2">
        <span className="text-sm text-muted-foreground">{label}</span>
        <span className="text-sm text-muted-foreground">
          Balance: {balance}
        </span>
      </div>

      <div className="flex gap-2">
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button
              variant="outline"
              className="flex items-center gap-2 hover:border-primary/50 transition-all duration-300"
            >
              <img
                src={tokenIcon}
                alt={tokenSymbol}
                className="w-5 h-5 rounded-full"
              />
              <span>{tokenSymbol}</span>
              <ChevronDown className="h-4 w-4" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent>
            {defaultTokens.map((token) => (
              <DropdownMenuItem
                key={token.symbol}
                onClick={() => onTokenSelect(token.symbol)}
                className="flex items-center gap-2"
              >
                <img
                  src={token.icon}
                  alt={token.symbol}
                  className="w-5 h-5 rounded-full"
                />
                <span>{token.symbol}</span>
                <span className="text-muted-foreground ml-2">
                  {token.balance}
                </span>
              </DropdownMenuItem>
            ))}
          </DropdownMenuContent>
        </DropdownMenu>
        {
          useInputField ?
          <Input
          type="number"
          placeholder="0"
          id="_this"
          onChange={(e) => 
            {
              document.getElementById('_this').innerHTML = e.target.value;
              onAmountChange(tokenSymbol, Number(e.target.value));
            }}
          className="flex-1 focus:ring-2 focus:ring-primary/30 transition-all duration-300"
        /> : 
          <Input
          type="number"
          placeholder="0"
          id="_this"
          value={amount}
          onChange={(e) => 
            {
              document.getElementById('_this').innerHTML = e.target.value;
              onAmountChange(tokenSymbol, Number(e.target.value));
            }}
          className="flex-1 focus:ring-2 focus:ring-primary/30 transition-all duration-300"
          />
        }
        
      </div>
    </div>
  );
};

export default TokenField;
