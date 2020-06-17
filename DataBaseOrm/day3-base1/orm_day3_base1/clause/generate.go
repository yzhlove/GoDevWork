package clause

import (
	"fmt"
	"orm_day3_base1/log"
	"strings"
)

type _GENERATE func(values ...interface{}) (string, []interface{})

var (
	_GenerateMap map[Type]_GENERATE
	_ResultEmpty []interface{}
)

func init() {
	_GenerateMap = make(map[Type]_GENERATE, 4)
	_GenerateMap[INSERT] = _insert
	_GenerateMap[VALUES] = _values
	_GenerateMap[SELECT] = _select
	_GenerateMap[LIMIT] = _limit
	_GenerateMap[WHERE] = _where
	_GenerateMap[ORDERBY] = _orderBy
}

func bindVars(num int) string {
	var vars []string
	for i := 0; i < num; i++ {
		vars = append(vars, "?")
	}
	return strings.Join(vars, ", ")
}

func _insert(values ...interface{}) (string, []interface{}) {
	//INSERT INTO $tableName ($fields)
	tableName := values[0]
	var fields string
	if v, ok := values[1].([]string); ok {
		fields = strings.Join(v, ",")
	} else {
		log.Error("bindVars type err")
	}
	return fmt.Sprintf("INSERT INTO %s (%v)", tableName, fields), _ResultEmpty
}

func _values(values ...interface{}) (string, []interface{}) {
	//VALUES ($1),($2)
	var bindStr string
	var sql strings.Builder
	var vars []interface{}
	sql.WriteString("VALUES ")
	for i, value := range values {
		if v, ok := value.([]interface{}); ok {
			if bindStr == "" {
				bindStr = bindVars(len(v))
			}
			sql.WriteString(fmt.Sprintf("(%v)", bindStr))
			if i+1 < len(values) {
				sql.WriteString(", ")
			}
			vars = append(vars, v...)
		} else {
			log.Error("_values type err")
		}
	}
	return sql.String(), vars
}

func _select(values ...interface{}) (string, []interface{}) {
	//SELECT $fields FROM $tableName
	tableName := values[0]
	var fields string
	if v, ok := values[1].([]string); ok {
		fields = strings.Join(v, ",")
	}
	return fmt.Sprintf("SELECT %v FROM %s", fields, tableName), _ResultEmpty
}

func _limit(values ...interface{}) (string, []interface{}) {
	//LIMIT $num
	return "LIMIT ?", values
}

func _where(values ...interface{}) (string, []interface{}) {
	//WHERE $desc
	desc, vars := values[0], values[1:]
	return fmt.Sprintf("WHERE %s ", desc), vars
}

func _orderBy(values ...interface{}) (string, []interface{}) {
	return fmt.Sprintf("ORDER BY %s", values[0]), _ResultEmpty
}
