package main

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/lucas-clemente/quic-go/http3"
	"io/ioutil"
	"net/http"
	"testing"
)

func returnX509Pool() *x509.CertPool {
	pool := x509.NewCertPool()
	if pem, err := ioutil.ReadFile(cert); err != nil {
		panic(err)
	} else {
		pool.AppendCertsFromPEM(pem)
	}
	return pool
}

const url = "https://abc.yzhdomain.com:1234/hello"

func BenchmarkHttpServer(b *testing.B) {

	for i := 0; i < 1000; i++ {
		go func() {
			c := http.Client{Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs: returnX509Pool()}}}
			defer c.CloseIdleConnections()
			if _, err := c.Get(url); err != nil {
				//b.Error(err)
			}
		}()
	}
}

// 1 - 4223 ns/op
// 2 - 0.00263 ns/op

func BenchmarkQuicServer(b *testing.B) {
	//c := http.Client{Transport: &http3.RoundTripper{
	//	TLSClientConfig: &tls.Config{RootCAs: returnX509Pool()},
	//}}
	for i := 0; i < 1000; i++ {
		go func() {
			c := http.Client{Transport: &http3.RoundTripper{
				TLSClientConfig: &tls.Config{RootCAs: returnX509Pool()},
			}}
			defer c.CloseIdleConnections()
			if _, err := c.Get(url); err != nil {
				//b.Error(err)
			}
		}()
	}
}

//1 - 3903
//2 - 0.000613 ns/op
