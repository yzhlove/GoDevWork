package bt

import (
	"behavior_day1_base1/xml"
	"fmt"
)

//选择节点

type Selector struct {
	Node
	CurIdx int
}

func NewSelector(xnode *xml.Node) NodeInterface {
	return &Selector{
		Node: Node{
			Name:  xnode.Name,
			Type:  NODE_TYPE_SELECTOR,
			State: NODE_STATE_INVALID,
			Idx:   xnode.Height,
		},
	}
}

func (s *Selector) OnEnter(agent AgentInterface) {
	fmt.Println("onEnter <selector> ", s.Idx)
	if s.GetState() == NODE_STATE_INVALID {
		s.State = NODE_STATE_RUNNING
	}
}

func (s *Selector) Exec(dt float64, agent AgentInterface) {
	if len(s.Children) <= 0 {
		panic("selector no children")
	}
	child := s.Children[s.CurIdx]
	if child.GetState() == NODE_STATE_INVALID || child.GetState() == NODE_STATE_RUNNING {
		child.TickUpdate(child, WithOpt(dt, agent))
	}
	switch state := child.GetState(); state {
	case NODE_STATE_RUNNING:
		s.State = NODE_STATE_RUNNING
		fmt.Println("---Running selector ", state, s.CurIdx)
	case NODE_STATE_FAILED:
		s.CurIdx++
		if s.CurIdx >= len(s.Children) {
			s.State = NODE_STATE_FAILED
		}
		fmt.Println("---Failed selector ", state, s.CurIdx)
	case NODE_STATE_SUCCESS:
		s.State = NODE_STATE_SUCCESS
		fmt.Println("--Succeed selector ", state, s.CurIdx)
	}
}

func (s *Selector) OnExit(agent AgentInterface) {
	fmt.Println("onExit selector ", s.Idx, s.State)
}

func (s *Selector) Reset() {
	s.CurIdx = 0
	s.State = NODE_STATE_INVALID
	for _, child := range s.Children {
		child.Reset()
	}
}
