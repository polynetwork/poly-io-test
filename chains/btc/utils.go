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
package btc

import (
	"bytes"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	types2 "github.com/cosmos/cosmos-sdk/types"
	"github.com/ontio/ontology-crypto/keypair"
	"github.com/ontio/ontology-crypto/signature"
	"github.com/ontio/ontology/common"
	"github.com/polynetwork/poly-go-sdk"
	"github.com/polynetwork/poly-io-test/config"
	"github.com/polynetwork/poly-io-test/log"
	"github.com/polynetwork/poly/account"
	common2 "github.com/polynetwork/poly/common"
	"github.com/polynetwork/poly/core/types"
	"github.com/polynetwork/poly/native/service/cross_chain_manager/btc"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type BtcSigner struct {
	WIF     *btcutil.WIF
	Address string
}

func NewBtcSigner(privateKey string) (*BtcSigner, error) {
	wif, err := btcutil.DecodeWIF(privateKey)
	if err != nil {
		return nil, fmt.Errorf("NewBtcSigner, failed to decode wif: %v", err)
	}

	addrPubk, err := btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeCompressed(), &chaincfg.TestNet3Params)
	if err != nil {
		return nil, fmt.Errorf("sendBtcCross, Failed to new an address pubkey: %v", err)
	}
	address := addrPubk.EncodeAddress()
	return &BtcSigner{
		WIF:     wif,
		Address: address,
	}, nil
}

func RandomInt64(min, max int64) int64 {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return r.Int63n(max-min) + min
}

type BtcItem struct {
	Item *ToBtcItem
	Mtx  *wire.MsgTx
}

type ToBtcItem struct {
	OntTxid string
	BtcAddr string
	Value   int64
}

type FromBtcItem struct {
	BtcTxid string
	OntAddr string
	Value   int64
}

func (item *FromBtcItem) String() string {
	return fmt.Sprintf("(btc-txid:%s, value:%d, ont-addr:%s)", item.BtcTxid, item.Value, item.OntAddr)
}

type Utxo struct {
	Txid         string
	Vout         uint32
	ScriptPubKey string
	Amount       int64
	Confs        int64
}

func (u *Utxo) String() string {
	return "\nutxo: " + u.Txid + ":" + strconv.FormatUint(uint64(u.Vout), 10) + "\n\tScriptPubKey: " +
		u.ScriptPubKey + "\n\tAmount: " + strconv.FormatInt(u.Amount, 10) + "\n\tConfs: " +
		strconv.FormatInt(u.Confs, 10) + "\n"
}

func SelectUtxos(utxos []*Utxo, value int64) ([]*Utxo, int64, error) {
	if value <= 0 {
		return nil, -1, fmt.Errorf("value must be positive")
	}

	ul := make(UtxoList, 0)
	for _, v := range utxos {
		ul = append(ul, &UtxoItem{
			key: v,
			val: v.Confs * v.Amount,
		})
	}

	sort.Sort(ul)
	if !sort.IsSorted(ul) {
		return nil, -1, fmt.Errorf("ItemList not sorted")
	}

	selected := make([]*Utxo, 0)
	selectedVal := int64(0)
	for i := len(ul) - 1; i >= 0; i-- {
		selected = append(selected, ul[i].key)
		if selectedVal += ul[i].key.Amount; selectedVal >= value {
			return selected, selectedVal, nil
		}
	}
	return nil, selectedVal, fmt.Errorf("not enough utxo for %d, all we have is %d", value, selectedVal)
}

type UtxoItem struct {
	key *Utxo
	val int64
}

type UtxoList []*UtxoItem

func (ul UtxoList) Swap(i, j int) {
	ul[i], ul[j] = ul[j], ul[i]
}

func (ul UtxoList) Len() int {
	return len(ul)
}

func (ul UtxoList) Less(i, j int) bool {
	return ul[i].val < ul[j].val
}

type Request struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      int           `json:"id"`
}

type Response struct {
	Result interface{}       `json:"result"`
	Error  *btcjson.RPCError `json:"error"` //maybe wrong
	Id     int               `json:"id"`
}

// Get tx in block; Get proof;
type RestCli struct {
	Addr string
	Cli  *http.Client
}

func NewRestCli(addr, user, pwd string) *RestCli {
	return &RestCli{
		Cli: &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   5,
				DisableKeepAlives:     false,
				IdleConnTimeout:       time.Second * 300,
				ResponseHeaderTimeout: time.Second * 300,
				TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
				Proxy: func(req *http.Request) (*url.URL, error) {
					req.SetBasicAuth(user, pwd)
					return nil, nil
				},
			},
			Timeout: time.Second * 300,
		},
		Addr: addr,
	}
}

