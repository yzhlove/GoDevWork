package dialect

import "reflect"

type sqlite3 struct{}

func init() {
	RegisterDialect("sqlite3", &sqlite3{})
}

func (s *sqlite3) TypeOf(typ reflect.Value) string {
	
}

func (s *sqlite3) TableExistsSQL(tableName string) (string, []interface{}) {

}
