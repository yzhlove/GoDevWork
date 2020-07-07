package composites

import (
	b3 "behaviorthree"
	"behaviorthree/core"
)

type MemPriority struct {
	core.Composite
}

func (this *MemPriority) OnOpen(tick *core.Tick) {
	tick.Blackboard.Set("runningChild", 0, tick.GetTree().GetID(), this.GetID())
}

func (this *MemPriority) OnTick(tick *core.Tick) b3.Status {
	var child = tick.Blackboard.GetInt("runningChild", tick.GetTree().GetID(), this.GetID())
	for i := child; i < this.GetChildCount(); i++ {
		var status = this.GetChild(i).Execute(tick)
		if status != b3.FAILURE {
			if status == b3.RUNNING {
				tick.Blackboard.Set("runningChild", i, tick.GetTree().GetID(), this.GetID())
			}
			return status
		}
	}
	return b3.FAILURE
}
