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
package cosmos

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcutil"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/mintkey"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ontio/ontology-go-sdk"
	common2 "github.com/ontio/ontology/common"
	"github.com/polynetwork/cosmos-poly-module/btcx"
	"github.com/polynetwork/cosmos-poly-module/headersync"
	"github.com/polynetwork/cosmos-poly-module/lockproxy"
	"github.com/polynetwork/poly-io-test/config"
	"github.com/polynetwork/poly-io-test/log"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/bytes"
	"github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/rpc/client/http"
	"github.com/tendermint/tendermint/rpc/core/types"
	"io/ioutil"
	"strings"
	"sync"
	"time"
)

type CosmosInvoker struct {
	RpcCli        *http.HTTP
	Acc           *CosmosAcc
	CosmosChainId string
	CMGas         uint64
	CMFees        types.Coins
	CMCdc         *codec.Codec
}

func NewCosmosInvoker() (*CosmosInvoker, error) {
	var (
		err      error
		gasPrice types.DecCoins
	)
	invoker := &CosmosInvoker{}
	conf := config.DefConfig

	invoker.CosmosChainId = conf.CMChainId
	switch conf.CMChainId {
	case "switcheochain":
		conf := types.GetConfig()
		conf.SetBech32PrefixForAccount("swth", "swthpub")
		conf.SetBech32PrefixForValidator("swthvaloper", "swthvaloperpub")
		conf.SetBech32PrefixForConsensusNode("swthvalcons", "swthvalconspub")
		conf.Seal()
	}

	invoker.RpcCli, err = http.New(conf.CMRpcUrl, "/websocket")
	if err != nil {
		return nil, err
	}
	invoker.CMCdc = NewCodec()

	invoker.Acc, err = NewCosmosAcc(conf.CMWalletPath, conf.CMWalletPwd, invoker.RpcCli, invoker.CMCdc)
	if err != nil {
		return nil, err
	}
	invoker.CMGas = conf.CMGas
	if gasPrice, err = types.ParseDecCoins(conf.CMGasPrice); err != nil {
		return nil, err
	}
	if invoker.CMFees, err = CalcCosmosFees(gasPrice, conf.CMGas); err != nil {
		return nil, err
	}

	return invoker, nil
}

func (invoker *CosmosInvoker) sendCosmosTx(msgs []types.Msg) (*coretypes.ResultBroadcastTx, error) {
	toSign := auth.StdSignMsg{
		Sequence:      invoker.Acc.Seq.GetAndAdd(),
		AccountNumber: invoker.Acc.AccNum,
		ChainID:       invoker.CosmosChainId,
		Msgs:          msgs,
		Fee:           auth.NewStdFee(invoker.CMGas, invoker.CMFees),
	}
	sig, err := invoker.Acc.PrivateKey.Sign(toSign.Bytes())
	if err != nil {
		return nil, fmt.Errorf("failed to sign raw tx: (error: %v, raw tx: %x)", err, toSign.Bytes())
	}

	tx := auth.NewStdTx(msgs, toSign.Fee, []auth.StdSignature{{invoker.Acc.PrivateKey.PubKey(),
		sig}}, toSign.Memo)
	encoder := auth.DefaultTxEncoder(invoker.CMCdc)
	rawTx, err := encoder(tx)
	if err != nil {
		return nil, fmt.Errorf("failed to encode signed tx: %v", err)
	}

	var res *coretypes.ResultBroadcastTx
	for {
		res, err = invoker.RpcCli.BroadcastTxSync(rawTx)
		if err != nil {
			return nil, fmt.Errorf("failed to broadcast tx: (error: %v, raw tx: %x)", err, rawTx)
		}
		if res.Code != 0 {
			if strings.Contains(res.Log, "verify correct account sequence and chain-id") {
				time.Sleep(time.Second)
				continue
			}
			return nil, fmt.Errorf("failed to check tx: (code: %d, log: %s)", res.Code, res.Log)
		} else {
			break
		}
	}

	return res, nil
}

