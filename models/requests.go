package models

type Address struct {
    Address string `json:"address"`
}

type Addresses struct {
    Addresses []string `json:"addresses"`
}

type AddressWithAmount struct {
    Address string `json:"address"`
    Amount  string `json:"amount"`
}
