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

var sale *Sale

type Sale struct {
    *geth.Contract
    Sale *contract.ChangeableRateCrowdsale
}

func InitSale() error {
    c := geth.NewContract(viper.GetString("geth.addr.sale"))
    c.InitEvents(contract.ChangeableRateCrowdsaleABI)

    s, err := contract.NewChangeableRateCrowdsale(c.Address, c.Wallet.Connection)
    if err != nil {
        return errors.New(fmt.Sprintf("Failed to instantiate a Sale contract: %v", err))
    }

    sale = &Sale{
        Contract: c,
        Sale:     s,
    }

    return nil
}

func GetSale() *Sale {
    return sale
}

func (s *Sale) Deploy(params *models.SaleDeployParams) (*common.Address, *types.Transaction, error) {
    tokenAddr := GetToken().Address

    tokenRate, ok := big.NewInt(0).SetString(params.TokenRate, 0)
    if !ok {
        return nil, nil, fmt.Errorf("wrong TokenRate provided: %s", params.TokenRate)
    }

    minBuyableAmount, ok := big.NewInt(0).SetString(params.MinBuyableAmount, 0)
    if !ok {
        return nil, nil, fmt.Errorf("wrong MinBuyableAmount provided: %s", params.MinBuyableAmount)
    }

    maxTokensAmount, ok := big.NewInt(0).SetString(params.MaxTokensAmount, 0)
    if !ok {
        return nil, nil, fmt.Errorf("wrong MaxTokensAmount provided: %s", params.MaxTokensAmount)
    }

    address, tx, _, err := contract.DeployChangeableRateCrowdsale(
        s.Wallet.Account,
        s.Wallet.Connection,
        tokenAddr,
        common.HexToAddress(params.BankAddress),
        common.HexToAddress(params.BeneficiaryAddress),
        tokenRate,
        minBuyableAmount,
        maxTokensAmount,
        big.NewInt(params.EndDate),
    )
    if err != nil {
        return nil, nil, fmt.Errorf("failed to deploy contract: %v", err)
    }
    return &address, tx, nil
}

func (s *Sale) Balance(addr string) (*big.Int, error) {
    return s.Sale.Tokens(nil, common.HexToAddress(addr))
}

func (s *Sale) Summary() (*models.SaleSummary, error) {
    max, err := s.Sale.MaxTokensAmount(nil)
    if err != nil {
        return nil, err
    }

    issued, err := s.Sale.IssuedTokensAmount(nil)
    if err != nil {
        return nil, err
    }

    nextBoundary, err := s.Sale.NextBoundaryAmount(nil)
    if err != nil {
        return nil, err
    }

    rate, err := s.Sale.TokenRate(nil)
    if err != nil {
        return nil, err
    }

    receivers, err := s.Sale.GetReceiversCount(&bind.CallOpts{
        From: s.Wallet.Account.From,
    })
    if err != nil {
        return nil, err
    }

    isFin, err := s.Sale.IsFinished(&bind.CallOpts{
        From: s.Wallet.Account.From,
    })
    if err != nil {
        return nil, err
    }

    return &models.SaleSummary{
        Address:            s.Address.String(),
        MaxTokensAmount:    max.String(),
        IssuedTokensAmount: issued.String(),
        NextBoundaryAmount: nextBoundary.String(),
        TokenRate:          rate.String(),
        ReceiversCount:     receivers,
        IsFinished:         isFin,
    }, nil
}

func (s *Sale) Approved(addr string) (bool, error) {
    return s.Sale.Approved(nil, common.HexToAddress(addr))
}

