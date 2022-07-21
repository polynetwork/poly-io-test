// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eccmp_abi

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ChainSQL/go-chainsql-api/abigen/abi"
	"github.com/ChainSQL/go-chainsql-api/abigen/abi/bind"
	"github.com/ChainSQL/go-chainsql-api/common"
	"github.com/ChainSQL/go-chainsql-api/core"
	"github.com/ChainSQL/go-chainsql-api/data"
	"github.com/ChainSQL/go-chainsql-api/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = bind.Bind
	_ = common.Big1
)

// EccmpAbiMetaData contains all meta data concerning the EccmpAbi contract.
var EccmpAbiMetaData = &core.CtrMetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethCrossChainManagerAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_newChainId\",\"type\":\"uint64\"}],\"name\":\"changeManagerChainID\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEthCrossChainManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pauseEthCrossChainManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpauseEthCrossChainManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newEthCrossChainManagerAddr\",\"type\":\"address\"}],\"name\":\"upgradeEthCrossChainManager\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051611ad1380380611ad18339818101604052602081101561003357600080fd5b8101908080519060200190929190505050600061005461015360201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35060008060146101000a81548160ff02191690831515021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505061015b565b600033905090565b6119678061016a6000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c806387939a7f1161007157806387939a7f1461016d5780638da5cb5b146101b75780638f32d59b14610201578063a2681d2814610223578063ab59d32d1461025b578063f2fde38b146102b7576100b4565b80633b9a80b8146100b95780633f4ba83a146100db5780634390c707146100fd5780635c975abb1461011f578063715018a6146101415780638456cb591461014b575b600080fd5b6100c16102fb565b604051808215151515815260200191505060405180910390f35b6100e361055a565b604051808215151515815260200191505060405180910390f35b6101056105fb565b604051808215151515815260200191505060405180910390f35b610127610859565b604051808215151515815260200191505060405180910390f35b61014961086f565b005b6101536109a8565b604051808215151515815260200191505060405180910390f35b610175610a4a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6101bf610af6565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610209610b1f565b604051808215151515815260200191505060405180910390f35b6102596004803603602081101561023957600080fd5b81019080803567ffffffffffffffff169060200190929190505050610b7d565b005b61029d6004803603602081101561027157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f16565b604051808215151515815260200191505060405180910390f35b6102f9600480360360208110156102cd57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506113cb565b005b6000610305610b1f565b610377576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600060149054906101000a900460ff16156103fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506104296109a8565b61047e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603081526020018061187c6030913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16638456cb596040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156104c657600080fd5b505af11580156104da573d6000803e3d6000fd5b505050506040513d60208110156104f057600080fd5b8101908080519060200190929190505050610556576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b815260200180611908602b913960400191505060405180910390fd5b5090565b6000610564610b1f565b6105d6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6105de610859565b6105eb57600190506105f8565b6105f3611451565b600190505b90565b6000610605610b1f565b610677576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600060149054906101000a900460ff166106f9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f5061757361626c653a206e6f742070617573656400000000000000000000000081525060200191505060405180910390fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff16633f4ba83a6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561076857600080fd5b505af115801561077c573d6000803e3d6000fd5b505050506040513d602081101561079257600080fd5b81019080805190602001909291905050506107f8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602d8152602001806118db602d913960400191505060405180910390fd5b61080061055a565b610855576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603281526020018061184a6032913960400191505060405180910390fd5b5090565b60008060149054906101000a900460ff16905090565b610877610b1f565b6108e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60006109b2610b1f565b610a24576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b610a2c610859565b15610a3a5760019050610a47565b610a42611559565b600190505b90565b60008060149054906101000a900460ff1615610ace576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16610b61611663565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b610b85610b1f565b610bf7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600060149054906101000a900460ff16610c79576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f5061757361626c653a206e6f742070617573656400000000000000000000000081525060200191505060405180910390fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff16635c975abb6040518163ffffffff1660e01b815260040160206040518083038186803b158015610ce657600080fd5b505afa158015610cfa573d6000803e3d6000fd5b505050506040513d6020811015610d1057600080fd5b8101908080519060200190929190505050610dfe578073ffffffffffffffffffffffffffffffffffffffff16638456cb596040518163ffffffff1660e01b8152600401602060405180830381600087803b158015610d6d57600080fd5b505af1158015610d81573d6000803e3d6000fd5b505050506040513d6020811015610d9757600080fd5b8101908080519060200190929190505050610dfd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f8152602001806118ac602f913960400191505060405180910390fd5b5b8073ffffffffffffffffffffffffffffffffffffffff16636f31031d836040518263ffffffff1660e01b8152600401808267ffffffffffffffff1667ffffffffffffffff168152602001915050602060405180830381600087803b158015610e6557600080fd5b505af1158015610e79573d6000803e3d6000fd5b505050506040513d6020811015610e8f57600080fd5b8101908080519060200190929190505050610f12576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f73657420636861696e204944206661696c65642e20000000000000000000000081525060200191505060405180910390fd5b5050565b6000610f20610b1f565b610f92576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600060149054906101000a900460ff16611014576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f5061757361626c653a206e6f742070617573656400000000000000000000000081525060200191505060405180910390fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff16635c975abb6040518163ffffffff1660e01b815260040160206040518083038186803b15801561108157600080fd5b505afa158015611095573d6000803e3d6000fd5b505050506040513d60208110156110ab57600080fd5b8101908080519060200190929190505050611199578073ffffffffffffffffffffffffffffffffffffffff16638456cb596040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561110857600080fd5b505af115801561111c573d6000803e3d6000fd5b505050506040513d602081101561113257600080fd5b8101908080519060200190929190505050611198576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f8152602001806118ac602f913960400191505060405180910390fd5b5b8073ffffffffffffffffffffffffffffffffffffffff16637e724ff3846040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b15801561121857600080fd5b505af115801561122c573d6000803e3d6000fd5b505050506040513d602081101561124257600080fd5b81019080805190602001909291905050506112a8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001806117fb6029913960400191505060405180910390fd5b60008390508073ffffffffffffffffffffffffffffffffffffffff16638f32d59b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156112f357600080fd5b505afa158015611307573d6000803e3d6000fd5b505050506040513d602081101561131d57600080fd5b8101908080519060200190929190505050611383576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252604b8152602001806117b0604b913960600191505060405180910390fd5b83600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050919050565b6113d3610b1f565b611445576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b61144e8161166b565b50565b600060149054906101000a900460ff166114d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f5061757361626c653a206e6f742070617573656400000000000000000000000081525060200191505060405180910390fd5b60008060146101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa611516611663565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1565b600060149054906101000a900460ff16156115dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b6001600060146101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258611620611663565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156116f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806118246026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505056fe45746843726f7373436861696e4d616e6167657250726f7879206973206e6f74206f776e6572206f66206e65772045746843726f7373436861696e4d616e6167657220636f6e747261637445746843726f7373436861696e4d616e616765722075706772616465546f4e6577206661696c6564214f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373756e70617573652045746843726f7373436861696e4d616e6167657250726f787920636f6e7472616374206661696c65642170617573652045746843726f7373436861696e4d616e6167657250726f787920636f6e7472616374206661696c6564215061757365206f6c642045746843726f7373436861696e4d616e6167657220636f6e7472616374206661696c656421756e70617573652045746843726f7373436861696e4d616e6167657220636f6e7472616374206661696c65642170617573652045746843726f7373436861696e4d616e6167657220636f6e7472616374206661696c656421a265627a7a72315820aa267dcfec48b655967c266b94d6d3439b90dac9b6d4cf4debfbad1510c641b264736f6c63430005100032",
}

