package lead

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
            "db: lead",
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

    conn.AutoMigrate(&models.Lead{})

    return nil
}

func (s *DbStorage) Op(ctx context.Context, opName string, input interface{}, output interface{}) (err error) {
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
        err = nil
    case entity.GET_BY_ID:
        err = s.GetById(spanCtx, input.(*models.TransactLeadRequest), output.(*models.LeadSuccessData))
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
        err = s.AfterCreate(spanCtx, input.(*models.AddLeadRequest), output.(*models.CreateResponseData))
    case TRANSACT:
        err = s.AfterTransact(spanCtx, input.(*models.TransactLeadRequest), output.(*models.TxSuccessData))
    case entity.GET_BY_ID:
        err = s.AfterGetById(spanCtx, input.(*models.TransactLeadRequest), output.(*models.LeadSuccessData))
    }
}

func (s *DbStorage) Create(spanCtx context.Context, input *models.AddLeadRequest, output *models.CreateResponseData) error {
    conn, err := s.db.NewConnection()
    if err != nil {
        return err
    }
    defer conn.Close()

    m := models.Lead{
        ID: input.Id.String(),
        CreatedAt: time.Now(),
        AdId: input.AdId,
        TrackerId: input.TrackerId,
        Meta: input.Meta,
        DataUrl: input.DataUrl,
        Price: input.Price,
        Intermediaries: make(map[string]float32),
        Status: models.STATUS_CREATED,
    }
    if err = conn.Create(&m).Error; err != nil {
        return err
    }

    return nil
}

func (s *DbStorage) AfterCreate(spanCtx context.Context, input *models.AddLeadRequest, output *models.CreateResponseData) error {
    if !output.Failed {
        return nil
    }

    conn, err := s.db.NewConnection()
    if err != nil {
        return err
    }
    defer conn.Close()

    m := models.Lead{
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

func (s *DbStorage) SetStatus(spanCtx context.Context, input *models.SetLeadStatusRequest, output *models.TxSuccessData) error {
    conn, err := s.db.NewConnection()
    if err != nil {
        return err
    }
    defer conn.Close()

    m := models.Lead{
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

func (s *DbStorage) SetPrice(spanCtx context.Context, input *models.SetLeadPriceRequest, output *models.TxSuccessData) error {
    conn, err := s.db.NewConnection()
    if err != nil {
        return err
    }
    defer conn.Close()

    m := models.Lead{
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

    m.Price = input.Price
    conn.Save(&m)
    return nil
}

func (s *DbStorage) SetDataUrl(spanCtx context.Context, input *models.SetLeadDataUrlRequest, output *models.TxSuccessData) error {
    conn, err := s.db.NewConnection()
    if err != nil {
        return err
    }
    defer conn.Close()

    m := models.Lead{
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

    m.DataUrl = input.DataUrl
    conn.Save(&m)
    return nil
}

func (s *DbStorage) AddIntermediary(spanCtx context.Context, input *models.AddLeadIntermediaryRequest, output *models.TxSuccessData) error {
    conn, err := s.db.NewConnection()
    if err != nil {
        return err
    }
    defer conn.Close()

    m := models.Lead{
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

    m.Intermediaries[input.Address] = input.Percent
    conn.Save(&m)
    return nil
}

func (s *DbStorage) AfterTransact(spanCtx context.Context, input *models.TransactLeadRequest, output *models.TxSuccessData) error {
    if output.Tx == "" {
        return nil
    }

    conn, err := s.db.NewConnection()
    if err != nil {
        return err
    }
    defer conn.Close()

    m := models.Lead{
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

    m.Status = models.STATUS_DONE
    conn.Save(&m)
    return nil
}

func (s *DbStorage) GetById(spanCtx context.Context, input *models.TransactLeadRequest, output *models.LeadSuccessData) error {
    conn, err := s.db.NewConnection()
    if err != nil {
        return err
    }
    defer conn.Close()

    m := models.Lead{
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

    output.Lead = &m
    s.OpDone(spanCtx)
    return nil
}

func (s *DbStorage) AfterGetById(spanCtx context.Context, input *models.TransactLeadRequest, output *models.LeadSuccessData) error {
    if !output.Update {
        return nil
    }

    conn, err := s.db.NewConnection()
    if err != nil {
        return err
    }
    defer conn.Close()

    m := *output.Lead
    if err = conn.Create(&m).Error; err != nil {
        return err
    }

    return nil
}