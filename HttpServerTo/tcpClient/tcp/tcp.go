package tcp

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

type TcpClient struct {
	net.Conn
	buf *bufio.Reader
}

type Result struct {
	Opt   string
	Key   string
	Value string
	Err   error
}

func readBuf(reader *bufio.Reader) (length int) {
	if s, err := reader.ReadString(' '); err != nil {
		log.Println("read buf err: " + err.Error())
	} else {
		if length, err = strconv.Atoi(strings.TrimSpace(s)); err != nil {
			log.Printf("read buf length transform err: " + s)
		}
	}
	return
}

func (c *TcpClient) send(s string) {
	_, _ = c.Write([]byte(s))
}

func (c *TcpClient) get(key string) {
	s := fmt.Sprintf("G%d %s", len(key), key)
	c.send(s)
}

func (c *TcpClient) set(key, value string) {
	s := fmt.Sprintf("S%d %d %s%s", len(key), len(value), key, value)
	c.send(s)
}

func (c *TcpClient) del(key string) {
	s := fmt.Sprintf("D%d %s", len(key), key)
	c.send(s)
}

func (c *TcpClient) recv() (str string, err error) {
	size := readBuf(c.buf)
	if size != 0 {
		status := false
		if size < 0 {
			status, size = true, -size
		}
		msg := make([]byte, size)
		if _, err = io.ReadFull(c.buf, msg); err != nil {
			return
		}
		if status {
			err = errors.New(string(msg))
		} else {
			str = string(msg)
		}
	}
	return
}

func (c *TcpClient) Run(result *Result) {
	switch result.Opt {
	case "get":
		c.get(result.Key)
		result.Value, result.Err = c.recv()
	case "set":
		c.set(result.Key, result.Value)
		_, result.Err = c.recv()
	case "del":
		c.del(result.Key)
		_, result.Err = c.recv()
	default:
		panic("unknown option type :" + result.Opt)
	}
	fmt.Println("Result => ", *result)
}

func (c *TcpClient) RunPipeline(rs []*Result) {
	if len(rs) > 0 {
		for _, result := range rs {
			switch result.Opt {
			case "set":
				c.set(result.Key, result.Value)
			case "get":
				c.get(result.Key)
			case "del":
				c.del(result.Key)
			}
		}
		for _, result := range rs {
			result.Value, result.Err = c.recv()
			fmt.Println("Result => ", *result)
		}
	}
}

func New() *TcpClient {
	if conn, err := net.Dial("tcp", ":1234"); err != nil {
		panic(err)
	} else {
		return &TcpClient{conn, bufio.NewReader(conn)}
	}
}
