package models

// status, one of 0:NotExists, 1:Created, 2:Pending, 3:Active, 4:Done, 5:Declined
// swagger:model
// example: 3
type Status uint8

const (
    STATUS_NOT_EXIST = Status(0)
    STATUS_CREATED = Status(1)
    STATUS_PENDING = Status(2)
    STATUS_ACTIVE = Status(3)
    STATUS_DONE = Status(4)
    STATUS_DECLINED = Status(5)

    KYC_UNDEFINED = uint8(0)
    KYC_TIER1 = uint8(1)
    KYC_TIER2 = uint8(2)
    KYC_TIER3 = uint8(3)
    KYC_TIER4 = uint8(4)
    KYC_TIER5 = uint8(5)
)

// swagger:parameters setAdStatus setOfferStatus setUserStatus setTrackerStatus setNetworkStatus setCompanyStatus
type SetStatusParams struct {
    // in: body
    Body SetStatusRequest `json:"body"`
}

type SetStatusRequest struct {
    // id of entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    Id string `json:"id"`
    // example: 3
    Status Status `json:"status"`
}

// swagger:parameters setOfferName
type SetNameParams struct {
    // in: body
    Body SetNameRequest `json:"body"`
}

type SetNameRequest struct {
    // id of entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    Id string `json:"id"`
    // example: HoQu LLC
    Name string `json:"name"`
}

// swagger:parameters addOfferTariff
type IdWithChildIdParams struct {
    // in: body
    Body IdWithChildIdRequest `json:"body"`
}

type IdWithChildIdRequest struct {
    // id of entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    Id string `json:"id"`
    // id of child entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    ChildId string `json:"childId"`
}

type OnlyAddressEventArgs struct {
    OwnerAddress string `json:"ownerAddress"`
    Id           string `json:"id"`
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

type IdWithChildIdEventArgs struct {
    OwnerAddress string `json:"ownerAddress"`
    Id           string `json:"id"`
    ChildId      string `json:"childId"`
}
