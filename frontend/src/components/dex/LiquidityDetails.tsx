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
  token0Name?: string;
  token1Name?: string;
  reserves0?: bigint;
  reserves1?: bigint;
}

const LiquidityDetails = ({
  priceImpact = "0.03%",
  token0Name,
  token1Name,
  reserves0,
  reserves1,
}: SwapDetailsProps) => {
  return (
    <Card className="p-4 bg-background border-border hover:border-primary/30 transition-all duration-300">
      <div className="space-y-4">
        {/* Price Impact and Settings Row */}
        <div className="flex justify-between items-center">
          <div className="flex items-center gap-2">
            <span className="text-sm text-muted-foreground">{token0Name} in pool: </span>
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
          <span className="text-sm font-medium">{Number(reserves0)} {token0Name}</span>
        </div>


        <div className="flex justify-between items-center">
          <div className="flex items-center gap-2">
            <span className="text-sm text-muted-foreground">{token1Name} in pool: </span>
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
          <span className="text-sm font-medium">{Number(reserves1)} {token1Name}</span>
        </div>
      </div>
    </Card>
  );
};

export default LiquidityDetails;
