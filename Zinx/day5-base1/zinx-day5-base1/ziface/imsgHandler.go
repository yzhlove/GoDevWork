package ziface

type MessageHandleInterface interface {
	DoMessageHandle(request RequestInterface)
	RegisterRouter(msgID uint32, router RouterInterface)
}
