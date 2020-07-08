package composites

import (
	b3 "behaviorthree"
	"behaviorthree/core"
)

type Parallel struct {
	core.Composite
}

func (this *Parallel) OnTick(tick *core.Tick) b3.Status {
	if count := this.GetChildCount(); count > 0 {
		var index int
		for i := 0; i < count; i++ {
			if this.GetChild(i).Execute(tick) == b3.SUCCESS {
				index++
			}
		}
		if index == count {
			return b3.SUCCESS
		}
	}
	return b3.FAILURE
}
