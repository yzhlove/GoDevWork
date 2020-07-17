package service

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

const ContextReqUUID = "req_uuid"

type DecoratorService func(service Service) Service

type logService struct {
	logger *zap.Logger
	next   Service
}

func NewLoggerServer(log *zap.Logger) DecoratorService {
	decorator := func(service Service) Service {
		l := &logService{logger: log, next: service}
		return l
	}
	return decorator
}

func (l logService) TestAdd(ctx context.Context, in Add) (out AddAck) {
	defer func() {
		l.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUID)),
			zap.Any("func", "logService->testAdd"),
			zap.Any("req", in),
			zap.Any("resp", out))
	}()

	out = l.next.TestAdd(ctx, in)
	return
}

func (l logService) Login(ctx context.Context, in Login) (out LoginAck, err error) {
	defer func() {
		l.logger.Debug(fmt.Sprint(ContextReqUUID),
			zap.Any("func", "logService->login"),
			zap.Any("req", in),
			zap.Any("resp", out),
			zap.Error(err))
	}()
	out, err = l.next.Login(ctx, in)
	return
}
