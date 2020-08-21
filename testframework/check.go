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
package testframework

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/btcsuite/btcd/wire"
	common2 "github.com/ethereum/go-ethereum/common"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/polynetwork/poly-io-test/chains/fisco/go_abi/eccm_abi"
	"github.com/polynetwork/poly-io-test/config"
	"github.com/polynetwork/poly-io-test/log"
	"github.com/polynetwork/poly-io-test/testframework/internal/github.com/hyperledger/fabric/protoutil"
	"github.com/polynetwork/poly/common"
	common3 "github.com/polynetwork/poly/native/service/cross_chain_manager/common"
	"io/ioutil"
	"math/big"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func MonitorOnt(ctx *TestFrameworkContext) {
	currentHeight, err := ctx.OntInvoker.OntSdk.GetCurrentBlockHeight()
	if err != nil {
		log.Errorf("ctx.OntSdk.GetCurrentBlockHeight error: %s", err)
		os.Exit(1)
	}
	ontHeight := currentHeight

	updateTicker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-updateTicker.C:
			currentHeight, err := ctx.OntInvoker.OntSdk.GetCurrentBlockHeight()
			if err != nil {
				log.Errorf("self.ontsdk.GetCurrentBlockHeight error: %s", err)
				continue
			}
			if currentHeight <= ontHeight {
				continue
			}
			for currentHeight > ontHeight {
				ontHeight++
				err = parseOntologyChainBlock(ctx, ontHeight)
				if err != nil {
					log.Errorf("parseOntologyChainBlock error: %s", err)
					ontHeight--
					break
				}
			}
		}
	}
}

func parseOntologyChainBlock(ctx *TestFrameworkContext, height uint32) error {
	events, err := ctx.OntInvoker.OntSdk.GetSmartContractEventByBlock(height)
	if err != nil {
		return err
	}

	for _, event := range events {
		for _, notify := range event.Notify {
			if notify.ContractAddress != "0900000000000000000000000000000000000000" {
				continue
			}

			states := notify.States.([]interface{})
			contractMethod, _ := states[0].(string)
			if contractMethod != "verifyToOntProof" {
				continue
			}
			// try to get all data
			//
			if len(states) < 5 {
				continue
			}
			allianceTxHash := states[1].(string)
			rawTxHash := states[2].(string)

			if ok, idx := ctx.Status.IsTxPending(rawTxHash); ok {
				log.Infof("receive cross chain tx on ontology, tx hash: %s, alliance tx hash: %s, "+
					"raw tx hash: %s", event.TxHash, allianceTxHash, rawTxHash)
				ctx.Status.DelWithIndex(rawTxHash, idx)
			}
		}
	}
	return nil
}

func MonitorEthChain(ctx *TestFrameworkContext) {
	currentHeight, err := ctx.EthInvoker.ETHUtil.GetNodeHeight()
	if err != nil {
		log.Errorf("ctx.EthTools.GetNodeHeight error: %s", err)
		os.Exit(1)
	}
	ethHeight := uint32(currentHeight) - 5
	updateTicker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-updateTicker.C:
			currentHeight, err := ctx.EthInvoker.ETHUtil.GetNodeHeight()
			if err != nil {
				log.Errorf("ctx.EthTools.GetNodeHeight error: %s", err)
				continue
			}
			if uint32(currentHeight) <= ethHeight+5 {
				continue
			}
			for uint32(currentHeight) > ethHeight+5 {
				ethHeight++
				err = parseEthChainBlock(ctx, ethHeight) // TODO: influenced by forks
				if err != nil {
					log.Errorf("parseEthChainBlock error: %s", err)
					ethHeight--
					break
				}
			}
		}
	}
}

func parseEthChainBlock(ctx *TestFrameworkContext, height uint32) error {
	// contract is different
	lockevents, unlockevents, err := ctx.EthInvoker.ETHUtil.GetSmartContractEventByBlock(config.DefConfig.Eccm, uint64(height))
	if err != nil {
		return err
	}
	//log.Infof("eth chain, block height: %d, unlock events num: %d, lock events num: %d", height, len(unlockevents), len(lockevents))
	for _, event := range lockevents {
		// try to get all data
		//
		ethTxHash := event.TxHash[2:]
		ok, idx := ctx.Status.IsTxPending(ethTxHash)
		if !ok {
			continue
		}

		var ethTxIdByte []byte
		indexInt := big.NewInt(0)
		indexInt.SetBytes(event.Txid)
		for i := len(indexInt.Bytes()); i < 32; i++ {
			ethTxIdByte = append(ethTxIdByte, 0)
		}
		ethTxIdByte = append(ethTxIdByte, indexInt.Bytes()...)
		ethTxIdStr := hex.EncodeToString(ethTxIdByte)
		log.Infof("send cross chain tx on eth, tx hash: %s, tx id: %s", ethTxHash, ethTxIdStr)
		caseStatus := ctx.Status.GetCaseStatus(idx)
		caseStatus.AddTx(ethTxIdStr, &TxInfo{ethTxHash, time.Now()})
		caseStatus.Del(ethTxHash)
	}

	for _, event := range unlockevents {
		// try to get all data
		//
		allianceTxHash := event.RTxid
		rawTxHash := event.FromTxId
		if ok, idx := ctx.Status.IsTxPending(rawTxHash); ok {
			log.Infof("receive cross chain tx on eth, txhash: %s, alliance tx hash: %s, raw tx hash: %s", event.Txid, allianceTxHash, rawTxHash)
			ctx.Status.DelWithIndex(rawTxHash, idx)
		}
	}
	return nil
}

