package main

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func Test_Client(t *testing.T) {

	pool := x509.NewCertPool()
	if pem, err := ioutil.ReadFile(cert); err != nil {
		t.Error(err)
		return
	} else {
		pool.AppendCertsFromPEM(pem)
	}

	c := &http.Client{
		Transport: &http3.RoundTripper{
			TLSClientConfig: &tls.Config{
				RootCAs: pool,
				//InsecureSkipVerify: true,
			},
			QuicConfig: &quic.Config{},
		},
	}

	//1234 ，1235，1236 ok
	if resp, err := c.Get("https://abc.yzhdomain.com:1236/demo/test"); err != nil {
		t.Error(err)
		return
	} else {
		io.Copy(os.Stdout, resp.Body)
	}
}
