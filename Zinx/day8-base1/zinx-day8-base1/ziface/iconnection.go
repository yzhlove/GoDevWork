package ziface

import "net"

type ConnectionInterface interface {
	Start()
	Stop()
	GetTcp() *net.TCPConn
	GetConnID() uint32
	RemoteAddr() net.Addr
	Send(msgID uint32, data []byte) error
	SendBuf(msgID uint32, data []byte) error
	SetAttribute(key string, value interface{})
	GetAttribute(key string) (interface{}, bool)
	DelAttribute(key string)
}
