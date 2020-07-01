package main

import "fmt"

func main() {

	c := &Child{}
	c.Tick(12)
	c.TickUpdate(14, c)
}

type RunInterface interface {
	Tick(dt float64)
}

type NodeInterface interface {
	OnEnter()
	Exec(dt float64)
	OnExit()
}

type Base struct{}

func (b *Base) Tick(dt float64) {
	b.OnEnter()
	b.Exec(dt)
	b.OnExit()
}

func (b *Base) TickUpdate(dt float64, n NodeInterface) {
	n.OnEnter()
	n.Exec(dt)
	n.OnExit()
}

func (b *Base) OnEnter() {
	fmt.Println("b onenter")
}

func (b *Base) Exec(dt float64) {
	fmt.Println("b exec")
}

func (b *Base) OnExit() {
	fmt.Println("b onexit")
}

type Child struct {
	Base
}

func (c *Child) OnEnter() {
	fmt.Println("c onenter")
}

func (c *Child) Exec(dt float64) {
	fmt.Println("c exec")
}

func (c *Child) OnExit() {
	fmt.Println("c exit")
}
