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
package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/btcsuite/btcd/wire"
	types3 "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	common3 "github.com/ethereum/go-ethereum/common"
	"github.com/joeqian10/neo-gogogo/block"
	"github.com/joeqian10/neo-gogogo/helper/io"
	"github.com/joeqian10/neo-gogogo/rpc"
	"github.com/ontio/ontology-crypto/keypair"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
	common2 "github.com/ontio/ontology/common"
	"github.com/ontio/ontology/core/types"
	"github.com/ontio/ontology/smartcontract/service/native/cross_chain/header_sync"
	"github.com/ontio/ontology/smartcontract/service/native/governance"
	utils2 "github.com/ontio/ontology/smartcontract/service/native/utils"

	"github.com/polynetwork/eth-contracts/go_abi/eccm_abi"
	poly_go_sdk "github.com/polynetwork/poly-go-sdk"

	"github.com/polynetwork/poly-io-test/chains/btc"
	cosmos2 "github.com/polynetwork/poly-io-test/chains/cosmos"
	"github.com/polynetwork/poly-io-test/chains/eth"
	"github.com/polynetwork/poly-io-test/chains/ont"
	"github.com/polynetwork/poly-io-test/config"
	"github.com/polynetwork/poly-io-test/log"
	"github.com/polynetwork/poly-io-test/testcase"
	"github.com/polynetwork/poly/common"
	vconfig "github.com/polynetwork/poly/consensus/vbft/config"
	"github.com/polynetwork/poly/native/service/governance/node_manager"
	"github.com/polynetwork/poly/native/service/governance/relayer_manager"
	"github.com/polynetwork/poly/native/service/governance/side_chain_manager"
	"github.com/polynetwork/poly/native/service/header_sync/bsc"
	"github.com/polynetwork/poly/native/service/header_sync/cosmos"
	"github.com/polynetwork/poly/native/service/utils"
	"github.com/tendermint/tendermint/rpc/client/http"
	types2 "github.com/tendermint/tendermint/types"
)

var (
	tool                                                                  string
	toolConf                                                              string
	pWalletFiles                                                          string
	pPwds                                                                 string
	oWalletFiles                                                          string
	oPwds                                                                 string
	newWallet                                                             string
	newPwd                                                                string
	amt                                                                   int64
	keyFile                                                               string
	stateFile                                                             string
	id                                                                    uint64
	blockMsgDelay, hashMsgDelay, peerHandshakeTimeout, maxBlockChangeView uint64
	rootca                                                                string
	chainId                                                               uint64
	fabricRelayerTy                                                       uint64
)

func init() {
	flag.StringVar(&tool, "tool", "", "choose a tool to run")
	flag.StringVar(&toolConf, "conf", "./config.json", "configuration file path")
	flag.StringVar(&pWalletFiles, "pwallets", "", "poly wallet files sep by ','")
	flag.StringVar(&pPwds, "ppwds", "", "poly pwd for every wallet, sep by ','")
	flag.StringVar(&oWalletFiles, "owallets", "", "ontology wallet files sep by ','")
	flag.StringVar(&oPwds, "opwds", "", "ontology pwd for every wallet, sep by ','")
	flag.StringVar(&newWallet, "newwallet", "", "new wallet adding to poly consensus")
	flag.StringVar(&newPwd, "newpwd", "", "password for new wallet")
	flag.Int64Var(&amt, "amt", 50, "amount to create new cosmos validator")
	flag.StringVar(&keyFile, "cosmos_val_privk_file", "", "cosmos validator's privk file")
	flag.StringVar(&stateFile, "cosmos_val_state_file", "", "cosmos validator's state file")
	flag.Uint64Var(&id, "id", 0, "chain id to quit")
	flag.Uint64Var(&blockMsgDelay, "blk_msg_delay", 5000, "")
	flag.Uint64Var(&hashMsgDelay, "hash_msg_delay", 5000, "")
	flag.Uint64Var(&peerHandshakeTimeout, "peer_handshake_timeout", 10, "")
	flag.Uint64Var(&maxBlockChangeView, "max_blk_change_view", 10000, "")
	flag.StringVar(&rootca, "rootca", "", "file path for root CA")
	flag.Uint64Var(&chainId, "chainid", 0, "default 0 means all chains")
	flag.Uint64Var(&fabricRelayerTy, "fab_relayer_type", 1, "the relayer of fabric type: how many orgs need to sign CA for relayer")

	flag.Parse()
}

