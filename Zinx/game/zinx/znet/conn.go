package znet

import (
	"errors"
	"io"
	"net"
	"sync"
	"zinx/config"
	"zinx/ziface"
	"zinx/zlog"
)

type Conn struct {
	tcpServer   ziface.ServerImp
	ID          uint32
	Conn        *net.TCPConn
	isClosed    bool
	msgHandle   ziface.MsgHandleImp
	msgChan     chan []byte
	attrs       map[string]interface{}
	mutex       sync.RWMutex
	closedMutex sync.RWMutex
	exitChan    chan struct{}
}

func NewConn(c *net.TCPConn, id uint32, server ziface.ServerImp) *Conn {
	conn := &Conn{
		tcpServer: server,
		Conn:      c,
		ID:        id,
		msgHandle: server.GetMsgHandle(),
		msgChan:   make(chan []byte, config.GlobalConfig.MaxMsgChanSize),
		exitChan:  make(chan struct{}, 1),
		attrs:     make(map[string]interface{}),
	}
	conn.tcpServer.GetConnMgr().Add(conn)
	return conn
}

func (c *Conn) reader() {
	zlog.Info("reader is starting ...")
	defer zlog.Info(c.Conn.RemoteAddr().String(), " conn reader exit.")
	defer c.Stop()

	for {
		pack := NewPack()
		data := make([]byte, pack.HeadSize())
		if _, err := io.ReadFull(c.Conn, data); err != nil {
			zlog.Error("read conn head err:", err)
			return
		}
		msg, err := pack.Unpack(data)
		if err != nil {
			zlog.Error("pack unpack err:", err)
			return
		}
		var values []byte
		if size := msg.GetSize(); size > 0 {
			values = make([]byte, size)
			if _, err := io.ReadFull(c.Conn, values); err != nil {
				zlog.Error("read data err:", err)
				return
			}
		}
		msg.SetData(values)
		c.msgHandle.AsyncTaskQueue(&Req{conn: c, msg: msg})
	}
}

func (c *Conn) writer() {
	zlog.Info("writer is starting ...")
	defer zlog.Info(c.Conn.RemoteAddr().String(), " conn write exit.")
	for {
		select {
		case msg, ok := <-c.msgChan:
			if !ok {
				return
			}
			if _, err := c.Conn.Write(msg); err != nil {
				zlog.Error("conn write err:", err)
				return
			}
		case <-c.exitChan:
			return
		}
	}
}

func (c *Conn) Send(msgID uint32, data []byte) error {
	if !c.getClosed() {
		pack := NewPack()
		if msg, err := pack.Pack(NewMsg(msgID, data)); err != nil {
			zlog.Info("package pack err:", err)
			return err
		} else {
			select {
			case <-c.exitChan:
				return errors.New("conn is closed")
			default:
				c.msgChan <- msg
			}
		}
		return nil
	}
	return errors.New("conn closed when send message")
}

func (c *Conn) Start() {
	if !c.isClosed {
		go c.writer()
		go c.reader()
		c.tcpServer.CallbackConnStart(c)
	}
}

func (c *Conn) Stop() {
	if !c.getClosed() {
		c.setClosed(true)
		c.tcpServer.CallbackConnStop(c)
		c.Conn.Close()
		c.exitChan <- struct{}{}
		c.tcpServer.GetConnMgr().Del(c)
		close(c.exitChan)
		close(c.msgChan)
	}
}

func (c *Conn) setClosed(closed bool) {
	c.closedMutex.Lock()
	defer c.closedMutex.Unlock()
	c.isClosed = closed
}

func (c *Conn) getClosed() bool {
	c.closedMutex.RLock()
	defer c.closedMutex.RUnlock()
	return c.isClosed
}

func (c *Conn) GetTcp() *net.TCPConn {
	return c.Conn
}

func (c *Conn) GetConnID() uint32 {
	return c.ID
}

func (c *Conn) SetAttr(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.attrs[key] = value
}

func (c *Conn) GetAttr(key string) (value interface{}, ok bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	value, ok = c.attrs[key]
	return
}

func (c *Conn) DelAttr(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.attrs, key)
}
