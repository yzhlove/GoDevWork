package ziface

type RequestInterface interface {
	GetConn() ConnectionInterface
	GetData() []byte
	GetMessageID() uint32
}
