package core

import (
	b3 "behaviorthree"
	"behaviorthree/config"
	"fmt"
)

type IComposite interface {
	IBaseNode
	GetChildCount() int
	GetChild(index int) IBaseNode
	AddChild(child IBaseNode)
}

type Composite struct {
	BaseNode
	BaseWorker

	children []IBaseNode
}

func (this *Composite) Ctor() {
	this.category = b3.COMPOSITE
}

func (this *Composite) Init(params *config.BTNodeCfg) {
	this.BaseNode.Init(params)
}

func (this *Composite) GetChildCount() int {
	return len(this.children)
}

func (this *Composite) GetChild(index int) IBaseNode {
	return this.children[index]
}

func (this *Composite) AddChild(child IBaseNode) {
	this.children = append(this.children, child)
}

func (this *Composite) tick(tick *Tick) b3.Status {
	fmt.Println("tick composite")
	return b3.ERROR
}
