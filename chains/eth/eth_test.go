package eth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/polynetwork/poly-io-test/chains/eth/abi/ongx"
	"github.com/polynetwork/poly-io-test/config"
	"testing"
)

func TestNewNonceManager(t *testing.T) {
	err := config.DefConfig.Init("config.json")
	if err != nil {
		t.Fatal(err)
	}

	invoker := NewEInvoker()

	contract, err := ongx.NewONGX(common.HexToAddress("0xFb37c160CFBd8BD4Ba6df6f70e2449b6EB83fc26"), invoker.ETHUtil.GetEthClient())
	if err != nil {
		t.Fatal(err)
	}

	amt, err := contract.Allowance(nil, common.HexToAddress("0x344cFc3B8635f72F14200aAf2168d9f75df86FD3"),
		common.HexToAddress("0x388Ed8B73bd707A78034E1d157fA08Da24095c18"))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(amt.Uint64(), 3%1)

}
