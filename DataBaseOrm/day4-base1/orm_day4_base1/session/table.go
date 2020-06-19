package session

import (
	"fmt"
	"orm_day4_base1/log"
	"orm_day4_base1/schema"
	"reflect"
	"strings"
)

func (sess *Session) Model(value interface{}) *Session {
	if sess.refTb == nil || reflect.TypeOf(value) != reflect.TypeOf(sess.refTb.Model) {
		sess.refTb = schema.Parse(value, sess.dialect)
	}
	return sess
}

func (sess *Session) RefTb() *schema.Schema {
	if sess.refTb == nil {
		log.Error("model is not set")
	}
	return sess.refTb
}

func (sess *Session) CreateTb() (err error) {
	tb := sess.RefTb()
	columns := make([]string, 0, len(tb.Fields))
	for _, field := range tb.Fields {
		columns = append(columns,
			fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}
	desc := strings.Join(columns, ",")
	sql := fmt.Sprintf("create table %s (%s)", tb.Name, desc)
	_, err = sess.BuildSQL(sql).Exec()
	return
}

func (sess *Session) DropTb() (err error) {
	sql := fmt.Sprintf("drop table if exists %s", sess.RefTb().Name)
	_, err = sess.BuildSQL(sql).Exec()
	return
}

func (sess *Session) HasTb() bool {
	sql, vars := sess.dialect.Check(sess.RefTb().Name)
	row := sess.BuildSQL(sql, vars...).QueryRow()
	var result string
	row.Scan(&result)
	return result == sess.RefTb().Name
}
