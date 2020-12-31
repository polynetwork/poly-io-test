package main

import (
	"flag"
	"fmt"
	ethComm "github.com/ethereum/go-ethereum/common"
	"github.com/ontio/ontology/common"
	"github.com/polynetwork/poly-io-test/chains/neo"
	"github.com/polynetwork/poly-io-test/chains/ont"
	"github.com/polynetwork/poly-io-test/config"
	"github.com/polynetwork/poly-io-test/log"
)

var (
	fnNeo       string
	neoConfFile string
)

func init() {
	flag.StringVar(&fnNeo, "func", "setup", "choose function to run: deploy or setup")
	flag.StringVar(&neoConfFile, "conf", "./config.json", "config path")
	flag.Parse()
}

func main() {
	err := config.DefConfig.Init(neoConfFile)
	if err != nil {
		panic(err)
	}
	invoker, err := neo.NewNeoInvoker()
	if err != nil {
		panic(err)
	}
	switch fnNeo {
	case "deploy":
		log.Errorf("Neo deploy contract not supported, pls deploy through gui")
		return
	case "setup":
		SetUpNeoContracts(invoker)

	}
}

func SetUpNeoContracts(invoker *neo.NeoInvoker) {
	neoProxyHash, err := neo.ParseNeoAddr(config.DefConfig.NeoLockProxy)
	if err != nil {
		log.Errorf("parse neo lockproxy:%s, err: %v", config.DefConfig.NeoLockProxy, err)
		return
	}
	if config.DefConfig.HecoChainID != 0 {
		SetupNeo2Heco(invoker, neoProxyHash)
	}

}

func SetupNeo2Heco(invoker *neo.NeoInvoker, neoProxyHash []byte) {
	// set up proxy contracts
	if config.DefConfig.HecoLockProxy != "" {
		err := invoker.BindProxyHash(neoProxyHash, config.DefConfig.HecoChainID, ethComm.HexToAddress(config.DefConfig.HecoLockProxy).Bytes())
		if err != nil {
			log.Errorf("Bind proxy from neo to heco  err: %v", err)
			return
		}
		log.Infof("BindProxy, from neo to heco successful")
		if config.DefConfig.HecoHt != "" && config.DefConfig.NeoHt != "" {
			fromAssetNeoHt, _ := neo.ParseNeoAddr(config.DefConfig.NeoHt)

			if _, err := invoker.BindAssetHash(neoProxyHash, fromAssetNeoHt, config.DefConfig.HecoChainID, ethComm.HexToAddress(config.DefConfig.HecoHt).Bytes()); err != nil {
				log.Errorf("Bind Ht from neo to heco  err: %v", err)
			} else {
				log.Infof("BindAsset ht, from neo to heco successful")
			}
		}
		if config.DefConfig.HecoHrc20 != "" && config.DefConfig.NeoHrc20 != "" {
			fromAssetNeoHrc20, _ := neo.ParseNeoAddr(config.DefConfig.NeoHrc20)

			if _, err := invoker.BindAssetHash(neoProxyHash, fromAssetNeoHrc20, config.DefConfig.HecoChainID, ethComm.HexToAddress(config.DefConfig.HecoHrc20).Bytes()); err != nil {
				log.Errorf("Bind Hrc20 from neo to heco  err: %v", err)
			}
			log.Infof("BindAsset hrc20, from neo to heco successful")
		}
	}
}

func GetInfo(addrs []common.Address) string {
	str := "=============================ONT info=============================\n"
	for i, name := range ont.ContractNames {
		str += fmt.Sprintf("{contract_name: %s, address: %s}\n", name, addrs[i].ToHexString())
	}
	str += "=================================================================="
	return str
}
