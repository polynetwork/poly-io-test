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
package ont

import (
	"fmt"
	goSdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/log"
	utils2 "github.com/ontio/ontology/smartcontract/service/native/utils"
	"github.com/polynetwork/poly-io-test/config"
	"io/ioutil"
	"math/big"
	"path"
	"strings"
	"time"
)

var (
	ContractNames = []string{
		"ZeroCopySink",
		"ZeroCopySource",
		"lock_proxy",
		"ERC20Template",
		"OEP4Template",
		"ethx",
	}
)

type OntInvoker struct {
	OntSdk     *goSdk.OntologySdk
	OntAcc     *goSdk.Account
	OntAvmPath string
}

func NewOntInvoker(rpc, avmPath, wallet, pwd string) (*OntInvoker, error) {
	sdk := goSdk.NewOntologySdk()
	sdk.NewRpcClient().SetAddress(rpc)
	acc, err := GetAccountByPassword(sdk, wallet, []byte(pwd))
	if err != nil {
		return nil, err
	}
	return &OntInvoker{
		OntSdk:     sdk,
		OntAcc:     acc,
		OntAvmPath: avmPath,
	}, nil
}

//func LockInvoke(cfg *Config, account *goSdk.Account) {
//	mutTx := GenerateLockParam(cfg, account)
//	sendTxSdk := goSdk.NewOntologySdk()
//	rpcClient := client.NewRpcClient()
//	rpcClient.SetAddress(cfg.Rpc[0])
//	sendTxSdk.SetDefaultClient(rpcClient)
//	if err := signTx(sendTxSdk, mutTx, cfg.StartNonce, account); err != nil {
//		log.Error(err)
//	}
//	hash, err := sendTxSdk.SendTransaction(mutTx)
//	if err != nil {
//		log.Errorf("send tx failed, err: %s********", err)
//	} else {
//		log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
//	}
//}

func (invoker *OntInvoker) DeployContracts() ([]common.Address, error) {
	addrs := make([]common.Address, 0)
	for _, name := range ContractNames {
		raw, err := ioutil.ReadFile(path.Join(invoker.OntAvmPath, name+".avm"))
		if err != nil {
			return nil, err
		}
		addr, err := utils.GetContractAddress(string(raw))
		if err != nil {
			return nil, err
		}
		val, err := invoker.OntSdk.GetSmartContract(addr.ToHexString())
		if err == nil && val != nil {
			addrs = append(addrs, addr)
			log.Warnf("contract %s already deployed", name)
			continue
		}
		_, err = invoker.OntSdk.NeoVM.DeployNeoVMSmartContract(config.DefConfig.GasPrice,
			config.DefConfig.GasLimit, invoker.OntAcc,
			true, string(raw), name, "", "cooltest", "",
			"for test")
		if err != nil {
			return nil, err
		}
		addrs = append(addrs, addr)
	}
	return addrs, nil
}

