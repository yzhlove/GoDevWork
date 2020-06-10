package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

func main() {
	f()
}

func panicFunc() {
	panic("test panic info.")
}

func f() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(trace(fmt.Sprint(err)))
			log.Println("==============================================")
			log.Println(stack(fmt.Sprint(err)))
		}
	}()
	panicFunc()
}

func trace(top string) string {
	var stack [32]uintptr
	i := runtime.Callers(0, stack[:])
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("{%s} trace:", top))
	for _, pc := range stack[:i] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		sb.WriteString(fmt.Sprintf("\n\t↓ [%s] %s:%d", fn.Name(), file, line))
	}
	return sb.String()
}

func stack(top string) string {
	sb := make([]byte, 2<<10)
	i := runtime.Stack(sb, true)
	return top + string(sb[:i])
}
