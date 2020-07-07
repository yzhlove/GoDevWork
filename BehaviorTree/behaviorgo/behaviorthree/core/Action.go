package core

import (
	b3 "behaviorthree"
	"behaviorthree/config"
)

type IAction interface {
	IBaseNode
}

type Action struct {
	BaseNode
	BaseWorker
}

func (this *Action) Ctor() {
	this.category = b3.ACTION
}

func (this *Action) Init(params *config.BTNodeCfg) {
	this.BaseNode.Init(params)
}
