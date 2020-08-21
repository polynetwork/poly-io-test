// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DaiABI is the input ABI used to generate the binding from.
const DaiABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"usr\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"usr\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"holder\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\"},{\"name\":\"allowed\",\"type\":\"bool\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"usr\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"usr\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"push\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"move\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"usr\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"pull\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"chainId_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"}]"

// DaiFuncSigs maps the 4-byte function signature to its string representation.
var DaiFuncSigs = map[string]string{
	"3644e515": "DOMAIN_SEPARATOR()",
	"30adf81f": "PERMIT_TYPEHASH()",
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"9dc29fac": "burn(address,uint256)",
	"313ce567": "decimals()",
	"9c52a7f1": "deny(address)",
	"40c10f19": "mint(address,uint256)",
	"bb35783b": "move(address,address,uint256)",
	"06fdde03": "name()",
	"7ecebe00": "nonces(address)",
	"8fcbaf0c": "permit(address,address,uint256,uint256,bool,uint8,bytes32,bytes32)",
	"f2d5d56b": "pull(address,uint256)",
	"b753a98c": "push(address,uint256)",
	"65fae35e": "rely(address)",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"54fd4d50": "version()",
	"bf353dbb": "wards(address)",
}

// DaiBin is the compiled bytecode used for deploying new contracts.
var DaiBin = "0x608060405234801561001057600080fd5b506040516020806114398339810160408181529151336000908152602081815290849020600190557f454950373132446f6d61696e28737472696e67206e616d652c737472696e672083527f76657273696f6e2c75696e7432353620636861696e49642c6164647265737320818401527f766572696679696e67436f6e74726163742900000000000000000000000000008484015283519283900360520183208385018552600e8085527f44616920537461626c65636f696e000000000000000000000000000000000000928501928352945192949093929182918083835b6020831061010e5780518252601f1990920191602091820191016100ef565b51815160209384036101000a600019018019909216911617905260408051929094018290038220828501855260018084527f3100000000000000000000000000000000000000000000000000000000000000928401928352945190965091945090928392508083835b602083106101965780518252601f199092019160209182019101610177565b51815160209384036101000a6000190180199092169116179052604080519290940182900382208282019890985281840196909652606081019690965250608085018690523060a0808701919091528151808703909101815260c09095019081905284519093849350850191508083835b602083106102265780518252601f199092019160209182019101610207565b5181516000196020949094036101000a939093019283169219169190911790526040519201829003909120600555505050506111d2806102676000396000f3006080604052600436106101275763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde03811461012c578063095ea7b3146101b857806318160ddd146101f057806323b872dd1461021757806330adf81f14610241578063313ce567146102565780633644e5151461028157806340c10f191461029657806354fd4d50146102bc57806365fae35e146102d157806370a08231146102f25780637ecebe00146103135780638fcbaf0c1461033457806395d89b41146103725780639c52a7f1146103875780639dc29fac146103a8578063a9059cbb146103cc578063b753a98c146103f0578063bb35783b14610414578063bf353dbb1461043e578063dd62ed3e1461045f578063f2d5d56b14610486575b600080fd5b34801561013857600080fd5b506101416104aa565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561017d578082015183820152602001610165565b50505050905090810190601f1680156101aa5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101c457600080fd5b506101dc600160a060020a03600435166024356104e1565b604051901515815260200160405180910390f35b3480156101fc57600080fd5b50610205610553565b60405190815260200160405180910390f35b34801561022357600080fd5b506101dc600160a060020a0360043581169060243516604435610559565b34801561024d57600080fd5b506102056107b1565b34801561026257600080fd5b5061026b6107d5565b60405160ff909116815260200160405180910390f35b34801561028d57600080fd5b506102056107da565b3480156102a257600080fd5b506102ba600160a060020a03600435166024356107e0565b005b3480156102c857600080fd5b506101416108d5565b3480156102dd57600080fd5b506102ba600160a060020a036004351661090c565b3480156102fe57600080fd5b50610205600160a060020a03600435166109c4565b34801561031f57600080fd5b50610205600160a060020a03600435166109d8565b34801561034057600080fd5b506102ba600160a060020a0360043581169060243516604435606435608435151560ff60a4351660c43560e4356109ec565b34801561037e57600080fd5b50610141610dff565b34801561039357600080fd5b506102ba600160a060020a0360043516610e36565b3480156103b457600080fd5b506102ba600160a060020a0360043516602435610eeb565b3480156103d857600080fd5b506101dc600160a060020a0360043516602435611110565b3480156103fc57600080fd5b506102ba600160a060020a0360043516602435611124565b34801561042057600080fd5b506102ba600160a060020a0360043581169060243516604435611134565b34801561044a57600080fd5b50610205600160a060020a0360043516611145565b34801561046b57600080fd5b50610205600160a060020a0360043581169060243516611159565b34801561049257600080fd5b506102ba600160a060020a036004351660243561117b565b60408051908101604052600e81527f44616920537461626c65636f696e000000000000000000000000000000000000602082015281565b33600090815260036020528160408220600160a060020a03851660009081526020919091526040902055600160a060020a038316337f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405190815260200160405180910390a35060015b92915050565b60015481565b600160a060020a0383166000908152600260205281604082205410156105c85760405160e560020a62461bcd02815260206004820152601860248201527f4461692f696e73756666696369656e742d62616c616e63650000000000000000604482015260640160405180910390fd5b600160a060020a038416331480159061060c5750600160a060020a038416600090815260036020526000199060409020336000908152602091909152604090205414155b156106ee57600160a060020a03841660009081526003602052829060409020336000908152602091909152604090205410156106915760405160e560020a62461bcd02815260206004820152601a60248201527f4461692f696e73756666696369656e742d616c6c6f77616e6365000000000000604482015260640160405180910390fd5b600160a060020a038416600090815260036020526106c49060409020336000908152602091909152604090205483611186565b600160a060020a038516600090815260036020526040902033600090815260209190915260409020555b600160a060020a0384166000908152600260205261071190604090205483611186565b600160a060020a038516600090815260026020526040902055600160a060020a0383166000908152600260205261074d90604090205483611196565b600160a060020a038416600090815260026020526040902055600160a060020a038084169085167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405190815260200160405180910390a35060019392505050565b7fea2aa0a1be11a07ed86d755c93467f4f82362b452371d1ba94d1715123511acb81565b601281565b60055481565b336000908152602081905260409020546001146108465760405160e560020a62461bcd02815260206004820152601260248201527f4461692f6e6f742d617574686f72697a65640000000000000000000000000000604482015260640160405180910390fd5b600160a060020a0382166000908152600260205261086990604090205482611196565b600160a060020a03831660009081526002602052604090205560015461088f9082611196565b600155600160a060020a03821660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405190815260200160405180910390a35050565b60408051908101604052600181527f3100000000000000000000000000000000000000000000000000000000000000602082015281565b336000908152602081905260409020546001146109725760405160e560020a62461bcd02815260206004820152601260248201527f4461692f6e6f742d617574686f72697a65640000000000000000000000000000604482015260640160405180910390fd5b600160a060020a0381166000908152602081905260019060409020555961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b600260205280600052604060002054905081565b600460205280600052604060002054905081565b6000806005547fea2aa0a1be11a07ed86d755c93467f4f82362b452371d1ba94d1715123511acb6001028b8b8b8b8b6040516020810196909652600160a060020a03948516604080880191909152939094166060860152608085019190915260a084015290151560c083015260e090910190516020818303038152906040526040518082805190602001908083835b60208310610a9a5780518252601f199092019160209182019101610a7b565b6001836020036101000a038019825116818451161790925250505091909101925060409150505180910390206040517f19010000000000000000000000000000000000000000000000000000000000006020820152602281019290925260428201526062016040516020818303038152906040526040518082805190602001908083835b60208310610b3d5780518252601f199092019160209182019101610b1e565b6001836020036101000a038019825116818451161790925250505091909101925060409150505180910390209150600160a060020a038a161515610bca5760405160e560020a62461bcd02815260206004820152601560248201527f4461692f696e76616c69642d616464726573732d300000000000000000000000604482015260640160405180910390fd5b6001828686866040516000815260200160405260405193845260ff9092166020808501919091526040808501929092526060840192909252608090920191516020810390808403906000865af1158015610c28573d6000803e3d6000fd5b50505060206040510351600160a060020a038b8116911614610c935760405160e560020a62461bcd02815260206004820152601260248201527f4461692f696e76616c69642d7065726d69740000000000000000000000000000604482015260640160405180910390fd5b861580610ca05750864211155b1515610cf55760405160e560020a62461bcd02815260206004820152601260248201527f4461692f7065726d69742d657870697265640000000000000000000000000000604482015260640160405180910390fd5b600160a060020a038a1660009081526004602052604090208054600181019091558814610d6b5760405160e560020a62461bcd02815260206004820152601160248201527f4461692f696e76616c69642d6e6f6e6365000000000000000000000000000000604482015260640160405180910390fd5b85610d77576000610d7b565b6000195b600160a060020a038b1660009081526003602052909150819060409020600160a060020a038b1660009081526020919091526040902055600160a060020a03808a16908b167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258360405190815260200160405180910390a350505050505050505050565b60408051908101604052600381527f4441490000000000000000000000000000000000000000000000000000000000602082015281565b33600090815260208190526040902054600114610e9c5760405160e560020a62461bcd02815260206004820152601260248201527f4461692f6e6f742d617574686f72697a65640000000000000000000000000000604482015260640160405180910390fd5b600160a060020a0381166000908152602081905260408120555961012081016040526020815260e0602082015260e0600060408301376024356004353360003560e01c60e01b61012085a45050565b600160a060020a03821660009081526002602052819060409020541015610f5b5760405160e560020a62461bcd02815260206004820152601860248201527f4461692f696e73756666696369656e742d62616c616e63650000000000000000604482015260640160405180910390fd5b600160a060020a0382163314801590610f9f5750600160a060020a038216600090815260036020526000199060409020336000908152602091909152604090205414155b1561108157600160a060020a03821660009081526003602052819060409020336000908152602091909152604090205410156110245760405160e560020a62461bcd02815260206004820152601a60248201527f4461692f696e73756666696369656e742d616c6c6f77616e6365000000000000604482015260640160405180910390fd5b600160a060020a038216600090815260036020526110579060409020336000908152602091909152604090205482611186565b600160a060020a038316600090815260036020526040902033600090815260209190915260409020555b600160a060020a038216600090815260026020526110a490604090205482611186565b600160a060020a0383166000908152600260205260409020556001546110ca9082611186565b6001556000600160a060020a0383167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405190815260200160405180910390a35050565b600061111d338484610559565b9392505050565b61112f338383610559565b505050565b61113f838383610559565b50505050565b600060205280600052604060002054905081565b6003602052816000526040600020602052806000526040600020549150829050565b61112f823383610559565b8082038281111561054d57600080fd5b8082018281101561054d57600080fd00a165627a7a72305820814685d838c3da258269bf19df60b0cd953fa297a551cd75937017bded09f05a0029"

