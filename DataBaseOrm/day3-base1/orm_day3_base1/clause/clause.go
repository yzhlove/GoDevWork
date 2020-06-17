package clause

import (
	"orm_day3_base1/log"
	"strings"
)

type Type int

const (
	INSERT Type = iota
	VALUES
	SELECT
	LIMIT
	WHERE
	ORDERBY
)

type Clause struct {
	sql     map[Type]string
	sqlVars map[Type][]interface{}
}

func (c *Clause) Set(name Type, vars ...interface{}) {
	if c.sql == nil {
		c.sql = make(map[Type]string, len(vars))
		c.sqlVars = make(map[Type][]interface{}, len(vars))
	}
	if fn, ok := _GenerateMap[name]; ok {
		sql, vars := fn(vars...)
		c.sql[name] = sql
		c.sqlVars[name] = vars
	} else {
		log.Error("not found func by:", name)
	}
}

func (c *Clause) Build(types ...Type) (string, []interface{}) {
	var sqls []string
	var vars []interface{}
	for _, typ := range types {
		if sql, ok := c.sql[typ]; ok {
			sqls = append(sqls, sql)
			vars = append(vars, c.sqlVars[typ]...)
		}
	}
	return strings.Join(sqls, " "), vars
}
