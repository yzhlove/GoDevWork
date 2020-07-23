package service

import (
	"context"
	"errors"
	"fmt"
	"go-kit-four/utils"
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
	return NewLogMiddle(log)(&baseService{logger: log})
}

func (s baseService) TestAdd(ctx context.Context, in Add) AddAck {
	time.Sleep(time.Millisecond * 100)
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextUID)),
		zap.Any("func-->", "baseService.TestAdd"),
		zap.Any("account", ctx.Value("name")))
	ack := AddAck{Res: in.A + in.B}
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextUID)),
		zap.Any("func-->", "baseService.TestAdd"),
		zap.Any("ack", ack))
	return ack
}

func (s baseService) Login(ctx context.Context, in Login) (LoginAck, error) {
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextUID)),
		zap.Any("func-->", "baseService.LoginEndpoint"),
		zap.Any("in", in))
	if in.Account != "yzh" && in.Passwd != "12345678" {
		return LoginAck{}, errors.New("login failed")
	}
	token, err := utils.GenericToken(in.Account, 1)
	if err != nil {
		return LoginAck{}, err
	}
	ack := LoginAck{Token: token}
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextUID)),
		zap.Any("func-->", "baseService.LoginEndpoint"),
		zap.Any("ack", ack))
	return ack, nil
}
