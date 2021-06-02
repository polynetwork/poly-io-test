package testframework

import (
	"context"
	"encoding/hex"
	"math/big"
	"os"
	"time"

	"github.com/polynetwork/poly-io-test/chains/kai"
	"github.com/polynetwork/poly-io-test/log"
)

func MonitorKai(ctx *TestFrameworkContext) {
	invoker := ctx.KaiInvoker
	ctx2 := context.Background()
	currentHeight, err := invoker.Client().BlockNumber(ctx2)
	if err != nil {
		log.Errorf("MonitorEthChain - ctx.EthTools.GetNodeHeight error: %s", err)
		os.Exit(1)
	}
	ethHeight := uint32(currentHeight) - 5
	updateTicker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-updateTicker.C:
			currentHeight, err := invoker.Client().BlockNumber(ctx2)
			if err != nil {
				log.Errorf("ctx.EthTools.GetNodeHeight error: %s", err)
				continue
			}
			if uint32(currentHeight) <= ethHeight+5 {
				continue
			}
			for uint32(currentHeight) > ethHeight+5 {
				ethHeight++
				err = parseKaiChainBlock(ctx, invoker, ethHeight) // TODO: influenced by forks
				if err != nil {
					log.Errorf("parseKaiChainBlock error: %s", err)
					ethHeight--
					break
				}
			}
		}
	}
}

func parseKaiChainBlock(ctx *TestFrameworkContext, invoker *kai.Invoker, height uint32) error {
	// contract is different
	lockevents, unlockevents, err := invoker.GetSmartContractEventByBlock(getEccm(invoker.ChainID), uint64(height))
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
		log.Infof("send cross chain tx on eth, tx hash: %s, tx id: %s chainID:%d", ethTxHash, ethTxIdStr, invoker.ChainID)
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
			log.Infof("receive cross chain tx on eth, txhash: %s, alliance tx hash: %s, raw tx hash: %s chainID:%d", event.Txid, allianceTxHash, rawTxHash, invoker.ChainID)
			ctx.Status.DelWithIndex(rawTxHash, idx)
		}
	}
	return nil
}
