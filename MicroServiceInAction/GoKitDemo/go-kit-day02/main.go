package main

import (
	"go-kit-two/endpoint"
	"go-kit-two/service"
	"go-kit-two/transport"
	"go-kit-two/utils"
	"net/http"
)

func main() {

	utils.NewLoggerServer()
	server := service.NewService(utils.GetLog())
	endpoints := endpoint.NewEndpointServer(server, utils.GetLog())
	httpHandler := transport.NewHttpHandler(endpoints, utils.GetLog())
	utils.GetLog().Info("server run 0.0.0.0:1234")
	if err := http.ListenAndServe(":1234", httpHandler); err != nil {
		panic(err)
	}
}
