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
	types2 "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/joeqian10/neo-gogogo/helper"
	"github.com/joeqian10/neo-gogogo/sc"
	"github.com/joeqian10/neo-gogogo/tx"
	"github.com/ontio/ontology-go-sdk"
	ontcommon "github.com/ontio/ontology/common"
	nutils "github.com/ontio/ontology/smartcontract/service/native/utils"
	"github.com/polynetwork/eth-contracts/go_abi/btcx_abi"
	"github.com/polynetwork/eth-contracts/go_abi/erc20_abi"
	"github.com/polynetwork/eth-contracts/go_abi/lock_proxy_abi"
	"github.com/polynetwork/eth-contracts/go_abi/oep4_abi"
	"github.com/polynetwork/eth-contracts/go_abi/ontx_abi"
	"github.com/polynetwork/eth-contracts/go_abi/usdt_abi"
	"github.com/polynetwork/poly-io-test/chains/eth"
	"github.com/polynetwork/poly-io-test/chains/neo"
	"github.com/polynetwork/poly-io-test/chains/ont"
	"github.com/polynetwork/poly-io-test/config"
	"github.com/polynetwork/poly-io-test/log"
	"github.com/polynetwork/poly-io-test/testframework"
	"math"
	"math/big"
	"strings"
	"time"
)

func GetAccountByPath(path string) (*ontology_go_sdk.Account, error) {
	wallet, err := ontology_go_sdk.OpenWallet(path)
	if err != nil {
		return nil, fmt.Errorf("ontology_go_sdk.OpenWallet %s error:%s", path, err)
	}
	account, err := wallet.GetDefaultAccount([]byte(config.DefConfig.OntWalletPassword))
	if err != nil {
		return nil, fmt.Errorf("wallet.GetDefaultAccount error:%s", err)
	}
	return account, nil
}

func SendOntCrossEth(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	proxyContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntLockProxy)
	if err != nil {
		return fmt.Errorf("SendOntCrossEth, ontcommon.AddressFromHexString error: %s", err)
	}
	to := ctx.EthInvoker.EthTestSigner.Address.Bytes()
	txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{nutils.OntContractAddress[:], ctx.OntInvoker.OntAcc.Address[:],
			config.DefConfig.EthChainID, to, amount}})
	if err != nil {
		return fmt.Errorf("SendOntCrossEth, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	status.AddTx(hex.EncodeToString(txHash[:]), &testframework.TxInfo{"OntToEth", time.Now()})
	return nil
}

func SendEOntCrossOnt(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, onte string, amount uint64) error {
	gasPrice, err := ctx.EthInvoker.ETHUtil.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("SendEOntCrossOnt, get suggest gas price failed error: %s", err.Error())
	}
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(5))

	contractabi, err := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
	if err != nil {
		return fmt.Errorf("SendEOntCrossOnt, abi.JSON error: %v", err)
	}

	assetaddress := ethcommon.HexToAddress(onte)
	contractAddr := ethcommon.HexToAddress(config.DefConfig.EthLockProxy)

	ontxContract, err := ontx_abi.NewONTX(assetaddress, ctx.EthInvoker.ETHUtil.GetEthClient())
	if err != nil {
		return fmt.Errorf("SendEOntCrossOnt, NewONTX error: %v", err)
	}
	nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	auth := MakeEthAuth(ctx.EthInvoker.EthTestSigner, nonce, gasPrice.Uint64(), uint64(eth.DefaultGasLimit))
	if err != nil {
		return fmt.Errorf("SendEOntCrossOnt, failed to get eth auth: %v", err)
	}
	// 1
	txhash, err := ontxContract.Approve(auth, contractAddr, big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendEOntCrossOnt, failed to approve: %v", err)
	}
	ctx.EthInvoker.ETHUtil.WaitTransactionConfirm(txhash.Hash())
	//2
	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.DefConfig.OntChainID), ctx.OntInvoker.OntAcc.Address[:],
		big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendEOntCrossOnt, contractabi.Pack error: %v", err)
	}

	callMsg := ethereum.CallMsg{
		From: ctx.EthInvoker.EthTestSigner.Address, To: &contractAddr, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(0), Data: txData,
	}
	gasLimit, err := ctx.EthInvoker.ETHUtil.GetEthClient().EstimateGas(context.Background(), callMsg)
	if err != nil {
		return fmt.Errorf("SendEOntCrossOnt, estimate gas limit error: %s", err.Error())
	}

	nonce = ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(0)), gasLimit, gasPrice, txData)
	bf := new(bytes.Buffer)
	rlp.Encode(bf, tx)

	rawtx := hexutil.Encode(bf.Bytes())
	unsignedTx, err := eth.DeserializeTx(rawtx)
	if err != nil {
		return fmt.Errorf("SendEOntCrossOnt, eth.DeserializeTx error: %s", err.Error())
	}
	signedtx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, ctx.EthInvoker.EthTestSigner.PrivateKey)
	if err != nil {
		return fmt.Errorf("SendEOntCrossOnt, types.SignTx error: %s", err.Error())
	}

	err = ctx.EthInvoker.ETHUtil.GetEthClient().SendTransaction(context.Background(), signedtx)
	if err != nil {
		return fmt.Errorf("SendEOntCrossOnt, send transaction error:%s", err.Error())
	}

	status.AddTx(signedtx.Hash().String()[2:], &testframework.TxInfo{"OnteToOnt", time.Now()})
	WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), signedtx.Hash())
	return nil
}

func SendOngCrossEth(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	proxyContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntLockProxy)
	if err != nil {
		return fmt.Errorf("SendOngCrossEth, ontcommon.AddressFromHexString error: %s", err)
	}
	to := ctx.EthInvoker.EthTestSigner.Address.Bytes()
	txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{nutils.OngContractAddress[:], ctx.OntInvoker.OntAcc.Address[:],
			config.DefConfig.EthChainID, to, amount}})
	if err != nil {
		return fmt.Errorf("SendOngCrossEth, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	status.AddTx(hex.EncodeToString(txHash[:]), &testframework.TxInfo{"OngToEth", time.Now()})
	log.Infof("SendOngCrossEth, tx success, txHash is: %s", txHash.ToHexString())
	return nil
}

func SendOngeCrossOnt(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, onge string, amount uint64) error {
	gasPrice, err := ctx.EthInvoker.ETHUtil.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("SendOngeCrossOnt, get suggest gas price failed error: %s", err.Error())
	}
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(5))

	contractabi, err := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
	if err != nil {
		return fmt.Errorf("SendOngeCrossOnt, abi.JSON error:" + err.Error())
	}
	assetaddress := ethcommon.HexToAddress(onge)
	contractAddr := ethcommon.HexToAddress(config.DefConfig.EthLockProxy)

	// approve
	token, err := eth.NewStandardToken(assetaddress, ctx.EthInvoker.ETHUtil.GetEthClient())
	if err != nil {
		return fmt.Errorf("SendOngeCrossOnt, NewStandardToken error: %v", err.Error())
	}
	val, err := token.Allowance(nil, ctx.EthInvoker.EthTestSigner.Address, contractAddr)
	if err != nil {
		return fmt.Errorf("SendOngeCrossOnt, failed to get allowance: %v", err)
	}
	if val.Uint64() < amount {
		nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
		auth := MakeEthAuth(ctx.EthInvoker.EthTestSigner, nonce, gasPrice.Uint64(), uint64(eth.DefaultGasLimit))
		if err != nil {
			return fmt.Errorf("SendOngeCrossOnt, failed to get eth auth: %v", err)
		}
		tx, err := token.Approve(auth, contractAddr, big.NewInt(math.MaxInt64))
		if err != nil {
			return fmt.Errorf("SendOngeCrossOnt, failed to approve: %v", err)
		}
		WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), tx.Hash())
	}

	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.DefConfig.OntChainID), ctx.OntInvoker.OntAcc.Address[:],
		big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendOngeCrossOnt, contractabi.Pack error:" + err.Error())
	}

	callMsg := ethereum.CallMsg{
		From: ctx.EthInvoker.EthTestSigner.Address, To: &contractAddr, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(0), Data: txData,
	}
	gasLimit, err := ctx.EthInvoker.ETHUtil.GetEthClient().EstimateGas(context.Background(), callMsg)
	if err != nil {
		return fmt.Errorf("SendOngeCrossOnt, estimate gas limit error: %s", err.Error())
	}

	nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(0)), gasLimit, gasPrice, txData)
	bf := new(bytes.Buffer)
	rlp.Encode(bf, tx)

	rawtx := hexutil.Encode(bf.Bytes())
	unsignedTx, err := eth.DeserializeTx(rawtx)
	if err != nil {
		return fmt.Errorf("SendOngeCrossOnt, eth.DeserializeTx error: %s", err.Error())
	}
	signedtx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, ctx.EthInvoker.EthTestSigner.PrivateKey)
	if err != nil {
		return fmt.Errorf("SendOngeCrossOnt, types.SignTx error: %s", err.Error())
	}

	err = ctx.EthInvoker.ETHUtil.GetEthClient().SendTransaction(context.Background(), signedtx)
	if err != nil {
		return fmt.Errorf("SendOngeCrossOnt, send transaction error:%s", err.Error())
	}
	status.AddTx(signedtx.Hash().String()[2:], &testframework.TxInfo{"OngeToOnt", time.Now()})
	WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), signedtx.Hash())
	return nil
}

func SendOEP4CrossEth(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, contractAddress string, amount uint64) error {
	proxyContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntLockProxy)
	if err != nil {
		return fmt.Errorf("SendOEP4CrossEth, ontcommon.AddressFromHexString error: %s", err)
	}
	assetContractAddress, err := ontcommon.AddressFromHexString(contractAddress) // TODO: OEP4 contract on eth
	if err != nil {
		return fmt.Errorf("SendOEP4CrossEth, ontcommon.AddressFromHexString error: %s", err)
	}

	txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		assetContractAddress,
		[]interface{}{"approve", []interface{}{ctx.OntInvoker.OntAcc.Address, proxyContractAddress, amount}})
	if err != nil {
		return fmt.Errorf("SendOEP4CrossEth, approve error: %v", err)
	}
	ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)

	to := ctx.EthInvoker.EthTestSigner.Address.Bytes()
	txHash, err = ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{assetContractAddress[:], ctx.OntInvoker.OntAcc.Address[:],
			config.DefConfig.EthChainID, to, amount}})
	if err != nil {
		return fmt.Errorf("SendOEP4CrossEth, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	status.AddTx(hex.EncodeToString(txHash[:]), &testframework.TxInfo{"OEP4ToEth", time.Now()})
	log.Infof("SendOEP4CrossEth, tx success, txHash is: %s", txHash.ToHexString())
	ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)
	return nil
}

func SendEOEP4CrossOnt(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, oep4 string, amount uint64) error {
	gasPrice, err := ctx.EthInvoker.ETHUtil.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("SendEOEP4CrossOnt, get suggest gas price failed error: %s", err.Error())
	}
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(5))
	contractabi, err := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
	if err != nil {
		return fmt.Errorf("SendEOEP4CrossOnt, abi.JSON error:" + err.Error())
	}
	assetaddress := ethcommon.HexToAddress(oep4)
	contractAddr := ethcommon.HexToAddress(config.DefConfig.EthLockProxy)

	oep4Contract, err := oep4_abi.NewOEP4Template(assetaddress, ctx.EthInvoker.ETHUtil.GetEthClient())
	if err != nil {
		return fmt.Errorf("SendEOEP4CrossOnt, NewONTX error:" + err.Error())
	}
	nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	auth := MakeEthAuth(ctx.EthInvoker.EthTestSigner, nonce, gasPrice.Uint64(), uint64(eth.DefaultGasLimit))
	if err != nil {
		return fmt.Errorf("SendEOEP4CrossOnt, failed to get eth auth: %v", err)
	}
	txhash, err := oep4Contract.Approve(auth, contractAddr, big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendEOEP4CrossOnt, approve error: %v", err.Error())
	}
	ctx.EthInvoker.ETHUtil.WaitTransactionConfirm(txhash.Hash())

	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.DefConfig.OntChainID), ctx.OntInvoker.OntAcc.Address[:],
		big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendEOEP4CrossOnt, contractabi.Pack error:" + err.Error())
	}
	nonce = ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(0)), uint64(eth.DefaultGasLimit), gasPrice, txData)
	bf := new(bytes.Buffer)
	rlp.Encode(bf, tx)

	rawtx := hexutil.Encode(bf.Bytes())
	unsignedTx, err := eth.DeserializeTx(rawtx)
	if err != nil {
		return fmt.Errorf("SendEOEP4CrossOnt, eth.DeserializeTx error: %s", err.Error())
	}
	signedtx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, ctx.EthInvoker.EthTestSigner.PrivateKey)
	if err != nil {
		return fmt.Errorf("SendEOEP4CrossOnt, types.SignTx error: %s", err.Error())
	}

	err = ctx.EthInvoker.ETHUtil.GetEthClient().SendTransaction(context.Background(), signedtx)
	if err != nil {
		return fmt.Errorf("SendEOEP4CrossOnt, send transaction error:%s", err.Error())
	}
	status.AddTx(signedtx.Hash().String()[2:], &testframework.TxInfo{"OEP4eToOnt", time.Now()})
	WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), signedtx.Hash())
	return nil
}

