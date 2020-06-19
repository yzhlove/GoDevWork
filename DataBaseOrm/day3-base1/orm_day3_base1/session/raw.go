package session

import (
	"database/sql"
	"orm_day3_base1/clause"
	"orm_day3_base1/dialect"
	"orm_day3_base1/log"
	"orm_day3_base1/schema"
	"strings"
)

type Session struct {
	db       *sql.DB
	dialect  dialect.Dialect
	refTable *schema.Schema
	clause   clause.Clause
	sql      strings.Builder
	sqlVars  []interface{}
}

func New(db *sql.DB, d dialect.Dialect) *Session {
	return &Session{
		db:      db,
		dialect: d,
	}
}

func (sess *Session) Clear() {
	sess.sql.Reset()
	sess.sqlVars = sess.sqlVars[:0]
	sess.clause = clause.Clause{}
}

func (sess *Session) DB() *sql.DB {
	return sess.db
}

func (sess *Session) Exec() (result sql.Result, err error) {
	defer sess.Clear()
	log.Info(sess.sql.String(), " -> ", sess.sqlVars)
	if result, err = sess.DB().Exec(sess.sql.String(), sess.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

func (sess *Session) QueryRow() *sql.Row {
	defer sess.Clear()
	log.Info(sess.sql.String(), " -> ", sess.sqlVars)
	return sess.DB().QueryRow(sess.sql.String(), sess.sqlVars...)
}

func (sess *Session) QueryRows() (rows *sql.Rows, err error) {
	defer sess.Clear()
	log.Info(sess.sql.String(), " -> ", sess.sqlVars)
	if rows, err = sess.DB().Query(sess.sql.String(), sess.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

func (sess *Session) Raw(sql string, values ...interface{}) *Session {
	sess.sql.WriteString(sql)
	sess.sql.WriteString(" ")
	sess.sqlVars = append(sess.sqlVars, values...)
	return sess
}
