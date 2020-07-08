package loader

import (
	"behaviorthree/config"
	"behaviorthree/core"
	"behaviorthree/share"
	"behaviorthree/util"
	"testing"
)

func TestTreeCfg(t *testing.T) {

	treeCfg, ok := config.LoadTreeCfg("tree.json")
	if !ok {
		panic("load err")
	}

	extMaps := util.NewRegisterStructMaps()
	extMaps.Register("Log", new(share.LogTest))

	tree := createTree(treeCfg, extMaps)
	tree.Print()

	board := core.NewBlackboard()
	for i := 0; i < 2; i++ {
		tree.Tick(i, board)
	}
}
