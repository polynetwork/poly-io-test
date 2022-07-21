// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EthCrossChainData

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

// EthCrossChainDataMetaData contains all meta data concerning the EthCrossChainData contract.
var EthCrossChainDataMetaData = &core.CtrMetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"ConKeepersPkBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CurEpochStartHeight\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EthToPolyTxHashIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"EthToPolyTxHashMap\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ExtraData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"fromChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"fromChainTx\",\"type\":\"bytes32\"}],\"name\":\"checkIfFromChainTxExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurEpochConPubKeyBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurEpochStartHeight\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ethTxHashIndex\",\"type\":\"uint256\"}],\"name\":\"getEthTxHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEthTxHashIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"key2\",\"type\":\"bytes32\"}],\"name\":\"getExtraData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"fromChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"fromChainTx\",\"type\":\"bytes32\"}],\"name\":\"markFromChainTxExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"curEpochPkBytes\",\"type\":\"bytes\"}],\"name\":\"putCurEpochConPubKeyBytes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"curEpochStartHeight\",\"type\":\"uint32\"}],\"name\":\"putCurEpochStartHeight\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ethTxHash\",\"type\":\"bytes32\"}],\"name\":\"putEthTxHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"key2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"putExtraData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260006100146100d160201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35060008060146101000a81548160ff0219169083151502179055506100d9565b600033905090565b611b59806100e86000396000f3fe608060405234801561001057600080fd5b506004361061014c5760003560e01c80635c975abb116100c35780638da5cb5b1161007c5780638da5cb5b146107425780638f32d59b1461078c578063e90bfdcf146107ae578063f2fde38b14610808578063fcbdc1e01461084c578063ff3d24a1146108cf5761014c565b80635c975abb146105fb57806367e31a741461061d57806369d4807414610647578063715018a6146106ca5780638456cb59146106d45780638a8bd17f146106f65761014c565b80633f4ba83a116101155780633f4ba83a146103a357806340602bb5146103c557806341973cd9146104765780634c3ccf6414610549578063529caad81461058f5780635ac40790146105d15761014c565b8062c5fff8146101515780630586763c1461016f5780631afe374e146101c957806320bbde38146102b05780632992787514610361575b600080fd5b6101596108ed565b6040518082815260200191505060405180910390f35b6101af6004803603604081101561018557600080fd5b81019080803567ffffffffffffffff169060200190929190803590602001909291905050506108f3565b604051808215151515815260200191505060405180910390f35b610296600480360360608110156101df57600080fd5b8101908080359060200190929190803590602001909291908035906020019064010000000081111561021057600080fd5b82018360208201111561022257600080fd5b8035906020019184600183028401116401000000008311171561024457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610943565b604051808215151515815260200191505060405180910390f35b6102e6600480360360408110156102c657600080fd5b810190808035906020019092919080359060200190929190505050610a85565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561032657808201518184015260208101905061030b565b50505050905090810190601f1680156103535780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61038d6004803603602081101561037757600080fd5b8101908080359060200190929190505050610b42565b6040518082815260200191505060405180910390f35b6103ab610b5f565b604051808215151515815260200191505060405180910390f35b6103fb600480360360408110156103db57600080fd5b810190808035906020019092919080359060200190929190505050610c6c565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561043b578082015181840152602081019050610420565b50505050905090810190601f1680156104685780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61052f6004803603602081101561048c57600080fd5b81019080803590602001906401000000008111156104a957600080fd5b8201836020820111156104bb57600080fd5b803590602001918460018302840111640100000000831117156104dd57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610d33565b604051808215151515815260200191505060405180910390f35b6105756004803603602081101561055f57600080fd5b8101908080359060200190929190505050610e51565b604051808215151515815260200191505060405180910390f35b6105bb600480360360208110156105a557600080fd5b8101908080359060200190929190505050610f7e565b6040518082815260200191505060405180910390f35b6105d9610f96565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b610603610fb0565b604051808215151515815260200191505060405180910390f35b610625610fc6565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b61064f610fdc565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561068f578082015181840152602081019050610674565b50505050905090810190601f1680156106bc5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6106d261107e565b005b6106dc6111b7565b604051808215151515815260200191505060405180910390f35b6107286004803603602081101561070c57600080fd5b81019080803563ffffffff1690602001909291905050506112c5565b604051808215151515815260200191505060405180910390f35b61074a6113ed565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610794611416565b604051808215151515815260200191505060405180910390f35b6107ee600480360360408110156107c457600080fd5b81019080803567ffffffffffffffff16906020019092919080359060200190929190505050611474565b604051808215151515815260200191505060405180910390f35b61084a6004803603602081101561081e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506115cd565b005b610854611653565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610894578082015181840152602081019050610879565b50505050905090810190601f1680156108c15780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6108d76116f1565b6040518082815260200191505060405180910390f35b60025481565b6000600560008467ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020600083815260200190815260200160002060009054906101000a900460ff16905092915050565b60008060149054906101000a900460ff16156109c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b6109cf611416565b610a41576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b816006600086815260200190815260200160002060008581526020019081526020016000209080519060200190610a79929190611a59565b50600190509392505050565b6006602052816000526040600020602052806000526040600020600091509150508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b3a5780601f10610b0f57610100808354040283529160200191610b3a565b820191906000526020600020905b815481529060010190602001808311610b1d57829003601f168201915b505050505081565b600060016000838152602001908152602001600020549050919050565b6000610b69611416565b610bdb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600060149054906101000a900460ff16610c5d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f5061757361626c653a206e6f742070617573656400000000000000000000000081525060200191505060405180910390fd5b610c656116fb565b6001905090565b60606006600084815260200190815260200160002060008381526020019081526020016000208054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d265780601f10610cfb57610100808354040283529160200191610d26565b820191906000526020600020905b815481529060010190602001808311610d0957829003601f168201915b5050505050905092915050565b60008060149054906101000a900460ff1615610db7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b610dbf611416565b610e31576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b8160039080519060200190610e47929190611a59565b5060019050919050565b60008060149054906101000a900460ff1615610ed5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b610edd611416565b610f4f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b816001600060025481526020019081526020016000208190555060016002540160028190555060019050919050565b60016020528060005260406000206000915090505481565b6000600460009054906101000a900463ffffffff16905090565b60008060149054906101000a900460ff16905090565b600460009054906101000a900463ffffffff1681565b606060038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110745780601f1061104957610100808354040283529160200191611074565b820191906000526020600020905b81548152906001019060200180831161105757829003601f168201915b5050505050905090565b611086611416565b6110f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60006111c1611416565b611233576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600060149054906101000a900460ff16156112b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b6112be611803565b6001905090565b60008060149054906101000a900460ff1615611349576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b611351611416565b6113c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b81600460006101000a81548163ffffffff021916908363ffffffff16021790555060019050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661145861190d565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b60008060149054906101000a900460ff16156114f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b611500611416565b611572576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6001600560008567ffffffffffffffff1667ffffffffffffffff168152602001908152602001600020600084815260200190815260200160002060006101000a81548160ff0219169083151502179055506001905092915050565b6115d5611416565b611647576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b61165081611915565b50565b60038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156116e95780601f106116be576101008083540402835291602001916116e9565b820191906000526020600020905b8154815290600101906020018083116116cc57829003601f168201915b505050505081565b6000600254905090565b600060149054906101000a900460ff1661177d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f5061757361626c653a206e6f742070617573656400000000000000000000000081525060200191505060405180910390fd5b60008060146101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6117c061190d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1565b600060149054906101000a900460ff1615611886576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260108152602001807f5061757361626c653a207061757365640000000000000000000000000000000081525060200191505060405180910390fd5b6001600060146101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586118ca61190d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561199b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180611aff6026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611a9a57805160ff1916838001178555611ac8565b82800160010185558215611ac8579182015b82811115611ac7578251825591602001919060010190611aac565b5b509050611ad59190611ad9565b5090565b611afb91905b80821115611af7576000816000905550600101611adf565b5090565b9056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373a265627a7a72315820abd2ff3f9fab911cceb20159f1034c6ad22cdcaec82eb7fbef2a40b6536ca43a64736f6c63430005110032",
}

