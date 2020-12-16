package fabric

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	config2 "github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/polynetwork/poly-io-test/config"
	"os"
)

type FabricInvoker struct {
	Sdk       *fabsdk.FabricSDK
	ChanCli   *channel.Client
	EventCli  *event.Client
	LedgerCLi *ledger.Client
}

func NewFabricInvoker() (*FabricInvoker, error) {
	_ = os.Setenv("FABRIC_RELAYER_PATH", config.DefConfig.FabricMSPPath)

	sdk, err := fabsdk.New(config2.FromFile(config.DefConfig.FabricSdkConfFile))
	if err != nil {
		return nil, err
	}
	ccp := sdk.ChannelContext(config.DefConfig.FabricChannel, fabsdk.WithUser(config.DefConfig.FabricUser), fabsdk.WithOrg(config.DefConfig.FabricOrg))
	cc, err := channel.New(ccp)
	if err != nil {
		return nil, err
	}
	eventClient, err := event.New(ccp, event.WithBlockEvents())
	if err != nil {
		return nil, err
	}
	ledgerClient, err := ledger.New(ccp)
	if err != nil {
		return nil, err
	}

	return &FabricInvoker{
		ChanCli:   cc,
		EventCli:  eventClient,
		LedgerCLi: ledgerClient,
		Sdk:       sdk,
	}, nil
}

func (inv *FabricInvoker) CallChainCode(chaincode, fnc string, args []string) (channel.Response, error) {
	req := channel.Request{
		ChaincodeID: chaincode,
		Fcn:         fnc,
		Args:        packArgs(args),
	}
	response, err := inv.ChanCli.Execute(req, channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		return channel.Response{}, err
	}
	return response, nil
}

func packArgs(args []string) [][]byte {
	ret := make([][]byte, 0)
	for _, arg := range args {
		ret = append(ret, []byte(arg))
	}
	return ret
}