func main() {
	log.InitLog(2, os.Stdout)

	err := config.DefConfig.Init(toolConf)
	if err != nil {
		panic(err)
	}
	poly := poly_go_sdk.NewPolySdk()
	if err := btc.SetUpPoly(poly, config.DefConfig.RchainJsonRpcAddress); err != nil {
		panic(err)
	}

	acc, err := btc.GetAccountByPassword(poly, config.DefConfig.RCWallet, []byte(config.DefConfig.RCWalletPwd))
	if err != nil {
		panic(err)
	}

	switch tool {
	case "register_side_chain":
		wArr := strings.Split(pWalletFiles, ",")
		pArr := strings.Split(pPwds, ",")

		accArr := make([]*poly_go_sdk.Account, len(wArr))
		for i, v := range wArr {
			accArr[i], err = btc.GetAccountByPassword(poly, v, []byte(pArr[i]))
			if err != nil {
				panic(fmt.Errorf("failed to decode no%d wallet %s with pwd %s", i, wArr[i], pArr[i]))
			}
		}
		switch chainId {
		case config.DefConfig.BtcChainID:
			if RegisterBtcChain(poly, acc) {
				ApproveRegisterSideChain(config.DefConfig.BtcChainID, poly, accArr)
			}
		case config.DefConfig.EthChainID:
			if RegisterEthChain(poly, acc) {
				ApproveRegisterSideChain(config.DefConfig.EthChainID, poly, accArr)
			}
		case config.DefConfig.OntChainID:
			if RegisterOntChain(poly, acc) {
				ApproveRegisterSideChain(config.DefConfig.OntChainID, poly, accArr)
			}
		case config.DefConfig.NeoChainID:
			if RegisterNeoChain(poly, acc) {
				ApproveRegisterSideChain(config.DefConfig.NeoChainID, poly, accArr)
			}
		case config.DefConfig.CMCrossChainId:
			if RegisterCosmos(poly, acc) {
				ApproveRegisterSideChain(config.DefConfig.CMCrossChainId, poly, accArr)
			}
		case config.DefConfig.BscChainID:
			if RegisterBSC(poly, acc) {
				ApproveRegisterSideChain(config.DefConfig.BscChainID, poly, accArr)
			}
		case 0:
			if RegisterBtcChain(poly, acc) {
				ApproveRegisterSideChain(config.DefConfig.BtcChainID, poly, accArr)
			}
			if RegisterOntChain(poly, acc) {
				ApproveRegisterSideChain(config.DefConfig.OntChainID, poly, accArr)
			}
			if RegisterEthChain(poly, acc) {
				ApproveRegisterSideChain(config.DefConfig.EthChainID, poly, accArr)
			}
			if RegisterCosmos(poly, acc) {
				ApproveRegisterSideChain(config.DefConfig.CMCrossChainId, poly, accArr)
			}
			if RegisterNeoChain(poly, acc) {
				ApproveRegisterSideChain(config.DefConfig.NeoChainID, poly, accArr)
			}
			if RegisterBSC(poly, acc) {
				ApproveRegisterSideChain(config.DefConfig.BscChainID, poly, accArr)
			}
		}

	case "sync_genesis_header":
		wArr := strings.Split(pWalletFiles, ",")
		pArr := strings.Split(pPwds, ",")

		accArr := make([]*poly_go_sdk.Account, len(wArr))
		for i, v := range wArr {
			accArr[i], err = btc.GetAccountByPassword(poly, v, []byte(pArr[i]))
			if err != nil {
				panic(fmt.Errorf("failed to decode no%d wallet %s with pwd %s", i, wArr[i], pArr[i]))
			}
		}

		switch chainId {
		case config.DefConfig.BtcChainID:
			SyncBtcGenesisHeader(poly, acc)
		case config.DefConfig.EthChainID:
			SyncEthGenesisHeader(poly, accArr)
		case config.DefConfig.OntChainID:
			SyncOntGenesisHeader(poly, accArr)
		case config.DefConfig.NeoChainID:
			SyncNeoGenesisHeader(poly, accArr)
		case config.DefConfig.CMCrossChainId:
			SyncCosmosGenesisHeader(poly, accArr)
		case config.DefConfig.BscChainID:
			SyncBSCGenesisHeader(poly, accArr)
		case 0:
			SyncBtcGenesisHeader(poly, acc)
			SyncEthGenesisHeader(poly, accArr)
			SyncOntGenesisHeader(poly, accArr)
			SyncCosmosGenesisHeader(poly, accArr)
			SyncNeoGenesisHeader(poly, accArr)
			SyncBSCGenesisHeader(poly, accArr)
		}

	case "update_btc":
		accArr := getPolyAccounts(poly)
		if UpdateBtc(poly, acc) {
			ApproveUpdateChain(config.DefConfig.BtcChainID, poly, accArr)
		}

	case "update_eth":
		accArr := getPolyAccounts(poly)
		if UpdateEth(poly, acc) {
			ApproveUpdateChain(config.DefConfig.EthChainID, poly, accArr)
		}

	case "update_neo":
		accArr := getPolyAccounts(poly)
		if UpdateNeo(poly, acc) {
			ApproveUpdateChain(config.DefConfig.NeoChainID, poly, accArr)
		}

	case "init_ont_acc":
		err := InitOntAcc()
		if err != nil {
			panic(err)
		}

	case "poly_add_node":
		accArr := getPolyAccounts(poly)
		acc, err = btc.GetAccountByPassword(poly, newWallet, []byte(newPwd))
		if err != nil {
			panic(fmt.Errorf("failed to get new account: %v", err))
		}
		if RegisterCandidate(poly, acc) {
			ApproveCandidate(acc, poly, accArr)
		}
		CommitPolyDpos(poly, accArr)

	case "cosmos_create_validator":
		CosmosCreateValidator(keyFile, stateFile, amt)
	case "cosmos_delegate_validator":
		CosmosDelegateToVal(amt)
	case "black_poly_node":
		accArr := getPolyAccounts(poly)
		acc, err = btc.GetAccountByPassword(poly, newWallet, []byte(newPwd))
		if err != nil {
			panic(fmt.Errorf("failed to get new account: %v", err))
		}
		BlackPolyNode(poly, acc, accArr)
	case "white_poly_node":
		accArr := getPolyAccounts(poly)
		acc, err = btc.GetAccountByPassword(poly, newWallet, []byte(newPwd))
		if err != nil {
			panic(fmt.Errorf("failed to get new account: %v", err))
		}
		WhitePolyNode(poly, acc, accArr)
	case "quit_poly_node":
		accArr := getPolyAccounts(poly)
		acc, err = btc.GetAccountByPassword(poly, newWallet, []byte(newPwd))
		if err != nil {
			panic(fmt.Errorf("failed to get new account: %v", err))
		}
		QuitNode(poly, acc)
		CommitPolyDpos(poly, accArr)
	case "get_poly_consensus":
		GetPolyConsensusInfo(poly)
	case "add_relayer":
		accArr := getPolyAccounts(poly)
		signer := acc
		acc = &poly_go_sdk.Account{}
		if newPwd != "" {
			acc, err = btc.GetAccountByPassword(poly, newWallet, []byte(newPwd))
			if err != nil {
				panic(fmt.Errorf("failed to get new account: %v", err))
			}
		} else {
			acc.Address, _ = common.AddressFromBase58(newWallet)
		}
		ApproveRelayer(poly, RegisterRelayer(poly, acc, signer), accArr)
	case "reg_poly_node":
		acc, err = btc.GetAccountByPassword(poly, newWallet, []byte(newPwd))
		if err != nil {
			panic(fmt.Errorf("failed to get new account: %v", err))
		}
		RegisterCandidate(poly, acc)
	case "unreg_poly_node":
		acc, err = btc.GetAccountByPassword(poly, newWallet, []byte(newPwd))
		if err != nil {
			panic(fmt.Errorf("failed to get new account: %v", err))
		}
		UnRegisterPolyCandidate(poly, acc)

	case "remove_relayer":
		accArr := getPolyAccounts(poly)
		acc, err = btc.GetAccountByPassword(poly, newWallet, []byte(newPwd))
		if err != nil {
			panic(fmt.Errorf("failed to get new account: %v", err))
		}
		ApproveRemoveRelayer(poly, RemoveRelayer(poly, acc), accArr)
	case "get_relayer":
		acc, err = btc.GetAccountByPassword(poly, newWallet, []byte(newPwd))
		if err != nil {
			panic(fmt.Errorf("failed to get new account: %v", err))
		}
		GetRelayer(poly, acc)
	case "quit_side_chain":
		accArr := getPolyAccounts(poly)
		QuitSideChain(poly, id, acc)
		ApproveQuitSideChain(poly, id, accArr)
	case "get_side_chain":
		GetSideChain(poly, id)
	case "get_poly_config":
		GetPolyConfig(poly)
	case "update_poly_config":
		accArr := getPolyAccounts(poly)
		UpdatePolyConfig(poly, uint32(blockMsgDelay), uint32(hashMsgDelay), uint32(peerHandshakeTimeout),
			uint32(maxBlockChangeView), accArr)
	case "commit_poly_dpos":
		accArr := getPolyAccounts(poly)
		CommitPolyDpos(poly, accArr)
	case "commit_ont_dpos":
		CommitOntDpos()
	}
}

func getPolyAccounts(poly *poly_go_sdk.PolySdk) []*poly_go_sdk.Account {
	wArr := strings.Split(pWalletFiles, ",")
	pArr := strings.Split(pPwds, ",")
	accArr := make([]*poly_go_sdk.Account, len(wArr))
	var err error
	for i, v := range wArr {
		accArr[i], err = btc.GetAccountByPassword(poly, v, []byte(pArr[i]))
		if err != nil {
			panic(fmt.Errorf("failed to decode no%d wallet %s with pwd %s", i, wArr[i], pArr[i]))
		}
	}
	return accArr
}

