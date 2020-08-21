// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package asset_abi

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/event"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AssetABI is the input ABI used to generate the binding from.
const AssetABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"string\"}],\"name\":\"select\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from_account\",\"type\":\"string\"},{\"name\":\"to_account\",\"type\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"string\"},{\"name\":\"asset_value\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from_account\",\"type\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"increase\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from_account\",\"type\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decrease\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"ret\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"account\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"asset_value\",\"type\":\"uint256\"}],\"name\":\"RegisterEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"ret\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"from_account\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"to_account\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"ret\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"from_account\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IncreaseEvent\",\"type\":\"event\"}]"

// AssetFuncSigs maps the 4-byte function signature to its string representation.
var AssetFuncSigs = map[string]string{
	"f8c6f6a6": "decrease(string,uint256)",
	"c68e3b0b": "increase(string,uint256)",
	"b433c7ca": "register(string,uint256)",
	"5b325d78": "select(string)",
	"612d2bff": "transfer(string,string,uint256)",
}

// AssetBin is the compiled bytecode used for deploying new contracts.
var AssetBin = "0x608060405234801561001057600080fd5b50610022640100000000610027810204565b61014b565b604080517fc92a7801000000000000000000000000000000000000000000000000000000008152606060048201526007606482018190527f745f617373657400000000000000000000000000000000000000000000000000608483015260a0602483015260a48201527f6163636f756e740000000000000000000000000000000000000000000000000060c482015260e06044820152600b60e48201527f61737365745f76616c7565000000000000000000000000000000000000000000610104820152905161100191829163c92a780191610124808201926020929091908290030181600087803b15801561011c57600080fd5b505af1158015610130573d6000803e3d6000fd5b505050506040513d602081101561014657600080fd5b505050565b61204d806200015b6000396000f3fe60806040526004361061004d5760e060020a60003504635b325d788114610052578063612d2bff1461011e578063b433c7ca1461026c578063c68e3b0b14610321578063f8c6f6a6146103d6575b600080fd5b34801561005e57600080fd5b506101056004803603602081101561007557600080fd5b81019060208101813564010000000081111561009057600080fd5b8201836020820111156100a257600080fd5b803590602001918460018302840111640100000000831117156100c457600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061048b945050505050565b6040805192835260208301919091528051918290030190f35b34801561012a57600080fd5b5061025a6004803603606081101561014157600080fd5b81019060208101813564010000000081111561015c57600080fd5b82018360208201111561016e57600080fd5b8035906020019184600183028401116401000000008311171561019057600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156101e357600080fd5b8201836020820111156101f557600080fd5b8035906020019184600183028401116401000000008311171561021757600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506107bb915050565b60408051918252519081900360200190f35b34801561027857600080fd5b5061025a6004803603604081101561028f57600080fd5b8101906020810181356401000000008111156102aa57600080fd5b8201836020820111156102bc57600080fd5b803590602001918460018302840111640100000000831117156102de57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506112fe915050565b34801561032d57600080fd5b5061025a6004803603604081101561034457600080fd5b81019060208101813564010000000081111561035f57600080fd5b82018360208201111561037157600080fd5b8035906020019184600183028401116401000000008311171561039357600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506116d5915050565b3480156103e257600080fd5b5061025a600480360360408110156103f957600080fd5b81019060208101813564010000000081111561041457600080fd5b82018360208201111561042657600080fd5b8035906020019184600183028401116401000000008311171561044857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250611c33915050565b6000806000610498611ee3565b9050600081600160a060020a031663d8ac59578684600160a060020a031663c74f8caf6040518163ffffffff1660e060020a02815260040160206040518083038186803b1580156104e857600080fd5b505afa1580156104fc573d6000803e3d6000fd5b505050506040513d602081101561051257600080fd5b50516040805160e060020a63ffffffff8616028152600160a060020a038316602482015260048101918252835160448201528351829160640190602086019080838360005b8381101561056f578181015183820152602001610557565b50505050905090810190601f16801561059c5780820380516001836020036101000a031916815260200191505b50935050505060206040518083038186803b1580156105ba57600080fd5b505afa1580156105ce573d6000803e3d6000fd5b505050506040513d60208110156105e457600080fd5b5051604080517fd3e9af5a0000000000000000000000000000000000000000000000000000000081529051919250600091600160a060020a0384169163d3e9af5a916004808301926020929190829003018186803b15801561064557600080fd5b505afa158015610659573d6000803e3d6000fd5b505050506040513d602081101561066f57600080fd5b5051151561068657600019945092506107b6915050565b600082600160a060020a0316633dd2b61460006040518263ffffffff1660e060020a0281526004018082815260200191505060206040518083038186803b1580156106d057600080fd5b505afa1580156106e4573d6000803e3d6000fd5b505050506040513d60208110156106fa57600080fd5b5051604080517f4900862e000000000000000000000000000000000000000000000000000000008152602060048201819052600b6024830152600080516020611fa283398151915260448301529151929350600092600160a060020a03851692634900862e9260648082019391829003018186803b15801561077b57600080fd5b505afa15801561078f573d6000803e3d6000fd5b505050506040513d60208110156107a557600080fd5b505190965094506107b69350505050565b915091565b6000808080806107ca8861048b565b909350915082156108e457600019935060008051602061200283398151915284898989604051808581526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360005b8381101561083a578181015183820152602001610822565b50505050905090810190601f1680156108675780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b8381101561089a578181015183820152602001610882565b50505050905090810190601f1680156108c75780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a1839450505050506112f7565b6108ed8761048b565b9093509050821561095c57600119935060008051602061200283398151915284898989604051808581526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360008381101561083a578181015183820152602001610822565b858210156109c857600219935060008051602061200283398151915284898989604051808581526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360008381101561083a578181015183820152602001610822565b808682011015610a3657600319935060008051602061200283398151915284898989604051808581526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360008381101561083a578181015183820152602001610822565b6000610a40611ee3565b9050600081600160a060020a0316635887ab246040518163ffffffff1660e060020a02815260040160206040518083038186803b158015610a8057600080fd5b505afa158015610a94573d6000803e3d6000fd5b505050506040513d6020811015610aaa57600080fd5b50516040805160e260020a63068e472d0281526004810191825260076044820152600080516020611fc283398151915260648201526080602482019081528d5160848301528d51939450600160a060020a03851693631a391cb4938f9390928392909160a40190602086019080838360005b83811015610b34578181015183820152602001610b1c565b50505050905090810190601f168015610b615780820380516001836020036101000a031916815260200191505b509350505050600060405180830381600087803b158015610b8157600080fd5b505af1158015610b95573d6000803e3d6000fd5b50506040805160e360020a631bde84d30281528b8803602482015260048101829052600b6044820152600080516020611fa283398151915260648201529051600160a060020a038516935063def426989250608480830192600092919082900301818387803b158015610c0757600080fd5b505af1158015610c1b573d6000803e3d6000fd5b50505050600082600160a060020a031663664b37d68c8486600160a060020a031663c74f8caf6040518163ffffffff1660e060020a02815260040160206040518083038186803b158015610c6e57600080fd5b505afa158015610c82573d6000803e3d6000fd5b505050506040513d6020811015610c9857600080fd5b505160405160e060020a63ffffffff8616028152600160a060020a0380841660248301528216604482015260606004820190815284516064830152845190918291608490910190602087019080838360005b83811015610d02578181015183820152602001610cea565b50505050905090810190601f168015610d2f5780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b158015610d5057600080fd5b505af1158015610d64573d6000803e3d6000fd5b505050506040513d6020811015610d7a57600080fd5b5051905060018114610e98576004199650600080516020612002833981519152878c8c8c604051808581526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360005b83811015610deb578181015183820152602001610dd3565b50505050905090810190601f168015610e185780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b83811015610e4b578181015183820152602001610e33565b50505050905090810190601f168015610e785780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a1869750505050505050506112f7565b600083600160a060020a0316635887ab246040518163ffffffff1660e060020a02815260040160206040518083038186803b158015610ed657600080fd5b505afa158015610eea573d6000803e3d6000fd5b505050506040513d6020811015610f0057600080fd5b8101908080519060200190929190505050905080600160a060020a0316631a391cb48c6040518263ffffffff1660e060020a02815260040180806020018060200183810383526007815260200180600080516020611fc2833981519152815250602001838103825284818151815260200191508051906020019080838360005b83811015610f98578181015183820152602001610f80565b50505050905090810190601f168015610fc55780820380516001836020036101000a031916815260200191505b509350505050600060405180830381600087803b158015610fe557600080fd5b505af1158015610ff9573d6000803e3d6000fd5b50506040805160e360020a631bde84d3028152888e01602482015260048101829052600b6044820152600080516020611fa283398151915260648201529051600160a060020a038516935063def426989250608480830192600092919082900301818387803b15801561106b57600080fd5b505af115801561107f573d6000803e3d6000fd5b5050505083600160a060020a031663664b37d68c8387600160a060020a031663c74f8caf6040518163ffffffff1660e060020a02815260040160206040518083038186803b1580156110d057600080fd5b505afa1580156110e4573d6000803e3d6000fd5b505050506040513d60208110156110fa57600080fd5b505160405160e060020a63ffffffff8616028152600160a060020a0380841660248301528216604482015260606004820190815284516064830152845190918291608490910190602087019080838360005b8381101561116457818101518382015260200161114c565b50505050905090810190601f1680156111915780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b1580156111b257600080fd5b505af11580156111c6573d6000803e3d6000fd5b505050506040513d60208110156111dc57600080fd5b810190808051906020019092919050505050600080516020612002833981519152888d8d8d604051808581526020018060200180602001848152602001838103835286818151815260200191508051906020019080838360005b8381101561124e578181015183820152602001611236565b50505050905090810190601f16801561127b5780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b838110156112ae578181015183820152602001611296565b50505050905090810190601f1680156112db5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390a1509596505050505050505b9392505050565b600080808061130c8661048b565b9092509050811561161a576000611321611ee3565b9050600081600160a060020a0316635887ab246040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561136157600080fd5b505afa158015611375573d6000803e3d6000fd5b505050506040513d602081101561138b57600080fd5b50516040805160e260020a63068e472d0281526004810191825260076044820152600080516020611fc283398151915260648201526080602482019081528b5160848301528b51939450600160a060020a03851693631a391cb4938d9390928392909160a40190602086019080838360005b838110156114155781810151838201526020016113fd565b50505050905090810190601f1680156114425780820380516001836020036101000a031916815260200191505b509350505050600060405180830381600087803b15801561146257600080fd5b505af1158015611476573d6000803e3d6000fd5b50506040805160e360020a631bde84d3028152602481018b905260048101829052600b6044820152600080516020611fa283398151915260648201529051600160a060020a038516935063def426989250608480830192600092919082900301818387803b1580156114e757600080fd5b505af11580156114fb573d6000803e3d6000fd5b5050604080517f4c6f30c0000000000000000000000000000000000000000000000000000000008152600160a060020a038581166024830152600482019283528c5160448301528c51600095509087169350634c6f30c0928d92879282916064019060208601908083838c5b8381101561157f578181015183820152602001611567565b50505050905090810190601f1680156115ac5780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b1580156115cc57600080fd5b505af11580156115e0573d6000803e3d6000fd5b505050506040513d60208110156115f657600080fd5b50519050600181141561160c5760009550611612565b60011995505b505050611620565b60001992505b7f7ac7a04970319ae8fc5b92fe177d000fee3c00c92f8e78aae13d6571f17c351f8387876040518084815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561168d578181015183820152602001611675565b50505050905090810190601f1680156116ba5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a150909150505b92915050565b60008080806116e38661048b565b90925090508115611795576000199250600080516020611fe28339815191528387876040518084815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561174e578181015183820152602001611736565b50505050905090810190601f16801561177b5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a18293505050506116cf565b600061179f611ee3565b9050600081600160a060020a0316635887ab246040518163ffffffff1660e060020a02815260040160206040518083038186803b1580156117df57600080fd5b505afa1580156117f3573d6000803e3d6000fd5b505050506040513d602081101561180957600080fd5b50516040805160e260020a63068e472d0281526004810191825260076044820152600080516020611fc283398151915260648201526080602482019081528b5160848301528b51939450600160a060020a03851693631a391cb4938d9390928392909160a40190602086019080838360005b8381101561189357818101518382015260200161187b565b50505050905090810190601f1680156118c05780820380516001836020036101000a031916815260200191505b509350505050600060405180830381600087803b1580156118e057600080fd5b505af11580156118f4573d6000803e3d6000fd5b50506040805160e360020a631bde84d3028152868b01602482015260048101829052600b6044820152600080516020611fa283398151915260648201529051600160a060020a038516935063def426989250608480830192600092919082900301818387803b15801561196657600080fd5b505af115801561197a573d6000803e3d6000fd5b50505050600082600160a060020a031663664b37d68a8486600160a060020a031663c74f8caf6040518163ffffffff1660e060020a02815260040160206040518083038186803b1580156119cd57600080fd5b505afa1580156119e1573d6000803e3d6000fd5b505050506040513d60208110156119f757600080fd5b505160405160e060020a63ffffffff8616028152600160a060020a0380841660248301528216604482015260606004820190815284516064830152845190918291608490910190602087019080838360005b83811015611a61578181015183820152602001611a49565b50505050905090810190601f168015611a8e5780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b158015611aaf57600080fd5b505af1158015611ac3573d6000803e3d6000fd5b505050506040513d6020811015611ad957600080fd5b5051905060018114611b8f576004199550600080516020611fe2833981519152868a8a6040518084815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b83811015611b45578181015183820152602001611b2d565b50505050905090810190601f168015611b725780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a18596505050505050506116cf565b600080516020611fe2833981519152868a8a6040518084815260200180602001838152602001828103825284818151815260200191508051906020019080838360005b83811015611bea578181015183820152602001611bd2565b50505050905090810190601f168015611c175780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a15093979650505050505050565b6000808080611c418661048b565b90925090508115611cab576000199250600080516020611fe28339815191528387876040518084815260200180602001838152602001828103825284818151815260200191508051906020019080838360008381101561174e578181015183820152602001611736565b84811015611d12576002199250600080516020611fe28339815191528387876040518084815260200180602001838152602001828103825284818151815260200191508051906020019080838360008381101561174e578181015183820152602001611736565b6000611d1c611ee3565b9050600081600160a060020a0316635887ab246040518163ffffffff1660e060020a02815260040160206040518083038186803b158015611d5c57600080fd5b505afa158015611d70573d6000803e3d6000fd5b505050506040513d6020811015611d8657600080fd5b50516040805160e260020a63068e472d0281526004810191825260076044820152600080516020611fc283398151915260648201526080602482019081528b5160848301528b51939450600160a060020a03851693631a391cb4938d9390928392909160a40190602086019080838360005b83811015611e10578181015183820152602001611df8565b50505050905090810190601f168015611e3d5780820380516001836020036101000a031916815260200191505b509350505050600060405180830381600087803b158015611e5d57600080fd5b505af1158015611e71573d6000803e3d6000fd5b50506040805160e360020a631bde84d30281528a8703602482015260048101829052600b6044820152600080516020611fa283398151915260648201529051600160a060020a038516935063def426989250608480830192600092919082900301818387803b15801561196657600080fd5b604080517f59a48b65000000000000000000000000000000000000000000000000000000008152602060048201819052600760248301527f745f6173736574000000000000000000000000000000000000000000000000006044830152915160009261100192849284926359a48b659260648082019391829003018186803b158015611f6e57600080fd5b505afa158015611f82573d6000803e3d6000fd5b505050506040513d6020811015611f9857600080fd5b5051925050509056fe61737365745f76616c75650000000000000000000000000000000000000000006163636f756e7400000000000000000000000000000000000000000000000000012fc7b91786496323158059812ae81fee87db419a88e70955aeaaf6201cd997105af2c562df33af7eaa9de5fb0c18d8d30f281a18f95a8f76b44353a322693ca165627a7a72305820400aae77db4eb99d99717423073740de973cd34526acb0870725dac3e0d06c6f0029"

