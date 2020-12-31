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
	"github.com/polynetwork/eth-contracts/go_abi/erc20_abi"

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
}

func main() {
	flag.Parse()
	err := config.DefConfig.Init(ethConfFile)
	if err != nil {
		panic(err)
	}

	switch fnEth {
	case "deploy":
		DeployHecoSmartContract()
	case "setup":
		SetUpHecoContracts()
	}
}

func DeployHecoSmartContract() {
	invoker := eth.NewEInvoker(config.DefConfig.HecoChainID)
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
		eccdAddr = common2.HexToAddress(config.DefConfig.HecoEccd)
		eccmAddr = common2.HexToAddress(config.DefConfig.HecoEccm)
		eccmpAddr = common2.HexToAddress(config.DefConfig.HecoEccmp)
	}

	var lockproxyAddrHex string = config.DefConfig.HecoLockProxy
	lockProxyAddr := common2.HexToAddress(lockproxyAddrHex)
	if lockproxyAddrHex == "" {
		lockProxyAddr, _, err = invoker.DeployLockProxyContract(eccmpAddr)
		if err != nil {
			panic(err)
		}
		lockproxyAddrHex = lockProxyAddr.Hex()
	}

	hrc20AddrHex := config.DefConfig.HecoHrc20
	hrc20Addr := common2.HexToAddress(hrc20AddrHex)
	if hrc20AddrHex == "" {
		var hrc20 *erc20_abi.ERC20Template
		hrc20Addr, hrc20, err = invoker.DeployERC20()
		if err != nil {
			panic(err)
		}
		total, err := hrc20.TotalSupply(nil)
		if err != nil {
			panic(fmt.Errorf("failed to get total supply for hrc20: %v", err))
		}
		auth, _ := invoker.MakeSmartContractAuth()
		tx, err := hrc20.Approve(auth, lockProxyAddr, total)
		if err != nil {
			panic(fmt.Errorf("failed to approve hrc20 to lockproxy: %v", err))
		}
		invoker.ETHUtil.WaitTransactionConfirm(tx.Hash())
	}

	oep4AddrHex := config.DefConfig.HecoOep4
	oep4Addr := common2.HexToAddress(oep4AddrHex)
	if oep4AddrHex == "" {
		oep4Addr, _, err = invoker.DeployOEP4(lockproxyAddrHex)
		if err != nil {
			panic(err)
		}
	}

	ongxAddrHex := config.DefConfig.HecoOngx
	ongxAddr := common2.HexToAddress(ongxAddrHex)
	if ongxAddrHex == "" {
		ongxAddr, _, err = invoker.DeployONGXContract(lockproxyAddrHex)
		if err != nil {
			panic(err)
		}
	}

	ontxAddrHex := config.DefConfig.HecoOntx
	ontxAddr := common2.HexToAddress(ontxAddrHex)
	if ontxAddrHex == "" {
		ontxAddr, _, err = invoker.DeployONTXContract(lockproxyAddrHex)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("=============================Heco info=============================")
	fmt.Println("heco hrc20                :", hrc20Addr.Hex())
	fmt.Println("heco ope4                 :", oep4Addr.Hex())
	fmt.Println("heco eccd address         :", eccdAddr.Hex())
	fmt.Println("heco eccm address         :", eccmAddr.Hex())
	fmt.Println("heco eccmp address        :", eccmpAddr.Hex())
	fmt.Println("heco lock proxy address   :", lockProxyAddr.Hex())
	fmt.Println("heco ongx address         :", ongxAddr.Hex())
	fmt.Println("heco ontx proxy address   :", ontxAddr.Hex())
	fmt.Println("===================================================================")

	config.DefConfig.HecoHrc20 = hrc20Addr.Hex()
	config.DefConfig.HecoOep4 = oep4Addr.Hex()
	config.DefConfig.HecoEccd = eccdAddr.Hex()
	config.DefConfig.HecoEccm = eccmAddr.Hex()
	config.DefConfig.HecoEccmp = eccmpAddr.Hex()
	config.DefConfig.HecoLockProxy = lockProxyAddr.Hex()
	config.DefConfig.HecoOngx = ongxAddr.Hex()
	config.DefConfig.HecoOntx = ontxAddr.Hex()

	if err := config.DefConfig.Save(ethConfFile); err != nil {
		panic(fmt.Errorf("failed to save config, you better save it youself: %v", err))
	}
}

