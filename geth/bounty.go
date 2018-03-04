package geth

import (
    "hoqu-geth-api/contract"
    "github.com/ethereum/go-ethereum/common"
    "github.com/spf13/viper"
    "errors"
    "fmt"
    "math/big"
    "hoqu-geth-api/sdk/geth"
    "hoqu-geth-api/geth/models"
    sdkModels "hoqu-geth-api/sdk/models"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/sirupsen/logrus"
)

var bounty *Bounty

type Bounty struct {
    *geth.Contract
    Bounty *contract.HoQuBounty
}

func InitBounty() error {
    c := geth.NewContract(viper.GetString("geth.addr.bounty"))
    c.InitEvents(contract.HoQuBountyABI)

    s, err := contract.NewHoQuBounty(c.Address, c.Wallet.Connection)
    if err != nil {
        return errors.New(fmt.Sprintf("Failed to instantiate a Bounty contract: %v", err))
    }

    bounty = &Bounty{
        Contract: c,
        Bounty:     s,
    }

    return nil
}

func GetBounty() *Bounty {
    return bounty
}

func (s *Bounty) Deploy(params *models.BountyDeployParams) (*common.Address, *types.Transaction, error) {
    tokenAddr := GetToken().Address

    address, tx, _, err := contract.DeployHoQuBounty(
        s.Wallet.Account,
        s.Wallet.Connection,
        tokenAddr,
        common.HexToAddress(params.BankAddress),
        common.HexToAddress(params.BeneficiaryAddress),
    )
    if err != nil {
        return nil, nil, fmt.Errorf("failed to deploy contract: %v", err)
    }
    return &address, tx, nil
}

func (s *Bounty) Balance(addr string) (*big.Int, error) {
    return s.Bounty.Tokens(nil, common.HexToAddress(addr))
}

func (s *Bounty) Summary() (*models.BountySummary, error) {
    issued, err := s.Bounty.IssuedTokensAmount(nil)
    if err != nil {
        return nil, err
    }

    receivers, err := s.Bounty.GetReceiversCount(&bind.CallOpts{
        From: s.Wallet.Account.From,
    })
    if err != nil {
        return nil, err
    }

    isFin, err := s.Bounty.IsFinished(&bind.CallOpts{
        From: s.Wallet.Account.From,
    })
    if err != nil {
        return nil, err
    }

    return &models.BountySummary{
        IssuedTokensAmount: issued.String(),
        ReceiversCount:     receivers,
        IsFinished:         isFin,
    }, nil
}

func (s *Bounty) Approved(addr string) (bool, error) {
    return s.Bounty.Approved(nil, common.HexToAddress(addr))
}

func (s *Bounty) Events(addrs []string) ([]sdkModels.ContractEvent, error) {
    hashAddrs := make([]common.Hash, len(addrs))
    for _, addr := range addrs {
        hashAddrs = append(hashAddrs, common.HexToHash(addr))
    }

    events, err := s.GetEventsByTopics(
        [][]common.Hash{{}, hashAddrs},
        big.NewInt(viper.GetInt64("geth.start_block.bounty")),
    )
    if err != nil {
        return nil, err
    }

    for key, event := range events {
        switch {
        case event.Name == "TokenBought" || event.Name == "TokenAdded" || event.Name == "TokenToppedUp" || event.Name == "TokenSubtracted":
            events[key].Args = models.TokenAddedEventArgs{
                Address:     common.BytesToAddress(event.RawArgs[0]).String(),
                TokenAmount: common.BytesToHash(event.RawArgs[1]).Big().String(),
                EtherAmount: common.BytesToHash(event.RawArgs[2]).Big().String(),
            }
        case event.Name == "TokenSent" || event.Name == "TokenAddedByBounty":
            events[key].Args = models.TokenSentEventAgrs{
                Address:     common.BytesToAddress(event.RawArgs[0]).String(),
                TokenAmount: common.BytesToHash(event.RawArgs[1]).Big().String(),
            }
        default:
            return nil, fmt.Errorf("unknown event type: %s", event.Name)
        }
    }

    return events, nil
}

func (s *Bounty) AddByBounty(addr string, tokens string) (common.Hash, error) {
    tokensAmount, ok := big.NewInt(0).SetString(tokens, 0)
    if !ok {
        return common.Hash{}, fmt.Errorf("wrong number provided: %s", tokens)
    }

    opts, err := s.Wallet.GetTransactOpts()
    if err != nil {
        s.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := s.Bounty.AddByBounty(opts, common.HexToAddress(addr), tokensAmount)
    if err != nil {
        s.Wallet.OnFailTransaction(err)

        if s.Wallet.ValidateRepeatableTransaction(err) {
            logrus.Warn("Repeat AddByBounty to ", addr)

            return s.AddByBounty(addr, tokens)
        }

        return common.Hash{}, err
    }
    s.Wallet.OnSuccessTransaction()

    return tx.Hash(), nil
}

func (s *Bounty) Approve(addr string) (common.Hash, error) {
    opts, err := s.Wallet.GetTransactOpts()
    if err != nil {
        s.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := s.Bounty.Approve(opts, common.HexToAddress(addr))
    if err != nil {
        s.Wallet.OnFailTransaction(err)

        if s.Wallet.ValidateRepeatableTransaction(err) {
            logrus.Warn("Repeat Approve to ", addr)

            return s.Approve(addr)
        }

        return common.Hash{}, err
    }
    s.Wallet.OnSuccessTransaction()

    return tx.Hash(), nil
}
