package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
)

const (
	absolute  = "/Users/yostar/workSpace/gowork/src/GoDevWork/MicroServiceInAction/JaegerDemo/chat_x509/"
	serverCrt = absolute + "ca_server.crt"
	serverKey = absolute + "ca_server.key"
	caCrt     = absolute + "ca.crt"
)

func main() {

	s := http.Server{
		Addr:    ":1234",
		Handler: hello{},
		TLSConfig: &tls.Config{
			ClientCAs:  loadCert(caCrt),
			ClientAuth: tls.RequireAndVerifyClientCert,
		}}

	if err := s.ListenAndServeTLS(serverCrt, serverKey); err != nil {
		panic(err)
	}
}

type hello struct{}

func (hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain;charset=utf-8")
	w.Write([]byte("hello <tls>!!!"))
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
