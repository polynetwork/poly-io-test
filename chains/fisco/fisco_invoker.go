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
package fisco

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum/common"
	common2 "github.com/ontio/ontology/common"
	"github.com/polynetwork/poly-io-test/chains/fisco/go_abi/eccd_abi"  // import eccd_abi
	"github.com/polynetwork/poly-io-test/chains/fisco/go_abi/eccm_abi"  // import eccm_abi
	"github.com/polynetwork/poly-io-test/chains/fisco/go_abi/eccmp_abi" // import eccmp_abi
	"github.com/polynetwork/poly-io-test/chains/fisco/go_abi/erc20_abi"
	"github.com/polynetwork/poly-io-test/chains/fisco/go_abi/fiscox_abi" // import eccd_abi
	"github.com/polynetwork/poly-io-test/chains/fisco/go_abi/lock_proxy_abi"
	"github.com/polynetwork/poly-io-test/chains/fisco/go_abi/oep4_abi"
	"github.com/polynetwork/poly-io-test/chains/fisco/go_abi/ongx_abi"
	"github.com/polynetwork/poly-io-test/chains/fisco/go_abi/ontx_abi"
	config2 "github.com/polynetwork/poly-io-test/config"
	"github.com/polynetwork/poly-io-test/log"
	"math/big"
	"strconv"
	"time"
	"unsafe"
)

type FiscoInvoker struct {
	FiscoSdk *client.Client
}

var (
	DefaultGasLimit = 6000000
)

func NewFiscoInvoker() (*FiscoInvoker, error) {
	instance := &FiscoInvoker{}
	config := &conf.ParseConfig(config2.DefConfig.FiscoSdkConfFile)[0]

	cli, err := client.Dial(config)
	if err != nil {
		return nil, err
	}
	instance.FiscoSdk = cli
	return instance, nil
}

func (fisInvoker *FiscoInvoker) DeployEthChainDataContract() (common.Address, *eccd_abi.EthCrossChainData, error) {
	contractAddress, tx, contract, err := eccd_abi.DeployEthCrossChainData(fisInvoker.FiscoSdk.GetTransactOpts(),
		fisInvoker.FiscoSdk)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("DeployEthChainDataContract, err: %v", err)
	}
	fmt.Println("DeployEthChainDataContract tx", tx.Hash().Hex())
	return contractAddress, contract, nil
}

func (fisInvoker *FiscoInvoker) DeployECCMContract(eccdAddress string) (common.Address, *eccm_abi.EthCrossChainManager, error) {

	address := common.HexToAddress(eccdAddress)
	contractAddress, tx, contract, err := eccm_abi.DeployEthCrossChainManager(fisInvoker.FiscoSdk.GetTransactOpts(),
		fisInvoker.FiscoSdk, address)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("DeployECCMContract, err: %v", err)
	}
	fmt.Println("DeployECCMContract tx", tx.Hash().Hex())

	return contractAddress, contract, nil
}

func (fisInvoker *FiscoInvoker) DeployECCMPContract(eccmAddress string) (common.Address, *eccmp_abi.EthCrossChainManagerProxy, error) {

	address := common.HexToAddress(eccmAddress)
	contractAddress, tx, contract, err := eccmp_abi.DeployEthCrossChainManagerProxy(fisInvoker.FiscoSdk.GetTransactOpts(),
		fisInvoker.FiscoSdk, address)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("DeployECCMPContract, err: %v", err)
	}
	fmt.Println("DeployECCMPContract tx", tx.Hash().Hex())

	return contractAddress, contract, nil
}

func (fisInvoker *FiscoInvoker) MakeSmartContractAuth() *bind.TransactOpts {
	return fisInvoker.FiscoSdk.GetTransactOpts()
}

func (fisInvoker *FiscoInvoker) MakeSmartClient() *client.Client {
	return fisInvoker.FiscoSdk
}

func (fisInvoker *FiscoInvoker) TransferOwnershipForECCD(eccdAddrHex, ownershipAddressHex string) (*types.Transaction, error) {
	ownershipAddress := common.HexToAddress(ownershipAddressHex)
	eccdAddr := common.HexToAddress(eccdAddrHex)
	eccdContract, err := eccd_abi.NewEthCrossChainData(eccdAddr, fisInvoker.FiscoSdk)
	if err != nil {
		return nil, fmt.Errorf("TransferOwnershipForECCD, err: %v", err)
	}

	tx, err := eccdContract.TransferOwnership(fisInvoker.FiscoSdk.GetTransactOpts(), ownershipAddress)
	fmt.Println("TransferOwnershipForECCD tx", tx.Hash().Hex())
	return tx, nil
}

