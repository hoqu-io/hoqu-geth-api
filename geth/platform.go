package geth

import (
    "hoqu-geth-api/contract/platform"
    "github.com/ethereum/go-ethereum/common"
    "github.com/spf13/viper"
    "errors"
    "fmt"
    "hoqu-geth-api/sdk/geth"
    "hoqu-geth-api/geth/models"
    sdkModels "hoqu-geth-api/sdk/models"
    "github.com/ethereum/go-ethereum/core/types"
    "math/big"
    "github.com/satori/go.uuid"
)

var hplatform *HoquPlatform

type HoquPlatform struct {
    *geth.Contract
    HoquPlatform *platform.HoQuPlatform
}

func InitHoquPlatform() error {
    c := geth.NewContract(viper.GetString("geth.addr.platform"))
    c.InitEvents(platform.HoQuPlatformABI)

    hp, err := platform.NewHoQuPlatform(c.Address, c.Wallet.Connection)
    if err != nil {
        return errors.New(fmt.Sprintf("Failed to instantiate a HoquPlatform contract: %v", err))
    }

    hplatform = &HoquPlatform{
        Contract:     c,
        HoquPlatform: hp,
    }

    return nil
}

func GetHoquPlatform() *HoquPlatform {
    return hplatform
}

func (hp *HoquPlatform) Deploy() (*common.Address, *types.Transaction, error) {
    address, tx, _, err := platform.DeployHoQuPlatform(
        hp.Wallet.Account,
        hp.Wallet.Connection,
        GetHoQuConfig().Address,
        GetHoQuStorage().Address,
    )
    if err != nil {
        return nil, nil, fmt.Errorf("failed to deploy contract: %v", err)
    }
    return &address, tx, nil
}

func (hp *HoquPlatform) RegisterUser(params *models.RegisterUserRequest) (common.Hash, uuid.UUID, error) {
    id, err := uuid.NewV2('U')
    if err != nil {
        return common.Hash{}, id, err
    }

    tx, err := hp.HoquPlatform.RegisterUser(
        hp.Wallet.Account,
        id,
        params.Role,
        common.HexToAddress(params.Address),
        params.PubKey,
    )
    if err != nil {
        return common.Hash{}, id, err
    }

    return tx.Hash(), id, nil
}

