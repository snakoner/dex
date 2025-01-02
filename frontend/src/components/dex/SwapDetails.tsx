import React from "react";
import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Slider } from "@/components/ui/slider";
import { Info, Settings } from "lucide-react";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";

interface SwapDetailsProps {
  priceImpact?: string;
  minimumReceived?: string;
  slippageTolerance?: number;
  onSlippageChange?: (value: number) => void;
  liquidityProviderFee?: string;
}

const SwapDetails = ({
  priceImpact = "0.03%",
  minimumReceived = "0.0976 ETH",
  slippageTolerance = 0.5,
  onSlippageChange = () => {},
  liquidityProviderFee = "0.001 ETH",
}: SwapDetailsProps) => {
  return (
    <Card className="p-4 bg-background border-border hover:border-primary/30 transition-all duration-300">
      <div className="space-y-4">
        {/* Price Impact and Settings Row */}
        <div className="flex justify-between items-center">
          <div className="flex items-center gap-2">
            <span className="text-sm text-muted-foreground">Price Impact</span>
            <TooltipProvider>
              <Tooltip>
                <TooltipTrigger>
                  <Info className="h-4 w-4 text-muted-foreground" />
                </TooltipTrigger>
                <TooltipContent>
                  <p>
                    The difference between market price and estimated price due
                    to trade size
                  </p>
                </TooltipContent>
              </Tooltip>
            </TooltipProvider>
          </div>
          <span className="text-sm font-medium">{priceImpact}</span>
        </div>

        {/* Minimum Received Row */}
        <div className="flex justify-between items-center">
          <div className="flex items-center gap-2">
            <span className="text-sm text-muted-foreground">
              Minimum Received
            </span>
            <TooltipProvider>
              <Tooltip>
                <TooltipTrigger>
                  <Info className="h-4 w-4 text-muted-foreground" />
                </TooltipTrigger>
                <TooltipContent>
                  <p>The minimum amount you are guaranteed to receive</p>
                </TooltipContent>
              </Tooltip>
            </TooltipProvider>
          </div>
          <span className="text-sm font-medium">{minimumReceived}</span>
        </div>

        {/* Slippage Tolerance Section */}
        <div className="space-y-2">
          <div className="flex justify-between items-center">
            <div className="flex items-center gap-2">
              <span className="text-sm text-muted-foreground">
                Slippage Tolerance
              </span>
              <TooltipProvider>
                <Tooltip>
                  <TooltipTrigger>
                    <Settings className="h-4 w-4 text-muted-foreground" />
                  </TooltipTrigger>
                  <TooltipContent>
                    <p>
                      Your transaction will revert if the price changes
                      unfavorably by more than this percentage
                    </p>
                  </TooltipContent>
                </Tooltip>
              </TooltipProvider>
            </div>
            <span className="text-sm font-medium">{slippageTolerance}%</span>
          </div>
          <Slider
            defaultValue={[slippageTolerance]}
            max={5}
            step={0.1}
            onValueChange={(values) => onSlippageChange(values[0])}
          />
          <div className="flex gap-2">
            <Button
              variant="outline"
              size="sm"
              className="flex-1 hover:bg-primary/10 transition-colors duration-300"
              onClick={() => onSlippageChange(0.1)}
            >
              0.1%
            </Button>
            <Button
              variant="outline"
              size="sm"
              className="flex-1 hover:bg-primary/10 transition-colors duration-300"
              onClick={() => onSlippageChange(0.5)}
            >
              0.5%
            </Button>
            <Button
              variant="outline"
              size="sm"
              className="flex-1 hover:bg-primary/10 transition-colors duration-300"
              onClick={() => onSlippageChange(1.0)}
            >
              1.0%
            </Button>
          </div>
        </div>

        {/* Liquidity Provider Fee Row */}
        <div className="flex justify-between items-center">
          <div className="flex items-center gap-2">
            <span className="text-sm text-muted-foreground">
              Liquidity Provider Fee
            </span>
            <TooltipProvider>
              <Tooltip>
                <TooltipTrigger>
                  <Info className="h-4 w-4 text-muted-foreground" />
                </TooltipTrigger>
                <TooltipContent>
                  <p>
                    A portion of each trade goes to liquidity providers as
                    incentive
                  </p>
                </TooltipContent>
              </Tooltip>
            </TooltipProvider>
          </div>
          <span className="text-sm font-medium">{liquidityProviderFee}</span>
        </div>
      </div>
    </Card>
  );
};

export default SwapDetails;
