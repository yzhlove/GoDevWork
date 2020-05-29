package znet

import "zinx/ziface"

type AbstractRouter struct{}

func (AbstractRouter) BeforeDo(imp ziface.ReqImp) {}
func (AbstractRouter) Handle(imp ziface.ReqImp)   {}
func (AbstractRouter) AfterDo(imp ziface.ReqImp)  {}
