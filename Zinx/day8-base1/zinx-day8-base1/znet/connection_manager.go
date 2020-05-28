package znet

import (
	"fmt"
	"sync"
	"zinx-day8-base1/ziface"
)

type ConnManager struct {
	connections map[uint32]ziface.ConnectionInterface
	mutex       sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.ConnectionInterface, 128),
	}
}

func (c *ConnManager) Add(conn ziface.ConnectionInterface) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.connections[conn.GetConnID()] = conn
	fmt.Println("conn add to manager succeed:conn num ", c.Len())
}

func (c *ConnManager) Del(conn ziface.ConnectionInterface) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.connections, conn.GetConnID())
	fmt.Println("conn del to manager succeed: id ", conn.GetConnID(), " num ", c.Len())
}

func (c *ConnManager) Get(connID uint32) (ziface.ConnectionInterface, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if conn, ok := c.connections[connID]; ok {
		return conn, true
	}
	return nil, false
}

func (c *ConnManager) Len() int {
	return len(c.connections)
}

func (c *ConnManager) ClearConn() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for cid, conn := range c.connections {
		conn.Stop()
		delete(c.connections, cid)
	}
	fmt.Println("clear all connections succeed num", c.Len())
}
