package decorators

import (
	b3 "behaviorthree"
	"behaviorthree/core"
)

type Inverter struct {
	core.Decorator
}

func (this *Inverter) OnTick(tick *core.Tick) b3.Status {
	if this.GetChild() != nil {
		var status = this.GetChild().Execute(tick)
		if status == b3.SUCCESS {
			status = b3.FAILURE
		} else if status == b3.FAILURE {
			status = b3.SUCCESS
		}
		return status
	}
	return b3.ERROR
}
