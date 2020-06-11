package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	flag     = log.Lshortfile | log.LstdFlags
	errorLog = log.New(os.Stdout, "\033[31m[ERROR]\033[0m", flag)
	infoLog  = log.New(os.Stdout, "\033[36m[INFO ]\033[0m", flag)
	loggers  = []*log.Logger{errorLog, infoLog}
	mutex    sync.Mutex
)

var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

const (
	InfoLevel = iota
	ErrorLevel
	Disable
)

func SetLevel(level int) {
	mutex.Lock()
	defer mutex.Unlock()
	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}
	if ErrorLevel < level {
		errorLog.SetOutput(ioutil.Discard)
	}
	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}
}
