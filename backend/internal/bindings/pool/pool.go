// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pool

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

// PoolMetaData contains all meta data concerning the Pool contract.
var PoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredAmount\",\"type\":\"uint256\"}],\"name\":\"BadSlippage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInputAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"WrongLiquidityAmout\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"}],\"name\":\"AddLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"zeroToOne\",\"type\":\"bool\"}],\"name\":\"SwapTokens\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"zeroToOne\",\"type\":\"bool\"}],\"name\":\"getAmountToAdd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getAmountsFromLp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outReserve\",\"type\":\"uint256\"}],\"name\":\"getOutputAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserve0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserve1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"zeroToOne\",\"type\":\"bool\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x610120604052348015610010575f5ffd5b503373ffffffffffffffffffffffffffffffffffffffff1663890357306040518163ffffffff1660e01b815260040160a060405180830381865afa15801561005a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061007e91906101fc565b8062ffffff166101009062ffffff168152508173ffffffffffffffffffffffffffffffffffffffff1660e09073ffffffffffffffffffffffffffffffffffffffff168152508273ffffffffffffffffffffffffffffffffffffffff1660c09073ffffffffffffffffffffffffffffffffffffffff168152508373ffffffffffffffffffffffffffffffffffffffff1660a09073ffffffffffffffffffffffffffffffffffffffff168152508473ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff168152505050505050610273565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101938261016a565b9050919050565b6101a381610189565b81146101ad575f5ffd5b50565b5f815190506101be8161019a565b92915050565b5f62ffffff82169050919050565b6101db816101c4565b81146101e5575f5ffd5b50565b5f815190506101f6816101d2565b92915050565b5f5f5f5f5f60a0868803121561021557610214610166565b5b5f610222888289016101b0565b9550506020610233888289016101b0565b9450506040610244888289016101b0565b9350506060610255888289016101b0565b9250506080610266888289016101e8565b9150509295509295909350565b60805160a05160c05160e051610100516116c061033c5f395f8181610ea3015261112301525f81816105950152818161076501528181610ac601528181610ce90152610f2401525f81816103b101528181610419015281816105e10152818161088c01528181610a2801528181610c4a01528181610e7d015261106701525f81816102c3015281816103d2015281816103f80152818161067f015281816107f00152818161098a01528181610bac0152610fb601525f81816105ba0152610e5901526116c05ff3fe608060405234801561000f575f5ffd5b50600436106100e8575f3560e01c80639cd441da1161008a578063d3fdbbb511610064578063d3fdbbb514610212578063d813478e14610242578063ddca3f4314610273578063ef01a1c314610291576100e8565b80639cd441da146101ba578063c45a0155146101d6578063d21220a7146101f4576100e8565b806388cc58e4116100c657806388cc58e41461014457806390dc67e61461016257806399215eef146101805780639c8f9f231461019e576100e8565b80630dfe1681146100ec5780632bec60fe1461010a5780635fcbd28514610126575b5f5ffd5b6100f46102c1565b60405161010191906111e7565b60405180910390f35b610124600480360381019061011f919061126c565b6102e5565b005b61012e610593565b60405161013b91906111e7565b60405180910390f35b61014c6105b7565b60405161015991906111e7565b60405180910390f35b61016a6105de565b60405161017791906112cb565b60405180910390f35b61018861067c565b60405161019591906112cb565b60405180910390f35b6101b860048036038101906101b391906112e4565b61071a565b005b6101d460048036038101906101cf919061130f565b61097a565b005b6101de610e57565b6040516101eb91906111e7565b60405180910390f35b6101fc610e7b565b60405161020991906111e7565b60405180910390f35b61022c6004803603810190610227919061134d565b610e9f565b60405161023991906112cb565b60405180910390f35b61025c600480360381019061025791906112e4565b610f1f565b60405161026a92919061139d565b60405180910390f35b61027b611121565b60405161028891906113e1565b60405180910390f35b6102ab60048036038101906102a691906113fa565b611145565b6040516102b891906112cb565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000081565b5f831161031e576040517f340dabef00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f8261033a5761032d6105de565b61033561067c565b61034b565b61034261067c565b61034a6105de565b5b915091505f61035b868484610e9f565b905084811015858290916103a6576040517f3fad68d000000000000000000000000000000000000000000000000000000000815260040161039d92919061139d565b60405180910390fd5b50505f5f856103f6577f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000610439565b7f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000005b915091508173ffffffffffffffffffffffffffffffffffffffff166323b872dd33308b6040518463ffffffff1660e01b815260040161047a93929190611438565b6020604051808303815f875af1158015610496573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104ba9190611481565b508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33856040518363ffffffff1660e01b81526004016104f69291906114ac565b6020604051808303815f875af1158015610512573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105369190611481565b503373ffffffffffffffffffffffffffffffffffffffff167f5aceb8ddbdcc435072e442e4dd594777e5575cb8c3e8ab1ff1e5b76d7aeaa6c4898589604051610581939291906114e2565b60405180910390a25050505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f7f0000000000000000000000000000000000000000000000000000000000000000905090565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161063891906111e7565b602060405180830381865afa158015610653573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610677919061152b565b905090565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016106d691906111e7565b602060405180830381865afa1580156106f1573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610715919061152b565b905090565b5f8111610753576040517f340dabef00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5f61075e83610f1f565b915091505f7f000000000000000000000000000000000000000000000000000000000000000090508073ffffffffffffffffffffffffffffffffffffffff16639dc29fac33866040518363ffffffff1660e01b81526004016107c19291906114ac565b5f604051808303815f87803b1580156107d8575f5ffd5b505af11580156107ea573d5f5f3e3d5ffd5b505050507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33856040518363ffffffff1660e01b81526004016108499291906114ac565b6020604051808303815f875af1158015610865573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108899190611481565b507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33846040518363ffffffff1660e01b81526004016108e59291906114ac565b6020604051808303815f875af1158015610901573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109259190611481565b503373ffffffffffffffffffffffffffffffffffffffff167fdfdd120ded9b7afc0c23dd5310e93aaa3e1c3b9f75af9b805fab3030842439f28560405161096c91906112cb565b60405180910390a250505050565b5f61098361067c565b03610b51577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b81526004016109e593929190611438565b6020604051808303815f875af1158015610a01573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a259190611481565b507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b8152600401610a8393929190611438565b6020604051808303815f875af1158015610a9f573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ac39190611481565b507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166340c10f1933846040518363ffffffff1660e01b8152600401610b1f9291906114ac565b5f604051808303815f87803b158015610b36575f5ffd5b505af1158015610b48573d5f5f3e3d5ffd5b50505050610e03565b5f610b5d836001611145565b90508082101581839091610ba8576040517ff9e88c81000000000000000000000000000000000000000000000000000000008152600401610b9f92919061139d565b60405180910390fd5b50507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd3330866040518463ffffffff1660e01b8152600401610c0793929190611438565b6020604051808303815f875af1158015610c23573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c479190611481565b507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b8152600401610ca593929190611438565b6020604051808303815f875af1158015610cc1573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ce59190611481565b505f7f000000000000000000000000000000000000000000000000000000000000000090505f610d1361067c565b858373ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d5d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d81919061152b565b610d8b9190611583565b610d9591906115f1565b90508173ffffffffffffffffffffffffffffffffffffffff166340c10f1933836040518363ffffffff1660e01b8152600401610dd29291906114ac565b5f604051808303815f87803b158015610de9575f5ffd5b505af1158015610dfb573d5f5f3e3d5ffd5b505050505050505b3373ffffffffffffffffffffffffffffffffffffffff167f06239653922ac7bea6aa2b19dc486b9361821d37712eb796adfd38d81de278ca8383604051610e4b92919061139d565b60405180910390a25050565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f5f7f00000000000000000000000000000000000000000000000000000000000000006103e8610ecf9190611621565b62ffffff1685610edf9190611583565b90505f816103e886610ef19190611583565b610efb9190611657565b8483610f079190611583565b610f1191906115f1565b905080925050509392505050565b5f5f5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f8b573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610faf919061152b565b90505f81857f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161100d91906111e7565b602060405180830381865afa158015611028573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061104c919061152b565b6110569190611583565b61106091906115f1565b90505f82867f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016110be91906111e7565b602060405180830381865afa1580156110d9573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110fd919061152b565b6111079190611583565b61111191906115f1565b9050818194509450505050915091565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f5f5f83611162576111556105de565b61115d61067c565b611173565b61116a61067c565b6111726105de565b5b915091505f85836111849190611657565b82876111909190611583565b61119a91906115f1565b905080935050505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6111d1826111a8565b9050919050565b6111e1816111c7565b82525050565b5f6020820190506111fa5f8301846111d8565b92915050565b5f5ffd5b5f819050919050565b61121681611204565b8114611220575f5ffd5b50565b5f813590506112318161120d565b92915050565b5f8115159050919050565b61124b81611237565b8114611255575f5ffd5b50565b5f8135905061126681611242565b92915050565b5f5f5f6060848603121561128357611282611200565b5b5f61129086828701611223565b93505060206112a186828701611223565b92505060406112b286828701611258565b9150509250925092565b6112c581611204565b82525050565b5f6020820190506112de5f8301846112bc565b92915050565b5f602082840312156112f9576112f8611200565b5b5f61130684828501611223565b91505092915050565b5f5f6040838503121561132557611324611200565b5b5f61133285828601611223565b925050602061134385828601611223565b9150509250929050565b5f5f5f6060848603121561136457611363611200565b5b5f61137186828701611223565b935050602061138286828701611223565b925050604061139386828701611223565b9150509250925092565b5f6040820190506113b05f8301856112bc565b6113bd60208301846112bc565b9392505050565b5f62ffffff82169050919050565b6113db816113c4565b82525050565b5f6020820190506113f45f8301846113d2565b92915050565b5f5f604083850312156114105761140f611200565b5b5f61141d85828601611223565b925050602061142e85828601611258565b9150509250929050565b5f60608201905061144b5f8301866111d8565b61145860208301856111d8565b61146560408301846112bc565b949350505050565b5f8151905061147b81611242565b92915050565b5f6020828403121561149657611495611200565b5b5f6114a38482850161146d565b91505092915050565b5f6040820190506114bf5f8301856111d8565b6114cc60208301846112bc565b9392505050565b6114dc81611237565b82525050565b5f6060820190506114f55f8301866112bc565b61150260208301856112bc565b61150f60408301846114d3565b949350505050565b5f815190506115258161120d565b92915050565b5f602082840312156115405761153f611200565b5b5f61154d84828501611517565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61158d82611204565b915061159883611204565b92508282026115a681611204565b915082820484148315176115bd576115bc611556565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6115fb82611204565b915061160683611204565b925082611616576116156115c4565b5b828204905092915050565b5f61162b826113c4565b9150611636836113c4565b9250828203905062ffffff81111561165157611650611556565b5b92915050565b5f61166182611204565b915061166c83611204565b925082820190508082111561168457611683611556565b5b9291505056fea2646970667358221220699c01f1500b19f452a2a87edc3a77c9b2f8566544d47abc50ce93ff598bd8f764736f6c634300081c0033",
}

// PoolABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolMetaData.ABI instead.
var PoolABI = PoolMetaData.ABI

// PoolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PoolMetaData.Bin instead.
var PoolBin = PoolMetaData.Bin

// DeployPool deploys a new Ethereum contract, binding an instance of Pool to it.
func DeployPool(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pool, error) {
	parsed, err := PoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PoolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// Pool is an auto generated Go binding around an Ethereum contract.
type Pool struct {
	PoolCaller     // Read-only binding to the contract
	PoolTransactor // Write-only binding to the contract
	PoolFilterer   // Log filterer for contract events
}

// PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolSession struct {
	Contract     *Pool             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolCallerSession struct {
	Contract *PoolCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolTransactorSession struct {
	Contract     *PoolTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolRaw struct {
	Contract *Pool // Generic contract binding to access the raw methods on
}

// PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolCallerRaw struct {
	Contract *PoolCaller // Generic read-only contract binding to access the raw methods on
}

// PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolTransactorRaw struct {
	Contract *PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPool creates a new instance of Pool, bound to a specific deployed contract.
func NewPool(address common.Address, backend bind.ContractBackend) (*Pool, error) {
	contract, err := bindPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// NewPoolCaller creates a new read-only instance of Pool, bound to a specific deployed contract.
func NewPoolCaller(address common.Address, caller bind.ContractCaller) (*PoolCaller, error) {
	contract, err := bindPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolCaller{contract: contract}, nil
}

// NewPoolTransactor creates a new write-only instance of Pool, bound to a specific deployed contract.
func NewPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolTransactor, error) {
	contract, err := bindPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolTransactor{contract: contract}, nil
}

// NewPoolFilterer creates a new log filterer instance of Pool, bound to a specific deployed contract.
func NewPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolFilterer, error) {
	contract, err := bindPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolFilterer{contract: contract}, nil
}

// bindPool binds a generic wrapper to an already deployed contract.
func bindPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Pool *PoolCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Pool *PoolSession) Factory() (common.Address, error) {
	return _Pool.Contract.Factory(&_Pool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Pool *PoolCallerSession) Factory() (common.Address, error) {
	return _Pool.Contract.Factory(&_Pool.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_Pool *PoolCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_Pool *PoolSession) Fee() (*big.Int, error) {
	return _Pool.Contract.Fee(&_Pool.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint24)
func (_Pool *PoolCallerSession) Fee() (*big.Int, error) {
	return _Pool.Contract.Fee(&_Pool.CallOpts)
}

// GetAmountToAdd is a free data retrieval call binding the contract method 0xef01a1c3.
//
// Solidity: function getAmountToAdd(uint256 amountIn, bool zeroToOne) view returns(uint256)
func (_Pool *PoolCaller) GetAmountToAdd(opts *bind.CallOpts, amountIn *big.Int, zeroToOne bool) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getAmountToAdd", amountIn, zeroToOne)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountToAdd is a free data retrieval call binding the contract method 0xef01a1c3.
//
// Solidity: function getAmountToAdd(uint256 amountIn, bool zeroToOne) view returns(uint256)
func (_Pool *PoolSession) GetAmountToAdd(amountIn *big.Int, zeroToOne bool) (*big.Int, error) {
	return _Pool.Contract.GetAmountToAdd(&_Pool.CallOpts, amountIn, zeroToOne)
}

// GetAmountToAdd is a free data retrieval call binding the contract method 0xef01a1c3.
//
// Solidity: function getAmountToAdd(uint256 amountIn, bool zeroToOne) view returns(uint256)
func (_Pool *PoolCallerSession) GetAmountToAdd(amountIn *big.Int, zeroToOne bool) (*big.Int, error) {
	return _Pool.Contract.GetAmountToAdd(&_Pool.CallOpts, amountIn, zeroToOne)
}

// GetAmountsFromLp is a free data retrieval call binding the contract method 0xd813478e.
//
// Solidity: function getAmountsFromLp(uint256 amount) view returns(uint256, uint256)
func (_Pool *PoolCaller) GetAmountsFromLp(opts *bind.CallOpts, amount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getAmountsFromLp", amount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetAmountsFromLp is a free data retrieval call binding the contract method 0xd813478e.
//
// Solidity: function getAmountsFromLp(uint256 amount) view returns(uint256, uint256)
func (_Pool *PoolSession) GetAmountsFromLp(amount *big.Int) (*big.Int, *big.Int, error) {
	return _Pool.Contract.GetAmountsFromLp(&_Pool.CallOpts, amount)
}

// GetAmountsFromLp is a free data retrieval call binding the contract method 0xd813478e.
//
// Solidity: function getAmountsFromLp(uint256 amount) view returns(uint256, uint256)
func (_Pool *PoolCallerSession) GetAmountsFromLp(amount *big.Int) (*big.Int, *big.Int, error) {
	return _Pool.Contract.GetAmountsFromLp(&_Pool.CallOpts, amount)
}

// GetFactory is a free data retrieval call binding the contract method 0x88cc58e4.
//
// Solidity: function getFactory() view returns(address)
func (_Pool *PoolCaller) GetFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFactory is a free data retrieval call binding the contract method 0x88cc58e4.
//
// Solidity: function getFactory() view returns(address)
func (_Pool *PoolSession) GetFactory() (common.Address, error) {
	return _Pool.Contract.GetFactory(&_Pool.CallOpts)
}

// GetFactory is a free data retrieval call binding the contract method 0x88cc58e4.
//
// Solidity: function getFactory() view returns(address)
func (_Pool *PoolCallerSession) GetFactory() (common.Address, error) {
	return _Pool.Contract.GetFactory(&_Pool.CallOpts)
}

// GetOutputAmount is a free data retrieval call binding the contract method 0xd3fdbbb5.
//
// Solidity: function getOutputAmount(uint256 amount, uint256 inReserve, uint256 outReserve) view returns(uint256)
func (_Pool *PoolCaller) GetOutputAmount(opts *bind.CallOpts, amount *big.Int, inReserve *big.Int, outReserve *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getOutputAmount", amount, inReserve, outReserve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOutputAmount is a free data retrieval call binding the contract method 0xd3fdbbb5.
//
// Solidity: function getOutputAmount(uint256 amount, uint256 inReserve, uint256 outReserve) view returns(uint256)
func (_Pool *PoolSession) GetOutputAmount(amount *big.Int, inReserve *big.Int, outReserve *big.Int) (*big.Int, error) {
	return _Pool.Contract.GetOutputAmount(&_Pool.CallOpts, amount, inReserve, outReserve)
}

// GetOutputAmount is a free data retrieval call binding the contract method 0xd3fdbbb5.
//
// Solidity: function getOutputAmount(uint256 amount, uint256 inReserve, uint256 outReserve) view returns(uint256)
func (_Pool *PoolCallerSession) GetOutputAmount(amount *big.Int, inReserve *big.Int, outReserve *big.Int) (*big.Int, error) {
	return _Pool.Contract.GetOutputAmount(&_Pool.CallOpts, amount, inReserve, outReserve)
}

// GetReserve0 is a free data retrieval call binding the contract method 0x99215eef.
//
// Solidity: function getReserve0() view returns(uint256)
func (_Pool *PoolCaller) GetReserve0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getReserve0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserve0 is a free data retrieval call binding the contract method 0x99215eef.
//
// Solidity: function getReserve0() view returns(uint256)
func (_Pool *PoolSession) GetReserve0() (*big.Int, error) {
	return _Pool.Contract.GetReserve0(&_Pool.CallOpts)
}

// GetReserve0 is a free data retrieval call binding the contract method 0x99215eef.
//
// Solidity: function getReserve0() view returns(uint256)
func (_Pool *PoolCallerSession) GetReserve0() (*big.Int, error) {
	return _Pool.Contract.GetReserve0(&_Pool.CallOpts)
}

// GetReserve1 is a free data retrieval call binding the contract method 0x90dc67e6.
//
// Solidity: function getReserve1() view returns(uint256)
func (_Pool *PoolCaller) GetReserve1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getReserve1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserve1 is a free data retrieval call binding the contract method 0x90dc67e6.
//
// Solidity: function getReserve1() view returns(uint256)
func (_Pool *PoolSession) GetReserve1() (*big.Int, error) {
	return _Pool.Contract.GetReserve1(&_Pool.CallOpts)
}

// GetReserve1 is a free data retrieval call binding the contract method 0x90dc67e6.
//
// Solidity: function getReserve1() view returns(uint256)
func (_Pool *PoolCallerSession) GetReserve1() (*big.Int, error) {
	return _Pool.Contract.GetReserve1(&_Pool.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_Pool *PoolCaller) LpToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "lpToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_Pool *PoolSession) LpToken() (common.Address, error) {
	return _Pool.Contract.LpToken(&_Pool.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x5fcbd285.
//
// Solidity: function lpToken() view returns(address)
func (_Pool *PoolCallerSession) LpToken() (common.Address, error) {
	return _Pool.Contract.LpToken(&_Pool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Pool *PoolCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Pool *PoolSession) Token0() (common.Address, error) {
	return _Pool.Contract.Token0(&_Pool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Pool *PoolCallerSession) Token0() (common.Address, error) {
	return _Pool.Contract.Token0(&_Pool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Pool *PoolCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Pool *PoolSession) Token1() (common.Address, error) {
	return _Pool.Contract.Token1(&_Pool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Pool *PoolCallerSession) Token1() (common.Address, error) {
	return _Pool.Contract.Token1(&_Pool.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9cd441da.
//
// Solidity: function addLiquidity(uint256 amount0In, uint256 amount1In) returns()
func (_Pool *PoolTransactor) AddLiquidity(opts *bind.TransactOpts, amount0In *big.Int, amount1In *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "addLiquidity", amount0In, amount1In)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9cd441da.
//
// Solidity: function addLiquidity(uint256 amount0In, uint256 amount1In) returns()
func (_Pool *PoolSession) AddLiquidity(amount0In *big.Int, amount1In *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.AddLiquidity(&_Pool.TransactOpts, amount0In, amount1In)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9cd441da.
//
// Solidity: function addLiquidity(uint256 amount0In, uint256 amount1In) returns()
func (_Pool *PoolTransactorSession) AddLiquidity(amount0In *big.Int, amount1In *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.AddLiquidity(&_Pool.TransactOpts, amount0In, amount1In)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 amount) returns()
func (_Pool *PoolTransactor) RemoveLiquidity(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "removeLiquidity", amount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 amount) returns()
func (_Pool *PoolSession) RemoveLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RemoveLiquidity(&_Pool.TransactOpts, amount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 amount) returns()
func (_Pool *PoolTransactorSession) RemoveLiquidity(amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RemoveLiquidity(&_Pool.TransactOpts, amount)
}

// Swap is a paid mutator transaction binding the contract method 0x2bec60fe.
//
// Solidity: function swap(uint256 amountIn, uint256 amountOutMin, bool zeroToOne) returns()
func (_Pool *PoolTransactor) Swap(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, zeroToOne bool) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "swap", amountIn, amountOutMin, zeroToOne)
}

// Swap is a paid mutator transaction binding the contract method 0x2bec60fe.
//
// Solidity: function swap(uint256 amountIn, uint256 amountOutMin, bool zeroToOne) returns()
func (_Pool *PoolSession) Swap(amountIn *big.Int, amountOutMin *big.Int, zeroToOne bool) (*types.Transaction, error) {
	return _Pool.Contract.Swap(&_Pool.TransactOpts, amountIn, amountOutMin, zeroToOne)
}

// Swap is a paid mutator transaction binding the contract method 0x2bec60fe.
//
// Solidity: function swap(uint256 amountIn, uint256 amountOutMin, bool zeroToOne) returns()
func (_Pool *PoolTransactorSession) Swap(amountIn *big.Int, amountOutMin *big.Int, zeroToOne bool) (*types.Transaction, error) {
	return _Pool.Contract.Swap(&_Pool.TransactOpts, amountIn, amountOutMin, zeroToOne)
}

// PoolAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the Pool contract.
type PoolAddLiquidityIterator struct {
	Event *PoolAddLiquidity // Event containing the contract specifics and raw log

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
func (it *PoolAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAddLiquidity)
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
		it.Event = new(PoolAddLiquidity)
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
func (it *PoolAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAddLiquidity represents a AddLiquidity event raised by the Pool contract.
type PoolAddLiquidity struct {
	Account   common.Address
	Amount0In *big.Int
	Amount1In *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x06239653922ac7bea6aa2b19dc486b9361821d37712eb796adfd38d81de278ca.
//
// Solidity: event AddLiquidity(address indexed account, uint256 amount0In, uint256 amount1In)
func (_Pool *PoolFilterer) FilterAddLiquidity(opts *bind.FilterOpts, account []common.Address) (*PoolAddLiquidityIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "AddLiquidity", accountRule)
	if err != nil {
		return nil, err
	}
	return &PoolAddLiquidityIterator{contract: _Pool.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x06239653922ac7bea6aa2b19dc486b9361821d37712eb796adfd38d81de278ca.
//
// Solidity: event AddLiquidity(address indexed account, uint256 amount0In, uint256 amount1In)
func (_Pool *PoolFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *PoolAddLiquidity, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "AddLiquidity", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAddLiquidity)
				if err := _Pool.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0x06239653922ac7bea6aa2b19dc486b9361821d37712eb796adfd38d81de278ca.
//
// Solidity: event AddLiquidity(address indexed account, uint256 amount0In, uint256 amount1In)
func (_Pool *PoolFilterer) ParseAddLiquidity(log types.Log) (*PoolAddLiquidity, error) {
	event := new(PoolAddLiquidity)
	if err := _Pool.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the Pool contract.
type PoolRemoveLiquidityIterator struct {
	Event *PoolRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *PoolRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolRemoveLiquidity)
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
		it.Event = new(PoolRemoveLiquidity)
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
func (it *PoolRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolRemoveLiquidity represents a RemoveLiquidity event raised by the Pool contract.
type PoolRemoveLiquidity struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0xdfdd120ded9b7afc0c23dd5310e93aaa3e1c3b9f75af9b805fab3030842439f2.
//
// Solidity: event RemoveLiquidity(address indexed account, uint256 amount)
func (_Pool *PoolFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, account []common.Address) (*PoolRemoveLiquidityIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "RemoveLiquidity", accountRule)
	if err != nil {
		return nil, err
	}
	return &PoolRemoveLiquidityIterator{contract: _Pool.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0xdfdd120ded9b7afc0c23dd5310e93aaa3e1c3b9f75af9b805fab3030842439f2.
//
// Solidity: event RemoveLiquidity(address indexed account, uint256 amount)
func (_Pool *PoolFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *PoolRemoveLiquidity, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "RemoveLiquidity", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolRemoveLiquidity)
				if err := _Pool.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0xdfdd120ded9b7afc0c23dd5310e93aaa3e1c3b9f75af9b805fab3030842439f2.
//
// Solidity: event RemoveLiquidity(address indexed account, uint256 amount)
func (_Pool *PoolFilterer) ParseRemoveLiquidity(log types.Log) (*PoolRemoveLiquidity, error) {
	event := new(PoolRemoveLiquidity)
	if err := _Pool.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolSwapTokensIterator is returned from FilterSwapTokens and is used to iterate over the raw logs and unpacked data for SwapTokens events raised by the Pool contract.
type PoolSwapTokensIterator struct {
	Event *PoolSwapTokens // Event containing the contract specifics and raw log

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
func (it *PoolSwapTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolSwapTokens)
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
		it.Event = new(PoolSwapTokens)
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
func (it *PoolSwapTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolSwapTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolSwapTokens represents a SwapTokens event raised by the Pool contract.
type PoolSwapTokens struct {
	Account   common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	ZeroToOne bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwapTokens is a free log retrieval operation binding the contract event 0x5aceb8ddbdcc435072e442e4dd594777e5575cb8c3e8ab1ff1e5b76d7aeaa6c4.
//
// Solidity: event SwapTokens(address indexed account, uint256 amountIn, uint256 amountOut, bool zeroToOne)
func (_Pool *PoolFilterer) FilterSwapTokens(opts *bind.FilterOpts, account []common.Address) (*PoolSwapTokensIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "SwapTokens", accountRule)
	if err != nil {
		return nil, err
	}
	return &PoolSwapTokensIterator{contract: _Pool.contract, event: "SwapTokens", logs: logs, sub: sub}, nil
}

// WatchSwapTokens is a free log subscription operation binding the contract event 0x5aceb8ddbdcc435072e442e4dd594777e5575cb8c3e8ab1ff1e5b76d7aeaa6c4.
//
// Solidity: event SwapTokens(address indexed account, uint256 amountIn, uint256 amountOut, bool zeroToOne)
func (_Pool *PoolFilterer) WatchSwapTokens(opts *bind.WatchOpts, sink chan<- *PoolSwapTokens, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "SwapTokens", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolSwapTokens)
				if err := _Pool.contract.UnpackLog(event, "SwapTokens", log); err != nil {
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

// ParseSwapTokens is a log parse operation binding the contract event 0x5aceb8ddbdcc435072e442e4dd594777e5575cb8c3e8ab1ff1e5b76d7aeaa6c4.
//
// Solidity: event SwapTokens(address indexed account, uint256 amountIn, uint256 amountOut, bool zeroToOne)
func (_Pool *PoolFilterer) ParseSwapTokens(log types.Log) (*PoolSwapTokens, error) {
	event := new(PoolSwapTokens)
	if err := _Pool.contract.UnpackLog(event, "SwapTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
