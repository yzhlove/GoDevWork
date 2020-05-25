package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx-day4-base1/ziface"
)

type Connection struct {
	Conn       *net.TCPConn
	ConnID     uint32
	isClosed   bool
	MsgHandler ziface.MessageHandleInterface
	ExitChan   chan struct{}
}

func (c *Connection) StartReader() {
	fmt.Println("[conn] reader conn data ...")
	defer fmt.Println(c.RemoteAddr().String(), " conn reader exit!")
	defer c.Stop()

	for {
		dp := NewDataPack()
		head := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.Conn, head); err != nil {
			fmt.Println("[conn] read conn head err :", err)
			break
		}
		msg, err := dp.Unpack(head)
		if err != nil {
			fmt.Println("[conn] unpack err:", err)
			break
		}

		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.Conn, data); err != nil {
				fmt.Println("[conn] read msg data err:", err)
				break
			}
		}

		msg.SetData(data)
		go c.MsgHandler.DoMessageHandle(&Request{conn: c, msg: msg})
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

func (c *Connection) Send(msgID uint32, data []byte) error {
	if c.isClosed {
		return errors.New("conn closed when send message")
	}
	dp := NewDataPack()
	msg, err := dp.Pack(NewMessagePackage(msgID, data))
	if err != nil {
		fmt.Printf("pack err msg id: %d , %s\n", msgID, err)
		return err
	}

	if _, err := c.Conn.Write(msg); err != nil {
		fmt.Printf("[conn] write data msg id:%d %s\n", msgID, err)
		c.ExitChan <- struct{}{}
		return errors.New("conn write error")
	}
	return nil
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

func NewConn(conn *net.TCPConn, connID uint32, handler ziface.MessageHandleInterface) *Connection {
	return &Connection{
		Conn:       conn,
		ConnID:     connID,
		MsgHandler: handler,
		ExitChan:   make(chan struct{}, 1),
	}
}
