package transport

import (
	"context"
	"fmt"
	"go-kit-five/service"
	"go.uber.org/zap"
)

type LogErrHandle struct {
	logger *zap.Logger
}

func NewZapHandle(logger *zap.Logger) *LogErrHandle {
	return &LogErrHandle{logger: logger}
}

func (h *LogErrHandle) Handle(ctx context.Context, err error) {
	h.logger.Warn(fmt.Sprint(ctx.Value(service.ContextReq)), zap.Error(err))
}
