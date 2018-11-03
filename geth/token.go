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
    sdkModels "hoqu-geth-api/sdk/models"
    "hoqu-geth-api/models"
)

var token *Token

type Token struct {
    *geth.Contract
    Token *contract.HoQuToken
}

func InitToken() error {
    c := geth.NewContract(viper.GetString("geth.addr.token"))
    c.InitEvents(contract.HoQuTokenABI)

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

func (t *Token) Allowance(owner string, spender string) (*big.Int, error) {
    return t.Token.Allowance(nil, common.HexToAddress(owner), common.HexToAddress(spender))
}

func (t *Token) Events(request *sdkModels.Events) ([]sdkModels.ContractEvent, error) {
    events, err := t.GetEventsByTopics(
        request,
        viper.GetInt64("geth.start_block.token"),
    )
    if err != nil {
        return nil, err
    }

    resEvents := make([]sdkModels.ContractEvent, 0)

    for _, event := range events {
        addr := common.BytesToAddress(event.RawArgs[0])

        switch {
        case event.Name == "Transfer":
            event.Args = models.TokenTransferEventArgs{
                From:   addr.String(),
                To:     common.BytesToAddress(event.RawArgs[1]).String(),
                Amount: common.BytesToHash(event.RawArgs[2]).Big().String(),
            }
        case event.Name == "Approval":
            event.Args = models.TokenApprovalEventArgs{
                Owner:   addr.String(),
                Spender: common.BytesToAddress(event.RawArgs[1]).String(),
                Amount:  common.BytesToHash(event.RawArgs[2]).Big().String(),
            }
        default:
            return nil, fmt.Errorf("unknown event type: %s", event.Name)
        }

        resEvents = append(resEvents, event)
    }

    return resEvents, nil
}

func (t *Token) Holders(request *sdkModels.Events) (map[string]string, error) {
    request.EventNames = []string{"Transfer"}
    events, err := t.GetEventsByTopics(
        request,
        viper.GetInt64("geth.start_block.token"),
    )
    if err != nil {
        return nil, err
    }

    holdersBalances := make(map[string]string)

    for _, event := range events {
        from := common.BytesToAddress(event.RawArgs[0])
        to := common.BytesToAddress(event.RawArgs[1])

        if _, ok := holdersBalances[from.String()]; !ok {
            balf, err := t.Balance(from.String())
            if err != nil {
                return nil, err
            }

            holdersBalances[from.String()] = balf.String()
        }

        if _, ok := holdersBalances[to.String()]; !ok {
            bal, err := t.Balance(to.String())
            if err != nil {
                return nil, err
            }

            holdersBalances[to.String()] = bal.String()
        }
    }

    return holdersBalances, nil
}
