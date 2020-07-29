package src

import (
	"context"
	"errors"
	"github.com/go-kit/kit/metrics"
	"go-kit-seven/agent/pb"
	"go-kit-seven/utils"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Service interface {
	Login(ctx context.Context, in *pb.Login) (out *pb.LoginAck, err error)
}

type server struct {
	log *zap.Logger
}

func NewService(log *zap.Logger, count metrics.Counter, histogram metrics.Histogram) Service {
	return NewMetricsMid(count, histogram)(NewLogMid(log)(&server{log: log}))
}

func (s server) Login(ctx context.Context, in *pb.Login) (*pb.LoginAck, error) {
	if in.Account != "yuzihan" || in.Passwd != "123456789" {
		return nil, errors.New("account or password is invalid")
	}
	t := rand.Int31n(10-1) + 1
	time.Sleep(time.Duration(t) * time.Millisecond * 100)
	token, err := utils.GenericToken(in.Account, 1)
	if err != nil {
		return nil, err
	}
	return &pb.LoginAck{Token: token}, nil
}
