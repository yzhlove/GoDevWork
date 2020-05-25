package ziface

type ServerInterface interface {
	Start()
	Stop()
	Run()
	RegisterRouter(msgID uint32, router RouterInterface)
}