func InitOntAcc() error {
	oi, err := ont.NewOntInvoker(config.DefConfig.OntJsonRpcAddress, config.DefConfig.OntContractsAvmPath,
		config.DefConfig.OntWallet, config.DefConfig.OntWalletPassword)
	if err != nil {
		return fmt.Errorf("failed to new ont invoker: %v", err)
	}

	ow := strings.Split(oWalletFiles, ",")
	op := strings.Split(oPwds, ",")
	oAccArr := make([]*ontology_go_sdk.Account, len(ow))
	pks := make([]keypair.PublicKey, len(ow))
	for i, v := range ow {
		oAccArr[i], err = ont.GetOntAccByPwd(v, op[i])
		if err != nil {
			return fmt.Errorf("failed to decode no%d wallet %s with pwd %s", i, ow[i], op[i])
		}
		pks[i] = oAccArr[i].PublicKey
	}

	multiSignAddr, _ := types.AddressFromBookkeepers(pks)
	tx, err := WithdrawONTFromConsensus(oi.OntSdk, 0, 20000000, oi.OntAcc, oAccArr, pks, multiSignAddr,
		oi.OntAcc.Address, 100000000)
	if err != nil {
		return err
	}
	oi.WaitTxConfirmation(tx)

	txHash, err := oi.OntSdk.Native.Ont.Transfer(0, 20000000, oi.OntAcc,
		oi.OntAcc, oi.OntAcc.Address, 100000000)
	if err != nil {
		return err
	}
	oi.WaitTxConfirmation(txHash)
	log.Infof("InitOntAcc, transfer ont to myself %s", oi.OntAcc.Address.ToBase58())

	for _, oa := range oAccArr {
		amount, err := oi.OntSdk.Native.Ong.BalanceOf(oa.Address)
		if err != nil {
			return err
		}
		tx, err := oi.OntSdk.Native.Ong.Transfer(0, 20000000, oa, oa, oi.OntAcc.Address, amount)
		if err != nil {
			return err
		}
		oi.WaitTxConfirmation(tx)
		log.Infof("InitOntAcc, Withdraw %d ONG to myself %s", amount, oi.OntAcc.Address.ToBase58())
	}

	return nil
}

func WithdrawONTFromConsensus(ontSdk *ontology_go_sdk.OntologySdk, gasPrice, gasLimit uint64,
	payer *ontology_go_sdk.Account, signers []*ontology_go_sdk.Account, pks []keypair.PublicKey, from common2.Address,
	to common2.Address, amount uint64) (common2.Uint256, error) {
	tx, err := ontSdk.Native.Ont.NewTransferTransaction(gasPrice, gasLimit, from, to, amount)
	if err != nil {
		return common2.UINT256_EMPTY, err
	}
	if payer != nil {
		ontSdk.SetPayer(tx, payer.Address)
		err = ontSdk.SignToTransaction(tx, payer)
		if err != nil {
			return common2.UINT256_EMPTY, err
		}
	}
	for _, singer := range signers {
		err = ontSdk.MultiSignToTransaction(tx, uint16((5*len(pks)+6)/7), pks, singer)
		if err != nil {
			return common2.UINT256_EMPTY, err
		}
	}

	return ontSdk.SendTransaction(tx)
}

func SyncBtcGenesisHeader(poly *poly_go_sdk.PolySdk, acc *poly_go_sdk.Account) {
	cli := btc.NewRestCli(config.DefConfig.BtcRestAddr, config.DefConfig.BtcRestUser, config.DefConfig.BtcRestPwd)
	curr, _, err := cli.GetCurrentHeightAndHash()
	if err != nil {
		panic(fmt.Errorf("SyncBtcGenesisHeader failed: %v", err))
	}
	start := curr - curr%2016
	hdr, err := cli.GetHeader(start)
	if err != nil {
		panic(fmt.Errorf("SyncBtcGenesisHeader falied: %v", err))
	}
	var buf bytes.Buffer
	err = hdr.BtcEncode(&buf, wire.ProtocolVersion, wire.LatestEncoding)
	if err != nil {
		panic(err)
	}

	hb := make([]byte, 4)
	binary.BigEndian.PutUint32(hb, uint32(start))
	txhash, err := poly.Native.Hs.SyncGenesisHeader(config.DefConfig.BtcChainID, append(buf.Bytes(), hb...),
		[]*poly_go_sdk.Account{acc})
	if err != nil {
		if strings.Contains(err.Error(), "had been initialized") {
			log.Info("btc already synced")
		} else {
			panic(fmt.Errorf("SyncBtcGenesisHeader failed: %v", err))
		}
	} else {
		testcase.WaitPolyTx(txhash, poly)
		blkHash := hdr.BlockHash()
		log.Infof("successful to sync btc genesis header: ( height: %d, block_hash: %s, txhash: %s )", start,
			blkHash.String(), txhash.ToHexString())
	}
}

func SyncEthGenesisHeader(poly *poly_go_sdk.PolySdk, accArr []*poly_go_sdk.Account) {
	tool := eth.NewEthTools(config.DefConfig.EthURL)
	curr, err := tool.GetNodeHeight()
	if err != nil {
		panic(err)
	}
	hdr, err := tool.GetBlockHeader(curr)
	if err != nil {
		panic(err)
	}
	raw, err := hdr.MarshalJSON()
	if err != nil {
		panic(err)
	}
	txhash, err := poly.Native.Hs.SyncGenesisHeader(config.DefConfig.EthChainID, raw, accArr)
	if err != nil {
		if strings.Contains(err.Error(), "had been initialized") {
			log.Info("eth already synced")
		} else {
			panic(fmt.Errorf("SyncEthGenesisHeader failed: %v", err))
		}
	} else {
		testcase.WaitPolyTx(txhash, poly)
		log.Infof("successful to sync eth genesis header: (height: %d, blk_hash: %s, txhash: %s )", curr,
			hdr.Hash().String(), txhash.ToHexString())
	}

	eccmContract, err := eccm_abi.NewEthCrossChainManager(common3.HexToAddress(config.DefConfig.Eccm), tool.GetEthClient())
	if err != nil {
		panic(err)
	}
	signer, err := eth.NewEthSigner(config.DefConfig.ETHPrivateKey)
	if err != nil {
		panic(err)
	}
	nonce := eth.NewNonceManager(tool.GetEthClient()).GetAddressNonce(signer.Address)
	gasPrice, err := tool.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		panic(fmt.Errorf("SyncEthGenesisHeader, get suggest gas price failed error: %s", err.Error()))
	}
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(5))

	gB, err := poly.GetBlockByHeight(config.DefConfig.RCEpoch)
	if err != nil {
		panic(err)
	}
	info := &vconfig.VbftBlockInfo{}
	if err := json.Unmarshal(gB.Header.ConsensusPayload, info); err != nil {
		panic(fmt.Errorf("SyncEthGenesisHeader - unmarshal blockInfo error: %s", err))
	}

	var bookkeepers []keypair.PublicKey
	for _, peer := range info.NewChainConfig.Peers {
		keystr, _ := hex.DecodeString(peer.ID)
		key, _ := keypair.DeserializePublicKey(keystr)
		bookkeepers = append(bookkeepers, key)
	}
	bookkeepers = keypair.SortPublicKeys(bookkeepers)

	publickeys := make([]byte, 0)
	for _, key := range bookkeepers {
		publickeys = append(publickeys, ont.GetOntNoCompressKey(key)...)
	}

	rawHdr := gB.Header.ToArray()

	contractabi, err := abi.JSON(strings.NewReader(eccm_abi.EthCrossChainManagerABI))
	if err != nil {
		log.Errorf("SyncEthGenesisHeader, abi.JSON error: %v", err)
		return
	}
	txData, err := contractabi.Pack("initGenesisBlock", rawHdr, publickeys)
	if err != nil {
		log.Errorf("SyncEthGenesisHeader, contractabi.Pack error: %v", err)
		return
	}

	eccm := common3.HexToAddress(config.DefConfig.Eccm)
	callMsg := ethereum.CallMsg{
		From: signer.Address, To: &eccm, Gas: 0, GasPrice: gasPrice,
		Value: big.NewInt(0), Data: txData,
	}
	gasLimit, err := tool.GetEthClient().EstimateGas(context.Background(), callMsg)
	if err != nil {
		log.Errorf("SyncEthGenesisHeader, estimate gas limit error: %s", err.Error())
		return
	}

	auth := testcase.MakeEthAuth(signer, nonce, gasPrice.Uint64(), gasLimit)
	tx, err := eccmContract.InitGenesisBlock(auth, rawHdr, publickeys)
	if err != nil {
		log.Errorf("SyncEthGenesisHeader, failed to sync poly header to ETH: %v", err)
		return
	}
	tool.WaitTransactionConfirm(tx.Hash())
	log.Infof("successful to sync poly genesis header to Ethereum: ( txhash: %s )", tx.Hash().String())
}

