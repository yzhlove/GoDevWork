package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	serverKey = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_x509/server.key"
	serverPem = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_x509/server.pem"
	clientPem = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_x509/client.pem"
)

func main() {

	c := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: loadCert(serverPem),
		}}}
	//abc.yzhdomain.com => etc/hosts配置
	if resp, err := c.Get("https://abc.yzhdomain.com:1234/hello"); err != nil {
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
