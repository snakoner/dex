import React from "react";
import { Button } from "@/components/ui/button";
import { WalletIcon, NetworkIcon } from "lucide-react";
import { ThemeToggle } from "@/components/ui/theme-toggle";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";

interface HeaderProps {
  isConnected?: boolean;
  walletAddress?: string;
  networkName?: string;
  onConnect?: () => void;
  onDisconnect?: () => void;
  onNetworkChange?: (network: string) => void;
}

const networks = [
  { id: "1", name: "Ethereum Mainnet" },
  { id: "5", name: "Goerli Testnet" },
  { id: "11155111", name: "Sepolia Testnet" },
];

const Header = ({
  isConnected = false,
  walletAddress = "0x1234...5678",
  networkName = "Ethereum Mainnet",
  onConnect = () => {},
  onDisconnect = () => {},
  onNetworkChange = () => {},
}: HeaderProps) => {
  return (
    <header className="w-full h-[72px] px-6 bg-background/80 backdrop-blur-sm border-b border-border flex items-center justify-between sticky top-0 z-50">
      <div className="flex items-center gap-4">
        <h1 className="text-xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-purple-600 to-blue-600">
          DEX
        </h1>

        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="outline" className="flex items-center gap-2">
              <NetworkIcon className="h-4 w-4" />
              <span>{networkName}</span>
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent>
            {networks.map((network) => (
              <DropdownMenuItem
                key={network.id}
                onClick={() => onNetworkChange(network.name)}
              >
                {network.name}
              </DropdownMenuItem>
            ))}
          </DropdownMenuContent>
        </DropdownMenu>
      </div>

      <div className="flex items-center gap-4">
        <ThemeToggle />

        {isConnected ? (
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button variant="outline" className="flex items-center gap-2">
                <WalletIcon className="h-4 w-4" />
                <span>{walletAddress}</span>
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent>
              <DropdownMenuItem onClick={onDisconnect}>
                Disconnect
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        ) : (
          <Button onClick={onConnect} className="flex items-center gap-2">
            <WalletIcon className="h-4 w-4" />
            Connect Wallet
          </Button>
        )}
      </div>
    </header>
  );
};

export default Header;
