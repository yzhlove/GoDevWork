package core

import (
	"fmt"
	"reflect"
)

type TreeData struct {
	NodeMem        *Memory
	OpenNodes      []IBaseNode
	TraversalDepth int
	TraversalCycle int
}

func NewTreeData() *TreeData {
	return &TreeData{NodeMem: NewMemory()}
}

/////////////////////////////////////////////////////////////////////

type Memory struct {
	_mem map[string]interface{}
}

func NewMemory() *Memory {
	return &Memory{_mem: make(map[string]interface{})}
}

func (this *Memory) Get(key string) interface{} {
	return this._mem[key]
}

func (this *Memory) Set(key string, val interface{}) {
	this._mem[key] = val
}

func (this *Memory) Remove(key string) {
	delete(this._mem, key)
}

/////////////////////////////////////////////////////////////////////

type TreeMemory struct {
	*Memory
	_treeData *TreeData
	_nodeMem  map[string]*Memory
}

func NewTreeMemory() *TreeMemory {
	return &TreeMemory{NewMemory(), NewTreeData(), make(map[string]*Memory)}
}

/////////////////////////////////////////////////////////////////////

type Blackboard struct {
	_baseMem *Memory
	_treeMem map[string]*TreeMemory
}

func NewBlackboard() *Blackboard {
	p := &Blackboard{}
	p.Init()
	return p
}

func (this *Blackboard) Init() {
	this._baseMem = NewMemory()
	this._treeMem = make(map[string]*TreeMemory)
}

func (this *Blackboard) _getTreeMem(treeScope string) *TreeMemory {
	if _, ok := this._treeMem[treeScope]; !ok {
		this._treeMem[treeScope] = NewTreeMemory()
	}
	return this._treeMem[treeScope]
}

func (this *Blackboard) _getNodeMem(treeMem *TreeMemory, nodeScope string) *Memory {
	mem := treeMem._nodeMem
	if _, ok := mem[nodeScope]; !ok {
		mem[nodeScope] = NewMemory()
	}
	return mem[nodeScope]
}

func (this *Blackboard) _getMem(treeScope, nodeScope string) *Memory {
	mem := this._baseMem
	if len(treeScope) > 0 {
		treeMem := this._getTreeMem(treeScope)
		mem = treeMem.Memory
		if len(nodeScope) > 0 {
			mem = this._getNodeMem(treeMem, nodeScope)
		}
	}
	return mem
}

func (this *Blackboard) Set(key string, value interface{}, treeScope string, nodeScope string) {
	memory := this._getMem(treeScope, nodeScope)
	memory.Set(key, value)
}

func (this *Blackboard) SetMem(key string, value interface{}) {
	memory := this._getMem("", "")
	memory.Set(key, value)
}

func (this *Blackboard) Remove(key string) {
	memory := this._getMem("", "")
	memory.Remove(key)
}

func (this *Blackboard) SetTree(key string, value interface{}, treeScope string) {
	memory := this._getMem(treeScope, "")
	memory.Set(key, value)
}

func (this *Blackboard) _getTreeData(treeScope string) *TreeData {
	treeMem := this._getTreeMem(treeScope)
	return treeMem._treeData
}

func (this *Blackboard) Get(key, treeScope, nodeScope string) interface{} {
	memory := this._getMem(treeScope, nodeScope)
	return memory.Get(key)
}

func (this *Blackboard) GetMem(key string) interface{} {
	memory := this._getMem("", "")
	return memory.Get(key)
}

func (this *Blackboard) GetFloat64(key, treeScope, nodeScope string) float64 {
	if v := this.Get(key, treeScope, nodeScope); v != nil {
		return v.(float64)
	}
	return 0
}

func (this *Blackboard) GetBool(key, treeScope, nodeScope string) bool {
	if v := this.Get(key, treeScope, nodeScope); v != nil {
		return v.(bool)
	}
	return false
}

func (this *Blackboard) GetInt(key, treeScope, nodeScope string) int {
	if v := this.Get(key, treeScope, nodeScope); v != nil {
		return v.(int)
	}
	return 0
}

func (this *Blackboard) GetInt64(key, treeScope, nodeScope string) int64 {
	if v := this.Get(key, treeScope, nodeScope); v != nil {
		return v.(int64)
	}
	return 0
}

func (this *Blackboard) GetUInt64(key, treeScope, nodeScope string) uint64 {
	if v := this.Get(key, treeScope, nodeScope); v != nil {
		return v.(uint64)
	}
	return 0
}

func (this *Blackboard) GetInt64Safe(key, treeScope, nodeScope string) int64 {
	if v := this.Get(key, treeScope, nodeScope); v != nil {
		return ToInt64(v)
	}
	return 0
}

func (this *Blackboard) GetUInt64Safe(key, treeScope, nodeScope string) uint64 {
	if v := this.Get(key, treeScope, nodeScope); v != nil {
		return ToUint64(v)
	}
	return 0
}

func (this *Blackboard) GetInt32(key, treeScope, nodeScope string) int32 {
	if v := this.Get(key, treeScope, nodeScope); v != nil {
		return v.(int32)
	}
	return 0
}

func ToInt64(v interface{}) int64 {
	switch tv := v.(type) {
	case uint64:
		return int64(tv)
	default:
		panic(fmt.Sprintf("type trans err:%v", reflect.TypeOf(v)))
	}
}

func ToUint64(v interface{}) uint64 {
	switch tv := v.(type) {
	case int64:
		return uint64(tv)
	default:
		panic(fmt.Sprintf("type trans err:%v", reflect.TypeOf(v)))
	}
}
