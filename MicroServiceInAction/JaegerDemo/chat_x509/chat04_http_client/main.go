package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

//双向验证的客户端

const (
	absolute  = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_x509/"
	clientCrt = absolute + "ca_client.crt"
	clientKey = absolute + "ca_client.key"
	caCrt     = absolute + "ca.crt"
)

func main() {

	pair, err := tls.LoadX509KeyPair(clientCrt, clientKey)
	if err != nil {
		panic(err)
	}
	c := &http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{
			RootCAs:      loadCert(caCrt),
			Certificates: []tls.Certificate{pair},
		}},
	}
	if resp, err := c.Get("https://localhost:1234/hello"); err != nil {
		panic(err)
	} else {
		defer resp.Body.Close()
		io.Copy(os.Stdout, resp.Body)
	}
}

func loadCert(f string) *x509.CertPool {
	pool := x509.NewCertPool()
	if pem, err := ioutil.ReadFile(f); err != nil {
		panic(err)
	} else {
		pool.AppendCertsFromPEM(pem)
	}
	return pool
}
