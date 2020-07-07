package share

import (
	b3 "behaviorthree"
	"behaviorthree/config"
	"behaviorthree/core"
	"fmt"
)

type LogTest struct {
	core.Action
	info string
}

func (this *LogTest) Init(setting *config.BTNodeCfg) {
	this.Action.Init(setting)
	this.info = setting.GetPropertyAsString("info")
}

func (this *LogTest) OnTick(tick *core.Tick) b3.Status {
	fmt.Println("[logTest]:", tick.GetLastSub(), this.info)
	return b3.SUCCESS
}