func MonitorRChain(ctx *TestFrameworkContext) {
	currentHeight, err := ctx.RcSdk.GetCurrentBlockHeight()
	if err != nil {
		log.Errorf("ctx.RcSdk.GetCurrentBlockHeight error: %s", err)
		os.Exit(1)
	}
	rcHeight := currentHeight

	updateTicker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-updateTicker.C:
			currentHeight, err := ctx.RcSdk.GetCurrentBlockHeight()
			if err != nil {
				log.Errorf("self.RcSdk.GetCurrentBlockHeight error: %s", err)
				continue
			}
			if currentHeight <= rcHeight {
				continue
			}
			for currentHeight > rcHeight {
				rcHeight++
				err = parseRelayChainBlock(ctx, rcHeight)
				if err != nil {
					log.Errorf("parseRelayChainBlock error: %s", err)
					rcHeight--
					break
				}
			}
		}
	}
}

func parseRelayChainBlock(ctx *TestFrameworkContext, height uint32) error {
	events, err := ctx.RcSdk.GetSmartContractEventByBlock(height)
	if err != nil {
		return err
	}

	for _, event := range events {
		for _, notify := range event.Notify {
			states, ok := notify.States.([]interface{})
			if !ok {
				continue
			}
			name, ok := states[0].(string)
			if ok && name == "btcTxToRelay" {
				txHash, _ := states[4].(string)
				if ok, idx := ctx.Status.IsTxPending(txHash); ok {
					log.Infof("receive cross chain tx on relay chain, tx hash: %s, raw tx hash: %s", event.TxHash, txHash)
					raw, _ := hex.DecodeString(states[3].(string))
					mtx := wire.NewMsgTx(wire.TxVersion)
					_ = mtx.BtcDecode(bytes.NewBuffer(raw), wire.ProtocolVersion, wire.LatestEncoding)
					txid := mtx.TxHash()
					caseStatus := ctx.Status.GetCaseStatus(idx)
					caseStatus.AddTx(txid.String(), &TxInfo{"RCToBtc", time.Now()})
					caseStatus.Del(txHash)
				}
			} else if ok && name == "makeProof" && uint64(states[2].(float64)) == config.DefConfig.NeoChainID {
				txHash, _ := states[3].(string)
				if ok, idx := ctx.Status.IsTxPending(txHash); ok {
					pHash, _ := common.Uint256FromHexString(event.TxHash)
					log.Infof("receive cross chain tx on relay chain, tx hash: %s, raw tx hash: %s", event.TxHash, txHash)
					caseStatus := ctx.Status.GetCaseStatus(idx)
					caseStatus.AddTx(hex.EncodeToString(pHash[:]), &TxInfo{"PolyToNeo", time.Now()})
					caseStatus.Del(txHash)
				}
			}
		}
	}
	return nil
}

func MonitorBtc(ctx *TestFrameworkContext) {
	updateTicker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-updateTicker.C:
			for idx, v := range ctx.Status.caseMap {
				for a, b := range v.GetMapCopy() {
					if b.Ty != "RCToBtc" {
						continue
					}
					raw, err := ctx.BtcInvoker.BtcCli.GetRawTransaction(a)
					if raw != "" && err == nil {
						log.Infof("receive cross chain tx on btc chain, tx hash: %s, info: %s", a, b)
						ctx.Status.DelWithIndex(a, idx)
					}
				}
			}
		}
	}
}