func SetupHrc20(ethInvoker *eth.EInvoker) {
	if config.DefConfig.OntHrc20 != "" {
		bindTx, err := ethInvoker.BindAssetHash(config.DefConfig.HecoLockProxy, config.DefConfig.HecoHrc20,
			config.DefConfig.OntHrc20, config.DefConfig.OntChainID, 0)
		if err != nil {
			panic(fmt.Errorf("SetupHrc20ToONT, failed to BindAssetHash: %v", err))
		}
		ethInvoker.ETHUtil.WaitTransactionConfirm(bindTx.Hash())
		hash := bindTx.Hash()
		fmt.Printf("binding hrc20 of ontology on bsc: ( txhash: %s )\n", hash.String())
	}
	if config.DefConfig.HecoHrc20 != "" && config.DefConfig.CMCrossChainId != 0 {
		bindTx, err := ethInvoker.BindAssetHash(config.DefConfig.HecoLockProxy, config.DefConfig.HecoHrc20,
			config.CM_HRC20, config.DefConfig.CMCrossChainId, 0)
		if err != nil {
			panic(fmt.Errorf("SetupHrc20ToONT, failed to BindAssetHash: %v", err))
		}
		ethInvoker.ETHUtil.WaitTransactionConfirm(bindTx.Hash())
		hash := bindTx.Hash()
		fmt.Printf("binding hrc20 of cosmos on heco: ( txhash: %s )\n", hash.String())
	}

}

