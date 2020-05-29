package utils

import (
	"fmt"
	"runtime"
	"strings"
)

func Trace(top string) string {
	var stack [32]uintptr
	i := runtime.Callers(0, stack[:])
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("{%s} TraceBack:\n", top))
	for _, pc := range stack[:i] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		sb.WriteString(fmt.Sprintf("\n\t↓ [%s] %s:%d", fn.Name(), file, line))
	}
	return sb.String()
}