func MonitorCosmos(ctx *TestFrameworkContext) {
	updateTicker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-updateTicker.C:
			cases := ctx.Status.GetCaseMap()
			for _, cs := range cases {
				keys := make([]string, 0)
				cmap := cs.GetMapCopy()
				for k := range cmap {
					res, err := ctx.CMInvoker.RpcCli.TxSearch(fmt.Sprintf(
						"verify_to_cosmos_proof.merkle_value:make_tx_param:txhash = '%s'", k),
						false, 1, 1, "asc") //TODO
					if err == nil && res.TotalCount == 1 {
						if res.Txs[0].TxResult.Code == 0 {
							keys = append(keys, k)
							log.Infof("MonitorCosmos, catch tx: %s", k)
						} else {
							log.Errorf("MonitorCosmos, catch tx %s and found it failed: %s", res.Txs[0].TxResult.Log)
						}
					}
					if err != nil {
						log.Warnf("query error: %v", err)
					}
				}
				cs.BatchDel(keys)
			}
		}
	}
}

func MonitorNeo(ctx *TestFrameworkContext) {
	res := ctx.NeoInvoker.Cli.GetBlockCount()
	if res.HasError() {
		log.Errorf("failed to get curr height: %s", res.Error.Message)
		os.Exit(1)
	}
	left := uint32(res.Result - 1)

	updateTicker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-updateTicker.C:
			res := ctx.NeoInvoker.Cli.GetBlockCount()
			if res.HasError() {
				continue
			}
			right := uint32(res.Result) - 3
			if left >= right {
				continue
			}
			for i := uint32(left); i < right; i++ {
				res := ctx.NeoInvoker.Cli.GetBlockByIndex(i)
				for _, tx := range res.Result.Tx {
					if tx.Type != "InvocationTransaction" {
						continue
					}
					appLogResp := ctx.NeoInvoker.Cli.GetApplicationLog(tx.Txid)
					if appLogResp.ErrorResponse.Error.Message != "" {
						continue
					}
					appLog := appLogResp.Result
					for _, exeitem := range appLog.Executions {
						for _, notify := range exeitem.Notifications {
							if notify.Contract != config.DefConfig.NeoCCMC {
								continue
							}
							if len(notify.State.Value) == 0 {
								continue
							}
							contractMethod, _ := hex.DecodeString(notify.State.Value[0].Value)
							if string(contractMethod) != "CrossChainUnlockEvent" {
								continue
							}
							if ok, idx := ctx.Status.IsTxPending(notify.State.Value[3].Value); ok {
								ctx.Status.DelWithIndex(notify.State.Value[3].Value, idx)
								log.Infof("neo unlock, txhash: %s, fromTxHash: %s", tx.Txid, notify.State.Value[3].Value)
							}
						}
					}
				}
			}
			left = right
		}
	}
}

func MonitorFisco(ctx *TestFrameworkContext) {
	left, err := ctx.FiscoInvoker.BlockNumber()
	if err != nil {
		panic(err)
	}
	left--

	updateTicker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-updateTicker.C:
			right, err := ctx.FiscoInvoker.BlockNumber()
			if err != nil {
				log.Errorf("MonitorFisco failed to BlockNumber: %v", err)
				continue
			}
			if right <= left {
				continue
			}
			for left < right {
				left++
				err := CheckFiscoHeight(ctx, uint64(left))
				if err != nil {
					log.Errorf("MonitorFisco error: %v", err)
				}
			}
		}
	}
}

type BlockRes struct {
	Transactions []string `json:"transactions"`
}

