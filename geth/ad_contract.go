package geth

import (
    "hoqu-geth-api/contract/ad"
    "github.com/ethereum/go-ethereum/common"
    "errors"
    "fmt"
    "hoqu-geth-api/sdk/geth"
    "github.com/ethereum/go-ethereum/core/types"
    "hoqu-geth-api/geth/models"
    "github.com/satori/go.uuid"
)

type HoQuAdCampaign struct {
    *geth.Contract
    HoQuAdCampaign *ad.HoQuAdCampaign
}

func GetAdCampaign(addr string) (*HoQuAdCampaign, error) {
    c := geth.NewContract(addr)
    c.InitEvents(ad.HoQuAdCampaignABI)

    s, err := ad.NewHoQuAdCampaign(c.Address, c.Wallet.Connection)
    if err != nil {
        return nil, errors.New(fmt.Sprintf("Failed to instantiate a HoQu Ad Campaign contract: %v", err))
    }

    return &HoQuAdCampaign{
        Contract:   c,
        HoQuAdCampaign: s,
    }, nil
}

func DeployAdCampaign(params *models.DeployAdContractRequest) (*common.Address, *types.Transaction, error) {
    aid, err := uuid.FromString(params.AdId)
    if err != nil {
        return nil, nil, err
    }

    oid, err := uuid.FromString(params.OwnerId)
    if err != nil {
        return nil, nil, err
    }

    offid, err := uuid.FromString(params.OfferId)
    if err != nil {
        return nil, nil, err
    }

    address, tx, _, err := ad.DeployHoQuAdCampaign(
        GetHoquPlatform().Contract.Wallet.Account,
        GetHoquPlatform().Contract.Wallet.Connection,
        GetHoQuConfig().Address,
        GetToken().Address,
        GetHoQuStorage().Address,
        GetHoQuRater().Address,
        aid,
        offid,
        oid,
        common.HexToAddress(params.BeneficiaryAddress),
        common.HexToAddress(params.PayerAddress),
    )
    if err != nil {
        return nil, nil, fmt.Errorf("failed to deploy HoQuAdCampaign contract: %v", err)
    }
    return &address, tx, nil
}

func (adc *HoQuAdCampaign) AdId() (uuid.UUID, error) {
    id, err := adc.HoQuAdCampaign.AdId(nil)
    if err != nil {
        return uuid.UUID{}, err
    }

    return uuid.FromBytes(id[:])
}

func (adc *HoQuAdCampaign) OfferId() (uuid.UUID, error) {
    id, err := adc.HoQuAdCampaign.OfferId(nil)
    if err != nil {
        return uuid.UUID{}, err
    }

    return uuid.FromBytes(id[:])
}

func (adc *HoQuAdCampaign) AffiliateId() (uuid.UUID, error) {
    id, err := adc.HoQuAdCampaign.AffiliateId(nil)
    if err != nil {
        return uuid.UUID{}, err
    }

    return uuid.FromBytes(id[:])
}

func (adc *HoQuAdCampaign) BeneficiaryAddress() (common.Address, error) {
    return adc.HoQuAdCampaign.BeneficiaryAddress(nil)
}

func (adc *HoQuAdCampaign) PayerAddress() (common.Address, error) {
    return adc.HoQuAdCampaign.PayerAddress(nil)
}

func (adc *HoQuAdCampaign) GetLead(id string) (*models.LeadData, error) {
    lid, err := uuid.FromString(id)
    if err != nil {
        return nil, err
    }

    lead, err := adc.HoQuAdCampaign.Leads(nil, lid)
    if err != nil {
        return nil, err
    }

    intermediaries := make(map[string]float32)
    for num := uint8(0); num < lead.NumOfIntermediaries; num++ {
        addr, err := adc.HoQuAdCampaign.GetLeadIntermediaryAddress(nil, lid, num)
        if err != nil {
            return nil, err
        }
        percent, err := adc.HoQuAdCampaign.GetLeadIntermediaryPercent(nil, lid, num)
        if err != nil {
            return nil, err
        }
        intermediaries[addr.String()] = float32(percent) / 1e6
    }

    aid, err := adc.AdId()
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