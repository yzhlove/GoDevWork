package clause

import (
	"orm_day4_base1/log"
	"strings"
)

type T int

const (
	INSERT T = iota
	VALUES
	SELECT
	LIMIT
	WHERE
	ORDER_BY
	UPDATE
	DELETE
	COUNT
)

type Clause struct {
	sql     map[T]string
	sqlVars map[T][]interface{}
}

func (c *Clause) Set(name T, values ...interface{}) {
	if c.sql == nil {
		c.sql = make(map[T]string, len(values))
		c.sqlVars = make(map[T][]interface{})
	}
	if fn, ok := _GenerateMap[name]; ok {
		sql, vars := fn(values...)
		c.sql[name] = sql
		c.sqlVars[name] = vars
	} else {
		log.Error("nof found type:", name)
	}
}

func (c *Clause) Build(tps ...T) (string, []interface{}) {
	var sqlstring []string
	var vars []interface{}
	for _, tp := range tps {
		if sql, ok := c.sql[tp]; ok {
			sqlstring = append(sqlstring, sql)
			vars = append(vars, c.sqlVars[tp])
		}
	}
	return strings.Join(sqlstring, " "), vars
}

func (c *Clause) Clear() {
	*c = Clause{}
}
