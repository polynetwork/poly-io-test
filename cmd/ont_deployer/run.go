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
	"github.com/polynetwork/poly-io-test/chains/ont"
	"github.com/polynetwork/poly-io-test/config"
)

var (
	fnOnt       string
	ontConfFile string
)

func init() {
	flag.StringVar(&fnOnt, "func", "deploy", "choose function to run: deploy or setup")
	flag.StringVar(&ontConfFile, "conf", "./config.json", "config path")
	flag.Parse()
}

func main() {
	err := config.DefConfig.Init(ontConfFile)
	if err != nil {
		panic(err)
	}

	invoker, err := ont.NewOntInvoker(config.DefConfig.OntJsonRpcAddress, config.DefConfig.OntContractsAvmPath,
		config.DefConfig.OntWallet, config.DefConfig.OntWalletPassword)
	if err != nil {
		panic(err)
	}
	switch fnOnt {
	case "deploy":
		addrs, err := invoker.DeployContracts()
		if err != nil {
			panic(err)
		}

		config.DefConfig.OntLockProxy = addrs[2].ToHexString()
		config.DefConfig.OntErc20 = addrs[3].ToHexString()
		config.DefConfig.OntOep4 = addrs[4].ToHexString()
		config.DefConfig.OntEth = addrs[5].ToHexString()
		config.DefConfig.OntWBTC = addrs[6].ToHexString()
		config.DefConfig.OntDai = addrs[7].ToHexString()
		config.DefConfig.OntUSDC = addrs[8].ToHexString()
		config.DefConfig.OntUSDT = addrs[9].ToHexString()

		err = config.DefConfig.Save(ontConfFile)
		if err != nil {
			panic(fmt.Errorf("failed to save config: %v", err))
		}

		fmt.Println(GetInfo(addrs))
	case "setup":
		txs, err := invoker.SetupOntAsset(config.DefConfig.OntLockProxy, config.DefConfig.EthOntx,
			config.DefConfig.EthOngx, config.DefConfig.OntOep4, config.DefConfig.EthOep4, config.DefConfig.GasPrice,
			config.DefConfig.GasLimit)
		if err != nil {
			panic(fmt.Errorf("failed to setup ont asset: %v", err))
		}
		fmt.Printf("set up ont asset:\n{ont of ethereum binding txhash: %s}\n{ont of cosmos binding txhash: %s}\n"+
			"{ong of ethereum bingding txhash: %s}\n{ong of cosmos bingding txhash: %s}\n"+
			"{oep4 of ethereum binding txhash: %s}\n{oep4 of cosmos binding txhash: %s}\n",
			txs[0].ToHexString(), txs[1].ToHexString(), txs[2].ToHexString(), txs[3].ToHexString(), txs[4].ToHexString(),
			txs[5].ToHexString())

		txs, err = invoker.SetupEthAsset(config.DefConfig.OntLockProxy, config.DefConfig.OntEth,
			config.DefConfig.EthErc20, config.DefConfig.OntErc20, config.DefConfig.GasPrice, config.DefConfig.GasLimit)
		if err != nil {
			panic(fmt.Errorf("failed to setup eth asset: %v", err))
		}
		fmt.Printf("set up eth asset:\n"+
			"{eth of ethereum binding txhash: %s}\n"+
			"{eth of cosmos binding txhash: %s}\n"+
			"{erc20 of ethereum bingding txhash: %s}\n"+
			"{erc20 of cosmos bingding txhash: %s}\n",
			txs[0].ToHexString(), txs[2].ToHexString(), txs[4].ToHexString(), txs[5].ToHexString())

		if config.DefConfig.EthLockProxy == "" {
			panic(fmt.Errorf("EthLockProxy is blank"))
		}

		otherAddr := common2.HexToAddress(config.DefConfig.EthLockProxy)
		txhash, err := invoker.SetOtherLockProxy(otherAddr.Bytes(), config.DefConfig.EthChainID)
		if err != nil {
			panic(fmt.Errorf("failed to bind eth lock proxy on ont_proxy: %v", err))
		}
		fmt.Printf("bind eth proxy on ont_proxy: {txhash: %s}\n", txhash.ToHexString())

		cmProxy, err := hex.DecodeString(config.DefConfig.CMLockProxy)
		if err != nil {
			panic(fmt.Errorf("failed to decode proxy: %v", err))
		}
		txhash, err = invoker.SetOtherLockProxy(cmProxy, config.DefConfig.CMCrossChainId)
		if err != nil {
			panic(fmt.Errorf("failed to bind cosmos lock proxy: %v", err))
		}

		fmt.Printf("bind cosmos proxy on ont_proxy: {txhash: %s}\n", txhash.ToHexString())
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
