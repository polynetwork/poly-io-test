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
