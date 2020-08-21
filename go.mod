module github.com/polynetwork/poly-io-test

go 1.14

require (
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/btcsuite/btcutil v1.0.2
	github.com/cosmos/cosmos-sdk v0.39.1
	github.com/ethereum/go-ethereum v1.9.15
	github.com/google/go-cmp v0.5.2 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/joeqian10/neo-gogogo v0.0.0-20200811090937-d8aab8600241
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/ontio/ontology v1.11.1-0.20200827103103-591f7ce1009c
	github.com/ontio/ontology-crypto v1.0.9
	github.com/ontio/ontology-go-sdk v1.11.9-0.20200819065441-7c4b7a8330f6
	github.com/polynetwork/btc-vendor-tools v0.0.0-20200813091748-3b19a5fd7666
	github.com/polynetwork/cosmos-poly-module v0.0.0-20200810030259-95d586518759
	github.com/polynetwork/eth-contracts v0.0.0-20200903021827-c9212e419943
	github.com/polynetwork/poly v0.0.0-20201216061550-50185057319d
	github.com/polynetwork/poly-go-sdk v0.0.0-20200817120957-365691ad3493
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/tendermint v0.33.7
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	golang.org/x/net v0.0.0-20200904194848-62affa334b73 // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/check.v1 v1.0.0-20200902074654-038fdea0a05b // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)

replace github.com/polynetwork/eth-contracts => github.com/zouxyan/eth-contracts v0.0.0-20201215112041-6532ab4e12f4

replace github.com/polynetwork/poly => github.com/zhiqiangxu/poly v0.0.0-20201216075543-df86073d2575

replace github.com/polynetwork/poly-go-sdk => github.com/zhiqiangxu/poly-go-sdk v0.0.0-20201215024222-5d728b68e651
