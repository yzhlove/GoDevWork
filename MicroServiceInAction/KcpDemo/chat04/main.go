package main

import (
	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"

	"log"
	"net/http"
	_ "net/http/pprof"
)

//http3-quic

const (
	key  = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_x509/server.key"
	cert = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_x509/server.pem"
)

func main() {

	//pprof
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	go func() {
		log.Println(http.ListenAndServeTLS(":1230", cert, key, setupHandle()))
	}()

	go func() {
		log.Println(http3.ListenAndServe(":1234", cert, key, setupHandle()))
	}()

	go func() {
		log.Println(http3.ListenAndServeQUIC(":1235", cert, key, setupHandle()))
	}()

	go func() {
		quicConfig := &quic.Config{}
		server := http3.Server{
			Server:     &http.Server{Handler: setupHandle(), Addr: ":1236"},
			QuicConfig: quicConfig,
		}
		log.Println(server.ListenAndServeTLS(cert, key))
	}()

	select {}
}

func setupHandle() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/demo/test", func(writer http.ResponseWriter, request *http.Request) {
		msg := "host:" + request.Host + " remote:" + request.RemoteAddr
		writer.Header().Add("alt-svc", `quic=":443"; ma=2592000; v="38,37,36"`)
		writer.Write([]byte(msg + " HELLO WORLD!"))
	})
	return mux
}
