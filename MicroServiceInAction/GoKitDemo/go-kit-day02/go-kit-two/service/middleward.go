package service

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

const ContextReqUUID = "req_uuid"

type NewMiddlewareServer func(service Service) Service

type loggerMiddlewareServer struct {
	logger *zap.Logger
	next   Service
}

func NewLoggerMiddlewareServer(logger *zap.Logger) NewMiddlewareServer {
	return func(service Service) Service {
		return loggerMiddlewareServer{
			logger: logger,
			next:   service,
		}
	}
}

func (l loggerMiddlewareServer) TestAdd(ctx context.Context, in Add) (out AddAck) {
	defer func() {
		l.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUID)),
			zap.Any("call service", "middleware test add"),
			zap.Any("call request", in),
			zap.Any("call response ", out))
	}()
	l.logger.Info("service middle ", zap.Any("in", in))
	out = l.next.TestAdd(ctx, in)
	return out
}