// EccmpAbiABI is the input ABI used to generate the binding from.
// Deprecated: Use EccmpAbiMetaData.ABI instead.
var EccmpAbiABI = EccmpAbiMetaData.ABI

// EccmpAbiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EccmpAbiMetaData.Bin instead.
var EccmpAbiBin = EccmpAbiMetaData.Bin

// DeployEccmpAbi deploys a new ChainSQL contract, binding an instance of EccmpAbi to it.
func DeployEccmpAbi(chainsql *core.Chainsql, auth *core.TransactOpts, _ethCrossChainManagerAddr common.Address) (*core.DeployTxRet, *EccmpAbi, error) {
	parsed, err := EccmpAbiMetaData.GetAbi()
	if err != nil {
		return &core.DeployTxRet{}, nil, err
	}
	if parsed == nil {
		return &core.DeployTxRet{}, nil, errors.New("GetABI returned nil")
	}

	deployRet, contract, err := core.DeployContract(chainsql, auth, *parsed, common.FromHex(EccmpAbiBin), _ethCrossChainManagerAddr)
	if err != nil {
		return &core.DeployTxRet{}, nil, err
	}
	return deployRet, &EccmpAbi{EccmpAbiCaller: EccmpAbiCaller{contract: contract}, EccmpAbiTransactor: EccmpAbiTransactor{contract: contract}, EccmpAbiFilterer: EccmpAbiFilterer{contract: contract}}, nil
}

