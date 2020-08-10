package src

import (
	"context"
	"errors"
	"github.com/go-kit/kit/metrics"
	"github.com/opentracing/opentracing-go"
	"go-kit-nine/agent/pb"
	"go-kit-nine/utils"
	"go.uber.org/zap"
	"hash/crc32"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Service interface {
	Login(ctx context.Context, in *pb.UserLogic_Login) (*pb.UserLogic_LoginAck, error)
}

func NewService(logger *zap.Logger, c metrics.Counter, h metrics.Histogram, tracer opentracing.Tracer) Service {
	var s Service
	{
		s = &server{logger: logger}
		s = NewTracerMiddle(tracer)(s)
		s = NewHystrixMiddle(c, h)(s)
		s = NewLoggerMiddle(logger)(s)
	}
	return s
}

type server struct {
	logger *zap.Logger
}

func (s server) Login(ctx context.Context, in *pb.UserLogic_Login) (*pb.UserLogic_LoginAck, error) {
	if in.Account != "yuzihan" || in.Password != "12345" {
		return nil, errors.New("login error:account or password invalid")
	}

	if token, err := utils.GenerToken(in.Account, int(crc32.ChecksumIEEE([]byte(utils.GetUID())))); err != nil {
		return nil, err
	} else {
		return &pb.UserLogic_LoginAck{Token: token}, nil
	}
}
