package src

import (
	"context"
	"fmt"
	"go-kit-six/agent/pb"
	"go.uber.org/zap"
	"time"
)

const CONTEXT_REQ_UID = "context_req_uid"

type MiddleFunc func(s Service) Service

type logMiddle struct {
	log  *zap.Logger
	next Service
}

func NewLogMiddle(log *zap.Logger) MiddleFunc {
	return func(s Service) Service {
		return &logMiddle{log: log, next: s}
	}
}

func (l logMiddle) Login(ctx context.Context, in *pb.Login) (ack *pb.LoginAck, err error) {
	defer func(start time.Time) {
		l.log.Debug(fmt.Sprint(ctx.Value(CONTEXT_REQ_UID)),
			zap.String("<func>", "service.middle.log"),
			zap.Any("in", in),
			zap.Any("out", ack),
			zap.Int64("time", time.Since(start).Milliseconds()),
			zap.Error(err))
	}(time.Now())
	ack, err = l.next.Login(ctx, in)
	return
}
