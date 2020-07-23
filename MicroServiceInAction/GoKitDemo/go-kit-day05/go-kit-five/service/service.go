package service

import (
	"context"
	"errors"
	"fmt"
	"go-kit-five/pb"
	"go-kit-five/utils"
	"go.uber.org/zap"
)

type Service interface {
	Login(ctx context.Context, in *pb.Login) (ack *pb.LoginAck, err error)
}

type service struct {
	logger *zap.Logger
}

func NewService(log *zap.Logger) Service {
	var s Service
	{
		s = &service{logger: log}
		s = NewLogMiddle(log)(s)
	}
	return s
}

func (s service) Login(ctx context.Context, in *pb.Login) (*pb.LoginAck, error) {
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReq)),
		zap.String("func-->", "service.service.LoginEndpoint"),
		zap.Any("in", in))
	if in.Account != "yzh" && in.Passwd != "123456789" {
		return nil, errors.New("login error , account or password invalid")
	}
	token, err := utils.GenericToken(in.Account, 1)
	if err != nil {
		s.logger.Debug(fmt.Sprint(ctx.Value(ContextReq)),
			zap.String("func-->", "service.service.LoginEndpoint"),
			zap.String("token error", "generate token err"),
			zap.Error(err))
		return nil, err
	}
	ack := &pb.LoginAck{Token: token}
	s.logger.Debug(fmt.Sprint(ctx.Value(ContextReq)),
		zap.String("func-->", "service.service.LoginEndpoint"),
		zap.Any("token succeed", "token generate succeed"),
		zap.Any("ack", ack))
	return ack, nil
}
