package ziface

import "net"

type ConnectionInterface interface {
	Start()
	Stop()
	GetTcp() *net.TCPConn
	GetConnID() uint32
	RemoteAddr() net.Addr
}

type HandleFunc func(*net.TCPConn, []byte, int) error