// EthCrossChainDataABI is the input ABI used to generate the binding from.
// Deprecated: Use EthCrossChainDataMetaData.ABI instead.
var EthCrossChainDataABI = EthCrossChainDataMetaData.ABI

// EthCrossChainDataBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthCrossChainDataMetaData.Bin instead.
var EthCrossChainDataBin = EthCrossChainDataMetaData.Bin

// DeployEthCrossChainData deploys a new ChainSQL contract, binding an instance of EthCrossChainData to it.
func DeployEthCrossChainData(chainsql *core.Chainsql, auth *core.TransactOpts) (*core.DeployTxRet, *EthCrossChainData, error) {
	parsed, err := EthCrossChainDataMetaData.GetAbi()
	if err != nil {
		return &core.DeployTxRet{}, nil, err
	}
	if parsed == nil {
		return &core.DeployTxRet{}, nil, errors.New("GetABI returned nil")
	}

	deployRet, contract, err := core.DeployContract(chainsql, auth, *parsed, common.FromHex(EthCrossChainDataBin))
	if err != nil {
		return &core.DeployTxRet{}, nil, err
	}
	return deployRet, &EthCrossChainData{EthCrossChainDataCaller: EthCrossChainDataCaller{contract: contract}, EthCrossChainDataTransactor: EthCrossChainDataTransactor{contract: contract}, EthCrossChainDataFilterer: EthCrossChainDataFilterer{contract: contract}}, nil
}

// EthCrossChainData is an auto generated Go binding around an ChainSQL contract.
type EthCrossChainData struct {
	EthCrossChainDataCaller     // Read-only binding to the contract
	EthCrossChainDataTransactor // Write-only binding to the contract
	EthCrossChainDataFilterer   // Log filterer for contract events
}

