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
    tokenAddr := GetToken().Address

    address, tx, _, err := platform.DeployHoQuPlatform(
        hp.Wallet.Account,
        hp.Wallet.Connection,
        GetHoQuConfig().Address,
        tokenAddr,
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

func (hp *HoquPlatform) AddUserKycReport(params *models.AddUserKycReportRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.AddUserKycReport(
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

func (hp *HoquPlatform) GetUser(id string) (*models.UserData, error) {
    uid, err := uuid.FromString(id)
    if err != nil {
        return nil, err
    }

    user, err := hp.HoquPlatform.Users(nil, uid)
    if err != nil {
        return nil, err
    }

    addresses := make(map[uint8]string)
    for num := uint8(0); num < user.NumOfAddresses; num++ {
        addr, err := hp.HoquPlatform.GetUserAddress(nil, uid, num)
        if err != nil {
            return nil, err
        }
        addresses[num] = addr.String()
    }

    reports := make(map[uint16]models.KycReport)
    for num := uint16(0); num < user.NumOfKycReports; num++ {
        createdAt, meta, kycLevel, dataUrl, err := hp.HoquPlatform.GetUserKycReport(nil, uid, num)
        if err != nil {
            return nil, err
        }
        reports[num] = models.KycReport{
            CreatedAt: createdAt.String(),
            Meta:      meta,
            KycLevel:  kycLevel,
            DataUrl:   dataUrl,
        }
    }

    userData := &models.UserData{
        CreatedAt:  user.CreatedAt.String(),
        Addresses:  addresses,
        Role:       user.Role,
        KycLevel:   user.KycLevel,
        KycReports: reports,
        PubKey:     user.PubKey,
        Status:     models.Status(user.Status),
    }

    return userData, nil
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
        common.HexToAddress(params.OwnerAddress),
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

func (hp *HoquPlatform) GetCompany(id string) (*models.CompanyData, error) {
    cid, err := uuid.FromString(id)
    if err != nil {
        return nil, err
    }

    company, err := hp.HoquPlatform.Companies(nil, cid)
    if err != nil {
        return nil, err
    }

    oid, err := uuid.FromBytes(company.OwnerId[:])
    if err != nil {
        return nil, err
    }

    companyData := &models.CompanyData{
        CreatedAt:    company.CreatedAt.String(),
        OwnerId:      oid.String(),
        OwnerAddress: company.OwnerAddress.String(),
        Name:         company.Name,
        DataUrl:      company.DataUrl,
        Status:       models.Status(company.Status),
    }

    return companyData, nil
}

func (hp *HoquPlatform) RegisterTracker(params *models.RegisterTrackerRequest) (common.Hash, uuid.UUID, error) {
    id, err := uuid.NewV2('T')
    if err != nil {
        return common.Hash{}, id, err
    }

    tx, err := hp.HoquPlatform.RegisterTracker(
        hp.Wallet.Account,
        id,
        common.HexToAddress(params.OwnerAddress),
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

func (hp *HoquPlatform) GetTracker(id string) (*models.TrackerData, error) {
    tid, err := uuid.FromString(id)
    if err != nil {
        return nil, err
    }

    tracker, err := hp.HoquPlatform.Trackers(nil, tid)
    if err != nil {
        return nil, err
    }

    trackerData := &models.TrackerData{
        CreatedAt:    tracker.CreatedAt.String(),
        OwnerAddress: tracker.OwnerAddress.String(),
        Name:         tracker.Name,
        DataUrl:      tracker.DataUrl,
        Status:       models.Status(tracker.Status),
    }

    return trackerData, nil
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

    cid, err := uuid.FromString(params.CompanyId)
    if err != nil {
        return common.Hash{}, id, err
    }

    tx, err := hp.HoquPlatform.AddOffer(
        hp.Wallet.Account,
        id,
        cid,
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

func (hp *HoquPlatform) GetOffer(id string) (*models.OfferData, error) {
    oid, err := uuid.FromString(id)
    if err != nil {
        return nil, err
    }

    offer, err := hp.HoquPlatform.Offers(nil, oid)
    if err != nil {
        return nil, err
    }

    cid, err := uuid.FromBytes(offer.CompanyId[:])
    if err != nil {
        return nil, err
    }

    offerData := &models.OfferData{
        CreatedAt:    offer.CreatedAt.String(),
        CompanyId:    cid.String(),
        PayerAddress: offer.PayerAddress.String(),
        Name:         offer.Name,
        DataUrl:      offer.DataUrl,
        Cost:         offer.Cost.String(),
        Status:       models.Status(offer.Status),
    }

    return offerData, nil
}

func (hp *HoquPlatform) AddAd(params *models.AddAdRequest) (common.Hash, uuid.UUID, error) {
    id, err := uuid.NewV2('A')
    if err != nil {
        return common.Hash{}, id, err
    }

    oid, err := uuid.FromString(params.OwnerId)
    if err != nil {
        return common.Hash{}, id, err
    }

    offid, err := uuid.FromString(params.OfferId)
    if err != nil {
        return common.Hash{}, id, err
    }

    tx, err := hp.HoquPlatform.AddAd(
        hp.Wallet.Account,
        id,
        oid,
        offid,
        common.HexToAddress(params.BeneficiaryAddress),
    )
    if err != nil {
        return common.Hash{}, id, err
    }

    return tx.Hash(), id, nil
}

func (hp *HoquPlatform) SetAdStatus(params *models.SetStatusRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.SetAdStatus(
        hp.Wallet.Account,
        id,
        uint8(params.Status),
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) GetAd(id string) (*models.AdData, error) {
    aid, err := uuid.FromString(id)
    if err != nil {
        return nil, err
    }

    ad, err := hp.HoquPlatform.Ads(nil, aid)
    if err != nil {
        return nil, err
    }

    oid, err := uuid.FromBytes(ad.OwnerId[:])
    if err != nil {
        return nil, err
    }

    offid, err := uuid.FromBytes(ad.OfferId[:])
    if err != nil {
        return nil, err
    }

    adData := &models.AdData{
        CreatedAt:          ad.CreatedAt.String(),
        OwnerId:            oid.String(),
        OfferId:            offid.String(),
        BeneficiaryAddress: ad.BeneficiaryAddress.String(),
    }

    return adData, nil
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

    tx, err := hp.HoquPlatform.AddLeadIntermediary(
        hp.Wallet.Account,
        id,
        common.HexToAddress(params.Address),
        uint32(params.Percent*1e6),
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) SetLeadStatus(params *models.SetStatusRequest) (common.Hash, error) {
    id, err := uuid.FromString(params.Id)
    if err != nil {
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.SetLeadStatus(
        hp.Wallet.Account,
        id,
        uint8(params.Status),
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) SellLead(id string) (common.Hash, error) {
    lid, err := uuid.FromString(id)
    if err != nil {
        return common.Hash{}, err
    }

    tx, err := hp.HoquPlatform.SellLead(
        hp.Wallet.Account,
        lid,
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) GetLead(id string) (*models.LeadData, error) {
    lid, err := uuid.FromString(id)
    if err != nil {
        return nil, err
    }

    lead, err := hp.HoquPlatform.Leads(nil, lid)
    if err != nil {
        return nil, err
    }

    intermediaries := make(map[string]float32)
    for num := uint8(0); num < lead.NumOfIntermediaries; num++ {
        addr, err := hp.HoquPlatform.GetLeadIntermediaryAddress(nil, lid, num)
        if err != nil {
            return nil, err
        }
        percent, err := hp.HoquPlatform.GetLeadIntermediaryPercent(nil, lid, num)
        if err != nil {
            return nil, err
        }
        intermediaries[addr.String()] = float32(percent) / 1e6
    }

    aid, err := uuid.FromBytes(lead.AdId[:])
    if err != nil {
        return nil, err
    }

    tid, err := uuid.FromBytes(lead.TrackerId[:])
    if err != nil {
        return nil, err
    }

    leadData := &models.LeadData{
        CreatedAt:      lead.CreatedAt.String(),
        AdId:           aid.String(),
        TrackerId:      tid.String(),
        Intermediaries: intermediaries,
        Meta:           lead.Meta,
        DataUrl:        lead.DataUrl,
        Price:          lead.Price.String(),
        Status:         models.Status(lead.Status),
    }

    return leadData, nil
}

func (hp *HoquPlatform) Events(addrs []string, latest int64) ([]sdkModels.ContractEvent, error) {
    hashAddrs := make([]common.Hash, len(addrs))
    for _, addr := range addrs {
        hashAddrs = append(hashAddrs, common.HexToHash(addr))
    }

    from := big.NewInt(viper.GetInt64("geth.start_block.platform"))
    if latest != 0 {
        b, err := hp.Wallet.GetBlockHeaderByNumber(nil)
        if err != nil {
            return nil, err
        }
        from = big.NewInt(0).Sub(b.Number, big.NewInt(latest))
    }

    events, err := hp.GetEventsByTopics(
        [][]common.Hash{{}, hashAddrs},
        from,
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
        case event.Name == "UserKycReportAdded":

            events[key].Args = models.UserKycReportAddedEventArgs{
                OwnerAddress: common.BytesToAddress(event.RawArgs[0]).String(),
                KycLevel:     uint8(common.BytesToHash(event.RawArgs[1]).Big().Uint64()),
            }
        case event.Name == "UserStatusChanged" || event.Name == "CompanyStatusChanged" ||
            event.Name == "TrackerStatusChanged" || event.Name == "OfferStatusChanged" ||
            event.Name == "LeadStatusChanged" || event.Name == "AdStatusChanged":

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
            event.Name == "OfferAdded":

            uuId, err := uuid.FromBytes(event.RawArgs[1][:16])
            if err != nil {
                return nil, err
            }

            events[key].Args = models.IdWithNameEventArgs{
                OwnerAddress: common.BytesToAddress(event.RawArgs[0]).String(),
                Id:           uuId.String(),
                Name:         string(event.RawArgs[2]),
            }
        case event.Name == "AdAdded":

            uuId, err := uuid.FromBytes(event.RawArgs[1][:16])
            if err != nil {
                return nil, err
            }

            events[key].Args = models.AdAddedEventArgs{
                OwnerAddress: common.BytesToAddress(event.RawArgs[0]).String(),
                Id:           uuId.String(),
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
