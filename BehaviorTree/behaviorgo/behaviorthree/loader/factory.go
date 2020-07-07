package loader

import (
	"behaviorthree/actions"
	"behaviorthree/composites"
	"behaviorthree/config"
	"behaviorthree/core"
	"behaviorthree/decorators"
	"behaviorthree/util"
)

func createMaps() *util.RegisterStructMaps {

	st := util.NewRegisterStructMaps()
	st.Register("Error", &actions.Error{})
	st.Register("Failer", &actions.Failer{})
	st.Register("Runner", &actions.Runner{})
	st.Register("Succeeder", &actions.Succeed{})
	st.Register("Wait", &actions.Wait{})
	st.Register("Log", &actions.Log{})

	st.Register("MemPriority", &composites.MemPriority{})
	st.Register("MemSequence", &composites.MemSequence{})
	st.Register("Priority", &composites.Priority{})
	st.Register("Sequence", &composites.Sequence{})

	st.Register("Inverter", &decorators.Inverter{})
	st.Register("Limiter", &decorators.Limiter{})
	st.Register("MaxTime", &decorators.MaxTime{})
	st.Register("Repeater", &decorators.Repeater{})
	st.Register("RepeatedUntilFailure", &decorators.RepeatedUntilFailure{})
	st.Register("RepeatedUntilSuccess", &decorators.RepeatedUntilSuccess{})

	return st
}

func createTree(config *config.BTTreeCfg, extMaps *util.RegisterStructMaps) *core.BehaviorTree {
	baseMaps := createMaps()
	tree := core.NewBeTree()
	tree.Load(config, baseMaps, extMaps)
	return tree
}
