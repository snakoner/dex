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

// LpMetaData contains all meta data concerning the Lp contract.
var LpMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b506040516118cb3803806118cb83398181016040528101906100319190610348565b808383816003908161004391906105e0565b50806004908161005391906105e0565b5050505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036100c6575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016100bd91906106be565b60405180910390fd5b6100d5816100de60201b60201c565b505050506106d7565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610200826101ba565b810181811067ffffffffffffffff8211171561021f5761021e6101ca565b5b80604052505050565b5f6102316101a1565b905061023d82826101f7565b919050565b5f67ffffffffffffffff82111561025c5761025b6101ca565b5b610265826101ba565b9050602081019050919050565b8281835e5f83830152505050565b5f61029261028d84610242565b610228565b9050828152602081018484840111156102ae576102ad6101b6565b5b6102b9848285610272565b509392505050565b5f82601f8301126102d5576102d46101b2565b5b81516102e5848260208601610280565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610317826102ee565b9050919050565b6103278161030d565b8114610331575f5ffd5b50565b5f815190506103428161031e565b92915050565b5f5f5f6060848603121561035f5761035e6101aa565b5b5f84015167ffffffffffffffff81111561037c5761037b6101ae565b5b610388868287016102c1565b935050602084015167ffffffffffffffff8111156103a9576103a86101ae565b5b6103b5868287016102c1565b92505060406103c686828701610334565b9150509250925092565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061041e57607f821691505b602082108103610431576104306103da565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026104937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610458565b61049d8683610458565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6104e16104dc6104d7846104b5565b6104be565b6104b5565b9050919050565b5f819050919050565b6104fa836104c7565b61050e610506826104e8565b848454610464565b825550505050565b5f5f905090565b610525610516565b6105308184846104f1565b505050565b5b81811015610553576105485f8261051d565b600181019050610536565b5050565b601f8211156105985761056981610437565b61057284610449565b81016020851015610581578190505b61059561058d85610449565b830182610535565b50505b505050565b5f82821c905092915050565b5f6105b85f198460080261059d565b1980831691505092915050565b5f6105d083836105a9565b9150826002028217905092915050565b6105e9826103d0565b67ffffffffffffffff811115610602576106016101ca565b5b61060c8254610407565b610617828285610557565b5f60209050601f831160018114610648575f8415610636578287015190505b61064085826105c5565b8655506106a7565b601f19841661065686610437565b5f5b8281101561067d57848901518255600182019150602085019450602081019050610658565b8683101561069a5784890151610696601f8916826105a9565b8355505b6001600288020188555050505b505050505050565b6106b88161030d565b82525050565b5f6020820190506106d15f8301846106af565b92915050565b6111e7806106e45f395ff3fe608060405234801561000f575f5ffd5b50600436106100e8575f3560e01c8063715018a61161008a5780639dc29fac116100645780639dc29fac14610238578063a9059cbb14610254578063dd62ed3e14610284578063f2fde38b146102b4576100e8565b8063715018a6146101f25780638da5cb5b146101fc57806395d89b411461021a576100e8565b806323b872dd116100c657806323b872dd14610158578063313ce5671461018857806340c10f19146101a657806370a08231146101c2576100e8565b806306fdde03146100ec578063095ea7b31461010a57806318160ddd1461013a575b5f5ffd5b6100f46102d0565b6040516101019190610e60565b60405180910390f35b610124600480360381019061011f9190610f11565b610360565b6040516101319190610f69565b60405180910390f35b610142610382565b60405161014f9190610f91565b60405180910390f35b610172600480360381019061016d9190610faa565b61038b565b60405161017f9190610f69565b60405180910390f35b6101906103b9565b60405161019d9190611015565b60405180910390f35b6101c060048036038101906101bb9190610f11565b6103c1565b005b6101dc60048036038101906101d7919061102e565b6103d7565b6040516101e99190610f91565b60405180910390f35b6101fa61041c565b005b61020461042f565b6040516102119190611068565b60405180910390f35b610222610457565b60405161022f9190610e60565b60405180910390f35b610252600480360381019061024d9190610f11565b6104e7565b005b61026e60048036038101906102699190610f11565b6104fd565b60405161027b9190610f69565b60405180910390f35b61029e60048036038101906102999190611081565b61051f565b6040516102ab9190610f91565b60405180910390f35b6102ce60048036038101906102c9919061102e565b6105a1565b005b6060600380546102df906110ec565b80601f016020809104026020016040519081016040528092919081815260200182805461030b906110ec565b80156103565780601f1061032d57610100808354040283529160200191610356565b820191905f5260205f20905b81548152906001019060200180831161033957829003601f168201915b5050505050905090565b5f5f61036a610625565b905061037781858561062c565b600191505092915050565b5f600254905090565b5f5f610395610625565b90506103a285828561063e565b6103ad8585856106d0565b60019150509392505050565b5f6012905090565b6103c96107c0565b6103d38282610847565b5050565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6104246107c0565b61042d5f6108c6565b565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b606060048054610466906110ec565b80601f0160208091040260200160405190810160405280929190818152602001828054610492906110ec565b80156104dd5780601f106104b4576101008083540402835291602001916104dd565b820191905f5260205f20905b8154815290600101906020018083116104c057829003601f168201915b5050505050905090565b6104ef6107c0565b6104f98282610989565b5050565b5f5f610507610625565b90506105148185856106d0565b600191505092915050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b6105a96107c0565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610619575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016106109190611068565b60405180910390fd5b610622816108c6565b50565b5f33905090565b6106398383836001610a08565b505050565b5f610649848461051f565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146106ca57818110156106bb578281836040517ffb8f41b20000000000000000000000000000000000000000000000000000000081526004016106b29392919061111c565b60405180910390fd5b6106c984848484035f610a08565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610740575f6040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526004016107379190611068565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036107b0575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016107a79190611068565b60405180910390fd5b6107bb838383610bd7565b505050565b6107c8610625565b73ffffffffffffffffffffffffffffffffffffffff166107e661042f565b73ffffffffffffffffffffffffffffffffffffffff161461084557610809610625565b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815260040161083c9190611068565b60405180910390fd5b565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036108b7575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016108ae9190611068565b60405180910390fd5b6108c25f8383610bd7565b5050565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036109f9575f6040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526004016109f09190611068565b60405180910390fd5b610a04825f83610bd7565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610a78575f6040517fe602df05000000000000000000000000000000000000000000000000000000008152600401610a6f9190611068565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610ae8575f6040517f94280d62000000000000000000000000000000000000000000000000000000008152600401610adf9190611068565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508015610bd1578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610bc89190610f91565b60405180910390a35b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610c27578060025f828254610c1b919061117e565b92505081905550610cf5565b5f5f5f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015610cb0578381836040517fe450d38c000000000000000000000000000000000000000000000000000000008152600401610ca79392919061111c565b60405180910390fd5b8181035f5f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610d3c578060025f8282540392505081905550610d86565b805f5f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610de39190610f91565b60405180910390a3505050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610e3282610df0565b610e3c8185610dfa565b9350610e4c818560208601610e0a565b610e5581610e18565b840191505092915050565b5f6020820190508181035f830152610e788184610e28565b905092915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610ead82610e84565b9050919050565b610ebd81610ea3565b8114610ec7575f5ffd5b50565b5f81359050610ed881610eb4565b92915050565b5f819050919050565b610ef081610ede565b8114610efa575f5ffd5b50565b5f81359050610f0b81610ee7565b92915050565b5f5f60408385031215610f2757610f26610e80565b5b5f610f3485828601610eca565b9250506020610f4585828601610efd565b9150509250929050565b5f8115159050919050565b610f6381610f4f565b82525050565b5f602082019050610f7c5f830184610f5a565b92915050565b610f8b81610ede565b82525050565b5f602082019050610fa45f830184610f82565b92915050565b5f5f5f60608486031215610fc157610fc0610e80565b5b5f610fce86828701610eca565b9350506020610fdf86828701610eca565b9250506040610ff086828701610efd565b9150509250925092565b5f60ff82169050919050565b61100f81610ffa565b82525050565b5f6020820190506110285f830184611006565b92915050565b5f6020828403121561104357611042610e80565b5b5f61105084828501610eca565b91505092915050565b61106281610ea3565b82525050565b5f60208201905061107b5f830184611059565b92915050565b5f5f6040838503121561109757611096610e80565b5b5f6110a485828601610eca565b92505060206110b585828601610eca565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061110357607f821691505b602082108103611116576111156110bf565b5b50919050565b5f60608201905061112f5f830186611059565b61113c6020830185610f82565b6111496040830184610f82565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61118882610ede565b915061119383610ede565b92508282019050808211156111ab576111aa611151565b5b9291505056fea26469706673582212200ff52428adf4fb258d2613ecb9f1a797504fedc523f527f26f7242fe1b6bc68d64736f6c634300081c0033",
}

