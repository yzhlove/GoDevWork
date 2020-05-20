package ziface

type ServerInterface interface {
	Start()
	Stop()
	Serve()
}
