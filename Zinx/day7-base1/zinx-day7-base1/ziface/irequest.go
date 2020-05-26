package ziface

type RequestInterface interface {
	GetConn() ConnectionInterface
	GetData() []byte
	GetMsgID() uint32
}
