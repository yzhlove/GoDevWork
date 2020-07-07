package core

import (
	b3 "behaviorthree"
	"behaviorthree/config"
)

type IDecorator interface {
	IBaseNode
	SetChild(child IBaseNode)
	GetChild() IBaseNode
}

type Decorator struct {
	BaseNode
	BaseWorker
	child IBaseNode
}

func (this *Decorator) Ctor() {
	this.category = b3.DECORATOR
}

func (this *Decorator) Init(params *config.BTNodeCfg) {
	this.BaseNode.Init(params)
}

func (this *Decorator) GetChild() IBaseNode {
	return this.child
}

func (this *Decorator) SetChild(child IBaseNode) {
	this.child = child
}
