package models

// swagger:parameters registerTracker
type RegisterTrackerParams struct {
    // in: body
    Body RegisterTrackerRequest `json:"body"`
}

type RegisterTrackerRequest struct {
    // Ethereum address of the owner
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    OwnerAddress string `json:"ownerAddress"`
    // the name of the Tracker
    Name         string `json:"name"`
    // the Tracker full data URL
    DataUrl      string `json:"dataUrl"`
}

// swagger:parameters getTracker
type GetTrackerParams struct {
    // an ID of the requested entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    // in: query
    Id string `json:"id"`
}

// Success
//
// swagger:response
type TrackerDataResponse struct {
    // in: body
    Body struct {
        Data struct {
            Lead TrackerData `json:"Tracker"`
        } `json:"data"`
    }
}

// Tracker Model
//
// swagger:model
type TrackerData struct {
    // unix timestamp
    // example: 1518957879
    CreatedAt    string `json:"createdAt"`
    // Ethereum address of the owner
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    OwnerAddress string `json:"ownerAddress"`
    // the name of the Tracker
    Name         string `json:"name"`
    // the Tracker full data URL
    DataUrl      string `json:"dataUrl"`
    // example: 3
    Status       Status  `json:"status"`
}
