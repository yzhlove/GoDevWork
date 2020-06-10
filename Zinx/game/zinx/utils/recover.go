package utils

import (
	"fmt"
	"runtime"
	"strings"
)

func Trace(top string) string {
	var stack [32]uintptr
	i := runtime.Callers(2, stack[:])
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("{%s} TraceBack:", top))
	for _, pc := range stack[:i] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		sb.WriteString(fmt.Sprintf("\n  \033[32m↓\033[0m [\033[31m%s\033[0m] %s:%d", fn.Name(), file, line))
	}
	return sb.String()
}
