package main

import (
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger

func InitLogger() {
	logger, _ = zap.NewProduction()
}

func main() {
	InitLogger()
	defer logger.Sync()
	simpleGet("www.google.com")
	simpleGet("http://google.com")
	simpleGet("http://123455.com")
}

func simpleGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("Get Err :", zap.String("url", url), zap.Error(err))
	} else {
		logger.Info("Get Succeed ", zap.String("code", resp.Status), zap.String("url", url))
		resp.Body.Close()
	}
}
