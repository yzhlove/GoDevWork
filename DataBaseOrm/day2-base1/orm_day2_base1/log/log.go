package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var flag = log.Lshortfile | log.LstdFlags
var errorLog = log.New(os.Stdout, "\033[31m[ERROR]\033[0m", flag)
var infoLog = log.New(os.Stdout, "\033[36m[INFO ]\033[0m", flag)
var loggers = []*log.Logger{errorLog, infoLog}
var mutex sync.Mutex

var ERROR = errorLog.Println
var ERRORF = errorLog.Printf
var INFO = infoLog.Println
var INFOF = infoLog.Printf

const (
	INFO_LEVEL = iota
	ERROR_LEVEL
	DISABLE
)

func SetLevel(level int) {
	mutex.Lock()
	defer mutex.Unlock()
	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}
	if ERROR_LEVEL < level {
		errorLog.SetOutput(ioutil.Discard)
	}
	if INFO_LEVEL < level {
		infoLog.SetOutput(ioutil.Discard)
	}
}
