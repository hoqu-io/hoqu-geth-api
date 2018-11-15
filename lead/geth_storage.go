package lead

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
            "ethereum: lead",
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
        err = s.Create(spanCtx, input.(*models.AddLeadRequest), output.(*models.CreateResponseData))
    case entity.SET_STATUS:
        err = s.SetStatus(spanCtx, input.(*models.SetLeadStatusRequest), output.(*models.TxSuccessData))
    case SET_PRICE:
        err = s.SetPrice(spanCtx, input.(*models.SetLeadPriceRequest), output.(*models.TxSuccessData))
    case entity.SET_URL:
        err = s.SetDataUrl(spanCtx, input.(*models.SetLeadDataUrlRequest), output.(*models.TxSuccessData))
    case ADD_INTERMEDIARY:
        err = s.AddIntermediary(spanCtx, input.(*models.AddLeadIntermediaryRequest), output.(*models.TxSuccessData))
    case TRANSACT:
        err = s.Transact(spanCtx, input.(*models.TransactLeadRequest), output.(*models.TxSuccessData))
    case entity.GET_BY_ID:
        err = s.GetById(spanCtx, input.(*models.TransactLeadRequest), output.(*models.LeadSuccessData))
    default:
        err = fmt.Errorf("%s: op %s is not supported", s.GetName(), opName)
    }

    return
}

func (s *GethStorage) Create(spanCtx context.Context, input *models.AddLeadRequest, output *models.CreateResponseData) error {
    tx, err := s.HoquPlatform.AddLead(input)
    if err != nil {
        output.Failed = true
        return err
    }

    output.Tx = tx.String()
    return nil
}

func (s *GethStorage) SetStatus(spanCtx context.Context, input *models.SetLeadStatusRequest, output *models.TxSuccessData) error {
    tx, err := s.HoquPlatform.SetLeadStatus(input)
    if err != nil {
        return err
    }

    output.Tx = tx.String()
    return nil
}

func (s *GethStorage) SetPrice(spanCtx context.Context, input *models.SetLeadPriceRequest, output *models.TxSuccessData) error {
    tx, err := s.HoquPlatform.SetLeadPrice(input)
    if err != nil {
        return err
    }

    output.Tx = tx.String()
    return nil
}

func (s *GethStorage) SetDataUrl(spanCtx context.Context, input *models.SetLeadDataUrlRequest, output *models.TxSuccessData) error {
    tx, err := s.HoquPlatform.SetLeadDataUrl(input)
    if err != nil {
        return err
    }

    output.Tx = tx.String()
    return nil
}

func (s *GethStorage) AddIntermediary(spanCtx context.Context, input *models.AddLeadIntermediaryRequest, output *models.TxSuccessData) error {
    tx, err := s.HoquPlatform.AddLeadIntermediary(input)
    if err != nil {
        return err
    }

    output.Tx = tx.String()
    return nil
}

func (s *GethStorage) Transact(spanCtx context.Context, input *models.TransactLeadRequest, output *models.TxSuccessData) error {
    tx, err := s.HoquPlatform.TransactLead(input.Id, input.AdId)
    if err != nil {
        return err
    }

    output.Tx = tx.String()
    return nil
}

func (s *GethStorage) GetById(spanCtx context.Context, input *models.TransactLeadRequest, output *models.LeadSuccessData) error {
    ad, err := s.HoquStorage.GetAdCampaign(input.AdId)
    if err != nil {
        return err
    }

    adContract, err := geth.GetAdCampaign(ad.ContractAddress)
    if err != nil {
        return err
    }

    lead, err := adContract.GetLead(input.Id)
    if err != nil {
        return err
    }

    output.Lead = lead
    output.Update = true
    s.OpDone(spanCtx)
    return nil
}