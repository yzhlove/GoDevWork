package composites

import (
	b3 "behaviorthree"
	"behaviorthree/core"
)

type Sequence struct {
	core.Composite
}

func (this *Sequence) OnTick(tick *core.Tick) b3.Status {
	for i := 0; i < this.GetChildCount(); i++ {
		var status = this.GetChild(i).Execute(tick)
		if status != b3.SUCCESS {
			return status
		}
	}
	return b3.SUCCESS
}
