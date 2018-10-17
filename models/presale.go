package models

type PresaleDeployParams struct {
    BankAddress        string `json:"bankAddress"`
    BeneficiaryAddress string `json:"beneficiaryAddress"`
    TokenRate          string `json:"tokenRate"`
    MinBuyableAmount   string `json:"minBuyableAmount"`
    MaxTokensAmount    string `json:"maxTokensAmount"`
    EndDate            int64  `json:"endDate"`
}

type TokenAddedEventArgs struct {
    Address     string `json:"address"`
    TokenAmount string `json:"tokenAmount"`
    EtherAmount string `json:"ethAmount"`
}

type TokenSentEventAgrs struct {
    Address     string `json:"address"`
    TokenAmount string `json:"tokenAmount"`
}

type Summary struct {
    Address            string `json:"address"`
    MinBuyableAmount   string `json:"minBuyableAmount"`
    MaxTokensAmount    string `json:"maxTokensAmount"`
    IssuedTokensAmount string `json:"issuedTokensAmount"`
    TokenRate          string `json:"tokenRate"`
    ReceiversCount     uint32 `json:"receiversCount"`
    IsFinished         bool   `json:"isFinished"`
}