func SyncBSCGenesisHeader(poly *poly_go_sdk.PolySdk, accArr []*poly_go_sdk.Account) {
	tool := eth.NewEthTools(config.DefConfig.BSCURL)
	height, err := tool.GetNodeHeight()
	if err != nil {
		panic(err)
	}

	epochHeight := height - height%200
	pEpochHeight := epochHeight - 200

	hdr, err := tool.GetBlockHeader(epochHeight)
	if err != nil {
		panic(err)
	}
	phdr, err := tool.GetBlockHeader(pEpochHeight)
	if err != nil {
		panic(err)
	}
	pvalidators, err := bsc.ParseValidators(phdr.Extra[32 : len(phdr.Extra)-65])
	if err != nil {
		panic(err)
	}

	if len(hdr.Extra) <= 65+32 {
		panic(fmt.Sprintf("invalid epoch header at height:%d", epochHeight))
	}
	if len(phdr.Extra) <= 65+32 {
		panic(fmt.Sprintf("invalid epoch header at height:%d", pEpochHeight))
	}

	genesisHeader := bsc.GenesisHeader{Header: *hdr, PrevValidators: []bsc.HeightAndValidators{
		{Height: big.NewInt(int64(pEpochHeight)), Validators: pvalidators},
	}}
	raw, err := json.Marshal(genesisHeader)
	if err != nil {
		panic(err)
	}
	txhash, err := poly.Native.Hs.SyncGenesisHeader(config.DefConfig.BscChainID, raw, accArr)
	if err != nil {
		if strings.Contains(err.Error(), "had been initialized") {
			log.Info("bsc already synced")
		} else {
			panic(fmt.Errorf("SyncBSCGenesisHeader failed: %v", err))
		}
	} else {
		testcase.WaitPolyTx(txhash, poly)
		log.Infof("successful to sync bsc genesis header: (height: %d, blk_hash: %s, txhash: %s )", epochHeight,
			hdr.Hash().String(), txhash.ToHexString())
	}

	eccmContract, err := eccm_abi.NewEthCrossChainManager(common3.HexToAddress(config.DefConfig.BscEccm), tool.GetEthClient())
	if err != nil {
		panic(err)
	}
	signer, err := eth.NewEthSigner(config.DefConfig.BSCPrivateKey)
	if err != nil {
		panic(err)
	}
	nonce := eth.NewNonceManager(tool.GetEthClient()).GetAddressNonce(signer.Address)
	gasPrice, err := tool.GetEthClient().SuggestGasPrice(context.Background())
	if err != nil {
		panic(fmt.Errorf("SyncBSCGenesisHeader, get suggest gas price failed error: %s", err.Error()))
	}
	gasPrice = gasPrice.Mul(gasPrice, big.NewInt(5))
	auth := testcase.MakeEthAuth(signer, nonce, gasPrice.Uint64(), uint64(8000000))

	gB, err := poly.GetBlockByHeight(config.DefConfig.RCEpoch)
	if err != nil {
		panic(err)
	}
	info := &vconfig.VbftBlockInfo{}
	if err := json.Unmarshal(gB.Header.ConsensusPayload, info); err != nil {
		panic(fmt.Errorf("commitGenesisHeader - unmarshal blockInfo error: %s", err))
	}

	var bookkeepers []keypair.PublicKey
	for _, peer := range info.NewChainConfig.Peers {
		keystr, _ := hex.DecodeString(peer.ID)
		key, _ := keypair.DeserializePublicKey(keystr)
		bookkeepers = append(bookkeepers, key)
	}
	bookkeepers = keypair.SortPublicKeys(bookkeepers)

	publickeys := make([]byte, 0)
	for _, key := range bookkeepers {
		publickeys = append(publickeys, ont.GetOntNoCompressKey(key)...)
	}

	tx, err := eccmContract.InitGenesisBlock(auth, gB.Header.ToArray(), publickeys)
	tool.WaitTransactionConfirm(tx.Hash())
	log.Infof("successful to sync poly genesis header to BSC: ( txhash: %s )", tx.Hash().String())
}

func SyncOntGenesisHeader(poly *poly_go_sdk.PolySdk, accArr []*poly_go_sdk.Account) {
	ontCli := ontology_go_sdk.NewOntologySdk()
	ontCli.NewRpcClient().SetAddress(config.DefConfig.OntJsonRpcAddress)

	genesisBlock, err := ontCli.GetBlockByHeight(config.DefConfig.OntEpoch)
	if err != nil {
		panic(err)
	}
	txhash, err := poly.Native.Hs.SyncGenesisHeader(config.DefConfig.OntChainID, genesisBlock.Header.ToArray(), accArr)
	if err != nil {
		if strings.Contains(err.Error(), "had been initialized") {
			log.Info("ont already synced")
		} else {
			panic(fmt.Errorf("SyncOntGenesisHeader failed: %v", err))
		}
	} else {
		testcase.WaitPolyTx(txhash, poly)
		log.Infof("successful to sync ont genesis header: ( txhash: %s )", txhash.ToHexString())
	}
	ow := strings.Split(oWalletFiles, ",")
	op := strings.Split(oPwds, ",")
	oAccArr := make([]*ontology_go_sdk.Account, len(ow))
	pks := make([]keypair.PublicKey, len(ow))
	for i, v := range ow {
		oAccArr[i], err = ont.GetOntAccByPwd(v, op[i])
		if err != nil {
			panic(fmt.Errorf("failed to decode no%d wallet %s with pwd %s", i, ow[i], op[i]))
		}
		pks[i] = oAccArr[i].PublicKey
	}
	gB, err := poly.GetBlockByHeight(config.DefConfig.RCEpoch)
	if err != nil {
		panic(err)
	}
	txHash, err := InvokeNativeContractWithMultiSign(ontCli, 0, 2000000, pks, oAccArr, byte(0),
		utils2.HeaderSyncContractAddress, header_sync.SYNC_GENESIS_HEADER,
		[]interface{}{
			&header_sync.SyncGenesisHeaderParam{
				GenesisHeader: gB.Header.ToArray(),
			}})
	if err != nil {
		panic(fmt.Errorf("faild to sync poly header to ontology: %v", err))
	}
	ont.WaitOntTx(txHash, ontCli)
	log.Infof("successful to sync poly genesis header to Ontology: ( txhash: %s )", txHash.ToHexString())
}

func SyncNeoGenesisHeader(poly *poly_go_sdk.PolySdk, accArr []*poly_go_sdk.Account) error {
	cli := rpc.NewClient(config.DefConfig.NeoUrl)
	resp := cli.GetBlockHeaderByIndex(config.DefConfig.NeoEpoch)
	if resp.HasError() {
		return fmt.Errorf("failed to get header: %v", resp.Error.Message)
	}
	header, err := block.NewBlockHeaderFromRPC(&resp.Result)
	if err != nil {
		return err
	}
	buf := io.NewBufBinaryWriter()
	header.Serialize(buf.BinaryWriter)
	if buf.Err != nil {
		return buf.Err
	}

	txhash, err := poly.Native.Hs.SyncGenesisHeader(config.DefConfig.NeoChainID, buf.Bytes(), accArr)
	if err != nil {
		if strings.Contains(err.Error(), "had been initialized") {
			log.Info("neo already synced")
		} else {
			panic(fmt.Errorf("SyncNeoGenesisHeader failed: %v", err))
		}
	} else {
		testcase.WaitPolyTx(txhash, poly)
		log.Infof("successful to sync neo genesis header: ( txhash: %s )", txhash.ToHexString())
	}

	return nil
}

