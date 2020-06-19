package dialect

import "reflect"

var _dialects = map[string]Dialect{}

type Dialect interface {
	GetDBType(typ reflect.Value) string
	Check(table string) (string, []interface{})
}

func Register(name string, d Dialect) {
	_dialects[name] = d
}

func Get(name string) (d Dialect, ok bool) {
	d, ok = _dialects[name]
	return
}
