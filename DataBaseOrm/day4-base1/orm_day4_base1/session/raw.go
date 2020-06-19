package session

import (
	"database/sql"
	"orm_day4_base1/clause"
	"orm_day4_base1/dialect"
	"orm_day4_base1/log"
	"orm_day4_base1/schema"
	"strings"
)

type Session struct {
	db      *sql.DB
	dialect dialect.Dialect
	refTb   *schema.Schema
	clause  clause.Clause
	sql     strings.Builder
	vars    []interface{}
}

func New(db *sql.DB, d dialect.Dialect) *Session {
	return &Session{
		db:      db,
		dialect: d,
	}
}

func (sess *Session) Clear() {
	sess.sql.Reset()
	sess.vars = sess.vars[:0]
	sess.clause.Clear()
}

func (sess *Session) DB() *sql.DB {
	return sess.db
}

func (sess *Session) Exec() (result sql.Result, err error) {
	defer sess.Clear()
	log.Info(sess.sql.String(), " ", sess.vars)
	if result, err = sess.DB().Exec(sess.sql.String(), sess.vars...); err != nil {
		log.Error(err)
	}
	return
}

func (sess *Session) BuildSQL(sql string, values ...interface{}) *Session {
	sess.sql.WriteString(sql)
	sess.sql.WriteString(" ")
	sess.vars = append(sess.vars, values...)
	return sess
}

func (sess *Session) QueryRow() *sql.Row {
	defer sess.Clear()
	log.Info(sess.sql.String(), " ", sess.vars)
	return sess.DB().QueryRow(sess.sql.String(), sess.vars...)
}

func (sess *Session) QueryRows() (rows *sql.Rows, err error) {
	defer sess.Clear()
	log.Info(sess.sql.String(), " ", sess.vars)
	if rows, err = sess.DB().Query(sess.sql.String(), sess.vars...); err != nil {
		log.Error(err)
	}
	return
}