func (cli *RestCli) sendPostReq(req []byte) (*Response, error) {
	resp, err := cli.Cli.Post(cli.Addr, "application/json;charset=UTF-8",
		bytes.NewReader(req))
	if err != nil {
		return nil, fmt.Errorf("failed to post: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error:%s", err)
	}

	response := new(Response)
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}
	return response, nil
}

func (cli *RestCli) GetTxsInBlock(hash string) ([]*wire.MsgTx, string, error) {
	req, err := json.Marshal(Request{
		Jsonrpc: "1.0",
		Method:  "getblock",
		Params:  []interface{}{hash, false},
		Id:      1,
	})
	if err != nil {
		return nil, "", fmt.Errorf("failed to marshal request: %v", err)
	}

	resp, err := cli.sendPostReq(req)
	if err != nil {
		return nil, "", fmt.Errorf("failed to send post: %v", err)
	}
	if resp.Error != nil {
		return nil, "", fmt.Errorf("response shows failure: %v", resp.Error.Message)
	}
	bhex := resp.Result.(string)
	bb, err := hex.DecodeString(bhex)
	if err != nil {
		return nil, "", fmt.Errorf("failed to decode hex string: %v", err)
	}

	block := wire.MsgBlock{}
	err = block.BtcDecode(bytes.NewBuffer(bb), wire.ProtocolVersion, wire.LatestEncoding)
	if err != nil {
		return nil, "", fmt.Errorf("failed to decode block: %v", err)
	}

	return block.Transactions, block.Header.PrevBlock.String(), nil
}

func (cli *RestCli) GetCurrentHeightAndHash() (int32, string, error) {
	reqTips, err := json.Marshal(Request{
		Jsonrpc: "1.0",
		Method:  "getchaintips",
		Params:  nil,
		Id:      1,
	})
	if err != nil {
		return -1, "", fmt.Errorf("failed to marshal request: %v", err)
	}

	resp, err := cli.sendPostReq(reqTips)
	if err != nil {
		return -1, "", fmt.Errorf("failed to send post: %v", err)
	}
	if resp.Error != nil {
		return -1, "", fmt.Errorf("response shows failure: %v", resp.Error.Message)
	}

	m := resp.Result.([]interface{})[0].(map[string]interface{})
	return int32(m["height"].(float64)), m["hash"].(string), nil
}

func (cli *RestCli) BroadcastTx(tx string) (string, error) {
	req, err := json.Marshal(Request{
		Jsonrpc: "1.0",
		Method:  "sendrawtransaction",
		Params:  []interface{}{tx},
		Id:      1,
	})
	if err != nil {
		return "", fmt.Errorf("[BroadcastTx] failed to marshal request: %v", err)
	}

	resp, err := cli.sendPostReq(req)
	if err != nil {
		return "", fmt.Errorf("[BroadcastTx] failed to send post: %v", err)
	}
	if resp.Error != nil {
		return "", fmt.Errorf("[BroadcastTx] response shows failure: %v", resp.Error.Message)
	}

	return resp.Result.(string), nil
}

func (cli *RestCli) GetProof(txids []string) (string, error) {
	req, err := json.Marshal(Request{
		Jsonrpc: "1.0",
		Method:  "gettxoutproof",
		Params:  []interface{}{txids},
		Id:      1,
	})
	if err != nil {
		return "", fmt.Errorf("failed to get proof: %v", err)
	}

	resp, err := cli.sendPostReq(req)
	if err != nil {
		return "", fmt.Errorf("failed to send post: %v", err)
	}
	if resp.Error != nil {
		if resp.Error.Code == -5 && resp.Error.Message == "Transaction not yet in block" {
			return "", fmt.Errorf(resp.Error.Message + ". Please check the setting of bitcoin " +
				"node, need -txindex")
		}
		return "", fmt.Errorf("response shows failure: %v", resp.Error.Message)
	}

	return resp.Result.(string), nil
}

func ReverseHexString(b string) string {
	var result string
	for i := len(b) - 2; i >= 0; i -= 2 {
		result += b[i : i+2]
	}
	return result
}

