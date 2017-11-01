package models


type Address struct {
    Address string `json:"address"`
}

type Addresses struct {
    Addresses []string `json:"addresses"`
}

type AddressWithAmount struct {
    Address string `json:"address"`
    Amount string `json:"amount"`
}

type PresaleDeployParams struct {
    BankAddress string `json:"bankAddress"`
    BeneficiaryAddress string `json:"beneficiaryAddress"`
    TokenRate string `json:"tokenRate"`
    MinBuyableAmount string `json:"minBuyableAmount"`
    MaxTokensAmount string `json:"maxTokensAmount"`
    EndDate int64 `json:"endDate"`
}