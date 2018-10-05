package models

// swagger:parameters registerCompany
type RegisterCompanyParams struct {
    // in: body
    Body RegisterCompanyRequest `json:"body"`
}

type RegisterCompanyRequest struct {
    // the ID of an User owned the entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OwnerId string `json:"ownerId"`
    // the name of the company
    // example: HOQU LLC
    Name string `json:"name"`
    // the company description URL
    DataUrl string `json:"dataUrl"`
}

// swagger:parameters getCompany
type GetCompanyParams struct {
    // an ID of the requested entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    // in: query
    Id string `json:"id"`
}

// Success
//
// swagger:response
type CompanyDataResponse struct {
    // in: body
    Body struct {
        Data struct {
            Lead CompanyData `json:"Company"`
        } `json:"data"`
    }
}

// Company Model
//
// swagger:model
type CompanyData struct {
    // unix timestamp
    // example: 1518957879
    CreatedAt string `json:"createdAt"`
    // the ID of an User owned the entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OwnerId string `json:"ownerId"`
    // the name of the company
    // example: HOQU LLC
    Name string `json:"name"`
    // the company description URL
    DataUrl string `json:"dataUrl"`
    // example: 3
    Status Status `json:"status"`
}