// DeployDai deploys a new Ethereum contract, binding an instance of Dai to it.
func DeployDai(auth *bind.TransactOpts, backend bind.ContractBackend, chainId_ *big.Int) (common.Address, *types.Transaction, *Dai, error) {
	parsed, err := abi.JSON(strings.NewReader(DaiABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DaiBin), backend, chainId_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Dai{DaiCaller: DaiCaller{contract: contract}, DaiTransactor: DaiTransactor{contract: contract}, DaiFilterer: DaiFilterer{contract: contract}}, nil
}

// Dai is an auto generated Go binding around an Ethereum contract.
type Dai struct {
	DaiCaller     // Read-only binding to the contract
	DaiTransactor // Write-only binding to the contract
	DaiFilterer   // Log filterer for contract events
}

// DaiCaller is an auto generated read-only Go binding around an Ethereum contract.
type DaiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DaiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DaiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DaiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DaiSession struct {
	Contract     *Dai              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DaiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DaiCallerSession struct {
	Contract *DaiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DaiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DaiTransactorSession struct {
	Contract     *DaiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DaiRaw is an auto generated low-level Go binding around an Ethereum contract.
type DaiRaw struct {
	Contract *Dai // Generic contract binding to access the raw methods on
}

// DaiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DaiCallerRaw struct {
	Contract *DaiCaller // Generic read-only contract binding to access the raw methods on
}

// DaiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DaiTransactorRaw struct {
	Contract *DaiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDai creates a new instance of Dai, bound to a specific deployed contract.
func NewDai(address common.Address, backend bind.ContractBackend) (*Dai, error) {
	contract, err := bindDai(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dai{DaiCaller: DaiCaller{contract: contract}, DaiTransactor: DaiTransactor{contract: contract}, DaiFilterer: DaiFilterer{contract: contract}}, nil
}

// NewDaiCaller creates a new read-only instance of Dai, bound to a specific deployed contract.
func NewDaiCaller(address common.Address, caller bind.ContractCaller) (*DaiCaller, error) {
	contract, err := bindDai(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DaiCaller{contract: contract}, nil
}

// NewDaiTransactor creates a new write-only instance of Dai, bound to a specific deployed contract.
func NewDaiTransactor(address common.Address, transactor bind.ContractTransactor) (*DaiTransactor, error) {
	contract, err := bindDai(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DaiTransactor{contract: contract}, nil
}

// NewDaiFilterer creates a new log filterer instance of Dai, bound to a specific deployed contract.
func NewDaiFilterer(address common.Address, filterer bind.ContractFilterer) (*DaiFilterer, error) {
	contract, err := bindDai(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DaiFilterer{contract: contract}, nil
}

// bindDai binds a generic wrapper to an already deployed contract.
func bindDai(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DaiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dai *DaiRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dai.Contract.DaiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dai *DaiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dai.Contract.DaiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dai *DaiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dai.Contract.DaiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dai *DaiCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dai.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dai *DaiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dai.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dai *DaiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dai.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Dai *DaiCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Dai.contract.Call(opts, out, "DOMAIN_SEPARATOR")
	return *ret0, err
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Dai *DaiSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Dai.Contract.DOMAINSEPARATOR(&_Dai.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Dai *DaiCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Dai.Contract.DOMAINSEPARATOR(&_Dai.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Dai *DaiCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Dai.contract.Call(opts, out, "PERMIT_TYPEHASH")
	return *ret0, err
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Dai *DaiSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Dai.Contract.PERMITTYPEHASH(&_Dai.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Dai *DaiCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Dai.Contract.PERMITTYPEHASH(&_Dai.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Dai *DaiCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dai.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Dai *DaiSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Dai.Contract.Allowance(&_Dai.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Dai *DaiCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Dai.Contract.Allowance(&_Dai.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Dai *DaiCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dai.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Dai *DaiSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Dai.Contract.BalanceOf(&_Dai.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Dai *DaiCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Dai.Contract.BalanceOf(&_Dai.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Dai *DaiCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Dai.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Dai *DaiSession) Decimals() (uint8, error) {
	return _Dai.Contract.Decimals(&_Dai.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Dai *DaiCallerSession) Decimals() (uint8, error) {
	return _Dai.Contract.Decimals(&_Dai.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dai *DaiCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Dai.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dai *DaiSession) Name() (string, error) {
	return _Dai.Contract.Name(&_Dai.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dai *DaiCallerSession) Name() (string, error) {
	return _Dai.Contract.Name(&_Dai.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Dai *DaiCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dai.contract.Call(opts, out, "nonces", arg0)
	return *ret0, err
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Dai *DaiSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Dai.Contract.Nonces(&_Dai.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Dai *DaiCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Dai.Contract.Nonces(&_Dai.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dai *DaiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Dai.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dai *DaiSession) Symbol() (string, error) {
	return _Dai.Contract.Symbol(&_Dai.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dai *DaiCallerSession) Symbol() (string, error) {
	return _Dai.Contract.Symbol(&_Dai.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dai *DaiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dai.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dai *DaiSession) TotalSupply() (*big.Int, error) {
	return _Dai.Contract.TotalSupply(&_Dai.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dai *DaiCallerSession) TotalSupply() (*big.Int, error) {
	return _Dai.Contract.TotalSupply(&_Dai.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Dai *DaiCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Dai.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Dai *DaiSession) Version() (string, error) {
	return _Dai.Contract.Version(&_Dai.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Dai *DaiCallerSession) Version() (string, error) {
	return _Dai.Contract.Version(&_Dai.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Dai *DaiCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dai.contract.Call(opts, out, "wards", arg0)
	return *ret0, err
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Dai *DaiSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Dai.Contract.Wards(&_Dai.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Dai *DaiCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Dai.Contract.Wards(&_Dai.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address usr, uint256 wad) returns(bool)
func (_Dai *DaiTransactor) Approve(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Dai.contract.Transact(opts, "approve", usr, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address usr, uint256 wad) returns(bool)
func (_Dai *DaiSession) Approve(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Dai.Contract.Approve(&_Dai.TransactOpts, usr, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address usr, uint256 wad) returns(bool)
func (_Dai *DaiTransactorSession) Approve(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Dai.Contract.Approve(&_Dai.TransactOpts, usr, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address usr, uint256 wad) returns()
func (_Dai *DaiTransactor) Burn(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Dai.contract.Transact(opts, "burn", usr, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address usr, uint256 wad) returns()
func (_Dai *DaiSession) Burn(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Dai.Contract.Burn(&_Dai.TransactOpts, usr, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address usr, uint256 wad) returns()
func (_Dai *DaiTransactorSession) Burn(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Dai.Contract.Burn(&_Dai.TransactOpts, usr, wad)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_Dai *DaiTransactor) Deny(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _Dai.contract.Transact(opts, "deny", guy)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_Dai *DaiSession) Deny(guy common.Address) (*types.Transaction, error) {
	return _Dai.Contract.Deny(&_Dai.TransactOpts, guy)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_Dai *DaiTransactorSession) Deny(guy common.Address) (*types.Transaction, error) {
	return _Dai.Contract.Deny(&_Dai.TransactOpts, guy)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address usr, uint256 wad) returns()
func (_Dai *DaiTransactor) Mint(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Dai.contract.Transact(opts, "mint", usr, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address usr, uint256 wad) returns()
func (_Dai *DaiSession) Mint(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Dai.Contract.Mint(&_Dai.TransactOpts, usr, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address usr, uint256 wad) returns()
func (_Dai *DaiTransactorSession) Mint(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Dai.Contract.Mint(&_Dai.TransactOpts, usr, wad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 wad) returns()
func (_Dai *DaiTransactor) Move(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Dai.contract.Transact(opts, "move", src, dst, wad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 wad) returns()
func (_Dai *DaiSession) Move(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Dai.Contract.Move(&_Dai.TransactOpts, src, dst, wad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 wad) returns()
func (_Dai *DaiTransactorSession) Move(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Dai.Contract.Move(&_Dai.TransactOpts, src, dst, wad)
}

// Permit is a paid mutator transaction binding the contract method 0x8fcbaf0c.
//
// Solidity: function permit(address holder, address spender, uint256 nonce, uint256 expiry, bool allowed, uint8 v, bytes32 r, bytes32 s) returns()
func (_Dai *DaiTransactor) Permit(opts *bind.TransactOpts, holder common.Address, spender common.Address, nonce *big.Int, expiry *big.Int, allowed bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Dai.contract.Transact(opts, "permit", holder, spender, nonce, expiry, allowed, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0x8fcbaf0c.
//
// Solidity: function permit(address holder, address spender, uint256 nonce, uint256 expiry, bool allowed, uint8 v, bytes32 r, bytes32 s) returns()
func (_Dai *DaiSession) Permit(holder common.Address, spender common.Address, nonce *big.Int, expiry *big.Int, allowed bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Dai.Contract.Permit(&_Dai.TransactOpts, holder, spender, nonce, expiry, allowed, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0x8fcbaf0c.
//
// Solidity: function permit(address holder, address spender, uint256 nonce, uint256 expiry, bool allowed, uint8 v, bytes32 r, bytes32 s) returns()
func (_Dai *DaiTransactorSession) Permit(holder common.Address, spender common.Address, nonce *big.Int, expiry *big.Int, allowed bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Dai.Contract.Permit(&_Dai.TransactOpts, holder, spender, nonce, expiry, allowed, v, r, s)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address usr, uint256 wad) returns()
func (_Dai *DaiTransactor) Pull(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Dai.contract.Transact(opts, "pull", usr, wad)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address usr, uint256 wad) returns()
func (_Dai *DaiSession) Pull(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Dai.Contract.Pull(&_Dai.TransactOpts, usr, wad)
}

// Pull is a paid mutator transaction binding the contract method 0xf2d5d56b.
//
// Solidity: function pull(address usr, uint256 wad) returns()
func (_Dai *DaiTransactorSession) Pull(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Dai.Contract.Pull(&_Dai.TransactOpts, usr, wad)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(address usr, uint256 wad) returns()
func (_Dai *DaiTransactor) Push(opts *bind.TransactOpts, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Dai.contract.Transact(opts, "push", usr, wad)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(address usr, uint256 wad) returns()
func (_Dai *DaiSession) Push(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Dai.Contract.Push(&_Dai.TransactOpts, usr, wad)
}

// Push is a paid mutator transaction binding the contract method 0xb753a98c.
//
// Solidity: function push(address usr, uint256 wad) returns()
func (_Dai *DaiTransactorSession) Push(usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Dai.Contract.Push(&_Dai.TransactOpts, usr, wad)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_Dai *DaiTransactor) Rely(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _Dai.contract.Transact(opts, "rely", guy)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_Dai *DaiSession) Rely(guy common.Address) (*types.Transaction, error) {
	return _Dai.Contract.Rely(&_Dai.TransactOpts, guy)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_Dai *DaiTransactorSession) Rely(guy common.Address) (*types.Transaction, error) {
	return _Dai.Contract.Rely(&_Dai.TransactOpts, guy)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Dai *DaiTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Dai.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Dai *DaiSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Dai.Contract.Transfer(&_Dai.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_Dai *DaiTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Dai.Contract.Transfer(&_Dai.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Dai *DaiTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Dai.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Dai *DaiSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Dai.Contract.TransferFrom(&_Dai.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_Dai *DaiTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Dai.Contract.TransferFrom(&_Dai.TransactOpts, src, dst, wad)
}

// DaiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Dai contract.
type DaiApprovalIterator struct {
	Event *DaiApproval // Event containing the contract specifics and raw log

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
func (it *DaiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaiApproval)
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
		it.Event = new(DaiApproval)
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
func (it *DaiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaiApproval represents a Approval event raised by the Dai contract.
type DaiApproval struct {
	Src common.Address
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_Dai *DaiFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, guy []common.Address) (*DaiApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Dai.contract.FilterLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return &DaiApprovalIterator{contract: _Dai.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_Dai *DaiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DaiApproval, src []common.Address, guy []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _Dai.contract.WatchLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaiApproval)
				if err := _Dai.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_Dai *DaiFilterer) ParseApproval(log types.Log) (*DaiApproval, error) {
	event := new(DaiApproval)
	if err := _Dai.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DaiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Dai contract.
type DaiTransferIterator struct {
	Event *DaiTransfer // Event containing the contract specifics and raw log

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
func (it *DaiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DaiTransfer)
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
		it.Event = new(DaiTransfer)
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
func (it *DaiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DaiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DaiTransfer represents a Transfer event raised by the Dai contract.
type DaiTransfer struct {
	Src common.Address
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_Dai *DaiFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*DaiTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Dai.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &DaiTransferIterator{contract: _Dai.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_Dai *DaiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DaiTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Dai.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DaiTransfer)
				if err := _Dai.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_Dai *DaiFilterer) ParseTransfer(log types.Log) (*DaiTransfer, error) {
	event := new(DaiTransfer)
	if err := _Dai.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LibNoteABI is the input ABI used to generate the binding from.
const LibNoteABI = "[{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"arg1\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"arg2\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"}]"

// LibNoteBin is the compiled bytecode used for deploying new contracts.
var LibNoteBin = "0x6080604052348015600f57600080fd5b50603580601d6000396000f3006080604052600080fd00a165627a7a72305820244b0a34a434591266b0795727f04e72dbad51025b190c5620231f99c7ea717c0029"

// DeployLibNote deploys a new Ethereum contract, binding an instance of LibNote to it.
func DeployLibNote(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LibNote, error) {
	parsed, err := abi.JSON(strings.NewReader(LibNoteABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LibNoteBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LibNote{LibNoteCaller: LibNoteCaller{contract: contract}, LibNoteTransactor: LibNoteTransactor{contract: contract}, LibNoteFilterer: LibNoteFilterer{contract: contract}}, nil
}

// LibNote is an auto generated Go binding around an Ethereum contract.
type LibNote struct {
	LibNoteCaller     // Read-only binding to the contract
	LibNoteTransactor // Write-only binding to the contract
	LibNoteFilterer   // Log filterer for contract events
}

// LibNoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type LibNoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibNoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LibNoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibNoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LibNoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LibNoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LibNoteSession struct {
	Contract     *LibNote          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LibNoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LibNoteCallerSession struct {
	Contract *LibNoteCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LibNoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LibNoteTransactorSession struct {
	Contract     *LibNoteTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LibNoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type LibNoteRaw struct {
	Contract *LibNote // Generic contract binding to access the raw methods on
}

// LibNoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LibNoteCallerRaw struct {
	Contract *LibNoteCaller // Generic read-only contract binding to access the raw methods on
}

// LibNoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LibNoteTransactorRaw struct {
	Contract *LibNoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLibNote creates a new instance of LibNote, bound to a specific deployed contract.
func NewLibNote(address common.Address, backend bind.ContractBackend) (*LibNote, error) {
	contract, err := bindLibNote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LibNote{LibNoteCaller: LibNoteCaller{contract: contract}, LibNoteTransactor: LibNoteTransactor{contract: contract}, LibNoteFilterer: LibNoteFilterer{contract: contract}}, nil
}

// NewLibNoteCaller creates a new read-only instance of LibNote, bound to a specific deployed contract.
func NewLibNoteCaller(address common.Address, caller bind.ContractCaller) (*LibNoteCaller, error) {
	contract, err := bindLibNote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LibNoteCaller{contract: contract}, nil
}

// NewLibNoteTransactor creates a new write-only instance of LibNote, bound to a specific deployed contract.
func NewLibNoteTransactor(address common.Address, transactor bind.ContractTransactor) (*LibNoteTransactor, error) {
	contract, err := bindLibNote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LibNoteTransactor{contract: contract}, nil
}

// NewLibNoteFilterer creates a new log filterer instance of LibNote, bound to a specific deployed contract.
func NewLibNoteFilterer(address common.Address, filterer bind.ContractFilterer) (*LibNoteFilterer, error) {
	contract, err := bindLibNote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LibNoteFilterer{contract: contract}, nil
}

// bindLibNote binds a generic wrapper to an already deployed contract.
func bindLibNote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LibNoteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibNote *LibNoteRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LibNote.Contract.LibNoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibNote *LibNoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibNote.Contract.LibNoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibNote *LibNoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibNote.Contract.LibNoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LibNote *LibNoteCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LibNote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LibNote *LibNoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LibNote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LibNote *LibNoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LibNote.Contract.contract.Transact(opts, method, params...)
}
