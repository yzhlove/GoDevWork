package core

import (
	b3 "behaviorthree"
	"behaviorthree/config"
)

type IBaseWrapper interface {
	_execute(tick *Tick) b3.Status
	_enter(tick *Tick)
	_open(tick *Tick)
	_tick(tick *Tick) b3.Status
	_close(tick *Tick)
	_exit(tick *Tick)
}

type IBaseNode interface {
	IBaseWrapper

	Ctor()
	Init(params *config.BTNodeCfg)
	GetCategory() string
	Execute(tick *Tick) b3.Status
	GetName() string
	GetTitle() string
	SetBaseNodeWorker(worker IBaseWorker)
	GetBaseNodeWorker() IBaseWorker
}

type BaseNode struct {
	IBaseWorker

	id          string
	name        string
	category    string
	title       string
	description string
	parameters  map[string]interface{}
	properties  map[string]interface{}
}

func (this *BaseNode) Ctor() {}

func (this *BaseNode) SetName(name string) {
	this.name = name
}

func (this *BaseNode) SetTitle(title string) {
	this.title = title
}

func (this *BaseNode) SetBaseNodeWorker(worker IBaseWorker) {
	this.IBaseWorker = worker
}

func (this *BaseNode) GetBaseNodeWorker() IBaseWorker {
	return this.IBaseWorker
}

func (this *BaseNode) Init(params *config.BTNodeCfg) {
	this.description = ""
	this.parameters = make(map[string]interface{})
	this.properties = make(map[string]interface{})

	this.id = params.Id
	this.name = params.Name
	this.title = params.Title
	this.description = params.Description
	this.properties = params.Properties

}

func (this *BaseNode) GetCategory() string {
	return this.category
}

func (this *BaseNode) GetID() string {
	return this.id
}

func (this *BaseNode) GetName() string {
	return this.name
}

func (this *BaseNode) GetTitle() string {
	return this.title
}

func (this *BaseNode) _execute(tick *Tick) b3.Status {
	this._enter(tick)
	if !tick.Blackboard.GetBool("isOpen", tick.tree.id, this.id) {
		this._open(tick)
	}
	var status = this._tick(tick)
	if status != b3.RUNNING {
		this._close(tick)
	}
	this._exit(tick)
	return status
}

func (this *BaseNode) Execute(tick *Tick) b3.Status {
	return this._execute(tick)
}

func (this *BaseNode) _enter(tick *Tick) {
	tick._enterNode(this)
	this.OnEnter(tick)
}

func (this *BaseNode) _open(tick *Tick) {
	tick._openNode(this)
	tick.Blackboard.Set("isOpen", true, tick.tree.id, this.id)
	this.OnOpen(tick)
}

func (this *BaseNode) _tick(tick *Tick) b3.Status {
	tick._tickNode(this)
	return this.OnTick(tick)
}

func (this *BaseNode) _close(tick *Tick) {
	tick._closeNode(this)
	tick.Blackboard.Set("isOpen", false, tick.tree.id, this.id)
	this.OnClose(tick)
}

func (this *BaseNode) _exit(tick *Tick) {
	tick._exitNode(this)
	this.OnExit(tick)
}
