package models

// status, one of 0:NotExists, 1:Created, 2:Pending, 3:Active, 4:Done, 5:Declined
// swagger:model
// example: 3
type Status uint8

// swagger:parameters setLeadStatus setAdStatus setOfferStatus setUserStatus setTrackerStatus setCompanyStatus
type SetStatusParams struct {
    // in: body
    Body SetStatusRequest `json:"body"`
}

type SetStatusRequest struct {
    // id of entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    Id     string `json:"id"`
    // example: 3
    Status Status `json:"status"`
}

type OnlyAddressEventArgs struct {
    OwnerAddress string `json:"ownerAddress"`
}

type StatusChangedEventArgs struct {
    OwnerAddress string `json:"ownerAddress"`
    Id           string `json:"id"`
    Status       uint8  `json:"status"`
}

type IdWithNameEventArgs struct {
    OwnerAddress string `json:"ownerAddress"`
    Id           string `json:"id"`
    Name         string `json:"name"`
}

// swagger:parameters sellLead
type IdRequestParams struct {
    // in: body
    Body IdRequest `json:"body"`
}

type IdRequest struct {
    // id of entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    Id string `json:"id"`
}

// Success
//
// swagger:response
type AddSuccessResponse struct {
    // in: body
    Body struct {
        Data struct {
            // id of created entity
            // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
            Id string `json:"id"`
            // Ethereum transaction hash
            // example: 0x23682ef776bfb465433e8f6a6e84eab71f039f039e86933aeca20ee46d01d576
            Tx string `json:"tx"`
        } `json:"data"`
    }
}

// Success
//
// swagger:response
type TxSuccessResponse struct {
    // in: body
    Body struct {
        Data struct {
            // Ethereum transaction hash
            // example: 0x23682ef776bfb465433e8f6a6e84eab71f039f039e86933aeca20ee46d01d576
            Tx string `json:"tx"`
        } `json:"data"`
    }
}
