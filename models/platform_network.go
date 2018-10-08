package models

// swagger:parameters registerNetwork
type RegisterNetworkParams struct {
    // in: body
    Body RegisterNetworkRequest `json:"body"`
}

type RegisterNetworkRequest struct {
    // the ID of an User owned the entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OwnerId string `json:"ownerId"`
    // the name of the network
    // example: HOQU Net
    Name string `json:"name"`
    // the company description URL
    DataUrl string `json:"dataUrl"`
}

// swagger:parameters getNetwork
type GetNetworkParams struct {
    // an ID of the requested entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    // in: query
    Id string `json:"id"`
}

// Success
//
// swagger:response
type NetworkDataResponse struct {
    // in: body
    Body struct {
        Data struct {
            Network NetworkData `json:"Network"`
        } `json:"data"`
    }
}

// Network Model
//
// swagger:model
type NetworkData struct {
    // unix timestamp
    // example: 1518957879
    CreatedAt string `json:"createdAt"`
    // the ID of an User owned the entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OwnerId string `json:"ownerId"`
    // the name of the Network
    Name string `json:"name"`
    // the Network full data URL
    DataUrl string `json:"dataUrl"`
    // example: 3
    Status Status `json:"status"`
}
