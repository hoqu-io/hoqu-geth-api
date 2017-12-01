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
)

var presale *Presale

type Presale struct {
    *geth.Contract
    Presale *contract.ClaimableCrowdsale
}

func InitPresale() error {
    c := geth.NewContract(viper.GetString("geth.addr.presale"))
    c.InitEvents(contract.ClaimableCrowdsaleABI)

    psale, err := contract.NewClaimableCrowdsale(c.Address, c.Wallet.Connection)
    if err != nil {
        return errors.New(fmt.Sprintf("Failed to instantiate a Presale contract: %v", err))
    }

    presale = &Presale{
        Contract: c,
        Presale:  psale,
    }

    return nil
}

func GetPresale() *Presale {
    return presale
}

func (p *Presale) Deploy(params *models.PresaleDeployParams) (*common.Address, *types.Transaction, error) {
    var tokenAddr common.Address
    tokenAddr, err := GetPrivatePlacement().TokenAddr()
    if err != nil {
        return nil, nil, err
    }

    tokenRate, ok := big.NewInt(0).SetString(params.TokenRate, 0)
    if !ok {
        return nil, nil, fmt.Errorf("wrong number provided: %s", params.TokenRate)
    }

    minBuyableAmount, ok := big.NewInt(0).SetString(params.MinBuyableAmount, 0)
    if !ok {
        return nil, nil, fmt.Errorf("wrong number provided: %s", params.MinBuyableAmount)
    }

    maxTokensAmount, ok := big.NewInt(0).SetString(params.MaxTokensAmount, 0)
    if !ok {
        return nil, nil, fmt.Errorf("wrong number provided: %s", params.MaxTokensAmount)
    }

    address, tx, _, err := contract.DeployClaimableCrowdsale(
        p.Wallet.Account,
        p.Wallet.Connection,
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

func (p *Presale) Balance(addr string) (*big.Int, error) {
    return p.Presale.Tokens(nil, common.HexToAddress(addr))
}

func (p *Presale) Summary() (*models.Summary, error) {
    max, err := p.Presale.MaxTokensAmount(nil)
    if err != nil {
        return nil, err
    }

    issued, err := p.Presale.IssuedTokensAmount(nil)
    if err != nil {
        return nil, err
    }

    min, err := p.Presale.MinBuyableAmount(nil)
    if err != nil {
        return nil, err
    }

    rate, err := p.Presale.TokenRate(nil)
    if err != nil {
        return nil, err
    }

    receivers, err := p.Presale.GetReceiversCount(&bind.CallOpts{
        From: p.Wallet.Account.From,
    })
    if err != nil {
        return nil, err
    }

    isFin, err := p.Presale.IsFinished(&bind.CallOpts{
        From: p.Wallet.Account.From,
    })
    if err != nil {
        return nil, err
    }

    return &models.Summary{
        Address:            p.Address.String(),
        MaxTokensAmount:    max.String(),
        MinBuyableAmount:   min.String(),
        IssuedTokensAmount: issued.String(),
        TokenRate:          rate.String(),
        ReceiversCount:     receivers,
        IsFinished:         isFin,
    }, nil
}

func (p *Presale) Approved(addr string) (bool, error) {
    return p.Presale.Approved(nil, common.HexToAddress(addr))
}

func (p *Presale) Events(addrs []string) ([]sdkModels.ContractEvent, error) {
    hashAddrs := make([]common.Hash, len(addrs))
    for _, addr := range addrs {
        hashAddrs = append(hashAddrs, common.HexToHash(addr))
    }

    events, err := p.GetEventsByTopics([][]common.Hash{{}, hashAddrs})
    if err != nil {
        return nil, err
    }

    for key, event := range events {
        switch {
        case event.Name == "TokenBought" || event.Name == "TokenAdded" || event.Name == "TokenToppedUp" ||
        event.Name == "TokenSubtracted":
            events[key].Args = models.TokenAddedEventArgs{
                Address:     common.BytesToAddress(event.RawArgs[0]).String(),
                TokenAmount: common.BytesToHash(event.RawArgs[1]).Big().String(),
                EtherAmount: common.BytesToHash(event.RawArgs[2]).Big().String(),
            }
        case event.Name == "TokenSent":
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

func (p *Presale) Add(addr string, tokens string) (common.Hash, error) {
    tokensAmount, ok := big.NewInt(0).SetString(tokens, 0)
    if !ok {
        return common.Hash{}, fmt.Errorf("wrong number provided: %s", tokens)
    }

    tx, err := p.Presale.Add(&bind.TransactOpts{
        From:     p.Wallet.Account.From,
        Signer:   p.Wallet.Account.Signer,
        GasLimit: big.NewInt(120000),
    }, common.HexToAddress(addr), tokensAmount)
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (p *Presale) TopUp(addr string, tokens string) (common.Hash, error) {
    tokensAmount, ok := big.NewInt(0).SetString(tokens, 0)
    if !ok {
        return common.Hash{}, fmt.Errorf("wrong number provided: %s", tokens)
    }

    tx, err := p.Presale.TopUp(&bind.TransactOpts{
        From:     p.Wallet.Account.From,
        Signer:   p.Wallet.Account.Signer,
        GasLimit: big.NewInt(120000),
    }, common.HexToAddress(addr), tokensAmount)
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (p *Presale) Approve(addr string) (common.Hash, error) {
    tx, err := p.Presale.Approve(p.Wallet.Account, common.HexToAddress(addr))
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}
