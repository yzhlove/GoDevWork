package ziface

import "net"

type ConnectionInterface interface {
	Start()
	Stop()
	GetTcp() *net.TCPConn
	GetConnID() uint32
	RemoteAddr() net.Addr
	Send(msgID uint32, data []byte) error
}