func (cli *RestCli) GenerateToAddr(n int, addr string) ([]string, error) {
	req, err := json.Marshal(Request{
		Jsonrpc: "1.0",
		Method:  "generatetoaddress",
		Params:  []interface{}{n, addr},
		Id:      1,
	})
	if err != nil {
		return nil, fmt.Errorf("[GenerateToAddr] failed to marshal request: %v", err)
	}

	resp, err := cli.sendPostReq(req)
	if err != nil {
		return nil, fmt.Errorf("[GenerateToAddr] failed to send post: %v", err)
	}
	if resp.Error != nil {
		return nil, fmt.Errorf("[GenerateToAddr] response shows failure: %v", resp.Error.Message)
	}

	hashes := make([]string, 0)
	for _, v := range resp.Result.([]interface{}) {
		hashes = append(hashes, v.(string))
	}

	return hashes, nil
}

func (cli *RestCli) GetMempoolInfo() (int32, error) {
	req, err := json.Marshal(Request{
		Jsonrpc: "1.0",
		Method:  "getmempoolinfo",
		Params:  []interface{}{},
		Id:      1,
	})
	if err != nil {
		return -1, fmt.Errorf("[GetMempoolInfo] failed to marshal request: %v", err)
	}

	resp, err := cli.sendPostReq(req)
	if err != nil {
		return -1, fmt.Errorf("[GetMempoolInfo] failed to send post: %v", err)
	}
	if resp.Error != nil {
		return -1, fmt.Errorf("[GetMempoolInfo] response shows failure: %v", resp.Error.Message)
	}

	return int32(resp.Result.(map[string]interface{})["size"].(float64)), nil
}

func (cli *RestCli) GetBlockHeight(hash string) (int32, error) {
	req, err := json.Marshal(Request{
		Jsonrpc: "1.0",
		Method:  "getblock",
		Params:  []interface{}{hash},
		Id:      1,
	})
	if err != nil {
		return -1, fmt.Errorf("[GetBlockHeight] failed to marshal request: %v", err)
	}

	resp, err := cli.sendPostReq(req)
	if err != nil {
		return -1, fmt.Errorf("[GetBlockHeight] failed to send post: %v", err)
	}
	if resp.Error != nil {
		return -1, fmt.Errorf("[GetBlockHeight] response shows failure: %v", resp.Error.Message)
	}

	return int32(resp.Result.(map[string]interface{})["height"].(float64)), nil
}

func (cli *RestCli) ListUnspent(minConfs, maxConfs int64, addr string) ([]*Utxo, error) {
	req, err := json.Marshal(Request{
		Jsonrpc: "1.0",
		Method:  "listunspent",
		Params:  []interface{}{minConfs, maxConfs, []string{addr}},
		Id:      1,
	})
	if err != nil {
		return nil, fmt.Errorf("[ListUnspent] failed to marshal request: %v", err)
	}

	resp, err := cli.sendPostReq(req)
	if err != nil {
		return nil, fmt.Errorf("[ListUnspent] failed to send post: %v", err)
	}
	if resp.Error != nil {
		return nil, fmt.Errorf("[ListUnspent] response shows failure: %v", resp.Error.Message)
	}

	utxos := make([]*Utxo, 0)
	arr, ok := resp.Result.([]interface{})
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	for _, v := range arr {
		item := v.(map[string]interface{})

		amount := item["amount"].(float64)
		txid := item["txid"].(string)
		a, err := btcutil.NewAmount(amount)
		if err != nil {
			return nil, fmt.Errorf("failed to get amount for %f(txid:%s)", amount, txid)
		}
		utxos = append(utxos, &Utxo{
			Txid:         txid,
			Vout:         uint32(item["vout"].(float64)),
			ScriptPubKey: item["scriptPubKey"].(string),
			Amount:       int64(a),
			Confs:        int64(item["confirmations"].(float64)),
		})
	}

	return utxos, err
}

func (cli *RestCli) ImportAddress(addr string) error {
	req, err := json.Marshal(Request{
		Jsonrpc: "1.0",
		Method:  "importaddress",
		Params:  []interface{}{addr},
		Id:      1,
	})
	if err != nil {
		return fmt.Errorf("[ImportAddress] failed to marshal request: %v", err)
	}

	resp, err := cli.sendPostReq(req)
	if err != nil {
		return fmt.Errorf("[ImportAddress] failed to send post: %v", err)
	}
	if resp.Error != nil {
		return fmt.Errorf("[ImportAddress] response shows failure: %v", resp.Error.Message)
	}

	return nil
}

