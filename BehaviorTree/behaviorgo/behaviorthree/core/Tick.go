package core

import "fmt"

type Tick struct {
	tree          *BehaviorTree
	debug         interface{}
	target        interface{}
	Blackboard    *Blackboard
	_openNodes    []IBaseNode
	_openSubNodes []*SubTree
	_nodeCount    int
}

func NewTick() *Tick {
	t := &Tick{}
	t.Init()
	return t
}

func (this *Tick) Init() {
	this.tree = nil
	this.debug = nil
	this.target = nil
	this.Blackboard = nil

	this._openNodes = nil
	this._openSubNodes = nil
	this._nodeCount = 0
}

func (this *Tick) GetTree() *BehaviorTree {
	return this.tree
}

func (this *Tick) _enterNode(node IBaseNode) {
	this._nodeCount++
	this._openNodes = append(this._openNodes, node)
}

func (this *Tick) _openNode(node *BaseNode) {}

func (this *Tick) _tickNode(node *BaseNode) {
	fmt.Println("Tick _tickNode:", this.debug, "id:", node.GetID(), node.GetTitle(), node.GetName())
}

func (this *Tick) _closeNode(node *BaseNode) {
	if l := len(this._openNodes); l > 0 {
		this._openNodes = this._openNodes[:l-1]
	}
}

func (this *Tick) pushSubNode(node *SubTree) {
	this._openSubNodes = append(this._openSubNodes, node)
}

func (this *Tick) popSubNode() {
	if l := len(this._openSubNodes); l > 0 {
		this._openSubNodes = this._openSubNodes[:l-1]
	}
}

func (this *Tick) GetLastSub() *SubTree {
	if l := len(this._openSubNodes); l > 0 {
		return this._openSubNodes[l-1]
	}
	return nil
}

func (this *Tick) _exitNode(node *BaseNode) {}

func (this *Tick) GetTarget() interface{} {
	return this.target
}
