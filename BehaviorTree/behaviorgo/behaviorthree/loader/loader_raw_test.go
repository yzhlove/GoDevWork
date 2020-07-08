package loader

import (
	"behaviorthree/config"
	"behaviorthree/core"
	"behaviorthree/share"
	"behaviorthree/util"
	"fmt"
	"testing"
)

func TestLoaderRawProject(t *testing.T) {

	projectCfg, ok := config.LoadRawProjectCfg("example.b3")
	if !ok {
		fmt.Println("LoadRawProjectCfg err")
		return
	}

	extMaps := util.NewRegisterStructMaps()
	extMaps.Register("Log", new(share.LogTest))

	var firstTree *core.BehaviorTree
	for _, v := range projectCfg.Data.Trees {
		tree := createTree(&v, extMaps)
		tree.Print()
		if firstTree == nil {
			firstTree = tree
		}
	}

	board := core.NewBlackboard()
	for i := 0; i < 5; i++ {
		if firstTree != nil {
			firstTree.Tick(i, board)
		}
	}

}
