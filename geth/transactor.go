package geth

import (
    transactor "hoqu-geth-api/contract/ad"
    "github.com/ethereum/go-ethereum/common"
    "github.com/spf13/viper"
    "errors"
    "fmt"
    "hoqu-geth-api/sdk/geth"
    "github.com/ethereum/go-ethereum/core/types"
)

var ht *HoQuTransactor

type HoQuTransactor struct {
    *geth.Contract
    HoQuTransactor *transactor.HoQuTransactor
}

func initHoQuTransactor() error {
    c := geth.NewContract(viper.GetString("geth.addr.transactor"))
    c.InitEvents(transactor.HoQuTransactorABI)

    s, err := transactor.NewHoQuTransactor(c.Address, c.Wallet.Connection)
    if err != nil {
        return errors.New(fmt.Sprintf("Failed to instantiate a HoQu Platform Transactor contract: %v", err))
    }

    ht = &HoQuTransactor{
        Contract:       c,
        HoQuTransactor: s,
    }

    return nil
}

func GetHoQuTransactor() *HoQuTransactor {
    return ht
}

func (s *HoQuTransactor) Deploy() (*common.Address, *types.Transaction, error) {
    address, tx, _, err := transactor.DeployHoQuTransactor(
        s.Wallet.Account,
        s.Wallet.Connection,
        GetHoQuConfig().Address,
        GetToken().Address,
    )
    if err != nil {
        return nil, nil, fmt.Errorf("failed to deploy HoQuTransactor contract: %v", err)
    }
    return &address, tx, nil
}