// DeployAsset deploys a new contract, binding an instance of Asset to it.
func DeployAsset(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Asset, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AssetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Asset{AssetCaller: AssetCaller{contract: contract}, AssetTransactor: AssetTransactor{contract: contract}, AssetFilterer: AssetFilterer{contract: contract}}, nil
}

// Asset is an auto generated Go binding around a Solidity contract.
type Asset struct {
	AssetCaller     // Read-only binding to the contract
	AssetTransactor // Write-only binding to the contract
	AssetFilterer   // Log filterer for contract events
}

// AssetCaller is an auto generated read-only Go binding around a Solidity contract.
type AssetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetTransactor is an auto generated write-only Go binding around a Solidity contract.
type AssetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type AssetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type AssetSession struct {
	Contract     *Asset            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type AssetCallerSession struct {
	Contract *AssetCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AssetTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type AssetTransactorSession struct {
	Contract     *AssetTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetRaw is an auto generated low-level Go binding around a Solidity contract.
type AssetRaw struct {
	Contract *Asset // Generic contract binding to access the raw methods on
}

// AssetCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type AssetCallerRaw struct {
	Contract *AssetCaller // Generic read-only contract binding to access the raw methods on
}

// AssetTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type AssetTransactorRaw struct {
	Contract *AssetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAsset creates a new instance of Asset, bound to a specific deployed contract.
func NewAsset(address common.Address, backend bind.ContractBackend) (*Asset, error) {
	contract, err := bindAsset(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Asset{AssetCaller: AssetCaller{contract: contract}, AssetTransactor: AssetTransactor{contract: contract}, AssetFilterer: AssetFilterer{contract: contract}}, nil
}

// NewAssetCaller creates a new read-only instance of Asset, bound to a specific deployed contract.
func NewAssetCaller(address common.Address, caller bind.ContractCaller) (*AssetCaller, error) {
	contract, err := bindAsset(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetCaller{contract: contract}, nil
}

// NewAssetTransactor creates a new write-only instance of Asset, bound to a specific deployed contract.
func NewAssetTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetTransactor, error) {
	contract, err := bindAsset(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetTransactor{contract: contract}, nil
}

// NewAssetFilterer creates a new log filterer instance of Asset, bound to a specific deployed contract.
func NewAssetFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetFilterer, error) {
	contract, err := bindAsset(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetFilterer{contract: contract}, nil
}

// bindAsset binds a generic wrapper to an already deployed contract.
func bindAsset(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Asset *AssetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Asset.Contract.AssetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Asset *AssetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Asset.Contract.AssetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Asset *AssetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Asset.Contract.AssetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Asset *AssetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Asset.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Asset *AssetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Asset.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Asset *AssetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Asset.Contract.contract.Transact(opts, method, params...)
}

// Select is a free data retrieval call binding the contract method 0x5b325d78.
//
// Solidity: function select(string account) constant returns(int256, uint256)
func (_Asset *AssetCaller) Select(opts *bind.CallOpts, account string) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Asset.contract.Call(opts, out, "select", account)
	return *ret0, *ret1, err
}

// Select is a free data retrieval call binding the contract method 0x5b325d78.
//
// Solidity: function select(string account) constant returns(int256, uint256)
func (_Asset *AssetSession) Select(account string) (*big.Int, *big.Int, error) {
	return _Asset.Contract.Select(&_Asset.CallOpts, account)
}

// Select is a free data retrieval call binding the contract method 0x5b325d78.
//
// Solidity: function select(string account) constant returns(int256, uint256)
func (_Asset *AssetCallerSession) Select(account string) (*big.Int, *big.Int, error) {
	return _Asset.Contract.Select(&_Asset.CallOpts, account)
}

// Decrease is a paid mutator transaction binding the contract method 0xf8c6f6a6.
//
// Solidity: function decrease(string from_account, uint256 amount) returns(int256)
func (_Asset *AssetTransactor) Decrease(opts *bind.TransactOpts, from_account string, amount *big.Int) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "decrease", from_account, amount)
}

// Decrease is a paid mutator transaction binding the contract method 0xf8c6f6a6.
//
// Solidity: function decrease(string from_account, uint256 amount) returns(int256)
func (_Asset *AssetSession) Decrease(from_account string, amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.Decrease(&_Asset.TransactOpts, from_account, amount)
}

// Decrease is a paid mutator transaction binding the contract method 0xf8c6f6a6.
//
// Solidity: function decrease(string from_account, uint256 amount) returns(int256)
func (_Asset *AssetTransactorSession) Decrease(from_account string, amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.Decrease(&_Asset.TransactOpts, from_account, amount)
}

// Increase is a paid mutator transaction binding the contract method 0xc68e3b0b.
//
// Solidity: function increase(string from_account, uint256 amount) returns(int256)
func (_Asset *AssetTransactor) Increase(opts *bind.TransactOpts, from_account string, amount *big.Int) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "increase", from_account, amount)
}

// Increase is a paid mutator transaction binding the contract method 0xc68e3b0b.
//
// Solidity: function increase(string from_account, uint256 amount) returns(int256)
func (_Asset *AssetSession) Increase(from_account string, amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.Increase(&_Asset.TransactOpts, from_account, amount)
}

