package decorators

import (
	b3 "behaviorthree"
	"behaviorthree/config"
	"behaviorthree/core"
	"time"
)

type MaxTime struct {
	core.Decorator
	maxTime int64
}

func (this *MaxTime) Init(setting *config.BTNodeCfg) {
	this.Decorator.Init(setting)
	this.maxTime = setting.GetPropertyAsInt64("maxTime")
	if this.maxTime < 1 {
		panic("must time type")
	}
}

func (this *MaxTime) OnOpen(tick *core.Tick) {
	var startTime = time.Now().UnixNano() / 1e6
	tick.Blackboard.Set("startTime", startTime, tick.GetTree().GetID(), this.GetID())
}

func (this *MaxTime) OnTick(tick *core.Tick) b3.Status {
	if this.GetChild() != nil {
		var curTime = time.Now().UnixNano() / 1e6
		var startTime = tick.Blackboard.GetInt64("startTime", tick.GetTree().GetID(), this.GetID())
		var status = this.GetChild().Execute(tick)
		if curTime-startTime > this.maxTime {
			return b3.FAILURE
		}
		return status
	}
	return b3.ERROR
}
