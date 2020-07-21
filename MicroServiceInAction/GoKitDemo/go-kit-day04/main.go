package main

import (
	ept "go-kit-four/endpoint"
	"go-kit-four/service"
	"go-kit-four/transport"
	"go-kit-four/utils"
	"go.uber.org/ratelimit"
	"golang.org/x/time/rate"
	"net/http"
)

func main() {
	utils.NewLoggerServer()
	golangRate := rate.NewLimiter(5, 2)
	uberRate := ratelimit.New(2) //一秒请求两次
	s := service.NewService(utils.GetLog())
	endpoint := ept.NewEndpoint(s, utils.GetLog(), golangRate, uberRate)
	httpHandle := transport.GetHttpHandle(endpoint, utils.GetLog())
	utils.GetLog().Info("server run :1234")
	if err := http.ListenAndServe(":1234", httpHandle); err != nil {
		panic(err)
	}
}