// EthCrossChainDataCaller is an auto generated read-only Go binding around an ChainSQL contract.
type EthCrossChainDataCaller struct {
	contract *core.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainDataTransactor is an auto generated write-only Go binding around an ChainSQL contract.
type EthCrossChainDataTransactor struct {
	contract *core.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainDataFilterer is an auto generated log filtering Go binding around an ChainSQL contract events.
type EthCrossChainDataFilterer struct {
	contract *core.BoundContract // Generic contract wrapper for the low level calls
}

// EthCrossChainDataSession is an auto generated Go binding around an ChainSQL contract,
// with pre-set call and transact options.
type EthCrossChainDataSession struct {
	Contract     *EthCrossChainData // Generic contract binding to set the session for
	CallOpts     core.CallOpts      // Call options to use throughout this session
	TransactOpts core.TransactOpts  // Transaction auth options to use throughout this session
}

// EthCrossChainDataCallerSession is an auto generated read-only Go binding around an ChainSQL contract,
// with pre-set call options.
type EthCrossChainDataCallerSession struct {
	Contract *EthCrossChainDataCaller // Generic contract caller binding to set the session for
	CallOpts core.CallOpts            // Call options to use throughout this session
}

// EthCrossChainDataTransactorSession is an auto generated write-only Go binding around an ChainSQL contract,
// with pre-set transact options.
type EthCrossChainDataTransactorSession struct {
	Contract     *EthCrossChainDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts core.TransactOpts            // Transaction auth options to use throughout this session
}

// EthCrossChainDataRaw is an auto generated low-level Go binding around an ChainSQL contract.
type EthCrossChainDataRaw struct {
	Contract *EthCrossChainData // Generic contract binding to access the raw methods on
}

// EthCrossChainDataCallerRaw is an auto generated low-level read-only Go binding around an ChainSQL contract.
type EthCrossChainDataCallerRaw struct {
	Contract *EthCrossChainDataCaller // Generic read-only contract binding to access the raw methods on
}

// EthCrossChainDataTransactorRaw is an auto generated low-level write-only Go binding around an ChainSQL contract.
type EthCrossChainDataTransactorRaw struct {
	Contract *EthCrossChainDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthCrossChainData creates a new instance of EthCrossChainData, bound to a specific deployed contract.
func NewEthCrossChainData(chainsql *core.Chainsql, address string) (*EthCrossChainData, error) {
	contract, err := bindEthCrossChainData(chainsql, address)
	if err != nil {
		return nil, err
	}
	return &EthCrossChainData{EthCrossChainDataCaller: EthCrossChainDataCaller{contract: contract}, EthCrossChainDataTransactor: EthCrossChainDataTransactor{contract: contract}, EthCrossChainDataFilterer: EthCrossChainDataFilterer{contract: contract}}, nil
}

// // NewEthCrossChainDataCaller creates a new read-only instance of EthCrossChainData, bound to a specific deployed contract.
// func NewEthCrossChainDataCaller(address common.Address, caller bind.ContractCaller) (*EthCrossChainDataCaller, error) {
//   contract, err := bindEthCrossChainData(address, caller, nil, nil)
//   if err != nil {
//     return nil, err
//   }
//   return &EthCrossChainDataCaller{contract: contract}, nil
// }

// // NewEthCrossChainDataTransactor creates a new write-only instance of EthCrossChainData, bound to a specific deployed contract.
// func NewEthCrossChainDataTransactor(address common.Address, transactor bind.ContractTransactor) (*EthCrossChainDataTransactor, error) {
//   contract, err := bindEthCrossChainData(address, nil, transactor, nil)
//   if err != nil {
//     return nil, err
//   }
//   return &EthCrossChainDataTransactor{contract: contract}, nil
// }

// // NewEthCrossChainDataFilterer creates a new log filterer instance of EthCrossChainData, bound to a specific deployed contract.
// func NewEthCrossChainDataFilterer(address common.Address, filterer bind.ContractFilterer) (*EthCrossChainDataFilterer, error) {
//   contract, err := bindEthCrossChainData(address, nil, nil, filterer)
//   if err != nil {
//     return nil, err
//   }
//   return &EthCrossChainDataFilterer{contract: contract}, nil
// }

// bindEthCrossChainData binds a generic wrapper to an already deployed contract.
func bindEthCrossChainData(chainsql *core.Chainsql, address string) (*core.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthCrossChainDataABI))
	if err != nil {
		return nil, err
	}
	return core.NewBoundContract(chainsql, address, parsed), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
// func (_EthCrossChainData *EthCrossChainDataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
// 	return _EthCrossChainData.Contract.EthCrossChainDataCaller.contract.Call(opts, result, method, params...)
// }

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
// func (_EthCrossChainData *EthCrossChainDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
// 	return _EthCrossChainData.Contract.EthCrossChainDataTransactor.contract.Transfer(opts)
// }

// Transact invokes the (paid) contract method with params as input values.
// func (_EthCrossChainData *EthCrossChainDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
// 	return _EthCrossChainData.Contract.EthCrossChainDataTransactor.contract.Transact(opts, method, params...)
// }

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
// func (_EthCrossChainData *EthCrossChainDataCallerRaw) Call(opts *core.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
// 	return _EthCrossChainData.Contract.contract.Call(opts, result, method, params...)
// }

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
// func (_EthCrossChainData *EthCrossChainDataTransactorRaw) Transfer(opts *core.TransactOpts) (*types.Transaction, error) {
// 	return _EthCrossChainData.Contract.contract.Transfer(opts)
// }

// Transact invokes the (paid) contract method with params as input values.
// func (_EthCrossChainData *EthCrossChainDataTransactorRaw) Transact(opts *core.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
// 	return _EthCrossChainData.Contract.contract.Transact(opts, method, params...)
// }

// ConKeepersPkBytes is a free data retrieval call binding the contract method 0xfcbdc1e0.
//
// Solidity: function ConKeepersPkBytes() view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCaller) ConKeepersPkBytes(opts *core.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "ConKeepersPkBytes")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ConKeepersPkBytes is a free data retrieval call binding the contract method 0xfcbdc1e0.
//
// Solidity: function ConKeepersPkBytes() view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataSession) ConKeepersPkBytes() ([]byte, error) {
	return _EthCrossChainData.Contract.ConKeepersPkBytes(&_EthCrossChainData.CallOpts)
}

// ConKeepersPkBytes is a free data retrieval call binding the contract method 0xfcbdc1e0.
//
// Solidity: function ConKeepersPkBytes() view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCallerSession) ConKeepersPkBytes() ([]byte, error) {
	return _EthCrossChainData.Contract.ConKeepersPkBytes(&_EthCrossChainData.CallOpts)
}

// CurEpochStartHeight is a free data retrieval call binding the contract method 0x67e31a74.
//
// Solidity: function CurEpochStartHeight() view returns(uint32)
func (_EthCrossChainData *EthCrossChainDataCaller) CurEpochStartHeight(opts *core.CallOpts) (uint32, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "CurEpochStartHeight")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CurEpochStartHeight is a free data retrieval call binding the contract method 0x67e31a74.
//
// Solidity: function CurEpochStartHeight() view returns(uint32)
func (_EthCrossChainData *EthCrossChainDataSession) CurEpochStartHeight() (uint32, error) {
	return _EthCrossChainData.Contract.CurEpochStartHeight(&_EthCrossChainData.CallOpts)
}

