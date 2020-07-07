package loader

import (
	b3 "behaviorthree"
	"behaviorthree/config"
	"behaviorthree/core"
	"behaviorthree/util"
	"fmt"
	"reflect"
	"testing"
)

type Test struct {
	value string
}

func (t *Test) Print() {
	fmt.Println(t.value)
}

func TestExample(t *testing.T) {
	maps := createMaps()
	if data, err := maps.New("Runner"); err != nil {
		t.Error("error:", err, data)
	} else {
		t.Log(reflect.TypeOf(data))
	}
}

type LogTest struct {
	core.Action
	info string
}

func (this *LogTest) Init(setting *config.BTNodeCfg) {
	this.Action.Init(setting)
	this.info = setting.GetPropertyAsString("info")
}

func (this *LogTest) OnTick(tick *core.Tick) b3.Status {
	fmt.Println("log test --> ", this.info)
	return b3.ERROR
}

func TestLoadTree(t *testing.T) {
	treeConfig, ok := config.LoadTreeCfg("tree.json")
	if ok {
		extMaps := util.NewRegisterStructMaps()
		extMaps.Register("Log", new(LogTest))

		tree := createTree(treeConfig, extMaps)
		tree.Print()

		board := core.NewBlackboard()
		for i := 0; i < 5; i++ {
			tree.Tick(i, board)
		}
	} else {
		t.Error("Load Config err!!!")
	}
}
