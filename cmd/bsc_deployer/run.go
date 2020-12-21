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
		DeployETHSmartContract()
	case "setup":
		SetUpEthContracts()
	}
}

func DeployETHSmartContract() {
	invoker := eth.NewEInvoker(config.DefConfig.BscChainID)
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
		eccdAddr = common2.HexToAddress(config.DefConfig.BscEccd)
		eccmAddr = common2.HexToAddress(config.DefConfig.BscEccm)
		eccmpAddr = common2.HexToAddress(config.DefConfig.BscEccmp)
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
	fmt.Println("bsc erc20:", erc20Addr.Hex())
	fmt.Println("bsc ope4:", oep4Addr.Hex())
	fmt.Println("bsc eccd address:", eccdAddr.Hex())
	fmt.Println("bsc eccm address:", eccmAddr.Hex())
	fmt.Println("bsc eccmp address:", eccmpAddr.Hex())
	fmt.Println("bsc lock proxy address: ", lockProxyAddr.Hex())
	fmt.Println("bsc ongx address: ", ongxAddr.Hex())
	fmt.Println("bsc ontx proxy address: ", ontxAddr.Hex())
	fmt.Println("==================================================================")

	config.DefConfig.Bep20 = erc20Addr.Hex()
	config.DefConfig.BscOep4 = oep4Addr.Hex()
	config.DefConfig.BscEccd = eccdAddr.Hex()
	config.DefConfig.BscEccm = eccmAddr.Hex()
	config.DefConfig.BscEccmp = eccmpAddr.Hex()
	config.DefConfig.BscLockProxy = lockProxyAddr.Hex()
	config.DefConfig.BscOngx = ongxAddr.Hex()
	config.DefConfig.BscOntx = ontxAddr.Hex()

	if err := config.DefConfig.Save(ethConfFile); err != nil {
		panic(fmt.Errorf("failed to save config, you better save it youself: %v", err))
	}
}

