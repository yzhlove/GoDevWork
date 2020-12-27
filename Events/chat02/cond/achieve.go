package cond

import (
	"WorkSpace/GoDevWork/Events/chat02/event"
	"fmt"
)

type Achieve struct {
	base
}

func (Achieve) Event() event.Event {
	return event.Achieve{}
}

func (a Achieve) Ok() bool {
	fmt.Println("achieve ==> ", a.target)
	return true
}
