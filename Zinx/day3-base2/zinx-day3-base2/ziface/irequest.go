package ziface

type RequestInterface interface {
	GetConn() ConnectionInterface
	GetData() []byte
}
