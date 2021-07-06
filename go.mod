module github.com/polynetwork/poly-io-test

go 1.14

require (
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/Zilliqa/gozilliqa-sdk v1.2.1-0.20210406170702-178fd8166bae
	github.com/aristanetworks/goarista v0.0.0-20200331225509-2cc472e8fbd6 // indirect
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/btcsuite/btcutil v1.0.2
	github.com/cosmos/cosmos-sdk v0.39.2
	github.com/ethereum/go-ethereum v1.9.18
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/golang/mock v1.5.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/joeqian10/neo-gogogo v1.1.0
	github.com/joeqian10/neo3-gogogo v0.3.6
	github.com/kardiachain/go-kardia v1.1.1-0.20210518091640-d13a5b7f7c4c
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/ontio/ontology v1.11.1-0.20200827103103-591f7ce1009c
	github.com/ontio/ontology-crypto v1.0.9
	github.com/ontio/ontology-go-sdk v1.11.9-0.20200819065441-7c4b7a8330f6
	github.com/polynetwork/btc-vendor-tools v0.0.0-20200813091748-3b19a5fd7666
	github.com/polynetwork/cosmos-poly-module v0.0.0-20200810030259-95d586518759
	github.com/polynetwork/eth-contracts v0.0.0-20200903021827-c9212e419943
	github.com/polynetwork/kai-relayer v0.0.0
	github.com/polynetwork/poly v0.0.0-20210112063446-24e3d053e9d6
	github.com/polynetwork/poly-go-sdk v0.0.0-20200817120957-365691ad3493
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/viper v1.7.1 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.33.9
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)

replace github.com/polynetwork/eth-contracts => github.com/zouxyan/eth-contracts v0.0.0-20210115072359-e4cac6edc20c

replace github.com/polynetwork/poly => github.com/joeqian10/poly v0.0.0-20210513061019-474879d3ddcd

//replace github.com/polynetwork/poly-go-sdk => github.com/zhiqiangxu/poly-go-sdk v0.0.0-20201215024222-5d728b68e651

replace github.com/polynetwork/poly-go-sdk => github.com/joeqian10/poly-go-sdk v0.0.0-20210517072349-71002ebfdf13

//replace github.com/ontio/ontology => github.com/ontio/ontology v1.11.0

replace github.com/polynetwork/kai-relayer => github.com/dogecoindev/kai-relayer v0.0.0-20210609112229-34bf794e78e7