func (hp *HoquPlatform) AddUserAddress(params *models.AddUserAddressRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.AddUserAddress(
        hp.Wallet.Account,
        id,
        common.HexToAddress(params.Address),
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) SetUserStatus(params *models.SetStatusRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.SetUserStatus(
        hp.Wallet.Account,
        id,
        uint8(params.Status),
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) AddIdentification(params *models.AddIdentificationRequest) (common.Hash, uuid.UUID, error) {
    id, err := uuid.NewV2('I')
    if err != nil {
        return common.Hash{}, id, err
    }

    uid, err := uuid.FromString(params.UserId)
    if err != nil {
        return common.Hash{}, id, err
    }

    var cid uuid.UUID
    if params.CompanyId != "" {
        cid, err = uuid.FromString(params.CompanyId)
        if err != nil {
            return common.Hash{}, id, err
        }
    }

    tx, err := hp.HoquPlatform.AddIdentification(
        hp.Wallet.Account,
        id,
        uid,
        params.IdType,
        params.Name,
        cid,
    )
    if err != nil {
        return common.Hash{}, id, err
    }

    return tx.Hash(), id, nil
}

func (hp *HoquPlatform) AddKycReport(params *models.AddKycReportRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.AddKycReport(
        hp.Wallet.Account,
        id,
        params.Meta,
        params.KycLevel,
        params.DataUrl,
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) RegisterCompany(params *models.RegisterCompanyRequest) (common.Hash, uuid.UUID, error) {
    id, err := uuid.NewV2('C')
    if err != nil {
        return common.Hash{}, id, err
    }

    oid, err := uuid.FromString(params.OwnerId)
    if err != nil {
        return common.Hash{}, id, err
    }

    tx, err := hp.HoquPlatform.RegisterCompany(
        hp.Wallet.Account,
        id,
        oid,
        params.Name,
        params.DataUrl,
    )
    if err != nil {
        return common.Hash{}, id, err
    }

    return tx.Hash(), id, nil
}

func (hp *HoquPlatform) SetCompanyStatus(params *models.SetStatusRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.SetCompanyStatus(
        hp.Wallet.Account,
        id,
        uint8(params.Status),
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) RegisterNetwork(params *models.RegisterNetworkRequest) (common.Hash, uuid.UUID, error) {
    id, err := uuid.NewV2('N')
    if err != nil {
        return common.Hash{}, id, err
    }

    oid, err := uuid.FromString(params.OwnerId)
    if err != nil {
        return common.Hash{}, id, err
    }

    tx, err := hp.HoquPlatform.RegisterNetwork(
        hp.Wallet.Account,
        id,
        oid,
        params.Name,
        params.DataUrl,
    )
    if err != nil {
        return common.Hash{}, id, err
    }

    return tx.Hash(), id, nil
}

func (hp *HoquPlatform) SetNetworkStatus(params *models.SetStatusRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.SetNetworkStatus(
        hp.Wallet.Account,
        id,
        uint8(params.Status),
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) RegisterTracker(params *models.RegisterTrackerRequest) (common.Hash, uuid.UUID, error) {
    id, err := uuid.NewV2('T')
    if err != nil {
        return common.Hash{}, id, err
    }

    oid, err := uuid.FromString(params.OwnerId)
    if err != nil {
        return common.Hash{}, id, err
    }

    nid, err := uuid.FromString(params.NetworkId)
    if err != nil {
        return common.Hash{}, id, err
    }

    tx, err := hp.HoquPlatform.RegisterTracker(
        hp.Wallet.Account,
        id,
        oid,
        nid,
        params.Name,
        params.DataUrl,
    )
    if err != nil {
        return common.Hash{}, id, err
    }

    return tx.Hash(), id, nil
}

func (hp *HoquPlatform) SetTrackerStatus(params *models.SetStatusRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.SetTrackerStatus(
        hp.Wallet.Account,
        id,
        uint8(params.Status),
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) AddOffer(params *models.AddOfferRequest) (common.Hash, uuid.UUID, error) {
    id, err := uuid.NewV2('O')
    if err != nil {
        return common.Hash{}, id, err
    }

    ethAmount, ok := big.NewInt(0).SetString(params.Cost, 0)
    if !ok {
        return common.Hash{}, id, fmt.Errorf("wrong offer cost provided: %s", params.Cost)
    }

    oid, err := uuid.FromString(params.OwnerId)
    if err != nil {
        return common.Hash{}, id, err
    }

    nid, err := uuid.FromString(params.NetworkId)
    if err != nil {
        return common.Hash{}, id, err
    }

    mid, err := uuid.FromString(params.MerchantId)
    if err != nil {
        return common.Hash{}, id, err
    }

    tx, err := hp.HoquPlatform.AddOffer(
        hp.Wallet.Account,
        id,
        oid,
        nid,
        mid,
        common.HexToAddress(params.PayerAddress),
        params.Name,
        params.DataUrl,
        ethAmount,
    )
    if err != nil {
        return common.Hash{}, id, err
    }

    return tx.Hash(), id, nil
}

func (hp *HoquPlatform) SetOfferStatus(params *models.SetStatusRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.SetOfferStatus(
        hp.Wallet.Account,
        id,
        uint8(params.Status),
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) AddAd(params *models.AddAdToStorageRequest) (common.Hash, error) {
    aid, err := uuid.FromString(params.AdId)
    if err != nil {
        return common.Hash{}, err
    }

    oid, err := uuid.FromString(params.OwnerId)
    if err != nil {
        return common.Hash{}, err
    }

    offid, err := uuid.FromString(params.OfferId)
    if err != nil {
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.AddAdCampaign(
        hp.Wallet.Account,
        aid,
        oid,
        offid,
        common.HexToAddress(params.ContractAddress),
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) SetAdStatus(params *models.SetStatusRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.SetAdCampaignStatus(
        hp.Wallet.Account,
        id,
        uint8(params.Status),
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) AddLead(params *models.AddLeadRequest) (common.Hash, uuid.UUID, error) {
    id, err := uuid.NewV2('L')
    if err != nil {
        return common.Hash{}, id, err
    }

    ethAmount, ok := big.NewInt(0).SetString(params.Price, 0)
    if !ok {
        return common.Hash{}, id, fmt.Errorf("wrong lead price provided: %s", params.Price)
    }

    aid, err := uuid.FromString(params.AdId)
    if err != nil {
        return common.Hash{}, id, err
    }

    tid, err := uuid.FromString(params.TrackerId)
    if err != nil {
        return common.Hash{}, id, err
    }

    tx, err := hp.HoquPlatform.AddLead(
        hp.Wallet.Account,
        id,
        aid,
        tid,
        params.Meta,
        params.DataUrl,
        ethAmount,
    )
    if err != nil {
        return common.Hash{}, id, err
    }

    return tx.Hash(), id, nil
}

func (hp *HoquPlatform) AddLeadIntermediary(params *models.AddLeadIntermediaryRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    aid, err := uuid.FromString(params.AdId)
    if err != nil {
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.AddLeadIntermediary(
        hp.Wallet.Account,
        id,
        aid,
        common.HexToAddress(params.Address),
        uint32(params.Percent*1e6),
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) SetLeadStatus(params *models.SetLeadStatusRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    aid, err := uuid.FromString(params.AdId)
    if err != nil {
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.SetLeadStatus(
        hp.Wallet.Account,
        id,
        aid,
        uint8(params.Status),
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) TransactLead(id string, adId string) (common.Hash, error) {
    lid, err := uuid.FromString(id)
    if err != nil {
        return common.Hash{}, err
    }

    aid, err := uuid.FromString(adId)
    if err != nil {
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.TransactLead(
        hp.Wallet.Account,
        lid,
        aid,
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) Events(request *sdkModels.Events) ([]sdkModels.ContractEvent, error) {
    events, err := hp.GetEventsByTopics(
        request,
        viper.GetInt64("geth.start_block.platform"),
    )
    if err != nil {
        return nil, err
    }

    for key, event := range events {
        switch {
        case event.Name == "UserRegistered":

            uuId, err := uuid.FromBytes(event.RawArgs[1][:16])
            if err != nil {
                return nil, err
            }

            events[key].Args = models.UserRegisteredEventArgs{
                OwnerAddress: common.BytesToAddress(event.RawArgs[0]).String(),
                Id:           uuId.String(),
                Role:         string(event.RawArgs[2]),
            }
        case event.Name == "UserAddressAdded":

            events[key].Args = models.UserAddressAddedEventArgs{
                OwnerAddress:      common.BytesToAddress(event.RawArgs[0]).String(),
                AdditionalAddress: common.BytesToHash(event.RawArgs[1]).String(),
            }
        case event.Name == "UserPubKeyUpdated" || event.Name == "LeadSold":

            events[key].Args = models.OnlyAddressEventArgs{
                OwnerAddress: common.BytesToAddress(event.RawArgs[0]).String(),
            }
        case event.Name == "KycReportAdded":

            events[key].Args = models.KycReportAddedEventArgs{
                OwnerAddress: common.BytesToAddress(event.RawArgs[0]).String(),
                KycLevel:     uint8(common.BytesToHash(event.RawArgs[1]).Big().Uint64()),
            }
        case event.Name == "UserStatusChanged" || event.Name == "CompanyStatusChanged" ||
            event.Name == "TrackerStatusChanged" || event.Name == "OfferStatusChanged" ||
            event.Name == "LeadStatusChanged" || event.Name == "AdCampaignStatusChanged" ||
            event.Name == "NetworkStatusChanged":

            uuId, err := uuid.FromBytes(event.RawArgs[1][:16])
            if err != nil {
                return nil, err
            }

            events[key].Args = models.StatusChangedEventArgs{
                OwnerAddress: common.BytesToAddress(event.RawArgs[0]).String(),
                Id: uuId.String(),
                Status:       uint8(common.BytesToHash(event.RawArgs[2]).Big().Uint64()),
            }
        case event.Name == "CompanyRegistered" || event.Name == "TrackerRegistered" ||
            event.Name == "OfferAdded" || event.Name == "IdentificationAdded" ||
            event.Name == "NetworkRegistered":

            uuId, err := uuid.FromBytes(event.RawArgs[1][:16])
            if err != nil {
                return nil, err
            }

            events[key].Args = models.IdWithNameEventArgs{
                OwnerAddress: common.BytesToAddress(event.RawArgs[0]).String(),
                Id:           uuId.String(),
                Name:         string(event.RawArgs[2]),
            }
        case event.Name == "AdCampaignAdded":

            uuId, err := uuid.FromBytes(event.RawArgs[1][:16])
            if err != nil {
                return nil, err
            }

            events[key].Args = models.AdAddedEventArgs{
                OwnerAddress: common.BytesToAddress(event.RawArgs[0]).String(),
                Id:           uuId.String(),
                ContractAddress: common.BytesToAddress(event.RawArgs[2]).String(),
            }
        case event.Name == "LeadAdded":

            uuId, err := uuid.FromBytes(event.RawArgs[1][:16])
            if err != nil {
                return nil, err
            }

            events[key].Args = models.LeadAddedEventArgs{
                OwnerAddress: common.BytesToAddress(event.RawArgs[0]).String(),
                Id:           uuId.String(),
                Price:        string(event.RawArgs[2]),
            }
        default:
            return nil, fmt.Errorf("unknown event type: %s", event.Name)
        }
    }

    return events, nil
}