func (invoker *OntInvoker) SetupEthAsset(lockProxy, etho, erc20, erc20o string, gasPrice, gasLimit uint64) ([]common.Uint256, error) {
	txs := make([]common.Uint256, 0)

	contractAddress, err := utils.AddressFromHexString(lockProxy)
	if err != nil {
		log.Errorf("parse contract addr failed, err: %s", err)
	}

	// etho
	ethoAddr, _ := common.AddressFromHexString(etho)
	ethAddr, _ := common.HexToBytes("0000000000000000000000000000000000000000")
	assetLimt := big.NewInt(1e18)

	res, err := invoker.OntSdk.NeoVM.PreExecInvokeNeoVMContract(ethoAddr, []interface{}{"balanceOf", []interface{}{contractAddress[:]}})
	if err != nil {
		return nil, fmt.Errorf("failed to check the ethx balance of proxy: %v", err)
	}
	val, err := res.Result.ToInteger()
	if err != nil {
		return nil, fmt.Errorf("failed to get value from result: %v", err)
	}
	tx2 := common.UINT256_EMPTY
	if val.Uint64() == 0 {
		tx2, err = invoker.OntSdk.NeoVM.InvokeNeoVMContract(gasPrice, gasLimit,
			invoker.OntAcc,
			invoker.OntAcc,
			ethoAddr,
			[]interface{}{"delegateToProxy", []interface{}{contractAddress[:], assetLimt}})
		if err != nil {
			return nil, fmt.Errorf("etho delegateToProxy error: %v", err)
		}
		invoker.WaitTxConfirmation(tx2)
	}

	tx1, err := invoker.OntSdk.NeoVM.InvokeNeoVMContract(gasPrice, gasLimit,
		invoker.OntAcc,
		invoker.OntAcc,
		contractAddress,
		[]interface{}{"bindAssetHash", []interface{}{ethoAddr[:], 2, ethAddr}})
	if err != nil {
		return nil, fmt.Errorf("etho bindAssetHash error: %v", err)
	}
	tx5, err := invoker.OntSdk.NeoVM.InvokeNeoVMContract(gasPrice, gasLimit,
		invoker.OntAcc,
		invoker.OntAcc,
		contractAddress,
		[]interface{}{"bindAssetHash", []interface{}{ethoAddr[:], config.DefConfig.CMCrossChainId, []byte(config.CM_ETHX)}})
	if err != nil {
		return nil, fmt.Errorf("etho bindAssetHash error: %v", err)
	}

	// erc20
	erc20oAddr, _ := common.AddressFromHexString(erc20o)
	erc20Addr, _ := common.HexToBytes(strings.Replace(erc20, "0x", "", 1))
	res, err = invoker.OntSdk.NeoVM.PreExecInvokeNeoVMContract(erc20oAddr, []interface{}{"balanceOf", []interface{}{contractAddress[:]}})
	if err != nil {
		return nil, fmt.Errorf("failed to check the erc20 balance of proxy: %v", err)
	}
	val, err = res.Result.ToInteger()
	if err != nil {
		return nil, fmt.Errorf("failed to get value from result: %v", err)
	}
	tx4 := common.UINT256_EMPTY
	if val.Uint64() == 0 {
		tx4, err = invoker.OntSdk.NeoVM.InvokeNeoVMContract(gasPrice, gasLimit,
			invoker.OntAcc,
			invoker.OntAcc,
			erc20oAddr,
			[]interface{}{"delegateToProxy", []interface{}{contractAddress[:], big.NewInt(1e13)}})
		if err != nil {
			return nil, fmt.Errorf("erc20 delegateToProxy error: %v", err)
		}
		invoker.WaitTxConfirmation(tx4)
	}

	tx3, err := invoker.OntSdk.NeoVM.InvokeNeoVMContract(gasPrice, gasLimit,
		invoker.OntAcc,
		invoker.OntAcc,
		contractAddress,
		[]interface{}{"bindAssetHash", []interface{}{erc20oAddr[:], 2, erc20Addr}})
	if err != nil {
		return nil, fmt.Errorf("erc20 bindAssetHash error: %v", err)
	}

	tx6, err := invoker.OntSdk.NeoVM.InvokeNeoVMContract(gasPrice, gasLimit,
		invoker.OntAcc,
		invoker.OntAcc,
		contractAddress,
		[]interface{}{"bindAssetHash", []interface{}{erc20oAddr[:], config.DefConfig.CMCrossChainId, []byte(config.CM_ERC20)}})
	if err != nil {
		return nil, fmt.Errorf("etho bindAssetHash error: %v", err)
	}

	txs = append(txs, tx1)
	txs = append(txs, tx2)
	txs = append(txs, tx3)
	txs = append(txs, tx4)
	txs = append(txs, tx5)
	txs = append(txs, tx6)

	return txs, nil
}

func (invoker *OntInvoker) GetLockedAmt(lockProxy, src common.Address) ([]byte, error) {
	res, err := invoker.OntSdk.NeoVM.PreExecInvokeNeoVMContract(lockProxy,
		[]interface{}{"getLockedAmt", []interface{}{src}})
	if err != nil {
		return nil, fmt.Errorf("get lock amt failed: %v", err)
	}
	raw, err := res.Result.ToByteArray()
	if err != nil {
		return nil, err
	}
	return raw, nil
}

