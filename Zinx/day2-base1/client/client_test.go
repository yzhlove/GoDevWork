package client

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func Test_Client(t *testing.T) {

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		t.Error(err)
		return
	}

	for {
		if _, err := conn.Write([]byte("Zinx v0.2")); err != nil {
			t.Error(err)
			return
		}
		buf := make([]byte, 512)
		if cnt, err := conn.Read(buf); err != nil {
			t.Error(err)
			return
		} else {
			fmt.Printf("server read back cnt:%d str:%s \n", cnt, buf)
		}
		time.Sleep(time.Second)
	}
}
