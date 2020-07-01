package bt

import (
	"behavior_day1_base1/xml"
	"fmt"
)

var GlobalTreeMap map[string]NodeInterface

func init() {
	GlobalTreeMap = make(map[string]NodeInterface)
}

type Node struct {
	BaseAction
	Name     string
	Type     NODE_TYPE
	State    NODE_STATE
	Idx      int
	Parent   NodeInterface
	Children []NodeInterface
}

func (n *Node) AddChild(node NodeInterface) {
	n.Children = append(n.Children, node)
}

func (n *Node) AddParent(node NodeInterface) {
	n.Parent = node
}

func (n *Node) GetName() string {
	return n.Name
}

func (n *Node) GetType() NODE_TYPE {
	return n.Type
}

func (n *Node) GetState() NODE_STATE {
	return n.State
}

func (n *Node) GetParent() NodeInterface {
	return n.Parent
}

func (n *Node) GetChildren() []NodeInterface {
	return n.Children
}

func (n *Node) Reset() {
	n.State = NODE_STATE_INVALID
}

type Root struct {
	Node
	CurIdx int
}

func NewRoot(xnode *xml.Node) NodeInterface {
	return &Root{
		Node: Node{
			Name:  xnode.Name,
			Type:  NODE_TYPE_ROOT,
			State: NODE_STATE_INVALID,
			Idx:   xnode.Height,
		},
	}
}

func (r *Root) OnEnter(agent AgentInterface) {
	fmt.Println("onEnter <root> ", r.Idx)
	if r.GetState() == NODE_STATE_INVALID {
		r.State = NODE_STATE_RUNNING
	}
}

func (r *Root) Exec(dt float64, agent AgentInterface) {
	if r.CurIdx == len(r.Children) {
		r.Reset()
	}
	child := r.Children[r.CurIdx]
	if child.GetState() != NODE_STATE_SUCCESS && child.GetState() != NODE_STATE_FAILED {
		child.TickUpdate(child, WithOpt(dt, agent))
	}
	if child.GetState() != NODE_STATE_RUNNING {
		r.CurIdx++
	}
}

func (r *Root) OnExit(agent AgentInterface) {
	fmt.Println("onExit root -> ", r.Idx, r.State)
}

func (r *Root) Reset() {
	r.CurIdx = 0
	for _, child := range r.Children {
		child.Reset()
	}
}