func (fisInvoker *FiscoInvoker) TransferOwnershipForECCM(eccmAddrHex, ownershipAddressHex string) (*types.Transaction, error) {
	ownershipAddress := common.HexToAddress(ownershipAddressHex)
	eccmAddr := common.HexToAddress(eccmAddrHex)
	eccmContract, err := eccm_abi.NewEthCrossChainManager(eccmAddr, fisInvoker.FiscoSdk)
	if err != nil {
		return nil, fmt.Errorf("TransferOwnershipForECCM err: %v", err)
	}

	tx, err := eccmContract.TransferOwnership(fisInvoker.FiscoSdk.GetTransactOpts(), ownershipAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("TransferOwnershipForECCM tx", tx.Hash().Hex())

	return tx, nil
}

func (fisInvoker *FiscoInvoker) DeployLockProxyContract(eccmp common.Address) (common.Address, *lock_proxy_abi.LockProxy, error) {
	contractAddress, tx, contract, err := lock_proxy_abi.DeployLockProxy(fisInvoker.FiscoSdk.GetTransactOpts(),
		fisInvoker.FiscoSdk)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("DeployLockProxyContract: %v", err)

	}

	fmt.Println("DeployLockProxy tx", tx.Hash().Hex())

	tx, err = contract.SetManagerProxy(fisInvoker.FiscoSdk.GetTransactOpts(), eccmp)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("SetManagerProxy: %v", err)
	}
	fmt.Println("SetManagerProxy tx", tx.Hash().Hex())

	return contractAddress, contract, nil
}

func (fisInvoker *FiscoInvoker) SendFiscoTx(fiscoxAddress common.Address, account common.Address) {
	fiscox, err := fiscox_abi.NewFISCOX(fiscoxAddress, fisInvoker.FiscoSdk)
	if err != nil {

		log.Fatal(err)
	}
	total, err := fiscox.TotalSupply(fisInvoker.FiscoSdk.GetCallOpts())
	fmt.Println("TotalSupply:", total)
	balance, err := fiscox.BalanceOf(fisInvoker.FiscoSdk.GetCallOpts(), account)
	fmt.Println("account "+account.Hex()+" balance is :", balance)
	//big1 := new(big.Int).SetUint64(uint64(100))
	address := common.HexToAddress("0xd280d750dDe9ab2f07cd524a28C81C54485fD1bc")
	balance0, err := fiscox.BalanceOf(fisInvoker.FiscoSdk.GetCallOpts(), address)
	fmt.Println("0xd280d750dDe9ab2f07cd524a28C81C54485fD1bc balance is :", balance0)

	address1 := common.HexToAddress("0x34f00110bad3236f01468799d44fe04d7deb25f0")
	balance1, err := fiscox.BalanceOf(fisInvoker.FiscoSdk.GetCallOpts(), address1)
	fmt.Println("0x34f00110bad3236f01468799d44fe04d7deb25f0 balance is :", balance1)

	decimal, err := fiscox.Decimals(fisInvoker.FiscoSdk.GetCallOpts())
	if err != nil {

		log.Fatal(err)
	}
	fmt.Println("decimal  is :", decimal)

	allow, err := fiscox.Allowance(fisInvoker.FiscoSdk.GetCallOpts(), address1, address)
	fmt.Println("allow  is :", allow)

}

func (fisInvoker *FiscoInvoker) DeployERC20() (common.Address, *erc20_abi.ERC20Template, error) {

	contractAddress, tx, contract, err := erc20_abi.DeployERC20Template(fisInvoker.FiscoSdk.GetTransactOpts(),
		fisInvoker.FiscoSdk)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DeployERC20Template tx", tx.Hash().Hex())
	return contractAddress, contract, nil
}

func (fisInvoker *FiscoInvoker) NewERC20Template(address common.Address) (*erc20_abi.ERC20TemplateCaller, error) {

	erc20template, err := erc20_abi.NewERC20TemplateCaller(address, fisInvoker.FiscoSdk)
	if err != nil {
		log.Fatal(err)
	}
	return erc20template, err
}

