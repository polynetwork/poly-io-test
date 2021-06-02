package kai

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethComm "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joeqian10/neo-gogogo/helper"
	ontcommon "github.com/ontio/ontology/common"
	"github.com/polynetwork/eth-contracts/go_abi/btcx_abi"
	"github.com/polynetwork/eth-contracts/go_abi/eccd_abi"
	"github.com/polynetwork/eth-contracts/go_abi/eccm_abi"
	"github.com/polynetwork/eth-contracts/go_abi/eccmp_abi"
	"github.com/polynetwork/eth-contracts/go_abi/erc20_abi"
	"github.com/polynetwork/eth-contracts/go_abi/lock_proxy_abi"
	"github.com/polynetwork/eth-contracts/go_abi/oep4_abi"
	"github.com/polynetwork/eth-contracts/go_abi/ongx_abi"
	"github.com/polynetwork/eth-contracts/go_abi/ontx_abi"
	"github.com/polynetwork/kai-relayer/kaiclient"
	"github.com/polynetwork/poly-io-test/chains/eth"
	"github.com/polynetwork/poly-io-test/config"
	"github.com/polynetwork/poly-io-test/log"
)

var (
	DefaultGasLimit = 5000000
)

type Invoker struct {
	PrivateKey     *ecdsa.PrivateKey
	ChainID        uint64
	TConfiguration *config.TestConfig
	client         *kaiclient.Client
	NM             *NonceManager
	Signer         *Signer
}

func NewInvoker(chainID uint64) *Invoker {
	var err error
	instance := &Invoker{}
	instance.ChainID = chainID
	instance.TConfiguration = config.DefConfig
	client, err := kaiclient.Dial(config.DefConfig.KaiUrl)
	if err != nil {
		panic(err)
	}
	instance.client = client
	instance.NM = NewNonceManager(client)
	instance.Signer, err = NewSigner(config.DefConfig.KaiPrivateKey)
	if err != nil {
		panic(err)
	}
	instance.PrivateKey = instance.Signer.PrivateKey
	return instance
}

func (ethInvoker *Invoker) BindAssetHash(lockProxyAddr, fromAssetHash, toAssetHash string,
	toChainId uint64, initAmt int64) (*types.Transaction, error) {
	auth, contract, err := ethInvoker.MakeLockProxy(lockProxyAddr)
	if err != nil {
		return nil, err
	}
	var toAddr []byte
	if uint64(toChainId) == config.DefConfig.OntChainID {
		addr, err := ontcommon.AddressFromHexString(toAssetHash)
		if err != nil {
			return nil, err
		}
		toAddr = addr[:]
	} else if uint64(toChainId) == config.DefConfig.CMCrossChainId {
		toAddr = []byte(toAssetHash)
	} else if uint64(toChainId) == config.DefConfig.EthChainID {
		toAddr = ethComm.HexToAddress(toAssetHash).Bytes()
	} else if uint64(toChainId) == config.DefConfig.BscChainID {
		toAddr = ethComm.HexToAddress(toAssetHash).Bytes()
	} else if uint64(toChainId) == config.DefConfig.MscChainID {
		toAddr = ethComm.HexToAddress(toAssetHash).Bytes()
	} else if uint64(toChainId) == config.DefConfig.OkChainID {
		toAddr = ethComm.HexToAddress(toAssetHash).Bytes()
	} else if uint64(toChainId) == config.DefConfig.HecoChainID {
		toAddr = ethComm.HexToAddress(toAssetHash).Bytes()
	} else if uint64(toChainId) == config.DefConfig.O3ChainID {
		toAddr = ethComm.HexToAddress(toAssetHash).Bytes()
	} else if uint64(toChainId) == config.DefConfig.NeoChainID {
		other, err := helper.UInt160FromString(toAssetHash)
		if err != nil {
			return nil, err
		}
		toAddr = other[:]
	} else {
		panic(fmt.Sprintf("unkown toChainId:%d", toChainId))
	}
	tx, err := contract.BindAssetHash(auth, ethComm.HexToAddress(fromAssetHash),
		uint64(toChainId), toAddr[:])
	if err != nil {
		return nil, err
	}
	ethInvoker.client.WaitTransactionConfirm(tx.Hash())
	return tx, nil
}

func (i *Invoker) MakeSmartContractAuth() (*bind.TransactOpts, error) {
	publicKey := i.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("MakeSmartContractAuth, cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := i.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, fmt.Errorf("MakeSmartContractAuth PendingNonceAt, %v", err)
	}
	gasPrice, err := i.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("MakeSmartContractAuth SuggestGasPrice, %v", err)
	}
	auth := bind.NewKeyedTransactor(i.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(int64(0))       // in wei
	auth.GasLimit = uint64(DefaultGasLimit) // in units
	auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(1))
	return auth, nil
}

