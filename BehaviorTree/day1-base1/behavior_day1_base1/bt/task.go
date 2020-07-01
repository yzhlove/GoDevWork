package bt

import "errors"

type AgentInterface interface {
	GetCurHp() int
	GetCurMp() int
	CastSkill(sid, mp int)
	IsSkillEnd() bool
	Run()
	Eat(id int)
}

type Task struct {
	Name  string
	Root  NodeInterface
	Agent AgentInterface
}

func NewTask(name string, agent AgentInterface) (*Task, error) {
	if node, ok := GlobalTreeMap[name]; !ok {
		return nil, errors.New("can't not found")
	} else {
		return &Task{Name: name, Root: node, Agent: agent}, nil
	}
}

func (t *Task) Update(dt float64) {
	t.Root.TickUpdate(t.Root, WithOpt(dt, t.Agent))
}
