package core

import (
	b3 "behaviorthree"
	"behaviorthree/config"
)

type ICondition interface {
	IBaseNode
}

type Condition struct {
	BaseNode
	BaseWorker
}

func (this *Condition) Ctor() {
	this.category = b3.CONDITION
}

func (this *Condition) Init(params *config.BTNodeCfg) {
	this.BaseNode.Init(params)
}
