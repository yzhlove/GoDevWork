package actions

import (
	b3 "behaviorthree"
	"behaviorthree/core"
)

type Error struct {
	core.Action
}

func (this *Error) OnTick(tick *core.Tick) b3.Status {
	return b3.ERROR
}
