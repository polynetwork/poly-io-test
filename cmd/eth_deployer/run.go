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
	common2 "github.com/ethereum/go-ethereum/common"
	"github.com/ontio/ontology/common"
	"github.com/polynetwork/poly-io-test/chains/eth"
	"github.com/polynetwork/poly-io-test/config"
)

var (
	fnEth        string
	ethConfFile  string
	eccmRedeploy int
)

func init() {
	flag.StringVar(&fnEth, "func", "deploy", "choose function to run: deploy or setup")
	flag.StringVar(&ethConfFile, "conf", "./config.json", "config file path")
	flag.IntVar(&eccmRedeploy, "redeploy_eccm", 1, "redeploy eccd, eccm and eccmp or not")
	flag.Parse()
}

func main() {
	err := config.DefConfig.Init(ethConfFile)
	if err != nil {
		panic(err)
	}

	switch fnEth {
	case "deploy":
		SetupETHSmartContract()
	case "setup":
		SetUpLockProxyForETH()
	}
}

func SetupETHSmartContract() {
	invoker := eth.NewEInvoker()
	var (
		eccdAddr  common2.Address
		eccmAddr  common2.Address
		eccmpAddr common2.Address
		err       error
	)
	if eccmRedeploy == 1 {
		eccdAddr, _, err = invoker.DeployEthChainDataContract()
		if err != nil {
			panic(err)
		}

		eccmAddr, _, err = invoker.DeployECCMContract(eccdAddr.Hex())
		if err != nil {
			panic(err)
		}
		eccmpAddr, _, err = invoker.DeployECCMPContract(eccmAddr.Hex())
		if err != nil {
			panic(err)
		}
		_, err = invoker.TransferOwnershipForECCD(eccdAddr.Hex(), eccmAddr.Hex())
		if err != nil {
			panic(err)
		}
		_, err = invoker.TransferOwnershipForECCM(eccmAddr.Hex(), eccmpAddr.Hex())
		if err != nil {
			panic(err)
		}
	} else {
		eccdAddr = common2.HexToAddress(config.DefConfig.Eccd)
		eccmAddr = common2.HexToAddress(config.DefConfig.Eccm)
		eccmpAddr = common2.HexToAddress(config.DefConfig.Eccmp)
	}

	lockProxyAddr, _, err := invoker.DeployLockProxyContract(eccmpAddr)
	if err != nil {
		panic(err)
	}

	lockproxyAddrHex := lockProxyAddr.Hex()
	erc20Addr, erc20, err := invoker.DeployERC20()
	if err != nil {
		panic(err)
	}

	total, err := erc20.TotalSupply(nil)
	if err != nil {
		panic(fmt.Errorf("failed to get total supply for erc20: %v", err))
	}
	auth, _ := invoker.MakeSmartContractAuth()
	tx, err := erc20.Approve(auth, lockProxyAddr, total)
	if err != nil {
		panic(fmt.Errorf("failed to approve erc20 to lockproxy: %v", err))
	}
	invoker.ETHUtil.WaitTransactionConfirm(tx.Hash())

	oep4Addr, _, err := invoker.DeployOEP4(lockproxyAddrHex)
	if err != nil {
		panic(err)
	}
	ongxAddr, _, err := invoker.DeployONGXContract(lockproxyAddrHex)
	if err != nil {
		panic(err)
	}
	ontxAddr, _, err := invoker.DeployONTXContract(lockproxyAddrHex)
	if err != nil {
		panic(err)
	}

	fmt.Println("=============================ETH info=============================")
	fmt.Println("erc20:", erc20Addr.Hex())
	fmt.Println("ope4:", oep4Addr.Hex())
	fmt.Println("eccd address:", eccdAddr.Hex())
	fmt.Println("eccm address:", eccmAddr.Hex())
	fmt.Println("eccmp address:", eccmpAddr.Hex())
	fmt.Println("lock proxy address: ", lockProxyAddr.Hex())
	fmt.Println("ongx address: ", ongxAddr.Hex())
	fmt.Println("ontx proxy address: ", ontxAddr.Hex())
	fmt.Println("==================================================================")

	config.DefConfig.EthErc20 = erc20Addr.Hex()
	config.DefConfig.EthOep4 = oep4Addr.Hex()
	config.DefConfig.Eccd = eccdAddr.Hex()
	config.DefConfig.Eccm = eccmAddr.Hex()
	config.DefConfig.Eccmp = eccmpAddr.Hex()
	config.DefConfig.EthLockProxy = lockProxyAddr.Hex()
	config.DefConfig.EthOngx = ongxAddr.Hex()
	config.DefConfig.EthOntx = ontxAddr.Hex()

	if err := config.DefConfig.Save(ethConfFile); err != nil {
		panic(fmt.Errorf("failed to save config, you better save it youself: %v", err))
	}
}

