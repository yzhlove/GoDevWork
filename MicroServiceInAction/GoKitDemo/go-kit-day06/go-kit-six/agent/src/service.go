package src

import (
	"context"
	"errors"
	"go-kit-six/agent/pb"
	"go-kit-six/utils"
	"go.uber.org/zap"
)

type Service interface {
	Login(ctx context.Context, in *pb.Login) (ack *pb.LoginAck, err error)
}

type service struct {
	log *zap.Logger
}

func NewService(log *zap.Logger) Service {
	return NewLogMiddle(log)(&service{log: log})
}

func (s service) Login(ctx context.Context, in *pb.Login) (ack *pb.LoginAck, err error) {
	if in.Account != "yzh" && in.Passwd != "123456789" {
		err = errors.New("account or password invalid")
		return
	}
	token, err := utils.GenericToken(in.Account, 1)
	if err != nil {
		return
	}
	ack = &pb.LoginAck{Token: token}
	return
}
