package main

import (
	"flag"
	"fmt"
	"github.com/Zilliqa/gozilliqa-sdk/account"
	"github.com/Zilliqa/gozilliqa-sdk/crosschain/polynetwork"
	"github.com/Zilliqa/gozilliqa-sdk/provider"
	zilliqa "github.com/polynetwork/poly-io-test/chains/zil"
	"github.com/polynetwork/poly-io-test/config"
	"log"
	"strconv"
)

var (
	fnZil       string
	zilConfFile string
)

func init() {
	flag.StringVar(&fnZil, "func", "deploy", "choose function to run: deploy or setup")
	flag.StringVar(&zilConfFile, "conf", "./config.json", "config file path")
	flag.Parse()
}

func main() {
	err := config.DefConfig.Init(zilConfFile)
	if err != nil {
		panic(err)
	}

	switch fnZil {
	case "deploy":
		DeployZILSmartContract()
	case "setup":
		SetupZILSmartContract()
	}
}

func DeployZILSmartContract() {
	deployer := zilliqa.NewDeployer(config.DefConfig.ZilPrivateKey, config.DefConfig.ZilURL)
	wallet := account.NewWallet()
	wallet.AddByPrivateKey(deployer.PrivateKey)
	client := provider.NewProvider(deployer.Host)
	zilliqaChainId, err := client.GetNetworkId()
	if err != nil {
		panic(err)
	}
	chainId, err := strconv.ParseInt(zilliqaChainId, 10, 64)
	if err != nil {
		panic(err)
	}

	zilSideChainId := strconv.FormatUint(config.DefConfig.ZilChainID, 10)

	proxy, impl, lockProxy, err := deployer.Deploy(wallet, client, zilSideChainId, int(chainId))

	log.Printf("corss chain manager proxy address: %s\n", proxy)
	log.Printf("cross chain manager address: %s\n", impl)
	log.Printf("lock proxy address: %s\n", lockProxy)
	if err != nil {
		log.Fatalln(err.Error())
	}

	p := &polynetwork.Proxy{
		ProxyAddr:  proxy,
		ImplAddr:   impl,
		Wallet:     wallet,
		Client:     client,
		ChainId:    int(chainId),
		MsgVersion: 1,
	}

	fmt.Println("4. upgrade cross chain manager proxy contract")
	_, err1 := p.UpgradeTo()
	if err1 != nil {
		log.Fatalln(err1.Error())
	}

	fmt.Println("5. unpause cross chain manager contract")
	_, err2 := p.Unpause()
	if err2 != nil {
		log.Fatalln(err2.Error())
	}
}

func SetupZILSmartContract() {
	wallet := account.NewWallet()
	wallet.AddByPrivateKey(config.DefConfig.ZilPrivateKey)
	client := provider.NewProvider(config.DefConfig.ZilURL)
	zilliqaChainId, err := client.GetNetworkId()
	if err != nil {
		panic(err)
	}
	chainId, err := strconv.ParseInt(zilliqaChainId, 10, 64)
	if err != nil {
		panic(err)
	}

	l := &polynetwork.LockProxy{
		Addr:       config.DefConfig.ZilLockProxy,
		Wallet:     wallet,
		Client:     client,
		ChainId:    int(chainId),
		MsgVersion: 1,
	}

	sideChainId := strconv.FormatUint(config.DefConfig.ZilChainID, 10)

	_, err = l.BindProxyHash(sideChainId, config.DefConfig.EthLockProxy)
	if err != nil {
		panic(err)
	}

	// _,err = l.BindAssetHash()
}
