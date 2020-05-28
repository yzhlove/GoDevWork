package ziface

type MsgHandleInterface interface {
	Do(request RequestInterface)
	RegisterRouter(msgID uint32, router RouterInterface)
	StartWorkerPool()
	AsyncTaskQueue(request RequestInterface)
}
