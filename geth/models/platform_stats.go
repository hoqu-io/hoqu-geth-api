package models

// swagger:parameters getStats
type GetStatsParams struct {
    // an ID of the requested entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    // in: query
    Id string `json:"id"`
}

// Success
//
// swagger:response
type StatsDataResponse struct {
    // in: body
    Body struct {
        Data struct {
            Stats StatsData `json:"Stats"`
        } `json:"data"`
    }
}

// Stats Model
//
// swagger:model
type StatsData struct {
    // current rating
    // example: 650
    Rating uint64 `json:"rating"`
    // the total volume in HQX
    // example: 542332
    Volume uint64 `json:"volume"`
    // number of related contragents
    // example: 34
    Contragents uint64 `json:"contragents"`
    // reserved stat for future needs
    // example: 500
    Stat1 uint64 `json:"stat1"`
    // reserved stat for future needs
    // example: 245
    Stat2 uint64 `json:"stat2"`
    // example: 3
    Status Status `json:"status"`
}
