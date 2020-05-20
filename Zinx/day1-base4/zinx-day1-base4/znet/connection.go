package znet

import (
	"fmt"
	"net"
	"zinx-day1-base4/ziface"
)

type Connection struct {
	Conn      *net.TCPConn
	ConnID    uint32
	isClosed  bool
	handleAPI ziface.HandleFunc
	die       chan struct{}
}

func NewConnection(conn *net.TCPConn, connID uint32, callback ziface.HandleFunc) *Connection {
	return &Connection{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		handleAPI: callback,
		die:       make(chan struct{}, 1),
	}
}

func (c *Connection) StartReader() {
	fmt.Println("Reader thread is running ...")
	defer fmt.Println(c.RemoteAddr().String(), " conn reader exit!")
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err ", err)
			c.die <- struct{}{}
			return
		}
		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			c.die <- struct{}{}
			return
		}
	}
}

func (c *Connection) Start() {
	go c.StartReader()
	for {
		select {
		case <-c.die:
			return
		}
	}
}

func (c *Connection) Stop() {
	if c.isClosed {
		return
	}
	c.isClosed = true
	c.Conn.Close()
	c.die <- struct{}{}
	close(c.die)
}

func (c *Connection) GetTcpConn() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
