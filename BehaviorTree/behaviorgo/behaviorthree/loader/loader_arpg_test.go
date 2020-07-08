package loader

import (
	"behaviorthree/config"
	"behaviorthree/share"
	"testing"
)

func TestLoadARPG(t *testing.T) {

	treeCfg, ok := config.LoadRawProjectCfg("zt.b3")
	if !ok {
		panic("load err")
	}

	extMaps := createMaps()
	extMaps.Register("Log", new(share.LogTest))
	
	for _, v := range treeCfg.Data.Trees {
		tree := createTree(&v, extMaps)
		tree.Print()
		t.Log()
	}
}
