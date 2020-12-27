package bind

import (
	"WorkSpace/GoDevWork/Events/chat02/cond"
	"WorkSpace/GoDevWork/Events/chat02/context"
	"WorkSpace/GoDevWork/Events/chat02/event"
)

func UpdateAchieve(ctx *context.Context, evt event.Event) error {

	for _, record := range ctx.Achieve.All() {
		if len(record.Conditions) == 0 {
			continue
		}
		for _, c := range record.Conditions {
			if creator, ok := cond.Get(c.Name); ok {
				cd := creator()
				cd.SetTarget(c.Params)
				cd.SetCancel(ctx.AddListener(cd.Event(), func(ctx *context.Context, evt event.Event) error {
					if cd.Ok() {
						cd.Cancel()
					}
					return nil
				}))
			}
		}
	}
	return nil
}
