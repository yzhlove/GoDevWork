package ziface

type ConnMgrImp interface {
	Add(conn ConnImp)
	Get(connID uint32) (ConnImp, bool)
	Del(connID ConnImp)
	Size() uint32
	Clear()
}
