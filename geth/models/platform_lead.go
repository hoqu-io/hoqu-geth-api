package models

type AddLeadRequest struct {
    Id                 uint32 `json:"id"`
    OwnerId            uint32 `json:"ownerId"`
    TrackerId          uint32 `json:"trackerId"`
    OfferId            uint32 `json:"offerId"`
    BeneficiaryAddress string `json:"beneficiaryAddress"`
    Meta               string `json:"meta"`
    DataUrl            string `json:"dataUrl"`
    Price              string `json:"price"`
}

type AddLeadIntermediaryRequest struct {
    Id      uint32  `json:"id"`
    Address string  `json:"address"`
    Percent float32 `json:"percent"`
}

type LeadAddedEventArgs struct {
    OwnerAddress string `json:"ownerAddress"`
    Id           string `json:"id"`
    Price        string `json:"price"`
}

type LeadData struct {
    CreatedAt          string             `json:"createdAt"`
    OwnerId            uint32             `json:"ownerId"`
    TrackerId          uint32             `json:"trackerId"`
    OfferId            uint32             `json:"offerId"`
    BeneficiaryAddress string             `json:"beneficiaryAddress"`
    Meta               string             `json:"meta"`
    DataUrl            string             `json:"dataUrl"`
    Price              string             `json:"price"`
    Intermediaries     map[string]float32 `json:"intermediaries"`
    Status             uint8              `json:"status"`
}
