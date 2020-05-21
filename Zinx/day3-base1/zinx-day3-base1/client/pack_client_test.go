package client

import (
	"net"
	"testing"
	"zinx-day3-base1/znet"
)

func Test_Client(t *testing.T) {

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		panic(err)
	}

	dp := znet.NewDataPack()
	msg1 := &znet.Message{
		ID:      0,
		DataLen: 5,
		Data:    []byte{'h', 'e', 'l', 'l', 'o'},
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		panic(err)
	}

	msg2 := &znet.Message{
		ID:      1,
		DataLen: 7,
		Data:    []byte{'w', 'o', 'r', 'l', 'd', '+', '!'},
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		panic(err)
	}

	sendData1 = append(sendData1, sendData2...)
	conn.Write(sendData1)
	select {}
}
