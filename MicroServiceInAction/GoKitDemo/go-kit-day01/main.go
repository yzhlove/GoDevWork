package main

import (
	"go-kit-one/endpoint"
	"go-kit-one/service"
	"go-kit-one/transport"
	"net/http"
)

func main() {
	server := service.NewService()
	endpoints := endpoint.NewEndpointServer(server)
	httpHandle := transport.NewHttpHandle(endpoints)
	if err := http.ListenAndServe(":6688", httpHandle); err != nil {
		panic(err)
	}
}
