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
package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ontio/ontology-go-sdk/utils"
	common2 "github.com/ontio/ontology/common"
	"github.com/ontio/ontology/common/log"
	"github.com/polynetwork/eth-contracts/go_abi/btcx_abi"
	"github.com/polynetwork/poly-io-test/chains/btc"
	"github.com/polynetwork/poly-io-test/chains/eth"
	"github.com/polynetwork/poly-io-test/chains/ont"
	"github.com/polynetwork/poly-io-test/config"
	"github.com/polynetwork/poly-io-test/testcase"
	"io/ioutil"
	"path"
)

var (
	btcFunc     string
	btcConfFile string
)

func init() {
	flag.StringVar(&btcFunc, "func", "new", "setup btc env from existing vendor or new one")
	flag.StringVar(&btcConfFile, "conf", "./config.json", "config file path")
	flag.Parse()
}

func main() {
	err := config.DefConfig.Init(btcConfFile)
	if err != nil {
		panic(err)
	}

	invoker, err := btc.NewBtcInvoker(config.DefConfig.RchainJsonRpcAddress, config.DefConfig.RCWallet,
		config.DefConfig.RCWalletPwd, config.DefConfig.BtcRestAddr, config.DefConfig.BtcRestUser,
		config.DefConfig.BtcRestPwd, config.DefConfig.BtcSignerPrivateKey)
	if err != nil {
		panic(fmt.Errorf("failed to new btc invoker: %v", err))
	}
	switch btcFunc {
	case "new":
		SetupNewVendor(invoker)
	case "existing":
		SetupExistingVendor(invoker)
	}
}

