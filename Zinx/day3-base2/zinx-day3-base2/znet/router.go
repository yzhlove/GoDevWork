package znet

import "zinx-day3-base2/ziface"

type BaseRouter struct{}

func (b BaseRouter) BeforeHandle(request ziface.RequestInterface) {

}

func (b BaseRouter) Handle(request ziface.RequestInterface) {

}

func (b BaseRouter) AfterHandle(request ziface.RequestInterface) {
	
}
