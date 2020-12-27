package context

import (
	"WorkSpace/GoDevWork/Events/chat02/data"
	"WorkSpace/GoDevWork/Events/chat02/handle"
)

type Context struct {
	Achieve data.Achieve
	*handle.EventContext
}

func Construct() *Context {
	return nil
}
