/*
* Copyright (C) 2020 The poly network Authors
* This file is part of The poly network library.
*
* The poly network is free software: you can redistribute it and/or modify
* it under the terms of the GNU Lesser General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* The poly network is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
* GNU Lesser General Public License for more details.
* You should have received a copy of the GNU Lesser General Public License
* along with The poly network . If not, see <http://www.gnu.org/licenses/>.
 */
package eth

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethComm "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	ontcommon "github.com/ontio/ontology/common"
	utils2 "github.com/ontio/ontology/smartcontract/service/native/utils"
	btcx_abi "github.com/polynetwork/poly-io-test/chains/eth/abi/btcx"
	eccd_abi "github.com/polynetwork/poly-io-test/chains/eth/abi/eccd"
	eccm_abi "github.com/polynetwork/poly-io-test/chains/eth/abi/eccm"
	eccmp_abi "github.com/polynetwork/poly-io-test/chains/eth/abi/eccmp"
	erc20_api "github.com/polynetwork/poly-io-test/chains/eth/abi/erc20"
	lockproxy_abi "github.com/polynetwork/poly-io-test/chains/eth/abi/lockproxy"
	oep4_api "github.com/polynetwork/poly-io-test/chains/eth/abi/oep4"
	ongx_api "github.com/polynetwork/poly-io-test/chains/eth/abi/ongx"
	ontx_api "github.com/polynetwork/poly-io-test/chains/eth/abi/ontx"
	"github.com/polynetwork/poly-io-test/config"
	"log"
	"math/big"
)

type EInvoker struct {
	PrivateKey     *ecdsa.PrivateKey
	ChainId        int8
	TConfiguration *config.TestConfig
	ETHUtil        *ETHTools
	NM             *NonceManager
	EthTestSigner  *EthSigner
}

var (
	DefaultGasLimit = 8000000
)

func NewEInvoker() *EInvoker {
	instance := &EInvoker{}
	instance.TConfiguration = config.DefConfig
	instance.ETHUtil = NewEthTools(instance.TConfiguration.EthURL)
	instance.NM = NewNonceManager(instance.ETHUtil.GetEthClient())
	instance.EthTestSigner, _ = NewEthSigner(instance.TConfiguration.ETHPrivateKey)
	instance.PrivateKey = instance.EthTestSigner.PrivateKey
	return instance
}

func (ethInvoker *EInvoker) MakeSmartContractAuth() (*bind.TransactOpts, error) {
	publicKey := ethInvoker.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("MakeSmartContractAuth, cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := ethInvoker.ETHUtil.GetEthClient().PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, fmt.Errorf("MakeSmartContractAuth, %v", err)
	}
	gasPrice, err := ethInvoker.ETHUtil.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("MakeSmartContractAuth, %v", err)
	}
	auth := bind.NewKeyedTransactor(ethInvoker.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(int64(0))       // in wei
	auth.GasLimit = uint64(DefaultGasLimit) // in units
	auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(10))
	return auth, nil
}

