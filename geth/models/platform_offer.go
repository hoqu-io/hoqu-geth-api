package models

// swagger:parameters addOffer
type AddOfferParams struct {
    // in: body
    Body AddOfferRequest `json:"body"`
}

type AddOfferRequest struct {
    // the ID of an User owned the entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OwnerId      string `json:"ownerId"`
    // the ID of the Network the offer is linked to
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    NetworkId string `json:"networkId"`
    // the ID of a Merchant owning this offer
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    MerchantId    string `json:"merchantId"`
    // Ethereum address of the Offer payer (who will pay for leads)
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    PayerAddress string `json:"payerAddress"`
    // the name of the Offer
    Name         string `json:"name"`
    // the offer full data URL
    DataUrl      string `json:"dataUrl"`
    // the cost of the leads in HQX
    // example: 25000000000000000
    Cost         string `json:"cost"`
}

// swagger:parameters getOffer
type GetOfferParams struct {
    // an ID of the requested entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    // in: query
    Id string `json:"id"`
}

// Success
//
// swagger:response
type OfferDataResponse struct {
    // in: body
    Body struct {
        Data struct {
            Offer OfferData `json:"Offer"`
        } `json:"data"`
    }
}

// Offer Model
//
// swagger:model
type OfferData struct {
    // unix timestamp
    // example: 1518957879
    CreatedAt    string `json:"createdAt"`
    // the ID of an User owned the entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OwnerId      string `json:"ownerId"`
    // the ID of the Network the offer is linked to
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    NetworkId string `json:"networkId"`
    // the ID of a Merchant owning this offer
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    MerchantId    string `json:"merchantId"`
    // Ethereum address of the Offer payer (who will pay for leads)
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    PayerAddress string `json:"payerAddress"`
    // the name of the Offer
    Name         string `json:"name"`
    // the offer full data URL
    DataUrl      string `json:"dataUrl"`
    // the cost of the leads in HQX
    // example: 25000000000000000
    Cost         string `json:"cost"`
    // example: 3
    Status       Status  `json:"status"`
}
