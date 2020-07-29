package main

import "crypto/tls"

const (
	serverKey = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_x509/server.key"
	serverPem = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_x509/server.pem"
	clientPem = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_x509/client.pem"
	clientKey = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_x509/client.key"
)

func main() {

	cert , err := tls.LoadX509KeyPair(clientPem , clientKey)
	if err != nil {
		panic(err)
	}



}