// CurEpochStartHeight is a free data retrieval call binding the contract method 0x67e31a74.
//
// Solidity: function CurEpochStartHeight() view returns(uint32)
func (_EthCrossChainData *EthCrossChainDataCallerSession) CurEpochStartHeight() (uint32, error) {
	return _EthCrossChainData.Contract.CurEpochStartHeight(&_EthCrossChainData.CallOpts)
}

// EthToPolyTxHashIndex is a free data retrieval call binding the contract method 0x00c5fff8.
//
// Solidity: function EthToPolyTxHashIndex() view returns(uint256)
func (_EthCrossChainData *EthCrossChainDataCaller) EthToPolyTxHashIndex(opts *core.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "EthToPolyTxHashIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthToPolyTxHashIndex is a free data retrieval call binding the contract method 0x00c5fff8.
//
// Solidity: function EthToPolyTxHashIndex() view returns(uint256)
func (_EthCrossChainData *EthCrossChainDataSession) EthToPolyTxHashIndex() (*big.Int, error) {
	return _EthCrossChainData.Contract.EthToPolyTxHashIndex(&_EthCrossChainData.CallOpts)
}

// EthToPolyTxHashIndex is a free data retrieval call binding the contract method 0x00c5fff8.
//
// Solidity: function EthToPolyTxHashIndex() view returns(uint256)
func (_EthCrossChainData *EthCrossChainDataCallerSession) EthToPolyTxHashIndex() (*big.Int, error) {
	return _EthCrossChainData.Contract.EthToPolyTxHashIndex(&_EthCrossChainData.CallOpts)
}

// EthToPolyTxHashMap is a free data retrieval call binding the contract method 0x529caad8.
//
// Solidity: function EthToPolyTxHashMap(uint256 ) view returns(bytes32)
func (_EthCrossChainData *EthCrossChainDataCaller) EthToPolyTxHashMap(opts *core.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "EthToPolyTxHashMap", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EthToPolyTxHashMap is a free data retrieval call binding the contract method 0x529caad8.
//
// Solidity: function EthToPolyTxHashMap(uint256 ) view returns(bytes32)
func (_EthCrossChainData *EthCrossChainDataSession) EthToPolyTxHashMap(arg0 *big.Int) ([32]byte, error) {
	return _EthCrossChainData.Contract.EthToPolyTxHashMap(&_EthCrossChainData.CallOpts, arg0)
}

// EthToPolyTxHashMap is a free data retrieval call binding the contract method 0x529caad8.
//
// Solidity: function EthToPolyTxHashMap(uint256 ) view returns(bytes32)
func (_EthCrossChainData *EthCrossChainDataCallerSession) EthToPolyTxHashMap(arg0 *big.Int) ([32]byte, error) {
	return _EthCrossChainData.Contract.EthToPolyTxHashMap(&_EthCrossChainData.CallOpts, arg0)
}

// ExtraData is a free data retrieval call binding the contract method 0x20bbde38.
//
// Solidity: function ExtraData(bytes32 , bytes32 ) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCaller) ExtraData(opts *core.CallOpts, arg0 [32]byte, arg1 [32]byte) ([]byte, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "ExtraData", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ExtraData is a free data retrieval call binding the contract method 0x20bbde38.
//
// Solidity: function ExtraData(bytes32 , bytes32 ) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataSession) ExtraData(arg0 [32]byte, arg1 [32]byte) ([]byte, error) {
	return _EthCrossChainData.Contract.ExtraData(&_EthCrossChainData.CallOpts, arg0, arg1)
}

// ExtraData is a free data retrieval call binding the contract method 0x20bbde38.
//
// Solidity: function ExtraData(bytes32 , bytes32 ) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCallerSession) ExtraData(arg0 [32]byte, arg1 [32]byte) ([]byte, error) {
	return _EthCrossChainData.Contract.ExtraData(&_EthCrossChainData.CallOpts, arg0, arg1)
}

// CheckIfFromChainTxExist is a free data retrieval call binding the contract method 0x0586763c.
//
// Solidity: function checkIfFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) view returns(bool)
func (_EthCrossChainData *EthCrossChainDataCaller) CheckIfFromChainTxExist(opts *core.CallOpts, fromChainId uint64, fromChainTx [32]byte) (bool, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "checkIfFromChainTxExist", fromChainId, fromChainTx)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckIfFromChainTxExist is a free data retrieval call binding the contract method 0x0586763c.
//
// Solidity: function checkIfFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) view returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) CheckIfFromChainTxExist(fromChainId uint64, fromChainTx [32]byte) (bool, error) {
	return _EthCrossChainData.Contract.CheckIfFromChainTxExist(&_EthCrossChainData.CallOpts, fromChainId, fromChainTx)
}

// CheckIfFromChainTxExist is a free data retrieval call binding the contract method 0x0586763c.
//
// Solidity: function checkIfFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) view returns(bool)
func (_EthCrossChainData *EthCrossChainDataCallerSession) CheckIfFromChainTxExist(fromChainId uint64, fromChainTx [32]byte) (bool, error) {
	return _EthCrossChainData.Contract.CheckIfFromChainTxExist(&_EthCrossChainData.CallOpts, fromChainId, fromChainTx)
}

