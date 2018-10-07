package models

import (
    "time"
    "github.com/satori/go.uuid"
    "github.com/jinzhu/gorm"
    "encoding/json"
)

// swagger:parameters addLead
type AddLeadParams struct {
    // in: body
    Body AddLeadRequest `json:"body"`
}

type AddLeadRequest struct {
    Id uuid.UUID `json:"-"`
    // the ID of an Ad the lead was created for
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    AdId string `json:"adId"`
    // the ID of a Tracker created the lead
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    TrackerId string `json:"trackerId"`
    // Lead meta data in JSON
    // example: {"country": "DE", "ip": "123.44.55.66"}
    Meta string `json:"meta"`
    // the full lead data URL
    DataUrl string `json:"dataUrl"`
    // the final price of the lead in HQX
    // example: 25000000000000000
    Price string `json:"price"`
}

// swagger:parameters addLeadIntermediary
type AddLeadIntermediaryParams struct {
    // in: body
    Body AddLeadIntermediaryRequest `json:"body"`
}

type AddLeadIntermediaryRequest struct {
    // the ID of an Ad Campaign
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    AdId string `json:"adId"`
    // Lead ID the intermediary is adding to
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    Id string `json:"id"`
    // Ethereum address of the intermediary
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    Address string `json:"address"`
    // the percent of a final price of the Lead this intermediary should receive
    // example: 0.045
    Percent float32 `json:"percent"`
}

// swagger:parameters setLeadStatus
type SetLeadStatusParams struct {
    // in: body
    Body SetLeadStatusRequest `json:"body"`
}

type SetLeadStatusRequest struct {
    // the ID of an Ad Campaign
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    AdId string `json:"adId"`
    // id of the lead
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    Id string `json:"id"`
    // example: 3
    Status Status `json:"status"`
}

// swagger:parameters setLeadPrice
type SetLeadPriceParams struct {
    // in: body
    Body SetLeadPriceRequest `json:"body"`
}

type SetLeadPriceRequest struct {
    // the ID of an Ad Campaign
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    AdId string `json:"adId"`
    // id of the lead
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    Id string `json:"id"`
    // example: 3
    Price string `json:"price"`
}

// swagger:parameters setLeadDataUrl
type SetLeadDataUrlParams struct {
    // in: body
    Body SetLeadDataUrlRequest `json:"body"`
}

type SetLeadDataUrlRequest struct {
    // the ID of an Ad Campaign
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    AdId string `json:"adId"`
    // id of the lead
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    Id string `json:"id"`
    // example: 3
    DataUrl string `json:"dataUrl"`
}

// swagger:parameters transactLead getLead
type TransactLeadParams struct {
    // in: body
    Body TransactLeadRequest `json:"body"`
}

type TransactLeadRequest struct {
    // the ID of an Ad Campaign
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    AdId string `json:"adId"`
    // id of entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    Id string `json:"id"`
}

type LeadAddedEventArgs struct {
    ContractAddress string `json:"contractAddress"`
    AdId            string `json:"adId"`
    Id              string `json:"id"`
}

type LeadStatusChangedEventArgs struct {
    ContractAddress string `json:"contractAddress"`
    AdId            string `json:"adId"`
    Id              string `json:"id"`
    Status          uint8 `json:"status"`
}

type LeadPriceChangedEventArgs struct {
    ContractAddress string `json:"contractAddress"`
    AdId            string `json:"adId"`
    Id              string `json:"id"`
    Price           string `json:"price"`
}

type LeadDataUrlChangedEventArgs struct {
    ContractAddress string `json:"contractAddress"`
    AdId            string `json:"adId"`
    Id              string `json:"id"`
    DataUrl         string `json:"dataUrl"`
}

// swagger:parameters getLead
type GetLeadParams struct {
    // the ID of an Ad Campaign
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    AdId string `json:"adId"`
    // an ID of the requested lead
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    // in: query
    Id string `json:"id"`
}

type LeadSuccessData struct {
    Lead   *Lead `json:"Lead"`
    Update bool  `json:"-"`
}

// Success
//
// swagger:response
type LeadDataResponse struct {
    // in: body
    Body struct {
        Data LeadSuccessData `json:"data"`
    }
}

// Lead Model
//
// swagger:model
type Lead struct {
    ID string `json:"id" gorm:"type:char(36)"`
    // ISO8601 datetime format
    // example: 2009-11-13T10:39:35Z
    CreatedAt time.Time `json:"createdAt"`
    // the ID of an Ad the lead was created for
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    AdId string `json:"adId"`
    // the ID of a Tracker created the lead
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    TrackerId string `json:"trackerId"`
    // Lead meta data in JSON
    Meta string `json:"meta"`
    // the full lead data URL
    DataUrl string `json:"dataUrl"`
    // the final price of the lead in HQX
    // example: 25000000000000000
    Price string `json:"price"`
    // a list of ethAddress -> percent of all money receivers (except owner itself)
    // example: {"0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD": 0.035}
    Intermediaries     map[string]float32 `json:"intermediaries" gorm:"-"`
    IntermediariesJson string             `json:"-" gorm:"intermediaries;type:text"`
    // example: 3
    Status Status `json:"status"`
}

func (model *Lead) BeforeSave(scope *gorm.Scope) error {
    tj, err := json.Marshal(model.Intermediaries)
    if err != nil {
        return err
    }

    model.IntermediariesJson = string(tj)
    return nil
}

func (model *Lead) AfterFind(scope *gorm.Scope) error {
    return json.Unmarshal([]byte(model.IntermediariesJson), &model.Intermediaries)
}
