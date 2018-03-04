package geth

import (
    "hoqu-geth-api/contract"
    "github.com/ethereum/go-ethereum/common"
    "github.com/spf13/viper"
    "errors"
    "fmt"
    "math/big"
    "hoqu-geth-api/sdk/geth"
    "github.com/ethereum/go-ethereum/core/types"
)

var token *Token

type Token struct {
    *geth.Contract
    Token *contract.HoQuToken
}

func InitToken() error {
    c := geth.NewContract(viper.GetString("geth.addr.token"))
    c.InitEvents(contract.ClaimableCrowdsaleABI)

    t, err := contract.NewHoQuToken(c.Address, c.Wallet.Connection)
    if err != nil {
        return errors.New(fmt.Sprintf("Failed to instantiate a Token contract: %v", err))
    }

    token = &Token{
        Contract: c,
        Token:    t,
    }

    return nil
}

func (t *Token) Deploy(totSupply string) (*common.Address, *types.Transaction, error) {
    totSupplyEth, ok := big.NewInt(0).SetString(totSupply, 0)
    if !ok {
        return nil, nil, fmt.Errorf("wrong number provided: %s", totSupply)
    }

    address, tx, _, err := contract.DeployHoQuToken(
        t.Wallet.Account,
        t.Wallet.Connection,
        totSupplyEth,
    )
    if err != nil {
        return nil, nil, fmt.Errorf("failed to deploy HoQuToken contract: %v", err)
    }
    return &address, tx, nil
}

func GetToken() *Token {
    return token
}

func (t *Token) Balance(addr string) (*big.Int, error) {
    return t.Token.BalanceOf(nil, common.HexToAddress(addr))
}
