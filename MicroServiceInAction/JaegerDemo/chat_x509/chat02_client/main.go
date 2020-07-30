package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"strconv"
)

const (
	serverKey = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_x509/server.key"
	serverPem = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_x509/server.pem"
	clientPem = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_x509/client.pem"
	clientKey = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_x509/client.key"
)

func main() {

	cert, err := tls.LoadX509KeyPair(clientPem, clientKey)
	if err != nil {
		panic(err)
	}

	certPem, err := ioutil.ReadFile(clientPem)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(certPem); !ok {
		panic("cert err")
	}

	conf := &tls.Config{
		RootCAs:            certPool,
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", ":1234", conf)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var count int
	var buf = make([]byte, 1024)
	for {
		count++
		if _, err = conn.Write([]byte("hello." + strconv.Itoa(count) + "\n")); err != nil {
			log.Println(err)
			return
		}
		if n, err := conn.Read(buf); err != nil {
			log.Println("read err:", err)
			continue
		} else {
			log.Println("read message:" + string(buf[:n]))
		}
	}

}
