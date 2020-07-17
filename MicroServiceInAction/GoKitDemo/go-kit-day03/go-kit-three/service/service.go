package service

import (
	"context"
	"fmt"
	"go-kit-three/utils"
	"go.uber.org/zap"
	"time"
)

type Service interface {
	TestAdd(ctx context.Context, in Add) AddAck
	Login(ctx context.Context, in Login) (LoginAck, error)
}

type baseService struct {
	logger *zap.Logger
}

func NewService(log *zap.Logger) Service {
	var s Service
	s = &baseService{logger: log}
	s = NewLoggerServer(log)(s)
	return s
}

func (s baseService) TestAdd(ctx context.Context, in Add) AddAck {
	time.Sleep(time.Millisecond * 2)
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUID)),
		zap.Any("func", "baseService->testAdd"),
		zap.Any("req", in),
		zap.Any("name", ctx.Value("name")))
	ack := AddAck{Res: in.A + in.B}
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUID)),
		zap.Any("func", "baseService->testAdd"),
		zap.Any("resp", ack))
	return ack
}

func (s baseService) Login(ctx context.Context, in Login) (ack LoginAck, err error) {
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUID)),
		zap.Any("func", "baseService->login"),
		zap.Any("req", in))
	ack.Token, err = utils.GenericToken(in.Account, 1)
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReqUUID)),
		zap.Any("func", "baseService->login"),
		zap.Any("resp", ack))
	return
}
