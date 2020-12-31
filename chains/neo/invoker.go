package neo

import (
	"encoding/hex"
	"fmt"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/rpc"
	"github.com/joeqian10/neo-gogogo/sc"
	"github.com/joeqian10/neo-gogogo/tx"
	"github.com/joeqian10/neo-gogogo/wallet"
	"github.com/ontio/ontology/common"
	"github.com/polynetwork/poly-io-test/config"
	"github.com/polynetwork/poly-io-test/log"
	"math/big"
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
	if config.DefConfig.NeoWallet != "" && config.DefConfig.NeoWalletPwd != "" {
		invoker.Acc = GetAccountByPassword(config.DefConfig.NeoWallet, config.DefConfig.NeoWalletPwd)
		if invoker.Acc == nil {
			log.Errorf("GetAccountByPassword to obtain pwd error")
			err = fmt.Errorf("GetAccountByPassword to obtain pwd error")
			return
		}
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

func (this *NeoInvoker) BindProxyHash(neoLockProxy []byte, toChainId uint64, toProxyHash []byte) error {
	from, err := ParseNeoAddr(this.Acc.Address)
	if err != nil {
		return fmt.Errorf("[BindProxyHash], ParseNeoAddr acct: %s,  err: %v", this.Acc.Address, err)
	}
	fromUint160, err := helper.UInt160FromBytes(from)
	if err != nil {
		return fmt.Errorf("[BindProxyHash], Uint160FromBytes err: %v", err)
	}
	tci := big.NewInt(int64(toChainId))
	toChainIdValue := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *tci,
	}
	toProxyHashValue := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: toProxyHash,
	}
	// build script
	scriptBuilder := sc.NewScriptBuilder()
	args := []sc.ContractParameter{toChainIdValue, toProxyHashValue}
	scriptBuilder.MakeInvocationScript(neoLockProxy, "bindProxyHash", args)
	script := scriptBuilder.ToArray()

	// create an InvocationTransaction
	tb := tx.NewTransactionBuilder(this.Cli.Endpoint.String())

	sysFee := helper.Fixed8FromFloat64(0)
	netFee := helper.Fixed8FromFloat64(0)

	itx, err := tb.MakeInvocationTransaction(script, fromUint160, nil, fromUint160, sysFee, netFee)
	if err != nil {
		return fmt.Errorf("[BindProxyHash] tb.MakeInvocationTransaction error: %s", err)
	}
	// sign transaction
	err = tx.AddSignature(itx, this.Acc.KeyPair)
	if err != nil {
		return fmt.Errorf("[BindProxyHash] tx.AddSignature error: %s", err)
	}

	rawTxString := itx.RawTransactionString()

	// send the raw transaction
	response := this.Cli.SendRawTransaction(rawTxString)
	if response.HasError() {
		return fmt.Errorf("[BindProxyHash] SendRawTransaction error: %s,  RawTransactionString: %s",
			response.ErrorResponse.Error.Message, rawTxString)
	}

	log.Infof("Neo bindProxyHash, txHash: %s", itx.HashString())
	WaitNeoTx(this.Cli, itx.Hash)

	return nil
}

