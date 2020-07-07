package share

import (
	b3 "behaviorthree"
	"behaviorthree/config"
	"behaviorthree/core"
)

type SetValue struct {
	core.Action
	value int
	key   string
}

func (this *SetValue) Init(setting *config.BTNodeCfg) {
	this.Action.Init(setting)
	this.value = setting.GetPropertyAsInt("value")
	this.key = setting.GetPropertyAsString("key")
}

func (this *SetValue) OnTick(tick *core.Tick) b3.Status {
	tick.Blackboard.SetMem(this.key, this.value)
	return b3.SUCCESS
}

type IsValue struct {
	core.Condition
	value int
	key   string
}

func (this *IsValue) Init(setting *config.BTNodeCfg) {
	this.Condition.Init(setting)
	this.value = setting.GetPropertyAsInt("value")
	this.key = setting.GetPropertyAsString("key")
}

func (this *IsValue) OnTick(tick *core.Tick) b3.Status {
	v := tick.Blackboard.GetInt(this.key, "", "")
	if v != this.value {
		return b3.FAILURE
	}
	return b3.SUCCESS
}
