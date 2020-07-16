package main

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"sync"
)

var sugarLogger *zap.SugaredLogger

func InitLogger() {

	writer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

func main() {

	InitLogger()
	defer sugarLogger.Sync()

	urls := []string{"www.google.com", "http://www.google.com"}
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		fmt.Println("fetch index --> ", i)
		go func(index uint) {
			defer wg.Done()
			httpGet(index, urls[index%2])
		}(uint(i))
	}
	wg.Wait()

}

func getEncoder() zapcore.Encoder {
	encodeCfg := zap.NewProductionEncoderConfig()
	encodeCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encodeCfg)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "lumberjack.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func httpGet(index uint, url string) {
	sugarLogger.Debugf("[index(%d)] Trying Get Request %s ", index, url)
	if resp, err := http.Get(url); err != nil {
		sugarLogger.Errorf("Error URL:%s Err:", url, err)
	} else {
		sugarLogger.Infof("Succeed Code:%s URL:%s", resp.Status, url)
		resp.Body.Close()
	}
}
