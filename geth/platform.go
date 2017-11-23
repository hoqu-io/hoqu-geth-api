package geth

import (
    "hoqu-api/contract/platform"
    "github.com/ethereum/go-ethereum/common"
    "github.com/spf13/viper"
    "errors"
    "fmt"
    "hoqu-api/sdk/geth"
    "hoqu-api/geth/models"
    sdkModels "hoqu-api/sdk/models"
    "github.com/ethereum/go-ethereum/core/types"
)

var hplatform *HoquPlatform

type HoquPlatform struct {
    *geth.Contract
    HoquPlatform *platform.HoQuPlatform
}

func InitHoquPlatform() error {
    c := geth.NewContract(viper.GetString("geth.addr.platform"))
    c.InitEvents(platform.HoQuPlatformABI)

    s, err := platform.NewHoQuPlatform(c.Address, c.Wallet.Connection)
    if err != nil {
        return errors.New(fmt.Sprintf("Failed to instantiate a HoquPlatform contract: %v", err))
    }

    hplatform = &HoquPlatform{
        Contract:     c,
        HoquPlatform: s,
    }

    return nil
}

func GetHoquPlatform() *HoquPlatform {
    return hplatform
}

func (s *HoquPlatform) Deploy() (*common.Address, *types.Transaction, error) {
    var tokenAddr common.Address
    tokenAddr, err := GetPrivatePlacement().TokenAddr()
    if err != nil {
        return nil, nil, err
    }

    address, tx, _, err := platform.DeployHoQuPlatform(
        s.Wallet.Account,
        s.Wallet.Connection,
        tokenAddr,
        GetHoQuConfig().Address,
    )
    if err != nil {
        return nil, nil, fmt.Errorf("failed to deploy contract: %v", err)
    }
    return &address, tx, nil
}

func (s *HoquPlatform) GetUser(id uint32) (*models.UserData, error) {
    user, err := s.HoquPlatform.Users(nil, id)
    if err != nil {
        return nil, err
    }

    addresses := make(map[uint8]string)
    for num := uint8(0); num < user.NumOfAddresses; num++ {
        addr, err := s.HoquPlatform.GetUserAddress(nil, id, num)
        if err != nil {
            return nil, err
        }
        addresses[num] = addr.String()
    }

    userData := &models.UserData{
        CreatedAt: user.CreatedAt.String(),
        Addresses: addresses,
        Role:      user.Role,
        KycLevel:  user.KycLevel,
        PubKey:    user.PubKey,
        Status:    user.Status,
    }

    return userData, nil
}

func (s *HoquPlatform) Events(addrs []string) ([]sdkModels.ContractEvent, error) {
    hashAddrs := make([]common.Hash, len(addrs))
    for _, addr := range addrs {
        hashAddrs = append(hashAddrs, common.HexToHash(addr))
    }

    events, err := s.GetEventsByTopics([][]common.Hash{{}, hashAddrs})
    if err != nil {
        return nil, err
    }

    for key, event := range events {
        switch {
        case event.Name == "UserRegistered":
            events[key].Args = models.UserRegisteredEventArgs{
                OwnerAddress: common.BytesToAddress(event.RawArgs[0]).String(),
                Id:           common.BytesToHash(event.RawArgs[1]).Big().String(),
                Role:         string(event.RawArgs[1]),
            }
        default:
            return nil, fmt.Errorf("unknown event type: %s", event.Name)
        }
    }

    return events, nil
}

func (s *HoquPlatform) AddUserAddress(id uint32, addr string) (common.Hash, error) {
    tx, err := s.HoquPlatform.AddUserAddress(s.Wallet.Account, id, common.HexToAddress(addr))
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}