func CheckFiscoHeight(ctx *TestFrameworkContext, height uint64) error {
	eccmAddress := common2.HexToAddress(config.DefConfig.FiscoCCMC)
	eccmContract, err := eccm_abi.NewEthCrossChainManager(eccmAddress, ctx.FiscoInvoker.FiscoSdk)
	if err != nil {
		return err
	}
	blk, err := ctx.FiscoInvoker.FiscoSdk.GetBlockByNumber(context.Background(), strconv.FormatUint(height, 10), false)
	if err != nil {
		return fmt.Errorf("CheckFiscoHeight - GetBlockByNumber error :%s", err.Error())
	}
	res := &BlockRes{}
	err = json.Unmarshal(blk, res)
	if err != nil {
		return fmt.Errorf("CheckFiscoHeight - Unmarshal error :%s", err.Error())
	}
	for _, tx := range res.Transactions {
		recp, err := ctx.FiscoInvoker.FiscoSdk.TransactionReceipt(context.Background(), common2.HexToHash(tx))
		if err != nil {
			log.Errorf("CheckFiscoHeight - tx %s TransactionReceipt error: %v", tx, err.Error())
			continue
		}
		if recp.Status != "0x0" {
			continue
		}
		for _, v := range recp.Logs {
			if v.Address != strings.ToLower(config.DefConfig.FiscoCCMC) {
				continue
			}
			topics := make([]common2.Hash, len(v.Topics))
			for i, t := range v.Topics {
				topics[i] = common2.HexToHash(t.(string))
			}
			rawData, _ := hex.DecodeString(strings.TrimPrefix(v.Data, "0x"))
			evt, _ := eccmContract.ParseCrossChainEvent(types.Log{
				Address: common2.HexToAddress(v.Address),
				Topics:  topics,
				Data:    rawData,
			})
			if evt != nil {
				ok, idx := ctx.Status.IsTxPending(strings.TrimPrefix(tx, "0x"))
				if !ok {
					continue
				}

				var ethTxIdByte []byte
				indexInt := big.NewInt(0)
				indexInt.SetBytes(evt.TxId)
				for i := len(indexInt.Bytes()); i < 32; i++ {
					ethTxIdByte = append(ethTxIdByte, 0)
				}
				ethTxIdByte = append(ethTxIdByte, indexInt.Bytes()...)
				ethTxIdStr := hex.EncodeToString(ethTxIdByte)
				log.Infof("send cross chain tx on Fisco, tx hash: %s, tx id: %s", tx, ethTxIdStr)
				caseStatus := ctx.Status.GetCaseStatus(idx)
				txUnPrefix := strings.TrimPrefix(tx, "0x")
				caseStatus.AddTx(ethTxIdStr, &TxInfo{txUnPrefix, time.Now()})
				caseStatus.Del(txUnPrefix)
				continue
			}
			evt1, _ := eccmContract.ParseVerifyHeaderAndExecuteTxEvent(types.Log{
				Address: common2.HexToAddress(v.Address),
				Topics:  topics,
				Data:    rawData,
			})
			if evt1 != nil {
				fromTx := hex.EncodeToString(evt1.FromChainTxHash)
				if ok, idx := ctx.Status.IsTxPending(fromTx); ok {
					ctx.Status.DelWithIndex(fromTx, idx)
					log.Infof("fisco VerifyHeaderAndExecuteTxEvent, fisco txhash: %s, fromTxHash: %s", tx, fromTx)
				}
			}
		}
	}

	return nil
}

func MonitorFabric(ctx *TestFrameworkContext) {
	info, err := ctx.FabricInvoker.LedgerCLi.QueryInfo()
	if err != nil {
		panic(err)
	}
	curr := info.BCI.Height - 1
	left := curr - 1

	updateTicker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-updateTicker.C:
			info, err := ctx.FabricInvoker.LedgerCLi.QueryInfo()
			if err != nil {
				panic(err)
			}
			curr = info.BCI.Height - 1
			if curr <= left {
				continue
			}

			for h := left + 1; h <= curr; h++ {
				blk, err := ctx.FabricInvoker.LedgerCLi.QueryBlock(h)
				if err != nil {
					log.Errorf("failed to get fabric block: %v", err)
					h--
					time.Sleep(time.Second)
					continue
				}
				for _, v := range blk.Data.Data {
					xx, _ := protoutil.GetEnvelopeFromBlock(v)
					cas, _ := protoutil.GetActionsFromEnvelopeMsg(xx)

					for _, v := range cas {
						chaincodeEvent := &peer.ChaincodeEvent{}
						_ = proto.Unmarshal(v.Events, chaincodeEvent)
						if len(chaincodeEvent.TxId) == 0 {
							continue
						}
						tx, _ := ctx.FabricInvoker.LedgerCLi.QueryTransaction(fab.TransactionID(chaincodeEvent.TxId))
						if tx.ValidationCode == 0 && strings.Contains(chaincodeEvent.EventName, "from_poly") {
							merkleValue := new(common3.ToMerkleValue)
							_ = merkleValue.Deserialization(common.NewZeroCopySource(chaincodeEvent.Payload))

							fromTx := hex.EncodeToString(merkleValue.MakeTxParam.TxHash)
							if ok, idx := ctx.Status.IsTxPending(fromTx); ok {
								ctx.Status.DelWithIndex(fromTx, idx)
								log.Infof("MonitorFabric, height: %d, txid: %s, validate code: %d, from_tx: %s",
									h, chaincodeEvent.TxId, tx.ValidationCode, fromTx)
							}
						}
					}
				}
			}

			left = curr
		}
	}
}

func ReportPending(ctx *TestFrameworkContext) {
	reportTicker := time.NewTicker(time.Second * time.Duration(config.DefConfig.ReportInterval))
	_ = os.RemoveAll(config.DefConfig.ReportDir)
	for {
		select {
		case <-reportTicker.C:
			info := ctx.Status.Info()
			for k, v := range ctx.Cases {
				content := info[k+1]
				if content != "no tx for now" && content != "success!" {
					log.Infof(content)
				}
				name := ctx.Framework.getTestCaseName(v)
				_ = os.Mkdir(config.DefConfig.ReportDir, os.ModePerm)
				err := ioutil.WriteFile(path.Join(config.DefConfig.ReportDir, name), []byte(content), os.ModePerm)
				if err != nil {
					log.Errorf("failed to write file: %v", err)
				}
			}
		}
	}
}