// Increase is a paid mutator transaction binding the contract method 0xc68e3b0b.
//
// Solidity: function increase(string from_account, uint256 amount) returns(int256)
func (_Asset *AssetTransactorSession) Increase(from_account string, amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.Increase(&_Asset.TransactOpts, from_account, amount)
}

// Register is a paid mutator transaction binding the contract method 0xb433c7ca.
//
// Solidity: function register(string account, uint256 asset_value) returns(int256)
func (_Asset *AssetTransactor) Register(opts *bind.TransactOpts, account string, asset_value *big.Int) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "register", account, asset_value)
}

// Register is a paid mutator transaction binding the contract method 0xb433c7ca.
//
// Solidity: function register(string account, uint256 asset_value) returns(int256)
func (_Asset *AssetSession) Register(account string, asset_value *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.Register(&_Asset.TransactOpts, account, asset_value)
}

// Register is a paid mutator transaction binding the contract method 0xb433c7ca.
//
// Solidity: function register(string account, uint256 asset_value) returns(int256)
func (_Asset *AssetTransactorSession) Register(account string, asset_value *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.Register(&_Asset.TransactOpts, account, asset_value)
}

// Transfer is a paid mutator transaction binding the contract method 0x612d2bff.
//
// Solidity: function transfer(string from_account, string to_account, uint256 amount) returns(int256)
func (_Asset *AssetTransactor) Transfer(opts *bind.TransactOpts, from_account string, to_account string, amount *big.Int) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "transfer", from_account, to_account, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x612d2bff.
//
// Solidity: function transfer(string from_account, string to_account, uint256 amount) returns(int256)
func (_Asset *AssetSession) Transfer(from_account string, to_account string, amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.Transfer(&_Asset.TransactOpts, from_account, to_account, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x612d2bff.
//
// Solidity: function transfer(string from_account, string to_account, uint256 amount) returns(int256)
func (_Asset *AssetTransactorSession) Transfer(from_account string, to_account string, amount *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.Transfer(&_Asset.TransactOpts, from_account, to_account, amount)
}

// AssetIncreaseEventIterator is returned from FilterIncreaseEvent and is used to iterate over the raw logs and unpacked data for IncreaseEvent events raised by the Asset contract.
type AssetIncreaseEventIterator struct {
	Event *AssetIncreaseEvent // Event containing the contract specifics and raw log

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
func (it *AssetIncreaseEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetIncreaseEvent)
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
		it.Event = new(AssetIncreaseEvent)
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
func (it *AssetIncreaseEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetIncreaseEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetIncreaseEvent represents a IncreaseEvent event raised by the Asset contract.
type AssetIncreaseEvent struct {
	Ret         *big.Int
	FromAccount string
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterIncreaseEvent is a free log retrieval operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000012fc7b9.
//
// Solidity: event IncreaseEvent(int256 ret, string from_account, uint256 amount)
func (_Asset *AssetFilterer) FilterIncreaseEvent(opts *bind.FilterOpts) (*AssetIncreaseEventIterator, error) {

	logs, sub, err := _Asset.contract.FilterLogs(opts, "IncreaseEvent")
	if err != nil {
		return nil, err
	}
	return &AssetIncreaseEventIterator{contract: _Asset.contract, event: "IncreaseEvent", logs: logs, sub: sub}, nil
}

// WatchIncreaseEvent is a free log subscription operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000012fc7b9.
//
// Solidity: event IncreaseEvent(int256 ret, string from_account, uint256 amount)
func (_Asset *AssetFilterer) WatchIncreaseEvent(opts *bind.WatchOpts, sink chan<- *AssetIncreaseEvent) (event.Subscription, error) {

	logs, sub, err := _Asset.contract.WatchLogs(opts, "IncreaseEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetIncreaseEvent)
				if err := _Asset.contract.UnpackLog(event, "IncreaseEvent", log); err != nil {
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

// ParseIncreaseEvent is a log parse operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000012fc7b9.
//
// Solidity: event IncreaseEvent(int256 ret, string from_account, uint256 amount)
func (_Asset *AssetFilterer) ParseIncreaseEvent(log types.Log) (*AssetIncreaseEvent, error) {
	event := new(AssetIncreaseEvent)
	if err := _Asset.contract.UnpackLog(event, "IncreaseEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AssetRegisterEventIterator is returned from FilterRegisterEvent and is used to iterate over the raw logs and unpacked data for RegisterEvent events raised by the Asset contract.
type AssetRegisterEventIterator struct {
	Event *AssetRegisterEvent // Event containing the contract specifics and raw log

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
func (it *AssetRegisterEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetRegisterEvent)
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
		it.Event = new(AssetRegisterEvent)
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
func (it *AssetRegisterEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetRegisterEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetRegisterEvent represents a RegisterEvent event raised by the Asset contract.
type AssetRegisterEvent struct {
	Ret        *big.Int
	Account    string
	AssetValue *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRegisterEvent is a free log retrieval operation binding the contract event 0x000000000000000000000000000000000000000000000000000000007ac7a049.
//
// Solidity: event RegisterEvent(int256 ret, string account, uint256 asset_value)
func (_Asset *AssetFilterer) FilterRegisterEvent(opts *bind.FilterOpts) (*AssetRegisterEventIterator, error) {

	logs, sub, err := _Asset.contract.FilterLogs(opts, "RegisterEvent")
	if err != nil {
		return nil, err
	}
	return &AssetRegisterEventIterator{contract: _Asset.contract, event: "RegisterEvent", logs: logs, sub: sub}, nil
}

// WatchRegisterEvent is a free log subscription operation binding the contract event 0x000000000000000000000000000000000000000000000000000000007ac7a049.
//
// Solidity: event RegisterEvent(int256 ret, string account, uint256 asset_value)
func (_Asset *AssetFilterer) WatchRegisterEvent(opts *bind.WatchOpts, sink chan<- *AssetRegisterEvent) (event.Subscription, error) {

	logs, sub, err := _Asset.contract.WatchLogs(opts, "RegisterEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetRegisterEvent)
				if err := _Asset.contract.UnpackLog(event, "RegisterEvent", log); err != nil {
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

// ParseRegisterEvent is a log parse operation binding the contract event 0x000000000000000000000000000000000000000000000000000000007ac7a049.
//
// Solidity: event RegisterEvent(int256 ret, string account, uint256 asset_value)
func (_Asset *AssetFilterer) ParseRegisterEvent(log types.Log) (*AssetRegisterEvent, error) {
	event := new(AssetRegisterEvent)
	if err := _Asset.contract.UnpackLog(event, "RegisterEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AssetTransferEventIterator is returned from FilterTransferEvent and is used to iterate over the raw logs and unpacked data for TransferEvent events raised by the Asset contract.
type AssetTransferEventIterator struct {
	Event *AssetTransferEvent // Event containing the contract specifics and raw log

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
func (it *AssetTransferEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetTransferEvent)
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
		it.Event = new(AssetTransferEvent)
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
func (it *AssetTransferEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetTransferEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetTransferEvent represents a TransferEvent event raised by the Asset contract.
type AssetTransferEvent struct {
	Ret         *big.Int
	FromAccount string
	ToAccount   string
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferEvent is a free log retrieval operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000105af2c5.
//
// Solidity: event TransferEvent(int256 ret, string from_account, string to_account, uint256 amount)
func (_Asset *AssetFilterer) FilterTransferEvent(opts *bind.FilterOpts) (*AssetTransferEventIterator, error) {

	logs, sub, err := _Asset.contract.FilterLogs(opts, "TransferEvent")
	if err != nil {
		return nil, err
	}
	return &AssetTransferEventIterator{contract: _Asset.contract, event: "TransferEvent", logs: logs, sub: sub}, nil
}

// WatchTransferEvent is a free log subscription operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000105af2c5.
//
// Solidity: event TransferEvent(int256 ret, string from_account, string to_account, uint256 amount)
func (_Asset *AssetFilterer) WatchTransferEvent(opts *bind.WatchOpts, sink chan<- *AssetTransferEvent) (event.Subscription, error) {

	logs, sub, err := _Asset.contract.WatchLogs(opts, "TransferEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetTransferEvent)
				if err := _Asset.contract.UnpackLog(event, "TransferEvent", log); err != nil {
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

// ParseTransferEvent is a log parse operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000105af2c5.
//
// Solidity: event TransferEvent(int256 ret, string from_account, string to_account, uint256 amount)
func (_Asset *AssetFilterer) ParseTransferEvent(log types.Log) (*AssetTransferEvent, error) {
	event := new(AssetTransferEvent)
	if err := _Asset.contract.UnpackLog(event, "TransferEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ConditionABI is the input ABI used to generate the binding from.
const ConditionABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"int256\"}],\"name\":\"LT\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"int256\"}],\"name\":\"LE\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"int256\"}],\"name\":\"limit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"int256\"}],\"name\":\"GE\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"name\":\"NE\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"name\":\"EQ\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"name\":\"limit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"int256\"}],\"name\":\"EQ\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"int256\"}],\"name\":\"GT\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"int256\"}],\"name\":\"NE\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ConditionFuncSigs maps the 4-byte function signature to its string representation.
var ConditionFuncSigs = map[string]string{
	"d62b54b4": "EQ(string,int256)",
	"ae763db5": "EQ(string,string)",
	"40b5f1ab": "GE(string,int256)",
	"e1f10ee0": "GT(string,int256)",
	"2ec346c9": "LE(string,int256)",
	"1ee88791": "LT(string,int256)",
	"f955264b": "NE(string,int256)",
	"966b0822": "NE(string,string)",
	"bd8cb043": "limit(int256)",
	"32492737": "limit(int256,int256)",
}

// Condition is an auto generated Go binding around a Solidity contract.
type Condition struct {
	ConditionCaller     // Read-only binding to the contract
	ConditionTransactor // Write-only binding to the contract
	ConditionFilterer   // Log filterer for contract events
}

// ConditionCaller is an auto generated read-only Go binding around a Solidity contract.
type ConditionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConditionTransactor is an auto generated write-only Go binding around a Solidity contract.
type ConditionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConditionFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type ConditionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConditionSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ConditionSession struct {
	Contract     *Condition        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConditionCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ConditionCallerSession struct {
	Contract *ConditionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ConditionTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ConditionTransactorSession struct {
	Contract     *ConditionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ConditionRaw is an auto generated low-level Go binding around a Solidity contract.
type ConditionRaw struct {
	Contract *Condition // Generic contract binding to access the raw methods on
}

// ConditionCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ConditionCallerRaw struct {
	Contract *ConditionCaller // Generic read-only contract binding to access the raw methods on
}

// ConditionTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ConditionTransactorRaw struct {
	Contract *ConditionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCondition creates a new instance of Condition, bound to a specific deployed contract.
func NewCondition(address common.Address, backend bind.ContractBackend) (*Condition, error) {
	contract, err := bindCondition(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Condition{ConditionCaller: ConditionCaller{contract: contract}, ConditionTransactor: ConditionTransactor{contract: contract}, ConditionFilterer: ConditionFilterer{contract: contract}}, nil
}

// NewConditionCaller creates a new read-only instance of Condition, bound to a specific deployed contract.
func NewConditionCaller(address common.Address, caller bind.ContractCaller) (*ConditionCaller, error) {
	contract, err := bindCondition(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConditionCaller{contract: contract}, nil
}

// NewConditionTransactor creates a new write-only instance of Condition, bound to a specific deployed contract.
func NewConditionTransactor(address common.Address, transactor bind.ContractTransactor) (*ConditionTransactor, error) {
	contract, err := bindCondition(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConditionTransactor{contract: contract}, nil
}

// NewConditionFilterer creates a new log filterer instance of Condition, bound to a specific deployed contract.
func NewConditionFilterer(address common.Address, filterer bind.ContractFilterer) (*ConditionFilterer, error) {
	contract, err := bindCondition(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConditionFilterer{contract: contract}, nil
}

// bindCondition binds a generic wrapper to an already deployed contract.
func bindCondition(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConditionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Condition *ConditionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Condition.Contract.ConditionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Condition *ConditionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Condition.Contract.ConditionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Condition *ConditionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Condition.Contract.ConditionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Condition *ConditionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Condition.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Condition *ConditionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Condition.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Condition *ConditionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Condition.Contract.contract.Transact(opts, method, params...)
}

// EQ is a paid mutator transaction binding the contract method 0xae763db5.
//
// Solidity: function EQ(string , string ) returns()
func (_Condition *ConditionTransactor) EQ(opts *bind.TransactOpts, arg0 string, arg1 string) (*types.Transaction, error) {
	return _Condition.contract.Transact(opts, "EQ", arg0, arg1)
}

// EQ is a paid mutator transaction binding the contract method 0xae763db5.
//
// Solidity: function EQ(string , string ) returns()
func (_Condition *ConditionSession) EQ(arg0 string, arg1 string) (*types.Transaction, error) {
	return _Condition.Contract.EQ(&_Condition.TransactOpts, arg0, arg1)
}

// EQ is a paid mutator transaction binding the contract method 0xae763db5.
//
// Solidity: function EQ(string , string ) returns()
func (_Condition *ConditionTransactorSession) EQ(arg0 string, arg1 string) (*types.Transaction, error) {
	return _Condition.Contract.EQ(&_Condition.TransactOpts, arg0, arg1)
}

// EQ0 is a paid mutator transaction binding the contract method 0xd62b54b4.
//
// Solidity: function EQ(string , int256 ) returns()
func (_Condition *ConditionTransactor) EQ0(opts *bind.TransactOpts, arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Condition.contract.Transact(opts, "EQ0", arg0, arg1)
}

// EQ0 is a paid mutator transaction binding the contract method 0xd62b54b4.
//
// Solidity: function EQ(string , int256 ) returns()
func (_Condition *ConditionSession) EQ0(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Condition.Contract.EQ0(&_Condition.TransactOpts, arg0, arg1)
}

// EQ0 is a paid mutator transaction binding the contract method 0xd62b54b4.
//
// Solidity: function EQ(string , int256 ) returns()
func (_Condition *ConditionTransactorSession) EQ0(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Condition.Contract.EQ0(&_Condition.TransactOpts, arg0, arg1)
}

// GE is a paid mutator transaction binding the contract method 0x40b5f1ab.
//
// Solidity: function GE(string , int256 ) returns()
func (_Condition *ConditionTransactor) GE(opts *bind.TransactOpts, arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Condition.contract.Transact(opts, "GE", arg0, arg1)
}

// GE is a paid mutator transaction binding the contract method 0x40b5f1ab.
//
// Solidity: function GE(string , int256 ) returns()
func (_Condition *ConditionSession) GE(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Condition.Contract.GE(&_Condition.TransactOpts, arg0, arg1)
}

// GE is a paid mutator transaction binding the contract method 0x40b5f1ab.
//
// Solidity: function GE(string , int256 ) returns()
func (_Condition *ConditionTransactorSession) GE(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Condition.Contract.GE(&_Condition.TransactOpts, arg0, arg1)
}

// GT is a paid mutator transaction binding the contract method 0xe1f10ee0.
//
// Solidity: function GT(string , int256 ) returns()
func (_Condition *ConditionTransactor) GT(opts *bind.TransactOpts, arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Condition.contract.Transact(opts, "GT", arg0, arg1)
}

// GT is a paid mutator transaction binding the contract method 0xe1f10ee0.
//
// Solidity: function GT(string , int256 ) returns()
func (_Condition *ConditionSession) GT(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Condition.Contract.GT(&_Condition.TransactOpts, arg0, arg1)
}

// GT is a paid mutator transaction binding the contract method 0xe1f10ee0.
//
// Solidity: function GT(string , int256 ) returns()
func (_Condition *ConditionTransactorSession) GT(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Condition.Contract.GT(&_Condition.TransactOpts, arg0, arg1)
}

// LE is a paid mutator transaction binding the contract method 0x2ec346c9.
//
// Solidity: function LE(string , int256 ) returns()
func (_Condition *ConditionTransactor) LE(opts *bind.TransactOpts, arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Condition.contract.Transact(opts, "LE", arg0, arg1)
}

// LE is a paid mutator transaction binding the contract method 0x2ec346c9.
//
// Solidity: function LE(string , int256 ) returns()
func (_Condition *ConditionSession) LE(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Condition.Contract.LE(&_Condition.TransactOpts, arg0, arg1)
}

// LE is a paid mutator transaction binding the contract method 0x2ec346c9.
//
// Solidity: function LE(string , int256 ) returns()
func (_Condition *ConditionTransactorSession) LE(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Condition.Contract.LE(&_Condition.TransactOpts, arg0, arg1)
}

// LT is a paid mutator transaction binding the contract method 0x1ee88791.
//
// Solidity: function LT(string , int256 ) returns()
func (_Condition *ConditionTransactor) LT(opts *bind.TransactOpts, arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Condition.contract.Transact(opts, "LT", arg0, arg1)
}

// LT is a paid mutator transaction binding the contract method 0x1ee88791.
//
// Solidity: function LT(string , int256 ) returns()
func (_Condition *ConditionSession) LT(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Condition.Contract.LT(&_Condition.TransactOpts, arg0, arg1)
}

// LT is a paid mutator transaction binding the contract method 0x1ee88791.
//
// Solidity: function LT(string , int256 ) returns()
func (_Condition *ConditionTransactorSession) LT(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Condition.Contract.LT(&_Condition.TransactOpts, arg0, arg1)
}

// NE is a paid mutator transaction binding the contract method 0x966b0822.
//
// Solidity: function NE(string , string ) returns()
func (_Condition *ConditionTransactor) NE(opts *bind.TransactOpts, arg0 string, arg1 string) (*types.Transaction, error) {
	return _Condition.contract.Transact(opts, "NE", arg0, arg1)
}

// NE is a paid mutator transaction binding the contract method 0x966b0822.
//
// Solidity: function NE(string , string ) returns()
func (_Condition *ConditionSession) NE(arg0 string, arg1 string) (*types.Transaction, error) {
	return _Condition.Contract.NE(&_Condition.TransactOpts, arg0, arg1)
}

// NE is a paid mutator transaction binding the contract method 0x966b0822.
//
// Solidity: function NE(string , string ) returns()
func (_Condition *ConditionTransactorSession) NE(arg0 string, arg1 string) (*types.Transaction, error) {
	return _Condition.Contract.NE(&_Condition.TransactOpts, arg0, arg1)
}

// NE0 is a paid mutator transaction binding the contract method 0xf955264b.
//
// Solidity: function NE(string , int256 ) returns()
func (_Condition *ConditionTransactor) NE0(opts *bind.TransactOpts, arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Condition.contract.Transact(opts, "NE0", arg0, arg1)
}

// NE0 is a paid mutator transaction binding the contract method 0xf955264b.
//
// Solidity: function NE(string , int256 ) returns()
func (_Condition *ConditionSession) NE0(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Condition.Contract.NE0(&_Condition.TransactOpts, arg0, arg1)
}

// NE0 is a paid mutator transaction binding the contract method 0xf955264b.
//
// Solidity: function NE(string , int256 ) returns()
func (_Condition *ConditionTransactorSession) NE0(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Condition.Contract.NE0(&_Condition.TransactOpts, arg0, arg1)
}

// Limit is a paid mutator transaction binding the contract method 0x32492737.
//
// Solidity: function limit(int256 , int256 ) returns()
func (_Condition *ConditionTransactor) Limit(opts *bind.TransactOpts, arg0 *big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _Condition.contract.Transact(opts, "limit", arg0, arg1)
}

// Limit is a paid mutator transaction binding the contract method 0x32492737.
//
// Solidity: function limit(int256 , int256 ) returns()
func (_Condition *ConditionSession) Limit(arg0 *big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _Condition.Contract.Limit(&_Condition.TransactOpts, arg0, arg1)
}

// Limit is a paid mutator transaction binding the contract method 0x32492737.
//
// Solidity: function limit(int256 , int256 ) returns()
func (_Condition *ConditionTransactorSession) Limit(arg0 *big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _Condition.Contract.Limit(&_Condition.TransactOpts, arg0, arg1)
}

// Limit0 is a paid mutator transaction binding the contract method 0xbd8cb043.
//
// Solidity: function limit(int256 ) returns()
func (_Condition *ConditionTransactor) Limit0(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Condition.contract.Transact(opts, "limit0", arg0)
}

// Limit0 is a paid mutator transaction binding the contract method 0xbd8cb043.
//
// Solidity: function limit(int256 ) returns()
func (_Condition *ConditionSession) Limit0(arg0 *big.Int) (*types.Transaction, error) {
	return _Condition.Contract.Limit0(&_Condition.TransactOpts, arg0)
}

// Limit0 is a paid mutator transaction binding the contract method 0xbd8cb043.
//
// Solidity: function limit(int256 ) returns()
func (_Condition *ConditionTransactorSession) Limit0(arg0 *big.Int) (*types.Transaction, error) {
	return _Condition.Contract.Limit0(&_Condition.TransactOpts, arg0)
}

// EntriesABI is the input ABI used to generate the binding from.
const EntriesABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"size\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EntriesFuncSigs maps the 4-byte function signature to its string representation.
var EntriesFuncSigs = map[string]string{
	"3dd2b614": "get(int256)",
	"d3e9af5a": "size()",
}

// Entries is an auto generated Go binding around a Solidity contract.
type Entries struct {
	EntriesCaller     // Read-only binding to the contract
	EntriesTransactor // Write-only binding to the contract
	EntriesFilterer   // Log filterer for contract events
}

// EntriesCaller is an auto generated read-only Go binding around a Solidity contract.
type EntriesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntriesTransactor is an auto generated write-only Go binding around a Solidity contract.
type EntriesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntriesFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type EntriesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntriesSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type EntriesSession struct {
	Contract     *Entries          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EntriesCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type EntriesCallerSession struct {
	Contract *EntriesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// EntriesTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type EntriesTransactorSession struct {
	Contract     *EntriesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EntriesRaw is an auto generated low-level Go binding around a Solidity contract.
type EntriesRaw struct {
	Contract *Entries // Generic contract binding to access the raw methods on
}

// EntriesCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type EntriesCallerRaw struct {
	Contract *EntriesCaller // Generic read-only contract binding to access the raw methods on
}

// EntriesTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type EntriesTransactorRaw struct {
	Contract *EntriesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEntries creates a new instance of Entries, bound to a specific deployed contract.
func NewEntries(address common.Address, backend bind.ContractBackend) (*Entries, error) {
	contract, err := bindEntries(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Entries{EntriesCaller: EntriesCaller{contract: contract}, EntriesTransactor: EntriesTransactor{contract: contract}, EntriesFilterer: EntriesFilterer{contract: contract}}, nil
}

// NewEntriesCaller creates a new read-only instance of Entries, bound to a specific deployed contract.
func NewEntriesCaller(address common.Address, caller bind.ContractCaller) (*EntriesCaller, error) {
	contract, err := bindEntries(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EntriesCaller{contract: contract}, nil
}

// NewEntriesTransactor creates a new write-only instance of Entries, bound to a specific deployed contract.
func NewEntriesTransactor(address common.Address, transactor bind.ContractTransactor) (*EntriesTransactor, error) {
	contract, err := bindEntries(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EntriesTransactor{contract: contract}, nil
}

// NewEntriesFilterer creates a new log filterer instance of Entries, bound to a specific deployed contract.
func NewEntriesFilterer(address common.Address, filterer bind.ContractFilterer) (*EntriesFilterer, error) {
	contract, err := bindEntries(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EntriesFilterer{contract: contract}, nil
}

// bindEntries binds a generic wrapper to an already deployed contract.
func bindEntries(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EntriesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Entries *EntriesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Entries.Contract.EntriesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Entries *EntriesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Entries.Contract.EntriesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Entries *EntriesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Entries.Contract.EntriesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Entries *EntriesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Entries.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Entries *EntriesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Entries.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Entries *EntriesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Entries.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x3dd2b614.
//
// Solidity: function get(int256 ) constant returns(address)
func (_Entries *EntriesCaller) Get(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Entries.contract.Call(opts, out, "get", arg0)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x3dd2b614.
//
// Solidity: function get(int256 ) constant returns(address)
func (_Entries *EntriesSession) Get(arg0 *big.Int) (common.Address, error) {
	return _Entries.Contract.Get(&_Entries.CallOpts, arg0)
}

// Get is a free data retrieval call binding the contract method 0x3dd2b614.
//
// Solidity: function get(int256 ) constant returns(address)
func (_Entries *EntriesCallerSession) Get(arg0 *big.Int) (common.Address, error) {
	return _Entries.Contract.Get(&_Entries.CallOpts, arg0)
}

// Size is a free data retrieval call binding the contract method 0xd3e9af5a.
//
// Solidity: function size() constant returns(int256)
func (_Entries *EntriesCaller) Size(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Entries.contract.Call(opts, out, "size")
	return *ret0, err
}

// Size is a free data retrieval call binding the contract method 0xd3e9af5a.
//
// Solidity: function size() constant returns(int256)
func (_Entries *EntriesSession) Size() (*big.Int, error) {
	return _Entries.Contract.Size(&_Entries.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0xd3e9af5a.
//
// Solidity: function size() constant returns(int256)
func (_Entries *EntriesCallerSession) Size() (*big.Int, error) {
	return _Entries.Contract.Size(&_Entries.CallOpts)
}

// EntryABI is the input ABI used to generate the binding from.
const EntryABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"string\"}],\"name\":\"getAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"string\"}],\"name\":\"getInt\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"string\"}],\"name\":\"getBytes64\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes1[64]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"string\"}],\"name\":\"getString\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"int256\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"string\"}],\"name\":\"getUInt\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"string\"}],\"name\":\"getBytes32\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EntryFuncSigs maps the 4-byte function signature to its string representation.
var EntryFuncSigs = map[string]string{
	"07afbf3a": "getAddress(string)",
	"fdebe414": "getBytes32(string)",
	"9139fa37": "getBytes64(string)",
	"4900862e": "getInt(string)",
	"9bca41e8": "getString(string)",
	"df7427af": "getUInt(string)",
	"517c4dd9": "set(string,address)",
	"def42698": "set(string,int256)",
	"1a391cb4": "set(string,string)",
	"f2f4ee6d": "set(string,uint256)",
}

// Entry is an auto generated Go binding around a Solidity contract.
type Entry struct {
	EntryCaller     // Read-only binding to the contract
	EntryTransactor // Write-only binding to the contract
	EntryFilterer   // Log filterer for contract events
}

// EntryCaller is an auto generated read-only Go binding around a Solidity contract.
type EntryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntryTransactor is an auto generated write-only Go binding around a Solidity contract.
type EntryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntryFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type EntryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EntrySession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type EntrySession struct {
	Contract     *Entry            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EntryCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type EntryCallerSession struct {
	Contract *EntryCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EntryTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type EntryTransactorSession struct {
	Contract     *EntryTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EntryRaw is an auto generated low-level Go binding around a Solidity contract.
type EntryRaw struct {
	Contract *Entry // Generic contract binding to access the raw methods on
}

// EntryCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type EntryCallerRaw struct {
	Contract *EntryCaller // Generic read-only contract binding to access the raw methods on
}

// EntryTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type EntryTransactorRaw struct {
	Contract *EntryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEntry creates a new instance of Entry, bound to a specific deployed contract.
func NewEntry(address common.Address, backend bind.ContractBackend) (*Entry, error) {
	contract, err := bindEntry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Entry{EntryCaller: EntryCaller{contract: contract}, EntryTransactor: EntryTransactor{contract: contract}, EntryFilterer: EntryFilterer{contract: contract}}, nil
}

// NewEntryCaller creates a new read-only instance of Entry, bound to a specific deployed contract.
func NewEntryCaller(address common.Address, caller bind.ContractCaller) (*EntryCaller, error) {
	contract, err := bindEntry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EntryCaller{contract: contract}, nil
}

// NewEntryTransactor creates a new write-only instance of Entry, bound to a specific deployed contract.
func NewEntryTransactor(address common.Address, transactor bind.ContractTransactor) (*EntryTransactor, error) {
	contract, err := bindEntry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EntryTransactor{contract: contract}, nil
}

// NewEntryFilterer creates a new log filterer instance of Entry, bound to a specific deployed contract.
func NewEntryFilterer(address common.Address, filterer bind.ContractFilterer) (*EntryFilterer, error) {
	contract, err := bindEntry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EntryFilterer{contract: contract}, nil
}

// bindEntry binds a generic wrapper to an already deployed contract.
func bindEntry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EntryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Entry *EntryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Entry.Contract.EntryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Entry *EntryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Entry.Contract.EntryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Entry *EntryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Entry.Contract.EntryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Entry *EntryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Entry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Entry *EntryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Entry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Entry *EntryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Entry.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0x07afbf3a.
//
// Solidity: function getAddress(string ) constant returns(address)
func (_Entry *EntryCaller) GetAddress(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Entry.contract.Call(opts, out, "getAddress", arg0)
	return *ret0, err
}

// GetAddress is a free data retrieval call binding the contract method 0x07afbf3a.
//
// Solidity: function getAddress(string ) constant returns(address)
func (_Entry *EntrySession) GetAddress(arg0 string) (common.Address, error) {
	return _Entry.Contract.GetAddress(&_Entry.CallOpts, arg0)
}

// GetAddress is a free data retrieval call binding the contract method 0x07afbf3a.
//
// Solidity: function getAddress(string ) constant returns(address)
func (_Entry *EntryCallerSession) GetAddress(arg0 string) (common.Address, error) {
	return _Entry.Contract.GetAddress(&_Entry.CallOpts, arg0)
}

// GetBytes32 is a free data retrieval call binding the contract method 0xfdebe414.
//
// Solidity: function getBytes32(string ) constant returns(bytes32)
func (_Entry *EntryCaller) GetBytes32(opts *bind.CallOpts, arg0 string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Entry.contract.Call(opts, out, "getBytes32", arg0)
	return *ret0, err
}

// GetBytes32 is a free data retrieval call binding the contract method 0xfdebe414.
//
// Solidity: function getBytes32(string ) constant returns(bytes32)
func (_Entry *EntrySession) GetBytes32(arg0 string) ([32]byte, error) {
	return _Entry.Contract.GetBytes32(&_Entry.CallOpts, arg0)
}

// GetBytes32 is a free data retrieval call binding the contract method 0xfdebe414.
//
// Solidity: function getBytes32(string ) constant returns(bytes32)
func (_Entry *EntryCallerSession) GetBytes32(arg0 string) ([32]byte, error) {
	return _Entry.Contract.GetBytes32(&_Entry.CallOpts, arg0)
}

// GetBytes64 is a free data retrieval call binding the contract method 0x9139fa37.
//
// Solidity: function getBytes64(string ) constant returns(bytes1[64])
func (_Entry *EntryCaller) GetBytes64(opts *bind.CallOpts, arg0 string) ([64][1]byte, error) {
	var (
		ret0 = new([64][1]byte)
	)
	out := ret0
	err := _Entry.contract.Call(opts, out, "getBytes64", arg0)
	return *ret0, err
}

// GetBytes64 is a free data retrieval call binding the contract method 0x9139fa37.
//
// Solidity: function getBytes64(string ) constant returns(bytes1[64])
func (_Entry *EntrySession) GetBytes64(arg0 string) ([64][1]byte, error) {
	return _Entry.Contract.GetBytes64(&_Entry.CallOpts, arg0)
}

// GetBytes64 is a free data retrieval call binding the contract method 0x9139fa37.
//
// Solidity: function getBytes64(string ) constant returns(bytes1[64])
func (_Entry *EntryCallerSession) GetBytes64(arg0 string) ([64][1]byte, error) {
	return _Entry.Contract.GetBytes64(&_Entry.CallOpts, arg0)
}

// GetInt is a free data retrieval call binding the contract method 0x4900862e.
//
// Solidity: function getInt(string ) constant returns(int256)
func (_Entry *EntryCaller) GetInt(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Entry.contract.Call(opts, out, "getInt", arg0)
	return *ret0, err
}

// GetInt is a free data retrieval call binding the contract method 0x4900862e.
//
// Solidity: function getInt(string ) constant returns(int256)
func (_Entry *EntrySession) GetInt(arg0 string) (*big.Int, error) {
	return _Entry.Contract.GetInt(&_Entry.CallOpts, arg0)
}

// GetInt is a free data retrieval call binding the contract method 0x4900862e.
//
// Solidity: function getInt(string ) constant returns(int256)
func (_Entry *EntryCallerSession) GetInt(arg0 string) (*big.Int, error) {
	return _Entry.Contract.GetInt(&_Entry.CallOpts, arg0)
}

// GetString is a free data retrieval call binding the contract method 0x9bca41e8.
//
// Solidity: function getString(string ) constant returns(string)
func (_Entry *EntryCaller) GetString(opts *bind.CallOpts, arg0 string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Entry.contract.Call(opts, out, "getString", arg0)
	return *ret0, err
}

// GetString is a free data retrieval call binding the contract method 0x9bca41e8.
//
// Solidity: function getString(string ) constant returns(string)
func (_Entry *EntrySession) GetString(arg0 string) (string, error) {
	return _Entry.Contract.GetString(&_Entry.CallOpts, arg0)
}

// GetString is a free data retrieval call binding the contract method 0x9bca41e8.
//
// Solidity: function getString(string ) constant returns(string)
func (_Entry *EntryCallerSession) GetString(arg0 string) (string, error) {
	return _Entry.Contract.GetString(&_Entry.CallOpts, arg0)
}

// GetUInt is a free data retrieval call binding the contract method 0xdf7427af.
//
// Solidity: function getUInt(string ) constant returns(int256)
func (_Entry *EntryCaller) GetUInt(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Entry.contract.Call(opts, out, "getUInt", arg0)
	return *ret0, err
}

// GetUInt is a free data retrieval call binding the contract method 0xdf7427af.
//
// Solidity: function getUInt(string ) constant returns(int256)
func (_Entry *EntrySession) GetUInt(arg0 string) (*big.Int, error) {
	return _Entry.Contract.GetUInt(&_Entry.CallOpts, arg0)
}

// GetUInt is a free data retrieval call binding the contract method 0xdf7427af.
//
// Solidity: function getUInt(string ) constant returns(int256)
func (_Entry *EntryCallerSession) GetUInt(arg0 string) (*big.Int, error) {
	return _Entry.Contract.GetUInt(&_Entry.CallOpts, arg0)
}

// Set is a paid mutator transaction binding the contract method 0x1a391cb4.
//
// Solidity: function set(string , string ) returns()
func (_Entry *EntryTransactor) Set(opts *bind.TransactOpts, arg0 string, arg1 string) (*types.Transaction, error) {
	return _Entry.contract.Transact(opts, "set", arg0, arg1)
}

// Set is a paid mutator transaction binding the contract method 0x1a391cb4.
//
// Solidity: function set(string , string ) returns()
func (_Entry *EntrySession) Set(arg0 string, arg1 string) (*types.Transaction, error) {
	return _Entry.Contract.Set(&_Entry.TransactOpts, arg0, arg1)
}

// Set is a paid mutator transaction binding the contract method 0x1a391cb4.
//
// Solidity: function set(string , string ) returns()
func (_Entry *EntryTransactorSession) Set(arg0 string, arg1 string) (*types.Transaction, error) {
	return _Entry.Contract.Set(&_Entry.TransactOpts, arg0, arg1)
}

// Set0 is a paid mutator transaction binding the contract method 0x517c4dd9.
//
// Solidity: function set(string , address ) returns()
func (_Entry *EntryTransactor) Set0(opts *bind.TransactOpts, arg0 string, arg1 common.Address) (*types.Transaction, error) {
	return _Entry.contract.Transact(opts, "set0", arg0, arg1)
}

// Set0 is a paid mutator transaction binding the contract method 0x517c4dd9.
//
// Solidity: function set(string , address ) returns()
func (_Entry *EntrySession) Set0(arg0 string, arg1 common.Address) (*types.Transaction, error) {
	return _Entry.Contract.Set0(&_Entry.TransactOpts, arg0, arg1)
}

// Set0 is a paid mutator transaction binding the contract method 0x517c4dd9.
//
// Solidity: function set(string , address ) returns()
func (_Entry *EntryTransactorSession) Set0(arg0 string, arg1 common.Address) (*types.Transaction, error) {
	return _Entry.Contract.Set0(&_Entry.TransactOpts, arg0, arg1)
}

// Set1 is a paid mutator transaction binding the contract method 0xdef42698.
//
// Solidity: function set(string , int256 ) returns()
func (_Entry *EntryTransactor) Set1(opts *bind.TransactOpts, arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Entry.contract.Transact(opts, "set1", arg0, arg1)
}

// Set1 is a paid mutator transaction binding the contract method 0xdef42698.
//
// Solidity: function set(string , int256 ) returns()
func (_Entry *EntrySession) Set1(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Entry.Contract.Set1(&_Entry.TransactOpts, arg0, arg1)
}

// Set1 is a paid mutator transaction binding the contract method 0xdef42698.
//
// Solidity: function set(string , int256 ) returns()
func (_Entry *EntryTransactorSession) Set1(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Entry.Contract.Set1(&_Entry.TransactOpts, arg0, arg1)
}

// Set2 is a paid mutator transaction binding the contract method 0xf2f4ee6d.
//
// Solidity: function set(string , uint256 ) returns()
func (_Entry *EntryTransactor) Set2(opts *bind.TransactOpts, arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Entry.contract.Transact(opts, "set2", arg0, arg1)
}

// Set2 is a paid mutator transaction binding the contract method 0xf2f4ee6d.
//
// Solidity: function set(string , uint256 ) returns()
func (_Entry *EntrySession) Set2(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Entry.Contract.Set2(&_Entry.TransactOpts, arg0, arg1)
}

// Set2 is a paid mutator transaction binding the contract method 0xf2f4ee6d.
//
// Solidity: function set(string , uint256 ) returns()
func (_Entry *EntryTransactorSession) Set2(arg0 string, arg1 *big.Int) (*types.Transaction, error) {
	return _Entry.Contract.Set2(&_Entry.TransactOpts, arg0, arg1)
}

// KVTableABI is the input ABI used to generate the binding from.
const KVTableABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"set\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newEntry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// KVTableFuncSigs maps the 4-byte function signature to its string representation.
var KVTableFuncSigs = map[string]string{
	"7b1b8e03": "get(string)",
	"5887ab24": "newEntry()",
	"517c4dd9": "set(string,address)",
}

// KVTable is an auto generated Go binding around a Solidity contract.
type KVTable struct {
	KVTableCaller     // Read-only binding to the contract
	KVTableTransactor // Write-only binding to the contract
	KVTableFilterer   // Log filterer for contract events
}

// KVTableCaller is an auto generated read-only Go binding around a Solidity contract.
type KVTableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KVTableTransactor is an auto generated write-only Go binding around a Solidity contract.
type KVTableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KVTableFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type KVTableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KVTableSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type KVTableSession struct {
	Contract     *KVTable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KVTableCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type KVTableCallerSession struct {
	Contract *KVTableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// KVTableTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type KVTableTransactorSession struct {
	Contract     *KVTableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// KVTableRaw is an auto generated low-level Go binding around a Solidity contract.
type KVTableRaw struct {
	Contract *KVTable // Generic contract binding to access the raw methods on
}

// KVTableCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type KVTableCallerRaw struct {
	Contract *KVTableCaller // Generic read-only contract binding to access the raw methods on
}

// KVTableTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type KVTableTransactorRaw struct {
	Contract *KVTableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKVTable creates a new instance of KVTable, bound to a specific deployed contract.
func NewKVTable(address common.Address, backend bind.ContractBackend) (*KVTable, error) {
	contract, err := bindKVTable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KVTable{KVTableCaller: KVTableCaller{contract: contract}, KVTableTransactor: KVTableTransactor{contract: contract}, KVTableFilterer: KVTableFilterer{contract: contract}}, nil
}

// NewKVTableCaller creates a new read-only instance of KVTable, bound to a specific deployed contract.
func NewKVTableCaller(address common.Address, caller bind.ContractCaller) (*KVTableCaller, error) {
	contract, err := bindKVTable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KVTableCaller{contract: contract}, nil
}

// NewKVTableTransactor creates a new write-only instance of KVTable, bound to a specific deployed contract.
func NewKVTableTransactor(address common.Address, transactor bind.ContractTransactor) (*KVTableTransactor, error) {
	contract, err := bindKVTable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KVTableTransactor{contract: contract}, nil
}

// NewKVTableFilterer creates a new log filterer instance of KVTable, bound to a specific deployed contract.
func NewKVTableFilterer(address common.Address, filterer bind.ContractFilterer) (*KVTableFilterer, error) {
	contract, err := bindKVTable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KVTableFilterer{contract: contract}, nil
}

// bindKVTable binds a generic wrapper to an already deployed contract.
func bindKVTable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KVTableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KVTable *KVTableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KVTable.Contract.KVTableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KVTable *KVTableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KVTable.Contract.KVTableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KVTable *KVTableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KVTable.Contract.KVTableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KVTable *KVTableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KVTable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KVTable *KVTableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KVTable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KVTable *KVTableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KVTable.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x7b1b8e03.
//
// Solidity: function get(string ) constant returns(bool, address)
func (_KVTable *KVTableCaller) Get(opts *bind.CallOpts, arg0 string) (bool, common.Address, error) {
	var (
		ret0 = new(bool)
		ret1 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _KVTable.contract.Call(opts, out, "get", arg0)
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x7b1b8e03.
//
// Solidity: function get(string ) constant returns(bool, address)
func (_KVTable *KVTableSession) Get(arg0 string) (bool, common.Address, error) {
	return _KVTable.Contract.Get(&_KVTable.CallOpts, arg0)
}

// Get is a free data retrieval call binding the contract method 0x7b1b8e03.
//
// Solidity: function get(string ) constant returns(bool, address)
func (_KVTable *KVTableCallerSession) Get(arg0 string) (bool, common.Address, error) {
	return _KVTable.Contract.Get(&_KVTable.CallOpts, arg0)
}

// NewEntry is a free data retrieval call binding the contract method 0x5887ab24.
//
// Solidity: function newEntry() constant returns(address)
func (_KVTable *KVTableCaller) NewEntry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KVTable.contract.Call(opts, out, "newEntry")
	return *ret0, err
}

// NewEntry is a free data retrieval call binding the contract method 0x5887ab24.
//
// Solidity: function newEntry() constant returns(address)
func (_KVTable *KVTableSession) NewEntry() (common.Address, error) {
	return _KVTable.Contract.NewEntry(&_KVTable.CallOpts)
}

// NewEntry is a free data retrieval call binding the contract method 0x5887ab24.
//
// Solidity: function newEntry() constant returns(address)
func (_KVTable *KVTableCallerSession) NewEntry() (common.Address, error) {
	return _KVTable.Contract.NewEntry(&_KVTable.CallOpts)
}

// Set is a paid mutator transaction binding the contract method 0x517c4dd9.
//
// Solidity: function set(string , address ) returns(int256)
func (_KVTable *KVTableTransactor) Set(opts *bind.TransactOpts, arg0 string, arg1 common.Address) (*types.Transaction, error) {
	return _KVTable.contract.Transact(opts, "set", arg0, arg1)
}

// Set is a paid mutator transaction binding the contract method 0x517c4dd9.
//
// Solidity: function set(string , address ) returns(int256)
func (_KVTable *KVTableSession) Set(arg0 string, arg1 common.Address) (*types.Transaction, error) {
	return _KVTable.Contract.Set(&_KVTable.TransactOpts, arg0, arg1)
}

// Set is a paid mutator transaction binding the contract method 0x517c4dd9.
//
// Solidity: function set(string , address ) returns(int256)
func (_KVTable *KVTableTransactorSession) Set(arg0 string, arg1 common.Address) (*types.Transaction, error) {
	return _KVTable.Contract.Set(&_KVTable.TransactOpts, arg0, arg1)
}

// KVTableFactoryABI is the input ABI used to generate the binding from.
const KVTableFactoryABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"string\"}],\"name\":\"openTable\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"name\":\"createTable\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// KVTableFactoryFuncSigs maps the 4-byte function signature to its string representation.
var KVTableFactoryFuncSigs = map[string]string{
	"c92a7801": "createTable(string,string,string)",
	"59a48b65": "openTable(string)",
}

// KVTableFactory is an auto generated Go binding around a Solidity contract.
type KVTableFactory struct {
	KVTableFactoryCaller     // Read-only binding to the contract
	KVTableFactoryTransactor // Write-only binding to the contract
	KVTableFactoryFilterer   // Log filterer for contract events
}

// KVTableFactoryCaller is an auto generated read-only Go binding around a Solidity contract.
type KVTableFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KVTableFactoryTransactor is an auto generated write-only Go binding around a Solidity contract.
type KVTableFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KVTableFactoryFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type KVTableFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KVTableFactorySession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type KVTableFactorySession struct {
	Contract     *KVTableFactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KVTableFactoryCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type KVTableFactoryCallerSession struct {
	Contract *KVTableFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// KVTableFactoryTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type KVTableFactoryTransactorSession struct {
	Contract     *KVTableFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// KVTableFactoryRaw is an auto generated low-level Go binding around a Solidity contract.
type KVTableFactoryRaw struct {
	Contract *KVTableFactory // Generic contract binding to access the raw methods on
}

// KVTableFactoryCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type KVTableFactoryCallerRaw struct {
	Contract *KVTableFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// KVTableFactoryTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type KVTableFactoryTransactorRaw struct {
	Contract *KVTableFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKVTableFactory creates a new instance of KVTableFactory, bound to a specific deployed contract.
func NewKVTableFactory(address common.Address, backend bind.ContractBackend) (*KVTableFactory, error) {
	contract, err := bindKVTableFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KVTableFactory{KVTableFactoryCaller: KVTableFactoryCaller{contract: contract}, KVTableFactoryTransactor: KVTableFactoryTransactor{contract: contract}, KVTableFactoryFilterer: KVTableFactoryFilterer{contract: contract}}, nil
}

// NewKVTableFactoryCaller creates a new read-only instance of KVTableFactory, bound to a specific deployed contract.
func NewKVTableFactoryCaller(address common.Address, caller bind.ContractCaller) (*KVTableFactoryCaller, error) {
	contract, err := bindKVTableFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KVTableFactoryCaller{contract: contract}, nil
}

// NewKVTableFactoryTransactor creates a new write-only instance of KVTableFactory, bound to a specific deployed contract.
func NewKVTableFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*KVTableFactoryTransactor, error) {
	contract, err := bindKVTableFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KVTableFactoryTransactor{contract: contract}, nil
}

// NewKVTableFactoryFilterer creates a new log filterer instance of KVTableFactory, bound to a specific deployed contract.
func NewKVTableFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*KVTableFactoryFilterer, error) {
	contract, err := bindKVTableFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KVTableFactoryFilterer{contract: contract}, nil
}

// bindKVTableFactory binds a generic wrapper to an already deployed contract.
func bindKVTableFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KVTableFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KVTableFactory *KVTableFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KVTableFactory.Contract.KVTableFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KVTableFactory *KVTableFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KVTableFactory.Contract.KVTableFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KVTableFactory *KVTableFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KVTableFactory.Contract.KVTableFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KVTableFactory *KVTableFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KVTableFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KVTableFactory *KVTableFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KVTableFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KVTableFactory *KVTableFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KVTableFactory.Contract.contract.Transact(opts, method, params...)
}

// OpenTable is a free data retrieval call binding the contract method 0x59a48b65.
//
// Solidity: function openTable(string ) constant returns(address)
func (_KVTableFactory *KVTableFactoryCaller) OpenTable(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KVTableFactory.contract.Call(opts, out, "openTable", arg0)
	return *ret0, err
}

// OpenTable is a free data retrieval call binding the contract method 0x59a48b65.
//
// Solidity: function openTable(string ) constant returns(address)
func (_KVTableFactory *KVTableFactorySession) OpenTable(arg0 string) (common.Address, error) {
	return _KVTableFactory.Contract.OpenTable(&_KVTableFactory.CallOpts, arg0)
}

// OpenTable is a free data retrieval call binding the contract method 0x59a48b65.
//
// Solidity: function openTable(string ) constant returns(address)
func (_KVTableFactory *KVTableFactoryCallerSession) OpenTable(arg0 string) (common.Address, error) {
	return _KVTableFactory.Contract.OpenTable(&_KVTableFactory.CallOpts, arg0)
}

// CreateTable is a paid mutator transaction binding the contract method 0xc92a7801.
//
// Solidity: function createTable(string , string , string ) returns(int256)
func (_KVTableFactory *KVTableFactoryTransactor) CreateTable(opts *bind.TransactOpts, arg0 string, arg1 string, arg2 string) (*types.Transaction, error) {
	return _KVTableFactory.contract.Transact(opts, "createTable", arg0, arg1, arg2)
}

// CreateTable is a paid mutator transaction binding the contract method 0xc92a7801.
//
// Solidity: function createTable(string , string , string ) returns(int256)
func (_KVTableFactory *KVTableFactorySession) CreateTable(arg0 string, arg1 string, arg2 string) (*types.Transaction, error) {
	return _KVTableFactory.Contract.CreateTable(&_KVTableFactory.TransactOpts, arg0, arg1, arg2)
}

// CreateTable is a paid mutator transaction binding the contract method 0xc92a7801.
//
// Solidity: function createTable(string , string , string ) returns(int256)
func (_KVTableFactory *KVTableFactoryTransactorSession) CreateTable(arg0 string, arg1 string, arg2 string) (*types.Transaction, error) {
	return _KVTableFactory.Contract.CreateTable(&_KVTableFactory.TransactOpts, arg0, arg1, arg2)
}

// TableABI is the input ABI used to generate the binding from.
const TableABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"remove\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"insert\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newEntry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"update\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newCondition\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"select\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TableFuncSigs maps the 4-byte function signature to its string representation.
var TableFuncSigs = map[string]string{
	"4c6f30c0": "insert(string,address)",
	"c74f8caf": "newCondition()",
	"5887ab24": "newEntry()",
	"09ff42f0": "remove(string,address)",
	"d8ac5957": "select(string,address)",
	"664b37d6": "update(string,address,address)",
}

// Table is an auto generated Go binding around a Solidity contract.
type Table struct {
	TableCaller     // Read-only binding to the contract
	TableTransactor // Write-only binding to the contract
	TableFilterer   // Log filterer for contract events
}

// TableCaller is an auto generated read-only Go binding around a Solidity contract.
type TableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TableTransactor is an auto generated write-only Go binding around a Solidity contract.
type TableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TableFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type TableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TableSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type TableSession struct {
	Contract     *Table            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TableCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type TableCallerSession struct {
	Contract *TableCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TableTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type TableTransactorSession struct {
	Contract     *TableTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TableRaw is an auto generated low-level Go binding around a Solidity contract.
type TableRaw struct {
	Contract *Table // Generic contract binding to access the raw methods on
}

// TableCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type TableCallerRaw struct {
	Contract *TableCaller // Generic read-only contract binding to access the raw methods on
}

// TableTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type TableTransactorRaw struct {
	Contract *TableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTable creates a new instance of Table, bound to a specific deployed contract.
func NewTable(address common.Address, backend bind.ContractBackend) (*Table, error) {
	contract, err := bindTable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Table{TableCaller: TableCaller{contract: contract}, TableTransactor: TableTransactor{contract: contract}, TableFilterer: TableFilterer{contract: contract}}, nil
}

// NewTableCaller creates a new read-only instance of Table, bound to a specific deployed contract.
func NewTableCaller(address common.Address, caller bind.ContractCaller) (*TableCaller, error) {
	contract, err := bindTable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TableCaller{contract: contract}, nil
}

// NewTableTransactor creates a new write-only instance of Table, bound to a specific deployed contract.
func NewTableTransactor(address common.Address, transactor bind.ContractTransactor) (*TableTransactor, error) {
	contract, err := bindTable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TableTransactor{contract: contract}, nil
}

// NewTableFilterer creates a new log filterer instance of Table, bound to a specific deployed contract.
func NewTableFilterer(address common.Address, filterer bind.ContractFilterer) (*TableFilterer, error) {
	contract, err := bindTable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TableFilterer{contract: contract}, nil
}

// bindTable binds a generic wrapper to an already deployed contract.
func bindTable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Table *TableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Table.Contract.TableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Table *TableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Table.Contract.TableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Table *TableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Table.Contract.TableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Table *TableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Table.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Table *TableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Table.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Table *TableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Table.Contract.contract.Transact(opts, method, params...)
}

// NewCondition is a free data retrieval call binding the contract method 0xc74f8caf.
//
// Solidity: function newCondition() constant returns(address)
func (_Table *TableCaller) NewCondition(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Table.contract.Call(opts, out, "newCondition")
	return *ret0, err
}

// NewCondition is a free data retrieval call binding the contract method 0xc74f8caf.
//
// Solidity: function newCondition() constant returns(address)
func (_Table *TableSession) NewCondition() (common.Address, error) {
	return _Table.Contract.NewCondition(&_Table.CallOpts)
}

// NewCondition is a free data retrieval call binding the contract method 0xc74f8caf.
//
// Solidity: function newCondition() constant returns(address)
func (_Table *TableCallerSession) NewCondition() (common.Address, error) {
	return _Table.Contract.NewCondition(&_Table.CallOpts)
}

// NewEntry is a free data retrieval call binding the contract method 0x5887ab24.
//
// Solidity: function newEntry() constant returns(address)
func (_Table *TableCaller) NewEntry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Table.contract.Call(opts, out, "newEntry")
	return *ret0, err
}

// NewEntry is a free data retrieval call binding the contract method 0x5887ab24.
//
// Solidity: function newEntry() constant returns(address)
func (_Table *TableSession) NewEntry() (common.Address, error) {
	return _Table.Contract.NewEntry(&_Table.CallOpts)
}

// NewEntry is a free data retrieval call binding the contract method 0x5887ab24.
//
// Solidity: function newEntry() constant returns(address)
func (_Table *TableCallerSession) NewEntry() (common.Address, error) {
	return _Table.Contract.NewEntry(&_Table.CallOpts)
}

// Select is a free data retrieval call binding the contract method 0xd8ac5957.
//
// Solidity: function select(string , address ) constant returns(address)
func (_Table *TableCaller) Select(opts *bind.CallOpts, arg0 string, arg1 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Table.contract.Call(opts, out, "select", arg0, arg1)
	return *ret0, err
}

// Select is a free data retrieval call binding the contract method 0xd8ac5957.
//
// Solidity: function select(string , address ) constant returns(address)
func (_Table *TableSession) Select(arg0 string, arg1 common.Address) (common.Address, error) {
	return _Table.Contract.Select(&_Table.CallOpts, arg0, arg1)
}

// Select is a free data retrieval call binding the contract method 0xd8ac5957.
//
// Solidity: function select(string , address ) constant returns(address)
func (_Table *TableCallerSession) Select(arg0 string, arg1 common.Address) (common.Address, error) {
	return _Table.Contract.Select(&_Table.CallOpts, arg0, arg1)
}

// Insert is a paid mutator transaction binding the contract method 0x4c6f30c0.
//
// Solidity: function insert(string , address ) returns(int256)
func (_Table *TableTransactor) Insert(opts *bind.TransactOpts, arg0 string, arg1 common.Address) (*types.Transaction, error) {
	return _Table.contract.Transact(opts, "insert", arg0, arg1)
}

// Insert is a paid mutator transaction binding the contract method 0x4c6f30c0.
//
// Solidity: function insert(string , address ) returns(int256)
func (_Table *TableSession) Insert(arg0 string, arg1 common.Address) (*types.Transaction, error) {
	return _Table.Contract.Insert(&_Table.TransactOpts, arg0, arg1)
}

// Insert is a paid mutator transaction binding the contract method 0x4c6f30c0.
//
// Solidity: function insert(string , address ) returns(int256)
func (_Table *TableTransactorSession) Insert(arg0 string, arg1 common.Address) (*types.Transaction, error) {
	return _Table.Contract.Insert(&_Table.TransactOpts, arg0, arg1)
}

// Remove is a paid mutator transaction binding the contract method 0x09ff42f0.
//
// Solidity: function remove(string , address ) returns(int256)
func (_Table *TableTransactor) Remove(opts *bind.TransactOpts, arg0 string, arg1 common.Address) (*types.Transaction, error) {
	return _Table.contract.Transact(opts, "remove", arg0, arg1)
}

// Remove is a paid mutator transaction binding the contract method 0x09ff42f0.
//
// Solidity: function remove(string , address ) returns(int256)
func (_Table *TableSession) Remove(arg0 string, arg1 common.Address) (*types.Transaction, error) {
	return _Table.Contract.Remove(&_Table.TransactOpts, arg0, arg1)
}

// Remove is a paid mutator transaction binding the contract method 0x09ff42f0.
//
// Solidity: function remove(string , address ) returns(int256)
func (_Table *TableTransactorSession) Remove(arg0 string, arg1 common.Address) (*types.Transaction, error) {
	return _Table.Contract.Remove(&_Table.TransactOpts, arg0, arg1)
}

// Update is a paid mutator transaction binding the contract method 0x664b37d6.
//
// Solidity: function update(string , address , address ) returns(int256)
func (_Table *TableTransactor) Update(opts *bind.TransactOpts, arg0 string, arg1 common.Address, arg2 common.Address) (*types.Transaction, error) {
	return _Table.contract.Transact(opts, "update", arg0, arg1, arg2)
}

// Update is a paid mutator transaction binding the contract method 0x664b37d6.
//
// Solidity: function update(string , address , address ) returns(int256)
func (_Table *TableSession) Update(arg0 string, arg1 common.Address, arg2 common.Address) (*types.Transaction, error) {
	return _Table.Contract.Update(&_Table.TransactOpts, arg0, arg1, arg2)
}

// Update is a paid mutator transaction binding the contract method 0x664b37d6.
//
// Solidity: function update(string , address , address ) returns(int256)
func (_Table *TableTransactorSession) Update(arg0 string, arg1 common.Address, arg2 common.Address) (*types.Transaction, error) {
	return _Table.Contract.Update(&_Table.TransactOpts, arg0, arg1, arg2)
}

// TableFactoryABI is the input ABI used to generate the binding from.
const TableFactoryABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"string\"}],\"name\":\"openTable\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"name\":\"createTable\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TableFactoryFuncSigs maps the 4-byte function signature to its string representation.
var TableFactoryFuncSigs = map[string]string{
	"c92a7801": "createTable(string,string,string)",
	"59a48b65": "openTable(string)",
}

// TableFactory is an auto generated Go binding around a Solidity contract.
type TableFactory struct {
	TableFactoryCaller     // Read-only binding to the contract
	TableFactoryTransactor // Write-only binding to the contract
	TableFactoryFilterer   // Log filterer for contract events
}

// TableFactoryCaller is an auto generated read-only Go binding around a Solidity contract.
type TableFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TableFactoryTransactor is an auto generated write-only Go binding around a Solidity contract.
type TableFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TableFactoryFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type TableFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TableFactorySession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type TableFactorySession struct {
	Contract     *TableFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TableFactoryCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type TableFactoryCallerSession struct {
	Contract *TableFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TableFactoryTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type TableFactoryTransactorSession struct {
	Contract     *TableFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TableFactoryRaw is an auto generated low-level Go binding around a Solidity contract.
type TableFactoryRaw struct {
	Contract *TableFactory // Generic contract binding to access the raw methods on
}

// TableFactoryCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type TableFactoryCallerRaw struct {
	Contract *TableFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// TableFactoryTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type TableFactoryTransactorRaw struct {
	Contract *TableFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTableFactory creates a new instance of TableFactory, bound to a specific deployed contract.
func NewTableFactory(address common.Address, backend bind.ContractBackend) (*TableFactory, error) {
	contract, err := bindTableFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TableFactory{TableFactoryCaller: TableFactoryCaller{contract: contract}, TableFactoryTransactor: TableFactoryTransactor{contract: contract}, TableFactoryFilterer: TableFactoryFilterer{contract: contract}}, nil
}

// NewTableFactoryCaller creates a new read-only instance of TableFactory, bound to a specific deployed contract.
func NewTableFactoryCaller(address common.Address, caller bind.ContractCaller) (*TableFactoryCaller, error) {
	contract, err := bindTableFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TableFactoryCaller{contract: contract}, nil
}

// NewTableFactoryTransactor creates a new write-only instance of TableFactory, bound to a specific deployed contract.
func NewTableFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*TableFactoryTransactor, error) {
	contract, err := bindTableFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TableFactoryTransactor{contract: contract}, nil
}

// NewTableFactoryFilterer creates a new log filterer instance of TableFactory, bound to a specific deployed contract.
func NewTableFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*TableFactoryFilterer, error) {
	contract, err := bindTableFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TableFactoryFilterer{contract: contract}, nil
}

// bindTableFactory binds a generic wrapper to an already deployed contract.
func bindTableFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TableFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TableFactory *TableFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TableFactory.Contract.TableFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TableFactory *TableFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TableFactory.Contract.TableFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TableFactory *TableFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TableFactory.Contract.TableFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TableFactory *TableFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TableFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TableFactory *TableFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TableFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TableFactory *TableFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TableFactory.Contract.contract.Transact(opts, method, params...)
}

// OpenTable is a free data retrieval call binding the contract method 0x59a48b65.
//
// Solidity: function openTable(string ) constant returns(address)
func (_TableFactory *TableFactoryCaller) OpenTable(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TableFactory.contract.Call(opts, out, "openTable", arg0)
	return *ret0, err
}

// OpenTable is a free data retrieval call binding the contract method 0x59a48b65.
//
// Solidity: function openTable(string ) constant returns(address)
func (_TableFactory *TableFactorySession) OpenTable(arg0 string) (common.Address, error) {
	return _TableFactory.Contract.OpenTable(&_TableFactory.CallOpts, arg0)
}

// OpenTable is a free data retrieval call binding the contract method 0x59a48b65.
//
// Solidity: function openTable(string ) constant returns(address)
func (_TableFactory *TableFactoryCallerSession) OpenTable(arg0 string) (common.Address, error) {
	return _TableFactory.Contract.OpenTable(&_TableFactory.CallOpts, arg0)
}

// CreateTable is a paid mutator transaction binding the contract method 0xc92a7801.
//
// Solidity: function createTable(string , string , string ) returns(int256)
func (_TableFactory *TableFactoryTransactor) CreateTable(opts *bind.TransactOpts, arg0 string, arg1 string, arg2 string) (*types.Transaction, error) {
	return _TableFactory.contract.Transact(opts, "createTable", arg0, arg1, arg2)
}

// CreateTable is a paid mutator transaction binding the contract method 0xc92a7801.
//
// Solidity: function createTable(string , string , string ) returns(int256)
func (_TableFactory *TableFactorySession) CreateTable(arg0 string, arg1 string, arg2 string) (*types.Transaction, error) {
	return _TableFactory.Contract.CreateTable(&_TableFactory.TransactOpts, arg0, arg1, arg2)
}

// CreateTable is a paid mutator transaction binding the contract method 0xc92a7801.
//
// Solidity: function createTable(string , string , string ) returns(int256)
func (_TableFactory *TableFactoryTransactorSession) CreateTable(arg0 string, arg1 string, arg2 string) (*types.Transaction, error) {
	return _TableFactory.Contract.CreateTable(&_TableFactory.TransactOpts, arg0, arg1, arg2)
}
