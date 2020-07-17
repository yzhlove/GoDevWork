package endpoint

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"go-kit-three/service"
	"go-kit-three/utils"
	"go.uber.org/zap"
	"time"
)

func LoggerDecorator(logger *zap.Logger) endpoint.Middleware {

	decorator := func(ept endpoint.Endpoint) endpoint.Endpoint {
		point := func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			defer func(begin time.Time) {
				logger.Debug(fmt.Sprint(service.ContextReqUUID),
					zap.Any("func", "loginDecorator"),
					zap.Int64("times", time.Since(begin).Milliseconds()))
			}(time.Now())
			return ept(ctx, req)
		}
		return point
	}
	return decorator
}

func AuthDecorator(logger *zap.Logger) endpoint.Middleware {

	decorator := func(ept endpoint.Endpoint) endpoint.Endpoint {
		point := func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			var tokenStr string
			if tokenStr = fmt.Sprint(ctx.Value(utils.JwtContextKey)); tokenStr == "" {
				err = errors.New("please wait login")
				logger.Debug(fmt.Sprint(ctx.Value(service.ContextReqUUID)),
					zap.Any("[AuthLogin]", "tokenStr is empty"),
					zap.Error(err))
				return
			}
			token, err := utils.ParseToken(tokenStr)
			if err != nil {
				logger.Debug(fmt.Sprint(ctx.Value(service.ContextReqUUID)),
					zap.Any("[AuthLogin]", "parse token err"),
					zap.Error(err))
				return
			}
			if name := token.Name; name != "" {
				ctx = context.WithValue(ctx, "name", name)
			}
			return ept(ctx, req)
		}
		return point
	}
	return decorator
}
