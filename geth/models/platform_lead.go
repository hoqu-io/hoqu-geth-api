package models

// swagger:parameters addLead
type AddLeadParams struct {
    // in: body
    Body AddLeadRequest `json:"body"`
}

type AddLeadRequest struct {
    // the ID of an Ad the lead was created for
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    AdId string `json:"adId"`
    // the ID of a Tracker created the lead
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    TrackerId string `json:"trackerId"`
    // Lead meta data in JSON
    // example: {"country": "DE", "ip": "123.44.55.66"}
    Meta string `json:"meta"`
    // the full lead data URL
    DataUrl string `json:"dataUrl"`
    // the final price of the lead in HQX
    // example: 25000000000000000
    Price string `json:"price"`
}

// swagger:parameters addLeadIntermediary
type AddLeadIntermediaryParams struct {
    // in: body
    Body AddLeadIntermediaryRequest `json:"body"`
}

type AddLeadIntermediaryRequest struct {
    // Lead ID the intermediary is adding to
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    Id string `json:"id"`
    // Ethereum address of the intermediary
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    Address string `json:"address"`
    // the percent of a final price of the Lead this intermediary should receive
    // example: 0.045
    Percent float32 `json:"percent"`
}

type LeadAddedEventArgs struct {
    OwnerAddress string `json:"ownerAddress"`
    Id           string `json:"id"`
    Price        string `json:"price"`
}

// swagger:parameters getLead
type GetLeadParams struct {
    // an ID of the requested lead
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    // in: query
    Id string `json:"id"`
}

// Success
//
// swagger:response
type LeadDataResponse struct {
    // in: body
    Body struct {
        Data struct {
            Lead LeadData `json:"Lead"`
        } `json:"data"`
    }
}

// Lead Model
//
// swagger:model
type LeadData struct {
    // unix timestamp
    // example: 1518957879
    CreatedAt string `json:"createdAt"`
    // the ID of an Ad the lead was created for
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    AdId string `json:"adId"`
    // the ID of a Tracker created the lead
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    TrackerId string `json:"trackerId"`
    // Lead meta data in JSON
    Meta string `json:"meta"`
    // the full lead data URL
    DataUrl string `json:"dataUrl"`
    // the final price of the lead in HQX
    // example: 25000000000000000
    Price string `json:"price"`
    // a list of ethAddress -> percent of all money receivers (except owner itself)
    // example: {"0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD": 0.035}
    Intermediaries map[string]float32 `json:"intermediaries"`
    // example: 3
    Status         Status             `json:"status"`
}

