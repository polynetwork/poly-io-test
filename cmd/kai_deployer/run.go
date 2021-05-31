package main

import (
	"encoding/hex"
	"flag"
	"fmt"

	common2 "github.com/ethereum/go-ethereum/common"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/ontio/ontology/common"
	"github.com/polynetwork/poly-io-test/chains/kai"
	"github.com/polynetwork/poly-io-test/config"
)

var (
	fnKai        string
	kaiConfFile  string
	eccmRedeploy int
)

func init() {
	flag.StringVar(&fnKai, "func", "deploy", "choose function to run: deploy or setup")
	flag.StringVar(&kaiConfFile, "conf", "./config.json", "config file path")
	flag.IntVar(&eccmRedeploy, "redeploy_eccm", 1, "redeploy eccd, eccm and eccmp or not")
	flag.Parse()
}

func main() {
	err := config.DefConfig.Init(kaiConfFile)
	if err != nil {
		panic(err)
	}

	switch fnKai {
	case "deploy":
		DeploySmartContract()
	case "setup":
		SetUpContracts()
	}
}

func DeploySmartContract() {
	invoker := kai.NewInvoker(config.DefConfig.KaiChainID)
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
		eccdAddr = common2.HexToAddress(config.DefConfig.KaiEccd)
		eccmAddr = common2.HexToAddress(config.DefConfig.KaiEccm)
		eccmpAddr = common2.HexToAddress(config.DefConfig.KaiEccmp)
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
	invoker.Client().WaitTransactionConfirm(tx.Hash())

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
	fmt.Println("kai krc20:", erc20Addr.Hex())
	fmt.Println("kai ope4:", oep4Addr.Hex())
	fmt.Println("kai eccd address:", eccdAddr.Hex())
	fmt.Println("kai eccm address:", eccmAddr.Hex())
	fmt.Println("kai eccmp address:", eccmpAddr.Hex())
	fmt.Println("kai lock proxy address: ", lockProxyAddr.Hex())
	fmt.Println("kai ongx address: ", ongxAddr.Hex())
	fmt.Println("kai ontx proxy address: ", ontxAddr.Hex())
	fmt.Println("==================================================================")

	config.DefConfig.Krc20 = erc20Addr.Hex()
	config.DefConfig.KaiOep4 = oep4Addr.Hex()
	config.DefConfig.KaiEccd = eccdAddr.Hex()
	config.DefConfig.KaiEccm = eccmAddr.Hex()
	config.DefConfig.KaiEccmp = eccmpAddr.Hex()
	config.DefConfig.KaiLockProxy = lockProxyAddr.Hex()
	config.DefConfig.KaiOngx = ongxAddr.Hex()
	config.DefConfig.KaiOntx = ontxAddr.Hex()

	if err := config.DefConfig.Save(kaiConfFile); err != nil {
		panic(fmt.Errorf("failed to save config, you better save it youself: %v", err))
	}
}

func SetupBep20(invoker *kai.Invoker) {
	if config.DefConfig.OntBep20 != "" {
		bindTx, err := invoker.BindAssetHash(config.DefConfig.KaiLockProxy, config.DefConfig.Bep20,
			config.DefConfig.OntBep20, config.DefConfig.OntChainID, 0)
		if err != nil {
			panic(fmt.Errorf("SetupBep20ToONT, failed to BindAssetHash: %v", err))
		}
		invoker.Client().WaitTransactionConfirm(bindTx.Hash())
		hash := bindTx.Hash()
		fmt.Printf("binding bep20 of ontology on kai: ( txhash: %s )\n", hash.String())
	}

	bindTx, err := invoker.BindAssetHash(config.DefConfig.KaiLockProxy, config.DefConfig.Bep20,
		config.CM_BEP20, config.DefConfig.CMCrossChainId, 0)
	if err != nil {
		panic(fmt.Errorf("SetupBep20ToONT, failed to BindAssetHash: %v", err))
	}
	invoker.Client().WaitTransactionConfirm(bindTx.Hash())
	hash := bindTx.Hash()
	fmt.Printf("binding bep20 of cosmos on kai: ( txhash: %s )\n", hash.String())
}