func (fisInvoker *FiscoInvoker) DeployOEP4(lockProxy string) (common.Address, *oep4_abi.OEP4Template, error) {

	lockProxyAddr := common.HexToAddress(lockProxy)
	contractAddress, tx, contract, err := oep4_abi.DeployOEP4Template(fisInvoker.FiscoSdk.GetTransactOpts(),
		fisInvoker.FiscoSdk, lockProxyAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DeployOEP4Template tx", tx.Hash().Hex())

	tx, err = contract.DeletageToProxy(fisInvoker.FiscoSdk.GetTransactOpts(), lockProxyAddr, big.NewInt(1e13))
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("failed to DeletageToProxy: %v", err)
	}
	fmt.Println("DeletageToProxy tx", tx.Hash().Hex())

	return contractAddress, contract, nil
}

func (fisInvoker *FiscoInvoker) DeployFISCOXContract(lockProxy string) (common.Address, *fiscox_abi.FISCOX, error) {

	lockProxyAddr := common.HexToAddress(lockProxy)
	contractAddress, tx, contract, err := fiscox_abi.DeployFISCOX(fisInvoker.FiscoSdk.GetTransactOpts(),
		fisInvoker.FiscoSdk, lockProxyAddr)
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("DeployFISCOXContract, failed to deploy: %v", err)
	}
	//fmt.Println("New Deployed ONTX Contract Address is", contractAddress)
	//fmt.Println("New Deployed ONTX Contract TX is", tx.Hash().Hex())
	fmt.Println("DeployFISCOX tx", tx.Hash().Hex())

	return contractAddress, contract, nil
}

func (fisInvoker *FiscoInvoker) BindLockProxy(lockProxy string, toChainId uint64, targetProxy string) error {
	address := common.HexToAddress(lockProxy)
	instance1, err := lock_proxy_abi.NewLockProxy(address, fisInvoker.FiscoSdk)
	if err != nil {
		return err
	}
	var toAddr []byte
	if toChainId == config2.DefConfig.OntChainID {
		addr, err := common2.AddressFromHexString(targetProxy)
		if err != nil {
			return err
		}
		toAddr = addr[:]
	} else if uint64(toChainId) == config2.DefConfig.CMCrossChainId {
		toAddr, err = hex.DecodeString(targetProxy)
		if err != nil {
			return err
		}
	} else if toChainId == config2.DefConfig.EthChainID {
		toAddr = common.HexToAddress(targetProxy).Bytes()
	} else if toChainId == config2.DefConfig.FabricChainID {
		toAddr = []byte(targetProxy)
	}
	trans, err := instance1.BindProxyHash(fisInvoker.FiscoSdk.GetTransactOpts(), toChainId, toAddr)
	if err != nil {
		return fmt.Errorf("BindProxyHash error: %v", err)
	}
	fmt.Printf("BindProxyHash tx: %v\n", trans.Hash().Hex())

	return nil
}

func (fisInvoker *FiscoInvoker) DeployONGXContract(lockProxyAddr string) (common.Address, *ongx_abi.ONGX, error) {

	contractAddress, tx, contract, err := ongx_abi.DeployONGX(fisInvoker.FiscoSdk.GetTransactOpts(),
		fisInvoker.FiscoSdk, common.HexToAddress(lockProxyAddr))
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("DeployONGXContract, failed to deploy: %v", err)
	}
	//fmt.Println("New Deployed ONTX Contract Address is", contractAddress)
	//fmt.Println("New Deployed ONTX Contract TX is", tx.Hash().Hex())
	fmt.Println("DeletageToProxy tx", tx.Hash().Hex())

	return contractAddress, contract, nil
}

func (fisInvoker *FiscoInvoker) DeployONTXContract(lockProxyAddr string) (common.Address, *ontx_abi.ONTX, error) {

	contractAddress, tx, contract, err := ontx_abi.DeployONTX(fisInvoker.FiscoSdk.GetTransactOpts(),
		fisInvoker.FiscoSdk, common.HexToAddress(lockProxyAddr))
	if err != nil {
		return common.Address{}, nil, fmt.Errorf("DeployONTXContract, failed to deploy: %v", err)
	}
	fmt.Println("DeletageToProxy tx", tx.Hash().Hex())
	//fmt.Println("New Deployed ONTX Contract Address is", contractAddress)
	//fmt.Println("New Deployed ONTX Contract TX is", tx.Hash().Hex())

	return contractAddress, contract, nil
}

