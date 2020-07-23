package utils

import (
	"go-kit-five/tool"
	"go.uber.org/zap"
)

var _logger *zap.Logger

func NewLoggerServer() {
	_logger = tool.NewLogger(
		tool.SetLogPrefix("go-kit"),
		tool.SetIsDevelopment(true),
		tool.SetLevel(zap.DebugLevel),
	)
}

func GetLog() *zap.Logger {
	return _logger
}