// GetCurEpochConPubKeyBytes is a free data retrieval call binding the contract method 0x69d48074.
//
// Solidity: function getCurEpochConPubKeyBytes() view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCaller) GetCurEpochConPubKeyBytes(opts *core.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "getCurEpochConPubKeyBytes")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetCurEpochConPubKeyBytes is a free data retrieval call binding the contract method 0x69d48074.
//
// Solidity: function getCurEpochConPubKeyBytes() view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataSession) GetCurEpochConPubKeyBytes() ([]byte, error) {
	return _EthCrossChainData.Contract.GetCurEpochConPubKeyBytes(&_EthCrossChainData.CallOpts)
}

// GetCurEpochConPubKeyBytes is a free data retrieval call binding the contract method 0x69d48074.
//
// Solidity: function getCurEpochConPubKeyBytes() view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCallerSession) GetCurEpochConPubKeyBytes() ([]byte, error) {
	return _EthCrossChainData.Contract.GetCurEpochConPubKeyBytes(&_EthCrossChainData.CallOpts)
}

// GetCurEpochStartHeight is a free data retrieval call binding the contract method 0x5ac40790.
//
// Solidity: function getCurEpochStartHeight() view returns(uint32)
func (_EthCrossChainData *EthCrossChainDataCaller) GetCurEpochStartHeight(opts *core.CallOpts) (uint32, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "getCurEpochStartHeight")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetCurEpochStartHeight is a free data retrieval call binding the contract method 0x5ac40790.
//
// Solidity: function getCurEpochStartHeight() view returns(uint32)
func (_EthCrossChainData *EthCrossChainDataSession) GetCurEpochStartHeight() (uint32, error) {
	return _EthCrossChainData.Contract.GetCurEpochStartHeight(&_EthCrossChainData.CallOpts)
}

// GetCurEpochStartHeight is a free data retrieval call binding the contract method 0x5ac40790.
//
// Solidity: function getCurEpochStartHeight() view returns(uint32)
func (_EthCrossChainData *EthCrossChainDataCallerSession) GetCurEpochStartHeight() (uint32, error) {
	return _EthCrossChainData.Contract.GetCurEpochStartHeight(&_EthCrossChainData.CallOpts)
}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 ethTxHashIndex) view returns(bytes32)
func (_EthCrossChainData *EthCrossChainDataCaller) GetEthTxHash(opts *core.CallOpts, ethTxHashIndex *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "getEthTxHash", ethTxHashIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 ethTxHashIndex) view returns(bytes32)
func (_EthCrossChainData *EthCrossChainDataSession) GetEthTxHash(ethTxHashIndex *big.Int) ([32]byte, error) {
	return _EthCrossChainData.Contract.GetEthTxHash(&_EthCrossChainData.CallOpts, ethTxHashIndex)
}

// GetEthTxHash is a free data retrieval call binding the contract method 0x29927875.
//
// Solidity: function getEthTxHash(uint256 ethTxHashIndex) view returns(bytes32)
func (_EthCrossChainData *EthCrossChainDataCallerSession) GetEthTxHash(ethTxHashIndex *big.Int) ([32]byte, error) {
	return _EthCrossChainData.Contract.GetEthTxHash(&_EthCrossChainData.CallOpts, ethTxHashIndex)
}

// GetEthTxHashIndex is a free data retrieval call binding the contract method 0xff3d24a1.
//
// Solidity: function getEthTxHashIndex() view returns(uint256)
func (_EthCrossChainData *EthCrossChainDataCaller) GetEthTxHashIndex(opts *core.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "getEthTxHashIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthTxHashIndex is a free data retrieval call binding the contract method 0xff3d24a1.
//
// Solidity: function getEthTxHashIndex() view returns(uint256)
func (_EthCrossChainData *EthCrossChainDataSession) GetEthTxHashIndex() (*big.Int, error) {
	return _EthCrossChainData.Contract.GetEthTxHashIndex(&_EthCrossChainData.CallOpts)
}

// GetEthTxHashIndex is a free data retrieval call binding the contract method 0xff3d24a1.
//
// Solidity: function getEthTxHashIndex() view returns(uint256)
func (_EthCrossChainData *EthCrossChainDataCallerSession) GetEthTxHashIndex() (*big.Int, error) {
	return _EthCrossChainData.Contract.GetEthTxHashIndex(&_EthCrossChainData.CallOpts)
}

// GetExtraData is a free data retrieval call binding the contract method 0x40602bb5.
//
// Solidity: function getExtraData(bytes32 key1, bytes32 key2) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCaller) GetExtraData(opts *core.CallOpts, key1 [32]byte, key2 [32]byte) ([]byte, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "getExtraData", key1, key2)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetExtraData is a free data retrieval call binding the contract method 0x40602bb5.
//
// Solidity: function getExtraData(bytes32 key1, bytes32 key2) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataSession) GetExtraData(key1 [32]byte, key2 [32]byte) ([]byte, error) {
	return _EthCrossChainData.Contract.GetExtraData(&_EthCrossChainData.CallOpts, key1, key2)
}

