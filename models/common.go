package models

// swagger:parameters getUser getOffer getTariff getAd
type IdRequest struct {
    // id of entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    Id string `json:"id"`
}

type CreateResponseData struct {
    // id of created entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    Id string `json:"id"`
    // Ethereum transaction hash
    // example: 0x23682ef776bfb465433e8f6a6e84eab71f039f039e86933aeca20ee46d01d576
    Tx string `json:"tx"`
    Failed bool `json:"-"`
}

// Success
//
// swagger:response
type AddSuccessResponse struct {
    // in: body
    Body struct {
        Data CreateResponseData `json:"data"`
    }
}

type TxSuccessData struct {
    // Ethereum transaction hash
    // example: 0x23682ef776bfb465433e8f6a6e84eab71f039f039e86933aeca20ee46d01d576
    Tx string `json:"tx"`
}

// Success
//
// swagger:response
type TxSuccessResponse struct {
    // in: body
    Body struct {
        Data TxSuccessData `json:"data"`
    }
}
