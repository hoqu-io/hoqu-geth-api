package models

type UserRegisteredEventArgs struct {
    OwnerAddress string `json:"ownerAddress"`
    Id           string `json:"id"`
    Role         string `json:"role"`
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
