package core

import (
	b3 "behaviorthree"
	"fmt"
)

type IBaseWorker interface {
	OnEnter(tick *Tick)
	OnOpen(tick *Tick)
	OnTick(tick *Tick) b3.Status
	OnClose(tick *Tick)
	OnExit(tick *Tick)
}

type BaseWorker struct{}

func (b BaseWorker) OnEnter(tick *Tick) {
}

func (b BaseWorker) OnOpen(tick *Tick) {
}

func (b BaseWorker) OnTick(tick *Tick) b3.Status {
	fmt.Println("tick BaseWorker")
	return b3.ERROR
}

func (b BaseWorker) OnClose(tick *Tick) {
}

func (b BaseWorker) OnExit(tick *Tick) {
}
