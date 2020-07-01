package bt

import (
	"behavior_day1_base1/xml"
	"fmt"
)

type CondHp struct {
	Node
	Info CondHpInfo
}

func NewCondHp(xnode *xml.Node) (*CondHp, error) {
	node := &CondHp{
		Node: Node{
			Name:  xnode.Name,
			Type:  NODE_TYPE_CONDITION,
			State: NODE_STATE_INVALID,
			Idx:   xnode.Height,
		},
	}
	if err := node.Info.Parse(xnode.Params); err != nil {
		return nil, err
	}
	return node, nil
}

func (c *CondHp) OnEnter(agent AgentInterface) {
	fmt.Println("onEnter <cond hp >", c.State, c.Idx)
	if c.GetState() == NODE_STATE_INVALID {
		c.State = NODE_STATE_RUNNING
	}
}

func (c *CondHp) Exec(dt float64, agent AgentInterface) {
	hp := agent.GetCurHp()
	fmt.Println("exec cond hp ", c.Idx, c.State, c.Info)
	if c.Info.Min <= hp && c.Info.Max > hp {
		c.State = NODE_STATE_SUCCESS
	} else {
		c.State = NODE_STATE_FAILED
	}
}

func (c *CondHp) OnExit(agent AgentInterface) {
	fmt.Println("onExit cond hp ", c.Idx, c.State)
}