// LpABI is the input ABI used to generate the binding from.
// Deprecated: Use LpMetaData.ABI instead.
var LpABI = LpMetaData.ABI

// LpBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LpMetaData.Bin instead.
var LpBin = LpMetaData.Bin

// DeployLp deploys a new Ethereum contract, binding an instance of Lp to it.
func DeployLp(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, owner common.Address) (common.Address, *types.Transaction, *Lp, error) {
	parsed, err := LpMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LpBin), backend, name, symbol, owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Lp{LpCaller: LpCaller{contract: contract}, LpTransactor: LpTransactor{contract: contract}, LpFilterer: LpFilterer{contract: contract}}, nil
}

// Lp is an auto generated Go binding around an Ethereum contract.
type Lp struct {
	LpCaller     // Read-only binding to the contract
	LpTransactor // Write-only binding to the contract
	LpFilterer   // Log filterer for contract events
}

// LpCaller is an auto generated read-only Go binding around an Ethereum contract.
type LpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LpSession struct {
	Contract     *Lp               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LpCallerSession struct {
	Contract *LpCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LpTransactorSession struct {
	Contract     *LpTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LpRaw is an auto generated low-level Go binding around an Ethereum contract.
type LpRaw struct {
	Contract *Lp // Generic contract binding to access the raw methods on
}

// LpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LpCallerRaw struct {
	Contract *LpCaller // Generic read-only contract binding to access the raw methods on
}

// LpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LpTransactorRaw struct {
	Contract *LpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLp creates a new instance of Lp, bound to a specific deployed contract.
func NewLp(address common.Address, backend bind.ContractBackend) (*Lp, error) {
	contract, err := bindLp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lp{LpCaller: LpCaller{contract: contract}, LpTransactor: LpTransactor{contract: contract}, LpFilterer: LpFilterer{contract: contract}}, nil
}

// NewLpCaller creates a new read-only instance of Lp, bound to a specific deployed contract.
func NewLpCaller(address common.Address, caller bind.ContractCaller) (*LpCaller, error) {
	contract, err := bindLp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LpCaller{contract: contract}, nil
}

// NewLpTransactor creates a new write-only instance of Lp, bound to a specific deployed contract.
func NewLpTransactor(address common.Address, transactor bind.ContractTransactor) (*LpTransactor, error) {
	contract, err := bindLp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LpTransactor{contract: contract}, nil
}

// NewLpFilterer creates a new log filterer instance of Lp, bound to a specific deployed contract.
func NewLpFilterer(address common.Address, filterer bind.ContractFilterer) (*LpFilterer, error) {
	contract, err := bindLp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LpFilterer{contract: contract}, nil
}

// bindLp binds a generic wrapper to an already deployed contract.
func bindLp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LpMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lp *LpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lp.Contract.LpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lp *LpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lp.Contract.LpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lp *LpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lp.Contract.LpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lp *LpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lp *LpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lp *LpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lp.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Lp *LpCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Lp *LpSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Lp.Contract.Allowance(&_Lp.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Lp *LpCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Lp.Contract.Allowance(&_Lp.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Lp *LpCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Lp *LpSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Lp.Contract.BalanceOf(&_Lp.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Lp *LpCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Lp.Contract.BalanceOf(&_Lp.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Lp *LpCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Lp *LpSession) Decimals() (uint8, error) {
	return _Lp.Contract.Decimals(&_Lp.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Lp *LpCallerSession) Decimals() (uint8, error) {
	return _Lp.Contract.Decimals(&_Lp.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Lp *LpCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Lp *LpSession) Name() (string, error) {
	return _Lp.Contract.Name(&_Lp.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Lp *LpCallerSession) Name() (string, error) {
	return _Lp.Contract.Name(&_Lp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lp *LpCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lp *LpSession) Owner() (common.Address, error) {
	return _Lp.Contract.Owner(&_Lp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lp *LpCallerSession) Owner() (common.Address, error) {
	return _Lp.Contract.Owner(&_Lp.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Lp *LpCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Lp *LpSession) Symbol() (string, error) {
	return _Lp.Contract.Symbol(&_Lp.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Lp *LpCallerSession) Symbol() (string, error) {
	return _Lp.Contract.Symbol(&_Lp.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lp *LpCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lp.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lp *LpSession) TotalSupply() (*big.Int, error) {
	return _Lp.Contract.TotalSupply(&_Lp.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lp *LpCallerSession) TotalSupply() (*big.Int, error) {
	return _Lp.Contract.TotalSupply(&_Lp.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Lp *LpTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Lp *LpSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Approve(&_Lp.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Lp *LpTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Approve(&_Lp.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_Lp *LpTransactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_Lp *LpSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Burn(&_Lp.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_Lp *LpTransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Burn(&_Lp.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Lp *LpTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Lp *LpSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Mint(&_Lp.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_Lp *LpTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Mint(&_Lp.TransactOpts, account, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lp *LpTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lp *LpSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lp.Contract.RenounceOwnership(&_Lp.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lp *LpTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lp.Contract.RenounceOwnership(&_Lp.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Lp *LpTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Lp *LpSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Transfer(&_Lp.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Lp *LpTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.Transfer(&_Lp.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Lp *LpTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Lp *LpSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.TransferFrom(&_Lp.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Lp *LpTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Lp.Contract.TransferFrom(&_Lp.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lp *LpTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Lp.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lp *LpSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lp.Contract.TransferOwnership(&_Lp.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lp *LpTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lp.Contract.TransferOwnership(&_Lp.TransactOpts, newOwner)
}

// LpApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Lp contract.
type LpApprovalIterator struct {
	Event *LpApproval // Event containing the contract specifics and raw log

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
func (it *LpApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpApproval)
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
		it.Event = new(LpApproval)
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
func (it *LpApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpApproval represents a Approval event raised by the Lp contract.
type LpApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Lp *LpFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LpApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Lp.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LpApprovalIterator{contract: _Lp.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Lp *LpFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LpApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Lp.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpApproval)
				if err := _Lp.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Lp *LpFilterer) ParseApproval(log types.Log) (*LpApproval, error) {
	event := new(LpApproval)
	if err := _Lp.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Lp contract.
type LpOwnershipTransferredIterator struct {
	Event *LpOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LpOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpOwnershipTransferred)
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
		it.Event = new(LpOwnershipTransferred)
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
func (it *LpOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpOwnershipTransferred represents a OwnershipTransferred event raised by the Lp contract.
type LpOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lp *LpFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LpOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lp.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LpOwnershipTransferredIterator{contract: _Lp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lp *LpFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LpOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lp.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpOwnershipTransferred)
				if err := _Lp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Lp *LpFilterer) ParseOwnershipTransferred(log types.Log) (*LpOwnershipTransferred, error) {
	event := new(LpOwnershipTransferred)
	if err := _Lp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LpTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Lp contract.
type LpTransferIterator struct {
	Event *LpTransfer // Event containing the contract specifics and raw log

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
func (it *LpTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LpTransfer)
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
		it.Event = new(LpTransfer)
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
func (it *LpTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LpTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LpTransfer represents a Transfer event raised by the Lp contract.
type LpTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Lp *LpFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LpTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Lp.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LpTransferIterator{contract: _Lp.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Lp *LpFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LpTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Lp.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LpTransfer)
				if err := _Lp.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Lp *LpFilterer) ParseTransfer(log types.Log) (*LpTransfer, error) {
	event := new(LpTransfer)
	if err := _Lp.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
