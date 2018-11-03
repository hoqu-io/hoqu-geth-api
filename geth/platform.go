package geth

import (
    "hoqu-geth-api/contract/platform"
    "github.com/ethereum/go-ethereum/common"
    "github.com/spf13/viper"
    "errors"
    "fmt"
    "hoqu-geth-api/sdk/geth"
    "hoqu-geth-api/models"
    sdkModels "hoqu-geth-api/sdk/models"
    "github.com/ethereum/go-ethereum/core/types"
    "math/big"
    "github.com/satori/go.uuid"
    "strings"
    "sync"
)

var (
    hplatform *HoquPlatform
    hpOnce    sync.Once
)

type HoquPlatform struct {
    *geth.Contract
    HoquPlatform *platform.HoQuPlatform
}

func initHoquPlatform() (err error) {
    hpOnce.Do(func() {
        c := geth.NewContract(viper.GetString("geth.addr.platform"))
        c.InitEvents(platform.HoQuPlatformABI)

        hp, err := platform.NewHoQuPlatform(c.Address, c.Wallet.Connection)
        if err != nil {
            err = errors.New(fmt.Sprintf("Failed to instantiate a HoquPlatform contract: %v", err))
        }

        hplatform = &HoquPlatform{
            Contract:     c,
            HoquPlatform: hp,
        }
    })
    return
}

func GetHoquPlatform() *HoquPlatform {
    return hplatform
}

func (hp *HoquPlatform) Deploy() (*common.Address, *types.Transaction, error) {
    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return nil, nil, err
    }

    address, tx, _, err := platform.DeployHoQuPlatform(
        opts,
        hp.Wallet.Connection,
        GetHoQuConfig().Address,
        GetHoQuStorage().Address,
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return nil, nil, fmt.Errorf("failed to deploy contract: %v", err)
    }

    hp.Wallet.OnSuccessTransaction()
    return &address, tx, nil
}

