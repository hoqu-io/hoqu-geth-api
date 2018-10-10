package models

type TokenDeployParams struct {
    TotalSupply string `json:"totalSupply"`
}

type TokenTransferEventArgs struct {
    From   string `json:"from"`
    To     string `json:"to"`
    Amount string `json:"amount"`
}

type TokenApprovalEventArgs struct {
    Owner   string `json:"owner"`
    Spender string `json:"spender"`
    Amount  string `json:"amount"`
}

// swagger:parameters getHqxBalance
type GetHqxBalanceParams struct {
    // Ethereum address
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    // in: query
    Address string `json:"address"`
}

// swagger:parameters getHqxBalances
type HqxAddressesParams struct {
    // in: body
    Body Addresses `json:"body"`
}

// swagger:parameters getHqxAllowance
type HqxAllowanceParams struct {
    // in: body
    Body AllowanceRequest `json:"body"`
}

type AllowanceRequest struct {
    // Ethereum address of token's owner
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    Owner string `json:"owner"`
    // Ethereum address of a spender
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    Spender string `json:"spender"`
}

// Success
//
// swagger:response
type GetAllowanceSuccessResponse struct {
    // in: body
    Body struct {
        Data struct {
            // HQX allowance in wei
            // example: 42340000000000000000
            Allowance string `json:"allowance"`
        } `json:"data"`
    }
}
