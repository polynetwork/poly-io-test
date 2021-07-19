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
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	poly_go_sdk "github.com/polynetwork/poly-go-sdk"
	"github.com/polynetwork/poly-io-test/chains/btc"
	"github.com/polynetwork/poly-io-test/chains/cosmos"
	"github.com/polynetwork/poly-io-test/chains/eth"
	"github.com/polynetwork/poly-io-test/chains/neo"
	"github.com/polynetwork/poly-io-test/chains/ont"
	"github.com/polynetwork/poly-io-test/config"
	"github.com/polynetwork/poly-io-test/log"
	_ "github.com/polynetwork/poly-io-test/testcase"
	"github.com/polynetwork/poly-io-test/testframework"
)

var (
	TestConfig string //Test config file
	TestCases  string //TestCase list in cmdline
	LoopNumber int
)

func init() {
	flag.StringVar(&TestConfig, "cfg", "./config.json", "Config of poly-io-test")
	flag.StringVar(&TestCases, "t", "", "Test case to run. use ',' to split test case")
	flag.IntVar(&LoopNumber, "loop", 1, " the number the whole test cases run")
	flag.Parse()
}

func main() {
	err := config.DefConfig.Init(TestConfig)
	if err != nil {
		log.Errorf("DefConfig.Init error:%s", err)
		return
	}

	rcSdk := poly_go_sdk.NewPolySdk()
	if err = btc.SetUpPoly(rcSdk, config.DefConfig.RchainJsonRpcAddress); err != nil {
		panic(err)
	}

	var (
		ethInvoker             *eth.EInvoker
		bscInvoker             *eth.EInvoker
		mscInvoker             *eth.EInvoker
		o3Invoker              *eth.EInvoker
		cmInvoker              *cosmos.CosmosInvoker
		kaiInvoker, borInvoker *eth.EInvoker
	)
	if config.DefConfig.EthChainID > 0 {
		ethInvoker = eth.NewEInvoker(config.DefConfig.EthChainID)
	}

	if config.DefConfig.BscChainID > 0 {
		bscInvoker = eth.NewEInvoker(config.DefConfig.BscChainID)
	}

	if config.DefConfig.MscChainID > 0 {
		mscInvoker = eth.NewEInvoker(config.DefConfig.MscChainID)
	}

	if config.DefConfig.O3ChainID > 0 {
		o3Invoker = eth.NewEInvoker(config.DefConfig.O3ChainID)
	}

	if config.DefConfig.KaiChainID > 0 {
		kaiInvoker = eth.NewEInvoker(config.DefConfig.KaiChainID)
	}

	if config.DefConfig.PolygonBorChainID > 0 {
		borInvoker = eth.NewEInvoker(config.DefConfig.PolygonBorChainID)
	}

	//btcInvoker, err := btc.NewBtcInvoker(config.DefConfig.RchainJsonRpcAddress, config.DefConfig.RCWallet,
	//	config.DefConfig.RCWalletPwd, config.DefConfig.BtcRestAddr, config.DefConfig.BtcRestUser,
	//	config.DefConfig.BtcRestPwd, config.DefConfig.BtcSignerPrivateKey)
	//if err != nil {
	//	log.Errorf("failed to new a btc invoker, do not test cases about BTC: %v", err)
	//}
	ontInvoker, err := ont.NewOntInvoker(config.DefConfig.OntJsonRpcAddress, config.DefConfig.OntContractsAvmPath,
		config.DefConfig.OntWallet, config.DefConfig.OntWalletPassword)
	if err != nil {
		log.Warnf("failed to new a ont invoker, do not test cases about ONT: %v", err)
	}
	if config.DefConfig.CMCrossChainId > 0 {
		cmInvoker, err = cosmos.NewCosmosInvoker()
		if err != nil {
			log.Warnf("failed to new a cosmos invoker, do not test cases about COSMOS: %v", err)
		}
	}

	neoInvoker, err := neo.NewNeoInvoker()
	if err != nil {
		log.Warnf("failed to new a neo invoker, do not test cases about NEO: %v", err)
	}

	testCases := make([]string, 0)
	if TestCases != "" {
		testCases = strings.Split(TestCases, ",")
	}
	testframework.TFramework.SetRcSdk(rcSdk)
	if ethInvoker != nil {
		testframework.TFramework.SetEthInvoker(ethInvoker)
	}
	if bscInvoker != nil {
		testframework.TFramework.SetBSCInvoker(bscInvoker)
	}
	if mscInvoker != nil {
		testframework.TFramework.SetMSCInvoker(mscInvoker)
	}
	if o3Invoker != nil {
		testframework.TFramework.SetO3Invoker(o3Invoker)
	}

	if kaiInvoker != nil {
		testframework.TFramework.SeKaiInvoker(kaiInvoker)
	}
	if borInvoker != nil {
		testframework.TFramework.SeBorInvoker(borInvoker)
	}

	//testframework.TFramework.SetBtcInvoker(btcInvoker)
	testframework.TFramework.SetOntInvoker(ontInvoker)
	testframework.TFramework.SetCosmosInvoker(cmInvoker)
	testframework.TFramework.SetNeoInvoker(neoInvoker)

	//Start run test case
	testframework.TFramework.Run(testCases, LoopNumber)
	waitToExit()
}

func waitToExit() {
	exit := make(chan bool, 0)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	go func() {
		for sig := range sc {
			fmt.Println("cross chain test received exit signal: ", sig.String())
			close(exit)
			break
		}
	}()
	<-exit
}