func (invoker *CosmosInvoker) TransferCoins(toAddr, denom string, amt int64) (*coretypes.ResultBroadcastTx, error) {
	toAcc, err := types.AccAddressFromBech32(toAddr)
	if err != nil {
		return nil, err
	}
	resTx, err := invoker.sendCosmosTx([]types.Msg{bank.NewMsgSend(invoker.Acc.Acc, toAcc, types.NewCoins(types.NewInt64Coin(denom, amt)))})
	if err != nil {
		return nil, err
	}
	invoker.WaitTx(resTx.Hash)
	return resTx, nil
}

// TODO: not working for now. How can we get an PrivKeyEd25519 ?
func (invoker *CosmosInvoker) CreateValidator(keyFile, stateFile, coin string, amt int64) (*coretypes.ResultBroadcastTx, error) {
	pri := privval.LoadOrGenFilePV(keyFile, stateFile)
	pk, err := pri.GetPubKey()
	if err != nil {
		return nil, err
	}
	valAcc := types.ValAddress(invoker.Acc.Acc)
	var res *coretypes.ResultBroadcastTx
	//res, err = invoker.TransferCoins(acc.String(), coin, 2 * amt)
	//if err != nil {
	//	return nil, err
	//}
	//log.Infof("send %d %s to %s from %s: txhash %s", 2 * amt, coin, acc.String(), invoker.Acc.Acc.String(), res.Hash.String())

	cRate := staking.NewCommissionRates(types.MustNewDecFromStr("0.20"), types.OneDec(), types.MustNewDecFromStr("0.10"))
	msg := staking.NewMsgCreateValidator(valAcc, pk, types.NewInt64Coin(coin, amt),
		staking.NewDescription(fmt.Sprintf("add_val: %s", valAcc.String()), "", "", "", ""),
		cRate, types.OneInt())
	res, err = invoker.sendCosmosTx([]types.Msg{msg})
	if err != nil {
		return nil, err
	}
	invoker.WaitTx(res.Hash)
	return res, nil
}

func (invoker *CosmosInvoker) DelegateValidator(coin string, amt int64) (*coretypes.ResultBroadcastTx, error) {
	res, err := invoker.sendCosmosTx([]types.Msg{
		staking.NewMsgDelegate(invoker.Acc.Acc, types.ValAddress(invoker.Acc.Acc), types.NewInt64Coin(coin, amt)),
	})
	if err != nil {
		return nil, err
	}
	invoker.WaitTx(res.Hash)
	return res, nil
}

func (invoker *CosmosInvoker) SyncPolyGenesisHdr(syner types.AccAddress, rawHdr []byte) (*coretypes.ResultBroadcastTx, error) {
	param := &headersync.MsgSyncGenesisParam{
		Syncer:        invoker.Acc.Acc,
		GenesisHeader: hex.EncodeToString(rawHdr),
	}
	resTx, err := invoker.sendCosmosTx([]types.Msg{param})
	if err != nil {
		return nil, err
	}

	invoker.WaitTx(resTx.Hash)
	return resTx, nil
}

func (invoker *CosmosInvoker) CreateLockProxy() (*coretypes.ResultBroadcastTx, error) {
	msg := lockproxy.NewMsgCreateLockProxy(invoker.Acc.Acc)
	res, err := invoker.sendCosmosTx([]types.Msg{msg})
	if err != nil {
		return nil, err
	}
	config.DefConfig.CMLockProxy = hex.EncodeToString(invoker.Acc.Acc.Bytes())
	return res, nil
}

