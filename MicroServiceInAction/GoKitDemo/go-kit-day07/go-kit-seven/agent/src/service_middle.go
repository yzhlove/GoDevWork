package src

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/metrics"
	"go-kit-seven/agent/pb"
	"go.uber.org/zap"
	"time"
)

const CONTEXT_UID = "request_uid"

type MidFunc func(s Service) Service

type logMid struct {
	log  *zap.Logger
	next Service
}

func NewLogMid(log *zap.Logger) MidFunc {
	return func(s Service) Service {
		return &logMid{log: log, next: s}
	}
}

func (l logMid) Login(ctx context.Context, in *pb.Login) (out *pb.LoginAck, err error) {
	defer func(start time.Time) {
		l.log.Debug(fmt.Sprint(ctx.Value(CONTEXT_UID)), zap.String("<func>", "mid.log.Login"),
			zap.Any("in", in), zap.Any("out", out), zap.Error(err))
	}(time.Now())
	out, err = l.next.Login(ctx, in)
	return
}

type metricsMid struct {
	count     metrics.Counter
	histogram metrics.Histogram
	next      Service
}

func NewMetricsMid(count metrics.Counter, histogram metrics.Histogram) MidFunc {
	return func(s Service) Service {
		return &metricsMid{count: count, histogram: histogram, next: s}
	}
}

func (m metricsMid) Login(ctx context.Context, in *pb.Login) (out *pb.LoginAck, err error) {
	defer func(start time.Time) {
		tags := []string{"method", "login"}
		m.count.With(tags...).Add(1)
		m.histogram.With(tags...).Observe(time.Since(start).Seconds())
	}(time.Now())
	out, err = m.next.Login(ctx, in)
	return
}