func SyncCosmosGenesisHeader(poly *poly_go_sdk.PolySdk, accArr []*poly_go_sdk.Account) {
	invoker, err := cosmos2.NewCosmosInvoker()
	if err != nil {
		panic(err)
	}

	res, err := invoker.RpcCli.Commit(&config.DefConfig.CMEpoch)
	if err != nil {
		panic(err)
	}
	vals, err := getValidators(invoker.RpcCli, config.DefConfig.CMEpoch)
	if err != nil {
		panic(err)
	}
	ch := &cosmos.CosmosHeader{
		Header:  *res.Header,
		Commit:  res.Commit,
		Valsets: vals,
	}
	raw, err := invoker.CMCdc.MarshalBinaryBare(ch)
	if err != nil {
		panic(err)
	}
	txhash, err := poly.Native.Hs.SyncGenesisHeader(config.DefConfig.CMCrossChainId, raw, accArr)
	if err != nil {
		if strings.Contains(err.Error(), "had been initialized") {
			log.Info("cosmos already synced")
		} else {
			panic(err)
		}
	} else {
		testcase.WaitPolyTx(txhash, poly)
		log.Infof("successful to sync cosmos genesis header: ( txhash: %s )", txhash.ToHexString())
	}

	header, err := poly.GetHeaderByHeight(config.DefConfig.RCEpoch)
	if err != nil {
		panic(err)
	}

	tx, err := invoker.SyncPolyGenesisHdr(invoker.Acc.Acc, header.ToArray())
	if err != nil {
		panic(err)
	}

	log.Infof("successful to sync poly genesis header to cosmos: ( txhash: %s )", tx.Hash.String())
}

func getValidators(rpc *http.HTTP, h int64) ([]*types2.Validator, error) {
	p := 1
	vSet := make([]*types2.Validator, 0)
	for {
		res, err := rpc.Validators(&h, p, 100)
		if err != nil {
			if strings.Contains(err.Error(), "page should be within") {
				return vSet, nil
			}
			return nil, err
		}
		// In case tendermint don't give relayer the right error
		if len(res.Validators) == 0 {
			return vSet, nil
		}
		vSet = append(vSet, res.Validators...)
		p++
	}
}

func InvokeNativeContractWithMultiSign(
	sdk *ontology_go_sdk.OntologySdk,
	gasPrice,
	gasLimit uint64,
	pubKeys []keypair.PublicKey,
	singers []*ontology_go_sdk.Account,
	cversion byte,
	contractAddress common2.Address,
	method string,
	params []interface{},
) (common2.Uint256, error) {
	tx, err := sdk.Native.NewNativeInvokeTransaction(gasPrice, gasLimit, cversion, contractAddress, method, params)
	if err != nil {
		return common2.UINT256_EMPTY, err
	}
	for _, singer := range singers {
		err = sdk.MultiSignToTransaction(tx, uint16((5*len(pubKeys)+6)/7), pubKeys, singer)
		if err != nil {
			return common2.UINT256_EMPTY, err
		}
	}
	return sdk.SendTransaction(tx)
}

func ApproveRegisterSideChain(id uint64, poly *poly_go_sdk.PolySdk, accArr []*poly_go_sdk.Account) {
	var (
		txhash common.Uint256
		err    error
	)

	for i, a := range accArr {
		txhash, err = poly.Native.Scm.ApproveRegisterSideChain(id, a)
		if err != nil {
			panic(fmt.Errorf("No%d ApproveRegisterSideChain failed: %v", i, err))
		}
		log.Infof("No%d: successful to approve: ( acc: %s, txhash: %s, chain-id: %d )", i, a.Address.ToBase58(), txhash.ToHexString(), id)
	}
	testcase.WaitPolyTx(txhash, poly)
}

func RegisterBtcChain(poly *poly_go_sdk.PolySdk, acc *poly_go_sdk.Account) bool {
	blkToWait := uint64(1)
	var tyNet utils.BtcNetType
	switch config.BtcNet.Name {
	case "testnet3":
		blkToWait = 6
		tyNet = utils.TyTestnet3
	case "mainnet":
		blkToWait = 6
		tyNet = utils.TyMainnet
	case "regtest":
		tyNet = utils.TyRegtest
	case "simnet":
		tyNet = utils.TySimnet
	}

	rawTy := make([]byte, 8)
	binary.LittleEndian.PutUint64(rawTy, uint64(tyNet))
	txhash, err := poly.Native.Scm.RegisterSideChain(acc.Address, config.DefConfig.BtcChainID, config.DefConfig.BtcChainID, "btc",
		blkToWait, rawTy, acc)
	if err != nil {
		if strings.Contains(err.Error(), "already registered") {
			log.Infof("btc chain %d already registered", config.DefConfig.BtcChainID)
			return false
		}
		if strings.Contains(err.Error(), "already requested") {
			log.Infof("btc chain %d already requested", config.DefConfig.BtcChainID)
			return true
		}
		panic(fmt.Errorf("RegisterBtcChain failed: %v", err))
	}

	testcase.WaitPolyTx(txhash, poly)
	log.Infof("successful to register btc chain: ( txhash: %s )", txhash.ToHexString())

	return true
}

func RegisterEthChain(poly *poly_go_sdk.PolySdk, acc *poly_go_sdk.Account) bool {
	blkToWait := uint64(1)
	if config.BtcNet.Name == "testnet3" {
		blkToWait = 12
	}
	eccd, err := hex.DecodeString(strings.Replace(config.DefConfig.Eccd, "0x", "", 1))
	if err != nil {
		panic(fmt.Errorf("RegisterEthChain, failed to decode eccd '%s' : %v", config.DefConfig.Eccd, err))
	}
	txhash, err := poly.Native.Scm.RegisterSideChain(acc.Address, config.DefConfig.EthChainID, 2, "eth",
		blkToWait, eccd, acc)
	if err != nil {
		if strings.Contains(err.Error(), "already registered") {
			log.Infof("eth chain %d already registered", config.DefConfig.EthChainID)
			return false
		}
		if strings.Contains(err.Error(), "already requested") {
			log.Infof("eth chain %d already requested", config.DefConfig.EthChainID)
			return true
		}
		panic(fmt.Errorf("RegisterEthChain failed: %v", err))
	}
	testcase.WaitPolyTx(txhash, poly)
	log.Infof("successful to register eth chain: ( txhash: %s )", txhash.ToHexString())

	return true
}

func RegisterNeoChain(poly *poly_go_sdk.PolySdk, acc *poly_go_sdk.Account) bool {
	blkToWait := uint64(1)
	eccd, err := common2.AddressFromHexString(strings.TrimPrefix(config.DefConfig.NeoCCMC, "0x"))
	if err != nil {
		panic(fmt.Errorf("RegisterNeoChain, failed to decode eccd '%s' : %v", config.DefConfig.Eccd, err))
	}
	txhash, err := poly.Native.Scm.RegisterSideChain(acc.Address, config.DefConfig.NeoChainID, 4, "NEO",
		blkToWait, eccd[:], acc)
	if err != nil {
		if strings.Contains(err.Error(), "already registered") {
			log.Infof("neo chain %d already registered", config.DefConfig.NeoChainID)
			return false
		}
		if strings.Contains(err.Error(), "already requested") {
			log.Infof("neo chain %d already requested", config.DefConfig.NeoChainID)
			return true
		}
		panic(fmt.Errorf("RegisterNeoChain failed: %v", err))
	}
	testcase.WaitPolyTx(txhash, poly)
	log.Infof("successful to register neo chain: ( txhash: %s )", txhash.ToHexString())

	return true
}

