package ziface

type MsgHandleImp interface {
	Do(req ReqImp)
	Register(msgID uint32, router RouterImp)
	RunWorkerPool()
	AsyncTaskQueue(req ReqImp)
}
