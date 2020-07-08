package loader

import (
	"behaviorthree/config"
	"behaviorthree/core"
	"behaviorthree/share"
	"behaviorthree/util"
	"sync"
	"testing"
)

var Tree = sync.Map{}

func init() {
	core.SetSubTreeLoadFunc(func(s string) *core.BehaviorTree {
		println("===+++ >> load subtree:", s)
		if t, ok := Tree.Load(s); ok {
			return t.(*core.BehaviorTree)
		}
		return nil
	})
}

func TestSubTree(t *testing.T) {

	projectCfg, ok := config.LoadRawProjectCfg("example.b3")
	if !ok {
		panic("load err")
	}

	extMaps := util.NewRegisterStructMaps()
	extMaps.Register("Log", new(share.LogTest))

	var first *core.BehaviorTree
	for _, v := range projectCfg.Data.Trees {
		tree := createTree(&v, extMaps)
		tree.Print()
		println("===++>>> store subtree:", v.ID)
		Tree.Store(v.ID, tree)
		if first == nil {
			first = tree
		}
	}

	board := core.NewBlackboard()
	for i := 0; i < 2; i++ {
		if first != nil {
			first.Tick(i, board)
		}
	}

}
