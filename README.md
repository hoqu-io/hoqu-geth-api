# HOQU blockchain API by GoLang

Provides all projects the opportunity to access Ethereum blockchain through RESTful API.

## Installation

Install [GoLang](https://www.goinggo.net/2016/05/installing-go-and-your-workspace.html) and setup local environment

### Install native Ethereum go bindings

```console
go get github.com/ethereum/go-ethereum
```

### Install govendor

```console
go get -u github.com/kardianos/govendor
```

Sync vendor packages

```console
govendor sync
```

## Building

Build Geth and utils

```console
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make all
```

## Development

### Generate new go bindings for solidity contracts

If you are going to modify solidity contracts or contracts are changed in main solidity repository follow next steps.

Install [Solidity compiler](http://solidity.readthedocs.io/en/develop/installing-solidity.html)

copy HOQU contracts to sol

```console
cd $GOPATH/src/hoqu-geth-api
rsync -vax vendor/github.com/hoqu-io/platform/contracts/ sol/
```

copy zeppelin contracts to sol

```console
cd $GOPATH/src/hoqu-geth-api
rsync -vax vendor/github.com/OpenZeppelin/zeppelin-solidity/contracts/ sol/zeppelin-solidity/contracts/
```

Generate go bindings from solidity contract (if sol sources modified)

```console
$GOPATH/src/github.com/ethereum/go-ethereum/build/bin/abigen --sol sol/ClaimableCrowdsale.sol --pkg=contract --out=contract/ClaimableCrowdsale.go
```

## Execution (production with own node)

Run Geth node

```console
cd $GOPATH/src/github.com/ethereum/go-ethereum/
./build/bin/geth
```

Use IPC endpoint provided by geth in your config!!!