func SendBtcoCrossBtc(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	btcxContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.BtcoContractAddress)
	if err != nil {
		return fmt.Errorf("SendBtcxCrossBtc, ontcommon.AddressFromHexString error: %s", err)
	}
	to := []byte(ctx.BtcInvoker.Signer.Address)
	txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		btcxContractAddress,
		[]interface{}{"lock", []interface{}{config.DefConfig.BtcChainID, ctx.OntInvoker.OntAcc.Address[:], to, amount}})
	if err != nil {
		return fmt.Errorf("SendBtcxCrossBtc, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	status.AddTx(hex.EncodeToString(txHash[:]), &testframework.TxInfo{"BtcoToBtc", time.Now()})
	log.Infof("SendBtcxCrossBtc, tx success, txHash is: %s", txHash.ToHexString())
	return nil
}

func SendBtcoCrossBtce(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	btcxContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.BtcoContractAddress)
	if err != nil {
		return fmt.Errorf("SendBtcoCrossBtce, ontcommon.AddressFromHexString error: %s", err)
	}
	txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		btcxContractAddress,
		[]interface{}{"lock", []interface{}{config.DefConfig.EthChainID, ctx.OntInvoker.OntAcc.Address[:],
			ctx.EthInvoker.EthTestSigner.Address.Bytes(), amount}})
	if err != nil {
		return fmt.Errorf("SendBtcoCrossBtce, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	status.AddTx(hex.EncodeToString(txHash[:]), &testframework.TxInfo{"BtcoToBtce", time.Now()})
	log.Infof("SendBtcoCrossBtce, tx success, txHash is: %s, val is: %d", txHash.ToHexString(), amount)
	return nil
}

func SendBtceCrossBtco(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	gasPrice, err := ctx.EthInvoker.ETHUtil.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("SendBtceCrossOnt, get suggest gas price failed error: %s", err.Error())
	}
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(5))

	contractabi, err := abi.JSON(strings.NewReader(btcx_abi.BTCXABI))
	if err != nil {
		return fmt.Errorf("SendBtceCrossOnt, abi.JSON error:" + err.Error())
	}
	assetaddress := ethcommon.HexToAddress(config.DefConfig.BtceContractAddress)
	txData, err := contractabi.Pack("lock", uint64(config.DefConfig.OntChainID), ctx.OntInvoker.OntAcc.Address[:], uint64(amount))
	if err != nil {
		return fmt.Errorf("SendBtceCrossOnt, contractabi.Pack error:" + err.Error())
	}

	callMsg := ethereum.CallMsg{
		From: ctx.EthInvoker.EthTestSigner.Address, To: &assetaddress, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(int64(0)), Data: txData,
	}
	gasLimit, err := ctx.EthInvoker.ETHUtil.GetEthClient().EstimateGas(context.Background(), callMsg)
	if err != nil {
		return fmt.Errorf("SendBtceCrossOnt, estimate gas limit error: %s", err.Error())
	}

	nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	tx := types.NewTransaction(nonce, assetaddress, big.NewInt(int64(0)), gasLimit, gasPrice, txData)
	bf := new(bytes.Buffer)
	rlp.Encode(bf, tx)

	rawtx := hexutil.Encode(bf.Bytes())
	unsignedTx, err := eth.DeserializeTx(rawtx)
	if err != nil {
		return fmt.Errorf("SendBtceCrossOnt, eth.DeserializeTx error: %s", err.Error())
	}
	signedtx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, ctx.EthInvoker.EthTestSigner.PrivateKey)
	if err != nil {
		return fmt.Errorf("SendBtceCrossOnt, types.SignTx error: %s", err.Error())
	}

	err = ctx.EthInvoker.ETHUtil.GetEthClient().SendTransaction(context.Background(), signedtx)
	if err != nil {
		return fmt.Errorf("SendBtceCrossOnt, send transaction error:%s", err.Error())
	}
	status.AddTx(signedtx.Hash().String()[2:], &testframework.TxInfo{"BtceToBtco", time.Now()})
	WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), signedtx.Hash())
	return nil
}

func SendBtcCrossOnt(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount int64) error {
	txid, err := sendBtcCross(ctx, config.DefConfig.OntChainID, ctx.BtcInvoker.Signer, ctx.OntInvoker.OntAcc.Address.ToBase58(), amount)
	if err != nil {
		return fmt.Errorf("SendBtcCrossOnt, sendBtcCross error: %s", err)
	}
	status.AddTx(txid, &testframework.TxInfo{"BtcToOnt", time.Now()})
	return nil
}

func SendBtcCrossEth(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount int64) error {
	txid, err := sendBtcCross(ctx, config.DefConfig.EthChainID, ctx.BtcInvoker.Signer, ctx.EthInvoker.EthTestSigner.Address.String(), amount)
	if err != nil {
		return fmt.Errorf("SendBtcCrossEth, sendBtcCross error: %s", err)
	}
	status.AddTx(txid, &testframework.TxInfo{"BtcToEth", time.Now()})
	return nil
}

func SendBtcCrossCosmos(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	txhash, err := sendBtcCross(ctx, config.DefConfig.CMCrossChainId, ctx.BtcInvoker.Signer,
		ctx.CMInvoker.Acc.Acc.String(), int64(amount))
	if err != nil {
		return err
	}
	status.AddTx(txhash, &testframework.TxInfo{"BtcToCosmos", time.Now()})
	return nil
}

func SendBtcFromCosmosToBitcoin(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amt uint64) error {
	lp, err := hex.DecodeString(config.DefConfig.CMLockProxy)
	if err != nil {
		return fmt.Errorf("SendBtcFromCosmosToBitcoin, failed to decode proxy: %v", err)
	}
	tx, err := ctx.CMInvoker.SendAsset(config.CM_BTCX, config.DefConfig.BtcChainID, int64(amt),
		[]byte(ctx.BtcInvoker.Signer.Address), lp)
	if err != nil {
		return fmt.Errorf("SendBtcFromCosmosToBitcoin, failed to send btc from cosmos to bitcoin: %v", err)
	}
	status.AddTx(strings.ToLower(tx.Hash.String()), &testframework.TxInfo{"BtcFromCosmosToBitcoin", time.Now()})
	return nil
}

func SendBtcFromCosmosToEthereum(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amt uint64) error {
	lp, err := hex.DecodeString(config.DefConfig.CMLockProxy)
	if err != nil {
		return fmt.Errorf("SendBtcFromCosmosToBitcoin, failed to decode proxy: %v", err)
	}
	tx, err := ctx.CMInvoker.SendAsset(config.CM_BTCX, config.DefConfig.EthChainID, int64(amt),
		ctx.EthInvoker.EthTestSigner.Address[:], lp)
	if err != nil {
		return fmt.Errorf("SendBtcFromCosmosToEthereum, failed to send btc from cosmos to ethereum: %v", err)
	}
	status.AddTx(strings.ToLower(tx.Hash.String()), &testframework.TxInfo{"BtcFromCosmosToEthereum", time.Now()})
	return nil
}

func SendBtcFromEthereumToCosmos(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amt uint64) error {
	gasPrice, err := ctx.EthInvoker.ETHUtil.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("SendBtcFromEthereumToCosmos, get suggest gas price failed error: %s", err.Error())
	}
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(5))

	contractabi, err := abi.JSON(strings.NewReader(btcx_abi.BTCXABI))
	if err != nil {
		return fmt.Errorf("SendBtcFromEthereumToCosmos, abi.JSON error:" + err.Error())
	}
	assetaddress := ethcommon.HexToAddress(config.DefConfig.BtceContractAddress)
	txData, err := contractabi.Pack("lock", uint64(config.DefConfig.CMCrossChainId), ctx.CMInvoker.Acc.Acc.Bytes(), amt)
	if err != nil {
		return fmt.Errorf("SendBtcFromEthereumToCosmos, contractabi.Pack error:" + err.Error())
	}

	callMsg := ethereum.CallMsg{
		From: ctx.EthInvoker.EthTestSigner.Address, To: &assetaddress, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(int64(0)), Data: txData,
	}
	gasLimit, err := ctx.EthInvoker.ETHUtil.GetEthClient().EstimateGas(context.Background(), callMsg)
	if err != nil {
		return fmt.Errorf("SendBtcFromEthereumToCosmos, estimate gas limit error: %s", err.Error())
	}

	nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	tx := types.NewTransaction(nonce, assetaddress, big.NewInt(int64(0)), gasLimit, gasPrice, txData)
	bf := new(bytes.Buffer)
	rlp.Encode(bf, tx)

	rawtx := hexutil.Encode(bf.Bytes())
	unsignedTx, err := eth.DeserializeTx(rawtx)
	if err != nil {
		return fmt.Errorf("SendBtcFromEthereumToCosmos, eth.DeserializeTx error: %s", err.Error())
	}
	signedtx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, ctx.EthInvoker.EthTestSigner.PrivateKey)
	if err != nil {
		return fmt.Errorf("SendBtcFromEthereumToCosmos, types.SignTx error: %s", err.Error())
	}

	err = ctx.EthInvoker.ETHUtil.GetEthClient().SendTransaction(context.Background(), signedtx)
	if err != nil {
		return fmt.Errorf("SendBtcFromEthereumToCosmos, send transaction error:%s", err.Error())
	}
	status.AddTx(signedtx.Hash().String()[2:], &testframework.TxInfo{"BtcFromEthereumToCosmos", time.Now()})
	WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), signedtx.Hash())
	return nil
}

func SendBtcFromCosmosToOntology(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amt uint64) error {
	lp, err := hex.DecodeString(config.DefConfig.CMLockProxy)
	if err != nil {
		return fmt.Errorf("SendBtcFromCosmosToBitcoin, failed to decode proxy: %v", err)
	}
	tx, err := ctx.CMInvoker.SendAsset(config.CM_BTCX, config.DefConfig.OntChainID, int64(amt), ctx.OntInvoker.OntAcc.Address[:], lp)
	if err != nil {
		return fmt.Errorf("SendBtcFromCosmosToOntology, failed to send btc from cosmos to ontology: %v", err)
	}
	status.AddTx(strings.ToLower(tx.Hash.String()), &testframework.TxInfo{"BtcFromCosmosToOntology", time.Now()})
	return nil
}

func SendBtcFromOntologyToCosmos(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amt uint64) error {
	btcxContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.BtcoContractAddress)
	if err != nil {
		return fmt.Errorf("SendBtcFromOntologyToCosmos, ontcommon.AddressFromHexString error: %s", err)
	}
	txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		btcxContractAddress,
		[]interface{}{"lock", []interface{}{config.DefConfig.CMCrossChainId, ctx.OntInvoker.OntAcc.Address[:],
			ctx.CMInvoker.Acc.Acc.Bytes(), amt}})
	if err != nil {
		return fmt.Errorf("SendBtcFromOntologyToCosmos, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	status.AddTx(hex.EncodeToString(txHash[:]), &testframework.TxInfo{"BtcFromOntologyToCosmos", time.Now()})
	log.Infof("SendBtcFromOntologyToCosmos, tx success, txHash is: %s, val is: %d", txHash.ToHexString(), amt)
	return nil
}

func SendEthCrossCosmos(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	gasPrice, err := ctx.EthInvoker.ETHUtil.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("SendEthToCosmos, get suggest gas price failed error: %s", err.Error())
	}
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(5))

	contractabi, err := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
	if err != nil {
		return fmt.Errorf("SendEthToCosmos, abi.JSON error:" + err.Error())
	}
	assetaddress := ethcommon.HexToAddress("0000000000000000000000000000000000000000")
	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.DefConfig.CMCrossChainId), ctx.CMInvoker.Acc.Acc.Bytes(),
		big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendEthToCosmos, contractabi.Pack error:" + err.Error())
	}

	contractAddr := ethcommon.HexToAddress(config.DefConfig.EthLockProxy)
	callMsg := ethereum.CallMsg{
		From: ctx.EthInvoker.EthTestSigner.Address, To: &contractAddr, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(int64(amount)), Data: txData,
	}
	gasLimit, err := ctx.EthInvoker.ETHUtil.GetEthClient().EstimateGas(context.Background(), callMsg)
	if err != nil {
		return fmt.Errorf("SendEthToCosmos, estimate gas limit error: %s", err.Error())
	}

	nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(amount)), gasLimit, gasPrice, txData)
	bf := new(bytes.Buffer)
	rlp.Encode(bf, tx)

	rawtx := hexutil.Encode(bf.Bytes())
	unsignedTx, err := eth.DeserializeTx(rawtx)
	if err != nil {
		return fmt.Errorf("SendEthToCosmos, eth.DeserializeTx error: %s", err.Error())
	}
	signedtx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, ctx.EthInvoker.EthTestSigner.PrivateKey)
	if err != nil {
		return fmt.Errorf("SendEthToCosmos, types.SignTx error: %s", err.Error())
	}

	err = ctx.EthInvoker.ETHUtil.GetEthClient().SendTransaction(context.Background(), signedtx)
	if err != nil {
		return fmt.Errorf("SendEthToCosmos, send transaction error:%s", err.Error())
	}
	status.AddTx(signedtx.Hash().String()[2:], &testframework.TxInfo{"EthToCosmos", time.Now()})
	WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), signedtx.Hash())
	return nil
}

func SendEthFromCosmosToEthereum(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amt uint64) error {
	lp, err := hex.DecodeString(config.DefConfig.CMLockProxy)
	if err != nil {
		return fmt.Errorf("SendBtcFromCosmosToBitcoin, failed to decode proxy: %v", err)
	}
	tx, err := ctx.CMInvoker.SendAsset(config.CM_ETHX, config.DefConfig.EthChainID, int64(amt),
		ctx.EthInvoker.EthTestSigner.Address[:], lp)
	if err != nil {
		return fmt.Errorf("SendEthFromCosmosToEthereum, failed to send eth: %v", err)
	}
	status.AddTx(strings.ToLower(tx.Hash.String()), &testframework.TxInfo{"EthFromCosmosToEthereum", time.Now()})
	return nil
}

func SendEthFromCosmosToOntology(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amt uint64) error {
	lp, err := hex.DecodeString(config.DefConfig.CMLockProxy)
	if err != nil {
		return fmt.Errorf("SendBtcFromCosmosToBitcoin, failed to decode proxy: %v", err)
	}
	tx, err := ctx.CMInvoker.SendAsset(config.CM_ETHX, config.DefConfig.OntChainID, int64(amt), ctx.OntInvoker.OntAcc.Address[:], lp)
	if err != nil {
		return fmt.Errorf("SendEthFromCosmosToOntology, failed to send eth: %v", err)
	}
	status.AddTx(strings.ToLower(tx.Hash.String()), &testframework.TxInfo{"EthFromCosmosToOntology", time.Now()})
	return nil
}

func SendEthFromOntologyToCosmos(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	proxyContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntLockProxy)
	if err != nil {
		return fmt.Errorf("SendEthFromOntologyToCosmos, ontcommon.AddressFromHexString error: %s", err)
	}

	assetaddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntEth)
	if err != nil {
		return fmt.Errorf("SendEthFromOntologyToCosmos, ontcommon.AddressFromHexString error: %s", err)
	}

	txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		assetaddress,
		[]interface{}{"approve", []interface{}{ctx.OntInvoker.OntAcc.Address, proxyContractAddress, amount}})
	if err != nil {
		return fmt.Errorf("SendEthFromOntologyToCosmos, approve error: %v", err)
	}
	ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)

	txHash, err = ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{assetaddress, ctx.OntInvoker.OntAcc.Address[:],
			config.DefConfig.CMCrossChainId, ctx.CMInvoker.Acc.Acc.Bytes(), amount}})
	if err != nil {
		return fmt.Errorf("SendEthFromOntologyToCosmos, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	status.AddTx(hex.EncodeToString(txHash[:]), &testframework.TxInfo{"EthFromOntologyToCosmos", time.Now()})
	ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)
	return nil
}

