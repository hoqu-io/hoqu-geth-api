package geth

import (
    "hoqu-geth-api/contract/platform"
    "github.com/ethereum/go-ethereum/common"
    "github.com/spf13/viper"
    "errors"
    "fmt"
    "hoqu-geth-api/sdk/geth"
    "github.com/ethereum/go-ethereum/core/types"
    "hoqu-geth-api/geth/models"
    "github.com/satori/go.uuid"
)

var storage *HoQuStorage

type HoQuStorage struct {
    *geth.Contract
    HoQuStorage *platform.HoQuStorage
}

func initHoQuStorage() error {
    c := geth.NewContract(viper.GetString("geth.addr.storage"))
    c.InitEvents(platform.HoQuStorageABI)

    s, err := platform.NewHoQuStorage(c.Address, c.Wallet.Connection)
    if err != nil {
        return errors.New(fmt.Sprintf("Failed to instantiate a HoQu Platform Storage contract: %v", err))
    }

    storage = &HoQuStorage{
        Contract:   c,
        HoQuStorage: s,
    }

    return nil
}

func GetHoQuStorage() *HoQuStorage {
    return storage
}

func (s *HoQuStorage) Deploy() (*common.Address, *types.Transaction, error) {
    address, tx, _, err := platform.DeployHoQuStorage(
        s.Wallet.Account,
        s.Wallet.Connection,
        GetHoQuConfig().Address,
    )
    if err != nil {
        return nil, nil, fmt.Errorf("failed to deploy HoQuStorage contract: %v", err)
    }
    return &address, tx, nil
}

func (s *HoQuStorage) GetUser(id string) (*models.UserData, error) {
    uid, err := uuid.FromString(id)
    if err != nil {
        return nil, err
    }

    user, err := s.HoQuStorage.Users(nil, uid)
    if err != nil {
        return nil, err
    }

    addresses := make(map[uint8]string)
    for num := uint8(0); num < user.NumOfAddresses; num++ {
        addr, err := s.HoQuStorage.GetUserAddress(nil, uid, num)
        if err != nil {
            return nil, err
        }
        addresses[num] = addr.String()
    }

    userData := &models.UserData{
        CreatedAt:  user.CreatedAt.String(),
        Addresses:  addresses,
        Role:       user.Role,
        KycLevel:   user.KycLevel,
        PubKey:     user.PubKey,
        Status:     models.Status(user.Status),
    }

    return userData, nil
}

func (s *HoQuStorage) GetIdentification(id string) (*models.IdentificationData, error) {
    uid, err := uuid.FromString(id)
    if err != nil {
        return nil, err
    }

    ident, err := s.HoQuStorage.Ids(nil, uid)
    if err != nil {
        return nil, err
    }

    reports := make(map[uint16]models.KycReport)
    for num := uint16(0); num < ident.NumOfKycReports; num++ {
        createdAt, meta, kycLevel, dataUrl, err := s.HoQuStorage.GetKycReport(nil, uid, num)
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

    cid, err := uuid.FromBytes(ident.CompanyId[:])
    if err != nil {
        return nil, err
    }

    userid, err := uuid.FromBytes(ident.UserId[:])
    if err != nil {
        return nil, err
    }

    identData := &models.IdentificationData{
        CreatedAt: ident.CreatedAt.String(),
        CompanyId: cid.String(),
        UserId: userid.String(),
        KycReports: reports,
        Name: ident.Name,
        IdType: ident.IdType,
        Status: models.Status(ident.Status),
    }

    return identData, nil
}

func (s *HoQuStorage) GetStats(id string) (*models.StatsData, error) {
    cid, err := uuid.FromString(id)
    if err != nil {
        return nil, err
    }

    stats, err := s.HoQuStorage.Stats(nil, cid)
    if err != nil {
        return nil, err
    }

    statsData := &models.StatsData{
        Rating: stats.Rating.Uint64(),
        Volume: stats.Volume.Uint64(),
        Contragents: stats.Contragents.Uint64(),
        Stat1: stats.Stat1.Uint64(),
        Stat2: stats.Stat2.Uint64(),
        Status: models.Status(stats.Status),
    }

    return statsData, nil
}

func (s *HoQuStorage) GetCompany(id string) (*models.CompanyData, error) {
    cid, err := uuid.FromString(id)
    if err != nil {
        return nil, err
    }

    company, err := s.HoQuStorage.Companies(nil, cid)
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
        Name:         company.Name,
        DataUrl:      company.DataUrl,
        Status:       models.Status(company.Status),
    }

    return companyData, nil
}

func (s *HoQuStorage) GetNetwork(id string) (*models.NetworkData, error) {
    cid, err := uuid.FromString(id)
    if err != nil {
        return nil, err
    }

    network, err := s.HoQuStorage.Networks(nil, cid)
    if err != nil {
        return nil, err
    }

    oid, err := uuid.FromBytes(network.OwnerId[:])
    if err != nil {
        return nil, err
    }

    networkData := &models.NetworkData{
        CreatedAt:    network.CreatedAt.String(),
        OwnerId:      oid.String(),
        Name:         network.Name,
        DataUrl:      network.DataUrl,
        Status:       models.Status(network.Status),
    }

    return networkData, nil
}

func (s *HoQuStorage) GetTracker(id string) (*models.TrackerData, error) {
    tid, err := uuid.FromString(id)
    if err != nil {
        return nil, err
    }

    tracker, err := s.HoQuStorage.Trackers(nil, tid)
    if err != nil {
        return nil, err
    }

    oid, err := uuid.FromBytes(tracker.OwnerId[:])
    if err != nil {
        return nil, err
    }

    nid, err := uuid.FromBytes(tracker.NetworkId[:])
    if err != nil {
        return nil, err
    }

    trackerData := &models.TrackerData{
        CreatedAt:    tracker.CreatedAt.String(),
        OwnerId: oid.String(),
        NetworkId: nid.String(),
        Name:         tracker.Name,
        DataUrl:      tracker.DataUrl,
        Status:       models.Status(tracker.Status),
    }

    return trackerData, nil
}

func (s *HoQuStorage) GetOffer(id string) (*models.OfferData, error) {
    uid, err := uuid.FromString(id)
    if err != nil {
        return nil, err
    }

    offer, err := s.HoQuStorage.Offers(nil, uid)
    if err != nil {
        return nil, err
    }

    oid, err := uuid.FromBytes(offer.OwnerId[:])
    if err != nil {
        return nil, err
    }

    nid, err := uuid.FromBytes(offer.NetworkId[:])
    if err != nil {
        return nil, err
    }

    mid, err := uuid.FromBytes(offer.MerchantId[:])
    if err != nil {
        return nil, err
    }

    offerData := &models.OfferData{
        CreatedAt:    offer.CreatedAt.String(),
        OwnerId: oid.String(),
        NetworkId: nid.String(),
        MerchantId:    mid.String(),
        PayerAddress: offer.PayerAddress.String(),
        Name:         offer.Name,
        DataUrl:      offer.DataUrl,
        Cost:         offer.Cost.String(),
        Status:       models.Status(offer.Status),
    }

    return offerData, nil
}

func (s *HoQuStorage) GetAd(id string) (*models.AdStorageData, error) {
    aid, err := uuid.FromString(id)
    if err != nil {
        return nil, err
    }

    ad, err := s.HoQuStorage.AdCampaigns(nil, aid)
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

    adData := &models.AdStorageData{
        CreatedAt:          ad.CreatedAt.String(),
        OwnerId:            oid.String(),
        OfferId:            offid.String(),
        ContractAddress: ad.ContractAddress.String(),
    }

    return adData, nil
}