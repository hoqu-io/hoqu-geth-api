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
    var tokenAddr common.Address
    tokenAddr, err := GetPrivatePlacement().TokenAddr()
    if err != nil {
        return nil, nil, err
    }

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

func (hp *HoquPlatform) RegisterUser(params *models.RegisterUserRequest) (common.Hash, error) {
    tx, err := hp.HoquPlatform.RegisterUser(
        hp.Wallet.Account,
        params.Id,
        params.Role,
        common.HexToAddress(params.Address),
        params.PubKey,
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) AddUserKycReport(params *models.AddUserKycReportRequest) (common.Hash, error) {
    tx, err := hp.HoquPlatform.AddUserKycReport(
        hp.Wallet.Account,
        params.Id,
        params.ReportId,
        params.KycLevel,
        params.DataUrl,
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) AddUserAddress(params *models.AddUserAddressRequest) (common.Hash, error) {
    tx, err := hp.HoquPlatform.AddUserAddress(
        hp.Wallet.Account,
        params.Id,
        common.HexToAddress(params.Address),
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) SetUserStatus(params *models.SetStatusRequest) (common.Hash, error) {
    tx, err := hp.HoquPlatform.SetUserStatus(
        hp.Wallet.Account,
        params.Id,
        params.Status,
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) GetUser(id uint32) (*models.UserData, error) {
    user, err := hp.HoquPlatform.Users(nil, id)
    if err != nil {
        return nil, err
    }

    addresses := make(map[uint8]string)
    for num := uint8(0); num < user.NumOfAddresses; num++ {
        addr, err := hp.HoquPlatform.GetUserAddress(nil, id, num)
        if err != nil {
            return nil, err
        }
        addresses[num] = addr.String()
    }

    reports := make(map[uint16]models.KycReport)
    for num := uint16(0); num < user.NumOfKycReports; num++ {
        createdAt, reportId, kycLevel, dataUrl, err := hp.HoquPlatform.GetUserKycReport(nil, id, num)
        if err != nil {
            return nil, err
        }
        reports[num] = models.KycReport{
            CreatedAt: createdAt.String(),
            ReportId:  reportId,
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
        Status:     user.Status,
    }

    return userData, nil
}

func (hp *HoquPlatform) RegisterCompany(params *models.RegisterCompanyRequest) (common.Hash, error) {
    tx, err := hp.HoquPlatform.RegisterCompany(
        hp.Wallet.Account,
        params.Id,
        params.OwnerId,
        common.HexToAddress(params.OwnerAddress),
        params.Name,
        params.DataUrl,
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) SetCompanyStatus(params *models.SetStatusRequest) (common.Hash, error) {
    tx, err := hp.HoquPlatform.SetCompanyStatus(
        hp.Wallet.Account,
        params.Id,
        params.Status,
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) GetCompany(id uint32) (*models.CompanyData, error) {
    company, err := hp.HoquPlatform.Companies(nil, id)
    if err != nil {
        return nil, err
    }

    companyData := &models.CompanyData{
        CreatedAt:    company.CreatedAt.String(),
        OwnerId:      company.OwnerId,
        OwnerAddress: company.OwnerAddress.String(),
        Name:         company.Name,
        DataUrl:      company.DataUrl,
        Status:       company.Status,
    }

    return companyData, nil
}

func (hp *HoquPlatform) RegisterTracker(params *models.RegisterTrackerRequest) (common.Hash, error) {
    tx, err := hp.HoquPlatform.RegisterTracker(
        hp.Wallet.Account,
        params.Id,
        common.HexToAddress(params.OwnerAddress),
        params.Name,
        params.DataUrl,
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) SetTrackerStatus(params *models.SetStatusRequest) (common.Hash, error) {
    tx, err := hp.HoquPlatform.SetTrackerStatus(
        hp.Wallet.Account,
        params.Id,
        params.Status,
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) GetTracker(id uint32) (*models.TrackerData, error) {
    tracker, err := hp.HoquPlatform.Trackers(nil, id)
    if err != nil {
        return nil, err
    }

    trackerData := &models.TrackerData{
        CreatedAt:    tracker.CreatedAt.String(),
        OwnerAddress: tracker.OwnerAddress.String(),
        Name:         tracker.Name,
        DataUrl:      tracker.DataUrl,
        Status:       tracker.Status,
    }

    return trackerData, nil
}

func (hp *HoquPlatform) AddOffer(params *models.AddOfferRequest) (common.Hash, error) {
    ethAmount, ok := big.NewInt(0).SetString(params.Cost, 0)
    if !ok {
        return common.Hash{}, fmt.Errorf("wrong offer cost provided: %s", params.Cost)
    }

    tx, err := hp.HoquPlatform.AddOffer(
        hp.Wallet.Account,
        params.Id,
        params.CompanyId,
        common.HexToAddress(params.PayerAddress),
        params.Name,
        params.DataUrl,
        ethAmount,
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) SetOfferStatus(params *models.SetStatusRequest) (common.Hash, error) {
    tx, err := hp.HoquPlatform.SetOfferStatus(
        hp.Wallet.Account,
        params.Id,
        params.Status,
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) GetOffer(id uint32) (*models.OfferData, error) {
    offer, err := hp.HoquPlatform.Offers(nil, id)
    if err != nil {
        return nil, err
    }

    offerData := &models.OfferData{
        CreatedAt:    offer.CreatedAt.String(),
        CompanyId:    offer.CompanyId,
        PayerAddress: offer.PayerAddress.String(),
        Name:         offer.Name,
        DataUrl:      offer.DataUrl,
        Cost:         offer.Cost.String(),
        Status:       offer.Status,
    }

    return offerData, nil
}

func (hp *HoquPlatform) AddLead(params *models.AddLeadRequest) (common.Hash, error) {
    ethAmount, ok := big.NewInt(0).SetString(params.Price, 0)
    if !ok {
        return common.Hash{}, fmt.Errorf("wrong lead price provided: %s", params.Price)
    }

    tx, err := hp.HoquPlatform.AddLead(
        hp.Wallet.Account,
        params.Id,
        params.OwnerId,
        params.TrackerId,
        params.OfferId,
        common.HexToAddress(params.BeneficiaryAddress),
        params.Meta,
        params.DataUrl,
        ethAmount,
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) AddLeadIntermediary(params *models.AddLeadIntermediaryRequest) (common.Hash, error) {
    tx, err := hp.HoquPlatform.AddLeadIntermediary(
        hp.Wallet.Account,
        params.Id,
        common.HexToAddress(params.Address),
        uint32(params.Percent*1e6),
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) SetLeadStatus(params *models.SetStatusRequest) (common.Hash, error) {
    tx, err := hp.HoquPlatform.SetLeadStatus(
        hp.Wallet.Account,
        params.Id,
        params.Status,
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) SellLead(id uint32) (common.Hash, error) {
    tx, err := hp.HoquPlatform.SellLead(
        hp.Wallet.Account,
        id,
    )
    if err != nil {
        return common.Hash{}, err
    }

    return tx.Hash(), nil
}

func (hp *HoquPlatform) GetLead(id uint32) (*models.LeadData, error) {
    offer, err := hp.HoquPlatform.Leads(nil, id)
    if err != nil {
        return nil, err
    }

    intermediaries := make(map[string]float32)
    for num := uint8(0); num < offer.NumOfIntermediaries; num++ {
        addr, err := hp.HoquPlatform.GetLeadIntermediaryAddress(nil, id, num)
        if err != nil {
            return nil, err
        }
        percent, err := hp.HoquPlatform.GetLeadIntermediaryPercent(nil, id, num)
        if err != nil {
            return nil, err
        }
        intermediaries[addr.String()] = float32(percent) / 1e6
    }

    offerData := &models.LeadData{
        CreatedAt:          offer.CreatedAt.String(),
        OwnerId:            offer.OwnerId,
        TrackerId:          offer.TrackerId,
        OfferId:            offer.OfferId,
        BeneficiaryAddress: offer.BeneficiaryAddress.String(),
        Intermediaries:     intermediaries,
        Meta:               offer.Meta,
        DataUrl:            offer.DataUrl,
        Price:              offer.Price.String(),
        Status:             offer.Status,
    }

    return offerData, nil
}

func (hp *HoquPlatform) Events(addrs []string) ([]sdkModels.ContractEvent, error) {
    hashAddrs := make([]common.Hash, len(addrs))
    for _, addr := range addrs {
        hashAddrs = append(hashAddrs, common.HexToHash(addr))
    }

    events, err := hp.GetEventsByTopics(
        [][]common.Hash{{}, hashAddrs},
        big.NewInt(viper.GetInt64("geth.start_block.platform")),
    )
    if err != nil {
        return nil, err
    }

    for key, event := range events {
        switch {
        case event.Name == "UserRegistered":

            events[key].Args = models.UserRegisteredEventArgs{
                OwnerAddress: common.BytesToAddress(event.RawArgs[0]).String(),
                Id:           common.BytesToHash(event.RawArgs[1]).Big().String(),
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
            event.Name == "LeadStatusChanged":

            events[key].Args = models.StatusChangedEventArgs{
                OwnerAddress: common.BytesToAddress(event.RawArgs[0]).String(),
                Status:       uint8(common.BytesToHash(event.RawArgs[1]).Big().Uint64()),
            }
        case event.Name == "CompanyRegistered" || event.Name == "TrackerRegistered" ||
            event.Name == "OfferAdded":

            events[key].Args = models.IdWithNameEventArgs{
                OwnerAddress: common.BytesToAddress(event.RawArgs[0]).String(),
                Id:           common.BytesToHash(event.RawArgs[1]).Big().String(),
                Name:         string(event.RawArgs[2]),
            }
        case event.Name == "LeadAdded":

            events[key].Args = models.LeadAddedEventArgs{
                OwnerAddress: common.BytesToAddress(event.RawArgs[0]).String(),
                Id:           common.BytesToHash(event.RawArgs[1]).Big().String(),
                Price:        string(event.RawArgs[2]),
            }
        default:
            return nil, fmt.Errorf("unknown event type: %s", event.Name)
        }
    }

    return events, nil
}
