package src

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/metrics"
	"go-kit-eight/agent/pb"
	"go.uber.org/zap"
	"time"
)

const CONTEXT_UID = "context_uid"

type MidFunc func(s Service) Service

type logMid struct {
	logger *zap.Logger
	next   Service
}

func NewLogMid(log *zap.Logger) MidFunc {
	return func(s Service) Service {
		return &logMid{logger: log, next: s}
	}
}

func (m logMid) Login(ctx context.Context, in *pb.Login) (out *pb.LoginAck, err error) {
	defer func(start time.Time) {
		m.logger.Debug(
			fmt.Sprint(ctx.Value(CONTEXT_UID)),
			zap.String("func", "mid.log.Login"),
			zap.Any("in", in),
			zap.Any("out", out),
			zap.Error(err),
		)
	}(time.Now())
	out, err = m.next.Login(ctx, in)
	return
}

type metricsMid struct {
	next Service
	c    metrics.Counter
	h    metrics.Histogram
}

func NewMetricsMid(c metrics.Counter, h metrics.Histogram) MidFunc {
	return func(s Service) Service {
		return &metricsMid{next: s, c: c, h: h}
	}
}

func (m metricsMid) Login(ctx context.Context, in *pb.Login) (out *pb.LoginAck, err error) {
	defer func(start time.Time) {
		methods := []string{"method", "login"}
		m.c.With(methods...).Add(1)
		m.h.With(methods...).Observe(time.Since(start).Seconds())
	}(time.Now())
	out, err = m.next.Login(ctx, in)
	return
}
