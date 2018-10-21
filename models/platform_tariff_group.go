package models

import (
    "github.com/satori/go.uuid"
    "time"
    "github.com/jinzhu/gorm"
    "encoding/json"
)

// swagger:parameters addTariffGroup
type AddTariffGroupParams struct {
    // in: body
    Body AddTariffGroupRequest `json:"body"`
}

type AddTariffGroupRequest struct {
    Id uuid.UUID `json:"-"`
    // the ID of an User owned the entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OwnerId string `json:"ownerId"`
    // the name of the TariffGroup
    Name string `json:"name"`
}

type TariffGroupSuccessData struct {
    TariffGroup *TariffGroup `json:"TariffGroup"`
    Update bool `json:"-"`
}

// Success
//
// swagger:response
type TariffGroupDataResponse struct {
    // in: body
    Body struct {
        Data TariffGroupSuccessData `json:"data"`
    }
}

// TariffGroup Model
//
// swagger:model
type TariffGroup struct {
    ID string `json:"id" gorm:"type:char(36)"`
    // ISO8601 datetime format
    // example: 2009-11-13T10:39:35Z
    CreatedAt time.Time `json:"createdAt"`
    // the ID of an User owned the entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    OwnerId string `json:"ownerId"`
    // the name of the TariffGroup
    Name string `json:"name"`
    // a list of tariff ids related to the offer
    // example: ["a6fdb91a-149d-11e8-b642-0ed5f89f718b"]
    Tariffs []string `json:"tariffs" gorm:"-"`
    TariffsJson string `json:"-" gorm:"tariffs;type:text"`
    // example: 3
    Status Status `json:"status"`
}

func (model *TariffGroup) BeforeSave(scope *gorm.Scope) error {
    tj, err := json.Marshal(model.Tariffs)
    if err != nil {
        return err
    }

    model.TariffsJson = string(tj)
    return nil
}

func (model *TariffGroup) AfterFind(scope *gorm.Scope) error {
    return json.Unmarshal([]byte(model.TariffsJson), &model.Tariffs)
}