// GetExtraData is a free data retrieval call binding the contract method 0x40602bb5.
//
// Solidity: function getExtraData(bytes32 key1, bytes32 key2) view returns(bytes)
func (_EthCrossChainData *EthCrossChainDataCallerSession) GetExtraData(key1 [32]byte, key2 [32]byte) ([]byte, error) {
	return _EthCrossChainData.Contract.GetExtraData(&_EthCrossChainData.CallOpts, key1, key2)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_EthCrossChainData *EthCrossChainDataCaller) IsOwner(opts *core.CallOpts) (bool, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) IsOwner() (bool, error) {
	return _EthCrossChainData.Contract.IsOwner(&_EthCrossChainData.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_EthCrossChainData *EthCrossChainDataCallerSession) IsOwner() (bool, error) {
	return _EthCrossChainData.Contract.IsOwner(&_EthCrossChainData.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthCrossChainData *EthCrossChainDataCaller) Owner(opts *core.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthCrossChainData *EthCrossChainDataSession) Owner() (common.Address, error) {
	return _EthCrossChainData.Contract.Owner(&_EthCrossChainData.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EthCrossChainData *EthCrossChainDataCallerSession) Owner() (common.Address, error) {
	return _EthCrossChainData.Contract.Owner(&_EthCrossChainData.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EthCrossChainData *EthCrossChainDataCaller) Paused(opts *core.CallOpts) (bool, error) {
	var out []interface{}
	err := _EthCrossChainData.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) Paused() (bool, error) {
	return _EthCrossChainData.Contract.Paused(&_EthCrossChainData.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_EthCrossChainData *EthCrossChainDataCallerSession) Paused() (bool, error) {
	return _EthCrossChainData.Contract.Paused(&_EthCrossChainData.CallOpts)
}

// MarkFromChainTxExist is a paid mutator transaction binding the contract method 0xe90bfdcf.
//
// Solidity: function markFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) MarkFromChainTxExist(opts *core.TransactOpts, fromChainId uint64, fromChainTx [32]byte) (*common.TxResult, error) {
	return _EthCrossChainData.contract.Transact(opts, "markFromChainTxExist", fromChainId, fromChainTx)
}

// MarkFromChainTxExist is a paid mutator transaction binding the contract method 0xe90bfdcf.
//
// Solidity: function markFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) MarkFromChainTxExist(fromChainId uint64, fromChainTx [32]byte) (*common.TxResult, error) {
	return _EthCrossChainData.Contract.MarkFromChainTxExist(&_EthCrossChainData.TransactOpts, fromChainId, fromChainTx)
}

// MarkFromChainTxExist is a paid mutator transaction binding the contract method 0xe90bfdcf.
//
// Solidity: function markFromChainTxExist(uint64 fromChainId, bytes32 fromChainTx) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) MarkFromChainTxExist(fromChainId uint64, fromChainTx [32]byte) (*common.TxResult, error) {
	return _EthCrossChainData.Contract.MarkFromChainTxExist(&_EthCrossChainData.TransactOpts, fromChainId, fromChainTx)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) Pause(opts *core.TransactOpts) (*common.TxResult, error) {
	return _EthCrossChainData.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) Pause() (*common.TxResult, error) {
	return _EthCrossChainData.Contract.Pause(&_EthCrossChainData.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) Pause() (*common.TxResult, error) {
	return _EthCrossChainData.Contract.Pause(&_EthCrossChainData.TransactOpts)
}

// PutCurEpochConPubKeyBytes is a paid mutator transaction binding the contract method 0x41973cd9.
//
// Solidity: function putCurEpochConPubKeyBytes(bytes curEpochPkBytes) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) PutCurEpochConPubKeyBytes(opts *core.TransactOpts, curEpochPkBytes []byte) (*common.TxResult, error) {
	return _EthCrossChainData.contract.Transact(opts, "putCurEpochConPubKeyBytes", curEpochPkBytes)
}

// PutCurEpochConPubKeyBytes is a paid mutator transaction binding the contract method 0x41973cd9.
//
// Solidity: function putCurEpochConPubKeyBytes(bytes curEpochPkBytes) returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) PutCurEpochConPubKeyBytes(curEpochPkBytes []byte) (*common.TxResult, error) {
	return _EthCrossChainData.Contract.PutCurEpochConPubKeyBytes(&_EthCrossChainData.TransactOpts, curEpochPkBytes)
}

// PutCurEpochConPubKeyBytes is a paid mutator transaction binding the contract method 0x41973cd9.
//
// Solidity: function putCurEpochConPubKeyBytes(bytes curEpochPkBytes) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) PutCurEpochConPubKeyBytes(curEpochPkBytes []byte) (*common.TxResult, error) {
	return _EthCrossChainData.Contract.PutCurEpochConPubKeyBytes(&_EthCrossChainData.TransactOpts, curEpochPkBytes)
}

// PutCurEpochStartHeight is a paid mutator transaction binding the contract method 0x8a8bd17f.
//
// Solidity: function putCurEpochStartHeight(uint32 curEpochStartHeight) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) PutCurEpochStartHeight(opts *core.TransactOpts, curEpochStartHeight uint32) (*common.TxResult, error) {
	return _EthCrossChainData.contract.Transact(opts, "putCurEpochStartHeight", curEpochStartHeight)
}

// PutCurEpochStartHeight is a paid mutator transaction binding the contract method 0x8a8bd17f.
//
// Solidity: function putCurEpochStartHeight(uint32 curEpochStartHeight) returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) PutCurEpochStartHeight(curEpochStartHeight uint32) (*common.TxResult, error) {
	return _EthCrossChainData.Contract.PutCurEpochStartHeight(&_EthCrossChainData.TransactOpts, curEpochStartHeight)
}

// PutCurEpochStartHeight is a paid mutator transaction binding the contract method 0x8a8bd17f.
//
// Solidity: function putCurEpochStartHeight(uint32 curEpochStartHeight) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) PutCurEpochStartHeight(curEpochStartHeight uint32) (*common.TxResult, error) {
	return _EthCrossChainData.Contract.PutCurEpochStartHeight(&_EthCrossChainData.TransactOpts, curEpochStartHeight)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x4c3ccf64.
//
// Solidity: function putEthTxHash(bytes32 ethTxHash) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) PutEthTxHash(opts *core.TransactOpts, ethTxHash [32]byte) (*common.TxResult, error) {
	return _EthCrossChainData.contract.Transact(opts, "putEthTxHash", ethTxHash)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x4c3ccf64.
