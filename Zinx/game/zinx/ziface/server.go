package ziface

type EventFunc func(conn ConnImp)

type EventImp interface {
	ConnStartEvent(event EventFunc)
	ConnStopEvent(event EventFunc)
	CallbackConnStart(conn ConnImp)
	CallbackConnStop(conn ConnImp)
}

type ServerImp interface {
	Start()
	Stop()
	Run()
	Register(msgID uint32, router RouterImp)
	GetConnMgr() ConnMgrImp
	GetMsgHandle() MsgHandleImp
	EventImp
}
