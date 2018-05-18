package models

// swagger:parameters addAd
type AddAdParams struct {
    // in: body
    Body AddAdRequest `json:"body"`
}

type AddAdRequest struct {
    // the ID of an User owned the entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OwnerId            string `json:"ownerId"`
    // the ID of an Offer the entity is linked to
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OfferId            string `json:"offerId"`
    // Ethereum address of the beneficiary (Ad money receiver)
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    BeneficiaryAddress string `json:"beneficiaryAddress"`
}

type DeployAdContractRequest struct {
    // the ID of an User owned the entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OwnerId            string `json:"ownerId"`
    // the ID of an Offer the entity is linked to
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OfferId            string `json:"offerId"`
    // the ID of an Ad Campaign
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    AdId            string `json:"adId"`
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
    AdId            string `json:"adId"`
    // the ID of an User owned the entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OwnerId            string `json:"ownerId"`
    // the ID of an Offer the entity is linked to
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OfferId            string `json:"offerId"`
    // Ethereum address of the contract serving the ad campaign
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    ContractAddress string `json:"contractAddress"`
}

type AdAddedEventArgs struct {
    OwnerAddress string `json:"ownerAddress"`
    Id           string `json:"id"`
    ContractAddress string `json:"contractAddress"`
}

// swagger:parameters getAd
type GetAdParams struct {
    // an ID of the requested entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    // in: query
    Id string `json:"id"`
}

// Success
//
// swagger:response
type AdDataResponse struct {
    // in: body
    Body struct {
        Data struct {
            Ad AdData `json:"Ad"`
        } `json:"data"`
    }
}

// Ad Model
//
// swagger:model
type AdData struct {
    // the ID of an Ad Campaign from Ad Contract
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    AdId            string `json:"adId"`
    // unix timestamp
    // example: 1518957879
    CreatedAt          string `json:"createdAt"`
    // the ID of an User owned the entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OwnerId            string `json:"ownerId"`
    // the ID of an Offer the entity is linked to
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OfferId            string `json:"offerId"`
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
    Status       Status  `json:"status"`
}

// Ad Model
//
// swagger:model
type AdStorageData struct {
    // unix timestamp
    // example: 1518957879
    CreatedAt          string `json:"createdAt"`
    // the ID of an User owned the entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OwnerId            string `json:"ownerId"`
    // the ID of an Offer the entity is linked to
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OfferId            string `json:"offerId"`
    // Ethereum address of the contract serving the ad campaign
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    ContractAddress string `json:"contractAddress"`
    // example: 3
    Status       Status  `json:"status"`
}