package models

import (
    "github.com/ethereum/go-ethereum/common"
)

type ContractEvent struct {
    Name string `json:"name"`
    Args interface{} `json:"args"`
    RawArgs [][]byte `json:"-"`
    TxHash common.Hash `json:"transactionHash"`
    BlockNumber uint64 `json:"blockNumber"`
    BlockHash common.Hash `json:"blockHash"`
    BlockTime string `json:"blockTime"`
    TxIndex uint `json:"transactionIndex"`
    Removed bool `json:"removed"`
}
