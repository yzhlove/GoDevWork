package utils

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

func TraceStack(extra ...interface{}) {
	if x := recover(); x != nil {
		log.Error(x)
		i := 0
		funcName, file, line, ok := runtime.Caller(i)
		for ok {
			log.Errorf("\n  \033[32m↓\033[0m frame:%v[func:%v,file:%v,line:%v]\n",
				i, runtime.FuncForPC(funcName).Name(), file, line)
			i++
			funcName, file, line, ok = runtime.Caller(i)
		}
		for k := range extra {
			log.Errorf("⌲ EXTRAS#%v DATA:%v \n", k, spew.Sdump(extra[k]))
		}
	}
}

func Trace(extra ...interface{}) {
	if x := recover(); x != nil {
		log.Error(x)
		var stack [32]uintptr
		i := runtime.Callers(3, stack[:])
		var sb strings.Builder
		sb.WriteString("☢︎ traceback:")
		for _, pc := range stack[:i] {
			funcName := runtime.FuncForPC(pc)
			file, line := funcName.FileLine(pc)
			sb.WriteString(fmt.Sprintf("\n  \033[32m↓\033[0m [\033[31m%s\033[0m] %s:%d", funcName.Name(), file, line))
		}
		for k := range extra {
			sb.WriteString(fmt.Sprintf("⌲ EXTRAS#%v DATA:%v \n", k, spew.Sdump(extra[k])))
		}
		log.Error(sb.String())
	}
}
