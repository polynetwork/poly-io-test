# CrossChain Test Framework


## 介绍

入口都在cmd/下：

| 入口           | 功能                                                         |
| -------------- | ------------------------------------------------------------ |
| btc_prepare    | 生成BTC多签地址等信息（也可使用现有的多签），部署并绑定各链上的BTCX合约，向Poly注册BTCX合约、交易参数。 |
| eth_deployer   | 两个功能：部署以太链上所有的合约，从ECCM到各个资产；设置合约的绑定，需要其他链先完成部署，确保合约hash已经填到config文件里。 |
| ont_deployer   | 同上                                                         |
| cosmos_prepare | 初始化gaia链，创建各个资产，完成资产绑定。                   |
| cctest         | 执行各种case，发送交易。                                     |
| tools          | 向Poly注册侧链，同步创世区块头等工作。                       |



## 配置

如下为配置文件，涉及各个端口、钱包等：

```
{
   ###
   # 这是BTC相关设置
   "BtcRestAddr": "http://172.168.3.10:20336", # BTC节点
   "BtcRestUser": "test",
   "BtcRestPwd": "test",
   "BtcFee": 1500, 
   "BtcRedeem": "522103c4564b837674de2482961a8d5f2a24a7e11e8a97aac5e92ac2e64500219144512102ccc07d3df7da58bb6fa5cfe5d7be415ff9463171b2600c93c080fcd0d49576a721036ec6299c1b14e57b45f1ad85eecbc48ad5447a05158a1bfb2ffb689ad69490d353ae", # 多签Redeem脚本
   "BtcNetType": "test", # 网络类型
   "BtcMultiSigNum": 3, # 多签总数
   "BtcMultiSigRequire": 2, # 多签最少签名数
   "BtcEncryptedPrivateKeyFile": "/data/poly_deployer/lib/vendor_tool/btcprivk", # 多签加密钱包存储路径，vendor_tool
   "BtcEncryptedPrivateKeyPwd": "123",
   "BtcVendorSigningToolConfFile": "/data/poly_deployer/lib/vendor_tool/vendor_tool.json", # vendor配置文件
   "BtcFeeRate": 10, # 多签的费率，BTC回比特币的时候，解锁交易的费率
   "BtcMinChange": 8000, # 多签在Poly链合约里允许的最小UTXO值
   "BtcMinOutputValFromContract": 10000, # BTCX合约里允许的最小出金金额
   "BtcSignerPrivateKey": "cRRMYvoHPNQu1tCz4ajPxytBVc2SN6GWLAVuyjzm4MVwyqZVrAcX", # BTC模拟发交易用的私钥
   "BtcExistingVendorPrivks": "cREJsmv4W9Lr4QhvnQrA77tTUZ5g488qNr2cvwRoSnpxryGRL92N,cVJqF57cdiSkFwpnTqEzjX7hCxZH95vxwenuJtBPgABPCwBDymEk,cTfnpP7CBFQ5mtK2BKPJLUnRJBYgnnRX7M8MnWpvWthPRoB3yGwR", # 现有的多签钱包私钥
   ###
   
   ###
   # eth
   "EthURL": "http://18.139.17.85:10331", # 以太节点
   "ETHPrivateKey": "AEC101ECDB5C86931E0CA5E635824F0D0F05240760C01DDEE64BE90B2A2608A9", # 以太模拟发交易的私钥
   ###
   
   ###
   # 本体
   "OntJsonRpcAddress": "http://172.168.3.73:20336", # 本体
   "OntWallet": ".wallets/ont_dev.dat",
   "OntWalletPassword": "admin",
   "GasPrice": 0,
   "GasLimit": 30000000,
   "OntContractsAvmPath": "./chains/ont/avm", # 所有本体合约的avm
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
   "ReportDir": "./report", # 测试用例的状态报告
   "BatchTxNum": 100, # 批量发交易：多少批
   "BatchInterval": 1, # 批量发的时间间隔
   "TxNumPerBatch": 100, # 每批多少交易
   ###
   
   ###
   # 这一部分是各个链的合约配置，部署完后，自动写入
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
   # 测试交易金额上限
   "BtcValLimit": 100000,
   "OntValLimit": 100,
   "OngValLimit": 100000,
   "EthValLimit": 10000000,
   "Oep4ValLimit": 10000000,
   "Erc20ValLimit": 10000000
   ###
}
```



## 开发

如果有新的链加入，需要开发对应的合约部署、配置，以及发交易的功能，主要从以下几个目录入手。

- chains/

  每条链的合约部署、调用、交易发送等方法写在这里

- cmd/

  各条链的功能入口写在这里，往往加一条链，其他链的代码也要修改，比如部署和绑定合约

- config/

  框架配置，合约端口等信息都写在里面。

- testcase

  发交易测试的用例写在这里。