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
package neo

import (
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/rpc"
	"github.com/joeqian10/neo-gogogo/wallet"
	"github.com/polynetwork/poly-io-test/config"
	"github.com/polynetwork/poly-io-test/log"
	"strings"
	"time"
)

type NeoInvoker struct {
	Cli *rpc.RpcClient
	Acc *wallet.Account
}

func NewNeoInvoker() (invoker *NeoInvoker, err error) {
	invoker = &NeoInvoker{}
	invoker.Cli = rpc.NewClient(config.DefConfig.NeoUrl)
	invoker.Acc, err = wallet.NewAccountFromWIF(config.DefConfig.NeoWif)
	if err != nil {
		return nil, err
	}
	return
}

func WaitNeoTx(cli *rpc.RpcClient, hash helper.UInt256) {
	tick := time.NewTicker(100 * time.Millisecond)
	for range tick.C {
		res := cli.GetTransactionHeight(hash.String())
		if res.HasError() {
			if strings.Contains(res.Error.Message, "Unknown") {
				continue
			}
			log.Errorf("failed to get neo tx: %s", res.Error.Message)
			continue
		}
		if res.Result <= 0 {
			continue
		}
		log.Infof("capture neo tx %s", hash.String())
		break
	}
}
