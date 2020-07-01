package bt

type NODE_TYPE uint8

const (
	NODE_TYPE_ROOT NODE_TYPE = iota
	NODE_TYPE_SEQUENCE
	NODE_TYPE_SELECTOR
	NODE_TYPE_PARALLEL
	NODE_TYPE_CONDITION
	NODE_TYPE_ACTION
)

type NODE_STATE uint8

const (
	NODE_STATE_INVALID NODE_STATE = iota
	NODE_STATE_SUCCESS
	NODE_STATE_RUNNING
	NODE_STATE_FAILED
)

type NodeInterface interface {
	AddChild(node NodeInterface)
	AddParent(node NodeInterface)
	GetName() string
	GetType() NODE_TYPE
	GetState() NODE_STATE
	GetParent() NodeInterface
	GetChildren() []NodeInterface
	Reset()
	OnEnter(agent AgentInterface)
	Exec(dt float64, agent AgentInterface)
	OnExit(agent AgentInterface)
	ActionInterface
}

type ActionInterface interface {
	TickUpdate(node NodeInterface, opts ...ModeOption)
}

type ModeOption func(opt *Option)

type Option struct {
	dt    float64
	agent AgentInterface
}

func WithOpt(dt float64, agent AgentInterface) ModeOption {
	return func(opt *Option) {
		opt.dt = dt
		opt.agent = agent
	}
}

func WithDt(dt float64) ModeOption {
	return func(opt *Option) {
		opt.dt = dt
	}
}

func WithAgent(agent AgentInterface) ModeOption {
	return func(opt *Option) {
		opt.agent = agent
	}
}

type BaseAction struct{}

func (*BaseAction) TickUpdate(node NodeInterface, opts ...ModeOption) {
	opt := &Option{dt: 10}
	for _, f := range opts {
		f(opt)
	}
	node.OnEnter(opt.agent)
	node.Exec(opt.dt, opt.agent)
	node.OnExit(opt.agent)
}
