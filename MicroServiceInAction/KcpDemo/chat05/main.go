package main

import (
	"github.com/lucas-clemente/quic-go/http3"
	"log"
	"net/http"
)

const (
	key  = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_x509/server.key"
	cert = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_x509/server.pem"
)

const addr = ":1234"

func main() {
	//HttpServer()
	QuicServer()
	select {}
}

func HttpServer() {
	log.Println(http.ListenAndServeTLS(addr, cert, key, handle()))
}

func QuicServer() {
	log.Println(http3.ListenAndServeQUIC(addr, cert, key, handle()))
}

func handle() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	return mux
}
