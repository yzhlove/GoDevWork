package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strconv"
)

const SERVER_KEY = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_x509/server.key"
const SERVER_PEM = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_x509/server.pem"

func main() {

	key, err := ioutil.ReadFile(SERVER_KEY)
	if err != nil {
		panic("key :" + err.Error())
	}

	pem, err := ioutil.ReadFile(SERVER_PEM)
	if err != nil {
		panic("pem " + err.Error())
	}

	cert, err := tls.X509KeyPair(pem, key)
	if err != nil {
		panic("cert " + err.Error())
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	ln, err := tls.Listen("tcp", ":1234", config)
	if err != nil {
		panic("listener " + err.Error())
	}
	defer ln.Close()
	for {
		if conn, err := ln.Accept(); err != nil {
			log.Println("Accept err:" + err.Error())
			continue
		} else {
			go hConn(conn)
		}
	}
}

func hConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	var count int
	for {
		count++
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Printf("conn read error: %s \n", err.Error())
			return
		}
		fmt.Println("msg->", msg)
		data := "succeed." + strconv.Itoa(count) + "\n"
		n, err := conn.Write([]byte(data))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}
