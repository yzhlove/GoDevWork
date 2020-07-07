package composites

import (
	b3 "behaviorthree"
	"behaviorthree/core"
)

type Priority struct {
	core.Composite
}

func (this *Priority) OnTick(tick *core.Tick) b3.Status {
	for i := 0; i < this.GetChildCount(); i++ {
		var status = this.GetChild(i).Execute(tick)
		if status != b3.FAILURE {
			return status
		}
	}
	return b3.FAILURE
}