func SendErc20CrossCosmos(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, erc20ContractAddress string, amount uint64) error {
	gasPrice, err := ctx.EthInvoker.ETHUtil.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("SendErc20CrossCosmos, get suggest gas price failed error: %s", err.Error())
	}
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(5))

	contractabi, err := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
	if err != nil {
		return fmt.Errorf("SendErc20CrossCosmos, abi.JSON error:" + err.Error())
	}
	assetaddress := ethcommon.HexToAddress(erc20ContractAddress)
	contractAddr := ethcommon.HexToAddress(config.DefConfig.EthLockProxy)
	// approve
	erc20Contract, err := erc20_abi.NewERC20(assetaddress, ctx.EthInvoker.ETHUtil.GetEthClient())
	if err != nil {
		return fmt.Errorf("SendErc20CrossCosmos, NewERC20 error:" + err.Error())
	}
	nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	auth := MakeEthAuth(ctx.EthInvoker.EthTestSigner, nonce, gasPrice.Uint64(), uint64(eth.DefaultGasLimit))
	if err != nil {
		return fmt.Errorf("SendErc20CrossCosmos, failed to get eth auth: %v", err)
	}
	txhash, err := erc20Contract.Approve(auth, contractAddr, big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendErc20CrossCosmos, approve error: %v", err.Error())
	}
	ctx.EthInvoker.ETHUtil.WaitTransactionConfirm(txhash.Hash())

	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.DefConfig.CMCrossChainId), ctx.CMInvoker.Acc.Acc.Bytes(),
		big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendErc20CrossCosmos, contractabi.Pack error:" + err.Error())
	}
	callMsg := ethereum.CallMsg{
		From: ctx.EthInvoker.EthTestSigner.Address, To: &contractAddr, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(int64(0)), Data: txData,
	}
	gasLimit, err := ctx.EthInvoker.ETHUtil.GetEthClient().EstimateGas(context.Background(), callMsg)
	if err != nil {
		return fmt.Errorf("SendErc20CrossCosmos, estimate gas limit error: %s", err.Error())
	}

	nonce = ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(0)), gasLimit, gasPrice, txData)
	bf := new(bytes.Buffer)
	rlp.Encode(bf, tx)

	rawtx := hexutil.Encode(bf.Bytes())
	unsignedTx, err := eth.DeserializeTx(rawtx)
	if err != nil {
		return fmt.Errorf("SendErc20CrossCosmos, eth.DeserializeTx error: %s", err.Error())
	}
	signedtx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, ctx.EthInvoker.EthTestSigner.PrivateKey)
	if err != nil {
		return fmt.Errorf("SendErc20CrossCosmos, types.SignTx error: %s", err.Error())
	}

	err = ctx.EthInvoker.ETHUtil.GetEthClient().SendTransaction(context.Background(), signedtx)
	if err != nil {
		return fmt.Errorf("SendErc20CrossCosmos, send transaction error:%s", err.Error())
	}
	status.AddTx(signedtx.Hash().String()[2:], &testframework.TxInfo{"ERC20ToCosmos", time.Now()})
	WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), signedtx.Hash())
	return nil
}

func SendErc20FromCosmosToEthereum(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amt uint64) error {
	lp, err := hex.DecodeString(config.DefConfig.CMLockProxy)
	if err != nil {
		return fmt.Errorf("SendBtcFromCosmosToBitcoin, failed to decode proxy: %v", err)
	}
	tx, err := ctx.CMInvoker.SendAsset(config.CM_ERC20, config.DefConfig.EthChainID, int64(amt),
		ctx.EthInvoker.EthTestSigner.Address[:], lp)
	if err != nil {
		return fmt.Errorf("failed to send erc20: %v", err)
	}
	status.AddTx(strings.ToLower(tx.Hash.String()), &testframework.TxInfo{"Erc20FromCosmosToEthereum", time.Now()})
	return nil
}

func SendErc20FromCosmosToOntology(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amt uint64) error {
	lp, err := hex.DecodeString(config.DefConfig.CMLockProxy)
	if err != nil {
		return fmt.Errorf("SendBtcFromCosmosToBitcoin, failed to decode proxy: %v", err)
	}
	tx, err := ctx.CMInvoker.SendAsset(config.CM_ERC20, config.DefConfig.OntChainID, int64(amt), ctx.OntInvoker.OntAcc.Address[:], lp)
	if err != nil {
		return fmt.Errorf("SendErc20FromCosmosToOntology, failed to send erc20: %v", err)
	}
	status.AddTx(strings.ToLower(tx.Hash.String()), &testframework.TxInfo{"Erc20FromCosmosToOntology", time.Now()})
	return nil
}

func SendErc20FromOntologyToCosmos(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	proxyContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntLockProxy)
	if err != nil {
		return fmt.Errorf("SendErc20FromOntologyToCosmos, ontcommon.AddressFromHexString error: %s", err)
	}

	assetaddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntErc20)
	if err != nil {
		return err
	}
	txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		assetaddress,
		[]interface{}{"approve", []interface{}{ctx.OntInvoker.OntAcc.Address, proxyContractAddress, amount}})
	if err != nil {
		return fmt.Errorf("SendErc20FromOntologyToCosmos, approve error: %v", err)
	}
	ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)

	txHash, err = ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{assetaddress, ctx.OntInvoker.OntAcc.Address[:],
			config.DefConfig.CMCrossChainId, ctx.CMInvoker.Acc.Acc.Bytes(), amount}})
	if err != nil {
		return fmt.Errorf("SendErc20FromOntologyToCosmos, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	status.AddTx(hex.EncodeToString(txHash[:]), &testframework.TxInfo{"Erc20FromOntologyToCosmos", time.Now()})
	ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)
	return nil
}

func SendOntCrossCosmos(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	proxyContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntLockProxy)
	if err != nil {
		return fmt.Errorf("SendOntCrossCosmos, ontcommon.AddressFromHexString error: %s", err)
	}
	to := ctx.CMInvoker.Acc.Acc.Bytes()
	txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{nutils.OntContractAddress[:], ctx.OntInvoker.OntAcc.Address[:],
			config.DefConfig.CMCrossChainId, to, amount}})
	if err != nil {
		return fmt.Errorf("SendOntCrossCosmos, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	status.AddTx(hex.EncodeToString(txHash[:]), &testframework.TxInfo{"OntToCosmos", time.Now()})
	log.Infof("send ont to cosmos: ( amount: %d, txhash: %s )", amount, txHash.ToHexString())
	return nil
}

func SendOntFromCosmosToOntology(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amt uint64) error {
	lp, err := hex.DecodeString(config.DefConfig.CMLockProxy)
	if err != nil {
		return fmt.Errorf("SendBtcFromCosmosToBitcoin, failed to decode proxy: %v", err)
	}
	tx, err := ctx.CMInvoker.SendAsset(config.CM_ONT, config.DefConfig.OntChainID, int64(amt), ctx.OntInvoker.OntAcc.Address[:], lp)
	if err != nil {
		return fmt.Errorf("SendOntFromCosmosToOntology, failed to send ont: %v", err)
	}
	status.AddTx(strings.ToLower(tx.Hash.String()), &testframework.TxInfo{"OntFromCosmosToOntology", time.Now()})
	return nil
}

func SendOntFromCosmosToEthereum(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amt uint64) error {
	lp, err := hex.DecodeString(config.DefConfig.CMLockProxy)
	if err != nil {
		return fmt.Errorf("SendBtcFromCosmosToBitcoin, failed to decode proxy: %v", err)
	}
	tx, err := ctx.CMInvoker.SendAsset(config.CM_ONT, config.DefConfig.EthChainID, int64(amt),
		ctx.EthInvoker.EthTestSigner.Address[:], lp)
	if err != nil {
		return fmt.Errorf("SendOntFromCosmosToEthereum, failed to send ont: %v", err)
	}
	status.AddTx(strings.ToLower(tx.Hash.String()), &testframework.TxInfo{"OntFromCosmosToEthereum", time.Now()})
	return nil
}

func SendOntFromEthereumToCosmos(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	gasPrice, err := ctx.EthInvoker.ETHUtil.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("SendOntFromEthereumToCosmos, get suggest gas price failed error: %s", err.Error())
	}
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(5))

	contractabi, err := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
	if err != nil {
		return fmt.Errorf("SendOntFromEthereumToCosmos, abi.JSON error:" + err.Error())
	}

	assetaddress := ethcommon.HexToAddress(config.DefConfig.EthOntx)
	contractAddr := ethcommon.HexToAddress(config.DefConfig.EthLockProxy)

	ontxContract, err := ontx_abi.NewONTX(assetaddress, ctx.EthInvoker.ETHUtil.GetEthClient())
	if err != nil {
		return fmt.Errorf("SendOntFromEthereumToCosmos, NewONTX error:" + err.Error())
	}
	nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	auth := MakeEthAuth(ctx.EthInvoker.EthTestSigner, nonce, gasPrice.Uint64(), uint64(eth.DefaultGasLimit))
	if err != nil {
		return fmt.Errorf("SendOntFromEthereumToCosmos, failed to get eth auth: %v", err)
	}
	txhash, err := ontxContract.Approve(auth, contractAddr, big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendOntFromEthereumToCosmos, approve error: %v", err.Error())
	}
	ctx.EthInvoker.ETHUtil.WaitTransactionConfirm(txhash.Hash())
	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.DefConfig.CMCrossChainId),
		ctx.CMInvoker.Acc.Acc.Bytes(), big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendOntFromEthereumToCosmos, contractabi.Pack error:" + err.Error())
	}

	callMsg := ethereum.CallMsg{
		From: ctx.EthInvoker.EthTestSigner.Address, To: &contractAddr, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(0), Data: txData,
	}
	gasLimit, err := ctx.EthInvoker.ETHUtil.GetEthClient().EstimateGas(context.Background(), callMsg)
	if err != nil {
		return fmt.Errorf("SendOntFromEthereumToCosmos, estimate gas limit error: %s", err.Error())
	}

	nonce = ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(0)), gasLimit, gasPrice, txData)
	bf := new(bytes.Buffer)
	rlp.Encode(bf, tx)

	rawtx := hexutil.Encode(bf.Bytes())
	unsignedTx, err := eth.DeserializeTx(rawtx)
	if err != nil {
		return fmt.Errorf("SendOntFromEthereumToCosmos, eth.DeserializeTx error: %s", err.Error())
	}
	signedtx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, ctx.EthInvoker.EthTestSigner.PrivateKey)
	if err != nil {
		return fmt.Errorf("SendOntFromEthereumToCosmos, types.SignTx error: %s", err.Error())
	}

	err = ctx.EthInvoker.ETHUtil.GetEthClient().SendTransaction(context.Background(), signedtx)
	if err != nil {
		return fmt.Errorf("SendOntFromEthereumToCosmos, send transaction error:%s", err.Error())
	}

	status.AddTx(signedtx.Hash().String()[2:], &testframework.TxInfo{"OntFromEthereumToCosmos", time.Now()})
	WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), signedtx.Hash())
	return nil
}

func SendOngCrossCosmos(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	proxyContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntLockProxy)
	if err != nil {
		return fmt.Errorf("SendOngCrossCosmos, ontcommon.AddressFromHexString error: %s", err)
	}
	to := ctx.CMInvoker.Acc.Acc.Bytes()
	txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{nutils.OngContractAddress[:], ctx.OntInvoker.OntAcc.Address[:],
			config.DefConfig.CMCrossChainId, to, amount}})
	if err != nil {
		return fmt.Errorf("SendOngCrossCosmos, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	status.AddTx(hex.EncodeToString(txHash[:]), &testframework.TxInfo{"OngToCosmos", time.Now()})
	log.Infof("SendOngCrossCosmos, tx success, txHash is: %s", txHash.ToHexString())
	return nil
}

func SendOngFromCosmosToOntology(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amt uint64) error {
	lp, err := hex.DecodeString(config.DefConfig.CMLockProxy)
	if err != nil {
		return fmt.Errorf("SendBtcFromCosmosToBitcoin, failed to decode proxy: %v", err)
	}
	tx, err := ctx.CMInvoker.SendAsset(config.CM_ONG, config.DefConfig.OntChainID, int64(amt), ctx.OntInvoker.OntAcc.Address[:], lp)
	if err != nil {
		return fmt.Errorf("SendOngFromCosmosToOntology, failed to send ong: %v", err)
	}
	status.AddTx(strings.ToLower(tx.Hash.String()), &testframework.TxInfo{"OngFromCosmosToOntology", time.Now()})
	return nil
}

func SendOngFromCosmosToEthereum(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amt uint64) error {
	lp, err := hex.DecodeString(config.DefConfig.CMLockProxy)
	if err != nil {
		return fmt.Errorf("SendBtcFromCosmosToBitcoin, failed to decode proxy: %v", err)
	}
	tx, err := ctx.CMInvoker.SendAsset(config.CM_ONG, config.DefConfig.EthChainID, int64(amt),
		ctx.EthInvoker.EthTestSigner.Address[:], lp)
	if err != nil {
		return fmt.Errorf("SendOngFromCosmosToEthereum, failed to send ong: %v", err)
	}
	status.AddTx(strings.ToLower(tx.Hash.String()), &testframework.TxInfo{"OngFromCosmosToEthereum", time.Now()})
	return nil
}

func SendOngFromEthereumToCosmos(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	gasPrice, err := ctx.EthInvoker.ETHUtil.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("SendOngFromEthereumToCosmos, get suggest gas price failed error: %s", err.Error())
	}
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(5))

	contractabi, err := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
	if err != nil {
		return fmt.Errorf("SendOngFromEthereumToCosmos, abi.JSON error:" + err.Error())
	}

	assetaddress := ethcommon.HexToAddress(config.DefConfig.EthOngx)
	contractAddr := ethcommon.HexToAddress(config.DefConfig.EthLockProxy)

	ontxContract, err := ontx_abi.NewONTX(assetaddress, ctx.EthInvoker.ETHUtil.GetEthClient())
	if err != nil {
		return fmt.Errorf("SendOngFromEthereumToCosmos, NewONGX error:" + err.Error())
	}
	nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	auth := MakeEthAuth(ctx.EthInvoker.EthTestSigner, nonce, gasPrice.Uint64(), uint64(eth.DefaultGasLimit))
	if err != nil {
		return fmt.Errorf("SendOngFromEthereumToCosmos, failed to get eth auth: %v", err)
	}
	txhash, err := ontxContract.Approve(auth, contractAddr, big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendOngFromEthereumToCosmos, approve error: %v", err.Error())
	}
	ctx.EthInvoker.ETHUtil.WaitTransactionConfirm(txhash.Hash())

	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.DefConfig.CMCrossChainId),
		ctx.CMInvoker.Acc.Acc.Bytes(), big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendOngFromEthereumToCosmos, contractabi.Pack error:" + err.Error())
	}

	callMsg := ethereum.CallMsg{
		From: ctx.EthInvoker.EthTestSigner.Address, To: &contractAddr, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(0), Data: txData,
	}
	gasLimit, err := ctx.EthInvoker.ETHUtil.GetEthClient().EstimateGas(context.Background(), callMsg)
	if err != nil {
		return fmt.Errorf("SendOngFromEthereumToCosmos, estimate gas limit error: %s", err.Error())
	}

	nonce = ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(0)), gasLimit, gasPrice, txData)
	bf := new(bytes.Buffer)
	rlp.Encode(bf, tx)

	rawtx := hexutil.Encode(bf.Bytes())
	unsignedTx, err := eth.DeserializeTx(rawtx)
	if err != nil {
		return fmt.Errorf("SendOngFromEthereumToCosmos, eth.DeserializeTx error: %s", err.Error())
	}
	signedtx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, ctx.EthInvoker.EthTestSigner.PrivateKey)
	if err != nil {
		return fmt.Errorf("SendOngFromEthereumToCosmos, types.SignTx error: %s", err.Error())
	}

	err = ctx.EthInvoker.ETHUtil.GetEthClient().SendTransaction(context.Background(), signedtx)
	if err != nil {
		return fmt.Errorf("SendOngFromEthereumToCosmos, send transaction error:%s", err.Error())
	}

	status.AddTx(signedtx.Hash().String()[2:], &testframework.TxInfo{"OngFromEthereumToCosmos", time.Now()})
	WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), signedtx.Hash())
	return nil
}

