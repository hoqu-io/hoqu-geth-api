package models

type RegisterCompanyRequest struct {
    Id           uint32 `json:"id"`
    OwnerId      uint32 `json:"ownerId"`
    OwnerAddress string `json:"ownerAddress"`
    Name         string `json:"name"`
    DataUrl      string `json:"dataUrl"`
}

type CompanyData struct {
    CreatedAt    string `json:"createdAt"`
    OwnerId      uint32 `json:"ownerId"`
    OwnerAddress string `json:"ownerAddress"`
    Name         string `json:"name"`
    DataUrl      string `json:"dataUrl"`
    Status       uint8  `json:"status"`
}
