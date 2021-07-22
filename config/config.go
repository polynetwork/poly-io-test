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
package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/polynetwork/poly-io-test/log"
)

const (
	CM_BTCX  = "btcx"
	CM_ETHX  = "ethx"
	CM_BNBX  = "bnbx"
	CM_ERC20 = "erc20x"
	CM_BEP20 = "bep20x"
	CM_ONT   = "ontx"
	CM_ONG   = "ongx"
	CM_OEP4  = "oep4x"
	CM_KRC20 = "krc20"
	CM_KAIX  = "kaix"
)

//Config object used by ontology-instance
type TestConfig struct {
	BtcChainID              uint64
	EthChainID              uint64
	OntChainID              uint64
	NeoChainID              uint64
	Neo3ChainID             uint64
	BscChainID              uint64
	ZilChainID              uint64
	HecoChainID             uint64
	O3ChainID               uint64
	MscChainID              uint64
	OkChainID               uint64
	KaiChainID              uint64
	PolygonHeimdallChainID  uint64
	PolygonBorChainID       uint64
	PolygonBorSignerChainID uint64

	BtcRestAddr                  string
	BtcRestUser                  string
	BtcRestPwd                   string
	BtcFee                       int64
	BtcRedeem                    string // auto set
	BtcNetType                   string
	BtcMultiSigNum               int // multi-sig vendor
	BtcMultiSigRequire           int
	BtcEncryptedPrivateKeyFile   string
	BtcEncryptedPrivateKeyPwd    string
	BtcVendorSigningToolConfFile string
	BtcFeeRate                   uint64
	BtcMinChange                 uint64
	BtcMinOutputValFromContract  uint64 // min btc val for btcx output
	BtcSignerPrivateKey          string
	BtcExistingVendorPrivks      string

	// eth urls
	EthURL        string
	ETHPrivateKey string

	// msc urls
	MSCURL        string
	MSCPrivateKey string

	// ok urls
	OKURL        string
	OKPrivateKey string
	// bsc urls
	BSCURL        string
	BSCPrivateKey string
	// heco urls
	HecoURL        string
	HecoPrivateKey string
	// o3 urls
	O3URL        string
	O3PrivateKey string
	// bor urls
	BorURL        string
	BorPrivateKey string

	// zil urls
	ZilURL        string
	ZilPrivateKey string

	// ontology
	OntJsonRpcAddress   string
	OntWallet           string
	OntWalletPassword   string
	GasPrice            uint64
	GasLimit            uint64
	OntContractsAvmPath string
	OntEpoch            uint32

	// cosmos
	CMWalletPath   string
	CMWalletPwd    string
	CMRpcUrl       string
	CMChainId      string
	CMGasPrice     string
	CMGas          uint64
	CMCrossChainId uint64
	CMEpoch        int64

	// neo chain conf
	NeoUrl   string
	NeoWif   string
	NeoEpoch uint32

	// neo3 chain
	Neo3Url            string
	Neo3Wallet         string
	Neo3Pwd            string
	Neo3Magic          uint32
	Neo3AddressVersion byte
	Neo3Epoch          uint32

	// Kai Chain
	KaiUrl        string
	KaiPrivateKey string
	KaiEpoch      int64

	// relayer chain
	RCWallet             string
	RCWalletPwd          string
	RchainJsonRpcAddress string
	RCEpoch              uint32

	ReportInterval uint64
	ReportDir      string
	BatchTxNum     uint64
	BatchInterval  uint64

	// Circle batch
	TxNumPerBatch uint64

	// msc contracts: auto set after deploy
	MscEccd      string
	MscEccm      string
	MscEccmp     string
	MscLockProxy string
	Mep20        string
	MscOep4      string
	MscOngx      string
	MscOntx      string
	MscWBTC      string
	MscUSDT      string
	MscDai       string
	MscUSDC      string
	MscNeo       string
	MscRenBTC    string
	// ok contracts: auto set after deploy
	OkEccd      string
	OkEccm      string
	OkEccmp     string
	OkLockProxy string
	OkErc20     string
	OkOep4      string
	OkOngx      string
	OkOntx      string
	OkWBTC      string
	OkUSDT      string
	OkDai       string
	OkUSDC      string
	OkNeo       string
	OkRenBTC    string
	// bsc contracts: auto set after deploy
	BscEccd      string
	BscEccm      string
	BscEccmp     string
	BscLockProxy string
	Bep20        string
	BscOep4      string
	BscOngx      string
	BscOntx      string
	BscWBTC      string
	BscUSDT      string
	BscDai       string
	BscUSDC      string
	BscNeo       string
	BscRenBTC    string
	// heco contracts: auto set after deploy
	HecoEccd      string
	HecoEccm      string
	HecoEccmp     string
	HecoLockProxy string
	HecoErc20     string
	HecoOep4      string
	HecoOngx      string
	HecoOntx      string
	HecoWBTC      string
	HecoUSDT      string
	HecoDai       string
	HecoUSDC      string
	HecoNeo       string
	HecoRenBTC    string

	// kai contracts: auto set after deploy
	KaiEccd      string
	KaiEccm      string
	KaiEccmp     string
	KaiLockProxy string
	Krc20        string
	KaiOep4      string
	KaiOngx      string
	KaiOntx      string
	KaiWBTC      string
	KaiUSDT      string
	KaiDai       string
	KaiUSDC      string
	KaiNeo       string
	KaiRenBTC    string

	// o3 contracts: auto set after deploy
	O3Eccd      string
	O3Eccm      string
	O3Eccmp     string
	O3LockProxy string
	O3Erc20     string
	O3Oep4      string
	O3Ongx      string
	O3Ontx      string
	O3WBTC      string
	O3USDT      string
	O3Dai       string
	O3USDC      string
	O3Neo       string
	O3RenBTC    string

	// bor contracts
	BorEccd      string
	BorEccm      string
	BorEccmp     string
	BorLockProxy string

	// zil contracts
	ZilEccdProxy string
	ZilEccdImpl  string
	ZilLockProxy string

	// eth contracts: auto set after deploy
	EthBnb              string
	EthErc20            string
	EthOep4             string
	Eccd                string
	Eccm                string
	Eccmp               string
	EthLockProxy        string
	EthOngx             string
	EthOntx             string
	EthOntd             string
	EthUSDT             string
	EthWBTC             string
	EthDai              string
	EthUSDC             string
	EthNeo              string
	EthRenBTC           string
	BtceContractAddress string

	// ont contracts: auto set after deploy
	OntErc20            string
	OntBep20            string
	OntKrc20            string
	OntOep4             string
	OntLockProxy        string
	OntEth              string
	OntBnb              string
	OntUSDT             string
	OntWBTC             string
	OntDai              string
	OntUSDC             string
	OntNeo              string
	OntONTD             string
	OntRenBTC           string
	BtcoContractAddress string
	OntKai              string

	// neo
	NeoCCMC      string
	NeoLockProxy string
	CNeo         string
	NeoOnt       string
	NeoOntd      string
	NeoBnb       string
	NeoEth       string

	// neo3 contracts
	Neo3CCMC      string
	Neo3LockProxy string
	Neo3Ont       string
	Neo3Ontd      string
	Neo3Bnb       string
	Neo3Eth       string
	Neo3Ht        string
	Neo3Hrc20     string

	// cosmos
	CMLockProxy string

	// transfer amount
	BtcValLimit    uint64
	OntValLimit    uint64
	OntdValLimit   uint64
	OngValLimit    uint64
	EthValLimit    uint64
	Oep4ValLimit   uint64
	Erc20ValLimit  uint64
	USDTValLimit   uint64
	NeoValLimit    uint64
	WBTCValLimit   uint64
	USDCValLimit   uint64
	RenBTCValLimit uint64
	KaiValLimit    uint64

	OntdValFloor uint64
}

