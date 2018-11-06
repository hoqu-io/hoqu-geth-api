package tracing

import (
	"io"
	"net/http"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"time"
	"net/http/httputil"
	"github.com/gin-gonic/gin"
	"fmt"
	"context"
    "strings"
)

var (
	tracer opentracing.Tracer
	closer io.Closer
)

func Init() (err error) {
	if viper.GetBool("tracer.enabled") {
		tracer, closer, err = initJaeger(viper.GetString("project.name"))
		opentracing.SetGlobalTracer(tracer)
	}
	return err
}

func Stop() error {
	if viper.GetBool("tracer.enabled") {
		return closer.Close()
	}
	return nil
}

func initJaeger(service string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			BufferFlushInterval: time.Millisecond * 300,
			LogSpans:            true,
			LocalAgentHostPort:  viper.GetString("tracer.jaeger.host"),
		},
	}
	return cfg.NewTracer(config.Logger(jaeger.StdLogger))
}

func DumpResponse(res *http.Response) (string) {
	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		return err.Error()
	}

	if len(dump) == 0 {
		dump = []byte("Empty response")
	}

	return string(dump)
}

func DumpRequest(req *http.Request) (string) {
	dump, err := httputil.DumpRequest(req, true)
	if err != nil {
		return err.Error()
	}

	if len(dump) == 0 {
		dump = []byte("Empty request")
	}

	return string(dump)
}

func GetServerSpan(ctx *gin.Context) context.Context {
	v, has := ctx.Get("serverSpan")
	if !has {
		return context.Background()
	}
	return v.(context.Context)
}

func SpanName(ctx *gin.Context) string {
    parts := strings.Split(ctx.HandlerName(), "/")
	return fmt.Sprintf("%s", parts[len(parts) - 1])
}