func (i *Invoker) DeployEthChainDataContract() (ethComm.Address, *eccd_abi.EthCrossChainData, error) {
	auth, err := i.MakeSmartContractAuth()
	if err != nil {
		return ethComm.Address{}, nil, fmt.Errorf("DeployEthChainDataContract, err: %v", err)
	}
	contractAddress, tx, contract, err := eccd_abi.DeployEthCrossChainData(auth,
		i.client)
	if err != nil {
		return ethComm.Address{}, nil, fmt.Errorf("DeployEthChainDataContract, err: %v", err)
	}

	i.client.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (i *Invoker) DeployECCMContract(eccdAddress string) (ethComm.Address, *eccm_abi.EthCrossChainManager, error) {
	auth, err := i.MakeSmartContractAuth()
	if err != nil {
		return ethComm.Address{}, nil, fmt.Errorf("DeployECCMContract, err: %v", err)
	}
	address := ethComm.HexToAddress(eccdAddress)
	contractAddress, tx, contract, err := eccm_abi.DeployEthCrossChainManager(auth,
		i.client, address, i.ChainID)
	if err != nil {
		return ethComm.Address{}, nil, fmt.Errorf("DeployECCMContract, err: %v", err)
	}
	i.client.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (i *Invoker) DeployECCMPContract(eccmAddress string) (ethComm.Address, *eccmp_abi.EthCrossChainManagerProxy, error) {
	auth, _ := i.MakeSmartContractAuth()
	address := ethComm.HexToAddress(eccmAddress)
	contractAddress, tx, contract, err := eccmp_abi.DeployEthCrossChainManagerProxy(auth,
		i.client, address)
	if err != nil {
		return ethComm.Address{}, nil, fmt.Errorf("DeployECCMPContract, err: %v", err)
	}
	i.client.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (i *Invoker) DeployERC20() (ethComm.Address, *erc20_abi.ERC20Template, error) {
	auth, _ := i.MakeSmartContractAuth()
	contractAddress, tx, contract, err := erc20_abi.DeployERC20Template(auth,
		i.client)
	if err != nil {
		log.Fatal(err)
	}
	i.client.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (i *Invoker) DeployOEP4(lockProxy string) (ethComm.Address, *oep4_abi.OEP4Template, error) {
	auth, _ := i.MakeSmartContractAuth()
	lockProxyAddr := ethComm.HexToAddress(lockProxy)
	contractAddress, tx, contract, err := oep4_abi.DeployOEP4Template(auth,
		i.client, lockProxyAddr)
	if err != nil {
		log.Fatal(err)
	}
	i.client.WaitTransactionConfirm(tx.Hash())

	auth, _ = i.MakeSmartContractAuth()
	tx, err = contract.DeletageToProxy(auth, lockProxyAddr, big.NewInt(1e13))
	if err != nil {
		return ethComm.Address{}, nil, fmt.Errorf("failed to DeletageToProxy: %v", err)
	}
	i.client.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (i *Invoker) MakeLockProxy(lockProxyAddr string) (*bind.TransactOpts, *lock_proxy_abi.LockProxy, error) {
	auth, _ := i.MakeSmartContractAuth()
	contract, err := lock_proxy_abi.NewLockProxy(ethComm.HexToAddress(lockProxyAddr),
		i.client)
	if err != nil {
		return nil, nil, err
	}
	return auth, contract, nil
}

func (i *Invoker) DeployLockProxyContract(eccmp ethComm.Address) (ethComm.Address, *lock_proxy_abi.LockProxy, error) {
	auth, _ := i.MakeSmartContractAuth()
	contractAddress, tx, contract, err := lock_proxy_abi.DeployLockProxy(auth,
		i.client)
	if err != nil {
		return ethComm.Address{}, nil, fmt.Errorf("DeployLockProxyContract: %v", err)

	}
	i.client.WaitTransactionConfirm(tx.Hash())

	auth, _ = i.MakeSmartContractAuth()
	tx, err = contract.SetManagerProxy(auth, eccmp)
	if err != nil {
		return ethComm.Address{}, nil, fmt.Errorf("SetManagerProxy: %v", err)
	}
	i.client.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (i *Invoker) DeployBTCXContract(redeemscript string) (ethComm.Address, *btcx_abi.BTCX, error) {
	auth, _ := i.MakeSmartContractAuth()
	redeemscriptBytes, err := hex.DecodeString(redeemscript)
	if err != nil {
		return ethComm.Address{}, nil, fmt.Errorf("DeployBTCXContract: %v", err)
	}
	contractAddress, tx, contract, err := btcx_abi.DeployBTCX(auth,
		i.client, redeemscriptBytes)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("New Deployed BTCX Contract Address is", contractAddress)
	//fmt.Println("New Deployed BTCX Contract TX is", tx.Hash().Hex())
	i.client.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (ethInvoker *Invoker) DeployONTXContract(lockProxyAddr string) (ethComm.Address, *ontx_abi.ONTX, error) {
	auth, _ := ethInvoker.MakeSmartContractAuth()
	contractAddress, tx, contract, err := ontx_abi.DeployONTX(auth,
		ethInvoker.client, ethComm.HexToAddress(lockProxyAddr))
	if err != nil {
		return ethComm.Address{}, nil, fmt.Errorf("DeployONTXContract, failed to deploy: %v", err)
	}
	//fmt.Println("New Deployed ONTX Contract Address is", contractAddress)
	//fmt.Println("New Deployed ONTX Contract TX is", tx.Hash().Hex())
	ethInvoker.client.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (i *Invoker) DeployONGXContract(lockProxyAddr string) (ethComm.Address, *ongx_abi.ONGX, error) {
	auth, _ := i.MakeSmartContractAuth()
	contractAddress, tx, contract, err := ongx_abi.DeployONGX(auth,
		i.client, ethComm.HexToAddress(lockProxyAddr))
	if err != nil {
		return ethComm.Address{}, nil, fmt.Errorf("DeployONGXContract, failed to deploy: %v", err)
	}
	//fmt.Println("New Deployed ONTX Contract Address is", contractAddress)
	//fmt.Println("New Deployed ONTX Contract TX is", tx.Hash().Hex())
	i.client.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (i *Invoker) SetManagerProxyForLockProxy(lockProxyAddrHex, eccmpAddressHex string) (*types.Transaction, error) {
	lockProxyAddr := ethComm.HexToAddress(lockProxyAddrHex)
	lockProxyContract, err := lock_proxy_abi.NewLockProxy(lockProxyAddr, i.client)
	if err != nil {
		return nil, fmt.Errorf("SetManagerProxyForLockProxy: %v", err)
	}
	auth, _ := i.MakeSmartContractAuth()
	tx, err := lockProxyContract.SetManagerProxy(auth, ethComm.HexToAddress(eccmpAddressHex))
	if err != nil {
		return nil, fmt.Errorf("SetManagerProxyForLockProxy: %v", err)
	}
	i.client.WaitTransactionConfirm(tx.Hash())
	return tx, nil
}

func (i *Invoker) TransferOwnershipForECCD(eccdAddrHex, ownershipAddressHex string) (*types.Transaction, error) {
	ownershipAddress := ethComm.HexToAddress(ownershipAddressHex)
	eccdAddr := ethComm.HexToAddress(eccdAddrHex)
	eccdContract, err := eccd_abi.NewEthCrossChainData(eccdAddr, i.client)
	if err != nil {
		return nil, fmt.Errorf("TransferOwnershipForECCD, err: %v", err)
	}
	auth, _ := i.MakeSmartContractAuth()
	tx, err := eccdContract.TransferOwnership(auth, ownershipAddress)
	i.client.WaitTransactionConfirm(tx.Hash())
	return tx, nil
}

func (i *Invoker) TransferOwnershipForECCM(eccmAddrHex, ownershipAddressHex string) (*types.Transaction, error) {
	ownershipAddress := ethComm.HexToAddress(ownershipAddressHex)
	eccmAddr := ethComm.HexToAddress(eccmAddrHex)
	eccmContract, err := eccm_abi.NewEthCrossChainManager(eccmAddr, i.client)
	if err != nil {
		return nil, fmt.Errorf("TransferOwnershipForECCM err: %v", err)
	}
	auth, _ := i.MakeSmartContractAuth()
	tx, err := eccmContract.TransferOwnership(auth, ownershipAddress)
	if err != nil {
		log.Fatal(err)
	}
	i.client.WaitTransactionConfirm(tx.Hash())
	return tx, nil
}

func (i *Invoker) Client() *kaiclient.Client {
	return i.client
}

func (i *Invoker) GetAccInfo() (string, error) {
	h, err := i.client.BlockNumber(context.Background())
	if err != nil {
		return "", err
	}
	val, err := i.client.BalanceAt(context.Background(), i.Signer.Address, big.NewInt(int64(h)))
	if err != nil {
		return "", err
	}
	ethInfo := fmt.Sprintf("eth: %d", val.Uint64())

	ontx, err := ontx_abi.NewONTX(ethComm.HexToAddress(i.TConfiguration.EthOntx), i.client)
	if err != nil {
		return "", err
	}
	val, err = ontx.BalanceOf(nil, i.Signer.Address)
	if err != nil {
		return "", err
	}
	ontInfo := fmt.Sprintf("ontx: %d", val.Uint64())

	ongx, err := ongx_abi.NewONGX(ethComm.HexToAddress(i.TConfiguration.EthOngx), i.client)
	if err != nil {
		return "", err
	}
	val, err = ongx.BalanceOf(nil, i.Signer.Address)
	if err != nil {
		return "", err
	}
	ongInfo := fmt.Sprintf("ongx: %d", val.Uint64())

	oep4x, err := oep4_abi.NewOEP4Template(ethComm.HexToAddress(i.TConfiguration.EthOep4), i.client)
	if err != nil {
		return "", err
	}
	val, err = oep4x.BalanceOf(nil, i.Signer.Address)
	if err != nil {
		return "", err
	}
	oep4Info := fmt.Sprintf("oep4x: %d", val.Uint64())

	erc20, err := erc20_abi.NewERC20(ethComm.HexToAddress(i.TConfiguration.EthErc20), i.client)
	if err != nil {
		return "", err
	}
	val, err = erc20.BalanceOf(nil, i.Signer.Address)
	if err != nil {
		return "", err
	}
	erc20Info := fmt.Sprintf("erc20: %d", val.Uint64())

	btcx, err := btcx_abi.NewBTCX(ethComm.HexToAddress(i.TConfiguration.BtceContractAddress), i.client)
	if err != nil {
		return "", err
	}
	val, err = btcx.BalanceOf(nil, i.Signer.Address)
	if err != nil {
		return "", err
	}
	btcxInfo := fmt.Sprintf("btcx: %d", val.Uint64())

	return fmt.Sprintf("ETHEREUM: acc: %s, asset: [ %s, %s, %s, %s, %s, %s ]",
		i.Signer.Address.String(), ethInfo, ontInfo, ongInfo, oep4Info, erc20Info, btcxInfo), nil
}

func (i *Invoker) GetSmartContractEventByBlock(contractAddr string, height uint64) ([]*eth.LockEvent, []*eth.UnlockEvent, error) {
	eccmAddr := common.HexToAddress(contractAddr)
	instance, err := eccm_abi.NewEthCrossChainManager(eccmAddr, i.client)
	if err != nil {
		return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, error: %s", err.Error())
	}

	opt := &bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: context.Background(),
	}

	ethlockevents := make([]*eth.LockEvent, 0)
	{
		events, err := instance.FilterCrossChainEvent(opt, nil)
		if err != nil {
			return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, error :%s", err.Error())
		}

		if events == nil {
			return nil, nil, fmt.Errorf("GetSmartContractEventByBlock - no events found on FilterCrossChainEvent")
		}

		for events.Next() {
			evt := events.Event
			ethlockevents = append(ethlockevents, &eth.LockEvent{
				Method:   "lock",
				TxHash:   evt.Raw.TxHash.String(),
				Txid:     evt.TxId,
				Saddress: evt.Sender.String(),
				Tchain:   uint32(evt.ToChainId),
				Value:    evt.Rawdata,
				Height:   height,
			})
		}
	}

	ethunlockevents := make([]*eth.UnlockEvent, 0)
	{
		events, err := instance.FilterVerifyHeaderAndExecuteTxEvent(opt)
		if err != nil {
			return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, error :%s", err.Error())
		}

		if events == nil {
			return nil, nil, fmt.Errorf("GetSmartContractEventByBlock - no events found on FilterCrossChainEvent")
		}

		for events.Next() {
			evt := events.Event
			ethunlockevents = append(ethunlockevents, &eth.UnlockEvent{
				Method:   "unlock",
				Txid:     evt.Raw.TxHash.String(),
				RTxid:    hex.EncodeToString(evt.CrossChainTxHash),
				FromTxId: hex.EncodeToString(evt.FromChainTxHash),
				Token:    hex.EncodeToString(evt.ToContract),
				Height:   height,
			})
		}
		if err != nil {
			return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, error :%s", err.Error())
		}

		if events == nil {
			return nil, nil, fmt.Errorf("GetSmartContractEventByBlock - no events found on FilterCrossChainEvent")
		}

		for events.Next() {
			evt := events.Event
			ethunlockevents = append(ethunlockevents, &eth.UnlockEvent{
				Method: "unlock",
				Txid:   evt.Raw.TxHash.String(),
				RTxid:  hex.EncodeToString(evt.CrossChainTxHash),
				Token:  hex.EncodeToString(evt.ToContract),
				Height: height,
			})
		}
	}
	return ethlockevents, ethunlockevents, nil
}
