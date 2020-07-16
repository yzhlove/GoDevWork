package transport

import (
	"context"
	"fmt"
	"go-kit-two/service"
	"go.uber.org/zap"
)

type LogErrorHandle struct {
	logger *zap.Logger
}

func NewZapLogErrorHandler(logger *zap.Logger) *LogErrorHandle {
	return &LogErrorHandle{
		logger: logger,
	}
}

func (h *LogErrorHandle) Handle(ctx context.Context, err error) {
	h.logger.Warn(fmt.Sprint(ctx.Value(service.ContextReqUUID)), zap.Error(err))
}
