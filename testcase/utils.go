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
package testcase

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/polynetwork/poly-go-sdk"
	"github.com/polynetwork/poly-io-test/chains/btc"
	"github.com/polynetwork/poly-io-test/chains/eth"
	"github.com/polynetwork/poly-io-test/config"
	"github.com/polynetwork/poly-io-test/log"
	"github.com/polynetwork/poly-io-test/testframework"
	"github.com/polynetwork/poly/common"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func WaitTransactionConfirm(ethclient *ethclient.Client, hash ethcommon.Hash) {
	//
	errNum := 0
	for errNum < 100 {
		time.Sleep(time.Second * 1)
		_, ispending, err := ethclient.TransactionByHash(context.Background(), hash)
		log.Infof("hash: %s, isPending: %t", hash.String(), ispending)
		if err != nil {
			errNum++
			continue
		}
		if ispending == true {
			continue
		} else {
			break
		}
	}
}

func WaitPolyTx(txhash common.Uint256, poly *poly_go_sdk.PolySdk) {
	tick := time.NewTicker(100 * time.Millisecond)
	var h uint32
	startTime := time.Now()
	for range tick.C {
		h, _ = poly.GetBlockHeightByTxHash(txhash.ToHexString())
		curr, _ := poly.GetCurrentBlockHeight()
		if h > 0 && curr > h {
			break
		}

		if startTime.Add(100 * time.Millisecond); startTime.Second() > 300 {
			panic(fmt.Errorf("tx( %s ) is not confirm for a long time ( over %d sec )",
				txhash.ToHexString(), 300))
		}
	}
}

func GetRandAmount(high, low uint64) uint64 {
	return uint64(rand.Int63n(int64(high-low))) + low
}

func WaitUntilClean(status *testframework.CaseStatus) {
	tick := time.NewTicker(time.Second)
	for range tick.C {
		if status.Len() == 0 {
			break
		}
	}
}

func MakeEthAuth(signer *eth.EthSigner, nonce, gasPrice, gasLimit uint64) *bind.TransactOpts {
	auth := bind.NewKeyedTransactor(signer.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(int64(0)) // in wei
	auth.GasLimit = gasLimit          // in units
	auth.GasPrice = big.NewInt(int64(gasPrice))

	return auth
}

func MakeEthAuthWithValue(signer *eth.EthSigner, nonce, gasPrice, gasLimit, value uint64) *bind.TransactOpts {
	auth := MakeEthAuth(signer, nonce, gasPrice, gasLimit)
	auth.Value = big.NewInt(int64(value))
	return auth
}

func sendBtcCross(ctx *testframework.TestFrameworkContext, chainID uint64, btcSigner *btc.BtcSigner,
	targetAddress string, amount int64) (string, error) {
	value := float64(amount) / btcutil.SatoshiPerBitcoin
	fee := float64(config.DefConfig.BtcFee) / btcutil.SatoshiPerBitcoin

	addrPubk, err := btcutil.NewAddressPubKey(btcSigner.WIF.PrivKey.PubKey().SerializeCompressed(), config.BtcNet)
	if err != nil {
		return "", fmt.Errorf("sendBtcCross, Failed to new an address pubkey: %v", err)
	}
	pubkScript, err := txscript.PayToAddrScript(addrPubk.AddressPubKeyHash())
	if err != nil {
		return "", fmt.Errorf("sendBtcCross, Failed to build pubk script: %v", err)
	}
	data, err := btc.BuildData(chainID, 0, targetAddress)
	if err != nil {
		return "", fmt.Errorf("sendBtcCross, Failed to ge data: %v", err)
	}

	addr := addrPubk.EncodeAddress()
	//err = ra.Cli.ImportAddress(addr)
	//if err != nil {
	//	log.Errorf("rpc failed: %v", err)
	//	return nil
	//}
HERE:
	cnt, err := ctx.BtcInvoker.BtcCli.GetBlockCount()
	if err != nil {
		return "", fmt.Errorf("sendBtcCross, rpc failed: %v", err)
	}
	utxos, err := ctx.BtcInvoker.BtcCli.ListUnspent(1, cnt, addr)
	if err != nil {
		return "", fmt.Errorf("sendBtcCross, rpc failed: %v", err)
	}
	total, err := btcutil.NewAmount(value + fee)
	if err != nil {
		return "", fmt.Errorf("sendBtcCross, failed to new amount: %v", err)
	}
	selected, sumVal, err := btc.SelectUtxos(utxos, int64(total))
	if err != nil {
		return "", fmt.Errorf("sendBtcCross, failed to select utxo when build btc tx: %v", err)
	}

	//var prevPkScripts [][]byte
	var ipts []btcjson.TransactionInput
	for _, v := range selected {
		ipts = append(ipts, btcjson.TransactionInput{
			Txid: v.Txid,
			Vout: v.Vout,
		})
	}

	b, err := btc.NewBuilder(&btc.BuildCrossChainTxParam{
		Redeem:       config.DefConfig.BtcRedeem,
		Data:         data,
		Inputs:       ipts,
		NetParam:     config.BtcNet,
		PrevPkScript: pubkScript,
		Privk:        btcSigner.WIF.PrivKey,
		Locktime:     nil,
		ToMultiValue: value,
		Changes: func() map[string]float64 {
			if changeVal := float64(sumVal)/btcutil.SatoshiPerBitcoin - value - fee; changeVal > 0 {
				return map[string]float64{addrPubk.EncodeAddress(): changeVal}
			} else {
				return map[string]float64{}
			}
		}(),
	})
	if err != nil {
		return "", fmt.Errorf("sendBtcCross, Failed to new an instance of Builder: %v", err)
	}

	var buf bytes.Buffer
	err = b.BuildSignedTx()
	if err != nil || !b.IsSigned {
		return "", fmt.Errorf("sendBtcCross, Failed to build signed transaction: %v", err)
	}
	err = b.Tx.BtcEncode(&buf, wire.ProtocolVersion, wire.LatestEncoding)
	if err != nil {
		return "", fmt.Errorf("sendBtcCross, Failed to encode transaction: %v", err)
	}

	txid, err := ctx.BtcInvoker.BtcCli.SendRawTx(hex.EncodeToString(buf.Bytes()))
	if err != nil {
		if strings.Contains(err.Error(), "min relay fee not met") {
			arr := strings.Split(err.Error(), " ")
			newFee, _ := strconv.ParseUint(arr[len(arr)-1], 10, 64)
			fee = float64(newFee) / btcutil.SatoshiPerBitcoin
			goto HERE
		}
		return "", fmt.Errorf("sendBtcCross, failed to send tx: %v", err)
	}
	txidBytes, err := hex.DecodeString(txid)
	if err != nil {
		return "", fmt.Errorf("sendBtcCross, hex.DecodeString error: %v", err)
	}
	log.Infof("sendBtcCross, send tx %s(%f btc) to regression net(%s)", txid, value, ctx.BtcInvoker.BtcCli.Addr)
	return hex.EncodeToString(ToArrayReverse(txidBytes)), nil
}
