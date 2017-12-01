package models

type RegisterUserRequest struct {
    Id      uint32 `json:"id"`
    Address string `json:"address"`
    Role    string `json:"role"`
    PubKey  string `json:"pubKey"`
}

type AddUserKycReportRequest struct {
    Id       uint32 `json:"id"`
    ReportId uint32 `json:"reportId"`
    KycLevel uint8  `json:"kycLevel"`
    DataUrl  string `json:"dataUrl"`
}

type AddUserAddressRequest struct {
    Id      uint32 `json:"id"`
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
    ReportId  uint32 `json:"reportId"`
    KycLevel  uint8  `json:"kycLevel"`
    DataUrl   string `json:"dataUrl"`
}

type UserData struct {
    CreatedAt  string               `json:"createdAt"`
    Addresses  map[uint8]string     `json:"addresses"`
    Role       string               `json:"role"`
    KycLevel   uint8                `json:"kycLevel"`
    KycReports map[uint16]KycReport `json:"kycReports"`
    PubKey     string               `json:"pubKey"`
    Status     uint8                `json:"status"`
}
