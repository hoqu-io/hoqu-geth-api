package tariff

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
			"db: tariff",
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

	conn.AutoMigrate(&models.Tariff{})

	return nil
}

func (s *DbStorage) Op(ctx context.Context, opName string, input interface{}, output interface{}) (err error) {
	span, spanCtx := s.CreateSpan(ctx, opName, input)
	defer s.CloseSpan(span, output, &err)

	switch opName {
	case entity.CREATE:
		err = s.Create(spanCtx, input.(*models.AddTariffRequest), output.(*models.CreateResponseData))
	case entity.SET_STATUS:
		err = s.SetStatus(spanCtx, input.(*models.SetStatusRequest), output.(*models.TxSuccessData))
	case entity.GET_BY_ID:
		err = s.GetById(spanCtx, input.(*models.IdRequest), output.(*models.TariffSuccessData))
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
		err = s.AfterCreate(spanCtx, input.(*models.AddTariffRequest), output.(*models.CreateResponseData))
	case entity.GET_BY_ID:
		err = s.AfterGetById(spanCtx, input.(*models.IdRequest), output.(*models.TariffSuccessData))
	}
}

func (s *DbStorage) Create(spanCtx context.Context, input *models.AddTariffRequest, output *models.CreateResponseData) error {
	conn, err := s.db.NewConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

    mt := models.TariffGroup{
        ID: input.TariffGroupId,
    }
    res := conn.First(&mt)

    if res.RecordNotFound() {
        opentracing.SpanFromContext(spanCtx).LogFields(
            log.String("message", res.Error.Error()),
        )
        return res.Error
    }

    if res.Error != nil {
        return res.Error
    }

	m := models.Tariff{
		ID: input.Id.String(),
		CreatedAt: time.Now(),
		TariffGroupId: input.TariffGroupId,
		Name: input.Name,
		Action: input.Action,
		CalcMethod: input.CalcMethod,
		Price: input.Price,
		Status: models.STATUS_CREATED,
	}
	if err = conn.Create(&m).Error; err != nil {
		return err
	}

	mt.Tariffs = append(mt.Tariffs, input.Id.String())
	conn.Save(&mt)

	return nil
}

func (s *DbStorage) AfterCreate(spanCtx context.Context, input *models.AddTariffRequest, output *models.CreateResponseData) error {
	if !output.Failed {
		return nil
	}

	conn, err := s.db.NewConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

    mt := models.TariffGroup{
        ID: input.TariffGroupId,
    }
    res := conn.First(&mt)

    if res.RecordNotFound() {
        opentracing.SpanFromContext(spanCtx).LogFields(
            log.String("message", res.Error.Error()),
        )
        return res.Error
    }

    if res.Error != nil {
        return res.Error
    }

    tariffs := make([]string, 0)
    for _, t := range mt.Tariffs {
        if t != input.Id.String() {
            tariffs = append(tariffs, t)
        }
    }
    mt.Tariffs = tariffs
    conn.Save(&mt)

    m := models.Tariff{
		ID: input.Id.String(),
	}
	res = conn.Delete(&m)

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

	m := models.Tariff{
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

func (s *DbStorage) GetById(spanCtx context.Context, input *models.IdRequest, output *models.TariffSuccessData) error {
	conn, err := s.db.NewConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	m := models.Tariff{
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

	output.Tariff = &m
	s.OpDone(spanCtx)
	return nil
}

func (s *DbStorage) AfterGetById(spanCtx context.Context, input *models.IdRequest, output *models.TariffSuccessData) error {
	if !output.Update {
		return nil
	}

	conn, err := s.db.NewConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	m := *output.Tariff
	if err = conn.Create(&m).Error; err != nil {
		return err
	}

	return nil
}