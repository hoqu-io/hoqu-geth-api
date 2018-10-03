package models

import (
    "time"
    "github.com/satori/go.uuid"
)

// swagger:parameters addAd
type AddAdCampaignParams struct {
    // in: body
    Body AddAdCampaignRequest `json:"body"`
}

type AddAdCampaignRequest struct {
    Id uuid.UUID `json:"-"`
    // the ID of an User owned the entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OwnerId string `json:"ownerId"`
    // the ID of an Offer the entity is linked to
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OfferId string `json:"offerId"`
    // Ethereum address of the beneficiary (Ad money receiver)
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    BeneficiaryAddress string `json:"beneficiaryAddress"`
}

type DeployAdContractRequest struct {
    // the ID of an User owned the entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OwnerId string `json:"ownerId"`
    // the ID of an Offer the entity is linked to
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OfferId string `json:"offerId"`
    // the ID of an Ad Campaign
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    AdId string `json:"adId"`
    // Ethereum address of the beneficiary (Ad money receiver)
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    BeneficiaryAddress string `json:"beneficiaryAddress"`
    // Ethereum address of the Offer payer (who will pay for leads)
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    PayerAddress string `json:"payerAddress"`
}

type AddAdToStorageRequest struct {
    // the ID of an Ad Campaign
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    AdId string `json:"adId"`
    // the ID of an User owned the entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OwnerId string `json:"ownerId"`
    // the ID of an Offer the entity is linked to
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OfferId string `json:"offerId"`
    // Ethereum address of the contract serving the ad campaign
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    ContractAddress string `json:"contractAddress"`
}

type AdAddedEventArgs struct {
    OwnerAddress    string `json:"ownerAddress"`
    Id              string `json:"id"`
    ContractAddress string `json:"contractAddress"`
}

type CreateAdCampaignResponseData struct {
    // id of created entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    Id string `json:"id"`
    // Ethereum transaction hash
    // example: 0x23682ef776bfb465433e8f6a6e84eab71f039f039e86933aeca20ee46d01d576
    Tx string `json:"tx"`
    // Ethereum address of the contract serving the ad campaign
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    ContractAddress string `json:"contractAddress"`
    Failed bool `json:"-"`
}

// Success
//
// swagger:response
type AddAdCampaignSuccessResponse struct {
    // in: body
    Body struct {
        Data CreateAdCampaignResponseData `json:"data"`
    }
}

type AdCampaignSuccessData struct {
    Ad *AdCampaign `json:"Ad"`
    Update bool `json:"-"`
}

// Success
//
// swagger:response
type AdDataResponse struct {
    // in: body
    Body struct {
        Data AdCampaignSuccessData `json:"data"`
    }
}

// Ad Model
//
// swagger:model
type AdCampaign struct {
    ID string `json:"id" gorm:"type:char(36)"`
    // the ID of an Ad Campaign from Ad Contract
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    AdId string `json:"adId"`
    // ISO8601 datetime format
    // example: 2009-11-13T10:39:35Z
    CreatedAt time.Time `json:"createdAt"`
    // the ID of an User owned the entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OwnerId string `json:"ownerId"`
    // the ID of an Offer the entity is linked to
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OfferId string `json:"offerId"`
    // Ethereum address of the beneficiary (Ad money receiver)
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    BeneficiaryAddress string `json:"beneficiaryAddress"`
    // Ethereum address of the Offer payer (who will pay for leads)
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    PayerAddress string `json:"payerAddress"`
    // Ethereum address of the contract serving the ad campaign
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    ContractAddress string `json:"contractAddress"`
    // example: 3
    Status Status `json:"status"`
}

// Ad Model
//
// swagger:model
type AdStorageData struct {
    // ISO8601 datetime format
    // example: 2009-11-13T10:39:35Z
    CreatedAt time.Time `json:"createdAt"`
    // the ID of an User owned the entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OwnerId string `json:"ownerId"`
    // the ID of an Offer the entity is linked to
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OfferId string `json:"offerId"`
    // Ethereum address of the contract serving the ad campaign
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    ContractAddress string `json:"contractAddress"`
    // example: 3
    Status Status `json:"status"`
}
