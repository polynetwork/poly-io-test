package neo

import (
	"encoding/hex"
	"fmt"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/sc"
	"github.com/joeqian10/neo-gogogo/tx"
	"github.com/polynetwork/poly-io-test/config"
	"github.com/polynetwork/poly-io-test/log"
	"github.com/polynetwork/poly/common"
	"math/big"
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

func Test_GetOperator(t *testing.T) {
	config.DefConfig.Init("./config_main.json")
	invoker, err := NewNeoInvoker()
	if err != nil {
		t.Fatal(err)
	}
	ccmc, _ := helper.UInt160FromString(config.DefConfig.NeoCCMC)
	neoLock, _ := helper.UInt160FromString(config.DefConfig.NeoLockProxy)

	// check if initialized
	{
		res, err := invoker.GetStorage(ccmc.Bytes(), hex.EncodeToString([]byte("IsInitGenesisBlock")))
		if err != nil {
			log.Errorf("Initialized height err: %v", err)
			return
		}
		log.Infof("ccmc initialized poly height on neo ccmc is: %s", res)
	}

	{
		res, err := invoker.GetProxyOperator(neoLock.Bytes())
		if err != nil {
			log.Errorf("GetOperator err: %v", err)
			return
		}
		log.Infof("neoLockProxy operator: %s", res)
	}

	var toChainId uint64 = config.DefConfig.HecoChainID
	{
		res, err := invoker.GetProxyHash(neoLock.Bytes(), toChainId)
		if err != nil {
			log.Errorf("GetProxyHash err: %v", err)
			return
		}
		log.Infof("neoLockProxy: %x, toChainId: %d, toChainProxyHash (little) : %s", neoLock.Bytes(), toChainId, res)
	}

	fromAssetHashs := []string{
		"0xf46719e2d16bf50cddcef9d4bbfece901f73cbb6",
	}
	{
		froms := make([][]byte, 0)
		for _, from := range fromAssetHashs {
			f, _ := ParseNeoAddr(from)
			froms = append(froms, f)
		}
		res, err := invoker.GetAssetHashs(neoLock.Bytes(), toChainId, froms)
		if err != nil {
			log.Errorf("GetProxyHash err: %v", err)
			return
		}
		bals, err := invoker.GetAssetBalances(neoLock.Bytes(), froms)
		if err != nil {
			log.Errorf("GetAssetBalances err: %v", err)
			return
		}

		for i, _ := range fromAssetHashs {
			log.Infof("neoLockProxy GetAssetHashs, toChainId: %d, from: %s, proxyBalance: %s, to(little): %s", toChainId, fromAssetHashs[i], bals[i].String(), res[i])
		}
	}
}

func Test_SendTxFromNeo(t *testing.T) {
	config.DefConfig.Init("./config_test.json")
	invoker, err := NewNeoInvoker()

	if err != nil {
		t.Fatal(err)
	}

	neoLockAddr, err := ParseNeoAddr(config.DefConfig.NeoLockProxy)
	if err != nil {
		return
	}
	hrc20AddrOnNeo, err := ParseNeoAddr(config.DefConfig.NeoHrc20)
	if err != nil {
		return
	}

	fromAsset := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: hrc20AddrOnNeo,
	}
	rawFrom, err := helper.AddressToScriptHash(invoker.Acc.Address)
	if err != nil {
		return
	}
	fromAddr := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: rawFrom.Bytes(),
	}
	toChainId := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *big.NewInt(int64(config.DefConfig.HecoChainID)),
	}
	toAddr := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: common.ADDRESS_EMPTY[:],
	}
	amt := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *big.NewInt(int64(1)),
	}
	tb := tx.NewTransactionBuilder(config.DefConfig.NeoUrl)
	sb := sc.NewScriptBuilder()
	sb.MakeInvocationScript(neoLockAddr, "lock", []sc.ContractParameter{fromAsset, fromAddr, toChainId, toAddr, amt})
	script := sb.ToArray()

	itx, err := tb.MakeInvocationTransaction(script, rawFrom, nil, rawFrom, helper.Zero, helper.Zero)
	if err != nil && err.Error() != "" {
		return
	}
	err = tx.AddSignature(itx, invoker.Acc.KeyPair)
	if err != nil {
		return
	}
	response := tb.Client.SendRawTransaction(itx.RawTransactionString())
	if response.HasError() {
		return
	}

	log.Infof("successful to send %d hrc20 from neo to heco", itx.HashString())
}