//Default config instance
var DefConfig = NewDefaultTestConfig()
var DefaultConfigFile = "./config.json"
var BtcNet *chaincfg.Params

//NewTestConfig retuen a TestConfig instance
func NewTestConfig() *TestConfig {
	return &TestConfig{}
}

func NewDefaultTestConfig() *TestConfig {
	var config = NewTestConfig()
	err := config.Init(DefaultConfigFile)
	if err != nil {
		return &TestConfig{}
	}
	return config
}

//Init TestConfig with a config file
func (conf *TestConfig) Init(fileName string) error {
	err := conf.loadConfig(fileName)
	if err != nil {
		return fmt.Errorf("loadConfig error:%s", err)
	}

	BtcNet = func(netTy string) *chaincfg.Params {
		switch netTy {
		case "regtest":
			return &chaincfg.RegressionNetParams
		case "test":
			return &chaincfg.TestNet3Params
		case "simnet":
			return &chaincfg.SimNetParams
		default:
			return &chaincfg.MainNetParams
		}
	}(conf.BtcNetType)

	return nil
}

/**
Load JSON Configuration
*/
func (conf *TestConfig) loadConfig(fileName string) error {
	data, err := conf.readFile(fileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, conf)
	if err != nil {
		return fmt.Errorf("json.Unmarshal TestConfig:%s error:%s", data, err)
	}
	return nil
}

/**
Read  File to bytes
*/
func (conf *TestConfig) readFile(fileName string) ([]byte, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("OpenFile %s error %s", fileName, err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Errorf("File %s close error %s", fileName, err)
		}
	}()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll %s error %s", fileName, err)
	}
	return data, nil
}

/**
Save Test Configuration To json file
*/
func (conf *TestConfig) Save(fileName string) error {
	data, err := json.MarshalIndent(conf, "", "\t")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(fileName, data, 0644); err != nil {
		return fmt.Errorf("failed to write conf file: %v", err)
	}
	return nil
}
