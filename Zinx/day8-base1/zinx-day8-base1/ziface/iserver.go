package ziface

type CallbackConnFunc func(conn ConnectionInterface)

type CallbackInterface interface {
	SetOnConnStart(fn CallbackConnFunc)
	SetOnConnStop(fn CallbackConnFunc)
	CallOnConnStart(conn ConnectionInterface)
	CallOnConnStop(conn ConnectionInterface)
}

type ServerInterface interface {
	Start()
	Stop()
	Run()
	RegisterRouter(msgID uint32, router RouterInterface)
	GetConnManager() ConnManagerInterface
	GetMsgHandle() MsgHandleInterface
	CallbackInterface
}
