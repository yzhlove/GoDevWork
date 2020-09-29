package registry

import "sync"

type Registry struct {
	records map[int32]interface{}
	sync.RWMutex
}

func (r *Registry) init() {
	r.records = make(map[int32]interface{})
}

func (r *Registry) Register(id int32, v interface{}) {
	r.Lock()
	defer r.Unlock()
	r.records[id] = v
}

func (r *Registry) UnRegister(id int32, v interface{}) {
	r.Lock()
	defer r.Unlock()
	if old, ok := r.records[id]; ok {
		if old == v {
			delete(r.records, id)
		}
	}
}

func (r *Registry) Query(id int32) interface{} {
	r.RLock()
	defer r.RUnlock()
	return r.records[id]
}

func (r *Registry) Count() int {
	r.RLock()
	defer r.RUnlock()
	return len(r.records)
}

var (
	_registry Registry
)

func init() {
	_registry.init()
}

func Register(id int32, v interface{}) {
	_registry.Register(id, v)
}

func UnRegister(id int32, v interface{}) {
	_registry.UnRegister(id, v)
}

func Query(id int32) interface{} {
	return _registry.Query(id)
}

func Count() int {
	return _registry.Count()
}
