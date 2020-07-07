package actions

import (
	b3 "behaviorthree"
	"behaviorthree/core"
)

type Failer struct {
	core.Action
}

func (this *Failer) OnTick(tick *core.Tick) b3.Status {
	return b3.FAILURE
}
