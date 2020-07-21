package main

import (
	"golang.org/x/time/rate"
	"net/http"
)

//golang 令牌桶限流

var limiter = rate.NewLimiter(3, 2)

func limitHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func okHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK.\n"))
}

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/limit", okHandle)
	http.ListenAndServe(":1234", limitHandle(m))
}
