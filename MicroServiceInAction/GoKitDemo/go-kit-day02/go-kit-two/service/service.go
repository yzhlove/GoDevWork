package service

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"time"
)

type Service interface {
	TestAdd(ctx context.Context, in Add) AddAck
}

type baseServer struct {
	logger *zap.Logger
}

func NewService(log *zap.Logger) Service {
	var server Service
	server = &baseServer{logger: log}
	server = NewLoggerMiddlewareServer(log)(server)
	return server
}

func (s *baseServer) TestAdd(ctx context.Context, in Add) AddAck {
	time.Sleep(time.Millisecond * 2)
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUID)), zap.Any("call request service", "test add"))
	ack := AddAck{Res: in.A + in.B}
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUID)), zap.Any("call response service", "test add"),
		zap.Any("response function ", "res ack"))
	return ack
}
