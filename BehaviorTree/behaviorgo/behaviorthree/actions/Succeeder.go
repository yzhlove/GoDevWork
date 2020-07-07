package actions

import (
	b3 "behaviorthree"
	"behaviorthree/core"
)

type Succeed struct {
	core.Action
}

func (this *Succeed) OnTick(tick *core.Tick) b3.Status {
	return b3.SUCCESS
}
