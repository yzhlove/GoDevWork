package endpoint

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"go-kit-five/service"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"time"
)

func LogMiddle(log *zap.Logger) endpoint.Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				log.Debug(fmt.Sprint(ctx.Value(service.ContextReq)),
					zap.String("func-->", "endpoint.middle.log"),
					zap.Int64("times", time.Since(begin).Milliseconds()))
			}(time.Now())
			return e(ctx, request)
		}
	}
}

func GolangRateMiddle(limit *rate.Limiter) endpoint.Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !limit.Allow() {
				return nil, errors.New("rate limit request")
			}
			return e(ctx, request)
		}
	}
}
