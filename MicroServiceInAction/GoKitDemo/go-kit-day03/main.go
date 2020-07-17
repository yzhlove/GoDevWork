package main

import (
	ept "go-kit-three/endpoint"
	"go-kit-three/service"
	"go-kit-three/transport"
	"go-kit-three/utils"
	"net/http"
)

func main() {

	utils.NewLoggerServer()
	server := service.NewService(utils.GetLog())
	endpoint := ept.NewEndpointServer(server, utils.GetLog())
	httpHandle := transport.NewHttpHandler(endpoint, utils.GetLog())
	if err := http.ListenAndServe(":1234", httpHandle); err != nil {
		panic(err)
	}
}
