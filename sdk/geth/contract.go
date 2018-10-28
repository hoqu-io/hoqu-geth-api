package geth

import (
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "strings"
    "context"
    "hoqu-geth-api/sdk/models"
    "fmt"
    "github.com/ethereum/go-ethereum"
    "math/big"
    "regexp"
)

type Contract struct {
    Wallet  *Wallet
    Address common.Address
    Abi     abi.ABI
    Events  map[string]string
    EventHashes  map[string]string
}

func NewContract(addr string) *Contract {
    return &Contract{
        Wallet:  GetWallet(),
        Address: common.HexToAddress(addr),
    }
}

func (c *Contract) InitEvents(contractAbi string) (err error) {
    c.Abi, err = abi.JSON(strings.NewReader(contractAbi))
    if err != nil {
        return
    }

    c.Events = make(map[string]string)
    c.EventHashes = make(map[string]string)
    for _, event := range c.Abi.Events {
        c.Events[event.Id().String()] = event.Name
        c.EventHashes[event.Name] = event.Id().String()
    }

    return
}

func (c *Contract) GetEventsByTopics(request *models.Events, fromBlock int64) ([]models.ContractEvent, error) {
    hashAddrs := make([]common.Hash, len(request.Addresses))
    for _, addr := range request.Addresses {
        hashAddrs = append(hashAddrs, common.HexToHash(addr))
    }

    hashEventNames := make([]common.Hash, len(request.EventNames))
    for _, eventName := range request.EventNames {
        name, ok := c.EventHashes[eventName]
        if !ok {
            return nil, fmt.Errorf("unknown event name provided: %s", eventName)
        }
        hashEventNames = append(hashEventNames, common.HexToHash(name))
    }

    query := ethereum.FilterQuery{
        Addresses: []common.Address{
            c.Address,
        },
        Topics: [][]common.Hash{hashEventNames, hashAddrs},
        FromBlock: big.NewInt(fromBlock),
    }

    if request.StartBlock != 0 {
        query.FromBlock = big.NewInt(request.StartBlock)
    }

    if request.EndBlock != 0 {
        query.ToBlock = big.NewInt(request.EndBlock)
    }

    res, err := c.Wallet.Connection.FilterLogs(context.TODO(), query)
    if err != nil {
        return []models.ContractEvent{}, err
    }

    events := make([]models.ContractEvent, len(res))
    for key, eventLog := range res {
        name, ok := c.Events[eventLog.Topics[0].String()]
        if !ok {
            return events, fmt.Errorf("unknown event parsed: %s", eventLog.Topics[0].String())
        }

        if len(eventLog.Data)%32 != 0 {
            return events, fmt.Errorf("wrong event unindexed data provided: %d", len(eventLog.Data))
        }

        numOfIndexed := len(eventLog.Topics) - 1
        numOfUnindexed := len(eventLog.Data) / 32
        rawArgs := make([][]byte, numOfIndexed+numOfUnindexed)
        if numOfIndexed > 0 {
            for i, indData := range eventLog.Topics[1:] {
                rawArgs[i] = indData.Bytes()
            }
        }
        for i := 0; i < numOfUnindexed; i++ {
            rawArgs[numOfIndexed+i] = eventLog.Data[i*32:(i+1)*32]
        }

        block, err := c.Wallet.Connection.HeaderByHash(context.TODO(), eventLog.BlockHash)
        if err != nil {
            return events, err
        }

        events[key] = models.ContractEvent{
            Name:        name,
            RawArgs:     rawArgs,
            TxHash:      eventLog.TxHash,
            BlockNumber: eventLog.BlockNumber,
            BlockHash:   eventLog.BlockHash,
            BlockTime:   block.Time.String(),
            TxIndex:     eventLog.TxIndex,
            Removed:     eventLog.Removed,
        }
    }

    return events, nil
}

func (c *Contract) FilterString(inp string) string {
    reg, _ := regexp.Compile("[^\\s\\tA-Za-z0-9]+")

    return reg.ReplaceAllString(inp, "")
}
