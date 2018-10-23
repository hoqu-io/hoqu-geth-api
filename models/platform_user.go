package models

import (
    "time"
    "github.com/satori/go.uuid"
    "github.com/jinzhu/gorm"
    "encoding/json"
)

// swagger:parameters registerUser
type RegisterUserParams struct {
    // in: body
    Body RegisterUserRequest `json:"body"`
}

type RegisterUserRequest struct {
    Id uuid.UUID `json:"-"`
    // Ethereum address of the user
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    Address string `json:"address"`
    // user role
    // example: affiliate
    Role string `json:"role"`
    // all private data available only for that user is encrypted by his public key
    PubKey string `json:"pubKey"`
}

// swagger:parameters addUserAddress
type AddUserAddressParams struct {
    // in: body
    Body AddUserAddressRequest `json:"body"`
}

type AddUserAddressRequest struct {
    // an ID of the user
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    Id string `json:"id"`
    // Ethereum address of the user
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    Address string `json:"address"`
}

type UserRegisteredEventArgs struct {
    OwnerAddress string `json:"ownerAddress"`
    Id           string `json:"id"`
    Role         string `json:"role"`
}

type UserAddressAddedEventArgs struct {
    OwnerAddress      string `json:"ownerAddress"`
    AdditionalAddress string `json:"additionalAddress"`
    Id                string `json:"id"`
}


type UserSuccessData struct {
    User   *User `json:"User"`
    Update bool  `json:"-"`
}

// Success
//
// swagger:response
type UserDataResponse struct {
    // in: body
    Body struct {
        Data UserSuccessData `json:"data"`
    }
}

// User Model
//
// swagger:model
type User struct {
    ID string `json:"id" gorm:"type:char(36)"`
    // ISO8601 datetime format
    // example: 2009-11-13T10:39:35Z
    CreatedAt time.Time `json:"createdAt"`
    // a list of Ethereum addresses of the user
    // example: ["0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD"]
    Addresses map[uint8]string `json:"addresses" gorm:"-"`
    AddressesJson string `json:"-" gorm:"addresses;type:text"`
    // user role
    // example: affiliate
    Role string `json:"role"`
    // Higher the level more trust this user has from other participants
    // example: 3
    KycLevel uint8 `json:"kycLevel"`
    // all private data available only for that user is encrypted by his public key
    PubKey string `json:"pubKey"`
    // example: 3
    Status Status `json:"status"`
}

func (model *User) BeforeSave(scope *gorm.Scope) error {
    tj, err := json.Marshal(model.Addresses)
    if err != nil {
        return err
    }

    model.AddressesJson = string(tj)
    return nil
}

func (model *User) AfterFind(scope *gorm.Scope) error {
    return json.Unmarshal([]byte(model.AddressesJson), &model.Addresses)
}
