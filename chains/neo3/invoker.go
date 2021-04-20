package neo3

import (
	"fmt"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/rpc"
	"github.com/joeqian10/neo3-gogogo/wallet"
	"github.com/polynetwork/poly-io-test/config"
)

type Neo3Invoker struct {
	*wallet.WalletHelper
}

func NewNeo3Invoker() (*Neo3Invoker, error) {
	invoker := &Neo3Invoker{}
	if config.DefConfig.Neo3Url == "" {
		return nil, fmt.Errorf("Neo3Url is empty")
	}
	cli := rpc.NewClient(config.DefConfig.NeoUrl)
	ps := helper.ProtocolSettings{
		Magic:          config.DefConfig.Neo3Magic,
		AddressVersion: config.DefConfig.Neo3AddressVersion,
	}
	name := "poly_io_test"
	wlt, err := wallet.NewNEP6Wallet(config.DefConfig.Neo3Wallet, &ps, &name, wallet.DefaultScryptParameters)
	if err != nil {
		return nil, err
	}
	err = wlt.Unlock(config.DefConfig.Neo3Pwd)
	if err != nil {
		return nil, err
	}

	invoker.WalletHelper = wallet.NewWalletHelperFromWallet(cli, wlt)

	return invoker, nil
}