func (cli *RestCli) GetBlockCount() (int64, error) {
	req, err := json.Marshal(Request{
		Jsonrpc: "1.0",
		Method:  "getblockcount",
		Params:  []interface{}{},
		Id:      1,
	})
	if err != nil {
		return -1, fmt.Errorf("[GetBlockCount] failed to marshal request: %v", err)
	}

	resp, err := cli.sendPostReq(req)
	if err != nil {
		return -1, fmt.Errorf("[GetBlockCount] failed to send post: %v", err)
	}
	if resp.Error != nil {
		return -1, fmt.Errorf("[GetBlockCount] response shows failure: %v", resp.Error.Message)
	}

	return int64(resp.Result.(float64)), nil
}

func (cli *RestCli) SendRawTx(rawTx string) (string, error) {
	req, err := json.Marshal(Request{
		Jsonrpc: "1.0",
		Method:  "sendrawtransaction",
		Params:  []interface{}{rawTx},
		Id:      1,
	})
	if err != nil {
		return "", fmt.Errorf("[SendRawTx] failed to marshal request: %v", err)
	}

	resp, err := cli.sendPostReq(req)
	if err != nil {
		return "", fmt.Errorf("[SendRawTx] failed to send post: %v", err)
	}
	if resp.Error != nil {
		return "", fmt.Errorf("[SendRawTx] response shows failure: %v", resp.Error.Message)
	}

	return resp.Result.(string), nil
}

func (cli *RestCli) GetRawTransaction(txid string) (string, error) {
	req, err := json.Marshal(Request{
		Jsonrpc: "1.0",
		Method:  "getrawtransaction",
		Params:  []interface{}{txid},
		Id:      1,
	})
	if err != nil {
		return "", fmt.Errorf("[GetRawTransaction] failed to marshal request: %v", err)
	}

	resp, err := cli.sendPostReq(req)
	if err != nil {
		return "", fmt.Errorf("[GetRawTransaction] failed to send post: %v", err)
	}
	if resp.Error != nil {
		return "", fmt.Errorf("[GetRawTransaction] response shows failure: %v", resp.Error.Message)
	}

	return resp.Result.(string), nil
}

func (cli *RestCli) GetTxOutVal(txid string, idx uint32) (int64, error) {
	req, err := json.Marshal(Request{
		Jsonrpc: "1.0",
		Method:  "getrawtransaction",
		Params:  []interface{}{txid},
		Id:      1,
	})
	if err != nil {
		return -1, fmt.Errorf("[GetTxOutVal] failed to marshal request: %v", err)
	}

	resp, err := cli.sendPostReq(req)
	if err != nil {
		return -1, fmt.Errorf("[GetTxOutVal] failed to send post: %v", err)
	}
	if resp.Error != nil {
		return -1, fmt.Errorf("[GetTxOutVal] response shows failure: %v", resp.Error.Message)
	}
	rawTxStr, ok := resp.Result.(string)
	if !ok {
		return -1, fmt.Errorf("[GetTxOutVal] tx not found: %s", txid)
	}
	mtx := wire.NewMsgTx(wire.TxVersion)
	rawTx, err := hex.DecodeString(rawTxStr)
	if err != nil {
		return -1, fmt.Errorf("[GetTxOutVal] failed to decode raw tx string: %v", err)
	}
	mtx.BtcDecode(bytes.NewBuffer(rawTx), wire.ProtocolVersion, wire.LatestEncoding)

	return mtx.TxOut[idx].Value, nil
}