func (fisInvoker *FiscoInvoker) SetManagerProxyForLockProxy(lockProxyAddrHex, eccmpAddressHex string) (*types.Transaction, error) {
	lockProxyAddr := common.HexToAddress(lockProxyAddrHex)
	lockProxyContract, err := lock_proxy_abi.NewLockProxy(lockProxyAddr, fisInvoker.FiscoSdk)
	if err != nil {
		return nil, fmt.Errorf("SetManagerProxyForLockProxy: %v", err)
	}
	auth := fisInvoker.MakeSmartContractAuth()
	tx, err := lockProxyContract.SetManagerProxy(auth, common.HexToAddress(eccmpAddressHex))
	if err != nil {
		return nil, fmt.Errorf("SetManagerProxyForLockProxy: %v", err)
	}

	fmt.Println("SetManagerProxy tx", tx.Hash().Hex())
	return tx, nil
}

func (fisInvoker *FiscoInvoker) MakeLockProxy(lockProxyAddr string) (*bind.TransactOpts, *lock_proxy_abi.LockProxy, error) {
	auth := fisInvoker.MakeSmartContractAuth()
	contract, err := lock_proxy_abi.NewLockProxy(common.HexToAddress(lockProxyAddr),
		fisInvoker.FiscoSdk)
	if err != nil {
		return nil, nil, err
	}
	return auth, contract, err
}

func (fisInvoker *FiscoInvoker) BindAssetHash(lockProxyAddr, fromAssetHash, toAssetHash string, toChainId uint64) (*types.Transaction, error) {
	auth, contract, err := fisInvoker.MakeLockProxy(lockProxyAddr)
	var toAddr []byte
	if uint64(toChainId) == config2.DefConfig.OntChainID {
		addr, err := common2.AddressFromHexString(toAssetHash)
		if err != nil {
			return nil, err
		}
		toAddr = addr[:]
	} else if uint64(toChainId) == config2.DefConfig.CMCrossChainId {
		toAddr = []byte(toAssetHash)
	} else if uint64(toChainId) == config2.DefConfig.EthChainID {
		toAddr = common.HexToAddress(toAssetHash).Bytes()
	} else if uint64(toChainId) == config2.DefConfig.FabricChainID {
		toAddr = []byte(toAssetHash)
	}
	tx, err := contract.BindAssetHash(auth, common.HexToAddress(fromAssetHash),
		uint64(toChainId), toAddr[:])
	if err != nil {
		return nil, err
	}

	fmt.Println("BindAssetHash tx", tx.Hash().Hex())
	return tx, nil
}

func (this *FiscoInvoker) StartTransact(toAccAddr string, amount int64) {
	lockAddr := common.HexToAddress(config2.DefConfig.FiscoLockProxy)
	instance1, err := lock_proxy_abi.NewLockProxy(lockAddr, this.FiscoSdk)
	if err != nil {
		log.Fatal(err)
	}
	lock_proxy := &lock_proxy_abi.LockProxySession{Contract: instance1, CallOpts: *this.FiscoSdk.GetCallOpts(), TransactOpts: *this.FiscoSdk.GetTransactOpts()}

	fiscoxAddress := common.HexToAddress(config2.DefConfig.FiscoEth)

	if err != nil {
		log.Fatal(err)
	}

	toAcc := common.HexToAddress(toAccAddr)
	tx1, err := lock_proxy.Lock(fiscoxAddress, 2, toAcc.Bytes(), big.NewInt(amount))

	if err != nil {
		log.Fatalf("Lock proxy: %v", err)
		return
	}
	log.Infof("Lock tx1 %s", tx1.Hash().Hex())

}

func (this *FiscoInvoker) BlockNumber() (int64, error) {
	bn, err := this.FiscoSdk.GetBlockNumber(context.Background())
	if err != nil {
		return 0, fmt.Errorf("block number not found: %v", err)
	}
	str, err := strconv.Unquote(*(*string)(unsafe.Pointer(&bn)))
	if err != nil {
		return 0, fmt.Errorf("ParseInt: %v", err)
	}
	height, err := strconv.ParseInt(str, 0, 0)
	if err != nil {
		return 0, fmt.Errorf("ParseInt: %v", err)
	}
	return height, nil
}

func (this *FiscoInvoker) WaitTransactionConfirm(hash common.Hash) bool {
	fiscoTx := hash.String()
	for {
		time.Sleep(time.Second * 1)
		_, err := this.FiscoSdk.GetTransactionByHash(context.Background(), fiscoTx)
		if err != nil {
			continue
		}
		receipt, err := this.FiscoSdk.TransactionReceipt(context.Background(), hash)
		if err != nil || receipt.GetBlockNumber() == "" {
			continue
		}
		log.Debugf("( fisco_transaction %s, Status %s ) is pending", fiscoTx, receipt.Status)
		return true
	}
}