func SetupWBTC(kaiInvoker *kai.Invoker) {
	bindTx, err := kaiInvoker.BindAssetHash(config.DefConfig.KaiLockProxy, config.DefConfig.KaiWBTC,
		config.DefConfig.OntWBTC, config.DefConfig.OntChainID, 0)
	if err != nil {
		panic(fmt.Errorf("SetupWBTC, failed to BindAssetHash: %v", err))
	}
	kaiInvoker.Client().WaitTransactionConfirm(bindTx.Hash())
	hash := bindTx.Hash()
	fmt.Printf("binding WBTC of ontology on kai: ( txhash: %s )\n", hash.String())
}

func SetupDAI(kaiInvoker *kai.Invoker) {
	bindTx, err := kaiInvoker.BindAssetHash(config.DefConfig.KaiLockProxy, config.DefConfig.KaiDai,
		config.DefConfig.OntDai, config.DefConfig.OntChainID, 0)
	if err != nil {
		panic(fmt.Errorf("SetupDAI, failed to BindAssetHash: %v", err))
	}
	kaiInvoker.Client().WaitTransactionConfirm(bindTx.Hash())
	hash := bindTx.Hash()
	fmt.Printf("binding DAI of ontology on kai: ( txhash: %s )\n", hash.String())
}

func SetupUSDT(kaiInvoker *kai.Invoker) {
	bindTx, err := kaiInvoker.BindAssetHash(config.DefConfig.KaiLockProxy, config.DefConfig.KaiUSDT,
		config.DefConfig.OntUSDT, config.DefConfig.OntChainID, 0)
	if err != nil {
		panic(fmt.Errorf("SetupUSDT, failed to BindAssetHash: %v", err))
	}
	kaiInvoker.Client().WaitTransactionConfirm(bindTx.Hash())
	hash := bindTx.Hash()
	fmt.Printf("binding USDT of ontology on kai: ( txhash: %s )\n", hash.String())
}

func SetupUSDC(kaiInvoker *kai.Invoker) {
	bindTx, err := kaiInvoker.BindAssetHash(config.DefConfig.KaiLockProxy, config.DefConfig.KaiUSDC,
		config.DefConfig.OntUSDC, config.DefConfig.OntChainID, 0)
	if err != nil {
		panic(fmt.Errorf("SetupUSDC, failed to BindAssetHash: %v", err))
	}
	kaiInvoker.Client().WaitTransactionConfirm(bindTx.Hash())
	hash := bindTx.Hash()
	fmt.Printf("binding USDC of ontology on kai: ( txhash: %s )\n", hash.String())
}

func SetupOntAsset(invoker *kai.Invoker) {
	if config.DefConfig.KaiLockProxy == "" {
		panic(fmt.Errorf("KaiLockProxy is blank"))
	}
	if config.DefConfig.KaiOntx == "" {
		panic(fmt.Errorf("KaicOntx is blank"))
	}
	if config.DefConfig.KaiOngx == "" {
		panic(fmt.Errorf("KaiOngx is blank"))
	}
	if config.DefConfig.KaiOep4 == "" {
		panic(fmt.Errorf("KaiOep4 is blank"))
	}
	if config.DefConfig.OntOep4 == "" {
		panic(fmt.Errorf("KaiOep4 is blank"))
	}

	txs, err := invoker.BindOntAsset(config.DefConfig.KaiLockProxy, config.DefConfig.KaiOntx, config.DefConfig.KaiOngx,
		config.DefConfig.KaiOep4, config.DefConfig.OntOep4)
	if err != nil {
		panic(err)
	}
	hash1, hash2, hash3 := txs[0].Hash(), txs[1].Hash(), txs[2].Hash()
	fmt.Printf("ont/ong/oep4 binding tx on ontology: %s/%s/%s\n", hash1.String(), hash2.String(), hash3.String())

	hash4, hash5, hash6 := txs[3].Hash(), txs[4].Hash(), txs[5].Hash()
	fmt.Printf("ont/ong/oep4 binding tx on cosmos: %s/%s/%s\n", hash4.String(), hash5.String(), hash6.String())
}

