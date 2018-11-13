package tariff_group

import (
    "hoqu-geth-api/geth"
    "sync"
    "context"
    "fmt"
    "hoqu-geth-api/models"
    "hoqu-geth-api/sdk/entity"
)

var (
    gs *GethStorage
    gsOnce sync.Once
)

type GethStorage struct {
    *entity.Storage
    HoquPlatform *geth.HoquPlatform
    HoquStorage *geth.HoQuStorage
}

func NewGethStorage(nm string, hp *geth.HoquPlatform, hs *geth.HoQuStorage) *GethStorage {
    return &GethStorage{
        Storage: entity.NewStorage(nm),
        HoquPlatform: hp,
        HoquStorage: hs,
    }
}

func InitGethStorage() (err error) {
    gsOnce.Do(func() {
        err = geth.InitGeth()
        gs = NewGethStorage(
            "ethereum: tariff_group",
            geth.GetHoquPlatform(),
            geth.GetHoQuStorage(),
        )
    })
    return
}

func GetGethStorage() *GethStorage {
    return gs
}

func (s *GethStorage) Op(ctx context.Context, opName string, input interface{}, output interface{}) (err error) {
    span, spanCtx := s.CreateSpan(ctx, opName, input)
    defer s.CloseSpan(span, output, &err)

    switch opName {
    case entity.CREATE:
        err = s.Create(spanCtx, input.(*models.AddTariffGroupRequest), output.(*models.CreateResponseData))
    case entity.SET_STATUS:
        err = s.SetStatus(spanCtx, input.(*models.SetStatusRequest), output.(*models.TxSuccessData))
    case entity.GET_BY_ID:
        err = s.GetById(spanCtx, input.(*models.IdRequest), output.(*models.TariffGroupSuccessData))
    default:
        err = fmt.Errorf("%s: op %s is not supported", s.GetName(), opName)
    }

    return
}

func (s *GethStorage) Create(spanCtx context.Context, input *models.AddTariffGroupRequest, output *models.CreateResponseData) error {
    tx, err := s.HoquPlatform.AddTariffGroup(input)
    if err != nil {
        output.Failed = true
        return err
    }

    output.Tx = tx.String()
    return nil
}

func (s *GethStorage) SetStatus(spanCtx context.Context, input *models.SetStatusRequest, output *models.TxSuccessData) error {
    tx, err := s.HoquPlatform.SetTariffGroupStatus(input)
    if err != nil {
        return err
    }

    output.Tx = tx.String()
    return nil
}

func (s *GethStorage) GetById(spanCtx context.Context, input *models.IdRequest, output *models.TariffGroupSuccessData) error {
    tariffGroup, err := s.HoquStorage.GetTariffGroup(input)
    if err != nil {
        return err
    }

    output.TariffGroup = tariffGroup
    output.Update = true
    s.OpDone(spanCtx)
    return nil
}