package models

type BountyDeployParams struct {
    BankAddress        string `json:"bankAddress"`
    BeneficiaryAddress string `json:"beneficiaryAddress"`
}

type BountySummary struct {
    IssuedTokensAmount string `json:"issuedTokensAmount"`
    ReceiversCount     uint32 `json:"receiversCount"`
    IsFinished         bool   `json:"isFinished"`
}
