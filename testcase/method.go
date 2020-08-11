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
	"github.com/ontio/ontology-go-sdk"
	ontcommon "github.com/ontio/ontology/common"
	nutils "github.com/ontio/ontology/smartcontract/service/native/utils"
	"github.com/polynetwork/cross_chain_test/chains/eth"
	btcx_abi "github.com/polynetwork/cross_chain_test/chains/eth/abi/btcx"
	"github.com/polynetwork/cross_chain_test/chains/eth/abi/erc20"
	lock_proxy_abi "github.com/polynetwork/cross_chain_test/chains/eth/abi/lockproxy"
	oep4_abi "github.com/polynetwork/cross_chain_test/chains/eth/abi/oep4"
	"github.com/polynetwork/cross_chain_test/chains/eth/abi/ongx"
	"github.com/polynetwork/cross_chain_test/chains/eth/abi/ontx"
	"github.com/polynetwork/cross_chain_test/chains/ont"
	"github.com/polynetwork/cross_chain_test/config"
	"github.com/polynetwork/cross_chain_test/log"
	"github.com/polynetwork/cross_chain_test/testframework"
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
			config.ETH_CHAIN_ID, to, amount}})
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
		return fmt.Errorf("SendEOntCrossOnt, abi.JSON error:" + err.Error())
	}

	assetaddress := ethcommon.HexToAddress(onte)
	contractAddr := ethcommon.HexToAddress(config.DefConfig.EthLockProxy)

	ontxContract, err := ontx.NewONTX(assetaddress, ctx.EthInvoker.ETHUtil.GetEthClient())
	if err != nil {
		return fmt.Errorf("SendEOntCrossOnt, NewONTX error:" + err.Error())
	}
	nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	auth := MakeEthAuth(ctx.EthInvoker.EthTestSigner, nonce, gasPrice.Uint64(), uint64(eth.DefaultGasLimit))
	if err != nil {
		return fmt.Errorf("SendEOntCrossOnt, failed to get eth auth: %v", err)
	}
	txhash, err := ontxContract.Approve(auth, contractAddr, big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendEOntCrossOnt, failed to approve: %v", err)
	}
	ctx.EthInvoker.ETHUtil.WaitTransactionConfirm(txhash.Hash())

	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.ONT_CHAIN_ID), ctx.OntInvoker.OntAcc.Address[:],
		big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendEOntCrossOnt, contractabi.Pack error:" + err.Error())
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
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(amount)), gasLimit, gasPrice, txData)
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
			config.ETH_CHAIN_ID, to, amount}})
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
	ongxContract, err := ongx.NewONGX(assetaddress, ctx.EthInvoker.ETHUtil.GetEthClient())
	if err != nil {
		return fmt.Errorf("SendOngeCrossOnt, NewONTX error:" + err.Error())
	}
	nonce := ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	auth := MakeEthAuth(ctx.EthInvoker.EthTestSigner, nonce, gasPrice.Uint64(), uint64(eth.DefaultGasLimit))
	if err != nil {
		return fmt.Errorf("SendOngeCrossOnt, failed to get eth auth: %v", err)
	}
	txhash, err := ongxContract.Approve(auth, contractAddr, big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendOngeCrossOnt, failed to approve: %v", err)
	}
	ctx.EthInvoker.ETHUtil.WaitTransactionConfirm(txhash.Hash())

	val, err := ongxContract.Allowance(nil, ctx.EthInvoker.EthTestSigner.Address, contractAddr)
	if err != nil {
		return fmt.Errorf("SendOngeCrossOnt, failed to Allowance: %v", err)
	}
	if val.Uint64() < amount {
		return fmt.Errorf("SendOngeCrossOnt, allowance %d is less than amount %d", val.Uint64(), amount)
	}

	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.ONT_CHAIN_ID), ctx.OntInvoker.OntAcc.Address[:],
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

	nonce = ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(amount)), gasLimit, gasPrice, txData)
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
			config.ETH_CHAIN_ID, to, amount}})
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

	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.ONT_CHAIN_ID), ctx.OntInvoker.OntAcc.Address[:],
		big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendEOEP4CrossOnt, contractabi.Pack error:" + err.Error())
	}
	nonce = ctx.EthInvoker.NM.GetAddressNonce(ctx.EthInvoker.EthTestSigner.Address)
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(amount)), uint64(eth.DefaultGasLimit), gasPrice, txData)
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
		[]interface{}{"lock", []interface{}{config.BTC_CHAIN_ID, ctx.OntInvoker.OntAcc.Address[:], to, amount}})
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
		[]interface{}{"lock", []interface{}{config.ETH_CHAIN_ID, ctx.OntInvoker.OntAcc.Address[:],
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
	txData, err := contractabi.Pack("lock", uint64(config.ONT_CHAIN_ID), ctx.OntInvoker.OntAcc.Address[:], uint64(amount))
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
	txid, err := sendBtcCross(ctx, config.ONT_CHAIN_ID, ctx.BtcInvoker.Signer, ctx.OntInvoker.OntAcc.Address.ToBase58(), amount)
	if err != nil {
		return fmt.Errorf("SendBtcCrossOnt, sendBtcCross error: %s", err)
	}
	status.AddTx(txid, &testframework.TxInfo{"BtcToOnt", time.Now()})
	return nil
}

func SendBtcCrossEth(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount int64) error {
	txid, err := sendBtcCross(ctx, config.ETH_CHAIN_ID, ctx.BtcInvoker.Signer, ctx.EthInvoker.EthTestSigner.Address.String(), amount)
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
	tx, err := ctx.CMInvoker.SendAsset(config.CM_BTCX, config.BTC_CHAIN_ID, int64(amt),
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
	tx, err := ctx.CMInvoker.SendAsset(config.CM_BTCX, config.ETH_CHAIN_ID, int64(amt),
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
	tx, err := ctx.CMInvoker.SendAsset(config.CM_BTCX, config.ONT_CHAIN_ID, int64(amt), ctx.OntInvoker.OntAcc.Address[:], lp)
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
	tx, err := ctx.CMInvoker.SendAsset(config.CM_ETHX, config.ETH_CHAIN_ID, int64(amt),
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
	tx, err := ctx.CMInvoker.SendAsset(config.CM_ETHX, config.ONT_CHAIN_ID, int64(amt), ctx.OntInvoker.OntAcc.Address[:], lp)
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
	erc20Contract, err := erc20.NewERC20(assetaddress, ctx.EthInvoker.ETHUtil.GetEthClient())
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
	tx, err := ctx.CMInvoker.SendAsset(config.CM_ERC20, config.ETH_CHAIN_ID, int64(amt),
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
	tx, err := ctx.CMInvoker.SendAsset(config.CM_ERC20, config.ONT_CHAIN_ID, int64(amt), ctx.OntInvoker.OntAcc.Address[:], lp)
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
	tx, err := ctx.CMInvoker.SendAsset(config.CM_ONT, config.ONT_CHAIN_ID, int64(amt), ctx.OntInvoker.OntAcc.Address[:], lp)
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
	tx, err := ctx.CMInvoker.SendAsset(config.CM_ONT, config.ETH_CHAIN_ID, int64(amt),
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

	ontxContract, err := ontx.NewONTX(assetaddress, ctx.EthInvoker.ETHUtil.GetEthClient())
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
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(amount)), gasLimit, gasPrice, txData)
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
	tx, err := ctx.CMInvoker.SendAsset(config.CM_ONG, config.ONT_CHAIN_ID, int64(amt), ctx.OntInvoker.OntAcc.Address[:], lp)
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
	tx, err := ctx.CMInvoker.SendAsset(config.CM_ONG, config.ETH_CHAIN_ID, int64(amt),
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

	ontxContract, err := ontx.NewONTX(assetaddress, ctx.EthInvoker.ETHUtil.GetEthClient())
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
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(amount)), gasLimit, gasPrice, txData)
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
	tx, err := ctx.CMInvoker.SendAsset(config.CM_OEP4, config.ONT_CHAIN_ID, int64(amt), ctx.OntInvoker.OntAcc.Address[:], lp)
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
	tx, err := ctx.CMInvoker.SendAsset("oep4", config.ETH_CHAIN_ID, int64(amt),
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

	ontxContract, err := ontx.NewONTX(assetaddress, ctx.EthInvoker.ETHUtil.GetEthClient())
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
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(amount)), gasLimit, gasPrice, txData)
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
	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.ONT_CHAIN_ID), ctx.OntInvoker.OntAcc.Address[:],
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

	to := ctx.EthInvoker.EthTestSigner.Address.Bytes()
	txHash, err = ctx.OntInvoker.OntSdk.NeoVM.InvokeNeoVMContract(config.DefConfig.GasPrice, config.DefConfig.GasLimit,
		ctx.OntInvoker.OntAcc,
		ctx.OntInvoker.OntAcc,
		proxyContractAddress,
		[]interface{}{"lock", []interface{}{assetaddress, ctx.OntInvoker.OntAcc.Address[:],
			config.ETH_CHAIN_ID, to, amount}})
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
	erc20Contract, err := erc20.NewERC20(assetaddress, ctx.EthInvoker.ETHUtil.GetEthClient())
	if err != nil {
		return fmt.Errorf("SendERC20CrossOnt, NewERC20 error:" + err.Error())
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

	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.ONT_CHAIN_ID), ctx.OntInvoker.OntAcc.Address[:],
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
			config.ETH_CHAIN_ID, to, amount}})
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
	txData, err := contractabi.Pack("lock", uint64(config.BTC_CHAIN_ID), []byte(ctx.BtcInvoker.Signer.Address), uint64(amount))
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

func ToArrayReverse(arr []byte) []byte {
	l := len(arr)
	x := make([]byte, 0)
	for i := l - 1; i >= 0; i-- {
		x = append(x, arr[i])
	}
	return x
}
