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

	return 0
}

func (this *SubTree) String() string {
	return "SBT_" + this.title
}

var subTreeLoadFunc func(string) *BehaviorTree

func SetSubTreeLoadFunc(f func(string) *BehaviorTree) {
	subTreeLoadFunc = f
}
