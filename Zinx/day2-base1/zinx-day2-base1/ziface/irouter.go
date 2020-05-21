package ziface

type RouterInterface interface {
	BeforeHandle(RequestInterface)
	Handle(RequestInterface)
	AfterHandle(RequestInterface)
}
