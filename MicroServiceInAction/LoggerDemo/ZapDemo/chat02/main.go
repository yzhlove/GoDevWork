package main

import (
	"go.uber.org/zap"
	"net/http"
)

var sugarLogger *zap.SugaredLogger

func InitLogger() {
	if logger, err := zap.NewProduction(); err != nil {
		panic(err)
	} else {
		sugarLogger = logger.Sugar()
	}
}

func main() {
	InitLogger()
	defer sugarLogger.Sync()

	simpleGet("www.google.com")
	simpleGet("http://www.google.com")
}

func simpleGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s ", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error :%s", url, err)
	} else {
		sugarLogger.Infof("Succeed Code : %s URL %s ", resp.Status, url)
		resp.Body.Close()
	}
}
