package main

import (
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/polynetwork/poly-io-test/chains/fisco"
	"github.com/polynetwork/poly-io-test/config"
)

var (
	fnEth        string
	configPath   string
	eccmRedeploy int
)

func init() {
	flag.StringVar(&configPath, "conf", "./config.json", "Config of poly-io-test")
	flag.StringVar(&fnEth, "func", "deploy", "choose function to run: deploy or setup")
	flag.IntVar(&eccmRedeploy, "redeploy_eccm", 1, "redeploy eccd, eccm and eccmp or not")
	flag.Parse()
}

func main() {
	err := config.DefConfig.Init(configPath)
	if err != nil {
		panic(err)
	}

	switch fnEth {
	case "deploy":
		DeployFiscoSmartContract()
	case "setup":
		SetUpFiscoContracts()
	}
}

func DeployFiscoSmartContract() {
	var (
		eccdAddr  common.Address
		eccmAddr  common.Address
		eccmpAddr common.Address
		err       error
	)
	invoker, err := fisco.NewFiscoInvoker()
	if err != nil {
		panic(err)
	}

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
		eccdAddr = common.HexToAddress(config.DefConfig.FiscoCCDC)
		eccmAddr = common.HexToAddress(config.DefConfig.FiscoCCMC)
		eccmpAddr = common.HexToAddress(config.DefConfig.FiscoCCMP)
	}

	lockProxyAddr, _, err := invoker.DeployLockProxyContract(eccmpAddr)
	if err != nil {
		panic(err)
	}

	lockproxyAddrHex := lockProxyAddr.Hex()
	//erc20Addr, erc20, err := invoker.DeployERC20()
	//if err != nil {
	//	panic(err)
	//}
	pethAddr, peth, err := invoker.DeployFISCOXContract(lockproxyAddrHex)
	if err != nil {
		panic(err)
	}

	total, err := peth.TotalSupply(invoker.MakeSmartClient().GetCallOpts())
	if err != nil {
		panic(fmt.Errorf("failed to get total supply for erc20: %v", err))
	}
	auth := invoker.MakeSmartContractAuth()
	_, err = peth.Approve(auth, lockProxyAddr, total)
	if err != nil {
		panic(fmt.Errorf("failed to approve erc20 to lockproxy: %v", err))
	}

	//fmt.Println("Approve:", tx.Hash().Hex())

	//oep4Addr, _, err := invoker.DeployOEP4(lockproxyAddrHex)
	//if err != nil {
	//	panic(err)
	//}
	//ongxAddr, _, err := invoker.DeployONGXContract(lockproxyAddrHex)
	//if err != nil {
	//	panic(err)
	//}
	//ontxAddr, _, err := invoker.DeployONTXContract(lockproxyAddrHex)
	//if err != nil {
	//	panic(err)
	//}
	//设置BindProxyHash和BindAssetHash
	//invoker.BindLockProxy(lockproxyAddrHex, fiscoxAddr.Hex())

	fmt.Println("=============================FISCO info=============================")
	//fmt.Println("erc20:", erc20Addr.Hex())
	//fmt.Println("ope4:", oep4Addr.Hex())
	fmt.Println("eccd address:", eccdAddr.Hex())
	fmt.Println("eccm address:", eccmAddr.Hex())
	fmt.Println("eccmp address:", eccmpAddr.Hex())
	fmt.Println("lock proxy address: ", lockproxyAddrHex)
	fmt.Println("pEth address: ", pethAddr.Hex())
	//fmt.Println("ongx address: ", ongxAddr.Hex())
	//fmt.Println("ontx proxy address: ", ontxAddr.Hex())
	fmt.Println("==================================================================")

	config.DefConfig.FiscoCCDC = eccdAddr.Hex()
	config.DefConfig.FiscoCCMC = eccmAddr.Hex()
	config.DefConfig.FiscoCCMP = eccmpAddr.Hex()
	config.DefConfig.FiscoLockProxy = lockproxyAddrHex
	config.DefConfig.FiscoEth = pethAddr.Hex()

	err = config.DefConfig.Save(configPath)
	if err != nil {
		panic(err)
	}
}

func SetUpFiscoContracts() {
	invoker, err := fisco.NewFiscoInvoker()
	if err != nil {
		panic(err)
	}

	err = invoker.BindLockProxy(config.DefConfig.FiscoLockProxy, config.DefConfig.EthChainID, config.DefConfig.EthLockProxy)
	if err != nil {
		panic(err)
	}

	_, err = invoker.BindAssetHash(config.DefConfig.FiscoLockProxy, config.DefConfig.FiscoEth, "0x0000000000000000000000000000000000000000", config.DefConfig.EthChainID)
	if err != nil {
		panic(err)
	}

	err = invoker.BindLockProxy(config.DefConfig.FiscoLockProxy, config.DefConfig.FabricChainID, config.DefConfig.FabricLockProxy)
	if err != nil {
		panic(err)
	}

	_, err = invoker.BindAssetHash(config.DefConfig.FiscoLockProxy, config.DefConfig.FiscoEth, config.DefConfig.FabricPEth, config.DefConfig.FabricChainID)
	if err != nil {
		panic(err)
	}
}
