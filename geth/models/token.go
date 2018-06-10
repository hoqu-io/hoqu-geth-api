package models

type TokenDeployParams struct {
    TotalSupply  string `json:"totalSupply"`
}

type TokenTransferEventArgs struct {
    From string `json:"from"`
    To string `json:"to"`
    Amount string `json:"amount"`
}

type TokenApprovalEventArgs struct {
    Owner string `json:"owner"`
    Spender string `json:"spender"`
    Amount string `json:"amount"`
}
