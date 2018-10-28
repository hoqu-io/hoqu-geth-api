# HOQU blockchain API by GoLang

Provides all projects the opportunity to access Ethereum blockchain through RESTful API.

## Installation

Install [GoLang](https://www.goinggo.net/2016/05/installing-go-and-your-workspace.html) and setup local environment
Install [Ethereum command line tools](https://www.ethereum.org/cli)

### Install govendor

```console
go get -u github.com/kardianos/govendor
```

Sync vendor packages

```console
govendor sync
```

Build HOQU API

```console
cd $GOPATH/src/hoqu-geth-api
go build
```

## Execution (test node via rinkeby.infura.io)

Run API

```console
./hoqu-geth-api
```

## Execution (production with own node)

Run Geth node

Use IPC endpoint provided by geth in your pro_net config (geth -> endpoint)

Run API

```console
APPLICATION_ENV=pro_net ./hoqu-geth-api
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
rsync -vax vendor/github.com/OpenZeppelin/zeppelin-solidity/contracts/ sol/zeppelin/
```

Install native Ethereum go node

```console
go get github.com/ethereum/go-ethereum
```

Build Geth and utils

```console
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make all
```

Generate go bindings from solidity contract (if sol sources modified)

```console
$GOPATH/src/github.com/ethereum/go-ethereum/build/bin/abigen --sol sol/HoQuPlatform.sol --pkg=platform --out=contract/platform/HoQuPlatform.go
```

Alternatively you can generate go bindings from assembled solidity contract

```console
$GOPATH/src/hoqu-geth-api/contract/bin/abigen --sol sol/HoQuPlatform.sol --pkg=platform --out=contract/platform/HoQuPlatform.go
$GOPATH/src/hoqu-geth-api/contract/bin/abigen --sol sol/HoQuStorage.sol --pkg=platform --out=contract/storage/HoQuStorage.go
$GOPATH/src/hoqu-geth-api/contract/bin/abigen --sol sol/HoQuConfig.sol --pkg=platform --out=contract/config/HoQuConfig.go
$GOPATH/src/hoqu-geth-api/contract/bin/abigen --sol sol/HoQuAdCampaign.sol --pkg=ad --out=contract/ad/HoQuAdCampaign.go
$GOPATH/src/hoqu-geth-api/contract/bin/abigen --sol sol/HoQuRater.sol --pkg=rater --out=contract/rater/HoQuRater.go
```

### Generate API documentation

Install [go-swagger](https://github.com/go-swagger/go-swagger)

Generate swagger.json

```console
cd $GOPATH/src/hoqu-geth-api
swagger generate spec -o ./data/docs/swagger.json
```
