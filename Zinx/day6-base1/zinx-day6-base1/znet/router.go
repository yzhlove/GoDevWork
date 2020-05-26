package znet

import "zinx-day6-base1/ziface"

type BaseRouter struct{}

func (BaseRouter) BeforeHandle(_ ziface.RequestInterface) {}
func (BaseRouter) Handle(_ ziface.RequestInterface)       {}
func (BaseRouter) AfterHandle(_ ziface.RequestInterface)  {}
