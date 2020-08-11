package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"github.com/lucas-clemente/quic-go"
	"log"
	"math/big"
	"strconv"
	"time"
)

func main() {

	go EchoServer()
	go EchoClient()
	select {}
}

func GenerTLSCfg() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})
	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{Certificates: []tls.Certificate{tlsCert}, NextProtos: []string{"quic"}}
}

func EchoServer() {
	lis, err := quic.ListenAddr(":1234", GenerTLSCfg(), nil)
	if err != nil {
		panic(err)
	}
	conn, err := lis.Accept(context.Background())
	if err != nil {
		panic(err)
	}
	stream, err := conn.AcceptStream(context.Background())
	if err != nil {
		panic(err)
	}

	go func(stream quic.Stream) {
		var buffer = make([]byte, 1024)
		for {
			if n, err := stream.Read(buffer); err != nil {
				panic(err)
			} else {
				log.Println("[server] read message:", string(buffer[:n]))
			}
		}
	}(stream)

	var count int
	for {
		count++
		msg := "[server] quic send message:" + strconv.Itoa(count)
		if n, err := stream.Write([]byte(msg)); err != nil {
			panic(err)
		} else {
			log.Println("[client] msg len:", len(msg), " send len:", n)
		}
		time.Sleep(time.Second * 2)
	}

}

func EchoClient() {

	tlsCfg := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic"},
	}
	conn, err := quic.DialAddr("localhost:1234", tlsCfg, nil)
	if err != nil {
		panic(err)
	}

	stream, err := conn.OpenStreamSync(context.Background())
	if err != nil {
		panic(err)
	}

	go func(stream quic.Stream) {
		var buffer = make([]byte, 1024)
		for {
			if n, err := stream.Read(buffer); err != nil {
				log.Println("client read err:", err)
				return
			} else {
				log.Println("[client] read message:", string(buffer[:n]))
			}
		}
	}(stream)

	var count int
	for {
		count++
		msg := "[client] quic send message:" + strconv.Itoa(count)
		if n, err := stream.Write([]byte(msg)); err != nil {
			panic(err)
		} else {
			log.Println("[client] msg len:", len(msg), " send len:", n)
		}
		time.Sleep(time.Second * 5)
	}

}
