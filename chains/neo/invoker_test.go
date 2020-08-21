package neo

import (
	"fmt"
	"github.com/polynetwork/poly-io-test/config"
	"testing"
)

func TestNewNeoInvoker(t *testing.T) {
	config.DefConfig.Init("./config-test-cp.json")
	invoker, err := NewNeoInvoker()
	if err != nil {
		t.Fatal(err)
	}
	res := invoker.Cli.GetBlockCount()
	if res.HasError() {
		t.Fatal(res.Error.Message)
	}
	fmt.Println(res.Result)

	res1 := invoker.Cli.GetContractState("0xa837ba329255884b40581ba8a3d29820acf44316")
	fmt.Println(res1)
}