func SetupNewVendor(invoker *btc.BtcInvoker) {
	vendor, err := invoker.GenerateVendor(config.DefConfig.BtcMultiSigNum, config.DefConfig.BtcMultiSigRequire)
	if err != nil {
		panic(fmt.Errorf("failed to new a vendor: %v", err))
	}

	info := fmt.Sprintf("=============================== vendor info ===============================\n"+
		"multisig-address: { p2sh: %s, p2wsh: %s }\nredeem-script: %s\nredeem-key: %s\nmembers: {\n",
		vendor.P2shAddr.EncodeAddress(), vendor.P2wshAddr.EncodeAddress(), hex.EncodeToString(vendor.Redeem),
		hex.EncodeToString(vendor.HashKey))
	for i := 0; i < len(vendor.PrivateKeys); i++ {
		info += fmt.Sprintf("\tNo %d member: [ private-key: %s, public-key: %s, p2pkh-address: %s ]\n",
			i, vendor.PrivateKeys[i].String(), hex.EncodeToString(vendor.PrivateKeys[i].SerializePubKey()),
			vendor.AddressSet[i].EncodeAddress())
	}

	info += "}\nencrypted private keys for vendor signing tool: {\n"
	files, err := vendor.EncryptPrivateKeys(config.DefConfig.BtcEncryptedPrivateKeyFile,
		config.DefConfig.BtcEncryptedPrivateKeyPwd)
	if err != nil {
		panic(fmt.Errorf("failed to encrypt private keys: %v", err))
	}
	for i, f := range files {
		info += fmt.Sprintf("\tNo %d file path: %s\n", i+1, f)
	}

	for i, f := range files {
		if err = vendor.UpdateConfigFile(config.DefConfig.BtcVendorSigningToolConfFile, fmt.Sprintf("_%d", i+1),
			f, config.DefConfig.BtcEncryptedPrivateKeyPwd); err != nil {
			panic(fmt.Errorf("failed to update vendor config file: %v", err))
		}
		info += fmt.Sprintf("}\nconfig file for vendor signing tool: %s_%d\n",
			config.DefConfig.BtcVendorSigningToolConfFile, i+1)
	}

	// get btcx
	var (
		ebtcx         common.Address
		ebtcxContract *btcx_abi.BTCX
		obtcx         common2.Address
	)
	ei := eth.NewEInvoker()
	if config.DefConfig.BtceContractAddress == "" {
		ebtcx, ebtcxContract, err = ei.DeployBTCXContract(hex.EncodeToString(vendor.Redeem))
		if err != nil {
			panic(err)
		}
		auth, _ := ei.MakeSmartContractAuth()
		tx, err := ebtcxContract.SetManagerProxy(auth, common.HexToAddress(config.DefConfig.Eccmp))
		if err != nil {
			panic(fmt.Errorf("failed to set manager proxy: %v", err))
		}
		ei.ETHUtil.WaitTransactionConfirm(tx.Hash())

		auth, _ = ei.MakeSmartContractAuth()
		tx, err = ebtcxContract.BindAssetHash(auth, 1, vendor.HashKey)
		if err != nil {
			panic(fmt.Errorf("failed to bind redeem key on btcx-eth contract: %v", err))
		}

		auth, _ = ei.MakeSmartContractAuth()
		tx, err = ebtcxContract.SetMinimumLimit(auth, config.DefConfig.BtcMinOutputValFromContract)
		ei.ETHUtil.WaitTransactionConfirm(tx.Hash())
	} else {
		ebtcx = common.HexToAddress(config.DefConfig.BtceContractAddress)
		ebtcxContract, err = btcx_abi.NewBTCX(ebtcx, ei.ETHUtil.GetEthClient())
		if err != nil {
			panic(err)
		}
	}
	info += fmt.Sprintf("btcx on ether: %s\n", ebtcx.String())

	oi, err := ont.NewOntInvoker(config.DefConfig.OntJsonRpcAddress, config.DefConfig.OntContractsAvmPath,
		config.DefConfig.OntWallet, config.DefConfig.OntWalletPassword)
	if err != nil {
		panic(fmt.Errorf("failed to new ont invoker: %v", err))
	}
	if config.DefConfig.BtcoContractAddress == "" {
		raw, err := ioutil.ReadFile(path.Join(oi.OntAvmPath, "btcx.avm"))
		if err != nil {
			panic(err)
		}
		addr, err := utils.GetContractAddress(string(raw))
		if err != nil {
			panic(err)
		}
		val, err := oi.OntSdk.GetSmartContract(addr.ToHexString())
		if err == nil && val != nil {
			log.Warnf("contract btcx %s already deployed", addr.ToHexString())
		} else {
			tx, err := oi.OntSdk.NeoVM.DeployNeoVMSmartContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit, oi.OntAcc,
				true, string(raw), "btcx", "", "cooltest", "",
				"for test")
			if err != nil {
				panic(err)
			}
			oi.WaitTxConfirmation(tx)
		}

		_, err = oi.SetupBtcx(addr.ToHexString(), vendor.Redeem, vendor.HashKey,
			config.DefConfig.BtcMinOutputValFromContract, config.DefConfig.GasPrice, config.DefConfig.GasLimit)
		if err != nil {
			panic(err)
		}
		obtcx = addr
	} else {
		obtcx, _ = common2.AddressFromHexString(config.DefConfig.BtcoContractAddress)
	}
	info += fmt.Sprintf("btcx on Ontology: %s\n", obtcx.ToHexString())

	// btcx set each other
	auth, _ := ei.MakeSmartContractAuth()
	tx, err := ebtcxContract.BindAssetHash(auth, config.DefConfig.OntChainID, obtcx[:])
	if err != nil {
		panic(fmt.Errorf("bind obtc on ebtc failed: %v", err))
	}
	ei.ETHUtil.WaitTransactionConfirm(tx.Hash())
	_, err = oi.BindBtcx(obtcx.ToHexString(), ebtcx.Bytes(), config.DefConfig.EthChainID,
		config.DefConfig.GasPrice, config.DefConfig.GasLimit)
	if err != nil {
		panic(fmt.Errorf("bind ebtc on obtc failed: %v", err))
	}

	txhash, err := invoker.BindBtcxWithVendor(ebtcx.String(), config.DefConfig.EthChainID, vendor)
	if err != nil {
		panic(fmt.Errorf("failed to bind ebtcx: %v", err))
	}
	testcase.WaitPolyTx(txhash, invoker.RChain)
	info += fmt.Sprintf("bind btcx on ether txhash: %s\n", txhash.ToHexString())

	txhash, err = invoker.BindBtcxWithVendor(obtcx.ToHexString(), config.DefConfig.OntChainID, vendor)
	if err != nil {
		panic(fmt.Errorf("failed to bind obtcx: %v", err))
	}
	testcase.WaitPolyTx(txhash, invoker.RChain)
	info += fmt.Sprintf("bind btcx on ontology txhash: %s\n", txhash.ToHexString())

	txhash, err = invoker.BindBtcTxParam(config.DefConfig.BtcFeeRate, config.DefConfig.BtcMinChange, vendor)
	if err != nil {
		panic(fmt.Errorf("failed to bind tx param: %v", err))
	}
	testcase.WaitPolyTx(txhash, invoker.RChain)
	info += fmt.Sprintf("bind tx param txhash: %s\n", txhash.ToHexString())

	info += "============================================================================\n"

	fmt.Println(info)

	config.DefConfig.BtceContractAddress = ebtcx.String()
	config.DefConfig.BtcoContractAddress = obtcx.ToHexString()
	config.DefConfig.BtcRedeem = hex.EncodeToString(vendor.Redeem)
	err = config.DefConfig.Save(btcConfFile)
	if err != nil {
		panic(fmt.Errorf("failed to save config: %v", err))
	}
}

