package bt

import (
	"behavior_day1_base1/xml"
	"fmt"
)

//并行节点

type Parallel struct {
	Node
}

func NewParallel(xnode *xml.Node) NodeInterface {
	return &Parallel{
		Node: Node{
			Name:  xnode.Name,
			Type:  NODE_TYPE_PARALLEL,
			State: NODE_STATE_INVALID,
			Idx:   xnode.Height,
		},
	}
}

func (p *Parallel) OnEnter(agent AgentInterface) {
	fmt.Println("onEnter <Parallel> ", p.State, p.Idx)
	if p.GetState() == NODE_STATE_INVALID {
		p.State = NODE_STATE_RUNNING
	}
}

func (p *Parallel) Exec(dt float64, agent AgentInterface) {
	if len(p.Children) <= 0 {
		panic("parallel no children")
	}
	var succeed, failed, count int
	for _, node := range p.Children {
		if node.GetState() == NODE_STATE_INVALID || node.GetState() == NODE_STATE_RUNNING {
			node.TickUpdate(node, WithOpt(dt, agent))
			if node.GetState() == NODE_STATE_SUCCESS {
				succeed++
			}
			if node.GetState() == NODE_STATE_FAILED {
				failed++
			}
			count++
		}
	}
	if succeed == len(p.Children) {
		p.State = NODE_STATE_SUCCESS
	}
	if count == len(p.Children) && failed > 0 {
		p.State = NODE_STATE_FAILED
	}
}

func (p *Parallel) OnExit(agent AgentInterface) {
	fmt.Println("onExit parallel ", p.Idx, p.State)
}

func (p *Parallel) Reset() {
	p.State = NODE_STATE_INVALID
	for _, child := range p.Children {
		child.Reset()
	}
}
