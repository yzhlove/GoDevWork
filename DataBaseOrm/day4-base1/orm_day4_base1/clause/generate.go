package clause

import (
	"fmt"
	"orm_day4_base1/log"
	"strings"
)

type _GENERATE func(values ...interface{}) (string, []interface{})

var (
	_GenerateMap map[T]_GENERATE
)

func init() {
	_GenerateMap = make(map[T]_GENERATE, 8)
	_GenerateMap[INSERT] = _insert
	_GenerateMap[VALUES] = _values
	_GenerateMap[COUNT] = _count
	_GenerateMap[DELETE] = _delete
	_GenerateMap[LIMIT] = _limit
	_GenerateMap[ORDER_BY] = _order_by
	_GenerateMap[SELECT] = _select
	_GenerateMap[UPDATE] = _update
	_GenerateMap[WHERE] = _where
}

func _bind(count int) string {
	values := make([]string, 0, count)
	for i := 0; i < count; i++ {
		values = append(values, "?")
	}
	return strings.Join(values, ",")
}

func _insert(values ...interface{}) (string, []interface{}) {
	//insert into $table ($fields)
	table, ok := values[0].(string)
	if !ok {
		log.Error("insert table type error")
		return "", nil
	}
	var fields string
	if strs, ok := values[1].([]string); ok {
		fields = strings.Join(strs, ",")
	} else {
		log.Error("insert values type error")
		return "", nil
	}
	return fmt.Sprintf("insert into %s (%s)", table, fields), nil
}

func _values(values ...interface{}) (string, []interface{}) {
	//values ($1) , ($2)
	var sql strings.Builder
	var vars []interface{}
	var bind string
	sql.WriteString("values ")
	for i, value := range values {
		if params, ok := value.([]interface{}); ok {
			if bind == "" {
				bind = _bind(len(params))
			}
			sql.WriteString(fmt.Sprintf("(%s)", bind))
			if i <= len(values)-1 {
				sql.WriteString(",")
			}
			vars = append(vars, params...)
		} else {
			log.Error("values type error")
		}
	}
	return sql.String(), vars
}

func _select(values ...interface{}) (string, []interface{}) {
	table, ok := values[0].(string)
	if !ok {
		log.Error("select table type error")
		return "", nil
	}
	var fields string
	if strs, ok := values[1].([]string); ok {
		fields = strings.Join(strs, ",")
	} else {
		log.Error("select values type error")
	}
	return fmt.Sprintf("select %s from %s", table, fields), nil
}

func _limit(values ...interface{}) (string, []interface{}) {
	//limit $num
	return "limit ?", values
}

func _where(values ...interface{}) (string, []interface{}) {
	//where $desc
	return fmt.Sprintf("where %s", values[0]), values[1:]
}

func _order_by(values ...interface{}) (string, []interface{}) {
	return fmt.Sprintf("order by %s", values[0]), nil
}

func _update(values ...interface{}) (string, []interface{}) {
	table, ok := values[0].(string)
	if !ok {
		log.Error("update table type error")
		return "", nil
	}
	var keys []string
	var vars []interface{}
	if params, ok := values[1].(map[string]interface{}); ok {
		for key, value := range params {
			keys = append(keys, key+" = ?")
			vars = append(vars, value)
		}
	} else {
		log.Error("update values type error")
		return "", nil
	}
	return fmt.Sprintf("update %s set %s", table, strings.Join(keys, ", ")), vars
}

func _delete(values ...interface{}) (string, []interface{}) {
	return fmt.Sprintf("delete from %s", values[0]), nil
}

func _count(values ...interface{}) (string, []interface{}) {
	return _select(values[0], []string{"count(*)"})
}
