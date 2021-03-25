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
	"github.com/joeqian10/neo-gogogo/helper"
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
		DeployETHSmartContract()
	case "setup":
		SetUpEthContracts()
	}
}

func DeployETHSmartContract() {
	invoker := eth.NewEInvoker(config.DefConfig.O3ChainID)
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
		eccdAddr = common2.HexToAddress(config.DefConfig.O3Eccd)
		eccmAddr = common2.HexToAddress(config.DefConfig.O3Eccm)
		eccmpAddr = common2.HexToAddress(config.DefConfig.O3Eccmp)
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
	fmt.Println("o3 erc20:", erc20Addr.Hex())
	fmt.Println("o3 ope4:", oep4Addr.Hex())
	fmt.Println("o3 eccd address:", eccdAddr.Hex())
	fmt.Println("o3 eccm address:", eccmAddr.Hex())
	fmt.Println("o3 eccmp address:", eccmpAddr.Hex())
	fmt.Println("o3 lock proxy address: ", lockProxyAddr.Hex())
	fmt.Println("o3 ongx address: ", ongxAddr.Hex())
	fmt.Println("o3 ontx proxy address: ", ontxAddr.Hex())
	fmt.Println("==================================================================")

	config.DefConfig.O3Erc20 = erc20Addr.Hex()
	config.DefConfig.O3Oep4 = oep4Addr.Hex()
	config.DefConfig.O3Eccd = eccdAddr.Hex()
	config.DefConfig.O3Eccm = eccmAddr.Hex()
	config.DefConfig.O3Eccmp = eccmpAddr.Hex()
	config.DefConfig.O3LockProxy = lockProxyAddr.Hex()
	config.DefConfig.O3Ongx = ongxAddr.Hex()
	config.DefConfig.O3Ontx = ontxAddr.Hex()

	if err := config.DefConfig.Save(ethConfFile); err != nil {
		panic(fmt.Errorf("failed to save config, you better save it youself: %v", err))
	}
}

