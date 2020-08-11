GOFMT=gofmt
GC=go build
PWD := $(shell pwd)

ARCH=$(shell uname -m)
SRC_FILES = $(shell git ls-files | grep -e .go$ | grep -v _test.go)

cct: $(SRC_FILES)
	CGO_ENABLED=1 $(GC) -o cct  cmd/cctest/main.go

cct-windows:
	GOOS=windows GOARCH=amd64 $(GC) -o cct-windows-amd64.exe cmd/cctest/main.go
cct-linux:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 $(GC) -o cct-linux-amd64 cmd/cctest/main.go
cct-mac:
	GOOS=darwin GOARCH=amd64 $(GC)  -o ccct-darwin-amd64 cmd/cctest/main.go
cct-btc-prepare:
	GOOS=linux GOARCH=amd64 $(GC) -o cct-btc-linux-amd64 cmd/btc_prepare/run.go
cct-eth-deployer:
	GOOS=linux GOARCH=amd64 $(GC)  -o cct-eth-linux-amd64 cmd/eth_deployer/run.go
cct-cosmos-deployer:
	GOOS=linux GOARCH=amd64 $(GC)  -o cct-cosmos-linux-amd64 cmd/cosmos_prepare/run.go
cct-ont-deployer:
	GOOS=linux GOARCH=amd64 $(GC)  -o cct-ont-linux-amd64 cmd/ont_deployer/run.go

format:
	$(GOFMT) -w cmd/cctest/main.go

clean:
	rm -rf *.8 *.o *.out *.6 *exe coverage
	rm -rf cct-*
