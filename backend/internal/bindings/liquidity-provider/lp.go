// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package liquidityprovider

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// LiquidityproviderMetaData contains all meta data concerning the Liquidityprovider contract.
var LiquidityproviderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b506040516118cb3803806118cb83398181016040528101906100319190610348565b808383816003908161004391906105e0565b50806004908161005391906105e0565b5050505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036100c6575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016100bd91906106be565b60405180910390fd5b6100d5816100de60201b60201c565b505050506106d7565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610200826101ba565b810181811067ffffffffffffffff8211171561021f5761021e6101ca565b5b80604052505050565b5f6102316101a1565b905061023d82826101f7565b919050565b5f67ffffffffffffffff82111561025c5761025b6101ca565b5b610265826101ba565b9050602081019050919050565b8281835e5f83830152505050565b5f61029261028d84610242565b610228565b9050828152602081018484840111156102ae576102ad6101b6565b5b6102b9848285610272565b509392505050565b5f82601f8301126102d5576102d46101b2565b5b81516102e5848260208601610280565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610317826102ee565b9050919050565b6103278161030d565b8114610331575f5ffd5b50565b5f815190506103428161031e565b92915050565b5f5f5f6060848603121561035f5761035e6101aa565b5b5f84015167ffffffffffffffff81111561037c5761037b6101ae565b5b610388868287016102c1565b935050602084015167ffffffffffffffff8111156103a9576103a86101ae565b5b6103b5868287016102c1565b92505060406103c686828701610334565b9150509250925092565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061041e57607f821691505b602082108103610431576104306103da565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026104937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610458565b61049d8683610458565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6104e16104dc6104d7846104b5565b6104be565b6104b5565b9050919050565b5f819050919050565b6104fa836104c7565b61050e610506826104e8565b848454610464565b825550505050565b5f5f905090565b610525610516565b6105308184846104f1565b505050565b5b81811015610553576105485f8261051d565b600181019050610536565b5050565b601f8211156105985761056981610437565b61057284610449565b81016020851015610581578190505b61059561058d85610449565b830182610535565b50505b505050565b5f82821c905092915050565b5f6105b85f198460080261059d565b1980831691505092915050565b5f6105d083836105a9565b9150826002028217905092915050565b6105e9826103d0565b67ffffffffffffffff811115610602576106016101ca565b5b61060c8254610407565b610617828285610557565b5f60209050601f831160018114610648575f8415610636578287015190505b61064085826105c5565b8655506106a7565b601f19841661065686610437565b5f5b8281101561067d57848901518255600182019150602085019450602081019050610658565b8683101561069a5784890151610696601f8916826105a9565b8355505b6001600288020188555050505b505050505050565b6106b88161030d565b82525050565b5f6020820190506106d15f8301846106af565b92915050565b6111e7806106e45f395ff3fe608060405234801561000f575f5ffd5b50600436106100e8575f3560e01c8063715018a61161008a5780639dc29fac116100645780639dc29fac14610238578063a9059cbb14610254578063dd62ed3e14610284578063f2fde38b146102b4576100e8565b8063715018a6146101f25780638da5cb5b146101fc57806395d89b411461021a576100e8565b806323b872dd116100c657806323b872dd14610158578063313ce5671461018857806340c10f19146101a657806370a08231146101c2576100e8565b806306fdde03146100ec578063095ea7b31461010a57806318160ddd1461013a575b5f5ffd5b6100f46102d0565b6040516101019190610e60565b60405180910390f35b610124600480360381019061011f9190610f11565b610360565b6040516101319190610f69565b60405180910390f35b610142610382565b60405161014f9190610f91565b60405180910390f35b610172600480360381019061016d9190610faa565b61038b565b60405161017f9190610f69565b60405180910390f35b6101906103b9565b60405161019d9190611015565b60405180910390f35b6101c060048036038101906101bb9190610f11565b6103c1565b005b6101dc60048036038101906101d7919061102e565b6103d7565b6040516101e99190610f91565b60405180910390f35b6101fa61041c565b005b61020461042f565b6040516102119190611068565b60405180910390f35b610222610457565b60405161022f9190610e60565b60405180910390f35b610252600480360381019061024d9190610f11565b6104e7565b005b61026e60048036038101906102699190610f11565b6104fd565b60405161027b9190610f69565b60405180910390f35b61029e60048036038101906102999190611081565b61051f565b6040516102ab9190610f91565b60405180910390f35b6102ce60048036038101906102c9919061102e565b6105a1565b005b6060600380546102df906110ec565b80601f016020809104026020016040519081016040528092919081815260200182805461030b906110ec565b80156103565780601f1061032d57610100808354040283529160200191610356565b820191905f5260205f20905b81548152906001019060200180831161033957829003601f168201915b5050505050905090565b5f5f61036a610625565b905061037781858561062c565b600191505092915050565b5f600254905090565b5f5f610395610625565b90506103a285828561063e565b6103ad8585856106d0565b60019150509392505050565b5f6012905090565b6103c96107c0565b6103d38282610847565b5050565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6104246107c0565b61042d5f6108c6565b565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b606060048054610466906110ec565b80601f0160208091040260200160405190810160405280929190818152602001828054610492906110ec565b80156104dd5780601f106104b4576101008083540402835291602001916104dd565b820191905f5260205f20905b8154815290600101906020018083116104c057829003601f168201915b5050505050905090565b6104ef6107c0565b6104f98282610989565b5050565b5f5f610507610625565b90506105148185856106d0565b600191505092915050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b6105a96107c0565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610619575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016106109190611068565b60405180910390fd5b610622816108c6565b50565b5f33905090565b6106398383836001610a08565b505050565b5f610649848461051f565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146106ca57818110156106bb578281836040517ffb8f41b20000000000000000000000000000000000000000000000000000000081526004016106b29392919061111c565b60405180910390fd5b6106c984848484035f610a08565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610740575f6040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526004016107379190611068565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036107b0575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016107a79190611068565b60405180910390fd5b6107bb838383610bd7565b505050565b6107c8610625565b73ffffffffffffffffffffffffffffffffffffffff166107e661042f565b73ffffffffffffffffffffffffffffffffffffffff161461084557610809610625565b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815260040161083c9190611068565b60405180910390fd5b565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036108b7575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016108ae9190611068565b60405180910390fd5b6108c25f8383610bd7565b5050565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036109f9575f6040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526004016109f09190611068565b60405180910390fd5b610a04825f83610bd7565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610a78575f6040517fe602df05000000000000000000000000000000000000000000000000000000008152600401610a6f9190611068565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610ae8575f6040517f94280d62000000000000000000000000000000000000000000000000000000008152600401610adf9190611068565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508015610bd1578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610bc89190610f91565b60405180910390a35b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610c27578060025f828254610c1b919061117e565b92505081905550610cf5565b5f5f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015610cb0578381836040517fe450d38c000000000000000000000000000000000000000000000000000000008152600401610ca79392919061111c565b60405180910390fd5b8181035f5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610d3c578060025f8282540392505081905550610d86565b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610de39190610f91565b60405180910390a3505050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610e3282610df0565b610e3c8185610dfa565b9350610e4c818560208601610e0a565b610e5581610e18565b840191505092915050565b5f6020820190508181035f830152610e788184610e28565b905092915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610ead82610e84565b9050919050565b610ebd81610ea3565b8114610ec7575f5ffd5b50565b5f81359050610ed881610eb4565b92915050565b5f819050919050565b610ef081610ede565b8114610efa575f5ffd5b50565b5f81359050610f0b81610ee7565b92915050565b5f5f60408385031215610f2757610f26610e80565b5b5f610f3485828601610eca565b9250506020610f4585828601610efd565b9150509250929050565b5f8115159050919050565b610f6381610f4f565b82525050565b5f602082019050610f7c5f830184610f5a565b92915050565b610f8b81610ede565b82525050565b5f602082019050610fa45f830184610f82565b92915050565b5f5f5f60608486031215610fc157610fc0610e80565b5b5f610fce86828701610eca565b9350506020610fdf86828701610eca565b9250506040610ff086828701610efd565b9150509250925092565b5f60ff82169050919050565b61100f81610ffa565b82525050565b5f6020820190506110285f830184611006565b92915050565b5f6020828403121561104357611042610e80565b5b5f61105084828501610eca565b91505092915050565b61106281610ea3565b82525050565b5f60208201905061107b5f830184611059565b92915050565b5f5f6040838503121561109757611096610e80565b5b5f6110a485828601610eca565b92505060206110b585828601610eca565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061110357607f821691505b602082108103611116576111156110bf565b5b50919050565b5f60608201905061112f5f830186611059565b61113c6020830185610f82565b6111496040830184610f82565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61118882610ede565b915061119383610ede565b92508282019050808211156111ab576111aa611151565b5b9291505056fea26469706673582212200ff52428adf4fb258d2613ecb9f1a797504fedc523f527f26f7242fe1b6bc68d64736f6c634300081c0033",
}

// LiquidityproviderABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquidityproviderMetaData.ABI instead.
var LiquidityproviderABI = LiquidityproviderMetaData.ABI

// LiquidityproviderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LiquidityproviderMetaData.Bin instead.
var LiquidityproviderBin = LiquidityproviderMetaData.Bin

// DeployLiquidityprovider deploys a new Ethereum contract, binding an instance of Liquidityprovider to it.
func DeployLiquidityprovider(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, owner common.Address) (common.Address, *types.Transaction, *Liquidityprovider, error) {
	parsed, err := LiquidityproviderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LiquidityproviderBin), backend, name, symbol, owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Liquidityprovider{LiquidityproviderCaller: LiquidityproviderCaller{contract: contract}, LiquidityproviderTransactor: LiquidityproviderTransactor{contract: contract}, LiquidityproviderFilterer: LiquidityproviderFilterer{contract: contract}}, nil
}

// Liquidityprovider is an auto generated Go binding around an Ethereum contract.
type Liquidityprovider struct {
	LiquidityproviderCaller     // Read-only binding to the contract
	LiquidityproviderTransactor // Write-only binding to the contract
	LiquidityproviderFilterer   // Log filterer for contract events
}

// LiquidityproviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquidityproviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityproviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquidityproviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityproviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquidityproviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityproviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquidityproviderSession struct {
	Contract     *Liquidityprovider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LiquidityproviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquidityproviderCallerSession struct {
	Contract *LiquidityproviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// LiquidityproviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquidityproviderTransactorSession struct {
	Contract     *LiquidityproviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// LiquidityproviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquidityproviderRaw struct {
	Contract *Liquidityprovider // Generic contract binding to access the raw methods on
}

// LiquidityproviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquidityproviderCallerRaw struct {
	Contract *LiquidityproviderCaller // Generic read-only contract binding to access the raw methods on
}

// LiquidityproviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquidityproviderTransactorRaw struct {
	Contract *LiquidityproviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquidityprovider creates a new instance of Liquidityprovider, bound to a specific deployed contract.
func NewLiquidityprovider(address common.Address, backend bind.ContractBackend) (*Liquidityprovider, error) {
	contract, err := bindLiquidityprovider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Liquidityprovider{LiquidityproviderCaller: LiquidityproviderCaller{contract: contract}, LiquidityproviderTransactor: LiquidityproviderTransactor{contract: contract}, LiquidityproviderFilterer: LiquidityproviderFilterer{contract: contract}}, nil
}

// NewLiquidityproviderCaller creates a new read-only instance of Liquidityprovider, bound to a specific deployed contract.
func NewLiquidityproviderCaller(address common.Address, caller bind.ContractCaller) (*LiquidityproviderCaller, error) {
	contract, err := bindLiquidityprovider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidityproviderCaller{contract: contract}, nil
}

// NewLiquidityproviderTransactor creates a new write-only instance of Liquidityprovider, bound to a specific deployed contract.
func NewLiquidityproviderTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquidityproviderTransactor, error) {
	contract, err := bindLiquidityprovider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidityproviderTransactor{contract: contract}, nil
}

// NewLiquidityproviderFilterer creates a new log filterer instance of Liquidityprovider, bound to a specific deployed contract.
func NewLiquidityproviderFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquidityproviderFilterer, error) {
	contract, err := bindLiquidityprovider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquidityproviderFilterer{contract: contract}, nil
}

