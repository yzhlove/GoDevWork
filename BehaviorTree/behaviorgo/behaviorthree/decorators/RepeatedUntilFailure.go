package decorators

import (
	b3 "behaviorthree"
	"behaviorthree/config"
	"behaviorthree/core"
)

type RepeatedUntilFailure struct {
	core.Decorator
	maxLoop int
}

func (this *RepeatedUntilFailure) Init(setting *config.BTNodeCfg) {
	this.Decorator.Init(setting)
	this.maxLoop = setting.GetPropertyAsInt("maxLoop")
	if this.maxLoop < 1 {
		panic("max loop must > 1")
	}
}

func (this *RepeatedUntilFailure) OnOpen(tick *core.Tick) {
	tick.Blackboard.Set("i", 0, tick.GetTree().GetID(), this.GetID())
}

func (this *RepeatedUntilFailure) OnTick(tick *core.Tick) b3.Status {
	if this.GetChild() == nil {
		return b3.ERROR
	}
	var i = tick.Blackboard.GetInt("i", tick.GetTree().GetID(), this.GetID())
	var status = b3.ERROR
	for this.maxLoop < 0 || i < this.maxLoop {
		status = this.GetChild().Execute(tick)
		if status == b3.SUCCESS {
			i++
		} else {
			break
		}
	}
	tick.Blackboard.Set("i", i, tick.GetTree().GetID(), this.GetID())
	return status
}
