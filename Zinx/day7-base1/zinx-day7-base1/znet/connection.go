package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx-day7-base1/config"
	"zinx-day7-base1/ziface"
)

type ConnContext struct {
	tcpServer  ziface.ServerInterface
	Conn       *net.TCPConn
	ID         uint32
	isClosed   bool
	msgHandle  ziface.MsgHandleInterface
	msgChan    chan []byte
	msgBufChan chan []byte
	exitChan   chan struct{}
}

func NewConnContext(conn *net.TCPConn, connID uint32, server ziface.ServerInterface) *ConnContext {
	context := &ConnContext{
		tcpServer:  server,
		Conn:       conn,
		ID:         connID,
		msgHandle:  server.GetMsgHandle(),
		msgChan:    make(chan []byte),
		msgBufChan: make(chan []byte, config.GlobalConfig.MaxWorkerTaskSize),
		exitChan:   make(chan struct{}, 1),
	}
	context.tcpServer.GetConnManager().Add(context)
	return context
}

func (c *ConnContext) startReader() {
	fmt.Println("[ConnContext] reader is starting ...")
	defer fmt.Println(c.RemoteAddr().String(), " conn reader exit.")
	defer c.Stop()

	for {
		pack := &Package{}
		data := make([]byte, pack.MaxHead())
		if _, err := io.ReadFull(c.Conn, data); err != nil {
			fmt.Println("read conn head data err:", err)
			return
		}

		msg, err := pack.Unpack(data)
		if err != nil {
			fmt.Println("unpack err:", err)
			return
		}

		var values []byte
		if size := msg.GetDataSize(); size > 0 {
			values = make([]byte, size)
			if _, err := io.ReadFull(c.Conn, values); err != nil {
				fmt.Println("read conn data err:", err)
				return
			}
		}

		msg.SetData(values)
		req := &Req{conn: c, msg: msg}
		if config.GlobalConfig.WorkerPoolSize > 0 {
			c.msgHandle.AsyncTaskQueue(req)
		} else {
			go c.msgHandle.Do(req)
		}
	}
}

func (c *ConnContext) startWriter() {
	defer fmt.Println(c.RemoteAddr().String(), " conn write exit.")
	for {
		select {
		case msg, ok := <-c.msgChan:
			if ok {
				if _, err := c.Conn.Write(msg); err != nil {
					fmt.Println("send message err:", err)
					return
				}
				break
			}
			return
		case <-c.exitChan:
			return
		}
	}
}

func (c *ConnContext) Send(msgID uint32, data []byte) error {
	if !c.isClosed {
		pack := &Package{}
		if msg, err := pack.Pack(NewMsgPackage(msgID, data)); err != nil {
			fmt.Printf("package pack err id:%d reason:%v \n", msgID, err)
			return err
		} else {
			c.msgChan <- msg
		}
		return nil
	}
	return errors.New("conn closed when send message")
}



func (c *ConnContext) Start() {
	go c.startWriter()
	go c.startReader()
}

func (c *ConnContext) Stop() {
	if !c.isClosed {
		c.isClosed = true
		c.Conn.Close()
		c.exitChan <- struct{}{}
		c.tcpServer.GetConnManager().Del(c)
		close(c.exitChan)
		close(c.msgChan)
		close(c.msgBufChan)
	}
}

func (c *ConnContext) GetTcp() *net.TCPConn {
	return c.Conn
}

func (c *ConnContext) GetConnID() uint32 {
	return c.ID
}

func (c *ConnContext) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
