package entity

import (
    "context"
    "fmt"
    "github.com/opentracing/opentracing-go"
    "github.com/opentracing/opentracing-go/log"
    "encoding/json"
    "hoqu-geth-api/sdk/app"
)

type Storage struct {
    Name string
}

func NewStorage(nm string) *Storage {
    return &Storage{
        Name: nm,
    }
}

type StorageInterface interface {
    Op(context.Context, string, interface{}, interface{}) error
    AfterOp(context.Context, string, interface{}, interface{})
    OpDone(context.Context)
    CreateSpan(context.Context, string, interface{}) (opentracing.Span, context.Context)
    CloseSpan(opentracing.Span, interface{}, *error)
    GetName() string
}

func (s *Storage) Op(ctx context.Context, opName string, input interface{}, output interface{}) error {
    return fmt.Errorf("%s: op %s is not implemented", s.GetName(), opName)
}

func (s *Storage) AfterOp(ctx context.Context, opName string, input interface{}, output interface{}) {
    // Do nothing by default
}

func (s *Storage) OpDone(doneCtx context.Context) {
    opDone := doneCtx.Value(OP_DONE).(*bool)
    *opDone = true
}

func (s *Storage) CreateSpan(ctx context.Context, opName string, input interface{}) (span opentracing.Span, spanCtx context.Context) {
    span, spanCtx = opentracing.StartSpanFromContext(ctx, fmt.Sprintf("%s -> %s", s.GetName(), opName))
    span.SetTag("app.env", app.Env())
    span.SetTag("span.kind", "client")
    i, _ := json.Marshal(input)
    span.LogFields(
        log.String("op.input", string(i)),
    )

    return
}

func (s *Storage) CloseSpan(span opentracing.Span, output interface{}, err *error) {
    o, _ := json.Marshal(output)
    span.LogFields(
        log.String("op.output", string(o)),
    )

    if *err != nil {
        span.SetTag("error", true)
        span.LogFields(
            log.Error(*err),
        )
    }

    span.Finish()
}

func (s *Storage) GetName() string {
    return s.Name
}