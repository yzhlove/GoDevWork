package geeseven

import (
	"log"
	"time"
)

func Logger() HandleFunc {
	return func(ctx *Context) {
		start := time.Now()
		ctx.Next()
		log.Printf("[time] (%d) <%s> -%v- ", ctx.StatusCode, ctx.Req.RequestURI, time.Since(start))
	}
}
