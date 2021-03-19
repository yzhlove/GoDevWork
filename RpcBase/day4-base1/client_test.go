package day4_base1

import (
	"context"
	"day4-base1/codec"
	"net"
	"testing"
	"time"
)

type Bar int

func (b Bar) Timeout(argv int, reply *int) error {
	time.Sleep(time.Second * 2)
	return nil
}

func startServer(addr chan string) {
	var b Bar
	_ = Register(&b)
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	addr <- l.Addr().String()
	Accept(l)
}

func TestClient_dialTimeout(t *testing.T) {
	t.Parallel()

	l, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	f := func(conn net.Conn, code *CheckCode) (*Client, error) {
		conn.Close()
		time.Sleep(time.Second * 2)
		return nil, nil
	}

	t.Run("timeout", func(t *testing.T) {
		_, err := dialTimout(f, "tcp", l.Addr().String(), &CheckCode{ConnectTimeout: time.Second})
		if err != nil {
			t.Log(err)
		}
	})
	t.Run("0", func(t *testing.T) {
		_, err := dialTimout(f, "tcp", l.Addr().String(), &CheckCode{ConnectTimeout: 0})
		if err != nil {
			t.Error(err)
		}
	})

}

func TestClient_Call(t *testing.T) {
	t.Parallel()
	chAddr := make(chan string)
	go startServer(chAddr)
	address := <-chAddr
	time.Sleep(time.Second)
	t.Run("client timeout", func(t *testing.T) {
		client, err := Dial("tcp", address, &CheckCode{
			Code: MagicCode,
			Type: codec.Gob,
		})
		if err != nil {
			t.Fatal(err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		var reply int
		if err := client.Call(ctx, "Bar.Timeout", 1, &reply); err != nil {
			t.Log("timeout ---> ", err)
		}
	})
	t.Run("server handle timeout", func(t *testing.T) {
		client, err := Dial("tcp", address, &CheckCode{
			Code:          MagicCode,
			Type:          codec.Gob,
			HandleTimeout: time.Second,
		})
		if err != nil {
			t.Fatal(err)
		}
		var reply int
		if err := client.Call(context.Background(), "Bar.Timeout", 1, &reply); err != nil {
			t.Log("handle ----->", err)
		}
	})
}
