package actions

import (
	b3 "behaviorthree"
	"behaviorthree/config"
	"behaviorthree/core"
	"time"
)

type Wait struct {
	core.Action
	endTime int64
}

func (this *Wait) Init(setting *config.BTNodeCfg) {
	this.Action.Init(setting)
	this.endTime = setting.GetPropertyAsInt64("milliseconds")
}

func (this *Wait) OnOpen(tick *core.Tick) {
	var start = time.Now().UnixNano() / 1e6
	tick.Blackboard.Set("startTime", start, tick.GetTree().GetID(), this.GetID())
}

func (this *Wait) OnTick(tick *core.Tick) b3.Status {
	var curTime = time.Now().UnixNano() / 1e6
	var startTime = tick.Blackboard.GetInt64("startTime", tick.GetTree().GetID(), this.GetID())
	if curTime-startTime > this.endTime {
		return b3.SUCCESS
	}
	return b3.RUNNING
}