func SendOep4CrossCosmos(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus,
	contractAddress string, amount uint64) error {
	proxyContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntLockProxy)
	if err != nil {
		return fmt.Errorf("SendOep4CrossCosmos, ontcommon.AddressFromHexString error: %s", err)
	}
	assetContractAddress, err := ontcommon.AddressFromHexString(contractAddress) // TODO: OEP4 contract on eth
	if err != nil {
		return fmt.Errorf("SendOep4CrossCosmos, ontcommon.AddressFromHexString error: %s", err)
	}
	txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		assetContractAddress,
		[]interface{}{"approve", []interface{}{ctx.OntInvoker.OntAcc.Address, proxyContractAddress, amount}})
	if err != nil {
		return fmt.Errorf("SendOep4CrossCosmos, approve error: %v", err)
	}
	ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)

	to := ctx.CMInvoker.Acc.Acc.Bytes()
	txHash, err = ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{assetContractAddress[:], ctx.OntInvoker.OntAcc.Address[:],
			config.DefConfig.CMCrossChainId, to, amount}})
	if err != nil {
		return fmt.Errorf("SendOep4CrossCosmos, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	status.AddTx(hex.EncodeToString(txHash[:]), &testframework.TxInfo{"OEP4ToCosmos", time.Now()})
	log.Infof("SendOep4CrossCosmos, tx success, txHash is: %s", txHash.ToHexString())
	ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)
	return nil
}

func SendOep4FromCosmosToOntology(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amt uint64) error {
	lp, err := hex.DecodeString(config.DefConfig.CMLockProxy)
	if err != nil {
		return fmt.Errorf("SendBtcFromCosmosToBitcoin, failed to decode proxy: %v", err)
	}
	tx, err := ctx.CMInvoker.SendAsset(config.CM_OEP4, config.DefConfig.OntChainID, int64(amt), ctx.OntInvoker.OntAcc.Address[:], lp)
	if err != nil {
		return fmt.Errorf("SendOep4FromCosmosToOntology, failed to send oep4: %v", err)
	}
	status.AddTx(strings.ToLower(tx.Hash.String()), &testframework.TxInfo{"Oep4FromCosmosToOntology", time.Now()})
	return nil
}

func SendOep4FromCosmosToEthereum(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus,
	ethAddress string, amt uint64) error {
	lp, err := hex.DecodeString(config.DefConfig.CMLockProxy)
	if err != nil {
		return fmt.Errorf("SendBtcFromCosmosToBitcoin, failed to decode proxy: %v", err)
	}
	addr := ethcommon.HexToAddress(ethAddress)
	tx, err := ctx.CMInvoker.SendAsset("oep4", config.DefConfig.EthChainID, int64(amt),
		addr[:], lp)
	if err != nil {
		return fmt.Errorf("SendOep4FromCosmosToEthereum, failed to send oep4: %v", err)
	}
	status.AddTx(strings.ToLower(tx.Hash.String()), &testframework.TxInfo{"Oep4FromCosmosToEthereum", time.Now()})
	return nil
}

func SendOep4FromEthereumToCosmos(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, signer *eth.EthSigner,
	cosmosAddr string, amount uint64) error {
	gasPrice, err := ctx.EthInvoker.ETHUtil.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("SendOep4FromEthereumToCosmos, get suggest gas price failed error: %s", err.Error())
	}
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(5))

	contractabi, err := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
	if err != nil {
		return fmt.Errorf("SendOep4FromEthereumToCosmos, abi.JSON error:" + err.Error())
	}

	assetaddress := ethcommon.HexToAddress(config.DefConfig.EthOep4)
	contractAddr := ethcommon.HexToAddress(config.DefConfig.EthLockProxy)

	ontxContract, err := ontx_abi.NewONTX(assetaddress, ctx.EthInvoker.ETHUtil.GetEthClient())
	if err != nil {
		return fmt.Errorf("SendOngFromEthereumToCosmos, NewONGX error:" + err.Error())
	}
	nonce := ctx.EthInvoker.NM.GetAddressNonce(signer.Address)
	auth := MakeEthAuth(signer, nonce, gasPrice.Uint64(), uint64(eth.DefaultGasLimit))
	if err != nil {
		return fmt.Errorf("SendOntFromEthereumToCosmos, failed to get eth auth: %v", err)
	}
	txhash, err := ontxContract.Approve(auth, contractAddr, big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendOntFromEthereumToCosmos, approve error: %v", err.Error())
	}
	ctx.EthInvoker.ETHUtil.WaitTransactionConfirm(txhash.Hash())

	toAddr, err := types2.AccAddressFromBech32(cosmosAddr)
	if err != nil {
		return fmt.Errorf("SendOep4FromEthereumToCosmos, failed to parse cosmos address %s: %v", cosmosAddr, err)
	}

	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.DefConfig.CMCrossChainId), toAddr.Bytes(),
		big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendOep4FromEthereumToCosmos, contractabi.Pack error:" + err.Error())
	}

	callMsg := ethereum.CallMsg{
		From: signer.Address, To: &contractAddr, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(0), Data: txData,
	}
	gasLimit, err := ctx.EthInvoker.ETHUtil.GetEthClient().EstimateGas(context.Background(), callMsg)
	if err != nil {
		return fmt.Errorf("SendOep4FromEthereumToCosmos, estimate gas limit error: %s", err.Error())
	}

	nonce = ctx.EthInvoker.NM.GetAddressNonce(signer.Address)
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(0)), gasLimit, gasPrice, txData)
	bf := new(bytes.Buffer)
	rlp.Encode(bf, tx)

	rawtx := hexutil.Encode(bf.Bytes())
	unsignedTx, err := eth.DeserializeTx(rawtx)
	if err != nil {
		return fmt.Errorf("SendOep4FromEthereumToCosmos, eth.DeserializeTx error: %s", err.Error())
	}
	signedtx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, signer.PrivateKey)
	if err != nil {
		return fmt.Errorf("SendOep4FromEthereumToCosmos, types.SignTx error: %s", err.Error())
	}

	err = ctx.EthInvoker.ETHUtil.GetEthClient().SendTransaction(context.Background(), signedtx)
	if err != nil {
		return fmt.Errorf("SendOep4FromEthereumToCosmos, send transaction error:%s", err.Error())
	}

	status.AddTx(signedtx.Hash().String()[2:], &testframework.TxInfo{"Oep4FromEthereumToCosmos", time.Now()})
	WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), signedtx.Hash())
	return nil
}

func SendEthCrossOnt(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	gasPrice, err := ctx.EthInvoker.ETHUtil.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("SendEthCrossOnt, get suggest gas price failed error: %s", err.Error())
	}
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(5))

	contractabi, err := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
	if err != nil {
		return fmt.Errorf("SendEthCrossOnt, abi.JSON error:" + err.Error())
	}

	assetaddress := ethcommon.HexToAddress("0000000000000000000000000000000000000000")
	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.DefConfig.OntChainID), ctx.OntInvoker.OntAcc.Address[:],
		big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendEthCrossOnt, contractabi.Pack error:" + err.Error())
	}

	contractAddr := ethcommon.HexToAddress(config.DefConfig.EthLockProxy)
	callMsg := ethereum.CallMsg{
		From: ctx.EthInvoker.EthTestSigner.Address, To: &contractAddr, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(int64(amount)), Data: txData,
	}
	gasLimit, err := ctx.EthInvoker.ETHUtil.GetEthClient().EstimateGas(context.Background(), callMsg)
	if err != nil {
		return fmt.Errorf("SendEthCrossOnt, estimate gas limit error: %s", err.Error())
	}

	nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(amount)), gasLimit, gasPrice, txData)
	bf := new(bytes.Buffer)
	rlp.Encode(bf, tx)

	rawtx := hexutil.Encode(bf.Bytes())
	unsignedTx, err := eth.DeserializeTx(rawtx)
	if err != nil {
		return fmt.Errorf("SendEthCrossOnt, eth.DeserializeTx error: %s", err.Error())
	}
	signedtx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, ctx.EthInvoker.EthTestSigner.PrivateKey)
	if err != nil {
		return fmt.Errorf("SendEthCrossOnt, types.SignTx error: %s", err.Error())
	}

	err = ctx.EthInvoker.ETHUtil.GetEthClient().SendTransaction(context.Background(), signedtx)
	if err != nil {
		return fmt.Errorf("SendEthCrossOnt, send transaction error:%s", err.Error())
	}
	status.AddTx(signedtx.Hash().String()[2:], &testframework.TxInfo{"EthToOnt", time.Now()})
	WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), signedtx.Hash())
	return nil
}

func SendEthoCrossEth(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, etho string, amount uint64) error {
	proxyContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntLockProxy)
	if err != nil {
		return fmt.Errorf("SendEthoCrossEth, ontcommon.AddressFromHexString error: %s", err)
	}

	assetaddress, err := ontcommon.AddressFromHexString(etho)
	if err != nil {
		return fmt.Errorf("SendEthoCrossEth, ontcommon.AddressFromHexString error: %s", err)
	}

	txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		assetaddress,
		[]interface{}{"approve", []interface{}{ctx.OntInvoker.OntAcc.Address, proxyContractAddress, amount}})
	if err != nil {
		return fmt.Errorf("SendEthoCrossEth, approve error: %v", err)
	}
	ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)

	to := ctx.EthInvoker.EthTestSigner.Address.String()
	txHash, err = ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{assetaddress, ctx.OntInvoker.OntAcc.Address[:],
			config.DefConfig.EthChainID, to, amount}})
	if err != nil {
		return fmt.Errorf("SendEthoCrossEth, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	status.AddTx(hex.EncodeToString(txHash[:]), &testframework.TxInfo{"EthoToEth", time.Now()})
	ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)
	return nil
}

func SendERC20CrossOnt(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, erc20ContractAddress string, amount uint64) error {
	gasPrice, err := ctx.EthInvoker.ETHUtil.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("SendERC20CrossOnt, get suggest gas price failed error: %s", err.Error())
	}
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(5))

	contractabi, err := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
	if err != nil {
		return fmt.Errorf("SendERC20CrossOnt, abi.JSON error:" + err.Error())
	}
	assetaddress := ethcommon.HexToAddress(erc20ContractAddress)
	contractAddr := ethcommon.HexToAddress(config.DefConfig.EthLockProxy)

	// approve
	erc20Contract, err := erc20_abi.NewERC20(assetaddress, ctx.EthInvoker.ETHUtil.GetEthClient())
	if err != nil {
		return fmt.Errorf("SendERC20CrossOnt, NewERC20 error:" + err.Error())
	}
	val, err := erc20Contract.BalanceOf(nil, ctx.EthInvoker.EthTestSigner.Address)
	if err != nil {
		return fmt.Errorf("SendERC20CrossOnt, failed to get the balance: %v", err)
	}
	if val.Uint64() < amount {
		return fmt.Errorf("SendERC20CrossOnt, balance %d is less than amount %d", val.Uint64(), amount)
	}
	nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	auth := MakeEthAuth(ctx.EthInvoker.EthTestSigner, nonce, gasPrice.Uint64(), uint64(eth.DefaultGasLimit))
	if err != nil {
		return fmt.Errorf("SendERC20CrossOnt, failed to get eth auth: %v", err)
	}
	txhash, err := erc20Contract.Approve(auth, contractAddr, big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendERC20CrossOnt, approve error: %v", err.Error())
	}
	ctx.EthInvoker.ETHUtil.WaitTransactionConfirm(txhash.Hash())

	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.DefConfig.OntChainID), ctx.OntInvoker.OntAcc.Address[:],
		big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendERC20CrossOnt, contractabi.Pack error:" + err.Error())
	}
	callMsg := ethereum.CallMsg{
		From: ctx.EthInvoker.EthTestSigner.Address, To: &contractAddr, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(int64(0)), Data: txData,
	}
	gasLimit, err := ctx.EthInvoker.ETHUtil.GetEthClient().EstimateGas(context.Background(), callMsg)
	if err != nil {
		return fmt.Errorf("SendERC20CrossOnt, estimate gas limit error: %s", err.Error())
	}

	nonce = ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(0)), gasLimit, gasPrice, txData)
	bf := new(bytes.Buffer)
	rlp.Encode(bf, tx)

	rawtx := hexutil.Encode(bf.Bytes())
	unsignedTx, err := eth.DeserializeTx(rawtx)
	if err != nil {
		return fmt.Errorf("SendERC20CrossOnt, eth.DeserializeTx error: %s", err.Error())
	}
	signedtx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, ctx.EthInvoker.EthTestSigner.PrivateKey)
	if err != nil {
		return fmt.Errorf("SendERC20CrossOnt, types.SignTx error: %s", err.Error())
	}

	err = ctx.EthInvoker.ETHUtil.GetEthClient().SendTransaction(context.Background(), signedtx)
	if err != nil {
		return fmt.Errorf("SendERC20CrossOnt, send transaction error:%s", err.Error())
	}
	status.AddTx(signedtx.Hash().String()[2:], &testframework.TxInfo{"ERC20ToOnt", time.Now()})
	WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), signedtx.Hash())
	return nil
}

