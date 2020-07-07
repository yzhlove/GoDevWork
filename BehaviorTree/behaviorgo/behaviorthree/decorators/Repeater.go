package decorators

import (
	b3 "behaviorthree"
	"behaviorthree/config"
	"behaviorthree/core"
)

type Repeater struct {
	core.Decorator
	maxLoop int
}

func (this *Repeater) Init(setting *config.BTNodeCfg) {
	this.Decorator.Init(setting)
	this.maxLoop = setting.GetPropertyAsInt("maxLoop")
	if this.maxLoop < 1 {
		panic("max loop must > 0")
	}
}

func (this *Repeater) OnOpen(tick *core.Tick) {
	tick.Blackboard.Set("i", 0, tick.GetTree().GetID(), this.GetID())
}

func (this *Repeater) OnTick(tick *core.Tick) b3.Status {
	if this.GetChild() != nil {
		var i = tick.Blackboard.GetInt("i", tick.GetTree().GetID(), this.GetID())
		var status = b3.SUCCESS

		for this.maxLoop < 0 || i < this.maxLoop {
			status = this.GetChild().Execute(tick)
			if status == b3.SUCCESS || status == b3.FAILURE {
				i++
			} else {
				break
			}
		}
		tick.Blackboard.Set("i", i, tick.GetTree().GetID(), this.GetID())
	}
	return b3.ERROR
}
