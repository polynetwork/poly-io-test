package main

import (
	"flag"
	"github.com/ChainSQL/go-chainsql-api/data"
	"github.com/polynetwork/poly-io-test/chains/chainsql"
	"github.com/polynetwork/poly-io-test/config"
	"github.com/polynetwork/poly-io-test/log"
)

var (
	fnEth        string
	configPath   string
)

func init() {
	flag.StringVar(&configPath, "conf", "./config.json", "Config of poly-io-test")
	flag.StringVar(&fnEth, "func", "deploy", "choose function to run: deploy or setup")
	flag.Parse()
}

func main() {
	err := config.DefConfig.Init(configPath)
	if err != nil {
		panic(err)
	}

	switch fnEth {
	case "deploy":
		DeployChainsqlSmartContract()
	case "setup":
		break
	}
}

func DeployChainsqlSmartContract() {

	var (
		eccdAddr  string
		eccmAddr  string
		err       error
	)

	invoker, err := chainsql.NewChainsqlInvoker()
	if err != nil {
		panic(err)
	}

	eccdAddr,err = invoker.DeployCrossChainDataContract()
	if err != nil{
		panic(err)
	}

	account, err := data.NewAccountFromAddress(eccdAddr)
	if err != nil {
		panic(err)
	}

	log.Infof("eccd_address:%s, hex:0x%x",eccdAddr,account.Bytes())

	eccmAddr,err = invoker.DeployCrossChainManagerContract(eccdAddr,config.DefConfig.ChainsqlChainID)

	result,err := invoker.TransaferOwnershipForECCD(eccdAddr,eccmAddr)

	if err != nil{
		panic(err)
	}
	if result.Status != "validate_success"{
		panic(result.ErrorMessage)
	}
	account, err = data.NewAccountFromAddress(eccmAddr)
	if err != nil {
		panic(err)
	}

	log.Infof("eccm_address:%s, hex:0x%x",eccmAddr,account.Bytes())
	eccmpAddr,err := invoker.DeployCrossChainManagerProxyContract(eccmAddr)
	if err != nil{
		panic(err)
	}
	result,err = invoker.TransferOwnershipForECCM(eccmAddr,eccmpAddr)
	if result.Status != "validate_success"{
		panic(result.ErrorMessage)
	}

	account, err = data.NewAccountFromAddress(eccmpAddr)
	if err != nil {
		panic(err)
	}
	log.Infof("eccmp_address:%s, hex:0x%x",eccmpAddr,account.Bytes())
}
