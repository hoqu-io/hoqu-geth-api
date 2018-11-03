package geth

import (
    "hoqu-geth-api/contract/rater"
    "github.com/ethereum/go-ethereum/common"
    "github.com/spf13/viper"
    "errors"
    "fmt"
    "hoqu-geth-api/sdk/geth"
    "github.com/ethereum/go-ethereum/core/types"
)

var r *HoQuRater

type HoQuRater struct {
    *geth.Contract
    HoQuRater *rater.HoQuRater
}

func initHoQuRater() error {
    c := geth.NewContract(viper.GetString("geth.addr.rater"))
    c.InitEvents(rater.HoQuRaterABI)

    s, err := rater.NewHoQuRater(c.Address, c.Wallet.Connection)
    if err != nil {
        return errors.New(fmt.Sprintf("Failed to instantiate a HoQu Platform Rater contract: %v", err))
    }

    r = &HoQuRater{
        Contract:  c,
        HoQuRater: s,
    }

    return nil
}

func GetHoQuRater() *HoQuRater {
    return r
}

func (s *HoQuRater) Deploy() (*common.Address, *types.Transaction, error) {
    address, tx, _, err := rater.DeployHoQuRater(
        s.Wallet.Account,
        s.Wallet.Connection,
        GetHoQuConfig().Address,
        GetHoQuStorage().Address,
    )
    if err != nil {
        return nil, nil, fmt.Errorf("failed to deploy HoQuRater contract: %v", err)
    }
    return &address, tx, nil
}
