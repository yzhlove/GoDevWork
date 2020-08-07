package src

import (
	"context"
	"errors"
	"github.com/go-kit/kit/metrics"
	"go-kit-eight/agent/pb"
	"go-kit-eight/utils"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Service interface {
	Login(ctx context.Context, in *pb.Login) (*pb.LoginAck, error)
}

type server struct {
	log *zap.Logger
}

func NewService(log *zap.Logger, c metrics.Counter, h metrics.Histogram) Service {
	return NewMetricsMid(c, h)(NewLogMid(log)(&server{log: log}))
}

func (server) Login(ctx context.Context, in *pb.Login) (*pb.LoginAck, error) {
	if !(in.Account == "yuzihan" && in.Passwd == "12345") {
		return nil, errors.New("login err: account or password failed")
	}
	if rand.Intn(10) > 3 {
		return nil, errors.New("service running err")
	}
	token, err := utils.GenericToken(in.Account, 1)
	if err != nil {
		return nil, err
	}
	return &pb.LoginAck{Token: token}, nil
}
