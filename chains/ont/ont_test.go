package ont

import (
	"fmt"
	"github.com/ontio/ontology/common"
	"github.com/polynetwork/poly-io-test/config"
	"math/big"
	"testing"
)

func TestGetAccountByPassword(t *testing.T) {
	config.DefConfig.Init("./config-test-cp.json")
	inv, err := NewOntInvoker(config.DefConfig.OntJsonRpcAddress, config.DefConfig.OntContractsAvmPath, config.DefConfig.OntWallet, config.DefConfig.OntWalletPassword)
	if err != nil {
		t.Fatal(err)
	}

	c, _ := common.AddressFromHexString(config.DefConfig.OntOep4)
	res, err := inv.OntSdk.NeoVM.PreExecInvokeNeoVMContract(c, []interface{}{"balanceOf", []interface{}{inv.OntAcc.Address[:]}})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res.Result.ToInteger())

	amt, err := res.Result.ToInteger()
	if err != nil {
		t.Fatal(err)
	}

	to, err := common.AddressFromBase58("AWuqAh23z13874ovnPW2BiHt9kMAEqs4ag")
	if err != nil {
		t.Fatal(err)
	}
	res, err = inv.OntSdk.NeoVM.PreExecInvokeNeoVMContract(c, []interface{}{"balanceOf", []interface{}{to[:]}})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res.Result.ToInteger())

	tx, err := inv.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		inv.OntAcc, inv.OntAcc,
		c,
		[]interface{}{inv.OntAcc.Address, to, amt.Div(amt, big.NewInt(2))})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tx.ToHexString())
}