func SetupERC20(ethInvoker *eth.EInvoker) {
	bindTx, err := ethInvoker.BindAssetHash(config.DefConfig.EthLockProxy, config.DefConfig.EthErc20,
		config.DefConfig.OntErc20, config.ONT_CHAIN_ID, 0)
	if err != nil {
		panic(fmt.Errorf("SetupERC20ToONT, failed to BindAssetHash: %v", err))
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(bindTx.Hash())
	hash := bindTx.Hash()
	fmt.Printf("binding erc20 of ontology on ethereum: ( txhash: %s )\n", hash.String())

	bindTx, err = ethInvoker.BindAssetHash(config.DefConfig.EthLockProxy, config.DefConfig.EthErc20,
		config.CM_ERC20, int(config.DefConfig.CMCrossChainId), 0)
	if err != nil {
		panic(fmt.Errorf("SetupERC20ToONT, failed to BindAssetHash: %v", err))
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(bindTx.Hash())
	hash = bindTx.Hash()
	fmt.Printf("binding erc20 of cosmos on ethereum: ( txhash: %s )\n", hash.String())
}

func SetupOntAsset(invoker *eth.EInvoker) {
	if config.DefConfig.EthLockProxy == "" {
		panic(fmt.Errorf("EthLockProxy is blank"))
	}
	if config.DefConfig.EthOntx == "" {
		panic(fmt.Errorf("EthOntx is blank"))
	}
	if config.DefConfig.EthOngx == "" {
		panic(fmt.Errorf("EthOntx is blank"))
	}
	if config.DefConfig.EthOep4 == "" {
		panic(fmt.Errorf("EthOep4 is blank"))
	}
	if config.DefConfig.OntOep4 == "" {
		panic(fmt.Errorf("OntOep4 is blank"))
	}

	txs, err := invoker.BindOntAsset(config.DefConfig.EthLockProxy, config.DefConfig.EthOntx, config.DefConfig.EthOngx,
		config.DefConfig.EthOep4, config.DefConfig.OntOep4)
	if err != nil {
		panic(err)
	}
	hash1, hash2, hash3 := txs[0].Hash(), txs[1].Hash(), txs[2].Hash()
	fmt.Printf("ont/ong/oep4 binding tx on ontology: %s/%s/%s\n", hash1.String(), hash2.String(), hash3.String())

	hash4, hash5, hash6 := txs[3].Hash(), txs[4].Hash(), txs[5].Hash()
	fmt.Printf("ont/ong/oep4 binding tx on cosmos: %s/%s/%s\n", hash4.String(), hash5.String(), hash6.String())
}

func SetupETH(ethInvoker *eth.EInvoker) {
	ethNativeAddr := "0x0000000000000000000000000000000000000000"
	if config.DefConfig.OntEth == "" {
		panic(fmt.Errorf("ethx on ontology is blank"))
	}
	tx, err := ethInvoker.BindAssetHash(config.DefConfig.EthLockProxy, ethNativeAddr, config.DefConfig.OntEth, config.ONT_CHAIN_ID, 0)
	if err != nil {
		panic(fmt.Errorf("SetupETH2ONT, failed to bind asset hash: %v", err))
	}
	hash := tx.Hash()
	fmt.Printf("binding ethx of ontology on ethereum: ( txhash: %s )\n", hash.String())

	tx, err = ethInvoker.BindAssetHash(config.DefConfig.EthLockProxy, ethNativeAddr, config.CM_ETHX, int(config.DefConfig.CMCrossChainId), 0)
	if err != nil {
		panic(fmt.Errorf("SetupETH2ONT, failed to bind asset hash: %v", err))
	}
	hash = tx.Hash()
	fmt.Printf("binding ethx of cosmos on ethereum: ( txhash: %s )\n", hash.String())
}

func SetOtherLockProxy(invoker *eth.EInvoker) {
	auth, contract, err := invoker.MakeLockProxy(config.DefConfig.EthLockProxy)
	if err != nil {
		panic(fmt.Errorf("failed to MakeLockProxy: %v", err))
	}
	if config.DefConfig.OntLockProxy == "" {
		panic(fmt.Errorf("OntLockProxy is blank"))
	}
	other, err := common.AddressFromHexString(config.DefConfig.OntLockProxy)
	if err != nil {
		panic(fmt.Errorf("failed to AddressFromHexString: %v", err))
	}
	tx, err := contract.BindProxyHash(auth, config.ONT_CHAIN_ID, other[:])
	if err != nil {
		panic(fmt.Errorf("failed to bind proxy: %"))
	}
	hash := tx.Hash()
	invoker.ETHUtil.WaitTransactionConfirm(hash)
	fmt.Printf("binding ont proxy: ( txhash: %s )\n", hash.String())

	auth, err = invoker.MakeSmartContractAuth()
	if err != nil {
		panic(fmt.Errorf("failed to get auth: %v", err))
	}
	if config.DefConfig.CMLockProxy == "" {
		panic(fmt.Errorf("COSMOS lockproxy is blank"))
	}
	raw, err := hex.DecodeString(config.DefConfig.CMLockProxy)
	if err != nil {
		panic(fmt.Errorf("failed to decode: %v", err))
	}
	tx, err = contract.BindProxyHash(auth, config.DefConfig.CMCrossChainId, raw)
	if err != nil {
		panic(fmt.Errorf("failed to bind COSMOS proxy: %v", err))
	}
	hash = tx.Hash()
	invoker.ETHUtil.WaitTransactionConfirm(hash)
	fmt.Printf("binding cosmos proxy: ( txhash: %s )\n", hash.String())
}

func SetUpLockProxyForETH() {
	invoker := eth.NewEInvoker()
	SetupETH(invoker)
	SetupERC20(invoker)
	SetupOntAsset(invoker)
	SetOtherLockProxy(invoker)
}
