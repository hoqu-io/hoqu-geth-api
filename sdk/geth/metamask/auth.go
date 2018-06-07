package metamask

import (
    "fmt"
    "github.com/miguelmota/go-solidity-sha3"
    "math/big"
    "github.com/ethereum/go-ethereum/common"
    "github.com/satori/go.uuid"
    "hoqu-geth-api/sdk/geth"
)

var auth *Auth

type Auth struct {
    Banner string
}

func NewAuth(banner string) *Auth {
    return &Auth{
        Banner: banner,
    }
}

func InitAuth(banner string) {
    auth = NewAuth(banner)
}

func GetAuth() *Auth {
    return auth
}

func (a *Auth) CreateChallenge(addr common.Address) ([]*TypedParam, error) {
    u, err := uuid.NewV4()
    if err != nil {
        return nil, err
    }

    challenge := solsha3.SoliditySHA3(
        append(addr.Bytes(), u[:]...),
    )

    return a.getChallengeData(common.Bytes2Hex(challenge)), nil
}

func (a *Auth) RecoverSignerAddr(challenge, signed string) (common.Address, error) {
    hash, err := a.GetChallengeHash(challenge)
    if err != nil {
        return common.Address{}, err
    }

    sig := common.FromHex(signed)

    return geth.EcRecover(hash, sig)
}

func (a *Auth) GetChallengeHash(challenge string) ([]byte, error) {
    return a.GetTypedParamsHash(
        a.getChallengeData(challenge),
    )
}

func (a *Auth) GetTypedParamsHash(params []*TypedParam) ([]byte, error) {
    schema := make([][]byte, 0)
    data := make([][]byte, 0)

    for _, param := range params {
        schema = append(schema, solsha3.String(
            fmt.Sprintf("%s %s", param.Type, param.Name),
        ))

        p, err := a.PackTypedParam(param)
        if err != nil {
            return nil, err
        }
        data = append(data, p)
    }

    return solsha3.SoliditySHA3(
        solsha3.SoliditySHA3(schema...),
        solsha3.SoliditySHA3(data...),
    ), nil
}

func (a *Auth) PackTypedParam(param *TypedParam) ([]byte, error) {
    switch {
    case param.Type == TypeString:
        return solsha3.String(param.Value.(string)), nil
    case param.Type == TypeAddress:
        return solsha3.Address(param.Value.(string)), nil
    case param.Type == TypeBytes32:
        return solsha3.Bytes32(param.Value.(string)), nil
    case param.Type == TypeBool:
        return solsha3.Bool(param.Value.(bool)), nil
    case param.Type == TypeUint || param.Type == TypeUint256:
        return solsha3.Uint256FromString(param.Value.(string)), nil
    case param.Type == TypeUint128:
        n := new(big.Int)
        n.SetString(param.Value.(string), 10)
        return solsha3.Uint128(n), nil
    case param.Type == TypeUint64:
        return solsha3.Uint64(param.Value.(uint64)), nil
    case param.Type == TypeUint32:
        return solsha3.Uint32(param.Value.(uint32)), nil
    case param.Type == TypeUint16:
        return solsha3.Uint16(param.Value.(uint16)), nil
    case param.Type == TypeUint8:
        return solsha3.Uint8(param.Value.(uint8)), nil
    case param.Type == TypeInt || param.Type == TypeInt256:
        n := new(big.Int)
        n.SetString(param.Value.(string), 10)
        return solsha3.Int256(n), nil
    case param.Type == TypeInt128:
        n := new(big.Int)
        n.SetString(param.Value.(string), 10)
        return solsha3.Int128(n), nil
    case param.Type == TypeInt64:
        return solsha3.Int64(param.Value.(int64)), nil
    case param.Type == TypeInt32:
        return solsha3.Int32(param.Value.(int32)), nil
    case param.Type == TypeInt16:
        return solsha3.Int16(param.Value.(int16)), nil
    case param.Type == TypeInt8:
        return solsha3.Int8(param.Value.(int8)), nil
    }

    return nil, fmt.Errorf("unknown type provided: %s", param.Type)
}

func (a *Auth) getChallengeData(challenge string) []*TypedParam {
    return []*TypedParam{
        &TypedParam{
            Name: "banner",
            Type: "string",
            Value: a.Banner,
        },
        &TypedParam{
            Name: "challenge",
            Type: "string",
            Value: challenge,
        },
    }
}
