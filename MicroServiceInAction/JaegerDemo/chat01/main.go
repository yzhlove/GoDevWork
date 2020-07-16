package main

///////////////////////////////////////////////
// Jaeger
///////////////////////////////////////////////

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	jaegerCfg "github.com/uber/jaeger-client-go/config"
	"io"
	"net/http"
	"time"
)

func NewJaegerTracer(serviceName, jaegerHostPort string) (opentracing.Tracer, io.Closer) {

	cfg := &jaegerCfg.Configuration{
		Sampler: &jaegerCfg.SamplerConfig{
			Type:  "const", //固定采样
			Param: 1,       //全采样
		},
		Reporter: &jaegerCfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: jaegerHostPort,
		}, ServiceName: serviceName,
	}
	if tracer, closer, err := cfg.NewTracer(jaegerCfg.Logger(jaeger.StdLogger)); err == nil {
		opentracing.SetGlobalTracer(tracer)
		return tracer, closer
	} else {
		panic(fmt.Sprintf("ERROR:cannect init jaeger:%v \n", err))
	}
}

func SetUp() gin.HandlerFunc {
	return func(context *gin.Context) {
		var parentSpan opentracing.Span
		tracer, closer := NewJaegerTracer("gin_test_tracer", "127.0.0.1:6831")
		defer closer.Close()
		spCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(context.Request.Header))
		if err != nil {
			parentSpan = tracer.StartSpan(context.Request.URL.Path)
			defer parentSpan.Finish()
		} else {
			parentSpan = opentracing.StartSpan(
				context.Request.URL.Path,
				opentracing.ChildOf(spCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
			)
			defer parentSpan.Finish()
		}
		context.Set("Tracer", tracer)
		context.Set("ParentSpanContext", parentSpan.Context())
		context.Next()
	}
}

func SetRouter(engine *gin.Engine) {
	engine.Use(SetUp())
	engine.GET("/jaeger", JaegerTest)
}

func JaegerTest(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "succeed",
		"mode":    "jaeger",
	})
}

func main() {

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	SetRouter(engine)
	server := &http.Server{
		Addr:              ":1234",
		Handler:           engine,
		ReadHeaderTimeout: 20 * time.Second,
		WriteTimeout:      20 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
