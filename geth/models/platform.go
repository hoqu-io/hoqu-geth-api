package models

type SetStatusRequest struct {
    Id     uint32 `json:"id"`
    Status uint8  `json:"status"`
}

type OnlyAddressEventArgs struct {
    OwnerAddress string `json:"ownerAddress"`
}

type StatusChangedEventArgs struct {
    OwnerAddress string `json:"ownerAddress"`
    Status       uint8  `json:"status"`
}

type IdWithNameEventArgs struct {
    OwnerAddress string `json:"ownerAddress"`
    Id           string `json:"id"`
    Name         string `json:"name"`
}

type IdRequest struct {
    Id uint32 `json:"id"`
}