func SetupBep20(ethInvoker *eth.EInvoker) {
	if config.DefConfig.OntBep20 != "" {
		bindTx, err := ethInvoker.BindAssetHash(config.DefConfig.BscLockProxy, config.DefConfig.Bep20,
			config.DefConfig.OntBep20, config.DefConfig.OntChainID, 0)
		if err != nil {
			panic(fmt.Errorf("SetupBep20ToONT, failed to BindAssetHash: %v", err))
		}
		ethInvoker.ETHUtil.WaitTransactionConfirm(bindTx.Hash())
		hash := bindTx.Hash()
		fmt.Printf("binding bep20 of ontology on bsc: ( txhash: %s )\n", hash.String())
	}

	bindTx, err := ethInvoker.BindAssetHash(config.DefConfig.BscLockProxy, config.DefConfig.Bep20,
		config.CM_BEP20, config.DefConfig.CMCrossChainId, 0)
	if err != nil {
		panic(fmt.Errorf("SetupBep20ToONT, failed to BindAssetHash: %v", err))
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(bindTx.Hash())
	hash := bindTx.Hash()
	fmt.Printf("binding bep20 of cosmos on bsc: ( txhash: %s )\n", hash.String())
}

func SetupWBTC(ethInvoker *eth.EInvoker) {
	bindTx, err := ethInvoker.BindAssetHash(config.DefConfig.BscLockProxy, config.DefConfig.BscWBTC,
		config.DefConfig.OntWBTC, config.DefConfig.OntChainID, 0)
	if err != nil {
		panic(fmt.Errorf("SetupWBTC, failed to BindAssetHash: %v", err))
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(bindTx.Hash())
	hash := bindTx.Hash()
	fmt.Printf("binding WBTC of ontology on bsc: ( txhash: %s )\n", hash.String())
}

func SetupDAI(ethInvoker *eth.EInvoker) {
	bindTx, err := ethInvoker.BindAssetHash(config.DefConfig.BscLockProxy, config.DefConfig.BscDai,
		config.DefConfig.OntDai, config.DefConfig.OntChainID, 0)
	if err != nil {
		panic(fmt.Errorf("SetupDAI, failed to BindAssetHash: %v", err))
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(bindTx.Hash())
	hash := bindTx.Hash()
	fmt.Printf("binding DAI of ontology on bsc: ( txhash: %s )\n", hash.String())
}

func SetupUSDT(ethInvoker *eth.EInvoker) {
	bindTx, err := ethInvoker.BindAssetHash(config.DefConfig.BscLockProxy, config.DefConfig.BscUSDT,
		config.DefConfig.OntUSDT, config.DefConfig.OntChainID, 0)
	if err != nil {
		panic(fmt.Errorf("SetupUSDT, failed to BindAssetHash: %v", err))
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(bindTx.Hash())
	hash := bindTx.Hash()
	fmt.Printf("binding USDT of ontology on bsc: ( txhash: %s )\n", hash.String())
}

func SetupUSDC(ethInvoker *eth.EInvoker) {
	bindTx, err := ethInvoker.BindAssetHash(config.DefConfig.BscLockProxy, config.DefConfig.BscUSDC,
		config.DefConfig.OntUSDC, config.DefConfig.OntChainID, 0)
	if err != nil {
		panic(fmt.Errorf("SetupUSDC, failed to BindAssetHash: %v", err))
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(bindTx.Hash())
	hash := bindTx.Hash()
	fmt.Printf("binding USDC of ontology on ethereum: ( txhash: %s )\n", hash.String())
}

func SetupOntAsset(invoker *eth.EInvoker) {
	if config.DefConfig.BscLockProxy == "" {
		panic(fmt.Errorf("BscLockProxy is blank"))
	}
	if config.DefConfig.BscOntx == "" {
		panic(fmt.Errorf("BscOntx is blank"))
	}
	if config.DefConfig.BscOngx == "" {
		panic(fmt.Errorf("BscOngx is blank"))
	}
	if config.DefConfig.BscOep4 == "" {
		panic(fmt.Errorf("BscOep4 is blank"))
	}
	if config.DefConfig.OntOep4 == "" {
		panic(fmt.Errorf("OntOep4 is blank"))
	}

	txs, err := invoker.BindOntAsset(config.DefConfig.BscLockProxy, config.DefConfig.BscOntx, config.DefConfig.BscOngx,
		config.DefConfig.BscOep4, config.DefConfig.OntOep4)
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
		tx, err := ethInvoker.BindAssetHash(config.DefConfig.BscLockProxy, ethNativeAddr, config.DefConfig.OntBnb, config.DefConfig.OntChainID, 0)
		if err != nil {
			panic(fmt.Errorf("SetupBnb2ONT, failed to bind asset hash: %v", err))
		}
		hash := tx.Hash()
		fmt.Printf("binding bnbx of ontology on bsc: ( txhash: %s )\n", hash.String())
	}

	tx, err := ethInvoker.BindAssetHash(config.DefConfig.BscLockProxy, ethNativeAddr, config.CM_BNBX, config.DefConfig.CMCrossChainId, 0)
	if err != nil {
		panic(fmt.Errorf("SetupBnb2COSMOS, failed to bind asset hash: %v", err))
	}
	hash := tx.Hash()
	fmt.Printf("binding bnbx of cosmos on bsc: ( txhash: %s )\n", hash.String())

	tx, err = ethInvoker.BindAssetHash(config.DefConfig.BscLockProxy, ethNativeAddr, ethNativeAddr, config.DefConfig.BscChainID, 0)
	if err != nil {
		panic(fmt.Errorf("BindAssetHash, failed to bind asset hash: %v", err))
	}
	hash = tx.Hash()
	fmt.Printf("binding bnb of bsc on bsc: ( txhash: %s )\n", hash.String())
}

func SetOtherLockProxy(invoker *eth.EInvoker) {
	_, contract, err := invoker.MakeLockProxy(config.DefConfig.BscLockProxy)
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
}

func SetUpEthContracts() {
	invoker := eth.NewEInvoker(config.DefConfig.BscChainID)
	SetupBnb(invoker)
	if config.DefConfig.Bep20 != "" {
		SetupBep20(invoker)
	}
	if config.DefConfig.OntLockProxy != "" {
		SetupOntAsset(invoker)
	}
	if config.DefConfig.EthWBTC != "" {
		SetupWBTC(invoker)
	}
	if config.DefConfig.EthDai != "" {
		SetupDAI(invoker)
	}
	if config.DefConfig.EthUSDT != "" {
		SetupUSDT(invoker)
	}
	//SetupUSDC(invoker)
	SetOtherLockProxy(invoker)
}