func SendOERC20CrossEth(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, erc20 string, amount uint64) error {
	proxyContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntLockProxy)
	if err != nil {
		return fmt.Errorf("SendOERC20CrossEth, ontcommon.AddressFromHexString error: %s", err)
	}

	assetaddress, err := ontcommon.AddressFromHexString(erc20)
	if err != nil {
		return err
	}
	txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		assetaddress,
		[]interface{}{"approve", []interface{}{ctx.OntInvoker.OntAcc.Address, proxyContractAddress, amount}})
	if err != nil {
		return fmt.Errorf("SendEthoCrossEth, approve error: %v", err)
	}
	ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)

	to := ctx.EthInvoker.EthTestSigner.Address.Bytes()
	txHash, err = ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{assetaddress, ctx.OntInvoker.OntAcc.Address[:],
			config.DefConfig.EthChainID, to, amount}})
	if err != nil {
		return fmt.Errorf("SendOERC20CrossEth, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	status.AddTx(hex.EncodeToString(txHash[:]), &testframework.TxInfo{"OERC20ToEth", time.Now()})
	ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)
	return nil
}

func SendBtceCrossBtc(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	gasPrice, err := ctx.EthInvoker.ETHUtil.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("SendERC20CrossOnt, get suggest gas price failed error: %s", err.Error())
	}
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(5))

	contractabi, err := abi.JSON(strings.NewReader(btcx_abi.BTCXABI))
	if err != nil {
		return fmt.Errorf("SendERC20CrossOnt, abi.JSON error:" + err.Error())
	}
	assetaddress := ethcommon.HexToAddress(config.DefConfig.BtceContractAddress)
	txData, err := contractabi.Pack("lock", uint64(config.DefConfig.BtcChainID), []byte(ctx.BtcInvoker.Signer.Address), uint64(amount))
	if err != nil {
		return fmt.Errorf("SendERC20CrossOnt, contractabi.Pack error:" + err.Error())
	}

	callMsg := ethereum.CallMsg{
		From: ctx.EthInvoker.EthTestSigner.Address, To: &assetaddress, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(int64(0)), Data: txData,
	}
	gasLimit, err := ctx.EthInvoker.ETHUtil.GetEthClient().EstimateGas(context.Background(), callMsg)
	if err != nil {
		return fmt.Errorf("SendERC20CrossOnt, estimate gas limit error: %s", err.Error())
	}

	nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	tx := types.NewTransaction(nonce, assetaddress, big.NewInt(int64(0)), gasLimit, gasPrice, txData)
	bf := new(bytes.Buffer)
	rlp.Encode(bf, tx)

	rawtx := hexutil.Encode(bf.Bytes())
	unsignedTx, err := eth.DeserializeTx(rawtx)
	if err != nil {
		return fmt.Errorf("SendERC20CrossOnt, eth.DeserializeTx error: %s", err.Error())
	}
	signedtx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, ctx.EthInvoker.EthTestSigner.PrivateKey)
	if err != nil {
		return fmt.Errorf("SendERC20CrossOnt, types.SignTx error: %s", err.Error())
	}

	err = ctx.EthInvoker.ETHUtil.GetEthClient().SendTransaction(context.Background(), signedtx)
	if err != nil {
		return fmt.Errorf("SendERC20CrossOnt, send transaction error:%s", err.Error())
	}
	status.AddTx(signedtx.Hash().String()[2:], &testframework.TxInfo{"BtceToBtc", time.Now()})
	WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), signedtx.Hash())
	return nil
}

func SendUSDTCrossOnt(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	gasPrice, err := ctx.EthInvoker.ETHUtil.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("SendUSDTCrossOnt, get suggest gas price failed error: %s", err.Error())
	}
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(5))

	contractabi, err := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
	if err != nil {
		return fmt.Errorf("SendUSDTCrossOnt, abi.JSON error:" + err.Error())
	}
	assetaddress := ethcommon.HexToAddress(config.DefConfig.EthUSDT)
	contractAddr := ethcommon.HexToAddress(config.DefConfig.EthLockProxy)

	// approve
	usdt, err := usdt_abi.NewTetherToken(assetaddress, ctx.EthInvoker.ETHUtil.GetEthClient())
	if err != nil {
		return fmt.Errorf("SendUSDTCrossOnt, NewERC20 error:" + err.Error())
	}
	val, err := usdt.BalanceOf(nil, ctx.EthInvoker.EthTestSigner.Address)
	if err != nil {
		return fmt.Errorf("SendUSDTCrossOnt, failed to get the balance: %v", err)
	}
	if val.Uint64() < amount {
		return fmt.Errorf("SendUSDTCrossOnt, balance %d is less than amount %d", val.Uint64(), amount)
	}
	nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	auth := MakeEthAuth(ctx.EthInvoker.EthTestSigner, nonce, gasPrice.Uint64(), uint64(eth.DefaultGasLimit))
	if err != nil {
		return fmt.Errorf("SendUSDTCrossOnt, failed to get eth auth: %v", err)
	}
	txhash, err := usdt.Approve(auth, contractAddr, val)
	if err != nil {
		return fmt.Errorf("SendUSDTCrossOnt, approve error: %v", err.Error())
	}
	ctx.EthInvoker.ETHUtil.WaitTransactionConfirm(txhash.Hash())

	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.DefConfig.OntChainID), ctx.OntInvoker.OntAcc.Address[:],
		big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendUSDTCrossOnt, contractabi.Pack error:" + err.Error())
	}
	callMsg := ethereum.CallMsg{
		From: ctx.EthInvoker.EthTestSigner.Address, To: &contractAddr, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(int64(0)), Data: txData,
	}
	gasLimit, err := ctx.EthInvoker.ETHUtil.GetEthClient().EstimateGas(context.Background(), callMsg)
	if err != nil {
		return fmt.Errorf("SendUSDTCrossOnt, estimate gas limit error: %s", err.Error())
	}

	nonce = ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(0)), gasLimit, gasPrice, txData)
	bf := new(bytes.Buffer)
	rlp.Encode(bf, tx)

	rawtx := hexutil.Encode(bf.Bytes())
	unsignedTx, err := eth.DeserializeTx(rawtx)
	if err != nil {
		return fmt.Errorf("SendUSDTCrossOnt, eth.DeserializeTx error: %s", err.Error())
	}
	signedtx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, ctx.EthInvoker.EthTestSigner.PrivateKey)
	if err != nil {
		return fmt.Errorf("SendUSDTCrossOnt, types.SignTx error: %s", err.Error())
	}

	err = ctx.EthInvoker.ETHUtil.GetEthClient().SendTransaction(context.Background(), signedtx)
	if err != nil {
		return fmt.Errorf("SendUSDTCrossOnt, send transaction error:%s", err.Error())
	}
	status.AddTx(signedtx.Hash().String()[2:], &testframework.TxInfo{"USDTToOnt", time.Now()})
	WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), signedtx.Hash())
	return nil
}

func SendOUSDTCrossEth(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	proxyContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntLockProxy)
	if err != nil {
		return fmt.Errorf("SendOUSDTCrossEth, ontcommon.AddressFromHexString error: %s", err)
	}

	assetaddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntUSDT)
	if err != nil {
		return err
	}
	txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		assetaddress,
		[]interface{}{"approve", []interface{}{ctx.OntInvoker.OntAcc.Address, proxyContractAddress, amount}})
	if err != nil {
		return fmt.Errorf("SendOUSDTCrossEth, approve error: %v", err)
	}
	ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)

	to := ctx.EthInvoker.EthTestSigner.Address.Bytes()
	txHash, err = ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{assetaddress, ctx.OntInvoker.OntAcc.Address[:],
			config.DefConfig.EthChainID, to, amount}})
	if err != nil {
		return fmt.Errorf("SendOUSDTCrossEth, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	status.AddTx(hex.EncodeToString(txHash[:]), &testframework.TxInfo{"OUSDTToEth", time.Now()})
	ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)
	return nil
}

func SendWBTCCrossOnt(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	gasPrice, err := ctx.EthInvoker.ETHUtil.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("SendWBTCCrossOnt, get suggest gas price failed error: %s", err.Error())
	}
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(5))

	contractabi, err := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
	if err != nil {
		return fmt.Errorf("SendWBTCCrossOnt, abi.JSON error:" + err.Error())
	}
	assetaddress := ethcommon.HexToAddress(config.DefConfig.EthWBTC)
	contractAddr := ethcommon.HexToAddress(config.DefConfig.EthLockProxy)

	// approve
	token, err := eth.NewStandardToken(assetaddress, ctx.EthInvoker.ETHUtil.GetEthClient())
	if err != nil {
		return fmt.Errorf("SendWBTCCrossOnt, NewStandardToken error: %v", err.Error())
	}
	val, err := token.Allowance(nil, ctx.EthInvoker.EthTestSigner.Address, contractAddr)
	if err != nil {
		return fmt.Errorf("SendWBTCCrossOnt, failed to get allowance: %v", err)
	}
	if val.Uint64() < amount {
		nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
		auth := MakeEthAuth(ctx.EthInvoker.EthTestSigner, nonce, gasPrice.Uint64(), uint64(eth.DefaultGasLimit))
		if err != nil {
			return fmt.Errorf("SendWBTCCrossOnt, failed to get eth auth: %v", err)
		}
		tx, err := token.Approve(auth, contractAddr, big.NewInt(math.MaxInt64))
		if err != nil {
			return fmt.Errorf("SendWBTCCrossOnt, failed to approve: %v", err)
		}
		WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), tx.Hash())
	}

	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.DefConfig.OntChainID), ctx.OntInvoker.OntAcc.Address[:],
		big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendWBTCCrossOnt, contractabi.Pack error:" + err.Error())
	}
	callMsg := ethereum.CallMsg{
		From: ctx.EthInvoker.EthTestSigner.Address, To: &contractAddr, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(int64(0)), Data: txData,
	}
	gasLimit, err := ctx.EthInvoker.ETHUtil.GetEthClient().EstimateGas(context.Background(), callMsg)
	if err != nil {
		return fmt.Errorf("SendWBTCCrossOnt, estimate gas limit error: %s", err.Error())
	}

	nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(0)), gasLimit, gasPrice, txData)
	bf := new(bytes.Buffer)
	rlp.Encode(bf, tx)

	rawtx := hexutil.Encode(bf.Bytes())
	unsignedTx, err := eth.DeserializeTx(rawtx)
	if err != nil {
		return fmt.Errorf("SendWBTCCrossOnt, eth.DeserializeTx error: %s", err.Error())
	}
	signedtx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, ctx.EthInvoker.EthTestSigner.PrivateKey)
	if err != nil {
		return fmt.Errorf("SendWBTCCrossOnt, types.SignTx error: %s", err.Error())
	}

	err = ctx.EthInvoker.ETHUtil.GetEthClient().SendTransaction(context.Background(), signedtx)
	if err != nil {
		return fmt.Errorf("SendWBTCCrossOnt, send transaction error:%s", err.Error())
	}
	status.AddTx(signedtx.Hash().String()[2:], &testframework.TxInfo{"WBTCToOnt", time.Now()})
	WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), signedtx.Hash())
	return nil
}

func SendOWBTCCrossEth(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	proxyContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntLockProxy)
	if err != nil {
		return fmt.Errorf("SendOWBTCCrossEth, ontcommon.AddressFromHexString error: %s", err)
	}

	assetaddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntWBTC)
	if err != nil {
		return err
	}
	res, err := ctx.OntInvoker.OntSdk.NeoVM.PreExecInvokeNeoVMContract(assetaddress, []interface{}{"allowance", []interface{}{ctx.OntInvoker.OntAcc.Address, proxyContractAddress}})
	if err != nil {
		return err
	}
	val, err := res.Result.ToInteger()
	if err != nil {
		return err
	}
	if val.Uint64() < amount {
		txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
			ctx.OntInvoker.OntAcc,
			ctx.OntInvoker.OntAcc,
			assetaddress,
			[]interface{}{"approve", []interface{}{ctx.OntInvoker.OntAcc.Address, proxyContractAddress, math.MaxInt64}})
		if err != nil {
			return fmt.Errorf("SendOWBTCCrossEth, approve error: %v", err)
		}
		ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)
	}

	to := ctx.EthInvoker.EthTestSigner.Address.Bytes()
	txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{assetaddress, ctx.OntInvoker.OntAcc.Address[:],
			config.DefConfig.EthChainID, to, amount}})
	if err != nil {
		return fmt.Errorf("SendOWBTCCrossEth, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	status.AddTx(hex.EncodeToString(txHash[:]), &testframework.TxInfo{"WBTCToEth", time.Now()})
	ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)
	return nil
}

