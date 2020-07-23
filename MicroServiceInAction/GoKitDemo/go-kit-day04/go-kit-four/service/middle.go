package service

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

const ContextUID = "request_uuid"

type MiddleFunc func(service Service) Service

type logMiddle struct {
	logger *zap.Logger
	next   Service
}

func NewLogMiddle(log *zap.Logger) MiddleFunc {
	return func(service Service) Service {
		return &logMiddle{
			logger: log,
			next:   service,
		}
	}
}

func (l logMiddle) TestAdd(ctx context.Context, in Add) (out AddAck) {
	defer func() {
		l.logger.Debug(fmt.Sprint(ctx.Value(ContextUID)),
			zap.Any("func-->", "logMiddle.TestAdd"),
			zap.Any("req", in),
			zap.Any("ack", out))
	}()
	out = l.next.TestAdd(ctx, in)
	return out
}

func (l logMiddle) Login(ctx context.Context, in Login) (out LoginAck, err error) {
	defer func() {
		l.logger.Debug(fmt.Sprint(ctx.Value(ContextUID)),
			zap.Any("func-->", "logMiddle.LoginEndpoint"),
			zap.Any("req", in),
			zap.Any("ack", out),
			zap.Error(err),
		)
	}()
	out, err = l.next.Login(ctx, in)
	return
}