//
// Solidity: function putEthTxHash(bytes32 ethTxHash) returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) PutEthTxHash(ethTxHash [32]byte) (*common.TxResult, error) {
	return _EthCrossChainData.Contract.PutEthTxHash(&_EthCrossChainData.TransactOpts, ethTxHash)
}

// PutEthTxHash is a paid mutator transaction binding the contract method 0x4c3ccf64.
//
// Solidity: function putEthTxHash(bytes32 ethTxHash) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) PutEthTxHash(ethTxHash [32]byte) (*common.TxResult, error) {
	return _EthCrossChainData.Contract.PutEthTxHash(&_EthCrossChainData.TransactOpts, ethTxHash)
}

// PutExtraData is a paid mutator transaction binding the contract method 0x1afe374e.
//
// Solidity: function putExtraData(bytes32 key1, bytes32 key2, bytes value) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) PutExtraData(opts *core.TransactOpts, key1 [32]byte, key2 [32]byte, value []byte) (*common.TxResult, error) {
	return _EthCrossChainData.contract.Transact(opts, "putExtraData", key1, key2, value)
}

// PutExtraData is a paid mutator transaction binding the contract method 0x1afe374e.
//
// Solidity: function putExtraData(bytes32 key1, bytes32 key2, bytes value) returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) PutExtraData(key1 [32]byte, key2 [32]byte, value []byte) (*common.TxResult, error) {
	return _EthCrossChainData.Contract.PutExtraData(&_EthCrossChainData.TransactOpts, key1, key2, value)
}

