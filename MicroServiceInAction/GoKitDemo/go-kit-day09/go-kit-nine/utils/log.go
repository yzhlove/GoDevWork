package utils

import (
	"go-kit-nine/tool"
	"go.uber.org/zap"
)

const CONTEXT_UID = "context_request_id"

var _logger *zap.Logger

func NewLogger() {
	_logger = tool.NewLogger(
		tool.SetLogPrefix("go-kit"),
		tool.SetIsDevelopment(true),
		tool.SetLevel(zap.DebugLevel))
}

func GetLog() *zap.Logger {
	return _logger
}
