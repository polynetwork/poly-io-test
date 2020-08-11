module github.com/polynetwork/cross_chain_test

go 1.14

require (
	github.com/FactomProject/basen v0.0.0-20150613233007-fe3947df716e // indirect
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/btcsuite/btcutil v1.0.2
	github.com/cosmos/cosmos-sdk v0.38.4
	github.com/ethereum/go-ethereum v1.9.15
	github.com/ontio/ontology v1.11.0
	github.com/ontio/ontology-crypto v1.0.9
	github.com/ontio/ontology-go-sdk v1.11.4
	github.com/polynetwork/cosmos-poly-module v0.0.0-20200624032945-f3e838f0d6c3
	github.com/polynetwork/poly v0.0.0-20200715030435-4f1d1a0adb44
	github.com/polynetwork/poly-go-sdk v0.0.0-20200730112529-d9c0c7ddf3d8
	github.com/polynetwork/vendortool v0.0.0-00010101000000-000000000000
	github.com/tendermint/tendermint v0.33.3
)

replace (
	github.com/polynetwork/vendortool => ../vendortool
	github.com/polynetwork/cosmos-poly-module => github.com/skyinglyh1/cosmos-poly-module v0.0.0-20200715085758-6038c006fe79
)
