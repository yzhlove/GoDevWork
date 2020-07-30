package main

import "net/http"

const (
	serverKey = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_x509/server.key"
	serverPem = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_x509/server.pem"
	clientPem = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_x509/client.pem"
)

func main() {

	http.HandleFunc("/hello", Hello)
	if err := http.ListenAndServeTLS(":1234", serverPem, serverKey, nil); err != nil {
		panic(err)
	}

}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")
	w.Write([]byte("hello tls!!!"))
}
