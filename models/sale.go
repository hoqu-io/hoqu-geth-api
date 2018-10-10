package models

type SaleDeployParams struct {
    BankAddress        string `json:"bankAddress"`
    BeneficiaryAddress string `json:"beneficiaryAddress"`
    TokenRate          string `json:"tokenRate"`
    MinBuyableAmount   string `json:"minBuyableAmount"`
    MaxTokensAmount    string `json:"maxTokensAmount"`
    EndDate            int64  `json:"endDate"`
}

type SaleSummary struct {
    Address            string `json:"address"`
    MaxTokensAmount    string `json:"maxTokensAmount"`
    IssuedTokensAmount string `json:"issuedTokensAmount"`
    NextBoundaryAmount string `json:"nextBoundaryAmount"`
    TokenRate          string `json:"tokenRate"`
    ReceiversCount     uint32 `json:"receiversCount"`
    IsFinished         bool   `json:"isFinished"`
}
