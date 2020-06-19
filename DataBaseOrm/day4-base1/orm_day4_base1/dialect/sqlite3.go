package dialect

import (
	"fmt"
	"reflect"
	"time"
)

type sqlite3 struct{}

func init() {
	Register("sqlite3", &sqlite3{})
}

func (*sqlite3) GetDBType(structField reflect.Value) string {
	switch structField.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return "integer"
	case reflect.Int64, reflect.Uint64:
		return "bigint"
	case reflect.Float32, reflect.Float64:
		return "real"
	case reflect.String:
		return "text"
	case reflect.Array, reflect.Slice:
		return "blob"
	case reflect.Struct:
		if _, ok := structField.Interface().(time.Time); ok {
			return "datetime"
		}
	}
	panic(fmt.Sprintf("sql type:%s kind:%s",
		structField.Type().Name(), structField.Type().Kind()))
}

func (*sqlite3) Check(table string) (string, []interface{}) {
	vars := []interface{}{table}
	return "select name from sql_master where type = 'table' and name = ?", vars
}
