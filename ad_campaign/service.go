package ad_campaign

import (
    "hoqu-geth-api/sdk/entity"
    "sync"
    "context"
    "hoqu-geth-api/models"
    "github.com/satori/go.uuid"
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
    req := input.(*models.AddAdCampaignRequest)
    resp := output.(*models.CreateAdCampaignResponseData)
    id, err := uuid.NewV2('A')
    if err != nil {
        return err
    }

    req.Id = id
    resp.Id = id.String()

    return s.Service.Create(ctx, req, output)
}