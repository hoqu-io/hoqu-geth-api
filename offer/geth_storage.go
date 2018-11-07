package offer

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
            "ethereum: offer",
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
        err = s.Create(spanCtx, input.(*models.AddOfferRequest), output.(*models.CreateResponseData))
    case entity.SET_STATUS:
        err = s.SetStatus(spanCtx, input.(*models.SetStatusRequest), output.(*models.TxSuccessData))
    case entity.SET_NAME:
        err = s.SetName(spanCtx, input.(*models.SetNameRequest), output.(*models.TxSuccessData))
    case SET_PAYER_ADDRESS:
        err = s.SetPayerAddress(spanCtx, input.(*models.SetOfferPayerAddressRequest), output.(*models.TxSuccessData))
    case entity.ADD_CHILD_BY_ID:
        err = s.AddOfferTariffGroup(spanCtx, input.(*models.IdWithChildIdRequest), output.(*models.TxSuccessData))
    case entity.GET_BY_ID:
        err = s.GetById(spanCtx, input.(*models.IdRequest), output.(*models.OfferSuccessData))
    default:
        err = fmt.Errorf("%s: op %s is not supported", s.GetName(), opName)
    }

    return
}

func (s *GethStorage) Create(spanCtx context.Context, input *models.AddOfferRequest, output *models.CreateResponseData) error {
    tx, err := s.HoquPlatform.AddOffer(input)
    if err != nil {
        output.Failed = true
        return err
    }

    output.Tx = tx.String()
    return nil
}

func (s *GethStorage) SetStatus(spanCtx context.Context, input *models.SetStatusRequest, output *models.TxSuccessData) error {
    tx, err := s.HoquPlatform.SetOfferStatus(input)
    if err != nil {
        return err
    }

    output.Tx = tx.String()
    return nil
}

func (s *GethStorage) SetName(spanCtx context.Context, input *models.SetNameRequest, output *models.TxSuccessData) error {
    tx, err := s.HoquPlatform.SetOfferName(input)
    if err != nil {
        return err
    }

    output.Tx = tx.String()
    return nil
}

func (s *GethStorage) SetPayerAddress(spanCtx context.Context, input *models.SetOfferPayerAddressRequest, output *models.TxSuccessData) error {
    tx, err := s.HoquPlatform.SetOfferPayerAddress(input)
    if err != nil {
        return err
    }

    output.Tx = tx.String()
    return nil
}

func (s *GethStorage) AddOfferTariffGroup(spanCtx context.Context, input *models.IdWithChildIdRequest, output *models.TxSuccessData) error {
    tx, err := s.HoquPlatform.AddOfferTariffGroup(input)
    if err != nil {
        return err
    }

    output.Tx = tx.String()
    return nil
}

func (s *GethStorage) GetById(spanCtx context.Context, input *models.IdRequest, output *models.OfferSuccessData) error {
    offer, err := s.HoquStorage.GetOffer(input)
    if err != nil {
        return err
    }

    output.Offer = offer
    output.Update = true
    s.OpDone(spanCtx)
    return nil
}