// PutExtraData is a paid mutator transaction binding the contract method 0x1afe374e.
//
// Solidity: function putExtraData(bytes32 key1, bytes32 key2, bytes value) returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) PutExtraData(key1 [32]byte, key2 [32]byte, value []byte) (*common.TxResult, error) {
	return _EthCrossChainData.Contract.PutExtraData(&_EthCrossChainData.TransactOpts, key1, key2, value)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthCrossChainData *EthCrossChainDataTransactor) RenounceOwnership(opts *core.TransactOpts) (*common.TxResult, error) {
	return _EthCrossChainData.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthCrossChainData *EthCrossChainDataSession) RenounceOwnership() (*common.TxResult, error) {
	return _EthCrossChainData.Contract.RenounceOwnership(&_EthCrossChainData.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EthCrossChainData *EthCrossChainDataTransactorSession) RenounceOwnership() (*common.TxResult, error) {
	return _EthCrossChainData.Contract.RenounceOwnership(&_EthCrossChainData.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthCrossChainData *EthCrossChainDataTransactor) TransferOwnership(opts *core.TransactOpts, newOwner common.Address) (*common.TxResult, error) {
	return _EthCrossChainData.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthCrossChainData *EthCrossChainDataSession) TransferOwnership(newOwner common.Address) (*common.TxResult, error) {
	return _EthCrossChainData.Contract.TransferOwnership(&_EthCrossChainData.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EthCrossChainData *EthCrossChainDataTransactorSession) TransferOwnership(newOwner common.Address) (*common.TxResult, error) {
	return _EthCrossChainData.Contract.TransferOwnership(&_EthCrossChainData.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactor) Unpause(opts *core.TransactOpts) (*common.TxResult, error) {
	return _EthCrossChainData.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_EthCrossChainData *EthCrossChainDataSession) Unpause() (*common.TxResult, error) {
	return _EthCrossChainData.Contract.Unpause(&_EthCrossChainData.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_EthCrossChainData *EthCrossChainDataTransactorSession) Unpause() (*common.TxResult, error) {
	return _EthCrossChainData.Contract.Unpause(&_EthCrossChainData.TransactOpts)
}

// EthCrossChainDataOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EthCrossChainData contract.
type EthCrossChainDataOwnershipTransferredIterator struct {
	Event *EthCrossChainDataOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EthCrossChainDataOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainDataOwnershipTransferred)
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
		it.Event = new(EthCrossChainDataOwnershipTransferred)
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
func (it *EthCrossChainDataOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainDataOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainDataOwnershipTransferred represents a OwnershipTransferred event raised by the EthCrossChainData contract.
type EthCrossChainDataOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           data.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
// func (_EthCrossChainData *EthCrossChainDataFilterer) FilterOwnershipTransferred(opts *core.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EthCrossChainDataOwnershipTransferredIterator, error) {
//
// 	var previousOwnerRule []interface{}
// 	for _, previousOwnerItem := range previousOwner {
// 		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
// 	}
// 	var newOwnerRule []interface{}
// 	for _, newOwnerItem := range newOwner {
// 		newOwnerRule = append(newOwnerRule, newOwnerItem)
// 	}

// 	logs, sub, err := _EthCrossChainData.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &EthCrossChainDataOwnershipTransferredIterator{contract: _EthCrossChainData.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
// }

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EthCrossChainData *EthCrossChainDataFilterer) WatchOwnershipTransferred(opts *core.WatchOpts, sink chan<- *EthCrossChainDataOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	sub, err := _EthCrossChainData.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.UnSubscribe()
		for {
			select {
			case log := <-sub.EventMsgCh:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainDataOwnershipTransferred)
				if err := _EthCrossChainData.contract.UnpackLog(event, "OwnershipTransferred", *log); err != nil {
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

func (_EthCrossChainData *EthCrossChainDataFilterer) GetOwnershipTransferredPastEvent(txHash string, ContractLogs string) ([]*EthCrossChainDataOwnershipTransferred, error) {
	var logRaws []*data.Log
	var err error
	if ContractLogs != "" {
		logRaws, err = _EthCrossChainData.contract.GetPastEventByCtrLog(ContractLogs)
	} else if txHash != "" {
		logRaws, err = _EthCrossChainData.contract.GetPastEventByTxHash(txHash)
	} else {
		return nil, errors.New("both txHash or ContractLogs is not provided for param")
	}

	if err != nil {
		return nil, err
	}
	var events []*EthCrossChainDataOwnershipTransferred
	for _, logRaw := range logRaws {
		event, err := _EthCrossChainData.ParseOwnershipTransferred(*logRaw)
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
func (_EthCrossChainData *EthCrossChainDataFilterer) ParseOwnershipTransferred(log data.Log) (*EthCrossChainDataOwnershipTransferred, error) {
	event := new(EthCrossChainDataOwnershipTransferred)
	if err := _EthCrossChainData.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthCrossChainDataPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the EthCrossChainData contract.
type EthCrossChainDataPausedIterator struct {
	Event *EthCrossChainDataPaused // Event containing the contract specifics and raw log

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
func (it *EthCrossChainDataPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainDataPaused)
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
		it.Event = new(EthCrossChainDataPaused)
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
func (it *EthCrossChainDataPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainDataPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainDataPaused represents a Paused event raised by the EthCrossChainData contract.
type EthCrossChainDataPaused struct {
	Account common.Address
	Raw     data.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
// func (_EthCrossChainData *EthCrossChainDataFilterer) FilterPaused(opts *core.FilterOpts) (*EthCrossChainDataPausedIterator, error) {
//
//

// 	logs, sub, err := _EthCrossChainData.contract.FilterLogs(opts, "Paused")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &EthCrossChainDataPausedIterator{contract: _EthCrossChainData.contract, event: "Paused", logs: logs, sub: sub}, nil
// }

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_EthCrossChainData *EthCrossChainDataFilterer) WatchPaused(opts *core.WatchOpts, sink chan<- *EthCrossChainDataPaused) (event.Subscription, error) {

	sub, err := _EthCrossChainData.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.UnSubscribe()
		for {
			select {
			case log := <-sub.EventMsgCh:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainDataPaused)
				if err := _EthCrossChainData.contract.UnpackLog(event, "Paused", *log); err != nil {
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

func (_EthCrossChainData *EthCrossChainDataFilterer) GetPausedPastEvent(txHash string, ContractLogs string) ([]*EthCrossChainDataPaused, error) {
	var logRaws []*data.Log
	var err error
	if ContractLogs != "" {
		logRaws, err = _EthCrossChainData.contract.GetPastEventByCtrLog(ContractLogs)
	} else if txHash != "" {
		logRaws, err = _EthCrossChainData.contract.GetPastEventByTxHash(txHash)
	} else {
		return nil, errors.New("both txHash or ContractLogs is not provided for param")
	}

	if err != nil {
		return nil, err
	}
	var events []*EthCrossChainDataPaused
	for _, logRaw := range logRaws {
		event, err := _EthCrossChainData.ParsePaused(*logRaw)
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
func (_EthCrossChainData *EthCrossChainDataFilterer) ParsePaused(log data.Log) (*EthCrossChainDataPaused, error) {
	event := new(EthCrossChainDataPaused)
	if err := _EthCrossChainData.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EthCrossChainDataUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the EthCrossChainData contract.
type EthCrossChainDataUnpausedIterator struct {
	Event *EthCrossChainDataUnpaused // Event containing the contract specifics and raw log

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
func (it *EthCrossChainDataUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthCrossChainDataUnpaused)
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
		it.Event = new(EthCrossChainDataUnpaused)
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
func (it *EthCrossChainDataUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthCrossChainDataUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthCrossChainDataUnpaused represents a Unpaused event raised by the EthCrossChainData contract.
type EthCrossChainDataUnpaused struct {
	Account common.Address
	Raw     data.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
// func (_EthCrossChainData *EthCrossChainDataFilterer) FilterUnpaused(opts *core.FilterOpts) (*EthCrossChainDataUnpausedIterator, error) {
//
//

// 	logs, sub, err := _EthCrossChainData.contract.FilterLogs(opts, "Unpaused")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &EthCrossChainDataUnpausedIterator{contract: _EthCrossChainData.contract, event: "Unpaused", logs: logs, sub: sub}, nil
// }

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_EthCrossChainData *EthCrossChainDataFilterer) WatchUnpaused(opts *core.WatchOpts, sink chan<- *EthCrossChainDataUnpaused) (event.Subscription, error) {

	sub, err := _EthCrossChainData.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.UnSubscribe()
		for {
			select {
			case log := <-sub.EventMsgCh:
				// New log arrived, parse the event and forward to the user
				event := new(EthCrossChainDataUnpaused)
				if err := _EthCrossChainData.contract.UnpackLog(event, "Unpaused", *log); err != nil {
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

func (_EthCrossChainData *EthCrossChainDataFilterer) GetUnpausedPastEvent(txHash string, ContractLogs string) ([]*EthCrossChainDataUnpaused, error) {
	var logRaws []*data.Log
	var err error
	if ContractLogs != "" {
		logRaws, err = _EthCrossChainData.contract.GetPastEventByCtrLog(ContractLogs)
	} else if txHash != "" {
		logRaws, err = _EthCrossChainData.contract.GetPastEventByTxHash(txHash)
	} else {
		return nil, errors.New("both txHash or ContractLogs is not provided for param")
	}

	if err != nil {
		return nil, err
	}
	var events []*EthCrossChainDataUnpaused
	for _, logRaw := range logRaws {
		event, err := _EthCrossChainData.ParseUnpaused(*logRaw)
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
func (_EthCrossChainData *EthCrossChainDataFilterer) ParseUnpaused(log data.Log) (*EthCrossChainDataUnpaused, error) {
	event := new(EthCrossChainDataUnpaused)
	if err := _EthCrossChainData.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
