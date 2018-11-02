package geth

import (
    "hoqu-geth-api/contract/burner"
    "github.com/ethereum/go-ethereum/common"
    "github.com/spf13/viper"
    "errors"
    "fmt"
    "math/big"
    "hoqu-geth-api/sdk/geth"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/sirupsen/logrus"
)

var burn *Burner

type Burner struct {
    *geth.Contract
    Burner *burner.HoQuBurner
}

func InitBurner() error {
    c := geth.NewContract(viper.GetString("geth.addr.burner"))

    s, err := burner.NewHoQuBurner(c.Address, c.Wallet.Connection)
    if err != nil {
        return errors.New(fmt.Sprintf("Failed to instantiate a Burner contract: %v", err))
    }

    burn = &Burner{
        Contract: c,
        Burner:   s,
    }

    return nil
}

func GetBurner() *Burner {
    return burn
}

func (s *Burner) Deploy() (*common.Address, *types.Transaction, error) {
    tokenAddr := GetToken().Address

    address, tx, _, err := burner.DeployHoQuBurner(
        s.Wallet.Account,
        s.Wallet.Connection,
        tokenAddr,
    )
    if err != nil {
        return nil, nil, fmt.Errorf("failed to deploy contract: %v", err)
    }
    return &address, tx, nil
}

func (s *Burner) BurnFrom(addr string, tokens string) (common.Hash, error) {
    tokensAmount, ok := big.NewInt(0).SetString(tokens, 0)
    if !ok {
        return common.Hash{}, fmt.Errorf("wrong number provided: %s", tokens)
    }

    opts, err := s.Wallet.GetTransactOpts()
    if err != nil {
        s.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := s.Burner.BurnFrom(opts, common.HexToAddress(addr), tokensAmount)
    if err != nil {
        s.Wallet.OnFailTransaction(err)

        if s.Wallet.ValidateRepeatableTransaction(err) {
            logrus.Warn("Repeat BurnFrom to ", addr)

            return s.BurnFrom(addr, tokens)
        }

        return common.Hash{}, err
    }
    s.Wallet.OnSuccessTransaction()

    return tx.Hash(), nil
}
