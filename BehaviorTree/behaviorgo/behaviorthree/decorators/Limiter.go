package decorators

import (
	b3 "behaviorthree"
	"behaviorthree/config"
	"behaviorthree/core"
)

type Limiter struct {
	core.Decorator
	maxLoop int
}

func (this *Limiter) Init(setting *config.BTNodeCfg) {
	this.Decorator.Init(setting)
	this.maxLoop = setting.GetPropertyAsInt("maxLoop")
	if this.maxLoop < 1 {
		panic("max loop must > 1")
	}
}

func (this *Limiter) OnTick(tick *core.Tick) b3.Status {
	if this.GetChild() != nil {
		var i = tick.Blackboard.GetInt("i", tick.GetTree().GetID(), this.GetID())
		if i < this.maxLoop {
			var status = this.GetChild().Execute(tick)
			if status == b3.SUCCESS || status == b3.FAILURE {
				tick.Blackboard.Set("i", i+1, tick.GetTree().GetID(), this.GetID())
			}
			return status
		}
	}
	return b3.FAILURE
}
