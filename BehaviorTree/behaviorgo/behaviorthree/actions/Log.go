package actions

import (
	b3 "behaviorthree"
	"behaviorthree/config"
	"behaviorthree/core"
	"fmt"
)

type Log struct {
	core.Action
	info string
}

func (this *Log) Init(setting *config.BTNodeCfg) {
	this.Action.Init(setting)
	this.info = setting.GetPropertyAsString("info")
}

func (this *Log) OnTick(tick *core.Tick) b3.Status {
	fmt.Println("log:", this.info)
	return b3.SUCCESS
}
