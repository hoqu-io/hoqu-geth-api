package models

type TokenBoughtEvent struct {
    Address string `json:"address"`
    TokenAmount string `json:"tokenAmount"`
    EtherAmount string `json:"ethAmount"`
}

type TokenAddedEvent struct {
    Address string `json:"address"`
    TokenAmount string `json:"tokenAmount"`
    EtherAmount string `json:"ethAmount"`
}

type TokenSentEvent struct {
    Address string `json:"address"`
    TokenAmount string `json:"tokenAmount"`
}

type PresaleEvents struct {
    TokenBought map[string]TokenBoughtEvent `json:"TokenBought"`
    TokenAdded map[string]TokenAddedEvent `json:"TokenAdded"`
    TokenSent map[string]TokenSentEvent `json:"TokenSent"`
}

type Summary struct {
    Address string `json:"address"`
    MinBuyableAmount string `json:"minBuyableAmount"`
    MaxTokensAmount string `json:"maxTokensAmount"`
    IssuedTokensAmount string `json:"issuedTokensAmount"`
    TokenRate string `json:"tokenRate"`
    ReceiversCount uint32 `json:"receiversCount"`
    IsFinished bool `json:"isFinished"`
}