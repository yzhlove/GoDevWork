package ziface

type ConnManagerInterface interface {
	Add(conn ConnectionInterface)
	Del(conn ConnectionInterface)
	Get(connID uint32) (ConnectionInterface, bool)
	Len() int
	ClearConn()
}
