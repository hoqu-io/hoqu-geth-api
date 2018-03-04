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

type AdAddedEventArgs struct {
    OwnerAddress string `json:"ownerAddress"`
    Id           string `json:"id"`
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
            Lead AdData `json:"Ad"`
        } `json:"data"`
    }
}

// Ad Model
//
// swagger:model
type AdData struct {
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
}
