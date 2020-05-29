package ziface

import "net"

type AttributeImp interface {
	SetAttr(key string, value interface{})
	GetAttr(key string) (interface{}, bool)
	DelAttr(key string)
}

type ConnImp interface {
	Start()
	Stop()
	GetTcp() *net.TCPConn
	GetConnID() uint32
	Send(msgID uint32, data []byte) error
	AttributeImp
}
