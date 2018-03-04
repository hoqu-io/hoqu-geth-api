package models

// swagger:parameters registerUser
type RegisterUserParams struct {
    // in: body
    Body RegisterUserRequest `json:"body"`
}

type RegisterUserRequest struct {
    // Ethereum address of the user
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    Address string `json:"address"`
    // user role
    // example: affiliate
    Role    string `json:"role"`
    // all private data available only for that user is encrypted by his public key
    PubKey  string `json:"pubKey"`
}

// swagger:parameters addUserKyc
type AddUserKycReportParams struct {
    // in: body
    Body AddUserKycReportRequest `json:"body"`
}

type AddUserKycReportRequest struct {
    // an ID of the user
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

// swagger:parameters addUserAddress
type AddUserAddressParams struct {
    // in: body
    Body AddUserAddressRequest `json:"body"`
}

type AddUserAddressRequest struct {
    // an ID of the user
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    Id      string `json:"id"`
    // Ethereum address of the user
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    Address string `json:"address"`
}

type UserRegisteredEventArgs struct {
    OwnerAddress string `json:"ownerAddress"`
    Id           string `json:"id"`
    Role         string `json:"role"`
}

type UserAddressAddedEventArgs struct {
    OwnerAddress      string `json:"ownerAddress"`
    AdditionalAddress string `json:"additionalAddress"`
}

type UserKycReportAddedEventArgs struct {
    OwnerAddress string `json:"ownerAddress"`
    KycLevel     uint8  `json:"kycLevel"`
}


type KycReport struct {
    CreatedAt string `json:"createdAt"`
    Meta      string `json:"meta"`
    KycLevel  uint8  `json:"kycLevel"`
    DataUrl   string `json:"dataUrl"`
}

// swagger:parameters getUser
type GetUserParams struct {
    // an ID of the requested entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    // in: query
    Id string `json:"id"`
}

// Success
//
// swagger:response
type UserDataResponse struct {
    // in: body
    Body struct {
        Data struct {
            Lead UserData `json:"User"`
        } `json:"data"`
    }
}

// User Model
//
// swagger:model
type UserData struct {
    // unix timestamp
    // example: 1518957879
    CreatedAt  string               `json:"createdAt"`
    // a list of Ethereum addresses of the user
    // example: ["0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD"]
    Addresses  map[uint8]string     `json:"addresses"`
    // user role
    // example: affiliate
    Role       string               `json:"role"`
    // Higher the level more trust this user has from other participants
    // example: 3
    KycLevel   uint8                `json:"kycLevel"`
    // a list of all KYC reports
    KycReports map[uint16]KycReport `json:"kycReports"`
    // all private data available only for that user is encrypted by his public key
    PubKey     string               `json:"pubKey"`
    // example: 3
    Status     Status                `json:"status"`
}