func SetupExistingVendor(invoker *btc.BtcInvoker) {
	vendor, err := btc.NewVendorFromConfig()
	if err != nil {
		panic(err)
	}

	info := fmt.Sprintf("=============================== vendor info ===============================\n"+
		"multisig-address: { p2sh: %s, p2wsh: %s }\nredeem-script: %s\nredeem-key: %s\nmembers: {\n",
		vendor.P2shAddr.EncodeAddress(), vendor.P2wshAddr.EncodeAddress(), hex.EncodeToString(vendor.Redeem),
		hex.EncodeToString(vendor.HashKey))
	for i := 0; i < len(vendor.PrivateKeys); i++ {
		info += fmt.Sprintf("\tNo %d member: [ private-key: %s, public-key: %s, p2pkh-address: %s ]\n",
			i, vendor.PrivateKeys[i].String(), hex.EncodeToString(vendor.PrivateKeys[i].SerializePubKey()),
			vendor.AddressSet[i].EncodeAddress())
	}

	// get btcx
	var (
		ebtcx         common.Address
		ebtcxContract *btcx_abi.BTCX
		obtcx         common2.Address
	)
	ei := eth.NewEInvoker()
	if config.DefConfig.BtceContractAddress == "" {
		ebtcx, ebtcxContract, err = ei.DeployBTCXContract(hex.EncodeToString(vendor.Redeem))
		if err != nil {
			panic(err)
		}
		auth, _ := ei.MakeSmartContractAuth()
		tx, err := ebtcxContract.SetManagerProxy(auth, common.HexToAddress(config.DefConfig.Eccmp))
		if err != nil {
			panic(fmt.Errorf("failed to set manager proxy: %v", err))
		}
		ei.ETHUtil.WaitTransactionConfirm(tx.Hash())

		auth, _ = ei.MakeSmartContractAuth()
		tx, err = ebtcxContract.BindAssetHash(auth, 1, vendor.HashKey)
		if err != nil {
			panic(fmt.Errorf("failed to bind redeem key on btcx-eth contract: %v", err))
		}

		auth, _ = ei.MakeSmartContractAuth()
		tx, err = ebtcxContract.SetMinimumLimit(auth, config.DefConfig.BtcMinOutputValFromContract)
		ei.ETHUtil.WaitTransactionConfirm(tx.Hash())
	} else {
		ebtcx = common.HexToAddress(config.DefConfig.BtceContractAddress)
		ebtcxContract, err = btcx_abi.NewBTCX(ebtcx, ei.ETHUtil.GetEthClient())
		if err != nil {
			panic(err)
		}
	}
	info += fmt.Sprintf("btcx on ether: %s\n", ebtcx.String())

	oi, err := ont.NewOntInvoker(config.DefConfig.OntJsonRpcAddress, config.DefConfig.OntContractsAvmPath,
		config.DefConfig.OntWallet, config.DefConfig.OntWalletPassword)
	if err != nil {
		panic(fmt.Errorf("failed to new ont invoker: %v", err))
	}
	if config.DefConfig.BtcoContractAddress == "" {
		raw, err := ioutil.ReadFile(path.Join(oi.OntAvmPath, "btcx.avm"))
		if err != nil {
			panic(err)
		}
		addr, err := utils.GetContractAddress(string(raw))
		if err != nil {
			panic(err)
		}
		val, err := oi.OntSdk.GetSmartContract(addr.ToHexString())
		if err == nil && val != nil {
			log.Warnf("contract btcx %s already deployed", addr.ToHexString())
		} else {
			tx, err := oi.OntSdk.NeoVM.DeployNeoVMSmartContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit, oi.OntAcc,
				true, string(raw), "btcx", "", "cooltest", "",
				"for test")
			if err != nil {
				panic(err)
			}
			oi.WaitTxConfirmation(tx)
		}

		_, err = oi.SetupBtcx(addr.ToHexString(), vendor.Redeem, vendor.HashKey,
			config.DefConfig.BtcMinOutputValFromContract, config.DefConfig.GasPrice, config.DefConfig.GasLimit)
		if err != nil {
			panic(err)
		}
		obtcx = addr
	} else {
		obtcx, _ = common2.AddressFromHexString(config.DefConfig.BtcoContractAddress)
	}
	info += fmt.Sprintf("btcx on Ontology: %s\n", obtcx.ToHexString())

	// btcx set each other
	auth, _ := ei.MakeSmartContractAuth()
	tx, err := ebtcxContract.BindAssetHash(auth, config.DefConfig.OntChainID, obtcx[:])
	if err != nil {
		panic(fmt.Errorf("bind obtc on ebtc failed: %v", err))
	}
	ei.ETHUtil.WaitTransactionConfirm(tx.Hash())

	auth, _ = ei.MakeSmartContractAuth()
	tx, err = ebtcxContract.BindAssetHash(auth, config.DefConfig.CMCrossChainId, []byte(config.CM_BTCX))
	if err != nil {
		panic(fmt.Errorf("bind cosmos-btc on ebtc failed: %v", err))
	}
	ei.ETHUtil.WaitTransactionConfirm(tx.Hash())

	_, err = oi.BindBtcx(obtcx.ToHexString(), ebtcx.Bytes(), config.DefConfig.EthChainID,
		config.DefConfig.GasPrice, config.DefConfig.GasLimit)
	if err != nil {
		panic(fmt.Errorf("bind ebtc on obtc failed: %v", err))
	}
	_, err = oi.BindBtcx(obtcx.ToHexString(), []byte(config.CM_BTCX), config.DefConfig.CMCrossChainId,
		config.DefConfig.GasPrice, config.DefConfig.GasLimit)
	if err != nil {
		panic(fmt.Errorf("bind cosmos-btc on obtc failed: %v", err))
	}

	txhash, err := invoker.BindBtcxWithVendor(ebtcx.String(), config.DefConfig.EthChainID, vendor)
	if err != nil {
		panic(fmt.Errorf("failed to bind ebtcx: %v", err))
	}
	testcase.WaitPolyTx(txhash, invoker.RChain)
	info += fmt.Sprintf("bind btcx on ether txhash: %s\n", txhash.ToHexString())

	txhash, err = invoker.BindBtcxWithVendor(obtcx.ToHexString(), config.DefConfig.OntChainID, vendor)
	if err != nil {
		panic(fmt.Errorf("failed to bind obtcx: %v", err))
	}
	testcase.WaitPolyTx(txhash, invoker.RChain)
	info += fmt.Sprintf("bind btcx on ontology txhash: %s\n", txhash.ToHexString())

	txhash, err = invoker.BindBtcxWithVendor(config.CM_BTCX, config.DefConfig.CMCrossChainId, vendor)
	if err != nil {
		panic(fmt.Errorf("failed to bind cosmos btcx: %v", err))
	}
	testcase.WaitPolyTx(txhash, invoker.RChain)
	info += fmt.Sprintf("bind btcx on cosmos txhash: %s\n", txhash.ToHexString())

	txhash, err = invoker.BindBtcTxParam(config.DefConfig.BtcFeeRate, config.DefConfig.BtcMinChange, vendor)
	if err != nil {
		panic(fmt.Errorf("failed to bind tx param: %v", err))
	}
	testcase.WaitPolyTx(txhash, invoker.RChain)
	info += fmt.Sprintf("bind tx param txhash: %s\n", txhash.ToHexString())

	info += "============================================================================\n"

	fmt.Println(info)

	config.DefConfig.BtceContractAddress = ebtcx.String()
	config.DefConfig.BtcoContractAddress = obtcx.ToHexString()
	err = config.DefConfig.Save(btcConfFile)
	if err != nil {
		panic(fmt.Errorf("failed to save config: %v", err))
	}
}