func (this *NeoInvoker) BindAssetHash(neoLockProxy []byte, fromAssetHash []byte, toChainId uint64, toAssetHash []byte) (string, error) {
	from, err := ParseNeoAddr(this.Acc.Address)
	if err != nil {
		return "", fmt.Errorf("[BindAssetHash], ParseNeoAddr acct: %s,  err: %v", this.Acc.Address, err)
	}
	fromUint160, err := helper.UInt160FromBytes(from)
	if err != nil {
		return "", fmt.Errorf("[BindAssetHash], Uint160FromBytes err: %v", err)
	}
	fromAssetHashValue := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: fromAssetHash,
	}
	toChainIdValue := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *big.NewInt(int64(toChainId)),
	}
	toAssetHashValue := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: toAssetHash,
	}
	// build script
	scriptBuilder := sc.NewScriptBuilder()
	args := []sc.ContractParameter{fromAssetHashValue, toChainIdValue, toAssetHashValue}
	scriptBuilder.MakeInvocationScript(neoLockProxy, "bindAssetHash", args)
	script := scriptBuilder.ToArray()

	// create an InvocationTransaction
	tb := tx.NewTransactionBuilder(this.Cli.Endpoint.String())

	sysFee := helper.Fixed8FromFloat64(0)
	netFee := helper.Fixed8FromFloat64(0)

	itx, err := tb.MakeInvocationTransaction(script, fromUint160, nil, fromUint160, sysFee, netFee)
	if err != nil {
		return "", fmt.Errorf("[BindAssetHash] tb.MakeInvocationTransaction error: %s", err)
	}
	// sign transaction
	err = tx.AddSignature(itx, this.Acc.KeyPair)
	if err != nil {
		return "", fmt.Errorf("[BindAssetHash] tx.AddSignature error: %s", err)
	}

	rawTxString := itx.RawTransactionString()

	// send the raw transaction
	response := this.Cli.SendRawTransaction(rawTxString)
	if response.HasError() {
		return "", fmt.Errorf("[BindAssetHash] SendRawTransaction error: %s,  RawTransactionString: %s",
			response.ErrorResponse.Error.Message, rawTxString)
	}
	log.Infof("Neo bindAssetHash, txHash: %s", itx.HashString())
	WaitNeoTx(this.Cli, itx.Hash)

	return itx.HashString(), nil
}

func (this *NeoInvoker) GetProxyOperator(neoLockProxy []byte) (string, error) {
	scriptBuilder := sc.NewScriptBuilder()
	args := []sc.ContractParameter{}
	scriptBuilder.MakeInvocationScript(neoLockProxy, "getOperator", args)
	script := scriptBuilder.ToArray()

	// create an InvocationTransaction
	response := this.Cli.InvokeScript(helper.BytesToHex(script), "0000000000000000000000000000000000000000")
	if response.HasError() || response.Result.State == "FAULT" {
		log.Errorf("invoke script error: %s", response.Error.Message)
		return "", fmt.Errorf("[GetProxyOperator], InvokeScript err: %v", response.Error)
	}
	for _, stack := range response.Result.Stack {
		stack.Convert()

		if stack.Type == "ByteArray" {
			x, e := hex.DecodeString(stack.Value.(string))
			if e != nil {
				return "", fmt.Errorf("bytearray: %s, decodestring err: %v", stack.Value.(string), e)
			}
			base58Addr, _ := common.AddressParseFromBytes(x)
			return base58Addr.ToBase58(), nil
		}
	}
	return "", fmt.Errorf("operator not found")
}

func (this *NeoInvoker) GetProxyHash(neoLockProxy []byte, toChainId uint64) (string, error) {
	scriptBuilder := sc.NewScriptBuilder()
	args := []sc.ContractParameter{
		sc.ContractParameter{
			Type:  sc.Integer,
			Value: *big.NewInt(int64(toChainId)),
		},
	}
	scriptBuilder.MakeInvocationScript(neoLockProxy, "getProxyHash", args)
	script := scriptBuilder.ToArray()

	// create an InvocationTransaction
	response := this.Cli.InvokeScript(helper.BytesToHex(script), "0000000000000000000000000000000000000000")
	if response.HasError() || response.Result.State == "FAULT" {
		log.Errorf("invoke script error: %s", response.Error.Message)
		return "", fmt.Errorf("[GetProxyHash], InvokeScript err: %v", response.Error)
	}
	for _, stack := range response.Result.Stack {
		stack.Convert()
		if stack.Type == "ByteArray" {
			return stack.Value.(string), nil
		}
	}
	return "", fmt.Errorf("GetProxyHash not found")
}

