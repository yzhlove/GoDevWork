package znet

import "zinx-day2-base1/ziface"

type BaseRouter struct{}

func (*BaseRouter) BeforeHandle(req ziface.RequestInterface) {}
func (*BaseRouter) Handle(req ziface.RequestInterface)       {}
func (*BaseRouter) AfterHandle(req ziface.RequestInterface)  {}


