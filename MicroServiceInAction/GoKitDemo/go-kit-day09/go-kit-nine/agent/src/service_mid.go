package src

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/metrics"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go-kit-nine/agent/pb"
	"go-kit-nine/utils"
	"go.uber.org/zap"
	"time"
)

type MiddleFunc func(service Service) Service

type LoggerMiddle struct {
	logger *zap.Logger
	next   Service
}

func NewLoggerMiddle(log *zap.Logger) MiddleFunc {
	return func(service Service) Service {
		return &LoggerMiddle{logger: log, next: service}
	}
}

func (m LoggerMiddle) Login(ctx context.Context, in *pb.UserLogic_Login) (out *pb.UserLogic_LoginAck, err error) {
	defer func(start time.Time) {
		m.logger.Debug(fmt.Sprint(ctx.Value(utils.CONTEXT_UID)),
			zap.String("func", "service.middle.login"),
			zap.Any("in", in),
			zap.Any("out", out),
			zap.Error(err))
	}(time.Now())
	out, err = m.next.Login(ctx, in)
	return
}

type HystrixMiddle struct {
	c    metrics.Counter
	h    metrics.Histogram
	next Service
}

func NewHystrixMiddle(c metrics.Counter, h metrics.Histogram) MiddleFunc {
	return func(service Service) Service {
		return &HystrixMiddle{c: c, h: h, next: service}
	}
}

func (m HystrixMiddle) Login(ctx context.Context, in *pb.UserLogic_Login) (*pb.UserLogic_LoginAck, error) {
	defer func(start time.Time) {
		methods := []string{"method", "login"}
		m.c.With(methods...).Add(1)
		m.h.With(methods...).Observe(time.Since(start).Seconds())
	}(time.Now())
	return m.next.Login(ctx, in)
}

type TracerMiddle struct {
	next   Service
	tracer opentracing.Tracer
}

func NewTracerMiddle(tracer opentracing.Tracer) MiddleFunc {
	return func(service Service) Service {
		return &TracerMiddle{tracer: tracer, next: service}
	}
}

func (m TracerMiddle) Login(ctx context.Context, in *pb.UserLogic_Login) (*pb.UserLogic_LoginAck, error) {
	span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, m.tracer, "service", opentracing.Tag{
		Key:   string(ext.Component),
		Value: "NewTracerServiceMiddle",
	})
	defer func() {
		span.LogKV("account", in.Account, "password", in.Password)
		span.Finish()
	}()
	return m.next.Login(ctx, in)
}
