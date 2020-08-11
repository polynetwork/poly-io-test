<h1 align="center">Poly IO Test</h1>

## Introduction

This is the testframework for Poly ecosystem which can send cross-chain transacions and check it status. Let's call this tool PIT for short.

There is two parts about PIT. First part including five functions is for create and setup cross-chain environment. Like creating accounts on chains and deploying contracts. Second part is for running testcases like sending cross-chain transactions. 

PIT can help you build your cross-chain environment and providing examples about writing code to send cross-chain transacion.

The code is still under development.

| Part | Func           | Desc                                                         |
| ---- | -------------- | ------------------------------------------------------------ |
| 1    | btc_prepare    | Generate BTC multi-signature address and other information (you can also use existing multi-signature), deploy and bind BTCX contracts on each chain, and register BTCX contracts and transaction parameters with Poly. |
| 1    | eth_deployer   | Two functions: deploy all the contracts on the Ethereum chain, from ECCM to each asset; to set up the binding of the contract, other chains need to complete the deployment first to ensure that the contract hash has been filled in the config file. |
| 1    | ont_deployer   | Same as `eth_deployer`                                       |
| 1    | cosmos_prepare | Initialize the chains based on COSMOS-SDK like Switcheo, create each asset and complete asset binding. |
| 1    | tools          | Register the sidechain with Poly and sync the genesis block between chains. |
| 2    | cctest         | Run testcases.                                               |

## Configuration

```
{
   ###
   # This is BTC related settings
   "BtcRestAddr": "http://172.168.3.10:20336", # BTC node
   "BtcRestUser": "test",
   "BtcRestPwd": "test",
   "BtcFee": 1500, 
   "BtcRedeem": "522103c4564b837674de2482961a8d5f2a24a7e11e8a97aac5e92ac2e64500219144512102ccc07d3df7da58bb6fa5cfe5d7be415ff9463171b2600c93c080fcd0d49576a721036ec6299c1b14e57b45f1ad85eecbc48ad5447a05158a1bfb2ffb689ad69490d353ae", # Multi-sign Redeem script
   "BtcNetType": "test", # Network Type
   "BtcMultiSigNum": 3, # Multi-signature total
   "BtcMultiSigRequire": 2, # Minimum number of multi-signatures
   "BtcEncryptedPrivateKeyFile": "/data/poly_deployer/lib/vendor_tool/btcprivk", # Multi-signature encrypted wallet storage path
   "BtcEncryptedPrivateKeyPwd": "123",
   "BtcVendorSigningToolConfFile": "/data/poly_deployer/lib/vendor_tool/vendor_tool.json", # vendor configuration file
   "BtcFeeRate": 10, # Multi-signature rate, when BTC returns to Bitcoin, unlock transaction rate
   "BtcMinChange": 8000, # Multi-sign the minimum UTXO value allowed in the Poly chain contract
   "BtcMinOutputValFromContract": 10000, # The minimum allowable withdrawal amount in the BTCX contract
   "BtcSignerPrivateKey": "cRRMYvoHPNQu1tCz4ajPxytBVc2SN6GWLAVuyjzm4MVwyqZVrAcX", # BTC simulates the private key for sending transactions
   "BtcExistingVendorPrivks": "cREJsmv4W9Lr4QhvnQrA77tTUZ5g488qNr2cvwRoSnpxryGRL92N,cVJqF57cdiSkFwpnTqEzjX7hCxZH95vxwenuJtBPgABPCwBDymEk,cTfnpP7CBFQ5mtK2BKPJLUnRJBYgnnRX7M8MnWpvWthPRoB3yGwR", # Existing multi-signature wallet private key
   ###
   
   ###
   # eth
   "EthURL": "http://18.139.17.85:10331", # Ethereum node
   "ETHPrivateKey": "AEC101ECDB5C86931E0CA5E635824F0D0F05240760C01DDEE64BE90B2A2608A9", # Etherem simulates the private key of the transaction
   ###
   
   ###
   # Ontology
   "OntJsonRpcAddress": "http://172.168.3.73:20336", # Ontology
   "OntWallet": ".wallets/ont_dev.dat",
   "OntWalletPassword": "admin",
   "GasPrice": 0,
   "GasLimit": 30000000,
   "OntContractsAvmPath": "./chains/ont/avm", # Avm of all ontology contracts
   ###
   
   ###
   # cosmos
   "CMWalletPath": ".wallets/cosmos_key",#
   "CMWalletPwd": "11111111",
   "CMRpcUrl": "http://172.168.3.95:26657",
   "CMChainId": "cc-cosmos",
   "CMGasPrice": "0.000001stake",
   "CMGas": 200000,
   "CMCrossChainId": 7,
   ###
   
   ###
   # Poly
   "RCWallet": "/Users/zou/Desktop/work/code/poly_d/poly_deployer_linux_testnet/lib/poly/wallet.dat",
   "RCWalletPwd": "4cUYqGj2yib718E7ZmGQc",
   "RchainJsonRpcAddress": "http://172.168.3.73:40336",
	 ###
	 
	 ###
	 # 发测试交易的配置
   "ReportInterval": 10,
   "ReportDir": "./report", # Test case status report
   "BatchTxNum": 100, # Batch send transaction: how many batches
   "BatchInterval": 1, # Time interval for batch sending
   "TxNumPerBatch": 100, # How many transactions per batch
   ###
   
   ###
   # This part is the contract configuration of each chain. After deployment, it is automatically written
   "EthErc20": "0xE4dc8fa991a3e0BC41f5E736BAc69b8186a1386D",
   "EthOep4": "0x15b55ca5C3082c95Bd4707651494384D2d2A6b50",
   "Eccd": "0x9770198364BF08C0fEd483E67C22551632500339",
   "Eccm": "0xF3A7B7B56E4935ACA162F187AD5526f37f8F4c52",
   "Eccmp": "0xDff4A9B177fa89394a10506b9f99AadA31B635f8",
   "EthLockProxy": "0x20eFad169D803D8d3Ca04f5a71C67A52de687c26",
   "EthOngx": "0x1ed03F921425883ceF5eE59ff129192Ba87DaC48",
   "EthOntx": "0x1666970b08DD77D67aB43441E519381d6bA9c49c",
   "BtceContractAddress": "0x8940fa5B2bb6e961D0bF3E8C3FBd9757ffCDDAf1",
   "OntErc20": "2bd086d90c282a26d021d01039b84a968ed57444",
   "OntOep4": "d6967fa9b11b0836ec9dc8572f4bab6f65d710f9",
   "OntLockProxy": "ebad45b887c6bf7cc4c1df8f72da156bc91b04b7",
   "OntEth": "d6b4cd930377e7e81d1ace85bbb9ce59f4dd9410",
   "BtcoContractAddress": "d177d904456479dd592d4ce368e5d97ec8ffbcee",
   "CMLockProxy": "f71b55ef55cedc91fd007f7a9ba386ec978f3aa8",
   ###
   
   ###
   # Test transaction limit
   "BtcValLimit": 100000,
   "OntValLimit": 100,
   "OngValLimit": 100000,
   "EthValLimit": 10000000,
   "Oep4ValLimit": 10000000,
   "Erc20ValLimit": 10000000
   ###
}
```

