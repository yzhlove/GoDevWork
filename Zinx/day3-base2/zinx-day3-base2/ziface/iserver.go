package ziface

type ServerInterface interface {
	Start()
	Stop()
	Run()
	RegisterRouter(router RouterInterface)
}
