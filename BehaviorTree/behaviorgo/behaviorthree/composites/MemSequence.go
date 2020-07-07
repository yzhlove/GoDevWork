package composites

import (
	b3 "behaviorthree"
	"behaviorthree/core"
	"fmt"
)

type MemSequence struct {
	core.Composite
}

func (this *MemSequence) OnOpen(tick *core.Tick) {
	fmt.Println("------------ MemSequence OnOpen ------------")
	tick.Blackboard.Set("runningChild", 0, tick.GetTree().GetID(), this.GetID())
}

func (this *MemSequence) OnTick(tick *core.Tick) b3.Status {
	var child = tick.Blackboard.GetInt("runningChild", tick.GetTree().GetID(), this.GetID())
	for i := child; i < this.GetChildCount(); i++ {
		var status = this.GetChild(i).Execute(tick)
		if status != b3.SUCCESS {
			if status == b3.RUNNING {
				tick.Blackboard.Set("runningChild", i, tick.GetTree().GetID(), this.GetID())
			}
			return status
		}
	}
	return b3.SUCCESS
}
