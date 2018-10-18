package models

import (
    "github.com/satori/go.uuid"
    "github.com/jinzhu/gorm"
    "encoding/json"
    "time"
)

// swagger:parameters addOffer
type AddOfferParams struct {
    // in: body
    Body AddOfferRequest `json:"body"`
}

type AddOfferRequest struct {
    Id uuid.UUID `json:"-"`
    // the ID of an User owned the entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OwnerId string `json:"ownerId"`
    // the ID of the Network the offer is linked to
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    NetworkId string `json:"networkId"`
    // the ID of a Merchant owning this offer
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    MerchantId string `json:"merchantId"`
    // Ethereum address of the Offer payer (who will pay for leads)
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    PayerAddress string `json:"payerAddress"`
    // the name of the Offer
    Name string `json:"name"`
    // the offer full data URL
    DataUrl string `json:"dataUrl"`
}

// swagger:parameters setOfferPayerAddress
type SetOfferPayerAddressParams struct {
    // in: body
    Body SetOfferPayerAddressRequest `json:"body"`
}

type SetOfferPayerAddressRequest struct {
    // id of the offer
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    Id string `json:"id"`
    // Ethereum address of the Offer payer (who will pay for leads)
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    PayerAddress string `json:"payerAddress"`
}

// swagger:parameters addOfferTariffGroup
type AddOfferTariffGroupParams struct {
    // in: body
    Body IdWithChildIdRequest `json:"body"`
}

type OfferSuccessData struct {
    Offer *Offer `json:"Offer"`
    Update bool `json:"-"`
}

// Success
//
// swagger:response
type OfferDataResponse struct {
    // in: body
    Body struct {
        Data OfferSuccessData `json:"data"`
    }
}

// Offer Model
//
// swagger:model
type Offer struct {
    ID string `json:"id" gorm:"type:char(36)"`
    // ISO8601 datetime format
    // example: 2009-11-13T10:39:35Z
    CreatedAt time.Time `json:"createdAt"`
    // the ID of an User owned the entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OwnerId string `json:"ownerId" gorm:"type:char(36)"`
    // the ID of the Network the offer is linked to
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    NetworkId string `json:"networkId" gorm:"type:char(36)"`
    // the ID of a Merchant owning this offer
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    MerchantId string `json:"merchantId" gorm:"type:char(36)"`
    // Ethereum address of the Offer payer (who will pay for leads)
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    PayerAddress string `json:"payerAddress"`
    // the name of the Offer
    Name string `json:"name"`
    // the offer full data URL
    DataUrl string `json:"dataUrl"`
    // the cost of the leads in HQX
    // example: ["a6fdb91a-149d-11e8-b642-0ed5f89f718b"]
    TariffGroups []string `json:"tariffGroups" gorm:"-"`
    TariffGroupsJson string `json:"-" gorm:"tariffGroups;type:text"`
    // example: 3
    Status Status `json:"status"`
}

func (model *Offer) BeforeSave(scope *gorm.Scope) error {
    tj, err := json.Marshal(model.TariffGroups)
    if err != nil {
        return err
    }

    model.TariffGroupsJson = string(tj)
    return nil
}

func (model *Offer) AfterFind(scope *gorm.Scope) error {
    return json.Unmarshal([]byte(model.TariffGroupsJson), &model.TariffGroups)
}