func (cli *RestCli) GetHeader(h int32) (*wire.BlockHeader, error) {
	req, err := json.Marshal(Request{
		Jsonrpc: "1.0",
		Method:  "getblockhash",
		Params:  []interface{}{h},
		Id:      1,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	resp, err := cli.sendPostReq(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send post: %v", err)
	}
	if resp.Error != nil {
		return nil, fmt.Errorf("response shows failure: %v", resp.Error.Message)
	}
	hash := resp.Result.(string)

	req, err = json.Marshal(Request{
		Jsonrpc: "1.0",
		Method:  "getblockheader",
		Params:  []interface{}{hash, false},
		Id:      1,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	resp, err = cli.sendPostReq(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send post: %v", err)
	}
	if resp.Error != nil {
		return nil, fmt.Errorf("response shows failure: %v", resp.Error.Message)
	}

	str, ok := resp.Result.(string)
	if !ok {
		return nil, errors.New("result is not string type")
	}
	hb, err := hex.DecodeString(str)
	if err != nil {
		return nil, fmt.Errorf("failed to decode string: %v", err)
	}
	header := &wire.BlockHeader{}
	if err := header.BtcDecode(bytes.NewBuffer(hb), wire.ProtocolVersion, wire.LatestEncoding); err != nil {
		return nil, fmt.Errorf("failed to decode header: %v", err)
	}

	return header, nil
}

func LenOfSyncMap(sm *sync.Map) int64 {
	lengh := int64(0)
	f := func(key, value interface{}) bool {
		lengh++
		return true
	}
	sm.Range(f)
	return lengh
}

func HexStringReverse(value string) string {
	aa, _ := hex.DecodeString(value)
	bb := hexReverse(aa)
	return hex.EncodeToString(bb)
}

func hexReverse(arr []byte) []byte {
	l := len(arr)
	x := make([]byte, 0)
	for i := l - 1; i >= 0; i-- {
		x = append(x, arr[i])
	}
	return x
}

func BuildData(toChainId uint64, ccFee int64, toAddr string) ([]byte, error) {
	var data []byte
	ccflag := byte(0xcc)
	var args *btc.Args
	switch toChainId {
	case config.DefConfig.EthChainID:
		toAddr = strings.Replace(toAddr, "0x", "", 1)
		toAddrBytes, err := hex.DecodeString(toAddr)
		if err != nil {
			return nil, err
		}
		args = &btc.Args{
			Address:   toAddrBytes[:],
			ToChainID: toChainId,
			Fee:       ccFee,
		}
	case config.DefConfig.OntChainID:
		addrBytes, _ := common.AddressFromBase58(toAddr)
		args = &btc.Args{
			Address:   addrBytes[:],
			ToChainID: toChainId,
			Fee:       ccFee,
		}
	case config.DefConfig.CMCrossChainId:
		addrBytes, _ := types2.AccAddressFromBech32(toAddr)
		args = &btc.Args{
			Address:   addrBytes[:],
			ToChainID: toChainId,
			Fee:       ccFee,
		}
	default:
		toAddrBytes, err := hex.DecodeString(toAddr)
		if err != nil {
			return nil, err
		}
		args = &btc.Args{
			Address:   toAddrBytes[:],
			ToChainID: toChainId,
			Fee:       ccFee,
		}
		log.Warn("not support address type, using hex.Decode")
	}
	var buf []byte
	sink := common2.NewZeroCopySink(buf)
	args.Serialization(sink)
	data = append(append(data, ccflag), sink.Bytes()...)

	return data, nil
}

func encryptBtcPrivk(path, privk, pwd string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}

	pri, err := keypair.GetP256KeyPairFromWIF([]byte(privk))
	privk = ""
	if err != nil {
		return err
	}

	wallet, err := account.Open(path)
	if err != nil {
		return err
	}

	pub := pri.Public()
	addr := types.AddressFromPubKey(pub)
	b58addr := addr.ToBase58()
	k, err := keypair.EncryptPrivateKey(pri, b58addr, []byte(pwd))
	if err != nil {
		return err
	}

	var accMeta account.AccountMetadata
	accMeta.Address = k.Address
	accMeta.KeyType = k.Alg
	accMeta.EncAlg = k.EncAlg
	accMeta.Hash = k.Hash
	accMeta.Key = k.Key
	accMeta.Curve = k.Param["curve"]
	accMeta.Salt = k.Salt
	accMeta.Label = ""
	accMeta.PubKey = hex.EncodeToString(keypair.SerializePublicKey(pub))
	accMeta.SigSch = signature.SHA256withECDSA.Name()

	err = wallet.ImportAccount(&accMeta)
	if err != nil {
		return err
	}

	return nil
}

func GetAccountByPassword(sdk *poly_go_sdk.PolySdk, path string, pwd []byte) (
	*poly_go_sdk.Account, error) {
	wallet, err := sdk.OpenWallet(path)
	if err != nil {
		return nil, fmt.Errorf("open wallet error: %v", err)
	}
	user, err := wallet.GetDefaultAccount(pwd)
	if err != nil {
		return nil, fmt.Errorf("getDefaultAccount error: %v", err)
	}
	return user, nil
}

func SetUpPoly(poly *poly_go_sdk.PolySdk, rpcAddr string) error {
	poly.NewRpcClient().SetAddress(rpcAddr)
	hdr, err := poly.GetHeaderByHeight(0)
	if err != nil {
		return err
	}
	poly.SetChainId(hdr.ChainID)
	return nil
}
