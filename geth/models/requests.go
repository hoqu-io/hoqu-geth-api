package models

// swagger:parameters getEthBalance
type GetEthBalanceParams struct {
    // Ethereum address
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    // in: query
    Address string `json:"address"`
}

// swagger:parameters getHqxBalance
type GetHqxBalanceParams struct {
    // Ethereum address
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    // in: query
    Address string `json:"address"`
}

type Address struct {
    Address string `json:"address"`
}

// swagger:parameters getEthBalances
type EthAddressesParams struct {
    // in: body
    Body Addresses `json:"body"`
}

// swagger:parameters getHqxBalances
type HqxAddressesParams struct {
    // in: body
    Body Addresses `json:"body"`
}

type Addresses struct {
    Addresses []string `json:"addresses"`
}

type AddressWithAmount struct {
    Address string `json:"address"`
    Amount  string `json:"amount"`
}

