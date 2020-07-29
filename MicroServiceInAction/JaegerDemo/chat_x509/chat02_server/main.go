package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strconv"
)

const (
	serverKey = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_x509/server.key"
	serverPem = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_x509/server.pem"
	clientPem = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_x509/client.pem"
)

func main() {

	var (
		//server_key, server_pem []byte
		client_pem []byte
		err        error
	)

	//if server_key, err = ioutil.ReadFile(serverKey); err != nil {
	//	panic(err)
	//}
	//
	//if server_pem, err = ioutil.ReadFile(serverPem); err != nil {
	//	panic(err)
	//}

	if client_pem, err = ioutil.ReadFile(clientPem); err != nil {
		panic(err)
	}

	cert, err := tls.LoadX509KeyPair(serverPem, serverKey)
	if err != nil {
		panic(err)
	}

	clientCertPool := x509.NewCertPool()
	if ok := clientCertPool.AppendCertsFromPEM(client_pem); !ok {
		panic("cert file err")
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    clientCertPool,
	}
	ln, err := tls.Listen("tcp", ":1234", config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()
	for {
		if conn, err := ln.Accept(); err != nil {
			fmt.Println("accept err:", err)
			continue
		} else {
			go hConn(conn)
		}
	}
}

func hConn(conn net.Conn) {
	defer conn.Close()
	var (
		reader = bufio.NewReader(conn)
		msg    string
		err    error
		count  int
	)
	for msg, err = reader.ReadString('\n'); err == nil; {
		count++
		fmt.Println("msg => ", msg)
		if _, err = conn.Write([]byte("send:" + strconv.Itoa(count) + "\n")); err != nil {
			fmt.Println("write message err:", err)
			return
		}
	}
	fmt.Println("conn err:", err)
}
