package models

// swagger:parameters addIdentification
type AddIdentificationParams struct {
    // in: body
    Body AddIdentificationRequest `json:"body"`
}

type AddIdentificationRequest struct {
    // an ID of the user
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    UserId       string `json:"userId"`
    // identification type
    // example: person|company
    IdType    string `json:"idType"`
    // name
    // example: John Snow
    Name    string `json:"name"`
    // an ID of the company for company type identifications
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    CompanyId       string `json:"companyId"`
}

// swagger:parameters addKyc
type AddKycReportParams struct {
    // in: body
    Body AddKycReportRequest `json:"body"`
}

type AddKycReportRequest struct {
    // an ID of the identification
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    Id       string `json:"id"`
    // Meta data of the report in string format (JSON for example)
    Meta string `json:"meta"`
    // Higher the level more trust this user has from other participants
    // example: 3
    KycLevel uint8  `json:"kycLevel"`
    // the URL with full report data
    DataUrl  string `json:"dataUrl"`
}

type IdentificationAddedEventArgs struct {
    OwnerAddress string `json:"ownerAddress"`
    Id           string `json:"id"`
    Name         string `json:"name"`
}

type KycReportAddedEventArgs struct {
    OwnerAddress string `json:"ownerAddress"`
    KycLevel     uint8  `json:"kycLevel"`
}


type KycReport struct {
    CreatedAt string `json:"createdAt"`
    Meta      string `json:"meta"`
    KycLevel  uint8  `json:"kycLevel"`
    DataUrl   string `json:"dataUrl"`
}

// swagger:parameters getIdentification
type GetIdentificationParams struct {
    // an ID of the requested entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    // in: query
    Id string `json:"id"`
}

// Success
//
// swagger:response
type IdentificationDataResponse struct {
    // in: body
    Body struct {
        Data struct {
            Identification IdentificationData `json:"identification"`
        } `json:"data"`
    }
}

// User Model
//
// swagger:model
type IdentificationData struct {
    // unix timestamp
    // example: 1518957879
    CreatedAt  string               `json:"createdAt"`
    // an ID of the user
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    UserId       string `json:"userId"`
    // identification type
    // example: person|company
    IdType    string `json:"idType"`
    // name
    // example: John Snow
    Name    string `json:"name"`
    // an ID of the company for company type identifications
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    CompanyId       string `json:"companyId"`
    // a list of all KYC reports
    KycReports map[uint16]KycReport `json:"kycReports"`
    // example: 3
    Status     Status                `json:"status"`
}
