package utils

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

type LogErrorHandle struct {
	log *zap.Logger
}

func NewZapErrorHandle(log *zap.Logger) *LogErrorHandle {
	return &LogErrorHandle{log: log}
}

func (h *LogErrorHandle) Handle(ctx context.Context, err error) {
	h.log.Warn(fmt.Sprint(ctx.Value("context_req_uid")), zap.Error(err))
}
