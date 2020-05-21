package znet

import (
	"fmt"
	"net"
	"zinx-day2-base1/ziface"
)

type Connection struct {
	Conn     *net.TCPConn
	ConnID   uint32
	isClosed bool
	Router   ziface.RouterInterface
	ExitChan chan struct{}
}

func NewConn(conn *net.TCPConn, connID uint32, r ziface.RouterInterface) *Connection {
	return &Connection{
		Conn:     conn,
		ConnID:   connID,
		Router:   r,
		ExitChan: make(chan struct{}, 1),
	}
}

func (c *Connection) GetTcp() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) StartReader() {
	fmt.Println("[conn] reader conn data ...")
	defer fmt.Println(c.RemoteAddr().String(), " reader exit!")
	defer c.Stop()
	for {
		buf := make([]byte, 512)
		if _, err := c.Conn.Read(buf); err != nil {
			fmt.Println("[conn] read err ", err)
			c.ExitChan <- struct{}{}
			return
		}
		request := &Request{conn: c, data: buf}
		go func(req ziface.RequestInterface) {
			if c.Router != nil {
				c.Router.BeforeHandle(req)
				c.Router.Handle(req)
				c.Router.AfterHandle(req)
			}
		}(request)
	}
}

func (c *Connection) Start() {
	go c.StartReader()
	for {
		select {
		case <-c.ExitChan:
			return
		}
	}
}

func (c *Connection) Stop() {
	if !c.isClosed {
		c.isClosed = true
		c.Conn.Close()
		c.ExitChan <- struct{}{}
		close(c.ExitChan)
	}
}
