package bt

import (
	"behavior_day1_base1/xml"
	"fmt"
)

//顺序节点

type Sequence struct {
	Node
	CurIdx int
}

func NewSequence(xnode *xml.Node) NodeInterface {
	return &Sequence{
		Node: Node{
			Name:  xnode.Name,
			Type:  NODE_TYPE_SEQUENCE,
			State: NODE_STATE_INVALID,
			Idx:   xnode.Height,
		},
	}
}

func (s *Sequence) OnEnter(agent AgentInterface) {
	fmt.Println("onEnter <sequence> ", s.Idx)
	if s.GetState() == NODE_STATE_INVALID {
		s.State = NODE_STATE_RUNNING
	}
}

func (s *Sequence) Exec(dt float64, agent AgentInterface) {
	if len(s.Children) <= 0 {
		panic("sequence no children")
	}
	child := s.Children[s.CurIdx]
	if child.GetState() == NODE_STATE_INVALID || child.GetState() == NODE_STATE_RUNNING {
		child.TickUpdate(child, WithOpt(dt, agent))
	}
	switch state := child.GetState(); state {
	case NODE_STATE_RUNNING:
		s.State = NODE_STATE_RUNNING
		fmt.Println("--- Running Sequence ", s.State, s.CurIdx)
	case NODE_STATE_FAILED:
		s.State = NODE_STATE_FAILED
		fmt.Println("--- Failed Sequence ", s.State, s.CurIdx)
	case NODE_STATE_SUCCESS:
		s.CurIdx++
		if s.CurIdx >= len(s.Children) {
			s.State = NODE_STATE_SUCCESS
		}
		fmt.Println("--- Succeed Sequence ", s.State, s.CurIdx)
	}
}

func (s *Sequence) OnExit(agent AgentInterface) {
	fmt.Println("onExit <Sequence> ", s.Idx, s.State)
}

func (s *Sequence) Reset() {
	s.CurIdx = 0
	s.State = NODE_STATE_FAILED
	for _, child := range s.Children {
		child.Reset()
	}
}
