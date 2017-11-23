package geth

import (
    "hoqu-api/contract"
    "github.com/ethereum/go-ethereum/common"
    "github.com/spf13/viper"
    "errors"
    "fmt"
    "hoqu-api/sdk/geth"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var privPlace *PrivatePlacement

type PrivatePlacement struct {
    *geth.Contract
    PrivatePlacement *contract.PrivatePlacement
}

func InitPrivatePlacement() error {
    c := geth.NewContract(viper.GetString("geth.addr.privPlace"))
    c.InitEvents(contract.ClaimableCrowdsaleABI)

    pp, err := contract.NewPrivatePlacement(c.Address, c.Wallet.Connection)
    if err != nil {
        return errors.New(fmt.Sprintf("Failed to instantiate a PrivatePlacement contract: %v", err))
    }

    privPlace = &PrivatePlacement{
        Contract:         c,
        PrivatePlacement: pp,
    }

    return nil
}

func GetPrivatePlacement() *PrivatePlacement {
    return privPlace
}

func (t *PrivatePlacement) TokenAddr() (common.Address, error) {
    return t.PrivatePlacement.Token(&bind.CallOpts{
        From: t.Wallet.Account.From,
    })
}
