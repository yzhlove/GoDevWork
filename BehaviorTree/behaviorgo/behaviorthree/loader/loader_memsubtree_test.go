package loader

import (
	"behaviorthree/config"
	"behaviorthree/core"
	"behaviorthree/share"
	"behaviorthree/util"
	"fmt"
	"sync"
	"testing"
	"time"
)

var MapTree = sync.Map{}
var ExtMaps = util.NewRegisterStructMaps()

func init() {

	ExtMaps.Register("Log", new(share.LogTest))
	ExtMaps.Register("SetValue", new(share.SetValue))
	ExtMaps.Register("IsValue", new(share.IsValue))

	core.SetSubTreeLoadFunc(func(s string) *core.BehaviorTree {
		println("==>load substree:", s)
		if t, ok := MapTree.Load(s); ok {
			return t.(*core.BehaviorTree)
		}
		return nil
	})

}

func TestMemSubTree(t *testing.T) {

	projectCfg, ok := config.LoadRawProjectCfg("memsubtree.b3")
	if !ok {
		panic("load err")
	}
	var first *core.BehaviorTree
	for _, v := range projectCfg.Data.Trees {
		tree := createTree(&v, ExtMaps)
		tree.Print()
		println("==>store subtree:", v.ID)
		MapTree.Store(v.ID, tree)
		if first == nil {
			first = tree
		}
	}

	board := core.NewBlackboard()
	for i := 0; i < 2; i++ {
		if first != nil {
			first.Tick(i, board)
			fmt.Println()
		}
		time.Sleep(time.Millisecond * 100)
	}

}
