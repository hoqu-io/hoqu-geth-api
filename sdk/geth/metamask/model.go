package metamask

const (
    TypeString = "string"
    TypeAddress = "address"
    TypeBytes32 = "bytes32"
    TypeBool = "bool"
    TypeUint = "uint"
    TypeUint8 = "uint8"
    TypeUint16 = "uint16"
    TypeUint32 = "uint32"
    TypeUint64 = "uint64"
    TypeUint128 = "uint128"
    TypeUint256 = "uint256"
    TypeInt = "int"
    TypeInt8 = "int8"
    TypeInt16 = "int16"
    TypeInt32 = "int32"
    TypeInt64 = "int64"
    TypeInt128 = "int128"
    TypeInt256 = "int256"
)

// TypedParam Model
//
// swagger:model
type TypedParam struct {
    // param name
    // example: challenge
    Name string `json:"name"`
    // param type
    // example: string
    Type string `json:"type"`
    // param value
    // example: 7b9415e1dbbd526c8496f66e3646bbe886245581530d6dfedc9d1e3eb88d879e
    Value interface{} `json:"value"`
}

// swagger:parameters createChallenge
type ChallengeParams struct {
    // in: body
    Body ChallengeRequest `json:"body"`
}

type ChallengeRequest struct {
    // Ethereum address of the signer
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    Address string `json:"address"`
}

// Success
//
// swagger:response
type ChallengeSuccessResponse struct {
    // in: body
    Body struct {
        Data struct {
            Result []TypedParam `json:"result"`
        } `json:"data"`
    }
}

// swagger:parameters recoverSigner
type RecoverParams struct {
    // in: body
    Body RecoverRequest `json:"body"`
}

type RecoverRequest struct {
    // Challenge value hash
    // example: 7b9415e1dbbd526c8496f66e3646bbe886245581530d6dfedc9d1e3eb88d879e
    Challenge string `json:"challenge"`
    // Signature hash
    // example: 0x5c71570e112156aec2091af830fe73256044a089b20d2aa209b93258504af2dd1eefa3bd584a701648fe24c79c4388671dad2470b74f67cfdd14c2f1dcb2fa0f1b
    Signature string `json:"sign"`
}

// Success
//
// swagger:response
type RecoverSuccessResponse struct {
    // in: body
    Body struct {
        Data struct {
            // Ethereum address of the signer
            // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
            Address string `json:"address"`
        } `json:"data"`
    }
}