package actions

import (
	b3 "behaviorthree"
	"behaviorthree/core"
)

type Runner struct {
	core.Action
}

func (this *Runner) OnTick(tick *core.Tick) b3.Status {
	return b3.RUNNING
}