func RegisterCosmos(poly *poly_go_sdk.PolySdk, acc *poly_go_sdk.Account) bool {
	blkToWait := uint64(1)
	txhash, err := poly.Native.Scm.RegisterSideChain(acc.Address, config.DefConfig.CMCrossChainId, 5, "switcheochain",
		blkToWait, []byte{}, acc)
	if err != nil {
		if strings.Contains(err.Error(), "already registered") {
			log.Infof("Cosmos chain %d already registered", config.DefConfig.CMCrossChainId)
			return false
		}
		if strings.Contains(err.Error(), "already requested") {
			log.Infof("Cosmos chain %d already requested", config.DefConfig.CMCrossChainId)
			return true
		}
		panic(fmt.Errorf("RegisterCosmosChain failed: %v", err))
	}

	testcase.WaitPolyTx(txhash, poly)
	log.Infof("successful to register cosmos chain: ( txhash: %s )", txhash.ToHexString())

	return true

}

func RegisterBSC(poly *poly_go_sdk.PolySdk, acc *poly_go_sdk.Account) bool {
	tool := eth.NewEthTools(config.DefConfig.BSCURL)
	chainID, err := tool.GetChainID()
	if err != nil {
		panic(err)
	}

	blkToWait := uint64(15)
	extra := bsc.ExtraInfo{
		ChainID: chainID,
	}

	extraBytes, _ := json.Marshal(extra)

	eccd, err := hex.DecodeString(strings.Replace(config.DefConfig.BscEccd, "0x", "", 1))
	if err != nil {
		panic(fmt.Errorf("RegisterBSC, failed to decode eccd '%s' : %v", config.DefConfig.BscEccd, err))
	}

	fmt.Println("config.DefConfig.BSCChainID", config.DefConfig.BscChainID, "extraBytes", string(extraBytes), "BscEccd", config.DefConfig.BscEccd)
	txhash, err := poly.Native.Scm.RegisterSideChainExt(acc.Address, config.DefConfig.BscChainID, 6, "bsc",
		blkToWait, eccd, extraBytes, acc)
	if err != nil {
		if strings.Contains(err.Error(), "already registered") {
			log.Infof("bsc chain %d already registered", config.DefConfig.BscChainID)
			return false
		}
		if strings.Contains(err.Error(), "already requested") {
			log.Infof("bsc chain %d already requested", config.DefConfig.BscChainID)
			return true
		}
		panic(fmt.Errorf("RegisterBSC failed: %v", err))
	}

	testcase.WaitPolyTx(txhash, poly)
	log.Infof("successful to register bsc chain: ( txhash: %s )", txhash.ToHexString())

	return true
}

func RegisterOntChain(poly *poly_go_sdk.PolySdk, acc *poly_go_sdk.Account) bool {
	txhash, err := poly.Native.Scm.RegisterSideChain(acc.Address, config.DefConfig.OntChainID, 3, "ont",
		1, []byte{}, acc)
	if err != nil {
		if strings.Contains(err.Error(), "already registered") {
			log.Infof("ont chain %d already registered", config.DefConfig.OntChainID)
			return false
		}
		if strings.Contains(err.Error(), "already requested") {
			log.Infof("ont chain %d already requested", config.DefConfig.OntChainID)
			return true
		}
		panic(fmt.Errorf("RegisterOntChain failed: %v", err))
	}

	testcase.WaitPolyTx(txhash, poly)
	log.Infof("successful to register ont chain: ( txhash: %s )", txhash.ToHexString())

	return true
}

func UpdateBtc(poly *poly_go_sdk.PolySdk, acc *poly_go_sdk.Account) bool {
	blkToWait := uint64(1)
	var tyNet utils.BtcNetType
	switch config.BtcNet.Name {
	case "testnet3":
		blkToWait = 6
		tyNet = utils.TyTestnet3
	case "mainnet":
		blkToWait = 6
		tyNet = utils.TyMainnet
	case "regtest":
		tyNet = utils.TyRegtest
	case "simnet":
		tyNet = utils.TySimnet
	}

	rawTy := make([]byte, 8)
	binary.LittleEndian.PutUint64(rawTy, uint64(tyNet))

	if err := updateSideChain(poly, acc, config.DefConfig.BtcChainID, 1, blkToWait, "btc", rawTy); err != nil {
		log.Errorf("failed to update btc: %v", err)
		return false
	}

	return true
}

func UpdateEth(poly *poly_go_sdk.PolySdk, acc *poly_go_sdk.Account) bool {
	blkToWait := uint64(1)
	if config.BtcNet.Name == "testnet3" {
		blkToWait = 12
	}
	eccd, err := hex.DecodeString(strings.Replace(config.DefConfig.Eccd, "0x", "", 1))
	if err != nil {
		log.Errorf("failed to decode eccd: %v", err)
		return false
	}
	if err = updateSideChain(poly, acc, config.DefConfig.EthChainID, 2, blkToWait, "eth", eccd); err != nil {
		log.Errorf("failed to update eth: %v", err)
		return false
	}
	return true
}

func UpdateNeo(poly *poly_go_sdk.PolySdk, acc *poly_go_sdk.Account) bool {
	blkToWait := uint64(1)
	eccd, err := common2.AddressFromHexString(strings.TrimPrefix(config.DefConfig.NeoCCMC, "0x"))
	if err != nil {
		log.Errorf("failed to decode eccd: %v", err)
		return false
	}
	if err = updateSideChain(poly, acc, config.DefConfig.NeoChainID, 4, blkToWait, "NEO", eccd[:]); err != nil {
		log.Errorf("failed to update neo: %v", err)
		return false
	}
	return true
}

func updateSideChain(poly *poly_go_sdk.PolySdk, acc *poly_go_sdk.Account, chainId, router, blkToWait uint64, name string,
	ccmc []byte) error {
	txhash, err := poly.Native.Scm.UpdateSideChain(acc.Address, chainId, router, name, blkToWait, ccmc, acc)
	if err != nil {
		return err
	}

	testcase.WaitPolyTx(txhash, poly)
	log.Infof("successful to update %s: ( txhash: %s )", name, txhash.ToHexString())
	return nil
}

func ApproveUpdateChain(id uint64, poly *poly_go_sdk.PolySdk, accArr []*poly_go_sdk.Account) {
	var (
		txhash common.Uint256
		err    error
	)

	for i, a := range accArr {
		txhash, err = poly.Native.Scm.ApproveUpdateSideChain(id, a)
		if err != nil {
			panic(fmt.Errorf("No%d ApproveUpdateChain failed: %v", i, err))
		}
		log.Infof("No%d: successful to approve update chain: ( acc: %s, txhash: %s, chain-id: %d )",
			i, a.Address.ToHexString(), txhash.ToHexString(), id)
	}
	testcase.WaitPolyTx(txhash, poly)
}

func RegisterCandidate(poly *poly_go_sdk.PolySdk, acc *poly_go_sdk.Account) bool {
	txHash, err := poly.Native.Nm.RegisterCandidate(vconfig.PubkeyID(acc.GetPublicKey()), acc)
	if err != nil {
		if strings.Contains(err.Error(), "already") {
			log.Warnf("candidate %s already registered: %v", acc.Address.ToBase58(), err)
			return true
		}
		log.Errorf("sendTransaction error: %v", err)
	}
	testcase.WaitPolyTx(txHash, poly)
	log.Infof("successful to register candidate: ( candidate: %s, txhash: %s )",
		acc.Address.ToHexString(), txHash.ToHexString())

	return true
}

