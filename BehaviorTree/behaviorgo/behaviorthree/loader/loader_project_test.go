package loader

import (
	"behaviorthree/config"
	"behaviorthree/core"
	"behaviorthree/share"
	"fmt"
	"testing"
)

func TestLoadProject(t *testing.T) {

	projectCfg, ok := config.LoadProjectCfg("project.json")
	if !ok {
		t.Error("load err")
		return
	}

	extMaps := createMaps()
	extMaps.Register("Log", new(share.LogTest))

	var firstTree *core.BehaviorTree
	for _, v := range projectCfg.Trees {
		tree := createTree(&v, extMaps)
		tree.Print()
		if firstTree == nil {
			firstTree = tree
		}
	}

	board := core.NewBlackboard()
	for i := 0; i < 5; i++ {
		firstTree.Tick(1, board)
		fmt.Println("========================")
	}

}
