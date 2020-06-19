package log

import (
	"log"
	"os"
)

var (
	tag      = log.Lshortfile | log.LstdFlags
	errorLog = log.New(os.Stdout, "\033[31m[ERROR]\033[0m", tag)
	infoLog  = log.New(os.Stdout, "\033[36m[INFO ]\033[0m", tag)

	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)
