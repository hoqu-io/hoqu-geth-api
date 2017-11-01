package geth

import (
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "fmt"
    "errors"
    "strings"
    "io/ioutil"
    "math/big"
    "github.com/ethereum/go-ethereum/common"
    "context"
)

var wallet *Wallet

type Wallet struct {
    Connection *ethclient.Client
    Account *bind.TransactOpts
}

func InitWallet(endpoint string, keyFile string, keyPass string) error {
    conn, err := ethclient.Dial(endpoint)
    if err != nil {
        return errors.New(fmt.Sprintf("Failed to connect to the Ethereum network: %v", err))
    }

    keyJson, err := ioutil.ReadFile(keyFile)
    if err != nil {
        return err
    }

    acc, err := bind.NewTransactor(strings.NewReader(string(keyJson)), keyPass)
    if err != nil {
        return errors.New(fmt.Sprintf("Failed to authorize main account: %v", err))
    }

    wallet = &Wallet{
        Connection: conn,
        Account: acc,
    }

    return nil
}

func GetWallet() *Wallet {
    return wallet
}

func (w *Wallet) BalanceAt(addr string) (*big.Int, error) {
    return w.Connection.BalanceAt(context.TODO(), common.HexToAddress(addr), nil)
}