// bindLiquidityprovider binds a generic wrapper to an already deployed contract.
func bindLiquidityprovider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LiquidityproviderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Liquidityprovider *LiquidityproviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Liquidityprovider.Contract.LiquidityproviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Liquidityprovider *LiquidityproviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquidityprovider.Contract.LiquidityproviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Liquidityprovider *LiquidityproviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Liquidityprovider.Contract.LiquidityproviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Liquidityprovider *LiquidityproviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Liquidityprovider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Liquidityprovider *LiquidityproviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquidityprovider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Liquidityprovider *LiquidityproviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Liquidityprovider.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Liquidityprovider *LiquidityproviderCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Liquidityprovider.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Liquidityprovider *LiquidityproviderSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Liquidityprovider.Contract.Allowance(&_Liquidityprovider.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Liquidityprovider *LiquidityproviderCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Liquidityprovider.Contract.Allowance(&_Liquidityprovider.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Liquidityprovider *LiquidityproviderCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Liquidityprovider.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Liquidityprovider *LiquidityproviderSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Liquidityprovider.Contract.BalanceOf(&_Liquidityprovider.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Liquidityprovider *LiquidityproviderCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Liquidityprovider.Contract.BalanceOf(&_Liquidityprovider.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Liquidityprovider *LiquidityproviderCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Liquidityprovider.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Liquidityprovider *LiquidityproviderSession) Decimals() (uint8, error) {
	return _Liquidityprovider.Contract.Decimals(&_Liquidityprovider.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Liquidityprovider *LiquidityproviderCallerSession) Decimals() (uint8, error) {
	return _Liquidityprovider.Contract.Decimals(&_Liquidityprovider.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Liquidityprovider *LiquidityproviderCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Liquidityprovider.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Liquidityprovider *LiquidityproviderSession) Name() (string, error) {
	return _Liquidityprovider.Contract.Name(&_Liquidityprovider.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Liquidityprovider *LiquidityproviderCallerSession) Name() (string, error) {
	return _Liquidityprovider.Contract.Name(&_Liquidityprovider.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Liquidityprovider *LiquidityproviderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liquidityprovider.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Liquidityprovider *LiquidityproviderSession) Owner() (common.Address, error) {
	return _Liquidityprovider.Contract.Owner(&_Liquidityprovider.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Liquidityprovider *LiquidityproviderCallerSession) Owner() (common.Address, error) {
	return _Liquidityprovider.Contract.Owner(&_Liquidityprovider.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Liquidityprovider *LiquidityproviderCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Liquidityprovider.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Liquidityprovider *LiquidityproviderSession) Symbol() (string, error) {
	return _Liquidityprovider.Contract.Symbol(&_Liquidityprovider.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Liquidityprovider *LiquidityproviderCallerSession) Symbol() (string, error) {
	return _Liquidityprovider.Contract.Symbol(&_Liquidityprovider.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Liquidityprovider *LiquidityproviderCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Liquidityprovider.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Liquidityprovider *LiquidityproviderSession) TotalSupply() (*big.Int, error) {
	return _Liquidityprovider.Contract.TotalSupply(&_Liquidityprovider.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Liquidityprovider *LiquidityproviderCallerSession) TotalSupply() (*big.Int, error) {
	return _Liquidityprovider.Contract.TotalSupply(&_Liquidityprovider.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Liquidityprovider *LiquidityproviderTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Liquidityprovider.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Liquidityprovider *LiquidityproviderSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Liquidityprovider.Contract.Approve(&_Liquidityprovider.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Liquidityprovider *LiquidityproviderTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Liquidityprovider.Contract.Approve(&_Liquidityprovider.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_Liquidityprovider *LiquidityproviderTransactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Liquidityprovider.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_Liquidityprovider *LiquidityproviderSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Liquidityprovider.Contract.Burn(&_Liquidityprovider.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_Liquidityprovider *LiquidityproviderTransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Liquidityprovider.Contract.Burn(&_Liquidityprovider.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Liquidityprovider *LiquidityproviderTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Liquidityprovider.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Liquidityprovider *LiquidityproviderSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Liquidityprovider.Contract.Mint(&_Liquidityprovider.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Liquidityprovider *LiquidityproviderTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Liquidityprovider.Contract.Mint(&_Liquidityprovider.TransactOpts, account, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Liquidityprovider *LiquidityproviderTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquidityprovider.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Liquidityprovider *LiquidityproviderSession) RenounceOwnership() (*types.Transaction, error) {
	return _Liquidityprovider.Contract.RenounceOwnership(&_Liquidityprovider.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Liquidityprovider *LiquidityproviderTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Liquidityprovider.Contract.RenounceOwnership(&_Liquidityprovider.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Liquidityprovider *LiquidityproviderTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Liquidityprovider.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Liquidityprovider *LiquidityproviderSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Liquidityprovider.Contract.Transfer(&_Liquidityprovider.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Liquidityprovider *LiquidityproviderTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Liquidityprovider.Contract.Transfer(&_Liquidityprovider.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Liquidityprovider *LiquidityproviderTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Liquidityprovider.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Liquidityprovider *LiquidityproviderSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Liquidityprovider.Contract.TransferFrom(&_Liquidityprovider.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Liquidityprovider *LiquidityproviderTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Liquidityprovider.Contract.TransferFrom(&_Liquidityprovider.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Liquidityprovider *LiquidityproviderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Liquidityprovider.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Liquidityprovider *LiquidityproviderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Liquidityprovider.Contract.TransferOwnership(&_Liquidityprovider.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Liquidityprovider *LiquidityproviderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Liquidityprovider.Contract.TransferOwnership(&_Liquidityprovider.TransactOpts, newOwner)
}

// LiquidityproviderApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Liquidityprovider contract.
type LiquidityproviderApprovalIterator struct {
	Event *LiquidityproviderApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LiquidityproviderApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityproviderApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LiquidityproviderApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LiquidityproviderApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityproviderApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityproviderApproval represents a Approval event raised by the Liquidityprovider contract.
type LiquidityproviderApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Liquidityprovider *LiquidityproviderFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LiquidityproviderApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Liquidityprovider.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityproviderApprovalIterator{contract: _Liquidityprovider.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Liquidityprovider *LiquidityproviderFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LiquidityproviderApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Liquidityprovider.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityproviderApproval)
				if err := _Liquidityprovider.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Liquidityprovider *LiquidityproviderFilterer) ParseApproval(log types.Log) (*LiquidityproviderApproval, error) {
	event := new(LiquidityproviderApproval)
	if err := _Liquidityprovider.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityproviderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Liquidityprovider contract.
type LiquidityproviderOwnershipTransferredIterator struct {
	Event *LiquidityproviderOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LiquidityproviderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityproviderOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LiquidityproviderOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LiquidityproviderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityproviderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityproviderOwnershipTransferred represents a OwnershipTransferred event raised by the Liquidityprovider contract.
type LiquidityproviderOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Liquidityprovider *LiquidityproviderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LiquidityproviderOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Liquidityprovider.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityproviderOwnershipTransferredIterator{contract: _Liquidityprovider.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Liquidityprovider *LiquidityproviderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LiquidityproviderOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Liquidityprovider.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityproviderOwnershipTransferred)
				if err := _Liquidityprovider.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Liquidityprovider *LiquidityproviderFilterer) ParseOwnershipTransferred(log types.Log) (*LiquidityproviderOwnershipTransferred, error) {
	event := new(LiquidityproviderOwnershipTransferred)
	if err := _Liquidityprovider.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityproviderTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Liquidityprovider contract.
type LiquidityproviderTransferIterator struct {
	Event *LiquidityproviderTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LiquidityproviderTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityproviderTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LiquidityproviderTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LiquidityproviderTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityproviderTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityproviderTransfer represents a Transfer event raised by the Liquidityprovider contract.
type LiquidityproviderTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Liquidityprovider *LiquidityproviderFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LiquidityproviderTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Liquidityprovider.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityproviderTransferIterator{contract: _Liquidityprovider.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Liquidityprovider *LiquidityproviderFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LiquidityproviderTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Liquidityprovider.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityproviderTransfer)
				if err := _Liquidityprovider.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Liquidityprovider *LiquidityproviderFilterer) ParseTransfer(log types.Log) (*LiquidityproviderTransfer, error) {
	event := new(LiquidityproviderTransfer)
	if err := _Liquidityprovider.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
