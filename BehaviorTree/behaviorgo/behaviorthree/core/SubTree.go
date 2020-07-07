package core

import (
	b3 "behaviorthree"
	"behaviorthree/config"
)

type SubTree struct {
	Action
}

func (this *SubTree) Init(setting *config.BTNodeCfg) {
	this.Action.Init(setting)
}

func (this *SubTree) OnTick(tick *Tick) b3.Status {
	stree := subTreeLoadFunc(this.GetName())
	if stree == nil {
		return b3.ERROR
	}
	if tick.GetTarget() == nil {
		panic("sub stree tick get target nil")
	}
	tick.pushSubNode(this)
	ret := stree.GetRoot().Execute(tick)
	tick.popSubNode()
	return ret
}

func (this *SubTree) String() string {
	return "SBT_" + this.title
}

var subTreeLoadFunc func(string) *BehaviorTree

func SetSubTreeLoadFunc(f func(string) *BehaviorTree) {
	subTreeLoadFunc = f
}
