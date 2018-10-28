package models

import (
    "github.com/ethereum/go-ethereum/common"
    "hoqu-geth-api/models"
)

// Success
//
// swagger:response
type ContractEventResponse struct {
    // in: body
    Body struct {
        Data struct {
            Events struct{
                // event name
                // example: LeadStatusChanged
                Name        string      `json:"name"`
                // event arguments, differ for each event
                Args        struct{
                    // Ethereum address this events is linked to
                    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
                    OwnerAddress string `json:"ownerAddress"`
                    // entity ID this event is linked to
                    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
                    Id           string `json:"id"`
                    // new status of entity
                    // example: 3
                    Status       models.Status  `json:"status"`
                } `json:"args"`
                // Ethereum transaction hash
                // example: 0x23682ef776bfb465433e8f6a6e84eab71f039f039e86933aeca20ee46d01d576
                TxHash      string `json:"transactionHash"`
                // Ethereum block number event was raised in
                // example: 4589232
                BlockNumber uint64      `json:"blockNumber"`
                BlockHash   string `json:"blockHash"`
                BlockTime   string      `json:"blockTime"`
                TxIndex     uint        `json:"transactionIndex"`
                Removed     bool        `json:"removed"`
            } `json:"events"`
        } `json:"data"`
    }
}

type ContractEvent struct {
    Name        string      `json:"name"`
    Args        interface{} `json:"args"`
    RawArgs     [][]byte    `json:"-"`
    TxHash      common.Hash `json:"transactionHash"`
    BlockNumber uint64      `json:"blockNumber"`
    BlockHash   common.Hash `json:"blockHash"`
    BlockTime   string      `json:"blockTime"`
    TxIndex     uint        `json:"transactionIndex"`
    Removed     bool        `json:"removed"`
}

// swagger:parameters events
type EventsParams struct {
    // in: body
    Body Events `json:"body"`
}

type Events struct {
    // filter events by list of Ethereum addresses
    Addresses []string `json:"addresses"`
    // filter events by list of event names
    EventNames []string `json:"eventNames"`
    // get events from specified block
    StartBlock int64 `json:"startBlock"`
    // get events to specified block
    EndBlock int64 `json:"endBlock"`
}