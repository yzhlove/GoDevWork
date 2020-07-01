package bt

import (
	"behavior_day1_base1/xml"
	"fmt"
)

type SkillAction struct {
	Node
	Info SkillInfo
}

func NewSkillAction(xnode *xml.Node) (NodeInterface, error) {
	node := &SkillAction{
		Node: Node{
			Name:  xnode.Name,
			Type:  NODE_TYPE_ACTION,
			State: NODE_STATE_INVALID,
			Idx:   xnode.Height,
		}}
	if err := node.Info.Parse(xnode.Params); err != nil {
		return nil, err
	}
	return node, nil
}

func (s *SkillAction) OnEnter(agent AgentInterface) {
	fmt.Println("onEnter <skill action> ", s.State, s.Idx)
	if s.GetState() == NODE_STATE_INVALID {
		s.State = NODE_STATE_RUNNING
	}
}

func (s *SkillAction) Exec(dt float64, agent AgentInterface) {
	fmt.Println("exec skill action", s.State, s.Idx)
	agent.CastSkill(s.Info.Sid, s.Info.Mp)
	if agent.IsSkillEnd() {
		s.State = NODE_STATE_SUCCESS
	} else {
		s.State = NODE_STATE_RUNNING
	}
}

func (s *SkillAction) OnExit(agent AgentInterface) {
	fmt.Println("onexit skill action ", s.Idx, s.State)
}

////////////////////////////////////////////////////////////////////////

type EscapeAction struct {
	Node
	Info EscapeInfo
}

func NewEscapeAction(xnode *xml.Node) NodeInterface {
	node := &EscapeAction{
		Node: Node{
			Name:  xnode.Name,
			Type:  NODE_TYPE_ACTION,
			State: NODE_STATE_INVALID,
			Idx:   xnode.Height,
		}}
	return node
}

func (e *EscapeAction) OnEnter(agent AgentInterface) {
	fmt.Println("onEnter <escape action>", e.State, e.Idx)
	if e.GetState() == NODE_STATE_INVALID {
		e.State = NODE_STATE_RUNNING
	}
}

func (e *EscapeAction) Exec(dt float64, agent AgentInterface) {
	fmt.Println("exec escape ", e.State, e.Idx)
	agent.Run()
	e.State = NODE_STATE_SUCCESS
}

func (e *EscapeAction) OnExit(agent AgentInterface) {
	fmt.Println("onexit escape ", e.Idx, e.State)
}

////////////////////////////////////////////////////////////////////////

type EatAction struct {
	Node
	Info EatInfo
}

func NewEatAction(xnode *xml.Node) (NodeInterface, error) {
	node := &EatAction{
		Node: Node{
			Name:  xnode.Name,
			Type:  NODE_TYPE_ACTION,
			State: NODE_STATE_INVALID,
			Idx:   xnode.Height,
		}}
	if err := node.Info.Parse(xnode.Params); err != nil {
		return nil, err
	}
	return node, nil
}

func (e *EatAction) OnEnter(agent AgentInterface) {
	fmt.Println("onEnter <eat action> ", e.State, e.Idx)
	if e.GetState() == NODE_STATE_INVALID {
		e.State = NODE_STATE_RUNNING
	}
}

func (e *EatAction) Exec(dt float64, agent AgentInterface) {
	fmt.Println("exec <eat action>", e.State, e.Idx)
	agent.Eat(e.Info.Tid)
	e.State = NODE_STATE_SUCCESS
}

func (e *EatAction) OnExit(agent AgentInterface) {
	fmt.Println("onexit eat action ", e.State, e.Idx)
}
