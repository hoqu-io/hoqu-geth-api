package offer

import (
    "sync"
    "context"
    "fmt"
    "hoqu-geth-api/models"
    "hoqu-geth-api/sdk/entity"
    "hoqu-geth-api/sdk/db"
    "github.com/spf13/viper"
    "time"
    "github.com/opentracing/opentracing-go"
    "github.com/opentracing/opentracing-go/log"
)

var (
    ds *DbStorage
    dsOnce sync.Once
)

type DbStorage struct {
    *entity.Storage
    db *db.ConnectionManager
}

func NewDbStorage(nm, driver, dsn string) *DbStorage {
    return &DbStorage{
        Storage: entity.NewStorage(nm),
        db: db.NewConnectionManager(driver, dsn),
    }
}

func InitDbStorage() (err error) {
    dsOnce.Do(func() {
        ds = NewDbStorage(
            "db: offer",
            viper.GetString("db.driver"),
            viper.GetString("db.dsn"),
        )

        err = ds.AutoMigrate()
    })
    return
}

func GetDbStorage() *DbStorage {
    return ds
}

func (s *DbStorage) AutoMigrate() error {
    conn, err := s.db.NewConnection()
    if err != nil {
        return err
    }
    defer conn.Close()

    conn.AutoMigrate(&models.Offer{})

    return nil
}

func (s *DbStorage) Op(ctx context.Context, opName string, input interface{}, output interface{}) (err error) {
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

func (s *DbStorage) AfterOp(ctx context.Context, opName string, input interface{}, output interface{}) {
    var err error
    span, spanCtx := s.CreateSpan(ctx, "after " + opName, input)
    defer s.CloseSpan(span, output, &err)

    switch opName {
    case entity.CREATE:
        err = s.AfterCreate(spanCtx, input.(*models.AddOfferRequest), output.(*models.CreateResponseData))
    case entity.GET_BY_ID:
        err = s.AfterGetById(spanCtx, input.(*models.IdRequest), output.(*models.OfferSuccessData))
    }
}

func (s *DbStorage) Create(spanCtx context.Context, input *models.AddOfferRequest, output *models.CreateResponseData) error {
    conn, err := s.db.NewConnection()
    if err != nil {
        return err
    }
    defer conn.Close()

    m := models.Offer{
        ID: input.Id.String(),
        CreatedAt: time.Now(),
        OwnerId: input.OwnerId,
        NetworkId: input.NetworkId,
        MerchantId: input.MerchantId,
        PayerAddress: input.PayerAddress,
        Name: input.Name,
        DataUrl: input.DataUrl,
        TariffGroups: make([]string,0),
        Status: models.STATUS_CREATED,
    }
    if err = conn.Create(&m).Error; err != nil {
        return err
    }

    return nil
}

func (s *DbStorage) AfterCreate(spanCtx context.Context, input *models.AddOfferRequest, output *models.CreateResponseData) error {
    if !output.Failed {
        return nil
    }

    conn, err := s.db.NewConnection()
    if err != nil {
        return err
    }
    defer conn.Close()

    m := models.Offer{
        ID: input.Id.String(),
    }
    res := conn.Delete(&m)

    if res.RecordNotFound() {
        opentracing.SpanFromContext(spanCtx).LogFields(
            log.String("message", res.Error.Error()),
        )
        return nil
    }

    if res.Error != nil {
        return res.Error
    }

    opentracing.SpanFromContext(spanCtx).LogFields(
        log.String("message", "record deleted from DB"),
    )
    return nil
}

func (s *DbStorage) SetStatus(spanCtx context.Context, input *models.SetStatusRequest, output *models.TxSuccessData) error {
    conn, err := s.db.NewConnection()
    if err != nil {
        return err
    }
    defer conn.Close()

    m := models.Offer{
        ID: input.Id,
    }
    res := conn.First(&m)

    if res.RecordNotFound() {
        opentracing.SpanFromContext(spanCtx).LogFields(
            log.String("message", res.Error.Error()),
        )
        return nil
    }

    if res.Error != nil {
        return res.Error
    }

    m.Status = input.Status
    conn.Save(&m)
    return nil
}

func (s *DbStorage) SetName(spanCtx context.Context, input *models.SetNameRequest, output *models.TxSuccessData) error {
    conn, err := s.db.NewConnection()
    if err != nil {
        return err
    }
    defer conn.Close()

    m := models.Offer{
        ID: input.Id,
    }
    res := conn.First(&m)

    if res.RecordNotFound() {
        opentracing.SpanFromContext(spanCtx).LogFields(
            log.String("message", res.Error.Error()),
        )
        return nil
    }

    if res.Error != nil {
        return res.Error
    }

    m.Name = input.Name
    conn.Save(&m)
    return nil
}

func (s *DbStorage) SetPayerAddress(spanCtx context.Context, input *models.SetOfferPayerAddressRequest, output *models.TxSuccessData) error {
    conn, err := s.db.NewConnection()
    if err != nil {
        return err
    }
    defer conn.Close()

    m := models.Offer{
        ID: input.Id,
    }
    res := conn.First(&m)

    if res.RecordNotFound() {
        opentracing.SpanFromContext(spanCtx).LogFields(
            log.String("message", res.Error.Error()),
        )
        return nil
    }

    if res.Error != nil {
        return res.Error
    }

    m.PayerAddress = input.PayerAddress
    conn.Save(&m)
    return nil
}

func (s *DbStorage) AddOfferTariffGroup(spanCtx context.Context, input *models.IdWithChildIdRequest, output *models.TxSuccessData) error {
    conn, err := s.db.NewConnection()
    if err != nil {
        return err
    }
    defer conn.Close()

    m := models.Offer{
        ID: input.Id,
    }
    res := conn.First(&m)

    if res.RecordNotFound() {
        opentracing.SpanFromContext(spanCtx).LogFields(
            log.String("message", res.Error.Error()),
        )
        return nil
    }

    if res.Error != nil {
        return res.Error
    }

    m.TariffGroups = append(m.TariffGroups, input.ChildId)
    conn.Save(&m)
    return nil
}

func (s *DbStorage) GetById(spanCtx context.Context, input *models.IdRequest, output *models.OfferSuccessData) error {
    conn, err := s.db.NewConnection()
    if err != nil {
        return err
    }
    defer conn.Close()

    m := models.Offer{
        ID: input.Id,
    }
    res := conn.First(&m)

    if res.RecordNotFound() {
        opentracing.SpanFromContext(spanCtx).LogFields(
            log.String("message", res.Error.Error()),
        )
        return nil
    }

    if res.Error != nil {
        return res.Error
    }

    output.Offer = &m
    s.OpDone(spanCtx)
    return nil
}

func (s *DbStorage) AfterGetById(spanCtx context.Context, input *models.IdRequest, output *models.OfferSuccessData) error {
    if !output.Update {
        return nil
    }

    conn, err := s.db.NewConnection()
    if err != nil {
        return err
    }
    defer conn.Close()

    m := *output.Offer
    if err = conn.Create(&m).Error; err != nil {
        return err
    }

    return nil
}