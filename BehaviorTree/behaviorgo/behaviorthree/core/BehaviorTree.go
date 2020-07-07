package core

import (
	b3 "behaviorthree"
	"behaviorthree/config"
	"behaviorthree/util"
	"fmt"
	"reflect"
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

func (this *BehaviorTree) Load(data *config.BTTreeCfg, maps, extMaps *util.RegisterStructMaps) {
	this.title = data.Title
	this.description = data.Description
	this.properties = data.Properties
	this.dumpInfo = data
	nodes := make(map[string]IBaseNode)
	for id, cfg := range data.Nodes {
		spec := &cfg
		var node IBaseNode
		if spec.Category == "tree" {
			node = new(SubTree)
		} else {
			if !extMaps.IsNIL() && extMaps.CheckElem(spec.Name) {
				if t, err := extMaps.New(spec.Name); err != nil {
					fmt.Println("ExtMaps Load Err:", err)
				} else {
					node = t.(IBaseNode)
				}
			} else {
				if t, err := maps.New(spec.Name); err != nil {
					fmt.Println("BaseMaps Load Err:", err)
				} else {
					node = t.(IBaseNode)
				}
			}
		}

		if node == nil {
			panic("behavior tree load err:invalid node name:" + spec.Name + " title:" + spec.Title)
		}
		node.Ctor()
		node.Init(spec)
		node.SetBaseNodeWorker(node.(IBaseWorker))
		nodes[id] = node
		fmt.Println("Load Type Node --> ", reflect.TypeOf(node), " Ctor:", node.GetCategory())
	}

	for id, spec := range data.Nodes {
		node := nodes[id]
		if node.GetCategory() == b3.COMPOSITE && len(spec.Children) > 0 {
			for i := 0; i < len(spec.Children); i++ {
				var cid = spec.Children[i]
				comp := node.(IComposite)
				comp.AddChild(nodes[cid])
			}
		} else if node.GetCategory() == b3.DECORATOR && len(spec.Children) > 0 {
			dec := node.(IDecorator)
			dec.SetChild(nodes[spec.Child])
		}
	}
	this.root = nodes[data.Root]
}

func (this *BehaviorTree) dump() *config.BTTreeCfg {
	return this.dumpInfo
}

func (this *BehaviorTree) Tick(target interface{}, blackboard *Blackboard) b3.Status {
	if blackboard != nil {
		var tick = NewTick()
		tick.debug = this.debug
		tick.target = target
		tick.Blackboard = blackboard
		tick.tree = this

		var state = this.root._execute(tick)
		var lastOpenNodes = blackboard._getTreeData(this.id).OpenNodes
		var currOpenNodes []IBaseNode
		currOpenNodes = append(currOpenNodes, tick._openNodes...)

		var start = 0
		for i := 0; i < util.MinInt(len(lastOpenNodes), len(currOpenNodes)); i++ {
			start = i + 1
			if lastOpenNodes[i] != currOpenNodes[i] {
				break
			}
		}

		for i := len(lastOpenNodes) - 1; i >= start; i-- {
			lastOpenNodes[i]._close(tick)
		}
		blackboard._getTreeData(this.id).OpenNodes = currOpenNodes
		blackboard.SetTree("nodeCount", tick._nodeCount, this.id)
		return state
	}
	panic("the blackboard must be an instance of")
}

func (this *BehaviorTree) Print() {
	printNode(this.root, 0)
}

func printNode(root IBaseNode, depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print(" ")
	}
	fmt.Print("|-", root.GetTitle(), "<", root.GetCategory(), ">")
	if root.GetCategory() == b3.DECORATOR {
		dec := root.(IDecorator)
		if dec.GetChild() != nil {
			printNode(dec.GetChild(), depth+3)
		}
	}
	fmt.Println()
	if root.GetCategory() == b3.COMPOSITE {
		comp := root.(IComposite)
		if count := comp.GetChildCount(); count > 0 {
			for i := 0; i < count; i++ {
				printNode(comp.GetChild(i), depth+3)
			}
		}
	}
}
