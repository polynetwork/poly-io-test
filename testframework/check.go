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
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/wire"
	"github.com/polynetwork/poly-io-test/config"
	"github.com/polynetwork/poly-io-test/log"
	"github.com/polynetwork/poly/common"
	"io/ioutil"
	"math/big"
	"os"
	"path"
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