// EccmpAbi is an auto generated Go binding around an ChainSQL contract.
type EccmpAbi struct {
	EccmpAbiCaller     // Read-only binding to the contract
	EccmpAbiTransactor // Write-only binding to the contract
	EccmpAbiFilterer   // Log filterer for contract events
}

// EccmpAbiCaller is an auto generated read-only Go binding around an ChainSQL contract.
type EccmpAbiCaller struct {
	contract *core.BoundContract // Generic contract wrapper for the low level calls
}

// EccmpAbiTransactor is an auto generated write-only Go binding around an ChainSQL contract.
type EccmpAbiTransactor struct {
	contract *core.BoundContract // Generic contract wrapper for the low level calls
}

// EccmpAbiFilterer is an auto generated log filtering Go binding around an ChainSQL contract events.
type EccmpAbiFilterer struct {
	contract *core.BoundContract // Generic contract wrapper for the low level calls
}

// EccmpAbiSession is an auto generated Go binding around an ChainSQL contract,
// with pre-set call and transact options.
type EccmpAbiSession struct {
	Contract     *EccmpAbi         // Generic contract binding to set the session for
	CallOpts     core.CallOpts     // Call options to use throughout this session
	TransactOpts core.TransactOpts // Transaction auth options to use throughout this session
}

// EccmpAbiCallerSession is an auto generated read-only Go binding around an ChainSQL contract,
// with pre-set call options.
type EccmpAbiCallerSession struct {
	Contract *EccmpAbiCaller // Generic contract caller binding to set the session for
	CallOpts core.CallOpts   // Call options to use throughout this session
}

// EccmpAbiTransactorSession is an auto generated write-only Go binding around an ChainSQL contract,
// with pre-set transact options.
type EccmpAbiTransactorSession struct {
	Contract     *EccmpAbiTransactor // Generic contract transactor binding to set the session for
	TransactOpts core.TransactOpts   // Transaction auth options to use throughout this session
}

// EccmpAbiRaw is an auto generated low-level Go binding around an ChainSQL contract.
type EccmpAbiRaw struct {
	Contract *EccmpAbi // Generic contract binding to access the raw methods on
}

// EccmpAbiCallerRaw is an auto generated low-level read-only Go binding around an ChainSQL contract.
type EccmpAbiCallerRaw struct {
	Contract *EccmpAbiCaller // Generic read-only contract binding to access the raw methods on
}

// EccmpAbiTransactorRaw is an auto generated low-level write-only Go binding around an ChainSQL contract.
type EccmpAbiTransactorRaw struct {
	Contract *EccmpAbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEccmpAbi creates a new instance of EccmpAbi, bound to a specific deployed contract.
func NewEccmpAbi(chainsql *core.Chainsql, address string) (*EccmpAbi, error) {
	contract, err := bindEccmpAbi(chainsql, address)
	if err != nil {
		return nil, err
	}
	return &EccmpAbi{EccmpAbiCaller: EccmpAbiCaller{contract: contract}, EccmpAbiTransactor: EccmpAbiTransactor{contract: contract}, EccmpAbiFilterer: EccmpAbiFilterer{contract: contract}}, nil
}

// // NewEccmpAbiCaller creates a new read-only instance of EccmpAbi, bound to a specific deployed contract.
// func NewEccmpAbiCaller(address common.Address, caller bind.ContractCaller) (*EccmpAbiCaller, error) {
//   contract, err := bindEccmpAbi(address, caller, nil, nil)
//   if err != nil {
//     return nil, err
//   }
//   return &EccmpAbiCaller{contract: contract}, nil
// }

// // NewEccmpAbiTransactor creates a new write-only instance of EccmpAbi, bound to a specific deployed contract.
// func NewEccmpAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*EccmpAbiTransactor, error) {
//   contract, err := bindEccmpAbi(address, nil, transactor, nil)
//   if err != nil {
//     return nil, err
//   }
//   return &EccmpAbiTransactor{contract: contract}, nil
// }

// // NewEccmpAbiFilterer creates a new log filterer instance of EccmpAbi, bound to a specific deployed contract.
// func NewEccmpAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*EccmpAbiFilterer, error) {
//   contract, err := bindEccmpAbi(address, nil, nil, filterer)
//   if err != nil {
//     return nil, err
//   }
//   return &EccmpAbiFilterer{contract: contract}, nil
// }

// bindEccmpAbi binds a generic wrapper to an already deployed contract.
func bindEccmpAbi(chainsql *core.Chainsql, address string) (*core.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EccmpAbiABI))
	if err != nil {
		return nil, err
	}
	return core.NewBoundContract(chainsql, address, parsed), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