func ApproveCandidate(acc *poly_go_sdk.Account, poly *poly_go_sdk.PolySdk, accArr []*poly_go_sdk.Account) {
	var (
		txhash common.Uint256
		err    error
	)
	for i, a := range accArr {
		txhash, err = poly.Native.Nm.ApproveCandidate(vconfig.PubkeyID(acc.GetPublicKey()), a)
		if err != nil {
			panic(fmt.Errorf("no%d sendTransaction error: %v", i, err))
		}
		log.Infof("No%d: successful to approve candidate: ( acc: %s, txhash: %s, candidate: %s )",
			i, a.Address.ToHexString(), txhash.ToHexString(), acc.Address.ToHexString())
	}
	testcase.WaitPolyTx(txhash, poly)
}

func BlackPolyNode(poly *poly_go_sdk.PolySdk, acc *poly_go_sdk.Account, accArr []*poly_go_sdk.Account) {
	var (
		txhash common.Uint256
		err    error
	)
	for i, v := range accArr {
		txhash, err = poly.Native.Nm.BlackNode([]string{vconfig.PubkeyID(acc.PublicKey)}, v)
		if err != nil {
			panic(fmt.Errorf("no%d - failed to black node %s: %v", i, acc.Address.ToHexString(), err))
		}
		log.Infof("No%d: successful to black node: ( acc: %s, txhash: %s, node_to_black: %s )",
			i, v.Address.ToHexString(), txhash.ToHexString(), acc.Address.ToHexString())
	}
	testcase.WaitPolyTx(txhash, poly)
}

func WhitePolyNode(poly *poly_go_sdk.PolySdk, acc *poly_go_sdk.Account, accArr []*poly_go_sdk.Account) {
	var (
		txhash common.Uint256
		err    error
	)
	for i, v := range accArr {
		txhash, err = poly.Native.Nm.WhiteNode(vconfig.PubkeyID(acc.PublicKey), v)
		if err != nil {
			panic(fmt.Errorf("no%d - failed to white node %s: %v", i, acc.Address.ToHexString(), err))
		}
		log.Infof("No%d: successful to white node: ( acc: %s, txhash: %s, node_to_white: %s )",
			i, v.Address.ToHexString(), txhash.ToHexString(), acc.Address.ToHexString())
	}
	testcase.WaitPolyTx(txhash, poly)
}

func CommitPolyDpos(poly *poly_go_sdk.PolySdk, accArr []*poly_go_sdk.Account) {
	txhash, err := poly.Native.Nm.CommitDpos(accArr)
	if err != nil {
		panic(err)
	}
	testcase.WaitPolyTx(txhash, poly)
	log.Infof("successful to commit dpos on Poly: txhash: %s", txhash.ToHexString())
}

func CommitOntDpos() {
	ontCli := ontology_go_sdk.NewOntologySdk()
	ontCli.NewRpcClient().SetAddress(config.DefConfig.OntJsonRpcAddress)
	tx, err := ontCli.Native.NewNativeInvokeTransaction(config.DefConfig.GasPrice, config.DefConfig.GasLimit, 0, ontology_go_sdk.GOVERNANCE_CONTRACT_ADDRESS, governance.COMMIT_DPOS, nil)
	if err != nil {
		panic(fmt.Errorf("NewNativeInvokeTransaction error: %+v\n", err))
	}
	ow := strings.Split(oWalletFiles, ",")
	op := strings.Split(oPwds, ",")
	oAccArr := make([]*ontology_go_sdk.Account, len(ow))
	pks := make([]keypair.PublicKey, len(ow))
	for i, v := range ow {
		oAccArr[i], err = ont.GetOntAccByPwd(v, op[i])
		if err != nil {
			panic(fmt.Errorf("failed to decode no%d wallet %s with pwd %s", i, ow[i], op[i]))
		}
		pks[i] = oAccArr[i].PublicKey
	}
	for _, signer := range oAccArr {
		err = ontCli.MultiSignToTransaction(tx, uint16((5*len(pks)+6)/7), pks, signer)
		if err != nil {
			panic(fmt.Errorf("multi sign failed, err: %s\n", err))
		}
	}
	txHash, err := ontCli.SendTransaction(tx)
	if err != nil {
		panic(fmt.Errorf("sendTransaction error: %+v\n", err))
	}
	ont.WaitOntTx(txHash, ontCli)
	log.Infof("successful to commit dpos on Ontology: txhash: %s", txHash.ToHexString())
}

func QuitNode(poly *poly_go_sdk.PolySdk, acc *poly_go_sdk.Account) {
	txhash, err := poly.Native.Nm.QuitNode(vconfig.PubkeyID(acc.PublicKey), acc)
	if err != nil {
		panic(fmt.Errorf("failed to quit %s: %v", acc.Address.ToBase58(), err))
	}
	testcase.WaitPolyTx(txhash, poly)
	log.Infof("successful to quit node %s on Poly: txhash: %s", acc.Address.ToBase58(), txhash.ToHexString())
}

func GetPolyConsensusInfo(poly *poly_go_sdk.PolySdk) {
	storeBs, err := poly.GetStorage(utils.NodeManagerContractAddress.ToHexString(), []byte(node_manager.GOVERNANCE_VIEW))
	if err != nil {
		panic(err)
	}
	source := common.NewZeroCopySource(storeBs)
	gv := new(node_manager.GovernanceView)
	if err := gv.Deserialization(source); err != nil {
		panic(err)
	}

	raw, err := poly.GetStorage(utils.NodeManagerContractAddress.ToHexString(),
		append([]byte(node_manager.PEER_POOL), utils.GetUint32Bytes(gv.View)...))
	if err != nil {
		panic(err)
	}
	m := &node_manager.PeerPoolMap{
		PeerPoolMap: make(map[string]*node_manager.PeerPoolItem),
	}
	if err := m.Deserialization(common.NewZeroCopySource(raw)); err != nil {
		panic(err)
	}
	str := ""
	for _, v := range m.PeerPoolMap {
		str += fmt.Sprintf("[ index: %d, address: %s, pubk: %s, status: %d ]\n",
			v.Index, v.Address.ToBase58(), v.PeerPubkey, v.Status)
	}

	log.Infof("get consensus info of poly: { view: %d, len_nodes: %d, info: \n%s }", gv.View, len(m.PeerPoolMap), str)
}

func RegisterRelayer(poly *poly_go_sdk.PolySdk, acc *poly_go_sdk.Account, signer *poly_go_sdk.Account) uint64 {
	txhash, err := poly.Native.Rm.RegisterRelayer([]common.Address{acc.Address}, signer)
	if err != nil {
		panic(err)
	}
	testcase.WaitPolyTx(txhash, poly)
	log.Infof("successful to register a relayer %s, txhash is %s", acc.Address.ToBase58(), txhash.ToHexString())
	event, err := poly.GetSmartContractEvent(txhash.ToHexString())
	if err != nil {
		panic(err)
	}
	var id uint64
	for _, e := range event.Notify {
		states := e.States.([]interface{})
		if states[0].(string) == "putRelayerApply" {
			id = uint64(states[1].(float64))
		}
	}
	return id
}

