package dialect

import "reflect"

var _DialectMap = map[string]Dialect{}

type Dialect interface {
	GetDataBaseType(typ reflect.Value) string
	TableExistsSQL(tableName string) (string, []interface{})
}

func RegisterDialect(name string, dialect Dialect) {
	_DialectMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = _DialectMap[name]
	return
}
