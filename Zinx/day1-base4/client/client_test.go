package client

import (
	"net"
	"testing"
	"time"
)

func Test_Client(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		t.Error(err)
		return
	}
	for {
		_, err := conn.Write([]byte("what are you doing !"))
		if err != nil {
			t.Error(err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log("server call ", " cnt = ", cnt, " str = ", string(buf))
		time.Sleep(time.Second)
	}
}