// func (_EccmpAbi *EccmpAbiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
// 	return _EccmpAbi.Contract.EccmpAbiCaller.contract.Call(opts, result, method, params...)
// }

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
// func (_EccmpAbi *EccmpAbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
// 	return _EccmpAbi.Contract.EccmpAbiTransactor.contract.Transfer(opts)
// }

// Transact invokes the (paid) contract method with params as input values.
// func (_EccmpAbi *EccmpAbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
// 	return _EccmpAbi.Contract.EccmpAbiTransactor.contract.Transact(opts, method, params...)
// }

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
// func (_EccmpAbi *EccmpAbiCallerRaw) Call(opts *core.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
// 	return _EccmpAbi.Contract.contract.Call(opts, result, method, params...)
// }

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
// func (_EccmpAbi *EccmpAbiTransactorRaw) Transfer(opts *core.TransactOpts) (*types.Transaction, error) {
// 	return _EccmpAbi.Contract.contract.Transfer(opts)
// }

// Transact invokes the (paid) contract method with params as input values.
// func (_EccmpAbi *EccmpAbiTransactorRaw) Transact(opts *core.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
// 	return _EccmpAbi.Contract.contract.Transact(opts, method, params...)
// }

// GetEthCrossChainManager is a free data retrieval call binding the contract method 0x87939a7f.
//
// Solidity: function getEthCrossChainManager() view returns(address)
func (_EccmpAbi *EccmpAbiCaller) GetEthCrossChainManager(opts *core.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EccmpAbi.contract.Call(opts, &out, "getEthCrossChainManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEthCrossChainManager is a free data retrieval call binding the contract method 0x87939a7f.
//
// Solidity: function getEthCrossChainManager() view returns(address)
func (_EccmpAbi *EccmpAbiSession) GetEthCrossChainManager() (common.Address, error) {
	return _EccmpAbi.Contract.GetEthCrossChainManager(&_EccmpAbi.CallOpts)
}

// GetEthCrossChainManager is a free data retrieval call binding the contract method 0x87939a7f.
//
// Solidity: function getEthCrossChainManager() view returns(address)
func (_EccmpAbi *EccmpAbiCallerSession) GetEthCrossChainManager() (common.Address, error) {
	return _EccmpAbi.Contract.GetEthCrossChainManager(&_EccmpAbi.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_EccmpAbi *EccmpAbiCaller) IsOwner(opts *core.CallOpts) (bool, error) {
	var out []interface{}
	err := _EccmpAbi.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_EccmpAbi *EccmpAbiSession) IsOwner() (bool, error) {
	return _EccmpAbi.Contract.IsOwner(&_EccmpAbi.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_EccmpAbi *EccmpAbiCallerSession) IsOwner() (bool, error) {
	return _EccmpAbi.Contract.IsOwner(&_EccmpAbi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EccmpAbi *EccmpAbiCaller) Owner(opts *core.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EccmpAbi.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EccmpAbi *EccmpAbiSession) Owner() (common.Address, error) {
	return _EccmpAbi.Contract.Owner(&_EccmpAbi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EccmpAbi *EccmpAbiCallerSession) Owner() (common.Address, error) {
	return _EccmpAbi.Contract.Owner(&_EccmpAbi.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EccmpAbi *EccmpAbiCaller) Paused(opts *core.CallOpts) (bool, error) {
	var out []interface{}
	err := _EccmpAbi.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EccmpAbi *EccmpAbiSession) Paused() (bool, error) {
	return _EccmpAbi.Contract.Paused(&_EccmpAbi.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EccmpAbi *EccmpAbiCallerSession) Paused() (bool, error) {
	return _EccmpAbi.Contract.Paused(&_EccmpAbi.CallOpts)
}

// ChangeManagerChainID is a paid mutator transaction binding the contract method 0xa2681d28.
//
// Solidity: function changeManagerChainID(uint64 _newChainId) returns()
func (_EccmpAbi *EccmpAbiTransactor) ChangeManagerChainID(opts *core.TransactOpts, _newChainId uint64) (*common.TxResult, error) {
	return _EccmpAbi.contract.Transact(opts, "changeManagerChainID", _newChainId)
}

// ChangeManagerChainID is a paid mutator transaction binding the contract method 0xa2681d28.
//
// Solidity: function changeManagerChainID(uint64 _newChainId) returns()
func (_EccmpAbi *EccmpAbiSession) ChangeManagerChainID(_newChainId uint64) (*common.TxResult, error) {
	return _EccmpAbi.Contract.ChangeManagerChainID(&_EccmpAbi.TransactOpts, _newChainId)
}

// ChangeManagerChainID is a paid mutator transaction binding the contract method 0xa2681d28.
//
// Solidity: function changeManagerChainID(uint64 _newChainId) returns()
func (_EccmpAbi *EccmpAbiTransactorSession) ChangeManagerChainID(_newChainId uint64) (*common.TxResult, error) {
	return _EccmpAbi.Contract.ChangeManagerChainID(&_EccmpAbi.TransactOpts, _newChainId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_EccmpAbi *EccmpAbiTransactor) Pause(opts *core.TransactOpts) (*common.TxResult, error) {
	return _EccmpAbi.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_EccmpAbi *EccmpAbiSession) Pause() (*common.TxResult, error) {
	return _EccmpAbi.Contract.Pause(&_EccmpAbi.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_EccmpAbi *EccmpAbiTransactorSession) Pause() (*common.TxResult, error) {
	return _EccmpAbi.Contract.Pause(&_EccmpAbi.TransactOpts)
}

// PauseEthCrossChainManager is a paid mutator transaction binding the contract method 0x3b9a80b8.
//
// Solidity: function pauseEthCrossChainManager() returns(bool)
func (_EccmpAbi *EccmpAbiTransactor) PauseEthCrossChainManager(opts *core.TransactOpts) (*common.TxResult, error) {
	return _EccmpAbi.contract.Transact(opts, "pauseEthCrossChainManager")
}

// PauseEthCrossChainManager is a paid mutator transaction binding the contract method 0x3b9a80b8.
//
// Solidity: function pauseEthCrossChainManager() returns(bool)
func (_EccmpAbi *EccmpAbiSession) PauseEthCrossChainManager() (*common.TxResult, error) {
	return _EccmpAbi.Contract.PauseEthCrossChainManager(&_EccmpAbi.TransactOpts)
}

// PauseEthCrossChainManager is a paid mutator transaction binding the contract method 0x3b9a80b8.
//
// Solidity: function pauseEthCrossChainManager() returns(bool)
func (_EccmpAbi *EccmpAbiTransactorSession) PauseEthCrossChainManager() (*common.TxResult, error) {
	return _EccmpAbi.Contract.PauseEthCrossChainManager(&_EccmpAbi.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EccmpAbi *EccmpAbiTransactor) RenounceOwnership(opts *core.TransactOpts) (*common.TxResult, error) {
	return _EccmpAbi.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EccmpAbi *EccmpAbiSession) RenounceOwnership() (*common.TxResult, error) {
	return _EccmpAbi.Contract.RenounceOwnership(&_EccmpAbi.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EccmpAbi *EccmpAbiTransactorSession) RenounceOwnership() (*common.TxResult, error) {
	return _EccmpAbi.Contract.RenounceOwnership(&_EccmpAbi.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EccmpAbi *EccmpAbiTransactor) TransferOwnership(opts *core.TransactOpts, newOwner common.Address) (*common.TxResult, error) {
	return _EccmpAbi.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EccmpAbi *EccmpAbiSession) TransferOwnership(newOwner common.Address) (*common.TxResult, error) {
	return _EccmpAbi.Contract.TransferOwnership(&_EccmpAbi.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EccmpAbi *EccmpAbiTransactorSession) TransferOwnership(newOwner common.Address) (*common.TxResult, error) {
	return _EccmpAbi.Contract.TransferOwnership(&_EccmpAbi.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_EccmpAbi *EccmpAbiTransactor) Unpause(opts *core.TransactOpts) (*common.TxResult, error) {
	return _EccmpAbi.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_EccmpAbi *EccmpAbiSession) Unpause() (*common.TxResult, error) {
	return _EccmpAbi.Contract.Unpause(&_EccmpAbi.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_EccmpAbi *EccmpAbiTransactorSession) Unpause() (*common.TxResult, error) {
	return _EccmpAbi.Contract.Unpause(&_EccmpAbi.TransactOpts)
}

// UnpauseEthCrossChainManager is a paid mutator transaction binding the contract method 0x4390c707.
//
// Solidity: function unpauseEthCrossChainManager() returns(bool)
func (_EccmpAbi *EccmpAbiTransactor) UnpauseEthCrossChainManager(opts *core.TransactOpts) (*common.TxResult, error) {
	return _EccmpAbi.contract.Transact(opts, "unpauseEthCrossChainManager")
}

// UnpauseEthCrossChainManager is a paid mutator transaction binding the contract method 0x4390c707.
//
// Solidity: function unpauseEthCrossChainManager() returns(bool)
func (_EccmpAbi *EccmpAbiSession) UnpauseEthCrossChainManager() (*common.TxResult, error) {
	return _EccmpAbi.Contract.UnpauseEthCrossChainManager(&_EccmpAbi.TransactOpts)
}

// UnpauseEthCrossChainManager is a paid mutator transaction binding the contract method 0x4390c707.
//
// Solidity: function unpauseEthCrossChainManager() returns(bool)
func (_EccmpAbi *EccmpAbiTransactorSession) UnpauseEthCrossChainManager() (*common.TxResult, error) {
	return _EccmpAbi.Contract.UnpauseEthCrossChainManager(&_EccmpAbi.TransactOpts)
}

// UpgradeEthCrossChainManager is a paid mutator transaction binding the contract method 0xab59d32d.
//
// Solidity: function upgradeEthCrossChainManager(address _newEthCrossChainManagerAddr) returns(bool)
func (_EccmpAbi *EccmpAbiTransactor) UpgradeEthCrossChainManager(opts *core.TransactOpts, _newEthCrossChainManagerAddr common.Address) (*common.TxResult, error) {
	return _EccmpAbi.contract.Transact(opts, "upgradeEthCrossChainManager", _newEthCrossChainManagerAddr)
}

// UpgradeEthCrossChainManager is a paid mutator transaction binding the contract method 0xab59d32d.
//
// Solidity: function upgradeEthCrossChainManager(address _newEthCrossChainManagerAddr) returns(bool)
func (_EccmpAbi *EccmpAbiSession) UpgradeEthCrossChainManager(_newEthCrossChainManagerAddr common.Address) (*common.TxResult, error) {
	return _EccmpAbi.Contract.UpgradeEthCrossChainManager(&_EccmpAbi.TransactOpts, _newEthCrossChainManagerAddr)
}

// UpgradeEthCrossChainManager is a paid mutator transaction binding the contract method 0xab59d32d.
//
// Solidity: function upgradeEthCrossChainManager(address _newEthCrossChainManagerAddr) returns(bool)
func (_EccmpAbi *EccmpAbiTransactorSession) UpgradeEthCrossChainManager(_newEthCrossChainManagerAddr common.Address) (*common.TxResult, error) {
	return _EccmpAbi.Contract.UpgradeEthCrossChainManager(&_EccmpAbi.TransactOpts, _newEthCrossChainManagerAddr)
}

// EccmpAbiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EccmpAbi contract.
type EccmpAbiOwnershipTransferredIterator struct {
	Event *EccmpAbiOwnershipTransferred // Event containing the contract specifics and raw log

	contract *core.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan data.Log      // Log channel receiving the found contract events
	sub  event.Subscription // Subscription for errors, completion and termination
	done bool               // Whether the subscription completed delivering logs
	fail error              // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EccmpAbiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EccmpAbiOwnershipTransferred)
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
		it.Event = new(EccmpAbiOwnershipTransferred)
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
func (it *EccmpAbiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EccmpAbiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EccmpAbiOwnershipTransferred represents a OwnershipTransferred event raised by the EccmpAbi contract.
type EccmpAbiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           data.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
// func (_EccmpAbi *EccmpAbiFilterer) FilterOwnershipTransferred(opts *core.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EccmpAbiOwnershipTransferredIterator, error) {
//
// 	var previousOwnerRule []interface{}
// 	for _, previousOwnerItem := range previousOwner {
// 		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
// 	}
// 	var newOwnerRule []interface{}
// 	for _, newOwnerItem := range newOwner {
// 		newOwnerRule = append(newOwnerRule, newOwnerItem)
// 	}

// 	logs, sub, err := _EccmpAbi.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &EccmpAbiOwnershipTransferredIterator{contract: _EccmpAbi.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
// }

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EccmpAbi *EccmpAbiFilterer) WatchOwnershipTransferred(opts *core.WatchOpts, sink chan<- *EccmpAbiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	sub, err := _EccmpAbi.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.UnSubscribe()
		for {
			select {
			case log := <-sub.EventMsgCh:
				// New log arrived, parse the event and forward to the user
				event := new(EccmpAbiOwnershipTransferred)
				if err := _EccmpAbi.contract.UnpackLog(event, "OwnershipTransferred", *log); err != nil {
					return err
				}
				event.Raw = *log

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

func (_EccmpAbi *EccmpAbiFilterer) GetOwnershipTransferredPastEvent(txHash string, ContractLogs string) ([]*EccmpAbiOwnershipTransferred, error) {
	var logRaws []*data.Log
	var err error
	if ContractLogs != "" {
		logRaws, err = _EccmpAbi.contract.GetPastEventByCtrLog(ContractLogs)
	} else if txHash != "" {
		logRaws, err = _EccmpAbi.contract.GetPastEventByTxHash(txHash)
	} else {
		return nil, errors.New("both txHash or ContractLogs is not provided for param")
	}

	if err != nil {
		return nil, err
	}
	var events []*EccmpAbiOwnershipTransferred
	for _, logRaw := range logRaws {
		event, err := _EccmpAbi.ParseOwnershipTransferred(*logRaw)
		if err != nil && err.Error() != "event signature mismatch" {
			return nil, err
		}
		if event != nil {
			events = append(events, event)
		}
	}
	return events, nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EccmpAbi *EccmpAbiFilterer) ParseOwnershipTransferred(log data.Log) (*EccmpAbiOwnershipTransferred, error) {
	event := new(EccmpAbiOwnershipTransferred)
	if err := _EccmpAbi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EccmpAbiPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the EccmpAbi contract.
type EccmpAbiPausedIterator struct {
	Event *EccmpAbiPaused // Event containing the contract specifics and raw log

	contract *core.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan data.Log      // Log channel receiving the found contract events
	sub  event.Subscription // Subscription for errors, completion and termination
	done bool               // Whether the subscription completed delivering logs
	fail error              // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EccmpAbiPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EccmpAbiPaused)
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
		it.Event = new(EccmpAbiPaused)
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
func (it *EccmpAbiPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EccmpAbiPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EccmpAbiPaused represents a Paused event raised by the EccmpAbi contract.
type EccmpAbiPaused struct {
	Account common.Address
	Raw     data.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
// func (_EccmpAbi *EccmpAbiFilterer) FilterPaused(opts *core.FilterOpts) (*EccmpAbiPausedIterator, error) {
//
//

// 	logs, sub, err := _EccmpAbi.contract.FilterLogs(opts, "Paused")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &EccmpAbiPausedIterator{contract: _EccmpAbi.contract, event: "Paused", logs: logs, sub: sub}, nil
// }

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EccmpAbi *EccmpAbiFilterer) WatchPaused(opts *core.WatchOpts, sink chan<- *EccmpAbiPaused) (event.Subscription, error) {

	sub, err := _EccmpAbi.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.UnSubscribe()
		for {
			select {
			case log := <-sub.EventMsgCh:
				// New log arrived, parse the event and forward to the user
				event := new(EccmpAbiPaused)
				if err := _EccmpAbi.contract.UnpackLog(event, "Paused", *log); err != nil {
					return err
				}
				event.Raw = *log

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

func (_EccmpAbi *EccmpAbiFilterer) GetPausedPastEvent(txHash string, ContractLogs string) ([]*EccmpAbiPaused, error) {
	var logRaws []*data.Log
	var err error
	if ContractLogs != "" {
		logRaws, err = _EccmpAbi.contract.GetPastEventByCtrLog(ContractLogs)
	} else if txHash != "" {
		logRaws, err = _EccmpAbi.contract.GetPastEventByTxHash(txHash)
	} else {
		return nil, errors.New("both txHash or ContractLogs is not provided for param")
	}

	if err != nil {
		return nil, err
	}
	var events []*EccmpAbiPaused
	for _, logRaw := range logRaws {
		event, err := _EccmpAbi.ParsePaused(*logRaw)
		if err != nil && err.Error() != "event signature mismatch" {
			return nil, err
		}
		if event != nil {
			events = append(events, event)
		}
	}
	return events, nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EccmpAbi *EccmpAbiFilterer) ParsePaused(log data.Log) (*EccmpAbiPaused, error) {
	event := new(EccmpAbiPaused)
	if err := _EccmpAbi.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EccmpAbiUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the EccmpAbi contract.
type EccmpAbiUnpausedIterator struct {
	Event *EccmpAbiUnpaused // Event containing the contract specifics and raw log

	contract *core.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan data.Log      // Log channel receiving the found contract events
	sub  event.Subscription // Subscription for errors, completion and termination
	done bool               // Whether the subscription completed delivering logs
	fail error              // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EccmpAbiUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EccmpAbiUnpaused)
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
		it.Event = new(EccmpAbiUnpaused)
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
func (it *EccmpAbiUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EccmpAbiUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EccmpAbiUnpaused represents a Unpaused event raised by the EccmpAbi contract.
type EccmpAbiUnpaused struct {
	Account common.Address
	Raw     data.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
// func (_EccmpAbi *EccmpAbiFilterer) FilterUnpaused(opts *core.FilterOpts) (*EccmpAbiUnpausedIterator, error) {
//
//

// 	logs, sub, err := _EccmpAbi.contract.FilterLogs(opts, "Unpaused")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &EccmpAbiUnpausedIterator{contract: _EccmpAbi.contract, event: "Unpaused", logs: logs, sub: sub}, nil
// }

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EccmpAbi *EccmpAbiFilterer) WatchUnpaused(opts *core.WatchOpts, sink chan<- *EccmpAbiUnpaused) (event.Subscription, error) {

	sub, err := _EccmpAbi.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.UnSubscribe()
		for {
			select {
			case log := <-sub.EventMsgCh:
				// New log arrived, parse the event and forward to the user
				event := new(EccmpAbiUnpaused)
				if err := _EccmpAbi.contract.UnpackLog(event, "Unpaused", *log); err != nil {
					return err
				}
				event.Raw = *log

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

func (_EccmpAbi *EccmpAbiFilterer) GetUnpausedPastEvent(txHash string, ContractLogs string) ([]*EccmpAbiUnpaused, error) {
	var logRaws []*data.Log
	var err error
	if ContractLogs != "" {
		logRaws, err = _EccmpAbi.contract.GetPastEventByCtrLog(ContractLogs)
	} else if txHash != "" {
		logRaws, err = _EccmpAbi.contract.GetPastEventByTxHash(txHash)
	} else {
		return nil, errors.New("both txHash or ContractLogs is not provided for param")
	}

	if err != nil {
		return nil, err
	}
	var events []*EccmpAbiUnpaused
	for _, logRaw := range logRaws {
		event, err := _EccmpAbi.ParseUnpaused(*logRaw)
		if err != nil && err.Error() != "event signature mismatch" {
			return nil, err
		}
		if event != nil {
			events = append(events, event)
		}
	}
	return events, nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EccmpAbi *EccmpAbiFilterer) ParseUnpaused(log data.Log) (*EccmpAbiUnpaused, error) {
	event := new(EccmpAbiUnpaused)
	if err := _EccmpAbi.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
