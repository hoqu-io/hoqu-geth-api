package models

type AddOfferRequest struct {
    Id           uint32 `json:"id"`
    CompanyId    uint32 `json:"companyId"`
    PayerAddress string `json:"payerAddress"`
    Name         string `json:"name"`
    DataUrl      string `json:"dataUrl"`
    Cost         string `json:"cost"`
}

type OfferData struct {
    CreatedAt    string `json:"createdAt"`
    CompanyId    uint32 `json:"companyId"`
    PayerAddress string `json:"payerAddress"`
    Name         string `json:"name"`
    DataUrl      string `json:"dataUrl"`
    Cost         string `json:"cost"`
    Status       uint8  `json:"status"`
}
