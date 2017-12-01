package geth

import (
    "hoqu-geth-api/contract"
    "github.com/ethereum/go-ethereum/common"
    "github.com/spf13/viper"
    "errors"
    "fmt"
    "math/big"
    "hoqu-geth-api/sdk/geth"
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

func GetToken() *Token {
    return token
}

func (t *Token) Balance(addr string) (*big.Int, error) {
    return t.Token.BalanceOf(nil, common.HexToAddress(addr))
}