func (this *NeoInvoker) GetAssetHashs(neoLockProxy []byte, toChainId uint64, fromAssetHashs [][]byte) ([]string, error) {
	scriptBuilder := sc.NewScriptBuilder()

	for _, from := range fromAssetHashs {
		scriptBuilder.MakeInvocationScript(neoLockProxy, "getAssetHash", []sc.ContractParameter{
			sc.ContractParameter{
				Type:  sc.ByteArray,
				Value: from,
			},
			sc.ContractParameter{
				Type:  sc.Integer,
				Value: *big.NewInt(int64(toChainId)),
			},
		})
	}
	script := scriptBuilder.ToArray()

	// create an InvocationTransaction
	response := this.Cli.InvokeScript(helper.BytesToHex(script), "0000000000000000000000000000000000000000")
	if response.HasError() || response.Result.State == "FAULT" {
		log.Errorf("invoke script error: %s", response.Error.Message)
		return nil, fmt.Errorf("[GetProxyHash], InvokeScript err: %v", response.Error)
	}
	res := make([]string, len(fromAssetHashs))
	for i, stack := range response.Result.Stack {
		stack.Convert()

		if stack.Type == "ByteArray" {
			res[i] = stack.Value.(string)
		}
	}
	return res, nil
}

func (this *NeoInvoker) GetAssetBalances(neoLockProxy []byte, fromAssetHashs [][]byte) ([]*big.Int, error) {
	scriptBuilder := sc.NewScriptBuilder()

	for _, from := range fromAssetHashs {
		scriptBuilder.MakeInvocationScript(from, "balanceOf", []sc.ContractParameter{
			sc.ContractParameter{
				Type:  sc.ByteArray,
				Value: neoLockProxy,
			},
		})
	}
	script := scriptBuilder.ToArray()

	// create an InvocationTransaction
	response := this.Cli.InvokeScript(helper.BytesToHex(script), "0000000000000000000000000000000000000000")
	if response.HasError() || response.Result.State == "FAULT" {
		log.Errorf("invoke script error: %s", response.Error.Message)
		return nil, fmt.Errorf("[GetAssetBalances], InvokeScript err: %v", response.Error)
	}
	res := make([]*big.Int, len(fromAssetHashs))
	for i, stack := range response.Result.Stack {
		stack.Convert()

		if stack.Type == "ByteArray" {
			hexNumBs, _ := hex.DecodeString(stack.Value.(string))
			res[i] = common.BigIntFromNeoBytes(hexNumBs)
		}
	}
	return res, nil
}

func (this *NeoInvoker) GetStorage(neoLockProxy []byte, key string) (string, error) {
	addr, _ := common.AddressParseFromBytes(neoLockProxy)

	resp := this.Cli.GetStorage(addr.ToHexString(), key)
	if resp.HasError() {
		return "", fmt.Errorf("resp.Error.Message: %s", resp.Error.Message)
	}
	return resp.Result, nil
}

func ParseNeoAddr(s string) ([]byte, error) {
	if strings.Contains(s, "0x") {
		rb, err := hex.DecodeString(strings.TrimPrefix(s, "0x"))
		if err != nil {
			return nil, fmt.Errorf("ParseNeoAddr err: %v", err)
		}
		return common.ToArrayReverse(rb), nil
	} else if strings.Contains(s, "A") {
		addr, err := common.AddressFromBase58(s)
		if err != nil {
			return nil, fmt.Errorf("AddressFromBase58 err: %v", err)
		}
		return addr[:], nil
	}
	return hex.DecodeString(s)
}
func GetAccountByPassword(walletPath, pwd string) *wallet.Account {
	// open the NEO wallet
	w, err := wallet.NewWalletFromFile(walletPath) //
	if err != nil {
		log.Errorf("[GetAccountByPassword] Failed to open NEO wallet")
		return nil
	}

	if pwd == "" {
		log.Errorf("pls provide neo wallet pwd")
		return nil
	}
	err = w.DecryptAll(pwd)
	if err != nil {
		log.Errorf("[GetAccountByPassword] Failed to decrypt NEO account")
		return nil
	}
	if len(w.Accounts) == 0 {
		log.Errorf("[GetAccountByPassword] empty account")
		return nil
	}
	return w.Accounts[0]
}
