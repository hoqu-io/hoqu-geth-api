package user

import (
    "hoqu-geth-api/sdk/entity"
    "sync"
    "context"
    "hoqu-geth-api/models"
    "github.com/satori/go.uuid"
)

const (
    ADD_ADDRESS = "add_address"
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
    req := input.(*models.RegisterUserRequest)
    resp := output.(*models.CreateResponseData)
    id, err := uuid.NewV2('U')
    if err != nil {
        return err
    }

    req.Id = id
    resp.Id = id.String()

    return s.Service.Create(ctx, req, output)
}

func (s *Service) AddAddress(ctx context.Context, input interface{}, output interface{}) error {
    return s.Op(ctx, ADD_ADDRESS, input, output)
}