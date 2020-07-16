package endpoint

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"go-kit-two/service"
	"go-kit-two/utils"
	"go.uber.org/zap"
	"reflect"
	"time"
)

func LoggerMiddleware(logger *zap.Logger) endpoint.Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				logger.Debug(fmt.Sprint(ctx.Value(service.ContextReqUUID)),
					zap.Any("loggerMiddleware", "succeed"),
					zap.Any("timeout", time.Since(begin).Milliseconds()))
			}(time.Now())
			utils.GetLog().Debug("LoggerMiddleware", zap.String("requestType", reflect.TypeOf(request).String()))
			return e(ctx, request)
		}
	}
}
