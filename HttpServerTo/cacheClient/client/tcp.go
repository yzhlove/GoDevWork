package client

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

type tcpClient struct {
	net.Conn
	r *bufio.Reader
}

func (c *tcpClient) sendGet(key string) {
	str := fmt.Sprintf("G%d %s", len(key), key)
	_, _ = c.Write([]byte(str))
}

func (c *tcpClient) sendSet(key, value string) {
	str := fmt.Sprintf("S%d %d %s%s", len(key), len(value), key, value)
	_, _ = c.Write([]byte(str))
}

func (c *tcpClient) sendDel(key string) {
	str := fmt.Sprintf("D%d %s", len(key), key)
	_, _ = c.Write([]byte(str))
}

func readLen(r *bufio.Reader) int {
	if str, err := r.ReadString(' '); err != nil {
		log.Println("read length err :" + err.Error())
		return 0
	} else {
		if length, err := strconv.Atoi(strings.TrimSpace(str)); err != nil {
			log.Println(str + " transform err:" + err.Error())
			return 0
		} else {
			return length
		}
	}
}

func (c *tcpClient) recvResp() (str string, err error) {
	length := readLen(c.r)
	if length == 0 {
		return
	}
	if length < 0 {
		errMsg := make([]byte, -length)
		if _, err = io.ReadFull(c.r, errMsg); err != nil {
			return
		}
		err = errors.New(string(errMsg))
	} else {
		value := make([]byte, length)
		if _, err = io.ReadFull(c.r, value); err != nil {
			return
		}
		str = string(value)
	}
	return
}

func (c *tcpClient) Run(msg *Message) {
	switch msg.Name {
	case "get":
		c.sendGet(msg.Key)
		msg.Key, msg.Error = c.recvResp()
		return
	case "set":
		c.sendSet(msg.Key, msg.Value)
		_, msg.Error = c.recvResp()
		return
	case "del":
		c.sendDel(msg.Key)
		_, msg.Error = c.recvResp()
		return
	}
	panic("unknown message type " + msg.Name)
}

func (c *tcpClient) PipeLineRun(msgs []*Message) {
	if len(msgs) > 0 {
		for _, msg := range msgs {
			switch msg.Name {
			case "get":
				c.sendGet(msg.Key)
			case "set":
				c.sendSet(msg.Key, msg.Value)
			case "del":
				c.sendDel(msg.Key)
			}
		}
		for _, msg := range msgs {
			msg.Value, msg.Error = c.recvResp()
		}
	}
}

func newTcpClient(server string) *tcpClient {
	conn, err := net.Dial("tcp", server+":1234")
	if err != nil {
		panic(err)
	}
	return &tcpClient{conn, bufio.NewReader(conn)}
}