func (invoker *OntInvoker) SetupOntAsset(lockProxy, onte, onge, oep4, oep4e string, gasPrice, gasLimit uint64) ([]common.Uint256, error) {
	txs := make([]common.Uint256, 0)

	contractAddress, err := utils.AddressFromHexString(lockProxy)
	if err != nil {
		return nil, fmt.Errorf("parse contract addr failed, err: %s", err)
	}

	// ont
	onteAddr, _ := common.HexToBytes(strings.Replace(onte, "0x", "", 1))
	tx1, err := invoker.OntSdk.NeoVM.InvokeNeoVMContract(gasPrice, gasLimit,
		invoker.OntAcc,
		invoker.OntAcc,
		contractAddress,
		[]interface{}{"bindAssetHash", []interface{}{utils2.OntContractAddress, 2, onteAddr}})
	if err != nil {
		return nil, fmt.Errorf("bindAssetHash Ont error: %s", err)
	}
	txs = append(txs, tx1)
	tx2, err := invoker.OntSdk.NeoVM.InvokeNeoVMContract(gasPrice, gasLimit,
		invoker.OntAcc,
		invoker.OntAcc,
		contractAddress,
		[]interface{}{"bindAssetHash", []interface{}{utils2.OntContractAddress, config.DefConfig.CMCrossChainId,
			[]byte(config.CM_ONT)}})
	if err != nil {
		return nil, fmt.Errorf("bindAssetHash Ont of cosmos error: %s", err)
	}
	txs = append(txs, tx2)

	// ong
	ongeAddr, _ := common.HexToBytes(strings.Replace(onge, "0x", "", 1))
	tx3, err := invoker.OntSdk.NeoVM.InvokeNeoVMContract(gasPrice, gasLimit,
		invoker.OntAcc,
		invoker.OntAcc,
		contractAddress,
		[]interface{}{"bindAssetHash", []interface{}{utils2.OngContractAddress[:], 2, ongeAddr}})
	if err != nil {
		return nil, fmt.Errorf("bindAssetHash ong error: %s", err)
	}
	txs = append(txs, tx3)
	tx4, err := invoker.OntSdk.NeoVM.InvokeNeoVMContract(gasPrice, gasLimit,
		invoker.OntAcc,
		invoker.OntAcc,
		contractAddress,
		[]interface{}{"bindAssetHash", []interface{}{utils2.OngContractAddress[:], config.DefConfig.CMCrossChainId,
			[]byte(config.CM_ONG)}})
	if err != nil {
		return nil, fmt.Errorf("bindAssetHash ong of cosmos error: %s", err)
	}
	txs = append(txs, tx4)

	// oep4
	oep4Addr, _ := common.AddressFromHexString(oep4)
	oep4eAddr, _ := common.HexToBytes(strings.Replace(oep4e, "0x", "", 1))
	tx5, err := invoker.OntSdk.NeoVM.InvokeNeoVMContract(gasPrice, gasLimit,
		invoker.OntAcc,
		invoker.OntAcc,
		contractAddress,
		[]interface{}{"bindAssetHash", []interface{}{oep4Addr[:], 2, oep4eAddr}})
	if err != nil {
		return nil, fmt.Errorf("bindAssetHash oep4 error: %s", err)
	}
	txs = append(txs, tx5)
	tx6, err := invoker.OntSdk.NeoVM.InvokeNeoVMContract(gasPrice, gasLimit,
		invoker.OntAcc,
		invoker.OntAcc,
		contractAddress,
		[]interface{}{"bindAssetHash", []interface{}{oep4Addr[:], config.DefConfig.CMCrossChainId, []byte(config.CM_OEP4)}})
	if err != nil {
		return nil, fmt.Errorf("bindAssetHash oep4 of cosmos error: %s", err)
	}
	txs = append(txs, tx6)
	tx7, err := invoker.OntSdk.NeoVM.InvokeNeoVMContract(gasPrice, gasLimit,
		invoker.OntAcc,
		invoker.OntAcc,
		oep4Addr,
		[]interface{}{"init", []interface{}{}})
	if err != nil {
		log.Errorf("init oep4 error, maybe already inited and approved: %s", err)
		return txs, nil
	}
	invoker.WaitTxConfirmation(tx4)
	txs = append(txs, tx7)

	tx8, err := invoker.OntSdk.NeoVM.InvokeNeoVMContract(gasPrice, gasLimit,
		invoker.OntAcc,
		invoker.OntAcc,
		oep4Addr,
		[]interface{}{"approve", []interface{}{invoker.OntAcc.Address[:], contractAddress[:], big.NewInt(1e13)}})
	if err != nil {
		return nil, fmt.Errorf("ApproveOEP4, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	txs = append(txs, tx8)

	return txs, nil
}

func (invoker *OntInvoker) SetupBtcx(btcx string, redeem, rk []byte, limit, gasPrice, gasLimit uint64) ([]common.Uint256, error) {
	txs := make([]common.Uint256, 0)

	btcxAddr, err := utils.AddressFromHexString(btcx)
	if err != nil {
		return nil, fmt.Errorf("parse contract addr failed, err: %s", err)
	}

	tx1, err := invoker.OntSdk.NeoVM.InvokeNeoVMContract(gasPrice, gasLimit, invoker.OntAcc, invoker.OntAcc, btcxAddr,
		[]interface{}{"init", []interface{}{redeem, rk}})
	if err != nil {
		return nil, err
	}
	invoker.WaitTxConfirmation(tx1)
	txs = append(txs, tx1)

	tx2, err := invoker.OntSdk.NeoVM.InvokeNeoVMContract(gasPrice, gasLimit, invoker.OntAcc, invoker.OntAcc, btcxAddr,
		[]interface{}{"setMinBackBTCLimit", []interface{}{limit}})
	if err != nil {
		return nil, err
	}
	invoker.WaitTxConfirmation(tx2)
	txs = append(txs, tx2)

	return txs, nil
}

func (invoker *OntInvoker) BindBtcx(btcx string, otherBtcx []byte, chainId int, gasPrice, gasLimit uint64) (common.Uint256, error) {
	btcxAddr, err := utils.AddressFromHexString(btcx)
	if err != nil {
		return common.UINT256_EMPTY, fmt.Errorf("parse contract addr failed, err: %s", err)
	}

	tx, err := invoker.OntSdk.NeoVM.InvokeNeoVMContract(gasPrice, gasLimit, invoker.OntAcc, invoker.OntAcc, btcxAddr,
		[]interface{}{"bindContractAddrWithChainId", []interface{}{chainId, otherBtcx}})
	if err != nil {
		return common.UINT256_EMPTY, err
	}

	return tx, nil
}

func (invoker *OntInvoker) SetOtherLockProxy(other []byte, toChainId int) (common.Uint256, error) {
	addr, err := utils.AddressFromHexString(config.DefConfig.OntLockProxy)
	if err != nil {
		return common.UINT256_EMPTY, err
	}
	txhash, err := invoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		invoker.OntAcc,
		invoker.OntAcc,
		addr,
		[]interface{}{"bindProxyHash", []interface{}{toChainId, other}})
	invoker.WaitTxConfirmation(txhash)
	return txhash, nil
}

func (invoker *OntInvoker) WaitTxConfirmation(tx common.Uint256) {
	tick := time.NewTicker(time.Second)
	for range tick.C {
		h, _ := invoker.OntSdk.GetBlockHeightByTxHash(tx.ToHexString())
		curr, _ := invoker.OntSdk.GetCurrentBlockHeight()
		if h > 0 && curr > h {
			break
		}
	}
}

func (invoker *OntInvoker) GetAccInfo() (string, error) {
	val, err := invoker.OntSdk.Native.Ont.BalanceOf(invoker.OntAcc.Address)
	if err != nil {
		return "", nil
	}
	ontInfo := fmt.Sprintf("ont: %d", val)

	val, err = invoker.OntSdk.Native.Ong.BalanceOf(invoker.OntAcc.Address)
	if err != nil {
		return "", nil
	}
	ongInfo := fmt.Sprintf("ong: %d", val)

	ethx, err := utils.AddressFromHexString(config.DefConfig.OntEth)
	if err != nil {
		return "", err
	}
	res, err := invoker.OntSdk.NeoVM.PreExecInvokeNeoVMContract(ethx, []interface{}{"balanceOf", []interface{}{invoker.OntAcc.Address[:]}})
	if err != nil {
		return "", err
	}
	bigVal, err := res.Result.ToInteger()
	if err != nil {
		return "", err
	}
	ethxInfo := fmt.Sprintf("ethx: %d", bigVal.Uint64())

	erc20x, err := utils.AddressFromHexString(config.DefConfig.OntErc20)
	if err != nil {
		return "", err
	}
	res, err = invoker.OntSdk.NeoVM.PreExecInvokeNeoVMContract(erc20x, []interface{}{"balanceOf", []interface{}{invoker.OntAcc.Address[:]}})
	if err != nil {
		return "", err
	}
	bigVal, err = res.Result.ToInteger()
	if err != nil {
		return "", err
	}
	erc20xInfo := fmt.Sprintf("erc20x: %d", bigVal.Uint64())

	oep4x, err := utils.AddressFromHexString(config.DefConfig.OntOep4)
	if err != nil {
		return "", err
	}
	res, err = invoker.OntSdk.NeoVM.PreExecInvokeNeoVMContract(oep4x, []interface{}{"balanceOf", []interface{}{invoker.OntAcc.Address[:]}})
	if err != nil {
		return "", err
	}
	bigVal, err = res.Result.ToInteger()
	if err != nil {
		return "", err
	}
	oep4xInfo := fmt.Sprintf("oep4x: %d", bigVal.Uint64())

	btco, err := utils.AddressFromHexString(config.DefConfig.BtcoContractAddress)
	if err != nil {
		return "", err
	}
	res, err = invoker.OntSdk.NeoVM.PreExecInvokeNeoVMContract(btco, []interface{}{"balanceOf", []interface{}{invoker.OntAcc.Address[:]}})
	if err != nil {
		return "", err
	}
	bigVal, err = res.Result.ToInteger()
	if err != nil {
		return "", err
	}
	btcoInfo := fmt.Sprintf("btco: %d", bigVal.Uint64())

	return fmt.Sprintf("ONTOLOGY: acc: %s, asset: [ %s, %s, %s, %s, %s, %s ]",
		invoker.OntAcc.Address.ToBase58(), ontInfo, ongInfo, ethxInfo, erc20xInfo, oep4xInfo, btcoInfo), nil
}

//func BindAssetHash(cfg *Config, account *goSdk.Account) {
//	fromAsstHash := cfg.EthX
//	toAssetHash := "0000000000000000000000000000000000000000"
//	toChainID := uint64(2)
//	Factor := int64(1000000000000000000)
//	assetLimt := common.BigIntToNeoBytes(big.NewInt(4 * Factor))
//	isTargetChainAsset := bool(true)
//	mutTx := GenerateBindAsstHashTx(cfg, fromAsstHash, toChainID, toAssetHash, assetLimt, isTargetChainAsset)
//	sendTxSdk := goSdk.NewOntologySdk()
//	rpcClient := client.NewRpcClient()
//	rpcClient.SetAddress(cfg.Rpc[0])
//	sendTxSdk.SetDefaultClient(rpcClient)
//	if err := signTx(sendTxSdk, mutTx, cfg.StartNonce, account); err != nil {
//		log.Error(err)
//	}
//	hash, err := sendTxSdk.SendTransaction(mutTx)
//	if err != nil {
//		log.Errorf("send tx failed, err: %s********", err)
//	} else {
//		log.Infof("send tx %s****sentnum:***%d", hash.ToHexString(), cfg.StartNonce)
//	}
//}

//func GetAssetHash(cfg *Config, account *goSdk.Account) {
//	sendTxSdk := goSdk.NewOntologySdk()
//	rpcClient := client.NewRpcClient()
//	rpcClient.SetAddress(cfg.Rpc[0])
//	sendTxSdk.SetDefaultClient(rpcClient)
//	contractAddr, err := utils.AddressFromHexString(cfg.LockProxy)
//	if err != nil {
//		log.Errorf("balanceOf: decode contract addr failed, err: %s", err)
//		return
//	}
//	FromAssetHash, _ := common.HexToBytes(cfg.EthX)
//	params := []interface{}{"getAssetHash", []interface{}{FromAssetHash, cfg.ToChainId}}
//	preResult, err := sendTxSdk.NeoVM.PreExecInvokeNeoVMContract(contractAddr, params)
//	if err != nil {
//		log.Errorf("balanceOf: pre-execute failed, err: %s", err)
//		return
//	}
//	log.Infof("result is %x", preResult.Result)
//}
//
//func GetCrossLimit(cfg *Config, account *goSdk.Account) {
//	sendTxSdk := goSdk.NewOntologySdk()
//	rpcClient := client.NewRpcClient()
//	rpcClient.SetAddress(cfg.Rpc[0])
//	sendTxSdk.SetDefaultClient(rpcClient)
//	contractAddr, err := utils.AddressFromHexString(cfg.LockProxy)
//	if err != nil {
//		log.Errorf("balanceOf: decode contract addr failed, err: %s", err)
//		return
//	}
//	FromAssetHash, _ := common.HexToBytes(cfg.EthX)
//	params := []interface{}{"getCrossedLimit", []interface{}{FromAssetHash, cfg.ToChainId}}
//	preResult, err := sendTxSdk.NeoVM.PreExecInvokeNeoVMContract(contractAddr, params)
//	if err != nil {
//		log.Errorf("balanceOf: pre-execute failed, err: %s", err)
//		return
//	}
//	log.Infof("result is %x", preResult.Result)
//}