func SetupBep20(ethInvoker *eth.EInvoker) {
	if config.DefConfig.OntBep20 != "" {
		bindTx, err := ethInvoker.BindAssetHash(config.DefConfig.O3LockProxy, config.DefConfig.O3Erc20,
			config.DefConfig.OntBep20, config.DefConfig.OntChainID, 0)
		if err != nil {
			panic(fmt.Errorf("SetupBep20ToONT, failed to BindAssetHash: %v", err))
		}
		ethInvoker.ETHUtil.WaitTransactionConfirm(bindTx.Hash())
		hash := bindTx.Hash()
		fmt.Printf("binding bep20 of ontology on o3: ( txhash: %s )\n", hash.String())
	}

	bindTx, err := ethInvoker.BindAssetHash(config.DefConfig.O3LockProxy, config.DefConfig.O3Erc20,
		config.CM_BEP20, config.DefConfig.CMCrossChainId, 0)
	if err != nil {
		panic(fmt.Errorf("SetupBep20ToONT, failed to BindAssetHash: %v", err))
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(bindTx.Hash())
	hash := bindTx.Hash()
	fmt.Printf("binding bep20 of cosmos on o3: ( txhash: %s )\n", hash.String())
}

func SetupWBTC(ethInvoker *eth.EInvoker) {
	bindTx, err := ethInvoker.BindAssetHash(config.DefConfig.O3LockProxy, config.DefConfig.O3WBTC,
		config.DefConfig.OntWBTC, config.DefConfig.OntChainID, 0)
	if err != nil {
		panic(fmt.Errorf("SetupWBTC, failed to BindAssetHash: %v", err))
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(bindTx.Hash())
	hash := bindTx.Hash()
	fmt.Printf("binding WBTC of ontology on o3: ( txhash: %s )\n", hash.String())
}

func SetupDAI(ethInvoker *eth.EInvoker) {
	bindTx, err := ethInvoker.BindAssetHash(config.DefConfig.O3LockProxy, config.DefConfig.O3Dai,
		config.DefConfig.OntDai, config.DefConfig.OntChainID, 0)
	if err != nil {
		panic(fmt.Errorf("SetupDAI, failed to BindAssetHash: %v", err))
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(bindTx.Hash())
	hash := bindTx.Hash()
	fmt.Printf("binding DAI of ontology on o3: ( txhash: %s )\n", hash.String())
}

func SetupUSDT(ethInvoker *eth.EInvoker) {
	bindTx, err := ethInvoker.BindAssetHash(config.DefConfig.O3LockProxy, config.DefConfig.O3USDT,
		config.DefConfig.OntUSDT, config.DefConfig.OntChainID, 0)
	if err != nil {
		panic(fmt.Errorf("SetupUSDT, failed to BindAssetHash: %v", err))
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(bindTx.Hash())
	hash := bindTx.Hash()
	fmt.Printf("binding USDT of ontology on o3: ( txhash: %s )\n", hash.String())
}

func SetupUSDC(ethInvoker *eth.EInvoker) {
	bindTx, err := ethInvoker.BindAssetHash(config.DefConfig.O3LockProxy, config.DefConfig.O3USDC,
		config.DefConfig.OntUSDC, config.DefConfig.OntChainID, 0)
	if err != nil {
		panic(fmt.Errorf("SetupUSDC, failed to BindAssetHash: %v", err))
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(bindTx.Hash())
	hash := bindTx.Hash()
	fmt.Printf("binding USDC of ontology on o3: ( txhash: %s )\n", hash.String())
}

func SetupOntAsset(invoker *eth.EInvoker) {
	if config.DefConfig.O3LockProxy == "" {
		panic(fmt.Errorf("BscLockProxy is blank"))
	}
	if config.DefConfig.O3Ontx == "" {
		panic(fmt.Errorf("BscOntx is blank"))
	}
	if config.DefConfig.O3Ongx == "" {
		panic(fmt.Errorf("BscOngx is blank"))
	}
	if config.DefConfig.O3Oep4 == "" {
		panic(fmt.Errorf("BscOep4 is blank"))
	}
	if config.DefConfig.OntOep4 == "" {
		panic(fmt.Errorf("OntOep4 is blank"))
	}

	txs, err := invoker.BindOntAsset(config.DefConfig.O3LockProxy, config.DefConfig.O3Ontx, config.DefConfig.O3Ongx,
		config.DefConfig.O3Oep4, config.DefConfig.OntOep4)
	if err != nil {
		panic(err)
	}
	hash1, hash2, hash3 := txs[0].Hash(), txs[1].Hash(), txs[2].Hash()
	fmt.Printf("ont/ong/oep4 binding tx on ontology: %s/%s/%s\n", hash1.String(), hash2.String(), hash3.String())

	hash4, hash5, hash6 := txs[3].Hash(), txs[4].Hash(), txs[5].Hash()
	fmt.Printf("ont/ong/oep4 binding tx on cosmos: %s/%s/%s\n", hash4.String(), hash5.String(), hash6.String())
}

func SetupBnb(ethInvoker *eth.EInvoker) {
	ethNativeAddr := "0x0000000000000000000000000000000000000000"
	if config.DefConfig.OntBnb != "" {
		tx, err := ethInvoker.BindAssetHash(config.DefConfig.O3LockProxy, ethNativeAddr, config.DefConfig.OntBnb, config.DefConfig.OntChainID, 0)
		if err != nil {
			panic(fmt.Errorf("SetupBnb2ONT, failed to bind asset hash: %v", err))
		}
		hash := tx.Hash()
		fmt.Printf("binding bnbx of ontology on o3: ( txhash: %s )\n", hash.String())
	}

	if config.DefConfig.EthBnb != "" {
		tx, err := ethInvoker.BindAssetHash(config.DefConfig.O3LockProxy, ethNativeAddr, config.DefConfig.EthBnb, config.DefConfig.EthChainID, 0)
		if err != nil {
			panic(fmt.Errorf("SetupBnb2ONT, failed to bind asset hash: %v", err))
		}
		hash := tx.Hash()
		fmt.Printf("binding bnb of o3 on ethereum: ( txhash: %s )\n", hash.String())
	}
	if config.DefConfig.NeoBnb != "" {
		tx, err := ethInvoker.BindAssetHash(config.DefConfig.O3LockProxy, ethNativeAddr, config.DefConfig.NeoBnb, config.DefConfig.NeoChainID, 0)
		if err != nil {
			panic(fmt.Errorf("SetupBnb2Neo, failed to bind asset hash: %v", err))
		}
		hash := tx.Hash()
		fmt.Printf("binding bnb of o3 on neo: ( txhash: %s )\n", hash.String())
	}

	tx, err := ethInvoker.BindAssetHash(config.DefConfig.O3LockProxy, ethNativeAddr, config.CM_BNBX, config.DefConfig.CMCrossChainId, 0)
	if err != nil {
		panic(fmt.Errorf("SetupBnb2COSMOS, failed to bind asset hash: %v", err))
	}
	hash := tx.Hash()
	fmt.Printf("binding bnbx of cosmos on o3: ( txhash: %s )\n", hash.String())

	tx, err = ethInvoker.BindAssetHash(config.DefConfig.O3LockProxy, ethNativeAddr, ethNativeAddr, config.DefConfig.BscChainID, 0)
	if err != nil {
		panic(fmt.Errorf("BindAssetHash, failed to bind asset hash: %v", err))
	}
	hash = tx.Hash()
	fmt.Printf("binding bnb of o3 on bsc: ( txhash: %s )\n", hash.String())
}

func SetOtherLockProxy(invoker *eth.EInvoker) {
	_, contract, err := invoker.MakeLockProxy(config.DefConfig.O3LockProxy)
	if err != nil {
		panic(fmt.Errorf("failed to MakeLockProxy: %v", err))
	}
	if config.DefConfig.OntLockProxy != "" {
		auth, err := invoker.MakeSmartContractAuth()
		if err != nil {
			panic(fmt.Errorf("failed to get auth: %v", err))
		}
		other, err := common.AddressFromHexString(config.DefConfig.OntLockProxy)
		if err != nil {
			panic(fmt.Errorf("failed to AddressFromHexString: %v", err))
		}
		tx, err := contract.BindProxyHash(auth, config.DefConfig.OntChainID, other[:])
		if err != nil {
			panic(fmt.Errorf("failed to bind proxy: %"))
		}
		hash := tx.Hash()
		invoker.ETHUtil.WaitTransactionConfirm(hash)
		fmt.Printf("binding ont proxy: ( txhash: %s )\n", hash.String())
	}

	if config.DefConfig.CMLockProxy != "" {
		auth, err := invoker.MakeSmartContractAuth()
		if err != nil {
			panic(fmt.Errorf("failed to get auth: %v", err))
		}
		raw, err := hex.DecodeString(config.DefConfig.CMLockProxy)
		if err != nil {
			panic(fmt.Errorf("failed to decode: %v", err))
		}
		tx, err := contract.BindProxyHash(auth, config.DefConfig.CMCrossChainId, raw)
		if err != nil {
			panic(fmt.Errorf("failed to bind COSMOS proxy: %v", err))
		}
		hash := tx.Hash()
		invoker.ETHUtil.WaitTransactionConfirm(hash)
		fmt.Printf("binding cosmos proxy: ( txhash: %s )\n", hash.String())
	}

	if config.DefConfig.BscLockProxy != "" {
		auth, err := invoker.MakeSmartContractAuth()
		if err != nil {
			panic(fmt.Errorf("failed to get auth: %v", err))
		}
		other := common2.HexToAddress(config.DefConfig.BscLockProxy)
		tx, err := contract.BindProxyHash(auth, config.DefConfig.BscChainID, other[:])
		if err != nil {
			panic(fmt.Errorf("failed to bind proxy: %v", err))
		}
		hash := tx.Hash()
		invoker.ETHUtil.WaitTransactionConfirm(hash)
		fmt.Printf("binding bsc proxy: ( txhash: %s )\n", hash.String())
	}

	if config.DefConfig.EthLockProxy != "" {
		auth, err := invoker.MakeSmartContractAuth()
		if err != nil {
			panic(fmt.Errorf("failed to get auth: %v", err))
		}
		other := common2.HexToAddress(config.DefConfig.EthLockProxy)
		tx, err := contract.BindProxyHash(auth, config.DefConfig.EthChainID, other[:])
		if err != nil {
			panic(fmt.Errorf("failed to bind proxy: %v", err))
		}
		hash := tx.Hash()
		invoker.ETHUtil.WaitTransactionConfirm(hash)
		fmt.Printf("binding eth proxy: ( txhash: %s )\n", hash.String())
	}

	if config.DefConfig.NeoLockProxy != "" {
		auth, err := invoker.MakeSmartContractAuth()
		if err != nil {
			panic(fmt.Errorf("failed to get auth: %v", err))
		}
		other, err := helper.UInt160FromString(config.DefConfig.NeoLockProxy)
		if err != nil {
			panic(fmt.Errorf("UInt160FromString error: %v", err))
		}
		tx, err := contract.BindProxyHash(auth, config.DefConfig.NeoChainID, other[:])
		if err != nil {
			panic(fmt.Errorf("failed to bind proxy: %v", err))
		}
		hash := tx.Hash()
		invoker.ETHUtil.WaitTransactionConfirm(hash)
		fmt.Printf("binding neo proxy: ( txhash: %s )\n", hash.String())
	}
}

func SetUpEthContracts() {
	invoker := eth.NewEInvoker(config.DefConfig.O3ChainID)
	SetupBnb(invoker)
	if config.DefConfig.O3Erc20 != "" {
		SetupBep20(invoker)
	}
	if config.DefConfig.OntLockProxy != "" {
		SetupOntAsset(invoker)
	}
	if config.DefConfig.O3WBTC != "" {
		SetupWBTC(invoker)
	}
	if config.DefConfig.O3Dai != "" {
		SetupDAI(invoker)
	}
	if config.DefConfig.O3USDT != "" {
		SetupUSDT(invoker)
	}

	//SetupUSDC(invoker)
	SetOtherLockProxy(invoker)
}
