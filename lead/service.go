package lead

import (
    "hoqu-geth-api/sdk/entity"
    "sync"
    "context"
    "hoqu-geth-api/models"
    "github.com/satori/go.uuid"
    "strings"
)

const (
    ADD_INTERMEDIARY = "add_intermediary"
    TRANSACT = "transact"
    SET_PRICE = "set_price"
)

var (
    s *Service
    sOnce sync.Once
)

type Service struct {
    *entity.Service
}

func NewService(storage entity.StorageInterface) *Service {
    return &Service{
        Service: entity.NewService(storage, true),
    }
}

func InitService() (err error) {
    sOnce.Do(func() {
        err = InitDbStorage()
        if err != nil {
            return
        }
        s = NewService(GetDbStorage())

        err = InitGethStorage()
        if err != nil {
            return
        }
        s.AppendStorage(GetGethStorage())
    })
    return
}

func GetService() *Service {
    return s
}

func (s *Service) Create(ctx context.Context, input interface{}, output interface{}) error {
    req := input.(*models.AddLeadRequest)
    resp := output.(*models.CreateResponseData)
    id, err := uuid.NewV2('L')
    if err != nil {
        return err
    }

    req.Id = id
    resp.Id = id.String()
    req.DataUrl = strings.Replace(req.DataUrl, ":id", id.String(), 1)

    return s.Service.Create(ctx, req, output)
}

func (s *Service) AddIntermediary(ctx context.Context, input interface{}, output interface{}) error {
    return s.Op(ctx, ADD_INTERMEDIARY, input, output)
}

func (s *Service) Transact(ctx context.Context, input interface{}, output interface{}) error {
    return s.Op(ctx, TRANSACT, input, output)
}

func (s *Service) SetPrice(ctx context.Context, input interface{}, output interface{}) error {
    return s.Op(ctx, SET_PRICE, input, output)
}