func SetupBnb(kaiInvoker *kai.Invoker) {
	kaiNativeAddr := "0x0000000000000000000000000000000000000000"
	if config.DefConfig.OntBnb != "" {
		tx, err := kaiInvoker.BindAssetHash(config.DefConfig.KaiLockProxy, kaiNativeAddr, config.DefConfig.OntBnb, config.DefConfig.OntChainID, 0)
		if err != nil {
			panic(fmt.Errorf("SetupBnb2ONT, failed to bind asset hash: %v", err))
		}
		hash := tx.Hash()
		fmt.Printf("binding bnbx of ontology on kai: ( txhash: %s )\n", hash.String())
	}

	if config.DefConfig.EthBnb != "" {
		tx, err := kaiInvoker.BindAssetHash(config.DefConfig.KaiLockProxy, kaiNativeAddr, config.DefConfig.EthBnb, config.DefConfig.EthChainID, 0)
		if err != nil {
			panic(fmt.Errorf("SetupBnb2ONT, failed to bind asset hash: %v", err))
		}
		hash := tx.Hash()
		fmt.Printf("binding bnb of kai on ethereum: ( txhash: %s )\n", hash.String())
	}
	if config.DefConfig.NeoBnb != "" {
		tx, err := kaiInvoker.BindAssetHash(config.DefConfig.KaiLockProxy, kaiNativeAddr, config.DefConfig.NeoBnb, config.DefConfig.NeoChainID, 0)
		if err != nil {
			panic(fmt.Errorf("SetupBnb2Neo, failed to bind asset hash: %v", err))
		}
		hash := tx.Hash()
		fmt.Printf("binding bnb of kai on neo: ( txhash: %s )\n", hash.String())
	}

	tx, err := kaiInvoker.BindAssetHash(config.DefConfig.KaiLockProxy, kaiNativeAddr, config.CM_BNBX, config.DefConfig.CMCrossChainId, 0)
	if err != nil {
		panic(fmt.Errorf("SetupBnb2COSMOS, failed to bind asset hash: %v", err))
	}
	hash := tx.Hash()
	fmt.Printf("binding bnbx of cosmos on kai: ( txhash: %s )\n", hash.String())

	tx, err = kaiInvoker.BindAssetHash(config.DefConfig.KaiLockProxy, kaiNativeAddr, kaiNativeAddr, config.DefConfig.KaiChainID, 0)
	if err != nil {
		panic(fmt.Errorf("BindAssetHash, failed to bind asset hash: %v", err))
	}
	hash = tx.Hash()
	fmt.Printf("binding bnb of kai on kai: ( txhash: %s )\n", hash.String())
}

func SetOtherLockProxy(invoker *kai.Invoker) {
	_, contract, err := invoker.MakeLockProxy(config.DefConfig.KaiLockProxy)
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
			panic(fmt.Errorf("failed to bind proxy: %s", err))
		}
		hash := tx.Hash()
		invoker.Client().WaitTransactionConfirm(hash)
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
		invoker.Client().WaitTransactionConfirm(hash)
		fmt.Printf("binding cosmos proxy: ( txhash: %s )\n", hash.String())
	}

	if config.DefConfig.KaiLockProxy != "" {
		auth, err := invoker.MakeSmartContractAuth()
		if err != nil {
			panic(fmt.Errorf("failed to get auth: %v", err))
		}
		other := common2.HexToAddress(config.DefConfig.KaiLockProxy)
		tx, err := contract.BindProxyHash(auth, config.DefConfig.KaiChainID, other[:])
		if err != nil {
			panic(fmt.Errorf("failed to bind proxy: %v", err))
		}
		hash := tx.Hash()
		invoker.Client().WaitTransactionConfirm(hash)
		fmt.Printf("binding kai proxy: ( txhash: %s )\n", hash.String())
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
		invoker.Client().WaitTransactionConfirm(hash)
		fmt.Printf("binding kai proxy: ( txhash: %s )\n", hash.String())
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
		invoker.Client().WaitTransactionConfirm(hash)
		fmt.Printf("binding neo proxy: ( txhash: %s )\n", hash.String())
	}
}

func SetUpContracts() {
	invoker := kai.NewInvoker(config.DefConfig.KaiChainID)
	SetupBnb(invoker)
	if config.DefConfig.Bep20 != "" {
		SetupBep20(invoker)
	}
	if config.DefConfig.OntLockProxy != "" {
		SetupOntAsset(invoker)
	}
	if config.DefConfig.KaiWBTC != "" {
		SetupWBTC(invoker)
	}
	if config.DefConfig.KaiDai != "" {
		SetupDAI(invoker)
	}
	if config.DefConfig.KaiUSDT != "" {
		SetupUSDT(invoker)
	}

	//SetupUSDC(invoker)
	SetOtherLockProxy(invoker)
}
