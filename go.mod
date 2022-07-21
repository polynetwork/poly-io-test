module github.com/polynetwork/poly-io-test

go 1.14

require (
	github.com/ChainSQL/go-chainsql-api v1.1.0
	github.com/FISCO-BCOS/go-sdk v0.9.0
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/btcsuite/btcutil v1.0.2
	github.com/cosmos/cosmos-sdk v0.39.1
	github.com/ethereum/go-ethereum v1.9.15
	github.com/golang/protobuf v1.4.2
	github.com/google/go-cmp v0.5.2 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/hyperledger/fabric-protos-go v0.0.0-20200707132912-fee30f3ccd23
	github.com/hyperledger/fabric-sdk-go v1.0.0-rc1
	github.com/joeqian10/neo-gogogo v0.0.0-20200811090937-d8aab8600241
	github.com/ontio/ontology v1.11.1-0.20200827103103-591f7ce1009c
	github.com/ontio/ontology-crypto v1.0.9
	github.com/ontio/ontology-go-sdk v1.11.9-0.20200819065441-7c4b7a8330f6
	github.com/pkg/errors v0.9.1
	github.com/polynetwork/btc-vendor-tools v0.0.0-20200813091748-3b19a5fd7666
	github.com/polynetwork/cosmos-poly-module v0.0.0-20200810030259-95d586518759
	github.com/polynetwork/eth-contracts v0.0.0-20200903021827-c9212e419943
	github.com/polynetwork/poly v0.0.0-20201126065907-da2c5521739e
	github.com/polynetwork/poly-go-sdk v0.0.0-20201216023150-7ff89c0e43f9
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/tendermint v0.33.7
	github.com/tjfoc/gmsm v1.3.2-0.20200914155643-24d14c7bd05c
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)

replace (
	github.com/go-kit/kit v0.10.0 => github.com/go-kit/kit v0.8.0
	github.com/polynetwork/eth-contracts => github.com/zouxyan/eth-contracts v0.0.0-20210118060440-a0c11940f74c
)
