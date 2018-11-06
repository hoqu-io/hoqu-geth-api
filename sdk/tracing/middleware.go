package tracing

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"golang.org/x/net/context"
	"hoqu-geth-api/sdk/app"
)

func Middleware(ctx *gin.Context) {
	done, has := ctx.Get("done")
	if !has {
		return
	}

	reqSpan, _ := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(ctx.Request.Header),
	)
	span := opentracing.GlobalTracer().StartSpan(SpanName(ctx), ext.RPCServerOption(reqSpan))
	span.SetTag("app.env", app.Env())

	reqDump := DumpRequest(ctx.Request)
	span.LogFields(
		log.String("http.incoming.request", reqDump),
	)
	span.SetTag("http.incoming.method", ctx.Request.Method)
	span.SetTag("http.incoming.URI", ctx.Request.URL.String())

	ctx.Set("serverSpan", opentracing.ContextWithSpan(context.Background(), span))

	go tickTracing(ctx, span, done.(chan bool))
}

func tickTracing(ctx *gin.Context, span opentracing.Span, done chan bool) {
	defer func() {
		span.SetTag("http.response.code", ctx.Writer.Status())
		if resp, exists := ctx.Get("responseBody"); exists {
			respBody, _ := json.Marshal(resp)
			span.LogFields(log.String("http.response.body", string(respBody)))
		}

		span.Finish()
	}()

	for {
		select {
		case <-done:
			return
		case <-time.After(300 * time.Second):
			return
		}
	}
}

