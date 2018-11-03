package geth

import (
    platform "hoqu-geth-api/contract/config"
    "github.com/ethereum/go-ethereum/common"
    "github.com/spf13/viper"
    "errors"
    "fmt"
    "hoqu-geth-api/sdk/geth"
    "hoqu-geth-api/models"
    sdkModels "hoqu-geth-api/sdk/models"
    "github.com/ethereum/go-ethereum/core/types"
)

var config *HoQuConfig

type HoQuConfig struct {
    *geth.Contract
    HoQuConfig *platform.HoQuConfig
}

func initHoQuConfig() error {
    c := geth.NewContract(viper.GetString("geth.addr.conf"))
    c.InitEvents(platform.HoQuConfigABI)

    s, err := platform.NewHoQuConfig(c.Address, c.Wallet.Connection)
    if err != nil {
        return errors.New(fmt.Sprintf("Failed to instantiate a HoQu Platform Config contract: %v", err))
    }

    config = &HoQuConfig{
        Contract:   c,
        HoQuConfig: s,
    }

    return nil
}

func GetHoQuConfig() *HoQuConfig {
    return config
}

func (s *HoQuConfig) Deploy(params *models.ConfigDeployParams) (*common.Address, *types.Transaction, error) {
    opts, err := s.Wallet.GetTransactOpts()
    if err != nil {
        s.Wallet.OnFailTransaction(err)
        return nil, nil, err
    }

    address, tx, _, err := platform.DeployHoQuConfig(
        opts,
        s.Wallet.Connection,
        common.HexToAddress(params.CommissionWallet),
    )
    if err != nil {
        s.Wallet.OnFailTransaction(err)
        return nil, nil, fmt.Errorf("failed to deploy HoQuConfig contract: %v", err)
    }

    s.Wallet.OnSuccessTransaction()
    return &address, tx, nil
}

func (s *HoQuConfig) Events(request *sdkModels.Events) ([]sdkModels.ContractEvent, error) {
    events, err := s.GetEventsByTopics(
        request,
        viper.GetInt64("geth.start_block.conf"),
    )
    if err != nil {
        return nil, err
    }

    for key, event := range events {
        switch {
        case event.Name == "SystemOwnerAdded":
            events[key].Args = models.SystemOwnerAddedEventArgs{
                NewOwner: common.BytesToAddress(event.RawArgs[0]).String(),
            }
        case event.Name == "SystemOwnerChanged":
            events[key].Args = models.SystemOwnerChangedEventArgs{
                PreviousOwner: common.BytesToAddress(event.RawArgs[0]).String(),
                NewOwner:      common.BytesToAddress(event.RawArgs[1]).String(),
            }
        case event.Name == "SystemOwnerDeleted":
            events[key].Args = models.SystemOwnerDeletedEventArgs{
                DeletedOwner: common.BytesToAddress(event.RawArgs[0]).String(),
            }
        case event.Name == "CommissionWalletChanged":
            events[key].Args = models.CommissionWalletChangedEventArgs{
                ChangedBy:        common.BytesToAddress(event.RawArgs[0]).String(),
                CommissionWallet: common.BytesToAddress(event.RawArgs[1]).String(),
            }
        case event.Name == "CommissionChanged":
            events[key].Args = models.CommissionChangedEventArgs{
                ChangedBy:  common.BytesToAddress(event.RawArgs[0]).String(),
                Commission: common.BytesToHash(event.RawArgs[1]).Big().String(),
            }
        default:
            return nil, fmt.Errorf("unknown event type: %s", event.Name)
        }
    }

    return events, nil
}

func (s *HoQuConfig) AddOwner(addr string) (common.Hash, error) {
    opts, err := s.Wallet.GetTransactOpts()
    if err != nil {
        s.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := s.HoQuConfig.AddOwner(opts, common.HexToAddress(addr))
    if err != nil {
        s.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    s.Wallet.OnSuccessTransaction()
    return tx.Hash(), nil
}
