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
package eth

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/polynetwork/eth-contracts/go_abi/eccm_abi"
	"github.com/polynetwork/poly-io-test/log"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ETHTools struct {
	restclient *RestClient
	ethclient  *ethclient.Client
}

type LockEvent struct {
	Method   string
	TxHash   string
	Txid     []byte
	Saddress string
	Tchain   uint32
	Taddress string
	Height   uint64
	Value    []byte
}
type UnlockEvent struct {
	Method   string
	Txid     string
	RTxid    string
	FromTxId string
	Height   uint64
	Token    string
}

type heightReq struct {
	JsonRpc string   `json:"jsonrpc"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	Id      uint     `json:"id"`
}

type heightRep struct {
	JsonRpc string `json:"jsonrpc"`
	Result  string `json:"result"`
	Id      uint   `json:"id"`
}

type BlockReq struct {
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      uint          `json:"id"`
}

type BlockRep struct {
	JsonRPC string        `json:"jsonrpc"`
	Result  *types.Header `json:"result"`
	Id      uint          `json:"id"`
}

func NewEthTools(url string) *ETHTools {
	ethclient, err := ethclient.Dial(url)
	if err != nil {
		log.Error("NewEthTools: cannot dial sync node, err: %s", err)
		return nil
	}
	restclient := NewRestClient()
	restclient.SetAddr(url)
	tool := &ETHTools{
		restclient: restclient,
		ethclient:  ethclient,
	}
	return tool
}

func (self *ETHTools) GetEthClient() *ethclient.Client {
	return self.ethclient
}

func (self *ETHTools) GetNodeHeight() (uint64, error) {
	req := &heightReq{
		JsonRpc: "2.0",
		Method:  "eth_blockNumber",
		Params:  make([]string, 0),
		Id:      1,
	}
	data, err := json.Marshal(req)
	if err != nil {
		return 0, fmt.Errorf("GetNodeHeight: marshal req err: %s", err)
	}
	resp, err := self.restclient.SendRestRequest(data)
	if err != nil {
		return 0, fmt.Errorf("GetNodeHeight err: %s", err)
	}
	rep := &heightRep{}
	err = json.Unmarshal(resp, rep)
	if err != nil {
		return 0, fmt.Errorf("GetNodeHeight, unmarshal resp err: %s", err)
	}
	height, err := strconv.ParseUint(rep.Result, 0, 64)
	if err != nil {
		return 0, fmt.Errorf("GetNodeHeight, parse resp height %s failed", rep.Result)
	} else {
		return height, nil
	}
}

func (self *ETHTools) GetBlockHeader(height uint64) (*types.Header, error) {
	params := []interface{}{fmt.Sprintf("0x%x", height), true}
	req := &BlockReq{
		JsonRpc: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  params,
		Id:      1,
	}
	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("GetNodeHeight: marshal req err: %s", err)
	}
	resp, err := self.restclient.SendRestRequest(data)
	if err != nil {
		return nil, fmt.Errorf("GetNodeHeight err: %s", err)
	}
	rsp := &BlockRep{}
	err = json.Unmarshal(resp, rsp)
	if err != nil {
		return nil, fmt.Errorf("GetNodeHeight, unmarshal resp err: %s", err)
	}

	return rsp.Result, nil
}

func (self *ETHTools) GetSmartContractEventByBlock(contractAddr string, height uint64) ([]*LockEvent, []*UnlockEvent, error) {
	eccmAddr := common.HexToAddress(contractAddr)
	instance, err := eccm_abi.NewEthCrossChainManager(eccmAddr, self.ethclient)
	if err != nil {
		return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, error: %s", err.Error())
	}

	opt := &bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: context.Background(),
	}

	ethlockevents := make([]*LockEvent, 0)
	{
		events, err := instance.FilterCrossChainEvent(opt, nil)
		if err != nil {
			return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, error :%s", err.Error())
		}

		if events == nil {
			return nil, nil, fmt.Errorf("GetSmartContractEventByBlock - no events found on FilterCrossChainEvent")
		}

		for events.Next() {
			evt := events.Event
			ethlockevents = append(ethlockevents, &LockEvent{
				Method:   "lock",
				TxHash:   evt.Raw.TxHash.String(),
				Txid:     evt.TxId,
				Saddress: evt.Sender.String(),
				Tchain:   uint32(evt.ToChainId),
				Value:    evt.Rawdata,
				Height:   height,
			})
		}
	}

	ethunlockevents := make([]*UnlockEvent, 0)
	{
		events, err := instance.FilterVerifyHeaderAndExecuteTxEvent(opt)
		if err != nil {
			return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, error :%s", err.Error())
		}

		if events == nil {
			return nil, nil, fmt.Errorf("GetSmartContractEventByBlock - no events found on FilterCrossChainEvent")
		}

		for events.Next() {
			evt := events.Event
			ethunlockevents = append(ethunlockevents, &UnlockEvent{
				Method:   "unlock",
				Txid:     evt.Raw.TxHash.String(),
				RTxid:    hex.EncodeToString(evt.CrossChainTxHash),
				FromTxId: hex.EncodeToString(evt.FromChainTxHash),
				Token:    hex.EncodeToString(evt.ToContract),
				Height:   height,
			})
		}
		if err != nil {
			return nil, nil, fmt.Errorf("GetSmartContractEventByBlock, error :%s", err.Error())
		}

		if events == nil {
			return nil, nil, fmt.Errorf("GetSmartContractEventByBlock - no events found on FilterCrossChainEvent")
		}

		for events.Next() {
			evt := events.Event
			ethunlockevents = append(ethunlockevents, &UnlockEvent{
				Method: "unlock",
				Txid:   evt.Raw.TxHash.String(),
				RTxid:  hex.EncodeToString(evt.CrossChainTxHash),
				Token:  hex.EncodeToString(evt.ToContract),
				Height: height,
			})
		}
	}
	return ethlockevents, ethunlockevents, nil
}

func EncodeBigInt(b *big.Int) string {
	if b.Uint64() == 0 {
		return "00"
	}
	return hex.EncodeToString(b.Bytes())
}

type getMoneyReq struct {
	Address string `json:"address"`
	Amount  uint   `json:"amount"`
}

type getMoneyRsp struct {
	Hash   string `json:"tx_hash"`
	Result string `json:"result"`
}

func (self *ETHTools) GetMoney(accounts []accounts.Account) []common.Hash {
	restclient := NewRestClient()
	restclient.SetAddr("https://api.bitaps.com/eth/testnet/v1/faucet/send/payment")
	hash := make([]common.Hash, 0)
	for _, account := range accounts {
		req := &getMoneyReq{
			Address: account.Address.String(),
			Amount:  1000000000000000000,
		}
		data, err := json.Marshal(req)
		if err != nil {
			log.Errorf("getMoney: marshal req err: %s", err)
			continue
		}
		resp, err := restclient.SendRestRequest([]byte(data))
		if err != nil {
			log.Errorf("getMoney err: %s", err)
			continue
		}
		rep := &getMoneyRsp{}
		err = json.Unmarshal(resp, rep)
		if err != nil {
			log.Errorf("getMoney, unmarshal resp err: %s", err)
			continue
		}
		hashBytes := common.HexToHash(rep.Hash)
		hash = append(hash, hashBytes)
	}
	return hash
}

func (self *ETHTools) WaitTransactionsConfirm(hashs []common.Hash) {
	hasPending := true
	for hasPending {
		time.Sleep(time.Second * 1)
		hasPending = false
		for _, hash := range hashs {
			_, ispending, err := self.ethclient.TransactionByHash(context.Background(), hash)
			log.Infof("transaction %s is pending: %d", hash.String(), ispending)
			if err != nil {
				hasPending = true
				continue
			}
			if ispending == true {
				hasPending = true
			} else {
			}
		}
	}
}

func (self *ETHTools) WaitTransactionConfirm(hash common.Hash) {
	for {
		time.Sleep(time.Millisecond * 100)
		_, ispending, err := self.ethclient.TransactionByHash(context.Background(), hash)
		if err != nil {
			log.Errorf("failed to call TransactionByHash: %v", err)
			continue
		}
		if ispending == true {
			continue
		} else {
			break
		}
	}

}

type RestClient struct {
	addr       string
	restClient *http.Client
	user       string
	passwd     string
}

func NewRestClient() *RestClient {
	return &RestClient{
		restClient: &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   5,
				DisableKeepAlives:     false,
				IdleConnTimeout:       time.Second * 300,
				ResponseHeaderTimeout: time.Second * 300,
				TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
			},
			Timeout: time.Second * 300,
		},
	}
}

func (self *RestClient) SetAddr(addr string) *RestClient {
	self.addr = addr
	return self
}

func (self *RestClient) SetAuth(user string, passwd string) *RestClient {
	self.user = user
	self.passwd = passwd
	return self
}

func (self *RestClient) SetRestClient(restClient *http.Client) *RestClient {
	self.restClient = restClient
	return self
}

func (self *RestClient) SendRestRequest(data []byte) ([]byte, error) {
	resp, err := self.restClient.Post(self.addr, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return nil, fmt.Errorf("http post request:%s error:%s", data, err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read rest response body error:%s", err)
	}
	return body, nil
}

func (self *RestClient) SendRestRequestWithAuth(data []byte) ([]byte, error) {
	url := self.addr
	bodyReader := bytes.NewReader(data)
	httpReq, err := http.NewRequest("POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("SendRestRequestWithAuth - build http request error:%s", err)
	}
	httpReq.Close = true
	httpReq.Header.Set("Content-Type", "application/json")

	httpReq.SetBasicAuth(self.user, self.passwd)

	rsp, err := self.restClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("SendRestRequestWithAuth - http post error:%s", err)
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil || len(body) == 0 {
		return nil, fmt.Errorf("SendRestRequestWithAuth - read rest response body error:%s", err)
	}
	return body, nil
}
