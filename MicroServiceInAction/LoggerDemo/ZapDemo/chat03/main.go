package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
)

var sugarLogger *zap.SugaredLogger

func InitLogger() {
	writeSyncer := getLogWrite()
	//encoder := getEncode()
	//encoder := getGenericEncode()
	encoder := getTimeEncode()
	core := zapcore.NewCore(encoder, writeSyncer, zap.InfoLevel)

	//logger := zap.New(core)
	logger := zap.New(core, zap.AddCaller()) // 添加调用者的信息
	sugarLogger = logger.Sugar()
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

func getEncode() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getGenericEncode() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
}

func getTimeEncode() zapcore.Encoder {
	encodeCfg := zap.NewProductionEncoderConfig()
	encodeCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encodeCfg)
}

func getLogWrite() zapcore.WriteSyncer {
	file, _ := os.Create("zap.log")
	return zapcore.AddSync(file)
}
