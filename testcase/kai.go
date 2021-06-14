package testcase

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/polynetwork/eth-contracts/go_abi/lock_proxy_abi"
	"github.com/polynetwork/poly-io-test/chains/eth"
	"github.com/polynetwork/poly-io-test/config"
	"github.com/polynetwork/poly-io-test/log"
	"github.com/polynetwork/poly-io-test/testframework"
)

func SendKaiToEthChain(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	err := SendKaiCrossEth(ctx, status, 1000000000000000000)
	if err != nil {
		log.Error("SendKaiToEthChain: %s", err)
		return false
	}
	WaitUntilClean(status)
	status.SetItSuccess(1)
	return true
}

func SendEthToKaiChain(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus) bool {
	err := SendEthCrossKai(ctx, status, GetRandAmount(config.DefConfig.OntValLimit, 1))
	if err != nil {
		log.Error("SendEthToKaiChain: %s", err)
		return false
	}
	WaitUntilClean(status)
	status.SetItSuccess(1)
	return true
}

func SendKaiCrossEth(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	kclient := ctx.KaiInvoker.ETHUtil.GetEthClient()
	gasPrice, err := kclient.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("SendKaiCrossEth, get suggest gas price failed error: %s", err.Error())
	}
	//gasPrice = gasPrice.Mul(gasPrice, big.NewInt(5))

	contractabi, err := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
	if err != nil {
		return fmt.Errorf("SendKaiCrossEth, abi.JSON error:" + err.Error())
	}
	rawFrom := ctx.KaiInvoker.EthTestSigner.Address.Bytes()
	assetaddress := ethcommon.HexToAddress("0000000000000000000000000000000000000000")
	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.DefConfig.KaiChainID), rawFrom[:],
		big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendKaiCrossEth, contractabi.Pack error:" + err.Error())
	}

	contractAddr := ethcommon.HexToAddress(config.DefConfig.KaiLockProxy)
	callMsg := ethereum.CallMsg{
		From: ctx.KaiInvoker.EthTestSigner.Address, To: &contractAddr, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(int64(amount)), Data: txData,
	}
	gasLimit, err := kclient.EstimateGas(context.Background(), callMsg)
	if err != nil {
		return fmt.Errorf("SendKaiCrossEth, estimate gas limit error: %s", err.Error())
	}

	nonce := ctx.KaiInvoker.NM.GetAddressNonce(ctx.KaiInvoker.EthTestSigner.Address)
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(amount)), gasLimit, gasPrice, txData)
	bf := new(bytes.Buffer)
	rlp.Encode(bf, tx)

	rawtx := hexutil.Encode(bf.Bytes())
	unsignedTx, err := eth.DeserializeTx(rawtx)
	if err != nil {
		return fmt.Errorf("SendKaiCrossEth, eth.DeserializeTx error: %s", err.Error())
	}
	signedtx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, ctx.KaiInvoker.PrivateKey)
	if err != nil {
		return fmt.Errorf("SendKaiCrossEth, types.SignTx error: %s", err.Error())
	}

	err = kclient.SendTransaction(context.Background(), signedtx)
	if err != nil {
		return fmt.Errorf("SendKaiCrossEth, send transaction error:%s", err.Error())
	}
	status.AddTx(signedtx.Hash().String()[2:], &testframework.TxInfo{"KaiToEth", time.Now()})
	ctx.KaiInvoker.ETHUtil.WaitTransactionConfirm(signedtx.Hash())
	return nil
}

func SendEthCrossKai(ctx *testframework.TestFrameworkContext, status *testframework.CaseStatus, amount uint64) error {
	client := ctx.EthInvoker.ETHUtil.GetEthClient()

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("SendEthCrossKai, get suggest gas price failed error: %s", err.Error())
	}
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(5))

	contractabi, err := abi.JSON(strings.NewReader(lock_proxy_abi.LockProxyABI))
	if err != nil {
		return fmt.Errorf("SendEthCrossKai, abi.JSON error:" + err.Error())
	}

	assetaddress := ethcommon.HexToAddress("0000000000000000000000000000000000000000")
	txData, err := contractabi.Pack("lock", assetaddress, uint64(config.DefConfig.KaiChainID), ctx.OntInvoker.OntAcc.Address[:],
		big.NewInt(int64(amount)))
	if err != nil {
		return fmt.Errorf("SendEthCrossKai, contractabi.Pack error:" + err.Error())
	}

	contractAddr := ethcommon.HexToAddress(config.DefConfig.KaiLockProxy)
	callMsg := ethereum.CallMsg{
		From: ctx.KaiInvoker.EthTestSigner.Address, To: &contractAddr, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(int64(amount)), Data: txData,
	}
	gasLimit, err := client.EstimateGas(context.Background(), callMsg)
	if err != nil {
		return fmt.Errorf("SendEthCrossKai, estimate gas limit error: %s", err.Error())
	}

	nonce := ctx.KaiInvoker.NM.GetAddressNonce(ctx.KaiInvoker.EthTestSigner.Address)
	tx := types.NewTransaction(nonce, contractAddr, big.NewInt(int64(amount)), gasLimit, gasPrice, txData)
	bf := new(bytes.Buffer)
	_ = rlp.Encode(bf, tx)

	rawtx := hexutil.Encode(bf.Bytes())
	unsignedTx, err := eth.DeserializeTx(rawtx)
	if err != nil {
		return fmt.Errorf("SendEthCrossKai, eth.DeserializeTx error: %s", err.Error())
	}
	signedtx, err := types.SignTx(unsignedTx, types.HomesteadSigner{}, ctx.KaiInvoker.EthTestSigner.PrivateKey)
	if err != nil {
		return fmt.Errorf("SendEthCrossKai, types.SignTx error: %s", err.Error())
	}

	err = client.SendTransaction(context.Background(), signedtx)
	if err != nil {
		return fmt.Errorf("SendEthCrossKai, send transaction error:%s", err.Error())
	}
	status.AddTx(signedtx.Hash().String()[2:], &testframework.TxInfo{"EthToKai", time.Now()})
	ctx.EthInvoker.ETHUtil.WaitTransactionConfirm(signedtx.Hash())
	return nil
}

func MakeKaiAuth(signer *eth.EthSigner, nonce, gasPrice, gasLimit uint64) *bind.TransactOpts {
	auth := bind.NewKeyedTransactor(signer.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(int64(0)) // in wei
	auth.GasLimit = gasLimit          // in units
	auth.GasPrice = big.NewInt(int64(gasPrice))

	return auth
}
