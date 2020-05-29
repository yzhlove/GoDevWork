package znet

import (
	"sync"
	"zinx/ziface"
)

type ConnMgr struct {
	cons  map[uint32]ziface.ConnImp
	mutex sync.RWMutex
}

func NewConnMgr() *ConnMgr {
	return &ConnMgr{
		cons: make(map[uint32]ziface.ConnImp, 128),
	}
}

func (c *ConnMgr) Add(conn ziface.ConnImp) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cons[conn.GetConnID()] = conn
}

func (c *ConnMgr) Get(id uint32) (ziface.ConnImp, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	conn, ok := c.cons[id]
	return conn, ok
}

func (c *ConnMgr) Del(conn ziface.ConnImp) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	conn.Stop()
	delete(c.cons, conn.GetConnID())
}

func (c *ConnMgr) Size() uint32 {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return uint32(len(c.cons))
}

func (c *ConnMgr) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for id, conn := range c.cons {
		conn.Stop()
		delete(c.cons, id)
	}
}
