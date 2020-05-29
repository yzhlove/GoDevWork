package ziface

type RouterImp interface {
	BeforeDo(req ReqImp)
	Handle(req ReqImp)
	AfterDo(req ReqImp)
}
