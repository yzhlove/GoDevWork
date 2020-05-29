package ziface

type ReqImp interface {
	GetConn() ConnImp
	GetMsgData() []byte
	GetMsgID() uint32
}
