package endpoint

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"go-kit-four/service"
	"go-kit-four/utils"
	"go.uber.org/ratelimit"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"strings"
	"time"
)

func LoginMiddle(log *zap.Logger) endpoint.Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				log.Debug(fmt.Sprint(ctx.Value(service.ContextUID)),
					zap.String("func-->", "endpoint.middle.Login"),
					zap.Any("times", time.Since(begin).Milliseconds()))
			}(time.Now())
			return e(ctx, request)
		}
	}
}

func AuthMiddle(log *zap.Logger) endpoint.Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			token := fmt.Sprint(ctx.Value(utils.JwtContextKey))
			if token = strings.TrimSpace(token); len(token) == 0 {
				err = errors.New("please login")
				log.Debug(fmt.Sprint(ctx.Value(service.ContextUID)),
					zap.String("func-->", "endpoint.middle.Auth"),
					zap.Error(err))
				return
			}
			jwtInfo, err := utils.ParseToken(token)
			if err != nil {
				log.Debug(fmt.Sprint(ctx.Value(service.ContextUID)),
					zap.String("func-->", "endpoint.middle.Auth"),
					zap.String("err reason", "parse token err"),
					zap.Error(err))
				return
			}
			if name := jwtInfo.Name; name != "" {
				ctx = context.WithValue(ctx, "name", name)
			}
			return e(ctx, request)
		}
	}
}

func GolangRateWaitMiddle(limit *rate.Limiter) endpoint.Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if err = limit.Wait(ctx); err != nil {
				return
			}
			return e(ctx, request)
		}
	}
}

func GolangRateAllowMiddle(limit *rate.Limiter) endpoint.Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !limit.Allow() {
				return nil, errors.New("limit req allow")
			}
			return e(ctx, request)
		}
	}
}

func UberRateMiddle(limit ratelimit.Limiter) endpoint.Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			limit.Take()
			return e(ctx, request)
		}
	}
}