func (hp *HoquPlatform) RegisterUser(params *models.RegisterUserRequest) (common.Hash, error) {
    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.RegisterUser(
        opts,
        params.Id,
        params.Role,
        common.HexToAddress(params.Address),
        params.PubKey,
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    hp.Wallet.OnSuccessTransaction()
    return tx.Hash(), nil
}

func (hp *HoquPlatform) AddUserAddress(params *models.AddUserAddressRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.AddUserAddress(
        opts,
        id,
        common.HexToAddress(params.Address),
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    hp.Wallet.OnSuccessTransaction()
    return tx.Hash(), nil
}

func (hp *HoquPlatform) SetUserStatus(params *models.SetStatusRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.SetUserStatus(
        opts,
        id,
        uint8(params.Status),
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    hp.Wallet.OnSuccessTransaction()
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

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, id, err
    }

    tx, err := hp.HoquPlatform.AddIdentification(
        opts,
        id,
        uid,
        params.IdType,
        params.Name,
        cid,
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, id, err
    }

    hp.Wallet.OnSuccessTransaction()
    return tx.Hash(), id, nil
}

func (hp *HoquPlatform) AddKycReport(params *models.AddKycReportRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    params.DataUrl = strings.Replace(params.DataUrl, ":id", id.String(), 1)

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.AddKycReport(
        opts,
        id,
        params.Meta,
        params.KycLevel,
        params.DataUrl,
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    hp.Wallet.OnSuccessTransaction()
    return tx.Hash(), nil
}

func (hp *HoquPlatform) RegisterCompany(params *models.RegisterCompanyRequest) (common.Hash, uuid.UUID, error) {
    id, err := uuid.NewV2('C')
    if err != nil {
        return common.Hash{}, id, err
    }

    params.DataUrl = strings.Replace(params.DataUrl, ":id", id.String(), 1)

    oid, err := uuid.FromString(params.OwnerId)
    if err != nil {
        return common.Hash{}, id, err
    }

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, id, err
    }

    tx, err := hp.HoquPlatform.RegisterCompany(
        opts,
        id,
        oid,
        params.Name,
        params.DataUrl,
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, id, err
    }

    hp.Wallet.OnSuccessTransaction()
    return tx.Hash(), id, nil
}

func (hp *HoquPlatform) SetCompanyStatus(params *models.SetStatusRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.SetCompanyStatus(
        opts,
        id,
        uint8(params.Status),
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    hp.Wallet.OnSuccessTransaction()
    return tx.Hash(), nil
}

func (hp *HoquPlatform) RegisterNetwork(params *models.RegisterNetworkRequest) (common.Hash, uuid.UUID, error) {
    id, err := uuid.NewV2('N')
    if err != nil {
        return common.Hash{}, id, err
    }

    params.DataUrl = strings.Replace(params.DataUrl, ":id", id.String(), 1)

    oid, err := uuid.FromString(params.OwnerId)
    if err != nil {
        return common.Hash{}, id, err
    }

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, id, err
    }

    tx, err := hp.HoquPlatform.RegisterNetwork(
        opts,
        id,
        oid,
        params.Name,
        params.DataUrl,
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, id, err
    }

    hp.Wallet.OnSuccessTransaction()
    return tx.Hash(), id, nil
}

func (hp *HoquPlatform) SetNetworkStatus(params *models.SetStatusRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.SetNetworkStatus(
        opts,
        id,
        uint8(params.Status),
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    hp.Wallet.OnSuccessTransaction()
    return tx.Hash(), nil
}

func (hp *HoquPlatform) RegisterTracker(params *models.RegisterTrackerRequest) (common.Hash, uuid.UUID, error) {
    id, err := uuid.NewV2('T')
    if err != nil {
        return common.Hash{}, id, err
    }

    params.DataUrl = strings.Replace(params.DataUrl, ":id", id.String(), 1)

    oid, err := uuid.FromString(params.OwnerId)
    if err != nil {
        return common.Hash{}, id, err
    }

    nid, err := uuid.FromString(params.NetworkId)
    if err != nil {
        return common.Hash{}, id, err
    }

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, id, err
    }

    tx, err := hp.HoquPlatform.RegisterTracker(
        opts,
        id,
        oid,
        nid,
        params.Name,
        params.DataUrl,
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, id, err
    }

    hp.Wallet.OnSuccessTransaction()
    return tx.Hash(), id, nil
}

func (hp *HoquPlatform) SetTrackerStatus(params *models.SetStatusRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.SetTrackerStatus(
        opts,
        id,
        uint8(params.Status),
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    hp.Wallet.OnSuccessTransaction()
    return tx.Hash(), nil
}

func (hp *HoquPlatform) AddOffer(params *models.AddOfferRequest) (common.Hash, error) {
    oid, err := uuid.FromString(params.OwnerId)
    if err != nil {
        return common.Hash{}, err
    }

    nid, err := uuid.FromString(params.NetworkId)
    if err != nil {
        return common.Hash{}, err
    }

    mid, err := uuid.FromString(params.MerchantId)
    if err != nil {
        return common.Hash{}, err
    }

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.AddOffer(
        opts,
        params.Id,
        oid,
        nid,
        mid,
        common.HexToAddress(params.PayerAddress),
        params.Name,
        params.DataUrl,
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    hp.Wallet.OnSuccessTransaction()
    return tx.Hash(), nil
}

func (hp *HoquPlatform) SetOfferStatus(params *models.SetStatusRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.SetOfferStatus(
        opts,
        id,
        uint8(params.Status),
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    hp.Wallet.OnSuccessTransaction()
    return tx.Hash(), nil
}

func (hp *HoquPlatform) SetOfferName(params *models.SetNameRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.SetOfferName(
        opts,
        id,
        params.Name,
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    hp.Wallet.OnSuccessTransaction()
    return tx.Hash(), nil
}

func (hp *HoquPlatform) SetOfferPayerAddress(params *models.SetOfferPayerAddressRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.SetOfferPayerAddress(
        opts,
        id,
        common.HexToAddress(params.PayerAddress),
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    hp.Wallet.OnSuccessTransaction()
    return tx.Hash(), nil
}

func (hp *HoquPlatform) AddOfferTariffGroup(params *models.IdWithChildIdRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    chid, err := uuid.FromString(params.ChildId)
    if err != nil {
        return common.Hash{}, err
    }

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.AddOfferTariffGroup(
        opts,
        id,
        chid,
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    hp.Wallet.OnSuccessTransaction()
    return tx.Hash(), nil
}

func (hp *HoquPlatform) AddAdCampaign(params *models.AddAdToStorageRequest) (common.Hash, error) {
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

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.AddAdCampaign(
        opts,
        aid,
        oid,
        offid,
        common.HexToAddress(params.ContractAddress),
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    hp.Wallet.OnSuccessTransaction()
    return tx.Hash(), nil
}

func (hp *HoquPlatform) SetAdCampaignStatus(params *models.SetStatusRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.SetAdCampaignStatus(
        opts,
        id,
        uint8(params.Status),
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    hp.Wallet.OnSuccessTransaction()
    return tx.Hash(), nil
}

func (hp *HoquPlatform) AddLead(params *models.AddLeadRequest) (common.Hash, error) {
    ethAmount, ok := big.NewInt(0).SetString(params.Price, 0)
    if !ok {
        return common.Hash{}, fmt.Errorf("wrong lead price provided: %s", params.Price)
    }

    aid, err := uuid.FromString(params.AdId)
    if err != nil {
        return common.Hash{}, err
    }

    tid, err := uuid.FromString(params.TrackerId)
    if err != nil {
        return common.Hash{}, err
    }

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.AddLead(
        opts,
        params.Id,
        aid,
        tid,
        params.Meta,
        params.DataUrl,
        ethAmount,
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    hp.Wallet.OnSuccessTransaction()
    return tx.Hash(), nil
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

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.AddLeadIntermediary(
        opts,
        id,
        aid,
        common.HexToAddress(params.Address),
        uint32(params.Percent*1e6),
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    hp.Wallet.OnSuccessTransaction()
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

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.SetLeadStatus(
        opts,
        id,
        aid,
        uint8(params.Status),
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    hp.Wallet.OnSuccessTransaction()
    return tx.Hash(), nil
}

func (hp *HoquPlatform) SetLeadPrice(params *models.SetLeadPriceRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    aid, err := uuid.FromString(params.AdId)
    if err != nil {
        return common.Hash{}, err
    }

    ethAmount, ok := big.NewInt(0).SetString(params.Price, 0)
    if !ok {
        return common.Hash{}, fmt.Errorf("wrong lead price provided: %s", params.Price)
    }

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.SetLeadPrice(
        opts,
        id,
        aid,
        ethAmount,
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    hp.Wallet.OnSuccessTransaction()
    return tx.Hash(), nil
}

func (hp *HoquPlatform) SetLeadDataUrl(params *models.SetLeadDataUrlRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    aid, err := uuid.FromString(params.AdId)
    if err != nil {
        return common.Hash{}, err
    }

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.SetLeadDataUrl(
        opts,
        id,
        aid,
        params.DataUrl,
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    hp.Wallet.OnSuccessTransaction()
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

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.TransactLead(
        opts,
        lid,
        aid,
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    hp.Wallet.OnSuccessTransaction()
    return tx.Hash(), nil
}

func (hp *HoquPlatform) AddTariffGroup(params *models.AddTariffGroupRequest) (common.Hash, error) {
    oid, err := uuid.FromString(params.OwnerId)
    if err != nil {
        return common.Hash{}, err
    }

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.AddTariffGroup(
        opts,
        params.Id,
        oid,
        params.Name,
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    hp.Wallet.OnSuccessTransaction()
    return tx.Hash(), nil
}

func (hp *HoquPlatform) AddTariff(params *models.AddTariffRequest) (common.Hash, error) {
    ethAmount, ok := big.NewInt(0).SetString(params.Price, 0)
    if !ok {
        return common.Hash{}, fmt.Errorf("wrong tariff price provided: %s", params.Price)
    }

    tgid, err := uuid.FromString(params.TariffGroupId)
    if err != nil {
        return common.Hash{}, err
    }

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.AddTariff(
        opts,
        params.Id,
        tgid,
        params.Name,
        params.Action,
        params.CalcMethod,
        ethAmount,
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    hp.Wallet.OnSuccessTransaction()
    return tx.Hash(), nil
}

func (hp *HoquPlatform) SetTariffGroupStatus(params *models.SetStatusRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.SetTariffGroupStatus(
        opts,
        id,
        uint8(params.Status),
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    hp.Wallet.OnSuccessTransaction()
    return tx.Hash(), nil
}

func (hp *HoquPlatform) SetTariffStatus(params *models.SetStatusRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    opts, err := hp.Wallet.GetTransactOpts()
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.SetTariffStatus(
        opts,
        id,
        uint8(params.Status),
    )
    if err != nil {
        hp.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    hp.Wallet.OnSuccessTransaction()
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

            r := make([]byte, 0)
            for _, arg := range event.RawArgs[4:] {
                r = append(r, arg...)
            }

            events[key].Args = models.UserRegisteredEventArgs{
                OwnerAddress: common.BytesToAddress(event.RawArgs[0]).String(),
                Id:           uuId.String(),
                Role:         hp.FilterString(string(r)),
            }

        case event.Name == "UserAddressAdded":

            uuId, err := uuid.FromBytes(event.RawArgs[2][:16])
            if err != nil {
                return nil, err
            }

            events[key].Args = models.UserAddressAddedEventArgs{
                OwnerAddress:      common.BytesToAddress(event.RawArgs[0]).String(),
                AdditionalAddress: common.BytesToHash(event.RawArgs[1]).String(),
                Id:                uuId.String(),
            }

        case event.Name == "UserPubKeyUpdated":

            uuId, err := uuid.FromBytes(event.RawArgs[1][:16])
            if err != nil {
                return nil, err
            }

            events[key].Args = models.OnlyAddressEventArgs{
                OwnerAddress: common.BytesToAddress(event.RawArgs[0]).String(),
                Id:           uuId.String(),
            }

        case event.Name == "KycReportAdded":

            uuId, err := uuid.FromBytes(event.RawArgs[2][:16])
            if err != nil {
                return nil, err
            }

            usId, err := uuid.FromBytes(event.RawArgs[3][:16])
            if err != nil {
                return nil, err
            }

            events[key].Args = models.KycReportAddedEventArgs{
                OwnerAddress: common.BytesToAddress(event.RawArgs[0]).String(),
                KycLevel:     uint8(common.BytesToHash(event.RawArgs[1]).Big().Uint64()),
                Id:           uuId.String(),
                UserId:       usId.String(),
            }

        case event.Name == "UserStatusChanged" || event.Name == "CompanyStatusChanged" ||
            event.Name == "TrackerStatusChanged" || event.Name == "OfferStatusChanged" ||
            event.Name == "LeadStatusChanged" || event.Name == "AdCampaignStatusChanged" ||
            event.Name == "NetworkStatusChanged" || event.Name == "TariffGroupStatusChanged":

            uuId, err := uuid.FromBytes(event.RawArgs[1][:16])
            if err != nil {
                return nil, err
            }

            events[key].Args = models.StatusChangedEventArgs{
                OwnerAddress: common.BytesToAddress(event.RawArgs[0]).String(),
                Id:           uuId.String(),
                Status:       uint8(common.BytesToHash(event.RawArgs[2]).Big().Uint64()),
            }

        case event.Name == "IdentificationAdded":

            uuId, err := uuid.FromBytes(event.RawArgs[1][:16])
            if err != nil {
                return nil, err
            }

            usId, err := uuid.FromBytes(event.RawArgs[2][:16])
            if err != nil {
                return nil, err
            }

            r := make([]byte, 0)
            for _, arg := range event.RawArgs[3:] {
                r = append(r, arg...)
            }

            events[key].Args = models.IdentificationAddedEventArgs{
                OwnerAddress: common.BytesToAddress(event.RawArgs[0]).String(),
                Id:           uuId.String(),
                UserId:       usId.String(),
                Name:         hp.FilterString(string(r)),
            }

        case event.Name == "CompanyRegistered" || event.Name == "TrackerRegistered" ||
            event.Name == "OfferAdded" || event.Name == "NetworkRegistered" ||
            event.Name == "TariffGroupAdded" || event.Name == "TariffAdded":

            uuId, err := uuid.FromBytes(event.RawArgs[1][:16])
            if err != nil {
                return nil, err
            }

            r := make([]byte, 0)
            for _, arg := range event.RawArgs[4:] {
                r = append(r, arg...)
            }

            events[key].Args = models.IdWithNameEventArgs{
                OwnerAddress: common.BytesToAddress(event.RawArgs[0]).String(),
                Id:           uuId.String(),
                Name:         hp.FilterString(string(r)),
            }

        case event.Name == "AdCampaignAdded":

            uuId, err := uuid.FromBytes(event.RawArgs[1][:16])
            if err != nil {
                return nil, err
            }

            events[key].Args = models.AdAddedEventArgs{
                OwnerAddress:    common.BytesToAddress(event.RawArgs[0]).String(),
                Id:              uuId.String(),
                ContractAddress: common.BytesToAddress(event.RawArgs[2]).String(),
            }

        case event.Name == "LeadAdded" || event.Name == "LeadTransacted" || event.Name == "LeadStatusChanged" ||
            event.Name == "LeadPriceChanged" || event.Name == "LeadDataUrlChanged":

            contractAddr := common.BytesToAddress(event.RawArgs[0]).String()

            adId, err := uuid.FromBytes(event.RawArgs[1][:16])
            if err != nil {
                return nil, err
            }

            uuId, err := uuid.FromBytes(event.RawArgs[2][:16])
            if err != nil {
                return nil, err
            }

            switch {
            case event.Name == "LeadStatusChanged":
                events[key].Args = models.LeadStatusChangedEventArgs{
                    ContractAddress: contractAddr,
                    AdId:            adId.String(),
                    Id:              uuId.String(),
                    Status:          uint8(common.BytesToHash(event.RawArgs[3]).Big().Uint64()),
                }
            case event.Name == "LeadPriceChanged":
                events[key].Args = models.LeadPriceChangedEventArgs{
                    ContractAddress: contractAddr,
                    AdId:            adId.String(),
                    Id:              uuId.String(),
                    Price:           common.BytesToHash(event.RawArgs[3]).Big().String(),
                }
            case event.Name == "LeadDataUrlChanged":
                r := make([]byte, 0)
                for _, arg := range event.RawArgs[3:] {
                    r = append(r, arg...)
                }

                events[key].Args = models.LeadDataUrlChangedEventArgs{
                    ContractAddress: contractAddr,
                    AdId:            adId.String(),
                    Id:              uuId.String(),
                    DataUrl:         hp.FilterString(string(r)),
                }
            default:
                events[key].Args = models.LeadAddedEventArgs{
                    ContractAddress: contractAddr,
                    AdId:            adId.String(),
                    Id:              uuId.String(),
                }
            }

        case event.Name == "OfferTariffGroupAdded":

            pId, err := uuid.FromBytes(event.RawArgs[1][:16])
            if err != nil {
                return nil, err
            }

            chId, err := uuid.FromBytes(event.RawArgs[2][:16])
            if err != nil {
                return nil, err
            }

            events[key].Args = models.IdWithChildIdEventArgs{
                OwnerAddress: common.BytesToAddress(event.RawArgs[0]).String(),
                Id:           pId.String(),
                ChildId:      chId.String(),
            }

        default:
            return nil, fmt.Errorf("unknown event type: %s", event.Name)
        }
    }

    return events, nil
}