func ApproveRelayer(poly *poly_go_sdk.PolySdk, id uint64, accArr []*poly_go_sdk.Account) {
	var (
		txhash common.Uint256
		err    error
	)
	for i, v := range accArr {
		txhash, err = poly.Native.Rm.ApproveRegisterRelayer(id, v)
		if err != nil {
			panic(fmt.Errorf("no%d - failed to approve %d: %v", i, id, err))
		}
		log.Infof("No%d: successful to approve relayer id %d: ( acc: %s, txhash: %s )",
			i, id, v.Address.ToHexString(), txhash.ToHexString())
	}
	testcase.WaitPolyTx(txhash, poly)
}

func RemoveRelayer(poly *poly_go_sdk.PolySdk, acc *poly_go_sdk.Account) uint64 {
	txhash, err := poly.Native.Rm.RemoveRelayer([]common.Address{acc.Address}, acc)
	if err != nil {
		panic(err)
	}
	testcase.WaitPolyTx(txhash, poly)
	log.Infof("successful to remove a relayer %s, txhash is %s", acc.Address.ToBase58(), txhash.ToHexString())
	event, err := poly.GetSmartContractEvent(txhash.ToHexString())
	if err != nil {
		panic(err)
	}
	var id uint64
	for _, e := range event.Notify {
		states := e.States.([]interface{})
		if states[0].(string) == "putRelayerRemove" {
			id = uint64(states[1].(float64))
		}
	}
	return id
}

func ApproveRemoveRelayer(poly *poly_go_sdk.PolySdk, id uint64, accArr []*poly_go_sdk.Account) {
	var (
		txhash common.Uint256
		err    error
	)
	for i, v := range accArr {
		txhash, err = poly.Native.Rm.ApproveRemoveRelayer(id, v)
		if err != nil {
			panic(fmt.Errorf("no%d - failed to approve %d: %v", i, id, err))
		}
		log.Infof("No%d: successful to approve remove relayer id %d: ( acc: %s, txhash: %s )",
			i, id, v.Address.ToHexString(), txhash.ToHexString())
	}
	testcase.WaitPolyTx(txhash, poly)
}

func GetRelayer(poly *poly_go_sdk.PolySdk, acc *poly_go_sdk.Account) {
	raw, err := poly.GetStorage(utils.RelayerManagerContractAddress.ToHexString(),
		append([]byte(relayer_manager.RELAYER), acc.Address[:]...))
	if err != nil {
		panic(err)
	}
	if len(raw) == 0 {
		log.Infof("no this relayer %s", acc.Address.ToBase58())
		return
	}
	addr, err := common.AddressParseFromBytes(raw)
	if err != nil {
		panic(err)
	}
	log.Infof("get relayer success: %s", addr.ToBase58())
}

func QuitSideChain(poly *poly_go_sdk.PolySdk, id uint64, acc *poly_go_sdk.Account) {
	txhash, err := poly.Native.Scm.QuitSideChain(id, acc)
	if err != nil {
		panic(fmt.Errorf("failed to quit %s: %v", acc.Address.ToBase58(), err))
	}
	testcase.WaitPolyTx(txhash, poly)
	log.Infof("successful to quit side chain %s on Poly: txhash: %s", acc.Address.ToBase58(), txhash.ToHexString())
}

func ApproveQuitSideChain(poly *poly_go_sdk.PolySdk, id uint64, accArr []*poly_go_sdk.Account) {
	var (
		txhash common.Uint256
		err    error
	)
	for i, v := range accArr {
		txhash, err = poly.Native.Scm.ApproveQuitSideChain(id, v)
		if err != nil {
			panic(fmt.Errorf("no%d - failed to approve %d: %v", i, id, err))
		}
		log.Infof("No%d: successful to approve quit side chain %d: ( acc: %s, txhash: %s )",
			i, id, v.Address.ToHexString(), txhash.ToHexString())
	}
	testcase.WaitPolyTx(txhash, poly)
}

func GetSideChain(poly *poly_go_sdk.PolySdk, id uint64) {
	store, err := poly.GetStorage(utils.SideChainManagerContractAddress.ToHexString(),
		append([]byte(side_chain_manager.SIDE_CHAIN), utils.GetUint64Bytes(id)...))
	if err != nil {
		panic(err)
	}
	if store == nil {
		log.Infof("no this %d side chain found", id)
		return
	}
	sideChain := new(side_chain_manager.SideChain)
	err = sideChain.Deserialization(common.NewZeroCopySource(store))
	if err != nil {
		panic(err)
	}
	log.Infof("side chain %d, name: %s, addr: %s", id, sideChain.Name, sideChain.Address.ToBase58())
}

func UpdatePolyConfig(poly *poly_go_sdk.PolySdk, blockMsgDelay, hashMsgDelay, peerHandshakeTimeout,
	maxBlockChangeView uint32, accArr []*poly_go_sdk.Account) {
	txhash, err := poly.Native.Nm.UpdateConfig(blockMsgDelay, hashMsgDelay, peerHandshakeTimeout, maxBlockChangeView, accArr)
	if err != nil {
		panic(err)
	}
	testcase.WaitPolyTx(txhash, poly)
	log.Infof("update poly config: "+
		"(blockMsgDelay: %d, hashMsgDelay: %d, peerHandshakeTimeout: %d, maxBlockChangeView: %d)",
		blockMsgDelay, hashMsgDelay, peerHandshakeTimeout, maxBlockChangeView)
}

func GetPolyConfig(poly *poly_go_sdk.PolySdk) {
	raw, err := poly.GetStorage(utils.NodeManagerContractAddress.ToHexString(), []byte(node_manager.VBFT_CONFIG))
	if err != nil {
		panic(err)
	}
	conf := &node_manager.Configuration{}
	if err = conf.Deserialization(common.NewZeroCopySource(raw)); err != nil {
		panic(err)
	}
	log.Infof("poly config: (blockMsgDelay: %d, hashMsgDelay: %d, peerHandshakeTimeout: %d, maxBlockChangeView: %d)",
		conf.BlockMsgDelay, conf.HashMsgDelay, conf.PeerHandshakeTimeout, conf.MaxBlockChangeView)
}

func UnRegisterPolyCandidate(poly *poly_go_sdk.PolySdk, acc *poly_go_sdk.Account) {
	txhash, err := poly.Native.Nm.UnRegisterCandidate(vconfig.PubkeyID(acc.PublicKey), acc)
	if err != nil {
		panic(err)
	}
	testcase.WaitPolyTx(txhash, poly)
	log.Infof("unregister %s success: txhash: %s", acc.Address.ToBase58(), txhash.ToHexString())
}

func CosmosCreateValidator(keyFile, stateFile string, amt int64) {
	invoker, err := cosmos2.NewCosmosInvoker()
	if err != nil {
		panic(err)
	}

	dc, err := types3.ParseDecCoins(config.DefConfig.CMGasPrice)
	if err != nil {
		panic(err)
	}
	demon := dc.GetDenomByIndex(0)
	res, err := invoker.CreateValidator(keyFile, stateFile, demon, amt)
	if err != nil {
		panic(err)
	}
	log.Infof("CosmosCreateValidator, create a new validator %s with %d %s: txhash %s",
		invoker.Acc.Acc.String(), amt, demon, res.Hash.String())
}

func CosmosDelegateToVal(amt int64) {
	invoker, err := cosmos2.NewCosmosInvoker()
	if err != nil {
		panic(err)
	}

	dc, err := types3.ParseDecCoins(config.DefConfig.CMGasPrice)
	if err != nil {
		panic(err)
	}
	demon := dc.GetDenomByIndex(0)
	res, err := invoker.DelegateValidator(demon, amt)
	if err != nil {
		panic(err)
	}
	log.Infof("CosmosDelegateToVal, delegate %d %s to validator %s: txhash %s",
		amt, demon, types3.ValAddress(invoker.Acc.Acc).String(), res.Hash.String())
}
