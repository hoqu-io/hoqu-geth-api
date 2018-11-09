package ad_campaign

import (
    "hoqu-geth-api/geth"
    "sync"
    "context"
    "fmt"
    "hoqu-geth-api/models"
    "hoqu-geth-api/sdk/entity"
)

var (
    gs *GethStorage
    gsOnce sync.Once
)

type GethStorage struct {
    *entity.Storage
    HoquPlatform *geth.HoquPlatform
    HoquStorage *geth.HoQuStorage
    HoquConfig *geth.HoQuConfig
}

func NewGethStorage(nm string, hp *geth.HoquPlatform, hs *geth.HoQuStorage, hc *geth.HoQuConfig) *GethStorage {
    return &GethStorage{
        Storage: entity.NewStorage(nm),
        HoquPlatform: hp,
        HoquStorage: hs,
        HoquConfig: hc,
    }
}

func InitGethStorage() (err error) {
    gsOnce.Do(func() {
        err = geth.InitGeth()
        gs = NewGethStorage(
            "ethereum: ad_campaign",
            geth.GetHoquPlatform(),
            geth.GetHoQuStorage(),
            geth.GetHoQuConfig(),
        )
    })
    return
}

func GetGethStorage() *GethStorage {
    return gs
}

func (s *GethStorage) Op(ctx context.Context, opName string, input interface{}, output interface{}) (err error) {
    span, spanCtx := s.CreateSpan(ctx, opName, input)
    defer s.CloseSpan(span, output, &err)

    switch opName {
    case entity.CREATE:
        err = s.Create(spanCtx, input.(*models.AddAdCampaignRequest), output.(*models.CreateAdCampaignResponseData))
    case entity.SET_STATUS:
        err = s.SetStatus(spanCtx, input.(*models.SetStatusRequest), output.(*models.TxSuccessData))
    case entity.GET_BY_ID:
        err = s.GetById(spanCtx, input.(*models.IdRequest), output.(*models.AdCampaignSuccessData))
    default:
        err = fmt.Errorf("%s: op %s is not supported", s.GetName(), opName)
    }

    return
}

func (s *GethStorage) Create(spanCtx context.Context, input *models.AddAdCampaignRequest, output *models.CreateAdCampaignResponseData) error {

    offReq := &models.IdRequest{
        Id: input.OfferId,
    }
    offer, err := s.HoquStorage.GetOffer(offReq)
    if err != nil {
        output.Failed = true
        return err
    }

    deployParams := &models.DeployAdContractRequest{
        OwnerId: input.OwnerId,
        OfferId: input.OfferId,
        AdId: input.Id.String(),
        BeneficiaryAddress: input.BeneficiaryAddress,
        PayerAddress: offer.PayerAddress,
    }
    contractAddress, tx, err := geth.DeployAdCampaign(deployParams)
    if err != nil {
        output.Failed = true
        return err
    }

    adToStorageReq := &models.AddAdToStorageRequest{
        OwnerId: input.OwnerId,
        OfferId: input.OfferId,
        AdId: input.Id.String(),
        ContractAddress: contractAddress.String(),
    }

    _, err = s.HoquPlatform.AddAdCampaign(adToStorageReq)
    if err != nil {
        output.Failed = true
        return err
    }

    _, err = s.HoquConfig.AddOwner(adToStorageReq.ContractAddress)
    if err != nil {
        output.Failed = true
        return err
    }

    output.Tx = tx.Hash().String()
    output.ContractAddress = contractAddress.String()
    return nil
}

func (s *GethStorage) SetStatus(spanCtx context.Context, input *models.SetStatusRequest, output *models.TxSuccessData) error {
    tx, err := s.HoquPlatform.SetAdCampaignStatus(input)
    if err != nil {
        return err
    }

    output.Tx = tx.String()
    return nil
}

func (s *GethStorage) GetById(spanCtx context.Context, input *models.IdRequest, output *models.AdCampaignSuccessData) error {
    adCampaign, err := s.HoquStorage.GetAdCampaign(input.Id)
    if err != nil {
        return err
    }

    adcontract, err := geth.GetAdCampaign(adCampaign.ContractAddress)
    if err != nil {
        return err
    }

    adId, err := adcontract.AdId()
    if err != nil {
        return err
    }

    offId, err := adcontract.OfferId()
    if err != nil {
        return err
    }

    affId, err := adcontract.AffiliateId()
    if err != nil {
        return err
    }

    baddr, err := adcontract.BeneficiaryAddress()
    if err != nil {
        return err
    }

    paddr, err := adcontract.PayerAddress()
    if err != nil {
        return err
    }

    ad := &models.AdCampaign{
        ID: input.Id,
        AdId: adId.String(),
        CreatedAt: adCampaign.CreatedAt,
        OwnerId: affId.String(),
        OfferId: offId.String(),
        ContractAddress: adCampaign.ContractAddress,
        BeneficiaryAddress: baddr.String(),
        PayerAddress: paddr.String(),
        Status: adCampaign.Status,
    }

    output.Ad = ad
    output.Update = true
    s.OpDone(spanCtx)
    return nil
}