func (ethInvoker *EInvoker) DeployEthChainDataContract() (ethComm.Address, *eccd_abi.EthCrossChainData, error) {
	auth, err := ethInvoker.MakeSmartContractAuth()
	if err != nil {
		return ethComm.Address{}, nil, fmt.Errorf("DeployEthChainDataContract, err: %v", err)
	}
	contractAddress, tx, contract, err := eccd_abi.DeployEthCrossChainData(auth,
		ethInvoker.ETHUtil.GetEthClient())
	if err != nil {
		return ethComm.Address{}, nil, fmt.Errorf("DeployEthChainDataContract, err: %v", err)
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (ethInvoker *EInvoker) DeployECCMContract(eccdAddress string) (ethComm.Address, *eccm_abi.EthCrossChainManager, error) {
	auth, err := ethInvoker.MakeSmartContractAuth()
	if err != nil {
		return ethComm.Address{}, nil, fmt.Errorf("DeployECCMContract, err: %v", err)
	}
	address := ethComm.HexToAddress(eccdAddress)
	contractAddress, tx, contract, err := eccm_abi.DeployEthCrossChainManager(auth,
		ethInvoker.ETHUtil.GetEthClient(), address)
	if err != nil {
		return ethComm.Address{}, nil, fmt.Errorf("DeployECCMContract, err: %v", err)
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (ethInvoker *EInvoker) DeployECCMPContract(eccmAddress string) (ethComm.Address, *eccmp_abi.EthCrossChainManagerProxy, error) {
	auth, _ := ethInvoker.MakeSmartContractAuth()
	address := ethComm.HexToAddress(eccmAddress)
	contractAddress, tx, contract, err := eccmp_abi.DeployEthCrossChainManagerProxy(auth,
		ethInvoker.ETHUtil.GetEthClient(), address)
	if err != nil {
		return ethComm.Address{}, nil, fmt.Errorf("DeployECCMPContract, err: %v", err)
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (ethInvoker *EInvoker) DeployERC20() (ethComm.Address, *erc20_api.ERC20Template, error) {
	auth, _ := ethInvoker.MakeSmartContractAuth()
	contractAddress, tx, contract, err := erc20_api.DeployERC20Template(auth,
		ethInvoker.ETHUtil.GetEthClient())
	if err != nil {
		log.Fatal(err)
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (ethInvoker *EInvoker) DeployOEP4(lockProxy string) (ethComm.Address, *oep4_api.OEP4Template, error) {
	auth, _ := ethInvoker.MakeSmartContractAuth()
	lockProxyAddr := ethComm.HexToAddress(lockProxy)
	contractAddress, tx, contract, err := oep4_api.DeployOEP4Template(auth,
		ethInvoker.ETHUtil.GetEthClient(), lockProxyAddr)
	if err != nil {
		log.Fatal(err)
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(tx.Hash())

	auth, _ = ethInvoker.MakeSmartContractAuth()
	tx, err = contract.DeletageToProxy(auth, lockProxyAddr, big.NewInt(1e13))
	if err != nil {
		return ethComm.Address{}, nil, fmt.Errorf("failed to DeletageToProxy: %v", err)
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (ethInvoker *EInvoker) BindAssetHash(lockProxyAddr, fromAssetHash, toAssetHash string,
	toChainId int, initAmt int64) (*types.Transaction, error) {
	auth, contract, err := ethInvoker.MakeLockProxy(lockProxyAddr)
	var toAddr []byte
	if toChainId == config.ONT_CHAIN_ID {
		addr, err := ontcommon.AddressFromHexString(toAssetHash)
		if err != nil {
			return nil, err
		}
		toAddr = addr[:]
	} else if uint64(toChainId) == config.DefConfig.CMCrossChainId {
		toAddr = []byte(toAssetHash)
	}
	tx, err := contract.BindAssetHash(auth, ethComm.HexToAddress(fromAssetHash),
		uint64(toChainId), toAddr[:])
	if err != nil {
		return nil, err
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(tx.Hash())
	return tx, nil
}

func (ethInvoker *EInvoker) BindOntAsset(lockProxy, ontOnEth, ongOnEth, oep4OnEth, oep4OnOnt string) ([]*types.Transaction, error) {
	auth, contract, err := ethInvoker.MakeLockProxy(lockProxy)
	if err != nil {
		return nil, err
	}
	txs := make([]*types.Transaction, 0)

	tx1, err := contract.BindAssetHash(auth, ethComm.HexToAddress(ontOnEth),
		config.ONT_CHAIN_ID, utils2.OntContractAddress[:])
	if err != nil {
		return nil, err
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(tx1.Hash())
	txs = append(txs, tx1)

	auth, _ = ethInvoker.MakeSmartContractAuth()
	tx4, err := contract.BindAssetHash(auth, ethComm.HexToAddress(ontOnEth),
		config.DefConfig.CMCrossChainId, []byte(config.CM_ONT))
	if err != nil {
		return nil, err
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(tx4.Hash())
	txs = append(txs, tx4)

	auth, _ = ethInvoker.MakeSmartContractAuth()
	tx2, err := contract.BindAssetHash(auth, ethComm.HexToAddress(ongOnEth),
		config.ONT_CHAIN_ID, utils2.OngContractAddress[:])
	if err != nil {
		return nil, err
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(tx2.Hash())
	txs = append(txs, tx2)

	auth, _ = ethInvoker.MakeSmartContractAuth()
	tx5, err := contract.BindAssetHash(auth, ethComm.HexToAddress(ongOnEth),
		config.DefConfig.CMCrossChainId, []byte(config.CM_ONG))
	if err != nil {
		return nil, err
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(tx5.Hash())
	txs = append(txs, tx5)

	oep4, err := ontcommon.AddressFromHexString(oep4OnOnt)
	if err != nil {
		return nil, err
	}
	auth, _ = ethInvoker.MakeSmartContractAuth()
	tx3, err := contract.BindAssetHash(auth, ethComm.HexToAddress(oep4OnEth), config.ONT_CHAIN_ID, oep4[:])
	if err != nil {
		return nil, err
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(tx3.Hash())
	txs = append(txs, tx3)

	auth, _ = ethInvoker.MakeSmartContractAuth()
	tx6, err := contract.BindAssetHash(auth, ethComm.HexToAddress(oep4OnEth), config.DefConfig.CMCrossChainId, []byte(config.CM_OEP4))
	if err != nil {
		return nil, err
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(tx6.Hash())
	txs = append(txs, tx6)

	return txs, nil
}

func (ethInvoker *EInvoker) MakeLockProxy(lockProxyAddr string) (*bind.TransactOpts, *lockproxy_abi.LockProxy, error) {
	auth, _ := ethInvoker.MakeSmartContractAuth()
	contract, err := lockproxy_abi.NewLockProxy(ethComm.HexToAddress(lockProxyAddr),
		ethInvoker.ETHUtil.GetEthClient())
	if err != nil {
		log.Fatal(err)
	}
	return auth, contract, err
}

func (ethInvoker *EInvoker) DeployLockProxyContract(eccmp ethComm.Address) (ethComm.Address, *lockproxy_abi.LockProxy, error) {
	auth, _ := ethInvoker.MakeSmartContractAuth()
	contractAddress, tx, contract, err := lockproxy_abi.DeployLockProxy(auth,
		ethInvoker.ETHUtil.GetEthClient())
	if err != nil {
		return ethComm.Address{}, nil, fmt.Errorf("DeployLockProxyContract: %v", err)

	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(tx.Hash())

	auth, _ = ethInvoker.MakeSmartContractAuth()
	tx, err = contract.SetManagerProxy(auth, eccmp)
	if err != nil {
		return ethComm.Address{}, nil, fmt.Errorf("SetManagerProxy: %v", err)
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (ethInvoker *EInvoker) DeployBTCXContract(redeemscript string) (ethComm.Address, *btcx_abi.BTCX, error) {
	auth, _ := ethInvoker.MakeSmartContractAuth()
	redeemscriptBytes, err := hex.DecodeString(redeemscript)
	if err != nil {
		return ethComm.Address{}, nil, fmt.Errorf("DeployBTCXContract: %v", err)
	}
	contractAddress, tx, contract, err := btcx_abi.DeployBTCX(auth,
		ethInvoker.ETHUtil.GetEthClient(), redeemscriptBytes)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("New Deployed BTCX Contract Address is", contractAddress)
	//fmt.Println("New Deployed BTCX Contract TX is", tx.Hash().Hex())
	ethInvoker.ETHUtil.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (ethInvoker *EInvoker) DeployONTXContract(lockProxyAddr string) (ethComm.Address, *ontx_api.ONTX, error) {
	auth, _ := ethInvoker.MakeSmartContractAuth()
	contractAddress, tx, contract, err := ontx_api.DeployONTX(auth,
		ethInvoker.ETHUtil.GetEthClient(), ethComm.HexToAddress(lockProxyAddr))
	if err != nil {
		return ethComm.Address{}, nil, fmt.Errorf("DeployONTXContract, failed to deploy: %v", err)
	}
	//fmt.Println("New Deployed ONTX Contract Address is", contractAddress)
	//fmt.Println("New Deployed ONTX Contract TX is", tx.Hash().Hex())
	ethInvoker.ETHUtil.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (ethInvoker *EInvoker) DeployONGXContract(lockProxyAddr string) (ethComm.Address, *ongx_api.ONGX, error) {
	auth, _ := ethInvoker.MakeSmartContractAuth()
	contractAddress, tx, contract, err := ongx_api.DeployONGX(auth,
		ethInvoker.ETHUtil.GetEthClient(), ethComm.HexToAddress(lockProxyAddr))
	if err != nil {
		return ethComm.Address{}, nil, fmt.Errorf("DeployONGXContract, failed to deploy: %v", err)
	}
	//fmt.Println("New Deployed ONTX Contract Address is", contractAddress)
	//fmt.Println("New Deployed ONTX Contract TX is", tx.Hash().Hex())
	ethInvoker.ETHUtil.WaitTransactionConfirm(tx.Hash())
	return contractAddress, contract, nil
}

func (ethInvoker *EInvoker) SetManagerProxyForLockProxy(lockProxyAddrHex, eccmpAddressHex string) (*types.Transaction, error) {
	lockProxyAddr := ethComm.HexToAddress(lockProxyAddrHex)
	lockProxyContract, err := lockproxy_abi.NewLockProxy(lockProxyAddr, ethInvoker.ETHUtil.GetEthClient())
	if err != nil {
		return nil, fmt.Errorf("SetManagerProxyForLockProxy: %v", err)
	}
	auth, _ := ethInvoker.MakeSmartContractAuth()
	tx, err := lockProxyContract.SetManagerProxy(auth, ethComm.HexToAddress(eccmpAddressHex))
	if err != nil {
		return nil, fmt.Errorf("SetManagerProxyForLockProxy: %v", err)
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(tx.Hash())
	return tx, nil
}

func (ethInvoker *EInvoker) TransferOwnershipForECCD(eccdAddrHex, ownershipAddressHex string) (*types.Transaction, error) {
	ownershipAddress := ethComm.HexToAddress(ownershipAddressHex)
	eccdAddr := ethComm.HexToAddress(eccdAddrHex)
	eccdContract, err := eccd_abi.NewEthCrossChainData(eccdAddr, ethInvoker.ETHUtil.GetEthClient())
	if err != nil {
		return nil, fmt.Errorf("TransferOwnershipForECCD, err: %v", err)
	}
	auth, _ := ethInvoker.MakeSmartContractAuth()
	tx, err := eccdContract.TransferOwnership(auth, ownershipAddress)
	ethInvoker.ETHUtil.WaitTransactionConfirm(tx.Hash())
	return tx, nil
}

func (ethInvoker *EInvoker) TransferOwnershipForECCM(eccmAddrHex, ownershipAddressHex string) (*types.Transaction, error) {
	ownershipAddress := ethComm.HexToAddress(ownershipAddressHex)
	eccmAddr := ethComm.HexToAddress(eccmAddrHex)
	eccmContract, err := eccm_abi.NewEthCrossChainManager(eccmAddr, ethInvoker.ETHUtil.GetEthClient())
	if err != nil {
		return nil, fmt.Errorf("TransferOwnershipForECCM err: %v", err)
	}
	auth, _ := ethInvoker.MakeSmartContractAuth()
	tx, err := eccmContract.TransferOwnership(auth, ownershipAddress)
	if err != nil {
		log.Fatal(err)
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(tx.Hash())
	return tx, nil
}

func (ethInvoker *EInvoker) GetAccInfo() (string, error) {
	h, err := ethInvoker.ETHUtil.GetNodeHeight()
	if err != nil {
		return "", err
	}
	val, err := ethInvoker.ETHUtil.ethclient.BalanceAt(context.Background(), ethInvoker.EthTestSigner.Address, big.NewInt(int64(h)))
	if err != nil {
		return "", err
	}
	ethInfo := fmt.Sprintf("eth: %d", val.Uint64())

	ontx, err := ontx_api.NewONTX(ethComm.HexToAddress(ethInvoker.TConfiguration.EthOntx), ethInvoker.ETHUtil.ethclient)
	if err != nil {
		return "", err
	}
	val, err = ontx.BalanceOf(nil, ethInvoker.EthTestSigner.Address)
	if err != nil {
		return "", err
	}
	ontInfo := fmt.Sprintf("ontx: %d", val.Uint64())

	ongx, err := ongx_api.NewONGX(ethComm.HexToAddress(ethInvoker.TConfiguration.EthOngx), ethInvoker.ETHUtil.ethclient)
	if err != nil {
		return "", err
	}
	val, err = ongx.BalanceOf(nil, ethInvoker.EthTestSigner.Address)
	if err != nil {
		return "", err
	}
	ongInfo := fmt.Sprintf("ongx: %d", val.Uint64())

	oep4x, err := oep4_api.NewOEP4Template(ethComm.HexToAddress(ethInvoker.TConfiguration.EthOep4), ethInvoker.ETHUtil.ethclient)
	if err != nil {
		return "", err
	}
	val, err = oep4x.BalanceOf(nil, ethInvoker.EthTestSigner.Address)
	if err != nil {
		return "", err
	}
	oep4Info := fmt.Sprintf("oep4x: %d", val.Uint64())

	erc20, err := erc20_api.NewERC20(ethComm.HexToAddress(ethInvoker.TConfiguration.EthErc20), ethInvoker.ETHUtil.ethclient)
	if err != nil {
		return "", err
	}
	val, err = erc20.BalanceOf(nil, ethInvoker.EthTestSigner.Address)
	if err != nil {
		return "", err
	}
	erc20Info := fmt.Sprintf("erc20: %d", val.Uint64())

	btcx, err := btcx_abi.NewBTCX(ethComm.HexToAddress(ethInvoker.TConfiguration.BtceContractAddress), ethInvoker.ETHUtil.ethclient)
	if err != nil {
		return "", err
	}
	val, err = btcx.BalanceOf(nil, ethInvoker.EthTestSigner.Address)
	if err != nil {
		return "", err
	}
	btcxInfo := fmt.Sprintf("btcx: %d", val.Uint64())

	return fmt.Sprintf("ETHEREUM: acc: %s, asset: [ %s, %s, %s, %s, %s, %s ]",
		ethInvoker.EthTestSigner.Address.String(), ethInfo, ontInfo, ongInfo, oep4Info, erc20Info, btcxInfo), nil
}
