package models

import (
    "github.com/satori/go.uuid"
    "time"
)

// swagger:parameters addTariff
type AddTariffParams struct {
    // in: body
    Body AddTariffRequest `json:"body"`
}

type AddTariffRequest struct {
    Id uuid.UUID `json:"-"`
    // the ID of a Tariff Group
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    TariffGroupId string `json:"tariffGroupId"`
    // the name of the Tariff
    Name string `json:"name"`
    // the name of the action when tariff is applied
    Action string `json:"action"`
    // the method of lead price calculation
    // example: percent|fixed
    CalcMethod string `json:"calcMethod"`
    // the price of a lead calculated by the tariff in wei
    // or the percent value
    // example: 25000000000000000
    Price string `json:"price"`
}

type TariffSuccessData struct {
    Tariff *Tariff `json:"Tariff"`
    Update bool `json:"-"`
}

// Success
//
// swagger:response
type TariffDataResponse struct {
    // in: body
    Body struct {
        Data TariffSuccessData `json:"data"`
    }
}

// Tariff Model
//
// swagger:model
type Tariff struct {
    ID string `json:"id" gorm:"type:char(36)"`
    // ISO8601 datetime format
    // example: 2009-11-13T10:39:35Z
    CreatedAt time.Time `json:"createdAt"`
    // the ID of a Tariff Group
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    TariffGroupId string `json:"tariffGroupId"`
    // the name of the Tariff
    Name string `json:"name"`
    // the name of the action when tariff is applied
    Action string `json:"action"`
    // the method of lead price calculation
    // example: percent|fixed
    CalcMethod string `json:"calcMethod"`
    // the price of a lead calculated by the tariff in wei
    // or the percent value
    // example: 25000000000000000
    Price string `json:"price"`
    // example: 3
    Status Status `json:"status"`
}