func SendDAICrossOnt(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	gasPrice, err := ctx.EthInvoker.ETHUtil.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		gasPrice = gasPrice.Mul(gasPrice, big.NewInt(5))
		return fmt.Errorf("SendDAICrossOnt, get suggest gas price failed error: %s", err.Error())
	}

	contractabi, err := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
	if err != nil {
		return fmt.Errorf("SendDAICrossOnt, abi.JSON error:" + err.Error())
	}
	assetaddress := ethcommon.HexToAddress(config.DefConfig.EthDai)
	contractAddr := ethcommon.HexToAddress(config.DefConfig.EthLockProxy)

	// approve
	dai, err := eth.NewDai(assetaddress, ctx.EthInvoker.ETHUtil.GetEthClient())
	if err != nil {
		return fmt.Errorf("SendDAICrossOnt, NewERC20 error:" + err.Error())
	}
	val, err := dai.BalanceOf(nil, ctx.EthInvoker.EthTestSigner.Address)
	if err != nil {
		return fmt.Errorf("SendDAICrossOnt, failed to get the balance: %v", err)
	}
	if val.Uint64() < amount {
		return fmt.Errorf("SendDAICrossOnt, balance %d is less than amount %d", val.Uint64(), amount)
	}
	nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	auth := MakeEthAuth(ctx.EthInvoker.EthTestSigner, nonce, gasPrice.Uint64(), uint64(eth.DefaultGasLimit))
	if err != nil {
		return fmt.Errorf("SendDAICrossOnt, failed to get eth auth: %v", err)
	}
	txhash, err := dai.Approve(auth, contractAddr, big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendDAICrossOnt, approve error: %v", err.Error())
	}
	ctx.EthInvoker.ETHUtil.WaitTransactionConfirm(txhash.Hash())

	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.DefConfig.OntChainID), ctx.OntInvoker.OntAcc.Address[:],
		big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendDAICrossOnt, contractabi.Pack error:" + err.Error())
	}
	callMsg := ethereum.CallMsg{
		From: ctx.EthInvoker.EthTestSigner.Address, To: &contractAddr, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(int64(0)), Data: txData,
	}
	gasLimit, err := ctx.EthInvoker.ETHUtil.GetEthClient().EstimateGas(context.Background(), callMsg)
	if err != nil {
		return fmt.Errorf("SendDAICrossOnt, estimate gas limit error: %s", err.Error())
	}

	nonce = ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(0)), gasLimit, gasPrice, txData)
	bf := new(bytes.Buffer)
	rlp.Encode(bf, tx)

	rawtx := hexutil.Encode(bf.Bytes())
	unsignedTx, err := eth.DeserializeTx(rawtx)
	if err != nil {
		return fmt.Errorf("SendDAICrossOnt, eth.DeserializeTx error: %s", err.Error())
	}
	signedtx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, ctx.EthInvoker.EthTestSigner.PrivateKey)
	if err != nil {
		return fmt.Errorf("SendDAICrossOnt, types.SignTx error: %s", err.Error())
	}

	err = ctx.EthInvoker.ETHUtil.GetEthClient().SendTransaction(context.Background(), signedtx)
	if err != nil {
		return fmt.Errorf("SendDAICrossOnt, send transaction error:%s", err.Error())
	}
	status.AddTx(signedtx.Hash().String()[2:], &testframework.TxInfo{"DAIToOnt", time.Now()})
	WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), signedtx.Hash())
	return nil
}

func SendODAICrossEth(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	proxyContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntLockProxy)
	if err != nil {
		return fmt.Errorf("SendODAICrossEth, ontcommon.AddressFromHexString error: %s", err)
	}

	assetaddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntDai)
	if err != nil {
		return err
	}
	txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		assetaddress,
		[]interface{}{"approve", []interface{}{ctx.OntInvoker.OntAcc.Address, proxyContractAddress, amount}})
	if err != nil {
		return fmt.Errorf("SendODAICrossEth, approve error: %v", err)
	}
	ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)

	to := ctx.EthInvoker.EthTestSigner.Address.Bytes()
	txHash, err = ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{assetaddress, ctx.OntInvoker.OntAcc.Address[:],
			config.DefConfig.EthChainID, to, amount}})
	if err != nil {
		return fmt.Errorf("SendODAICrossEth, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	status.AddTx(hex.EncodeToString(txHash[:]), &testframework.TxInfo{"ODAIToEth", time.Now()})
	ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)
	return nil
}

func SendNeoCrossEth(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	lpAddr, err := helper.UInt160FromString(config.DefConfig.NeoLockProxy)
	if err != nil {
		return err
	}

	rawCNeo := helper.ReverseBytes(helper.HexToBytes(config.DefConfig.CNeo))
	fromAsset := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: rawCNeo,
	}
	rawFrom, err := helper.AddressToScriptHash(ctx.NeoInvoker.Acc.Address)
	if err != nil {
		return err
	}
	fromAddr := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: rawFrom.Bytes(),
	}
	toChainId := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *big.NewInt(int64(config.DefConfig.EthChainID)),
	}
	toAddr := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: ctx.EthInvoker.EthTestSigner.Address.Bytes(),
	}
	amt := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *big.NewInt(int64(amount)),
	}
	idx := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *big.NewInt(1),
	}

	tb := tx.NewTransactionBuilder(config.DefConfig.NeoUrl)
	sb := sc.NewScriptBuilder()
	sb.MakeInvocationScript(lpAddr.Bytes(), "lock", []sc.ContractParameter{fromAsset, fromAddr, toChainId, toAddr, amt, idx})
	script := sb.ToArray()

	from, err := helper.AddressToScriptHash(ctx.NeoInvoker.Acc.Address)
	if err != nil {
		return err
	}
	itx, err := tb.MakeInvocationTransaction(script, from, nil, helper.UInt160{}, helper.Zero)
	if err != nil {
		return err
	}
	err = tx.AddSignature(itx, ctx.NeoInvoker.Acc.KeyPair)
	if err != nil {
		return err
	}
	response := tb.Client.SendRawTransaction(itx.RawTransactionString())
	if response.HasError() {
		return fmt.Errorf(response.ErrorResponse.Error.Message)
	}
	_ = itx.HashString()

	status.AddTx(hex.EncodeToString(itx.Hash.Bytes()), &testframework.TxInfo{"NeoToEth", time.Now()})
	neo.WaitNeoTx(ctx.NeoInvoker.Cli, itx.Hash)
	return nil
}

func SendNeoCrossOnt(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	lpAddr, err := helper.UInt160FromString(config.DefConfig.NeoLockProxy)
	if err != nil {
		return err
	}

	rawCNeo, _ := hex.DecodeString(config.DefConfig.CNeo)
	fromAsset := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: helper.ReverseBytes(rawCNeo),
	}
	rawFrom, err := helper.AddressToScriptHash(ctx.NeoInvoker.Acc.Address)
	if err != nil {
		return err
	}
	fromAddr := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: rawFrom.Bytes(),
	}
	toChainId := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *big.NewInt(int64(config.DefConfig.OntChainID)),
	}
	toAddr := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: ctx.OntInvoker.OntAcc.Address[:],
	}
	amt := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *big.NewInt(int64(amount)),
	}

	tb := tx.NewTransactionBuilder(config.DefConfig.NeoUrl)
	sb := sc.NewScriptBuilder()
	sb.MakeInvocationScript(lpAddr.Bytes(), "lock", []sc.ContractParameter{fromAsset, fromAddr, toChainId, toAddr, amt})
	script := sb.ToArray()

	from, err := helper.AddressToScriptHash(ctx.NeoInvoker.Acc.Address)
	if err != nil {
		return err
	}
	itx, err := tb.MakeInvocationTransaction(script, from, nil, helper.UInt160{}, helper.Zero)
	if err != nil {
		return err
	}
	err = tx.AddSignature(itx, ctx.NeoInvoker.Acc.KeyPair)
	if err != nil {
		return err
	}
	response := tb.Client.SendRawTransaction(itx.RawTransactionString())
	if response.HasError() {
		return fmt.Errorf(response.ErrorResponse.Error.Message)
	}
	_ = itx.HashString()

	status.AddTx(hex.EncodeToString(itx.Hash.Bytes()), &testframework.TxInfo{"NeoToOnt", time.Now()})
	neo.WaitNeoTx(ctx.NeoInvoker.Cli, itx.Hash)

	log.Infof("successful to send %d CNEO to Ontology", amount)
	return nil
}

func SendENeoCrossNeo(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	gasPrice, err := ctx.EthInvoker.ETHUtil.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		gasPrice = gasPrice.Mul(gasPrice, big.NewInt(5))
		return fmt.Errorf("SendENeoCrossNeo, get suggest gas price failed error: %s", err.Error())
	}

	contractabi, err := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
	if err != nil {
		return fmt.Errorf("SendENeoCrossNeo, abi.JSON error:" + err.Error())
	}
	assetaddress := ethcommon.HexToAddress(config.DefConfig.EthNeo)
	contractAddr := ethcommon.HexToAddress(config.DefConfig.EthLockProxy)

	// approve
	neox, err := erc20_abi.NewERC20(assetaddress, ctx.EthInvoker.ETHUtil.GetEthClient())
	if err != nil {
		return fmt.Errorf("SendENeoCrossNeo, NewERC20 error:" + err.Error())
	}
	val, err := neox.Allowance(nil, ctx.EthInvoker.EthTestSigner.Address, contractAddr)
	if err != nil {
		return fmt.Errorf("SendENeoCrossNeo, failed to get the balance: %v", err)
	}
	if val.Uint64() < amount {
		nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
		auth := MakeEthAuth(ctx.EthInvoker.EthTestSigner, nonce, gasPrice.Uint64(), uint64(eth.DefaultGasLimit))
		if err != nil {
			return fmt.Errorf("SendEOntdCrossOnt, failed to get eth auth: %v", err)
		}
		tx, err := neox.Approve(auth, contractAddr, big.NewInt(math.MaxInt64))
		if err != nil {
			return fmt.Errorf("SendEOntdCrossOnt, failed to approve: %v", err)
		}
		WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), tx.Hash())
	}

	rawFrom, err := helper.AddressToScriptHash(ctx.NeoInvoker.Acc.Address)
	if err != nil {
		return err
	}
	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.DefConfig.NeoChainID), rawFrom[:],
		big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendENeoCrossNeo, contractabi.Pack error:" + err.Error())
	}
	callMsg := ethereum.CallMsg{
		From: ctx.EthInvoker.EthTestSigner.Address, To: &contractAddr, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(int64(0)), Data: txData,
	}
	gasLimit, err := ctx.EthInvoker.ETHUtil.GetEthClient().EstimateGas(context.Background(), callMsg)
	if err != nil {
		return fmt.Errorf("SendENeoCrossNeo, estimate gas limit error: %s", err.Error())
	}

	nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(0)), gasLimit, gasPrice, txData)
	bf := new(bytes.Buffer)
	rlp.Encode(bf, tx)

	rawtx := hexutil.Encode(bf.Bytes())
	unsignedTx, err := eth.DeserializeTx(rawtx)
	if err != nil {
		return fmt.Errorf("SendENeoCrossNeo, eth.DeserializeTx error: %s", err.Error())
	}
	signedtx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, ctx.EthInvoker.EthTestSigner.PrivateKey)
	if err != nil {
		return fmt.Errorf("SendENeoCrossNeo, types.SignTx error: %s", err.Error())
	}

	err = ctx.EthInvoker.ETHUtil.GetEthClient().SendTransaction(context.Background(), signedtx)
	if err != nil {
		return fmt.Errorf("SendENeoCrossNeo, send transaction error:%s", err.Error())
	}
	status.AddTx(signedtx.Hash().String()[2:], &testframework.TxInfo{"ENeoToNeo", time.Now()})
	WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), signedtx.Hash())
	return nil
}

func SendONeoCrossNeo(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	proxyContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntLockProxy)
	if err != nil {
		return fmt.Errorf("SendONeoCrossNeo, ontcommon.AddressFromHexString error: %s", err)
	}

	assetaddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntNeo)
	if err != nil {
		return err
	}
	res, err := ctx.OntInvoker.OntSdk.NeoVM.PreExecInvokeNeoVMContract(assetaddress, []interface{}{"allowance", []interface{}{ctx.OntInvoker.OntAcc.Address, proxyContractAddress}})
	if err != nil {
		return err
	}
	val, err := res.Result.ToInteger()
	if err != nil {
		return err
	}
	if val.Uint64() < amount {
		txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
			ctx.OntInvoker.OntAcc,
			ctx.OntInvoker.OntAcc,
			assetaddress,
			[]interface{}{"approve", []interface{}{ctx.OntInvoker.OntAcc.Address, proxyContractAddress, math.MaxInt64}})
		if err != nil {
			return fmt.Errorf("SendONeoCrossNeo, approve error: %v", err)
		}
		ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)
	}

	rawFrom, err := helper.AddressToScriptHash(ctx.NeoInvoker.Acc.Address)
	if err != nil {
		return err
	}
	txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{assetaddress, ctx.OntInvoker.OntAcc.Address[:],
			config.DefConfig.NeoChainID, rawFrom.Bytes(), amount}})
	if err != nil {
		return fmt.Errorf("SendONeoCrossNeo, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	status.AddTx(hex.EncodeToString(txHash[:]), &testframework.TxInfo{"ONeoToNeo", time.Now()})
	ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)
	return nil
}

func SendOntCrossNeo(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	proxyContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntLockProxy)
	if err != nil {
		return fmt.Errorf("SendOntCrossNeo, ontcommon.AddressFromHexString error: %s", err)
	}
	rawFrom, err := helper.AddressToScriptHash(ctx.NeoInvoker.Acc.Address)
	if err != nil {
		return err
	}
	txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{nutils.OntContractAddress[:], ctx.OntInvoker.OntAcc.Address[:],
			config.DefConfig.NeoChainID, rawFrom.Bytes(), amount}})
	if err != nil {
		return fmt.Errorf("SendOntCrossNeo, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	status.AddTx(hex.EncodeToString(txHash[:]), &testframework.TxInfo{"OntToNeo", time.Now()})
	return nil
}

func SendNOntCrossOnt(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	lpAddr, err := helper.UInt160FromString(config.DefConfig.NeoLockProxy)
	if err != nil {
		return err
	}

	rawNOnt, _ := hex.DecodeString(config.DefConfig.NeoOnt)
	rawNOnt = helper.ReverseBytes(rawNOnt)
	fromAsset := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: rawNOnt,
	}
	rawFrom, err := helper.AddressToScriptHash(ctx.NeoInvoker.Acc.Address)
	if err != nil {
		return err
	}
	fromAddr := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: rawFrom.Bytes(),
	}
	toChainId := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *big.NewInt(int64(config.DefConfig.OntChainID)),
	}
	toAddr := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: ctx.OntInvoker.OntAcc.Address[:],
	}
	amt := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *big.NewInt(int64(amount)),
	}
	tb := tx.NewTransactionBuilder(config.DefConfig.NeoUrl)
	sb := sc.NewScriptBuilder()
	sb.MakeInvocationScript(lpAddr.Bytes(), "lock", []sc.ContractParameter{fromAsset, fromAddr, toChainId, toAddr, amt})
	script := sb.ToArray()

	from, err := helper.AddressToScriptHash(ctx.NeoInvoker.Acc.Address)
	if err != nil {
		return err
	}
	itx, err := tb.MakeInvocationTransaction(script, from, nil, helper.UInt160{}, helper.Zero)
	if err != nil {
		return err
	}
	err = tx.AddSignature(itx, ctx.NeoInvoker.Acc.KeyPair)
	if err != nil {
		return err
	}
	response := tb.Client.SendRawTransaction(itx.RawTransactionString())
	if response.HasError() {
		return fmt.Errorf(response.ErrorResponse.Error.Message)
	}
	_ = itx.HashString()

	status.AddTx(hex.EncodeToString(itx.Hash.Bytes()), &testframework.TxInfo{"NOntToOnt", time.Now()})
	neo.WaitNeoTx(ctx.NeoInvoker.Cli, itx.Hash)

	log.Infof("successful to send %d NOnt to ontology", amount)
	return nil
}

