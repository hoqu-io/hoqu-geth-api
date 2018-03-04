package geth

import (
    "hoqu-geth-api/contract"
    "github.com/ethereum/go-ethereum/common"
    "github.com/spf13/viper"
    "errors"
    "fmt"
    "math/big"
    "hoqu-geth-api/sdk/geth"
    "hoqu-geth-api/geth/models"
    sdkModels "hoqu-geth-api/sdk/models"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/sirupsen/logrus"
    "github.com/satori/go.uuid"
    "time"
    "sync"
)

var claim *Claim

type Claim struct {
    *geth.Contract
    Claim *contract.HoQuClaim
    batchLimit int
    batchTimeout time.Duration
    batchStart time.Time
    batch map[common.Address]*big.Int
    BatchId uuid.UUID
    mutex *sync.Mutex
}

func InitClaim() error {
    c := geth.NewContract(viper.GetString("geth.addr.claim"))
    c.InitEvents(contract.HoQuClaimABI)

    s, err := contract.NewHoQuClaim(c.Address, c.Wallet.Connection)
    if err != nil {
        return errors.New(fmt.Sprintf("Failed to instantiate a Claim contract: %v", err))
    }

    u, err := uuid.NewV4()
    if err != nil {
        return err
    }

    claim = &Claim{
        Contract: c,
        Claim:     s,
        batchLimit: viper.GetInt("geth.claim.batch_limit"),
        batchTimeout: time.Duration(viper.GetInt("geth.claim.batch_timeout")) * time.Minute,
        batchStart: time.Now().Add(24 * 365 * time.Hour),
        BatchId: u,
        mutex: &sync.Mutex{},
    }

    return nil
}

func GetClaim() *Claim {
    return claim
}

func (c *Claim) Deploy(params *models.ClaimDeployParams) (*common.Address, *types.Transaction, error) {
    tokenAddr := GetToken().Address

    address, tx, _, err := contract.DeployHoQuClaim(
        c.Wallet.Account,
        c.Wallet.Connection,
        tokenAddr,
        common.HexToAddress(params.BankAddress),
    )
    if err != nil {
        return nil, nil, fmt.Errorf("failed to deploy contract: %v", err)
    }
    return &address, tx, nil
}

func (c *Claim) Events(addrs []string) ([]sdkModels.ContractEvent, error) {
    hashAddrs := make([]common.Hash, len(addrs))
    for _, addr := range addrs {
        hashAddrs = append(hashAddrs, common.HexToHash(addr))
    }

    events, err := c.GetEventsByTopics(
        [][]common.Hash{{}, hashAddrs},
        big.NewInt(viper.GetInt64("geth.start_block.claim")),
    )
    if err != nil {
        return nil, err
    }

    resEvents := make([]sdkModels.ContractEvent, 0)

    for _, event := range events {
        addr := common.BytesToAddress(event.RawArgs[0])

        switch {
        case event.Name == "TokenSent" || event.Name == "AlreadyClaimed":
            event.Args = models.TokenSentEventAgrs{
                Address:     addr.String(),
                TokenAmount: common.BytesToHash(event.RawArgs[1]).Big().String(),
            }
        default:
            return nil, fmt.Errorf("unknown event type: %s", event.Name)
        }

        resEvents = append(resEvents, event)
    }

    return resEvents, nil
}

func (c *Claim) ClaimAddress(addr string, amount string) (common.Hash, error) {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    if c.batchStart.After(time.Now()) {
        c.batch = make(map[common.Address]*big.Int)
        c.batchStart = time.Now()
    }

    err := c.addToBatch(addr, amount)
    if err != nil {
        return common.Hash{}, err
    }

    l1 := len(c.batch)
    if l1 >= c.batchLimit || c.batchStart.Add(c.batchTimeout).Before(time.Now()) {
        defer func(){
            c.batch = nil
            c.batchStart = time.Now().Add(24 * 365 * time.Hour)
            c.BatchId, _ = uuid.NewV4()
        }()

        return c.ClaimAccumulated(c.batch)
    }

    return common.Hash{}, nil
}

func (c *Claim) addToBatch(addr string, amount string) error {
    tokenAmount, ok := big.NewInt(0).SetString(amount, 0)
    if !ok {
        return fmt.Errorf("wrong number provided: %s", amount)
    }

    if ex, ok := c.batch[common.HexToAddress(addr)]; ok {
        //c.batch[common.HexToAddress(addr)] = big.NewInt(0).Add(ex, tokenAmount)
        //logrus.Info("Append ", addr, " to batch")
        return fmt.Errorf("%s already claimed with amount %s", addr, ex.String())
    } else {
        c.batch[common.HexToAddress(addr)] = tokenAmount
        logrus.Info("Add ", addr, " to batch")
    }
    return nil
}

func (c *Claim) ClaimAccumulated(batch map[common.Address]*big.Int) (common.Hash, error) {
    addrs := make([]common.Address, 0)
    amounts := make([]*big.Int, 0)

    for addr, amount := range batch {
        addrs = append(addrs, addr)
        amounts = append(amounts, amount)
    }

    opts, err := c.Wallet.GetTransactOpts()
    if err != nil {
        c.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := c.Claim.ClaimMany(opts, addrs, amounts)
    if err != nil {
        logrus.Error("ClaimMany failed ", err.Error())
        c.Wallet.OnFailTransaction(err)

        if c.Wallet.ValidateRepeatableTransaction(err) {
            logrus.Warn("Repeat ClaimMany")

            return c.ClaimAccumulated(batch)
        }

        return common.Hash{}, err
    }
    c.Wallet.OnSuccessTransaction()

    logrus.Info("Claimed ", len(batch), " addresses, tx ", tx.Hash().String())

    return tx.Hash(), nil
}

func (c *Claim) ClaimOne(addr string, amount string) (common.Hash, error) {
    tokenAmount, ok := big.NewInt(0).SetString(amount, 0)
    if !ok {
        return common.Hash{}, fmt.Errorf("wrong number provided: %s", amount)
    }

    opts, err := c.Wallet.GetTransactOpts()
    if err != nil {
        c.Wallet.OnFailTransaction(err)
        return common.Hash{}, err
    }

    tx, err := c.Claim.ClaimOne(opts, common.HexToAddress(addr), tokenAmount)
    if err != nil {
        logrus.Error("ClaimOne failed ", err.Error())
        c.Wallet.OnFailTransaction(err)

        if c.Wallet.ValidateRepeatableTransaction(err) {
            logrus.Warn("Repeat ClaimOne")

            return c.ClaimOne(addr, amount)
        }

        return common.Hash{}, err
    }
    c.Wallet.OnSuccessTransaction()

    logrus.Info("Claimed ", addr, ", tx ", tx.Hash().String())

    return tx.Hash(), nil
}