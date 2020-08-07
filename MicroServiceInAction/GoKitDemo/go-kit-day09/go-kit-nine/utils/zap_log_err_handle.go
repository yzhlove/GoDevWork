package utils

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

type LogErrHandle struct {
	logger *zap.Logger
}

func NewZapLogErrHandle(logger *zap.Logger) *LogErrHandle {
	return &LogErrHandle{logger: logger}
}

func (h *LogErrHandle) Handle(ctx context.Context, err error) {
	h.logger.Warn(fmt.Sprint(ctx.Value(CONTEXT_UID), zap.Error(err)))
}
