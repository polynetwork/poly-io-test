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

Before you try the testcase or other functions, you need to finish the configuration file: 

```
{
   ###
   # This is BTC related settings
   "BtcRestAddr": "http://ip:port", # BTC node
   "BtcRestUser": "test",
   "BtcRestPwd": "test",
   "BtcFee": 1500, 
   "BtcRedeem": "552102dec9a415b6384ec0a9331d0cdf02020f0f1e5731c327b86e2b5a92455a289748210365b1066bcfa21987c3e207b92e309b95ca6bee5f1133cf04d6ed4ed265eafdbc21031104e387cd1a103c27fdc8a52d5c68dec25ddfb2f574fbdca405edfd8c5187de21031fdb4b44a9f20883aff505009ebc18702774c105cb04b1eecebcb294d404b1cb210387cda955196cc2b2fc0adbbbac1776f8de77b563c6d2a06a77d96457dc3d0d1f2102dd7767b6a7cc83693343ba721e0f5f4c7b4b8d85eeb7aec20d227625ec0f59d321034ad129efdab75061e8d4def08f5911495af2dae6d3e9a4b6e7aeb5186fa432fc57ae", # Multi-sign Redeem script(testnet)
   "BtcNetType": "test", # Network Type
   "BtcMultiSigNum": 7, # Multi-signature total
   "BtcMultiSigRequire": 5, # Minimum number of multi-signatures
   "BtcEncryptedPrivateKeyFile": "./btcprivk", # Multi-signature encrypted wallet storage path
   "BtcEncryptedPrivateKeyPwd": "123",
   "BtcVendorSigningToolConfFile": "./vendor_tool.json", # vendor configuration file
   "BtcFeeRate": 10, # Multi-signature rate, when BTC returns to Bitcoin, unlock transaction rate
   "BtcMinChange": 8000, # Multi-sign the minimum UTXO value allowed in the Poly chain contract
   "BtcMinOutputValFromContract": 10000, # The minimum allowable withdrawal amount in the BTCX contract
   "BtcSignerPrivateKey": "cRRMYvoHPN...MVwyqZVrAcX", # BTC simulates the private key for sending transactions
   "BtcExistingVendorPrivks": "cREJsmv4W9Lr4Qh...wRoSnpxryGRL92N,cVJqF57c...PCwBDymEk,cTfnpP7C...pvWthPRoB3yGwR", # Existing multi-signature wallet private key
   ###
   
   ###
   # eth
   "EthURL": "http://ip:port", # Ethereum node
   "ETHPrivateKey": "AEC101ECD...BE90B2A2608A9", # Etherem simulates the private key of the transaction
   ###
   
   ###
   # Ontology
   "OntJsonRpcAddress": "http://ip:port", # Ontology
   "OntWallet": "./wallet.dat",
   "OntWalletPassword": "pwd",
   "GasPrice": 2500,
   "GasLimit": 30000000,
   "OntContractsAvmPath": "./avm", # Avm of all ontology contracts
   ###
   
   ###
   # cosmos
   "CMWalletPath": "./cosmos_key",#
   "CMWalletPwd": "pwd",
   "CMRpcUrl": "http://ip:port",
   "CMChainId": "cosmos-gaia",
   "CMGasPrice": "0.00001stake",
   "CMGas": 200000,
   "CMCrossChainId": 1000,
   "CMEpoch": 1,
   ###
   
   ###
   # NEO
   "NeoUrl": "http://seed9.ngd.network:20332",
	"NeoWif": "",
	"NeoEpoch": 0,
   ###
   
   ###
   # Poly
   "RCWallet": "./wallet.dat",
   "RCWalletPwd": "pwd",
   "RchainJsonRpcAddress": "http://ip:port",
   "RCEpoch": 60000,
	 ###
	 
	 ###
	 # configuration for sending tx
   "ReportInterval": 10,
   "ReportDir": "./report", # Test case status report
   "BatchTxNum": 100, # Batch send transaction: how many batches
   "BatchInterval": 1, # Time interval for batch sending
   "TxNumPerBatch": 100, # How many transactions per batch
   ###
   
   ###
   # This part is the contract configuration for each chain. 
   # After deployment, it is automatically written.
   # follows are contracts for testnet
   "EthErc20": "0x276788aF4a803781267c84692416311DE1F761f9",
	"EthOep4": "0x3105A14F7956D33a51F12eF3AE50A3f1eF161Dff",
	"Eccd": "0xA38366d552672556CE82426Da5031E2Ae0598dcD",
	"Eccm": "0x726532586C50ec9f4080B71f906a3d9779bbd64F",
	"Eccmp": "0xb600c8a2e8852832B75DB9Da1A3A1c173eAb28d8",
	"EthLockProxy": "0xD8aE73e06552E270340b63A8bcAbf9277a1aac99",
	"EthOngx": "0x42d9feF0Cbd9c3000CECe9764d99A4a6fE9E1B34",
	"EthOntx": "0x530aae4C0859894023906e28467f2a7F111B6ff3",
	"EthOntd": "0x76130c293AA35bf7B3e5fED1E9aE1E5DF12C6A92",
	"EthRenBTC": "0x239100e629a9Ca8e0BF45C7892b0fc72d78AA97A",
	"EthUSDT": "0xad3f96ae966ad60347f31845b7e4b333104c52fb",
	"EthWBTC": "0x557563dc4ed3fd256eBA55B9622f53331ab97c2f",
	"EthDai": "0x8Cad2301F7348DFc10C65778197028F432d51e76",
	"EthUSDC": "0x0d9c8723b343a8368bebe0b5e89273ff8d712e3c",
	"EthNeo": "0x7E269f2f33A97C64192e9889FAeEC72A6fcdB397",
	"BtceContractAddress": "0x92705a16815A3d1AEC3cE9Cc273C5aa302961FcC",
	"OntErc20": "e930755b130dccb25dc3cfee2b2e30d9370c1a75",
	"OntOep4": "969850e009b5e2a061694f3479ec8e44bc68bcd3",
	"OntLockProxy": "33c439c502cb4b6ac5a1e8057a65fe1fa7c300e2",
	"OntEth": "7009a2f7c8a2e45fa386a6078c7bfeaf518be487",
	"OntUSDT": "c6f91c11d740d39943b99a6b1c6fd2b5f476e2a3",
	"OntWBTC": "aede525f05065306423a5522bfcd31b5847ffa52",
	"OntDai": "96cf88356123592835a2fa75068a242260be1791",
	"OntUSDC": "07a12c0a6bdce4df04ef4b2045d1b0fd63a56e25",
	"OntNeo": "13eef3e184d878038317d806796b3af2d9f9b36d",
	"OntONTD": "869951e3397550e800d5faf579857cdb637a0051",
	"OntRenBTC": "ec547bc4486dea97cb659f1fe73407922f9e63c8",
	 "BtcoContractAddress": "814d32455c21bfc25c33b75ccbfc34fe8e79bff1",
	 "NeoCCMC": "0xe1695b1314a1331e3935481620417ed835669407",
	 "NeoLockProxy": "0x229e8fe772d5f0cbb408d58b593e42a1d1dfa3a9",
	 "CNeo": "c074a05e9dcf0141cbe6b4b3475dd67baf4dcb60",
	 "NeoOnt": "5a9222225f1bdb135123b74354c7248200c440aa",
	 "NeoOntd": "5a9222225f1bdb135123b74354c7248200c440aa",
	 "NeoEth": "d7b32de37ad906df80805c2419ff5560d20f9cbf",
	 "CMLockProxy": "f71b55ef55cedc91fd007f7a9ba386ec978f3aa8",
   ###
   
   ###
   # Test transaction limit. The cross-chain amount would not be over the limit.
   "BtcValLimit": 100000,
   "OntValLimit": 10,
   "OntdValLimit": 1000,
   "OngValLimit": 10000,
   "EthValLimit": 100000,
   "Oep4ValLimit": 1000,
   "Erc20ValLimit": 1000,
   "USDTValLimit": 10000,
   "NeoValLimit": 10000,
   "USDCValLimit": 1000,
   "WBTCValLimit": 1000,
   "RenBTCValLimit": 1000,
   "OntdValFloor": 100
   ###
}
```

## Send Transactions To Testnet

Build the `cctest` like follow:

```
go build -o cctest cmd/cctest/main.go
```

You can run a testcase like: 

```
./cctest -cfg=your_config_file -t case_name
```

Some cases:

| Case Name          | Desc                                                         |
| ------------------ | ------------------------------------------------------------ |
| SendOntToEthChain  | Send ONT to ethereum. Contract `EthOntx` will receive your ONT. |
| SendOnteToOntChain | Send ONT back to ontology. It would transfer from `EthOntx` to Ontology. |
| SendEthToOntChain  | Send ETH to ontology. Contract `OntEth` will receive your ETH. |
| SendEthoToEthChain | Send ETH back to ethereum.                                   |
| SendBtcToOntChain  | Send BTC to ontology. Contract `BtcoContractAddress` will mint a reflection coin BTCX for you. |

More case see [here](https://github.com/polynetwork/poly-io-test/blob/master/testcase/init.go).

If you want more details about sending cross-chain transactions, you can read these [documents](https://github.com/polynetwork/docs/tree/master/examples)