func SendEthCrossNeo(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	gasPrice, err := ctx.EthInvoker.ETHUtil.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("SendEthCrossNeo, get suggest gas price failed error: %s", err.Error())
	}
	//gasPrice = gasPrice.Mul(gasPrice, big.NewInt(5))

	contractabi, err := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
	if err != nil {
		return fmt.Errorf("SendEthCrossNeo, abi.JSON error:" + err.Error())
	}
	rawFrom, err := helper.AddressToScriptHash(ctx.NeoInvoker.Acc.Address)
	if err != nil {
		return err
	}
	assetaddress := ethcommon.HexToAddress("0000000000000000000000000000000000000000")
	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.DefConfig.NeoChainID), rawFrom[:],
		big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendEthCrossNeo, contractabi.Pack error:" + err.Error())
	}

	contractAddr := ethcommon.HexToAddress(config.DefConfig.EthLockProxy)
	callMsg := ethereum.CallMsg{
		From: ctx.EthInvoker.EthTestSigner.Address, To: &contractAddr, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(int64(amount)), Data: txData,
	}
	gasLimit, err := ctx.EthInvoker.ETHUtil.GetEthClient().EstimateGas(context.Background(), callMsg)
	if err != nil {
		return fmt.Errorf("SendEthCrossNeo, estimate gas limit error: %s", err.Error())
	}

	nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(amount)), gasLimit, gasPrice, txData)
	bf := new(bytes.Buffer)
	rlp.Encode(bf, tx)

	rawtx := hexutil.Encode(bf.Bytes())
	unsignedTx, err := eth.DeserializeTx(rawtx)
	if err != nil {
		return fmt.Errorf("SendEthCrossNeo, eth.DeserializeTx error: %s", err.Error())
	}
	signedtx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, ctx.EthInvoker.EthTestSigner.PrivateKey)
	if err != nil {
		return fmt.Errorf("SendEthCrossNeo, types.SignTx error: %s", err.Error())
	}

	err = ctx.EthInvoker.ETHUtil.GetEthClient().SendTransaction(context.Background(), signedtx)
	if err != nil {
		return fmt.Errorf("SendEthCrossNeo, send transaction error:%s", err.Error())
	}
	status.AddTx(signedtx.Hash().String()[2:], &testframework.TxInfo{"EthToNeo", time.Now()})
	WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), signedtx.Hash())
	return nil
}

func SendNEthCrossEth(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	lpAddr, err := helper.UInt160FromString(config.DefConfig.NeoLockProxy)
	if err != nil {
		return err
	}

	rawNEth, _ := hex.DecodeString(config.DefConfig.NeoEth)
	rawNEth = helper.ReverseBytes(rawNEth)
	fromAsset := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: rawNEth,
	}
	rawFrom, err := helper.AddressToScriptHash(ctx.NeoInvoker.Acc.Address)
	if err != nil {
		return err
	}
	fromAddr := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: rawFrom.Bytes(),
	}
	toChainId := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *big.NewInt(int64(config.DefConfig.EthChainID)),
	}
	toAddr := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: ctx.EthInvoker.EthTestSigner.Address.Bytes(),
	}
	amt := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *big.NewInt(int64(amount)),
	}
	tb := tx.NewTransactionBuilder(config.DefConfig.NeoUrl)
	sb := sc.NewScriptBuilder()
	sb.MakeInvocationScript(lpAddr.Bytes(), "lock", []sc.ContractParameter{fromAsset, fromAddr, toChainId, toAddr, amt})
	script := sb.ToArray()

	from, err := helper.AddressToScriptHash(ctx.NeoInvoker.Acc.Address)
	if err != nil {
		return err
	}
	itx, err := tb.MakeInvocationTransaction(script, from, nil, helper.UInt160{}, helper.Zero)
	if err != nil {
		return err
	}
	err = tx.AddSignature(itx, ctx.NeoInvoker.Acc.KeyPair)
	if err != nil {
		return err
	}
	response := tb.Client.SendRawTransaction(itx.RawTransactionString())
	if response.HasError() {
		return fmt.Errorf(response.ErrorResponse.Error.Message)
	}
	_ = itx.HashString()

	status.AddTx(hex.EncodeToString(itx.Hash.Bytes()), &testframework.TxInfo{"NEthToEth", time.Now()})
	neo.WaitNeoTx(ctx.NeoInvoker.Cli, itx.Hash)

	log.Infof("successful to send %d NEth to Ethereum", amount)
	return nil
}

func SendOntdCrossNeo(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	proxyContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntLockProxy)
	if err != nil {
		return fmt.Errorf("SendOntdCrossNeo, ontcommon.AddressFromHexString error: %s", err)
	}
	rawFrom, err := helper.AddressToScriptHash(ctx.NeoInvoker.Acc.Address)
	if err != nil {
		return err
	}
	asset, err := ontcommon.AddressFromHexString(config.DefConfig.OntONTD)
	if err != nil {
		return err
	}

	res, err := ctx.OntInvoker.OntSdk.NeoVM.PreExecInvokeNeoVMContract(asset, []interface{}{"allowance", []interface{}{ctx.OntInvoker.OntAcc.Address, proxyContractAddress}})
	if err != nil {
		return err
	}
	val, err := res.Result.ToInteger()
	if err != nil {
		return err
	}
	if val.Uint64() < amount {
		txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
			ctx.OntInvoker.OntAcc,
			ctx.OntInvoker.OntAcc,
			asset,
			[]interface{}{"approve", []interface{}{ctx.OntInvoker.OntAcc.Address, proxyContractAddress, math.MaxInt64}})
		if err != nil {
			return fmt.Errorf("SendONeoCrossNeo, approve error: %v", err)
		}
		ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)
	}

	txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{asset[:], ctx.OntInvoker.OntAcc.Address[:],
			config.DefConfig.NeoChainID, rawFrom.Bytes(), amount}})
	if err != nil {
		return fmt.Errorf("SendOntdCrossNeo, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	status.AddTx(hex.EncodeToString(txHash[:]), &testframework.TxInfo{"OntdToNeo", time.Now()})

	return nil
}

func SendOntdCrossEth(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	proxyContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntLockProxy)
	if err != nil {
		return fmt.Errorf("SendOntdCrossEth, ontcommon.AddressFromHexString error: %s", err)
	}
	to := ctx.EthInvoker.EthTestSigner.Address.Bytes()
	asset, err := ontcommon.AddressFromHexString(config.DefConfig.OntONTD)
	if err != nil {
		return err
	}

	res, err := ctx.OntInvoker.OntSdk.NeoVM.PreExecInvokeNeoVMContract(asset, []interface{}{"allowance", []interface{}{ctx.OntInvoker.OntAcc.Address, proxyContractAddress}})
	if err != nil {
		return err
	}
	val, err := res.Result.ToInteger()
	if err != nil {
		return err
	}
	if val.Uint64() < amount {
		txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
			ctx.OntInvoker.OntAcc,
			ctx.OntInvoker.OntAcc,
			asset,
			[]interface{}{"approve", []interface{}{ctx.OntInvoker.OntAcc.Address, proxyContractAddress, math.MaxInt64}})
		if err != nil {
			return fmt.Errorf("SendONeoCrossNeo, approve error: %v", err)
		}
		ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)
	}

	txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{asset[:], ctx.OntInvoker.OntAcc.Address[:],
			config.DefConfig.EthChainID, to, amount}})
	if err != nil {
		return fmt.Errorf("SendOntdCrossEth, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	status.AddTx(hex.EncodeToString(txHash[:]), &testframework.TxInfo{"OntdToEth", time.Now()})

	return nil
}

func SendEOntdCrossOnt(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	gasPrice, err := ctx.EthInvoker.ETHUtil.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("SendEOntdCrossOnt, get suggest gas price failed error: %s", err.Error())
	}
	//gasPrice = gasPrice.Mul(gasPrice, big.NewInt(5))

	assetaddress := ethcommon.HexToAddress(config.DefConfig.EthOntd)
	contractAddr := ethcommon.HexToAddress(config.DefConfig.EthLockProxy)
	token, err := eth.NewStandardToken(assetaddress, ctx.EthInvoker.ETHUtil.GetEthClient())
	if err != nil {
		return fmt.Errorf("SendEOntdCrossOnt, NewStandardToken error: %v", err.Error())
	}
	val, err := token.Allowance(nil, ctx.EthInvoker.EthTestSigner.Address, contractAddr)
	if err != nil {
		return fmt.Errorf("SendEOntdCrossOnt, failed to get allowance: %v", err)
	}
	if val.Uint64() < amount {
		nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
		auth := MakeEthAuth(ctx.EthInvoker.EthTestSigner, nonce, gasPrice.Uint64(), uint64(eth.DefaultGasLimit))
		if err != nil {
			return fmt.Errorf("SendEOntdCrossOnt, failed to get eth auth: %v", err)
		}
		tx, err := token.Approve(auth, contractAddr, big.NewInt(math.MaxInt64))
		if err != nil {
			return fmt.Errorf("SendEOntdCrossOnt, failed to approve: %v", err)
		}
		WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), tx.Hash())
	}

	contractabi, err := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
	if err != nil {
		return fmt.Errorf("SendEOntdCrossOnt, abi.JSON error:" + err.Error())
	}

	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.DefConfig.OntChainID), ctx.OntInvoker.OntAcc.Address[:],
		big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendEOntdCrossOnt, contractabi.Pack error:" + err.Error())
	}

	callMsg := ethereum.CallMsg{
		From: ctx.EthInvoker.EthTestSigner.Address, To: &contractAddr, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(int64(0)), Data: txData,
	}
	gasLimit, err := ctx.EthInvoker.ETHUtil.GetEthClient().EstimateGas(context.Background(), callMsg)
	if err != nil {
		return fmt.Errorf("SendEOntdCrossOnt, estimate gas limit error: %s", err.Error())
	}

	nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(0)), gasLimit, gasPrice, txData)
	bf := new(bytes.Buffer)
	rlp.Encode(bf, tx)

	rawtx := hexutil.Encode(bf.Bytes())
	unsignedTx, err := eth.DeserializeTx(rawtx)
	if err != nil {
		return fmt.Errorf("SendEOntdCrossOnt, eth.DeserializeTx error: %s", err.Error())
	}
	signedtx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, ctx.EthInvoker.EthTestSigner.PrivateKey)
	if err != nil {
		return fmt.Errorf("SendEOntdCrossOnt, types.SignTx error: %s", err.Error())
	}

	err = ctx.EthInvoker.ETHUtil.GetEthClient().SendTransaction(context.Background(), signedtx)
	if err != nil {
		return fmt.Errorf("SendEOntdCrossOnt, send transaction error:%s", err.Error())
	}
	status.AddTx(signedtx.Hash().String()[2:], &testframework.TxInfo{"EONTDToOnt", time.Now()})
	WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), signedtx.Hash())
	return nil
}

func SendEOntdCrossNeo(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	gasPrice, err := ctx.EthInvoker.ETHUtil.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("SendEOntdCrossNeo, get suggest gas price failed error: %s", err.Error())
	}
	//gasPrice = gasPrice.Mul(gasPrice, big.NewInt(5))

	assetaddress := ethcommon.HexToAddress(config.DefConfig.EthOntd)
	contractAddr := ethcommon.HexToAddress(config.DefConfig.EthLockProxy)
	token, err := eth.NewStandardToken(assetaddress, ctx.EthInvoker.ETHUtil.GetEthClient())
	if err != nil {
		return fmt.Errorf("SendEOntdCrossNeo, NewStandardToken error: %v", err.Error())
	}
	val, err := token.Allowance(nil, ctx.EthInvoker.EthTestSigner.Address, contractAddr)
	if err != nil {
		return fmt.Errorf("SendEOntdCrossNeo, failed to get allowance: %v", err)
	}
	if val.Uint64() < amount {
		nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
		auth := MakeEthAuth(ctx.EthInvoker.EthTestSigner, nonce, gasPrice.Uint64(), uint64(eth.DefaultGasLimit))
		if err != nil {
			return fmt.Errorf("SendEOntdCrossNeo, failed to get eth auth: %v", err)
		}
		tx, err := token.Approve(auth, contractAddr, big.NewInt(math.MaxInt64))
		if err != nil {
			return fmt.Errorf("SendEOntdCrossNeo, failed to approve: %v", err)
		}
		WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), tx.Hash())
	}

	contractabi, err := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
	if err != nil {
		return fmt.Errorf("SendEOntdCrossNeo, abi.JSON error:" + err.Error())
	}
	rawFrom, err := helper.AddressToScriptHash(ctx.NeoInvoker.Acc.Address)
	if err != nil {
		return err
	}
	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.DefConfig.NeoChainID), rawFrom[:],
		big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendEOntdCrossNeo, contractabi.Pack error:" + err.Error())
	}

	callMsg := ethereum.CallMsg{
		From: ctx.EthInvoker.EthTestSigner.Address, To: &contractAddr, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(int64(0)), Data: txData,
	}
	gasLimit, err := ctx.EthInvoker.ETHUtil.GetEthClient().EstimateGas(context.Background(), callMsg)
	if err != nil {
		return fmt.Errorf("SendEOntdCrossNeo, estimate gas limit error: %s", err.Error())
	}

	nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(0)), gasLimit, gasPrice, txData)
	bf := new(bytes.Buffer)
	rlp.Encode(bf, tx)

	rawtx := hexutil.Encode(bf.Bytes())
	unsignedTx, err := eth.DeserializeTx(rawtx)
	if err != nil {
		return fmt.Errorf("SendEOntdCrossNeo, eth.DeserializeTx error: %s", err.Error())
	}
	signedtx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, ctx.EthInvoker.EthTestSigner.PrivateKey)
	if err != nil {
		return fmt.Errorf("SendEOntdCrossNeo, types.SignTx error: %s", err.Error())
	}

	err = ctx.EthInvoker.ETHUtil.GetEthClient().SendTransaction(context.Background(), signedtx)
	if err != nil {
		return fmt.Errorf("SendEOntdCrossNeo, send transaction error:%s", err.Error())
	}
	status.AddTx(signedtx.Hash().String()[2:], &testframework.TxInfo{"EONTDToNeo", time.Now()})
	WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), signedtx.Hash())
	return nil
}

