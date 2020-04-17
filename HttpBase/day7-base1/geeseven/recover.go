package geeseven

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func trace(err string) string {
	var stack [32]uintptr
	i := runtime.Callers(3, stack[:])
	var str strings.Builder
	str.WriteString(err + "\nTraceBack:")
	for _, pc := range stack[:i] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t[%s]%s:%d", fn.Name(), file, line))
	}
	return str.String()
}

func Recovery() HandleFunc {
	return func(ctx *Context) {
		defer func() {
			if err := recover(); err != nil {
				str := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(str))
				ctx.Error(http.StatusInternalServerError, "Server Error")
			}
		}()
		ctx.Next()
	}
}
