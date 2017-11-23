package models

type ConfigDeployParams struct {
    CommissionWallet string `json:"commissionWallet"`
}

type SystemOwnerChangedEventArgs struct {
    PreviousOwner string `json:"previousOwner"`
    NewOwner      string `json:"newOwner"`
}

type CommissionWalletChangedEventArgs struct {
    ChangedBy        string `json:"changedBy"`
    CommissionWallet string `json:"commissionWallet"`
}

type CommissionChangedEventArgs struct {
    ChangedBy  string `json:"changedBy"`
    Commission string `json:"commission"`
}