func SendNOntdCrossOnt(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	lpAddr, err := helper.UInt160FromString(config.DefConfig.NeoLockProxy)
	if err != nil {
		return err
	}

	rawNOntd, _ := hex.DecodeString(config.DefConfig.NeoOntd)
	rawNOntd = helper.ReverseBytes(rawNOntd)
	fromAsset := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: rawNOntd,
	}
	rawFrom, err := helper.AddressToScriptHash(ctx.NeoInvoker.Acc.Address)
	if err != nil {
		return err
	}
	fromAddr := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: rawFrom.Bytes(),
	}
	toChainId := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *big.NewInt(int64(config.DefConfig.OntChainID)),
	}
	toAddr := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: ctx.OntInvoker.OntAcc.Address[:],
	}
	amt := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *big.NewInt(int64(amount)),
	}
	tb := tx.NewTransactionBuilder(config.DefConfig.NeoUrl)
	sb := sc.NewScriptBuilder()
	sb.MakeInvocationScript(lpAddr.Bytes(), "lock", []sc.ContractParameter{fromAsset, fromAddr, toChainId, toAddr, amt})
	script := sb.ToArray()

	from, err := helper.AddressToScriptHash(ctx.NeoInvoker.Acc.Address)
	if err != nil {
		return err
	}
	itx, err := tb.MakeInvocationTransaction(script, from, nil, helper.UInt160{}, helper.Zero)
	if err != nil {
		return err
	}
	err = tx.AddSignature(itx, ctx.NeoInvoker.Acc.KeyPair)
	if err != nil {
		return err
	}
	response := tb.Client.SendRawTransaction(itx.RawTransactionString())
	if response.HasError() {
		return fmt.Errorf(response.ErrorResponse.Error.Message)
	}
	_ = itx.HashString()

	status.AddTx(hex.EncodeToString(itx.Hash.Bytes()), &testframework.TxInfo{"NOntdToOnt", time.Now()})
	neo.WaitNeoTx(ctx.NeoInvoker.Cli, itx.Hash)

	log.Infof("successful to send %d NOntd to ontology", amount)
	return nil
}

func SendNOntdCrossEth(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	lpAddr, err := helper.UInt160FromString(config.DefConfig.NeoLockProxy)
	if err != nil {
		return err
	}

	rawNOntd, _ := hex.DecodeString(config.DefConfig.NeoOntd)
	rawNOntd = helper.ReverseBytes(rawNOntd)
	fromAsset := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: rawNOntd,
	}
	rawFrom, err := helper.AddressToScriptHash(ctx.NeoInvoker.Acc.Address)
	if err != nil {
		return err
	}
	fromAddr := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: rawFrom.Bytes(),
	}
	toChainId := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *big.NewInt(int64(config.DefConfig.EthChainID)),
	}
	toAddr := sc.ContractParameter{
		Type:  sc.ByteArray,
		Value: ctx.EthInvoker.EthTestSigner.Address.Bytes(),
	}
	amt := sc.ContractParameter{
		Type:  sc.Integer,
		Value: *big.NewInt(int64(amount)),
	}
	tb := tx.NewTransactionBuilder(config.DefConfig.NeoUrl)
	sb := sc.NewScriptBuilder()
	sb.MakeInvocationScript(lpAddr.Bytes(), "lock", []sc.ContractParameter{fromAsset, fromAddr, toChainId, toAddr, amt})
	script := sb.ToArray()

	from, err := helper.AddressToScriptHash(ctx.NeoInvoker.Acc.Address)
	if err != nil {
		return err
	}
	itx, err := tb.MakeInvocationTransaction(script, from, nil, helper.UInt160{}, helper.Zero)
	if err != nil {
		return err
	}
	err = tx.AddSignature(itx, ctx.NeoInvoker.Acc.KeyPair)
	if err != nil {
		return err
	}
	response := tb.Client.SendRawTransaction(itx.RawTransactionString())
	if response.HasError() {
		return fmt.Errorf(response.ErrorResponse.Error.Message)
	}
	_ = itx.HashString()

	status.AddTx(hex.EncodeToString(itx.Hash.Bytes()), &testframework.TxInfo{"NOntdToEth", time.Now()})
	neo.WaitNeoTx(ctx.NeoInvoker.Cli, itx.Hash)

	log.Infof("successful to send %d NOntd to ethereum", amount)
	return nil
}

func SendUSDCCrossOnt(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	gasPrice, err := ctx.EthInvoker.ETHUtil.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("SendUSDCCrossOnt, get suggest gas price failed error: %s", err.Error())
	}

	contractabi, err := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
	if err != nil {
		return fmt.Errorf("SendUSDCCrossOnt, abi.JSON error:" + err.Error())
	}
	assetaddress := ethcommon.HexToAddress(config.DefConfig.EthUSDC)
	contractAddr := ethcommon.HexToAddress(config.DefConfig.EthLockProxy)

	// approve
	token, err := eth.NewStandardToken(assetaddress, ctx.EthInvoker.ETHUtil.GetEthClient())
	if err != nil {
		return fmt.Errorf("SendUSDCCrossOnt, NewStandardToken error: %v", err.Error())
	}
	val, err := token.Allowance(nil, ctx.EthInvoker.EthTestSigner.Address, contractAddr)
	if err != nil {
		return fmt.Errorf("SendUSDCCrossOnt, failed to get allowance: %v", err)
	}
	if val.Uint64() < amount {
		nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
		auth := MakeEthAuth(ctx.EthInvoker.EthTestSigner, nonce, gasPrice.Uint64(), uint64(eth.DefaultGasLimit))
		if err != nil {
			return fmt.Errorf("SendUSDCCrossOnt, failed to get eth auth: %v", err)
		}
		tx, err := token.Approve(auth, contractAddr, big.NewInt(math.MaxInt64))
		if err != nil {
			return fmt.Errorf("SendUSDCCrossOnt, failed to approve: %v", err)
		}
		WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), tx.Hash())
	}

	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.DefConfig.OntChainID), ctx.OntInvoker.OntAcc.Address[:],
		big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendUSDCCrossOnt, contractabi.Pack error:" + err.Error())
	}
	callMsg := ethereum.CallMsg{
		From: ctx.EthInvoker.EthTestSigner.Address, To: &contractAddr, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(int64(0)), Data: txData,
	}
	gasLimit, err := ctx.EthInvoker.ETHUtil.GetEthClient().EstimateGas(context.Background(), callMsg)
	if err != nil {
		return fmt.Errorf("SendUSDCCrossOnt, estimate gas limit error: %s", err.Error())
	}

	nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(0)), gasLimit, gasPrice, txData)
	bf := new(bytes.Buffer)
	rlp.Encode(bf, tx)

	rawtx := hexutil.Encode(bf.Bytes())
	unsignedTx, err := eth.DeserializeTx(rawtx)
	if err != nil {
		return fmt.Errorf("SendUSDCCrossOnt, eth.DeserializeTx error: %s", err.Error())
	}
	signedtx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, ctx.EthInvoker.EthTestSigner.PrivateKey)
	if err != nil {
		return fmt.Errorf("SendUSDCCrossOnt, types.SignTx error: %s", err.Error())
	}

	err = ctx.EthInvoker.ETHUtil.GetEthClient().SendTransaction(context.Background(), signedtx)
	if err != nil {
		return fmt.Errorf("SendUSDCCrossOnt, send transaction error:%s", err.Error())
	}
	status.AddTx(signedtx.Hash().String()[2:], &testframework.TxInfo{"USDCToOnt", time.Now()})
	WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), signedtx.Hash())
	return nil
}

func SendOUSDCCrossEth(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	proxyContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntLockProxy)
	if err != nil {
		return fmt.Errorf("SendOUSDCCrossEth, ontcommon.AddressFromHexString error: %s", err)
	}

	assetaddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntUSDC)
	if err != nil {
		return err
	}
	res, err := ctx.OntInvoker.OntSdk.NeoVM.PreExecInvokeNeoVMContract(assetaddress,
		[]interface{}{"allowance", []interface{}{ctx.OntInvoker.OntAcc.Address, proxyContractAddress}})
	if err != nil {
		return err
	}
	val, err := res.Result.ToInteger()
	if err != nil {
		return err
	}
	if val.Uint64() < amount {
		txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
			ctx.OntInvoker.OntAcc,
			ctx.OntInvoker.OntAcc,
			assetaddress,
			[]interface{}{"approve", []interface{}{ctx.OntInvoker.OntAcc.Address, proxyContractAddress, math.MaxInt64}})
		if err != nil {
			return fmt.Errorf("SendOUSDCCrossEth, approve error: %v", err)
		}
		ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)
	}

	to := ctx.EthInvoker.EthTestSigner.Address.Bytes()
	txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{assetaddress, ctx.OntInvoker.OntAcc.Address[:],
			config.DefConfig.EthChainID, to, amount}})
	if err != nil {
		return fmt.Errorf("SendOUSDCCrossEth, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	status.AddTx(hex.EncodeToString(txHash[:]), &testframework.TxInfo{"OUSDCToEth", time.Now()})
	ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)
	return nil
}

func SendRenBTCCrossOnt(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	gasPrice, err := ctx.EthInvoker.ETHUtil.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("SendRenBTCCrossOnt, get suggest gas price failed error: %s", err.Error())
	}

	contractabi, err := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
	if err != nil {
		return fmt.Errorf("SendRenBTCCrossOnt, abi.JSON error:" + err.Error())
	}
	assetaddress := ethcommon.HexToAddress(config.DefConfig.EthRenBTC)
	contractAddr := ethcommon.HexToAddress(config.DefConfig.EthLockProxy)

	// approve
	token, err := eth.NewStandardToken(assetaddress, ctx.EthInvoker.ETHUtil.GetEthClient())
	if err != nil {
		return fmt.Errorf("SendRenBTCCrossOnt, NewStandardToken error: %v", err.Error())
	}
	val, err := token.Allowance(nil, ctx.EthInvoker.EthTestSigner.Address, contractAddr)
	if err != nil {
		return fmt.Errorf("SendRenBTCCrossOnt, failed to get allowance: %v", err)
	}
	if val.Uint64() < amount {
		nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
		auth := MakeEthAuth(ctx.EthInvoker.EthTestSigner, nonce, gasPrice.Uint64(), uint64(eth.DefaultGasLimit))
		if err != nil {
			return fmt.Errorf("SendRenBTCCrossOnt, failed to get eth auth: %v", err)
		}
		tx, err := token.Approve(auth, contractAddr, big.NewInt(math.MaxInt64))
		if err != nil {
			return fmt.Errorf("SendRenBTCCrossOnt, failed to approve: %v", err)
		}
		WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), tx.Hash())
	}

	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.DefConfig.OntChainID), ctx.OntInvoker.OntAcc.Address[:],
		big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendRenBTCCrossOnt, contractabi.Pack error:" + err.Error())
	}
	callMsg := ethereum.CallMsg{
		From: ctx.EthInvoker.EthTestSigner.Address, To: &contractAddr, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(int64(0)), Data: txData,
	}
	gasLimit, err := ctx.EthInvoker.ETHUtil.GetEthClient().EstimateGas(context.Background(), callMsg)
	if err != nil {
		return fmt.Errorf("SendRenBTCCrossOnt, estimate gas limit error: %s", err.Error())
	}

	nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(0)), gasLimit, gasPrice, txData)
	bf := new(bytes.Buffer)
	rlp.Encode(bf, tx)

	rawtx := hexutil.Encode(bf.Bytes())
	unsignedTx, err := eth.DeserializeTx(rawtx)
	if err != nil {
		return fmt.Errorf("SendRenBTCCrossOnt, eth.DeserializeTx error: %s", err.Error())
	}
	signedtx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, ctx.EthInvoker.EthTestSigner.PrivateKey)
	if err != nil {
		return fmt.Errorf("SendRenBTCCrossOnt, types.SignTx error: %s", err.Error())
	}

	err = ctx.EthInvoker.ETHUtil.GetEthClient().SendTransaction(context.Background(), signedtx)
	if err != nil {
		return fmt.Errorf("SendRenBTCCrossOnt, send transaction error:%s", err.Error())
	}
	status.AddTx(signedtx.Hash().String()[2:], &testframework.TxInfo{"RenBTCToOnt", time.Now()})
	WaitTransactionConfirm(ctx.EthInvoker.ETHUtil.GetEthClient(), signedtx.Hash())
	return nil
}

func SendORenBTCCrossEth(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	proxyContractAddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntLockProxy)
	if err != nil {
		return fmt.Errorf("SendORenBTCCrossEth, ontcommon.AddressFromHexString error: %s", err)
	}

	assetaddress, err := ontcommon.AddressFromHexString(config.DefConfig.OntRenBTC)
	if err != nil {
		return err
	}
	res, err := ctx.OntInvoker.OntSdk.NeoVM.PreExecInvokeNeoVMContract(assetaddress,
		[]interface{}{"allowance", []interface{}{ctx.OntInvoker.OntAcc.Address, proxyContractAddress}})
	if err != nil {
		return err
	}
	val, err := res.Result.ToInteger()
	if err != nil {
		return err
	}
	if val.Uint64() < amount {
		txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
			ctx.OntInvoker.OntAcc,
			ctx.OntInvoker.OntAcc,
			assetaddress,
			[]interface{}{"approve", []interface{}{ctx.OntInvoker.OntAcc.Address, proxyContractAddress, math.MaxInt64}})
		if err != nil {
			return fmt.Errorf("SendORenBTCCrossEth, approve error: %v", err)
		}
		ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)
	}

	to := ctx.EthInvoker.EthTestSigner.Address.Bytes()
	txHash, err := ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{assetaddress, ctx.OntInvoker.OntAcc.Address[:],
			config.DefConfig.EthChainID, to, amount}})
	if err != nil {
		return fmt.Errorf("SendORenBTCCrossEth, ctx.Ont.NeoVM.InvokeNeoVMContract error: %s", err)
	}
	status.AddTx(hex.EncodeToString(txHash[:]), &testframework.TxInfo{"ORenBTCToEth", time.Now()})
	ont.WaitOntTx(txHash, ctx.OntInvoker.OntSdk)
	return nil
}

func ToArrayReverse(arr []byte) []byte {
	l := len(arr)
	x := make([]byte, 0)
	for i := l - 1; i >= 0; i-- {
		x = append(x, arr[i])
	}
	return x
}