func (s *Sale) Events(addrs []string) ([]sdkModels.ContractEvent, error) {
    hashAddrs := make([]common.Hash, len(addrs))
    for _, addr := range addrs {
        hashAddrs = append(hashAddrs, common.HexToHash(addr))
    }

    events, err := s.GetEventsByTopics(
        [][]common.Hash{{}, hashAddrs},
        big.NewInt(viper.GetInt64("geth.start_block.sale")),
    )
    if err != nil {
        return nil, err
    }

    resEvents := make([]sdkModels.ContractEvent, 0)

    for _, event := range events {
        addr := common.BytesToAddress(event.RawArgs[0])

        switch {
        case event.Name == "TokenBought" || event.Name == "TokenAdded" || event.Name == "TokenToppedUp" ||
            event.Name == "TokenSubtracted":
            event.Args = models.TokenAddedEventArgs{
                Address:     addr.String(),
                TokenAmount: common.BytesToHash(event.RawArgs[1]).Big().String(),
                EtherAmount: common.BytesToHash(event.RawArgs[2]).Big().String(),
            }
        case event.Name == "TokenSent":
            event.Args = models.TokenSentEventAgrs{
                Address:     addr.String(),
                TokenAmount: common.BytesToHash(event.RawArgs[1]).Big().String(),
            }
        default:
            return nil, fmt.Errorf("unknown event type: %s", event.Name)
        }

        resEvents = append(resEvents, event)
    }

    return resEvents, nil
}

func (s *Sale) Add(addr string, eqEthAmount string) (common.Hash, error) {
    ethAmount, ok := big.NewInt(0).SetString(eqEthAmount, 0)
    if !ok {
        return common.Hash{}, fmt.Errorf("wrong number provided: %s", eqEthAmount)
    }

    opts, err := s.Wallet.GetTransactOpts()
    if err != nil {
        s.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := s.Sale.Add(opts, common.HexToAddress(addr), ethAmount)
    if err != nil {
        logrus.Error("Add failed ", err.Error())
        s.Wallet.OnFailTransaction(err)

        if s.Wallet.ValidateRepeatableTransaction(err) {
            logrus.Warn("Repeat Add to ", addr)

            return s.Add(addr, eqEthAmount)
        }

        return common.Hash{}, err
    }
    s.Wallet.OnSuccessTransaction()

    logrus.Info("Add to ", addr, " tx ", tx.Hash().String())

    return tx.Hash(), nil
}

func (s *Sale) TopUp(addr string, eqEthAmount string) (common.Hash, error) {
    return common.Hash{}, fmt.Errorf("disabled")

    ethAmount, ok := big.NewInt(0).SetString(eqEthAmount, 0)
    if !ok {
        return common.Hash{}, fmt.Errorf("wrong number provided: %s", eqEthAmount)
    }

    opts, err := s.Wallet.GetTransactOpts()
    if err != nil {
        s.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := s.Sale.TopUp(opts, common.HexToAddress(addr), ethAmount)
    if err != nil {
        s.Wallet.OnFailTransaction(err)

        if s.Wallet.ValidateRepeatableTransaction(err) {
            logrus.Warn("Repeat TopUp to ", addr)

            return s.TopUp(addr, eqEthAmount)
        }

        return common.Hash{}, err
    }
    s.Wallet.OnSuccessTransaction()

    logrus.Info("TopUp to ", addr, " tx ", tx.Hash().String())

    return tx.Hash(), nil
}

func (s *Sale) Approve(addr string) (common.Hash, error) {
    opts, err := s.Wallet.GetTransactOpts()
    if err != nil {
        s.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := s.Sale.Approve(opts, common.HexToAddress(addr))
    if err != nil {
        s.Wallet.OnFailTransaction(err)

        if s.Wallet.ValidateRepeatableTransaction(err) {
            logrus.Warn("Repeat Approve to ", addr)

            return s.Approve(addr)
        }

        return common.Hash{}, err
    }
    s.Wallet.OnSuccessTransaction()

    logrus.Info("Approve ", addr, " tx ", tx.Hash().String())

    return tx.Hash(), nil
}

func (s *Sale) ClaimAll() (common.Hash, error) {
    opts, err := s.Wallet.GetTransactOpts()
    if err != nil {
        s.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := s.Sale.ClaimAll(opts)
    if err != nil {
        s.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }
    s.Wallet.OnSuccessTransaction()

    logrus.Info("ClaimAll tx ", tx.Hash().String())

    return tx.Hash(), nil
}