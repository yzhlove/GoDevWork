package core

import (
	"behaviorthree/config"
	"behaviorthree/util"
)

type BehaviorTree struct {
	id          string
	title       string
	description string
	properties  map[string]interface{}
	root        IBaseNode
	debug       interface{}
	dumpInfo    *config.BTTreeCfg
}

func NewBeTree() *BehaviorTree {
	tree := &BehaviorTree{}
	return tree
}

func (this *BehaviorTree) Init() {
	this.id = util.GetUID()
	this.title = "the behavior tree"
	this.properties = make(map[string]interface{})
	this.root = nil
	this.debug = nil
}

func (this *BehaviorTree) GetID() string {
	return this.id
}

func (this *BehaviorTree) GetTitle() string {
	return this.title
}

func (this *BehaviorTree) SetDebug(debug interface{}) {
	this.debug = debug
}

func (this *BehaviorTree) GetRoot() IBaseNode {
	return this.root
}

func (this *BehaviorTree) Load(data *config.BTTreeCfg,maps ,extMaps *util.RegisterStructMaps) {
	this.title = data.Title
	this.description = data.Description
	this.properties = data.Properties
	this.dumpInfo = data
	nodes := make(map[string]IBaseNode)

	for id , cfg := range data.Nodes {
		spec := &cfg
		var node IBaseNode

		if spec.Category == "tree" {

		} else {

		}
	}
}