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
    // number of related members
    // example: 34
    Members uint64 `json:"members"`
    // reserved stat for future needs
    // example: 500
    Alfa uint64 `json:"alfa"`
    // reserved stat for future needs
    // example: 245
    Beta uint64 `json:"beta"`
    // example: 3
    Status Status `json:"status"`
}
