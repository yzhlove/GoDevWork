package cond

import (
	"WorkSpace/GoDevWork/Events/chat02/event"
	"fmt"
)

type Login struct {
	base
}

func (l Login) Event() event.Event {
	return event.Login{}
}

func (l Login) Ok() bool {
	fmt.Println("login ==> ", l.target)
	return true
}