func SetupErc20(ethInvoker *eth.EInvoker) {
	if config.DefConfig.OntChainID != 0 && config.DefConfig.OntHrc20 != "" {
		bindTx, err := ethInvoker.BindAssetHash(config.DefConfig.HecoLockProxy, config.DefConfig.HecoErc20,
			config.DefConfig.OntHrc20, config.DefConfig.OntChainID, 0)
		if err != nil {
			panic(fmt.Errorf("SetupErc20ToONT, failed to BindAssetHash: %v", err))
		}
		ethInvoker.ETHUtil.WaitTransactionConfirm(bindTx.Hash())
		hash := bindTx.Hash()
		fmt.Printf("binding erc20 of ontology on bsc: ( txhash: %s )\n", hash.String())
	}
	if config.DefConfig.CMCrossChainId != 0 {
		bindTx, err := ethInvoker.BindAssetHash(config.DefConfig.HecoLockProxy, config.DefConfig.HecoErc20,
			config.CM_HRC20, config.DefConfig.CMCrossChainId, 0)
		if err != nil {
			panic(fmt.Errorf("SetupErc20ToONT, failed to BindAssetHash: %v", err))
		}
		ethInvoker.ETHUtil.WaitTransactionConfirm(bindTx.Hash())
		hash := bindTx.Hash()
		fmt.Printf("binding erc20 of cosmos on heco: ( txhash: %s )\n", hash.String())
	}

	if config.DefConfig.EthChainID != 0 && config.DefConfig.EthErc20 != "" {
		bindTx, err := ethInvoker.BindAssetHash(config.DefConfig.HecoLockProxy, config.DefConfig.HecoErc20,
			config.DefConfig.EthErc20, config.DefConfig.EthChainID, 0)
		if err != nil {
			panic(fmt.Errorf("SetupErc20ToONT, failed to BindAssetHash: %v", err))
		}
		ethInvoker.ETHUtil.WaitTransactionConfirm(bindTx.Hash())
		hash := bindTx.Hash()
		fmt.Printf("binding erc20 of eth on heco: ( txhash: %s )\n", hash.String())
	}

}
func SetupWBTC(ethInvoker *eth.EInvoker) {
	bindTx, err := ethInvoker.BindAssetHash(config.DefConfig.HecoLockProxy, config.DefConfig.HecoWBTC,
		config.DefConfig.OntWBTC, config.DefConfig.OntChainID, 0)
	if err != nil {
		panic(fmt.Errorf("SetupWBTC, failed to BindAssetHash: %v", err))
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(bindTx.Hash())
	hash := bindTx.Hash()
	fmt.Printf("binding WBTC of ontology on heco: ( txhash: %s )\n", hash.String())
}

func SetupDAI(ethInvoker *eth.EInvoker) {
	bindTx, err := ethInvoker.BindAssetHash(config.DefConfig.HecoLockProxy, config.DefConfig.HecoDai,
		config.DefConfig.OntDai, config.DefConfig.OntChainID, 0)
	if err != nil {
		panic(fmt.Errorf("SetupDAI, failed to BindAssetHash: %v", err))
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(bindTx.Hash())
	hash := bindTx.Hash()
	fmt.Printf("binding DAI of ontology on heco: ( txhash: %s )\n", hash.String())
}

func SetupUSDT(ethInvoker *eth.EInvoker) {
	bindTx, err := ethInvoker.BindAssetHash(config.DefConfig.HecoLockProxy, config.DefConfig.HecoUSDT,
		config.DefConfig.OntUSDT, config.DefConfig.OntChainID, 0)
	if err != nil {
		panic(fmt.Errorf("SetupUSDT, failed to BindAssetHash: %v", err))
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(bindTx.Hash())
	hash := bindTx.Hash()
	fmt.Printf("binding USDT of ontology on heco: ( txhash: %s )\n", hash.String())
}

func SetupUSDC(ethInvoker *eth.EInvoker) {
	bindTx, err := ethInvoker.BindAssetHash(config.DefConfig.HecoLockProxy, config.DefConfig.HecoUSDC,
		config.DefConfig.OntUSDC, config.DefConfig.OntChainID, 0)
	if err != nil {
		panic(fmt.Errorf("SetupUSDC, failed to BindAssetHash: %v", err))
	}
	ethInvoker.ETHUtil.WaitTransactionConfirm(bindTx.Hash())
	hash := bindTx.Hash()
	fmt.Printf("binding USDC of ontology on heco: ( txhash: %s )\n", hash.String())
}

func SetupOntAsset(invoker *eth.EInvoker) {
	if config.DefConfig.HecoLockProxy == "" {
		panic(fmt.Errorf("HecoLockProxy is blank"))
	}
	if config.DefConfig.HecoOntx == "" {
		panic(fmt.Errorf("HecoOntx is blank"))
	}
	if config.DefConfig.HecoOngx == "" {
		panic(fmt.Errorf("HecoOngx is blank"))
	}
	if config.DefConfig.HecoOep4 == "" {
		panic(fmt.Errorf("HecoOep4 is blank"))
	}
	if config.DefConfig.OntOep4 == "" {
		panic(fmt.Errorf("OntOep4 is blank"))
	}

	txs, err := invoker.BindOntAsset(config.DefConfig.HecoLockProxy, config.DefConfig.HecoOntx, config.DefConfig.HecoOngx,
		config.DefConfig.HecoOep4, config.DefConfig.OntOep4)
	if err != nil {
		panic(err)
	}
	hash1, hash2, hash3 := txs[0].Hash(), txs[1].Hash(), txs[2].Hash()
	fmt.Printf("ont/ong/oep4 binding tx on ontology: %s/%s/%s\n", hash1.String(), hash2.String(), hash3.String())

	hash4, hash5, hash6 := txs[3].Hash(), txs[4].Hash(), txs[5].Hash()
	fmt.Printf("ont/ong/oep4 binding tx on cosmos: %s/%s/%s\n", hash4.String(), hash5.String(), hash6.String())
}

func SetupHt(ethInvoker *eth.EInvoker) {
	ethNativeAddr := "0x0000000000000000000000000000000000000000"
	if config.DefConfig.OntHt != "" && config.DefConfig.OntChainID != 0 {
		tx, err := ethInvoker.BindAssetHash(config.DefConfig.HecoLockProxy, ethNativeAddr, config.DefConfig.OntHt, config.DefConfig.OntChainID, 0)
		if err != nil {
			panic(fmt.Errorf("SetupHt2ONT, failed to bind asset hash: %v", err))
		}
		hash := tx.Hash()
		fmt.Printf("binding htx of ontology on heco: ( txhash: %s )\n", hash.String())
	}

	if config.DefConfig.EthHt != "" && config.DefConfig.EthChainID != 0 {
		tx, err := ethInvoker.BindAssetHash(config.DefConfig.HecoLockProxy, ethNativeAddr, config.DefConfig.EthHt, config.DefConfig.EthChainID, 0)
		if err != nil {
			panic(fmt.Errorf("SetupHt2ONT, failed to bind asset hash: %v", err))
		}
		hash := tx.Hash()
		fmt.Printf("binding ht of ethereum on heco: ( txhash: %s )\n", hash.String())
	}
	if config.DefConfig.CMCrossChainId != 0 {
		tx, err := ethInvoker.BindAssetHash(config.DefConfig.HecoLockProxy, ethNativeAddr, config.CM_HTX, config.DefConfig.CMCrossChainId, 0)
		if err != nil {
			panic(fmt.Errorf("SetupHt2COSMOS, failed to bind asset hash: %v", err))
		}
		hash := tx.Hash()
		fmt.Printf("binding htx of cosmos on heco: ( txhash: %s )\n", hash.String())
	}

	tx, err := ethInvoker.BindAssetHash(config.DefConfig.HecoLockProxy, ethNativeAddr, ethNativeAddr, config.DefConfig.HecoChainID, 0)
	if err != nil {
		panic(fmt.Errorf("BindAssetHash, failed to bind asset hash: %v", err))
	}
	hash := tx.Hash()
	fmt.Printf("binding ht of heco on heco: ( txhash: %s )\n", hash.String())
}

func SetOtherLockProxy(invoker *eth.EInvoker) {
	_, contract, err := invoker.MakeLockProxy(config.DefConfig.HecoLockProxy)
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

	if config.DefConfig.HecoLockProxy != "" {
		auth, err := invoker.MakeSmartContractAuth()
		if err != nil {
			panic(fmt.Errorf("failed to get auth: %v", err))
		}
		other := common2.HexToAddress(config.DefConfig.HecoLockProxy)
		tx, err := contract.BindProxyHash(auth, config.DefConfig.HecoChainID, other[:])
		if err != nil {
			panic(fmt.Errorf("failed to bind proxy: %v", err))
		}
		hash := tx.Hash()
		invoker.ETHUtil.WaitTransactionConfirm(hash)
		fmt.Printf("binding Heco proxy: ( txhash: %s )\n", hash.String())
	}
}

func SetUpHecoContracts() {
	invoker := eth.NewEInvoker(config.DefConfig.HecoChainID)
	SetupHt(invoker)
	if config.DefConfig.HecoHrc20 != "" {
		SetupHrc20(invoker)
	}
	if config.DefConfig.HecoErc20 != "" {
		SetupErc20(invoker)
	}
	if config.DefConfig.OntLockProxy != "" {
		SetupOntAsset(invoker)
	}
	if config.DefConfig.HecoWBTC != "" {
		SetupWBTC(invoker)
	}
	if config.DefConfig.HecoDai != "" {
		SetupDAI(invoker)
	}
	if config.DefConfig.HecoUSDT != "" {
		SetupUSDT(invoker)
	}

	//SetupUSDC(invoker)
	SetOtherLockProxy(invoker)
}
