package models

type RegisterTrackerRequest struct {
    Id           uint32 `json:"id"`
    OwnerAddress string `json:"ownerAddress"`
    Name         string `json:"name"`
    DataUrl      string `json:"dataUrl"`
}

type TrackerData struct {
    CreatedAt    string `json:"createdAt"`
    OwnerAddress string `json:"ownerAddress"`
    Name         string `json:"name"`
    DataUrl      string `json:"dataUrl"`
    Status       uint8  `json:"status"`
}
