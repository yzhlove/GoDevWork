package service

import (
	"context"
	"fmt"
	"go-kit-five/pb"
	"go.uber.org/zap"
)

const ContextReq = "req_uuid"

type MiddleFunc func(s Service) Service

type logMiddle struct {
	logger *zap.Logger
	next   Service
}

func NewLogMiddle(log *zap.Logger) MiddleFunc {
	return func(s Service) Service {
		return &logMiddle{logger: log, next: s}
	}
}

func (l logMiddle) Login(ctx context.Context, in *pb.Login) (ack *pb.LoginAck, err error) {
	defer func() {
		l.logger.Debug(fmt.Sprint(ctx.Value(ContextReq)),
			zap.String("func-->", "service.middle.LoginEndpoint"),
			zap.Any("in", in),
			zap.Any("ack", ack),
			zap.Error(err))
	}()
	ack, err = l.next.Login(ctx, in)
	return
}
