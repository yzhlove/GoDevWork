package session

import (
	"fmt"
	"orm_day3_base1/log"
	"orm_day3_base1/schema"
	"reflect"
	"strings"
)

func (sess *Session) Model(value interface{}) *Session {
	if sess.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(sess.refTable.Model) {
		sess.refTable = schema.Parse(value, sess.dialect)
	}
	return sess
}

func (sess *Session) RefTable() *schema.Schema {
	if sess.refTable == nil {
		log.Error("model is not set")
	}
	return sess.refTable
}

func (sess *Session) CreateTable() error {
	table := sess.RefTable()
	columns := make([]string, 0, len(table.Fields))
	for _, field := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}
	desc := strings.Join(columns, ",")
	sql := fmt.Sprintf("CREATE TABLE %s (%s)", sess.RefTable().Name, desc)
	_, err := sess.Raw(sql).Exec()
	return err
}

func (sess *Session) DropTable() error {
	sql := fmt.Sprintf("DROP TABLE IF EXISTS %s ", sess.RefTable().Name)
	_, err := sess.Raw(sql).Exec()
	return err
}

func (sess *Session) HasTable() bool {
	sql, values := sess.dialect.TableExistsSQL(sess.RefTable().Name)
	row := sess.Raw(sql, values...).QueryRow()
	var result string
	row.Scan(&result)
	return result == sess.RefTable().Name
}
