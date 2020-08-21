package eth

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joeqian10/neo-gogogo/wallet"
	"github.com/ontio/ontology-crypto/keypair"
	"github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/consensus/vbft/config"
	"github.com/polynetwork/poly-go-sdk"
	"github.com/polynetwork/poly-io-test/config"
	"github.com/polynetwork/poly-io-test/log"
	vconfig2 "github.com/polynetwork/poly/consensus/vbft/config"
	common3 "github.com/polynetwork/poly/native/service/cross_chain_manager/common"
	"github.com/polynetwork/poly/native/service/utils"
	"math/big"
	"testing"
	"time"
)

func MakeEthAuth(signer *EthSigner, nonce, gasPrice, gasLimit uint64) *bind.TransactOpts {
	auth := bind.NewKeyedTransactor(signer.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(int64(0)) // in wei
	auth.GasLimit = gasLimit          // in units
	auth.GasPrice = big.NewInt(int64(gasPrice))

	return auth
}

func TestNewNonceManager(t *testing.T) {
	//err := config.DefConfig.Init("config.json")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//invoker := NewEInvoker()
	//
	//contract, err := ongx_abi.NewONGX(common.HexToAddress("0xFb37c160CFBd8BD4Ba6df6f70e2449b6EB83fc26"), invoker.ETHUtil.GetEthClient())
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//amt, err := contract.Allowance(nil, common.HexToAddress("0x344cFc3B8635f72F14200aAf2168d9f75df86FD3"),
	//	common.HexToAddress("0x388Ed8B73bd707A78034E1d157fA08Da24095c18"))
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//fmt.Println(amt.Uint64(), 3%1)

	err := config.DefConfig.Init("/Users/zou/go/src/github.com/ontio/poly-io-test/config-main.json")
	if err != nil {
		t.Fatal(err)
	}

	//poly := poly_go_sdk.NewPolySdk()
	//poly.NewRpcClient().SetAddress(config.DefConfig.RchainJsonRpcAddress)
	//
	//raw, err := poly.GetStorage(utils.SideChainManagerContractAddress.ToHexString(), append([]byte(side_chain_manager.SIDE_CHAIN), utils.GetUint64Bytes(4)...))
	//if err != nil {
	//	t.Fatal(err)
	//}
	//sideChain := new(side_chain_manager.SideChain)
	//if err := sideChain.Deserialization(common2.NewZeroCopySource(raw)); err != nil {
	//	t.Fatal(fmt.Errorf("getSideChain, deserialize sideChain error: %v", err))
	//}
	//fmt.Println(hex.EncodeToString(sideChain.CCMCAddress))

	nUSDTAddr := "0x2205d2F559ef91580090011Aa4E0eF68Ec33da44"
	tx := "0x05a873aba68598585b0d22553de8a86a158992b2ddca77e82eea95d06205d550"
	txhash := common.HexToHash(tx)

	invoker := NewEInvoker()
	invoker.ETHUtil.WaitTransactionConfirm(txhash)

	gasPrice := big.NewInt(100 * 1e9)

	token, err := NewERC20(common.HexToAddress(nUSDTAddr), invoker.ETHUtil.GetEthClient())
	if err != nil {
		t.Fatal(err)
	}

	bal, err := token.BalanceOf(nil, invoker.EthTestSigner.Address)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(bal.Uint64())

	auth := MakeEthAuth(invoker.EthTestSigner, 3, gasPrice.Uint64(), uint64(80000))
	fmt.Println(auth.GasPrice.Uint64(), auth.GasLimit)
	txn, err := token.Approve(auth, common.HexToAddress("0x8B083BbB4c5200fb9653F28d010e92daB71CFa71"), bal)
	if err != nil {
		t.Fatal(err)
	}
	hash := txn.Hash()
	log.Infof("get ETH now and sending tx %s with gasPrice %d", hash.String(), gasPrice.Uint64())
	invoker.ETHUtil.WaitTransactionConfirm(hash)

}

func Test_Scan_ONT_hegit_To_grab_latest_Consensus_Switch_Height1(t *testing.T) {
	ontSdk := ontology_go_sdk.NewOntologySdk()
	url := "http://dappnode1.ont.io:20336"
	ontSdk.NewRpcClient().SetAddress(url)
	curHeight, err := ontSdk.GetCurrentBlockHeight()
	if err != nil {
		fmt.Printf("ontology, get current block height error: %v\n", err)
		return
	}
	fmt.Printf("Current ontlogy net: %s, height: %d\n", url, curHeight)
	var height uint32 = 9354300
	found := false
	for !found && height <= height+1000 {
		block, err := ontSdk.GetBlockByHeight(height)
		if err != nil {
			fmt.Printf("GetBLockByHeight err: %v\n", err)
		}
		header := block.Header

		blkInfo := &vconfig.VbftBlockInfo{}
		if err := json.Unmarshal(header.ConsensusPayload, blkInfo); err != nil {
			fmt.Printf("unmarshal blockInfo error: %v\n", err)
			return
		}
		if blkInfo.NewChainConfig != nil {
			fmt.Printf("===========FOUND============at height: %d\n", height)
			found = true
			return
		}
		height++
		if height%5 == 0 {
			fmt.Printf("height: %d not found\n", height)
		}
	}

}

func Test_Scan_ONT_hegit_To_grab_latest_Consensus_Switch_Height(t *testing.T) {
	ontSdk := ontology_go_sdk.NewOntologySdk()
	url := "http://dappnode1.ont.io:20336"
	ontSdk.NewRpcClient().SetAddress(url)
	curHeight, err := ontSdk.GetCurrentBlockHeight()
	if err != nil {
		fmt.Printf("ontology, get current block height error: %v\n", err)
		return
	}
	fmt.Printf("Current ontlogy net: %s, height: %d\n", url, curHeight)
	var height uint32 = 9295092
	found := false
	oneByOne := false

	block, err := ontSdk.GetBlockByHeight(height)
	if err != nil {
		t.Fatalf("GetBLockByHeight err: %v\n", err)
	}
	header := block.Header
	old := header.Bookkeepers
	for !found {
		block, err := ontSdk.GetBlockByHeight(height)
		if err != nil {
			fmt.Printf("GetBLockByHeight err: %v\n", err)
		}
		header := block.Header

		blkInfo := &vconfig.VbftBlockInfo{}
		if err := json.Unmarshal(header.ConsensusPayload, blkInfo); err != nil {
			fmt.Printf("unmarshal blockInfo error: %v\n", err)
			return
		}
		if blkInfo.NewChainConfig != nil {
			fmt.Printf("===========FOUND============at height: %d\n", height)
			found = true
			return
		}
		if !oneByOne && !EqualKeepers(old, header.Bookkeepers) {
			height = height - 1000
			fmt.Printf("set %d to restart\n", height)
			oneByOne = true
			continue
		}

		if !oneByOne {
			height = height + 1000
		} else {
			height++
		}
		if time.Now().Second()%60 == 0 {
			fmt.Printf("height: %d not found\n", height)
		}
	}
}

func EqualKeepers(oldBkers, newBkers []keypair.PublicKey) bool {
	if len(oldBkers) != len(newBkers) {
		return false
	}
	idx := 0
	for _, v := range oldBkers {
		rawV := keypair.SerializePublicKey(v)
		for _, u := range newBkers {
			rawU := keypair.SerializePublicKey(u)
			if bytes.Equal(rawV, rawU) {
				idx++
				break
			}
		}
	}
	if idx != len(oldBkers) {
		return false
	}

	return true
}

func TestDeployBasicToken(t *testing.T) {
	poly := poly_go_sdk.NewPolySdk()
	poly.NewRpcClient().SetAddress("http://138.91.6.125:20336")
	contractAddress := utils.CrossChainManagerContractAddress
	c, _ := hex.DecodeString("0100000000000000000000000000000000000000000000000000000000000000")
	key := append(append([]byte(common3.DONE_TX), utils.GetUint64Bytes(2)...), c...)
	// try to get storage

	result, err := poly.GetStorage(contractAddress.ToHexString(), key)
	if err != nil {
		t.Fatal(err)
	}
	if result == nil {
		t.Fatal("err")
	}
	fmt.Println("res", hex.EncodeToString(result))
}

func Test_Scan_POLY_hegit_To_grab_latest_Consensus_Switch_Height(t *testing.T) {
	polySdk := poly_go_sdk.NewPolySdk()
	url := "http://40.115.182.238:40336"
	polySdk.NewRpcClient().SetAddress(url)
	curHeight, err := polySdk.GetCurrentBlockHeight()
	if err != nil {
		fmt.Printf("ontology, get current block height error: %v\n", err)
		return
	}
	wallet.NewAccount()
	fmt.Printf("Current ontlogy net: %s, height: %d\n", url, curHeight)
	startHeight := uint32(168168)
	height := startHeight
	//current scan height: 96757 -> 85185, no switch
	found := false
	backSearchHeight := uint32(startHeight)
	backSearchThreshold := 0
	searchOneByOneFlag := false

	for !found {
		block, err := polySdk.GetBlockByHeight(height)
		if err != nil {
			fmt.Printf("GetBLockByHeight err: %v\n", err)
		}
		header := block.Header

		blkInfo := &vconfig2.VbftBlockInfo{}
		if err := json.Unmarshal(header.ConsensusPayload, blkInfo); err != nil {
			fmt.Printf("unmarshal blockInfo error: %v\n", err)
			return
		}
		if blkInfo.NewChainConfig != nil {
			fmt.Printf("===========Switch consensus ============at height: %d\n", height)
			blkInfo := &vconfig.VbftBlockInfo{}
			if err := json.Unmarshal(header.ConsensusPayload, blkInfo); err != nil {
				fmt.Printf("Header Unmarshal(header.ConsensusPayload, blkInfo), err: %v", err)
				return
			}
			if blkInfo.NewChainConfig != nil {
				fmt.Printf("Height: %d, blkInfo.LastConfigBlockNum is %d\n", header.Height, blkInfo.LastConfigBlockNum)
				for i, p := range blkInfo.NewChainConfig.Peers {
					fmt.Printf("blkInfo.NewChainConfig.Peers[%d], Index: %d, Id: %s\n", i, p.Index, p.ID)
				}
			}
			fmt.Printf("\n\n")

			if len(blkInfo.NewChainConfig.Peers) > 1 {
				fmt.Printf("===========Found ============at height: %d\n", height)
				found = true
				return
			}
			backSearchHeight = height
			backSearchThreshold = 0
			searchOneByOneFlag = false
		}
		if backSearchThreshold > 10 && !searchOneByOneFlag {
			height = backSearchHeight
			fmt.Printf("back to height: %d to scan one by one", height)
			searchOneByOneFlag = true
		} else {
			backSearchThreshold++
		}
		if searchOneByOneFlag {
			height++
		} else {
			height = height + 1000
		}
		fmt.Printf("height: %d not found\n", height)
	}

}
