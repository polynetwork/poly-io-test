package chainsql

import (
	"encoding/json"
	"fmt"
	"github.com/ChainSQL/go-chainsql-api/common"
	"github.com/ChainSQL/go-chainsql-api/core"
	data "github.com/ChainSQL/go-chainsql-api/data"
	eccd_abi "github.com/polynetwork/poly-io-test/chains/chainsql/eccd_abi"
	eccm_abi "github.com/polynetwork/poly-io-test/chains/chainsql/eccm_abi"
	eccmp_abi "github.com/polynetwork/poly-io-test/chains/chainsql/eccmp_abi"
	config2 "github.com/polynetwork/poly-io-test/config"
	"github.com/polynetwork/poly-io-test/log"
	"io/ioutil"
	"os"
)
type ChainsqlInvoker struct {
	ChainsqlSdk *core.Chainsql
	TransOpts *core.TransactOpts
}

type Account struct {
	Address string
	Secrect string
}
type Config struct {
	URL            string
	ServerName     string
	RootCertPath   string
	ClientCertPath string
	ClientKeyPath  string
	Account        Account
}


func NewConfig(configFilePath string) *Config {

	fileContent, err := ReadFile(configFilePath)
	if err != nil {
		log.Errorf("NewServiceConfig: failed, err: %s", err)
		return nil
	}
	config := &Config{}
	err = json.Unmarshal(fileContent, config)
	if err != nil {
		log.Errorf("NewServiceConfig: failed, err: %s", err)
		return nil
	}

	return config
}

// Dial connects a client to the given URL and groupID.
func Dial(config *Config) (*core.Chainsql, error) {
	node := core.NewChainsql()
	node.Connect(
		config.URL,
		config.RootCertPath,
		config.ClientCertPath,
		config.ClientKeyPath,
		config.ServerName)

	node.As(config.Account.Address, config.Account.Secrect)
	return node, nil
}

func NewChainsqlInvoker() (*ChainsqlInvoker, error) {
	instance := &ChainsqlInvoker{}
	cfg := NewConfig(config2.DefConfig.ChainsqlSdkConfFile)

	chainsql, err := Dial(cfg)
	if err != nil {
		return nil, err
	}
	instance.ChainsqlSdk = chainsql
	instance.TransOpts = &core.TransactOpts{
		ContractValue: 0,
		Gas:           30000000,
		Expectation:   "validate_success",
	}
	return instance, nil
}


func ReadFile(fileName string) ([]byte, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("ReadFile: open file %s error %s", fileName, err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Errorf("ReadFile: File %s close error %s", fileName, err)
		}
	}()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ReadFile: ioutil.ReadAll %s error %s", fileName, err)
	}
	return data, nil
}

func (invoker *ChainsqlInvoker) DeployCrossChainDataContract() (string, error) {

	ret,_,err := eccd_abi.DeployEthCrossChainData(invoker.ChainsqlSdk,invoker.TransOpts)
	if err != nil{
		return "",err
	}

	return ret.ContractAddress,nil
}

func (invoker *ChainsqlInvoker) DeployCrossChainManagerContract(ccdcAddress string, chainID uint64) (string, error) {
	account, err := data.NewAccountFromAddress(ccdcAddress)
	if err != nil {
		return "", err
	}
	eccdAddress := common.BytesToAddress(account.Bytes())
	ret,_,err := eccm_abi.DeployEthCrossChainManager(invoker.ChainsqlSdk,invoker.TransOpts,eccdAddress,chainID)
	if err != nil{
		return "",err
	}

	return ret.ContractAddress,nil
}

func (invoker *ChainsqlInvoker) TransaferOwnershipForECCD(ccdcAddress string,ownerAddress string) (*common.TxResult,error){
	account, err := data.NewAccountFromAddress(ownerAddress)
	if err != nil {
		return nil, err
	}
	ccmcAddress := common.BytesToAddress(account.Bytes())

	eccdContract,_ := eccd_abi.NewEthCrossChainData(invoker.ChainsqlSdk,ccdcAddress)
	return eccdContract.TransferOwnership(invoker.TransOpts,ccmcAddress)
}

func (invoker *ChainsqlInvoker) DeployCrossChainManagerProxyContract(ccmcAddress string) (string, error) {
	account, err := data.NewAccountFromAddress(ccmcAddress)
	if err != nil {
		return "", err
	}
	eccmAddress := common.BytesToAddress(account.Bytes())
	ret,_,err := eccmp_abi.DeployEccmpAbi(invoker.ChainsqlSdk,invoker.TransOpts,eccmAddress)
	if err != nil{
		return "",err
	}

	return ret.ContractAddress,nil
}

func (invoker *ChainsqlInvoker) TransferOwnershipForECCM(ccmcAddress string,ccmp string) (*common.TxResult,error){
	account, err := data.NewAccountFromAddress(ccmp)
	if err != nil {
		return nil, err
	}
	ccmpAddress := common.BytesToAddress(account.Bytes())

	eccmContract,_ := eccm_abi.NewEthCrossChainManager(invoker.ChainsqlSdk,ccmcAddress)
	return eccmContract.TransferOwnership(invoker.TransOpts,ccmpAddress)
}