func (invoker *CosmosInvoker) CreateAsset(denom, amt string, lockProxy []byte) (*coretypes.ResultBroadcastTx, error) {
	amount, _ := types.NewIntFromString(amt)
	msg := lockproxy.NewMsgCreateCoinAndDelegateToProxy(invoker.Acc.Acc, types.NewCoin(denom, amount), lockProxy)
	res, err := invoker.sendCosmosTx([]types.Msg{msg})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (invoker *CosmosInvoker) BindProxy(toChainId uint64, proxy []byte) (*coretypes.ResultBroadcastTx, error) {
	res, err := invoker.sendCosmosTx([]types.Msg{lockproxy.NewMsgBindProxyHash(invoker.Acc.Acc, toChainId, proxy)})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (invoker *CosmosInvoker) BtcxBindAsset(sourceAsset string, toChainId uint64, toAssetHash []byte) (*coretypes.ResultBroadcastTx, error) {
	res, err := invoker.sendCosmosTx([]types.Msg{btcx.NewMsgBindAssetHash(invoker.Acc.Acc, sourceAsset, toChainId, toAssetHash)})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (invoker *CosmosInvoker) BindAsset(sourceAsset string, toChainId uint64, toAssetHash []byte) (*coretypes.ResultBroadcastTx, error) {
	res, err := invoker.sendCosmosTx([]types.Msg{lockproxy.NewMsgBindAssetHash(invoker.Acc.Acc, sourceAsset, toChainId, toAssetHash)})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (invoker *CosmosInvoker) SetupBtcx(denom, redeem string) error {
	rawRdm, err := hex.DecodeString(redeem)
	if err != nil {
		return err
	}
	tx, err := invoker.sendCosmosTx([]types.Msg{btcx.NewMsgCreateDenom(invoker.Acc.Acc, denom, redeem)})
	if err != nil && !strings.Contains(err.Error(), "already") {
		return err
	}
	invoker.WaitTx(tx.Hash)
	// bind btcx asset hash
	tx, err = invoker.BtcxBindAsset(denom, config.DefConfig.BtcChainID, btcutil.Hash160(rawRdm))
	if err != nil {
		return fmt.Errorf("asset: %s, BtcxBindAsset error: %v", denom, err)
	}
	tx, err = invoker.BtcxBindAsset(denom, config.DefConfig.EthChainID, common.HexToAddress(config.DefConfig.BtceContractAddress).Bytes())
	if err != nil {
		return fmt.Errorf("asset: %s, BtcxBindAsset error: %v", denom, err)
	}
	btcOnt, _ := common2.AddressFromHexString(config.DefConfig.BtcoContractAddress)
	tx, err = invoker.BtcxBindAsset(denom, config.DefConfig.OntChainID, btcOnt[:])
	if err != nil {
		return fmt.Errorf("asset: %s, BtcxBindAsset error: %v", denom, err)
	}
	invoker.WaitTx(tx.Hash)

	return nil
}

// eth, erc20, ont, ong, oep4
func (invoker *CosmosInvoker) SetupAllAssets(proxy []byte) error {
	// create coins and deletage to lockproxy
	tx, err := invoker.CreateAsset(config.CM_ETHX, "1000000000000000000000000", proxy)
	if err != nil && !strings.Contains(err.Error(), "already") {
		return fmt.Errorf("create ethx failed: %v", err)
	}
	{
		// bind ethx asset hash
		var err error
		_, err = invoker.BindAsset(config.CM_ETHX, config.DefConfig.EthChainID, common.HexToAddress("0x0000000000000000000000000000000000000000").Bytes())
		if err != nil {
			return fmt.Errorf("asset: %s, bindAsset error: %v", config.CM_ETHX, err)
		}
		ontEth, _ := common2.AddressFromHexString(config.DefConfig.OntEth)
		_, err = invoker.BindAsset(config.CM_ETHX, config.DefConfig.OntChainID, ontEth[:])
		if err != nil {
			return fmt.Errorf("asset: %s, bindAsset error: %v", config.CM_ETHX, err)
		}
	}
	tx, err = invoker.CreateAsset(config.CM_ERC20, "1000000000000000000000000", proxy)
	if err != nil && !strings.Contains(err.Error(), "already") {
		return fmt.Errorf("create erc20x failed: %v", err)
	}
	{
		// bind erc20x asset hash
		var err error
		_, err = invoker.BindAsset(config.CM_ERC20, config.DefConfig.EthChainID, common.HexToAddress(config.DefConfig.EthErc20).Bytes())
		if err != nil {
			return fmt.Errorf("asset: %s, bindAsset error: %v", config.CM_ERC20, err)
		}
		ontErc20, _ := common2.AddressFromHexString(config.DefConfig.OntErc20)
		_, err = invoker.BindAsset(config.CM_ERC20, config.DefConfig.OntChainID, ontErc20[:])
		if err != nil {
			return fmt.Errorf("asset: %s, bindAsset error: %v", config.CM_ERC20, err)
		}
	}
	tx, err = invoker.CreateAsset(config.CM_ONT, "1000000000", proxy)
	if err != nil && !strings.Contains(err.Error(), "already") {
		return fmt.Errorf("create ontx failed: %v", err)
	}
	{
		// bind ontx asset hash
		var err error
		_, err = invoker.BindAsset(config.CM_ONT, config.DefConfig.EthChainID, common.HexToAddress(config.DefConfig.EthOntx).Bytes())
		if err != nil {
			return fmt.Errorf("asset: %s, bindAsset error: %v", config.CM_ONT, err)
		}
		_, err = invoker.BindAsset(config.CM_ONT, config.DefConfig.OntChainID, ontology_go_sdk.ONT_CONTRACT_ADDRESS[:])
		if err != nil {
			return fmt.Errorf("asset: %s, bindAsset error: %v", config.CM_ONT, err)
		}
	}
	tx, err = invoker.CreateAsset(config.CM_ONG, "1000000000000000000", proxy)
	if err != nil && !strings.Contains(err.Error(), "already") {
		return fmt.Errorf("create ongx failed: %v", err)
	}
	{
		// bind ongx asset hash
		var err error
		_, err = invoker.BindAsset(config.CM_ONG, config.DefConfig.EthChainID, common.HexToAddress(config.DefConfig.EthOngx).Bytes())
		if err != nil {
			return fmt.Errorf("asset: %s, bindAsset error: %v", config.CM_ONG, err)
		}
		_, err = invoker.BindAsset(config.CM_ONG, config.DefConfig.OntChainID, ontology_go_sdk.ONG_CONTRACT_ADDRESS[:])
		if err != nil {
			return fmt.Errorf("asset: %s, bindAsset error: %v", config.CM_ONG, err)
		}
	}
	tx, err = invoker.CreateAsset(config.CM_OEP4, "10000000000000", proxy)
	if err != nil && !strings.Contains(err.Error(), "already") {
		return fmt.Errorf("create oep4x failed: %v", err)
	}
	{
		// bind oep4 asset hash
		var err error
		_, err = invoker.BindAsset(config.CM_OEP4, config.DefConfig.EthChainID, common.HexToAddress(config.DefConfig.EthOep4).Bytes())
		if err != nil {
			return fmt.Errorf("asset: %s, bindAsset error: %v", config.CM_OEP4, err)
		}
		ontOep4, _ := common2.AddressFromHexString(config.DefConfig.OntOep4)
		_, err = invoker.BindAsset(config.CM_OEP4, config.DefConfig.OntChainID, ontOep4[:])
		if err != nil {
			return fmt.Errorf("asset: %s, bindAsset error: %v", config.CM_OEP4, err)
		}
	}
	// bind eth proxy hash
	ontProxyHash, _ := common2.AddressFromHexString(config.DefConfig.OntLockProxy)
	_, err = invoker.BindProxy(config.DefConfig.OntChainID, ontProxyHash[:])
	if err != nil {
		return fmt.Errorf("BindProxy ethx failed: %v", err)
	}
	// bind ont proxy hash
	_, err = invoker.BindProxy(config.DefConfig.EthChainID, common.HexToAddress(config.DefConfig.EthLockProxy).Bytes())
	if err != nil {
		return fmt.Errorf("BindProxy ethx failed: %v", err)
	}
	invoker.WaitTx(tx.Hash)
	return nil
}

func (invoker *CosmosInvoker) WaitTx(txhash bytes.HexBytes) {
	tick := time.NewTicker(time.Second)
	for range tick.C {
		res, err := invoker.RpcCli.Tx(txhash, false)
		if err == nil && res.Height > 0 {
			break
		}
	}
}

func (invoker *CosmosInvoker) SendAsset(asset string, toChainId uint64, value int64, toAddr, lockProxy []byte) (*coretypes.ResultBroadcastTx, error) {
	var msg types.Msg
	if asset == config.CM_BTCX {
		msg = btcx.NewMsgLock(invoker.Acc.Acc, asset, toChainId, toAddr, types.NewInt(value))
	} else {
		msg = lockproxy.NewMsgLock(lockProxy, invoker.Acc.Acc, asset, toChainId, toAddr, types.NewInt(value))
	}
	res, err := invoker.sendCosmosTx([]types.Msg{msg})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (invoker *CosmosInvoker) GetAccInfo() (string, error) {
	param := bank.NewQueryBalanceParams(invoker.Acc.Acc)
	raw, err := invoker.CMCdc.MarshalJSON(param)
	if err != nil {
		return "", err
	}
	res, err := invoker.RpcCli.ABCIQuery("custom/bank/balances", raw)
	if err != nil {
		return "", err
	}
	coins := types.NewCoins()
	if err := invoker.CMCdc.UnmarshalJSON(res.Response.GetValue(), &coins); err != nil {
		return "", err
	}
	return fmt.Sprintf("COSMOS: acc: %s, asset: [ %s ]", invoker.Acc.Acc.String(), coins.String()), nil
}

type CosmosAcc struct {
	Acc        types.AccAddress
	PrivateKey crypto.PrivKey
	Seq        *CosmosSeq
	AccNum     uint64
}

func NewCosmosAcc(wallet, pwd string, cli *http.HTTP, cdc *codec.Codec) (*CosmosAcc, error) {
	acc := &CosmosAcc{}
	bz, err := ioutil.ReadFile(wallet)
	if err != nil {
		return nil, err
	}

	privKey, _, err := mintkey.UnarmorDecryptPrivKey(string(bz), string(pwd))
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt private key: v", err)
	}

	acc.PrivateKey = privKey
	acc.Acc = types.AccAddress(privKey.PubKey().Address().Bytes())
	log.Infof("cosmos address: %s", acc.Acc.String())
	var eAcc exported.Account
	rawParam, err := cdc.MarshalJSON(auth.NewQueryAccountParams(acc.Acc))
	if err != nil {
		return nil, err
	}
	res, err := cli.ABCIQuery("/custom/acc/account", rawParam)
	if err != nil {
		return nil, err
	}
	if !res.Response.IsOK() {
		return nil, fmt.Errorf("failed to get response for accout-query: %v", res.Response)
	}
	if err := cdc.UnmarshalJSON(res.Response.Value, &eAcc); err != nil {
		return nil, fmt.Errorf("unmarshal query-account-resp failed, err: %v", err)
	}
	acc.Seq = &CosmosSeq{
		lock: sync.Mutex{},
		val:  eAcc.GetSequence(),
	}
	acc.AccNum = eAcc.GetAccountNumber()

	return acc, nil
}

type CosmosSeq struct {
	lock sync.Mutex
	val  uint64
}

func (seq *CosmosSeq) GetAndAdd() uint64 {
	seq.lock.Lock()
	defer func() {
		seq.val += 1
		seq.lock.Unlock()
	}()
	return seq.val
}
