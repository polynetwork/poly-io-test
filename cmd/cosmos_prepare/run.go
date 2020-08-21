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
	"github.com/polynetwork/poly-io-test/chains/cosmos"
	"github.com/polynetwork/poly-io-test/config"
	"github.com/polynetwork/poly-io-test/log"
	"strings"
)

var (
	cmConfFile string
)

func init() {
	flag.StringVar(&cmConfFile, "conf", "./config.json", "config file path")
	flag.Parse()
}

func main() {
	err := config.DefConfig.Init(cmConfFile)
	if err != nil {
		panic(err)
	}
	invoker, err := cosmos.NewCosmosInvoker()
	if err != nil {
		panic(err)
	}
	res, err := invoker.CreateLockProxy()
	if err != nil {
		if !strings.Contains(err.Error(), "already") {
			panic(err)
		}
		log.Infof("already created lockproxy %s", hex.EncodeToString(invoker.Acc.Acc.Bytes()))
	} else {
		invoker.WaitTx(res.Hash)
	}
	err = invoker.SetupAllAssets(invoker.Acc.Acc.Bytes())
	if err != nil {
		panic(err)
	}

	err = invoker.SetupBtcx(config.CM_BTCX, config.DefConfig.BtcRedeem)
	if err != nil {
		panic(err)
	}

	err = config.DefConfig.Save(cmConfFile)
	if err != nil {
		panic(err)
	}

	log.Info("successful to set cosmos up")
}
