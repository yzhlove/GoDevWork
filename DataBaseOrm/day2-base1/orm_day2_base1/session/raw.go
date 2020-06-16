package session

import (
	"database/sql"
	"errors"
	"fmt"
	"orm_day2_base1/dialect"
	"orm_day2_base1/log"
	"orm_day2_base1/schema"
	"reflect"
	"strings"
)

type Session struct {
	db       *sql.DB
	dialect  dialect.Dialect
	refTable *schema.Schema
	sql      strings.Builder
	sqlVars  []interface{}
}

func New(db *sql.DB, d dialect.Dialect) *Session {
	return &Session{
		db:      db,
		dialect: d,
	}
}

func (sess *Session) Model(value interface{}) *Session {
	if sess.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(sess.refTable.Model) {
		sess.refTable = schema.Parse(value, sess.dialect)
	}
	return sess
}

func (sess *Session) RefTable() *schema.Schema {
	if sess.refTable == nil {
		log.ERROR("model is not set")
	}
	return sess.refTable
}

func (sess *Session) CreateTable() error {
	table := sess.RefTable()
	var columns []string
	for _, field := range table.Fields {
		columns = append(columns,
			fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}
	desc := strings.Join(columns, ",")
	_, err := sess.Raw(fmt.Sprintf("CREATE TABLE %s (%s);", table.Name, desc)).Exec()
	return err
}

func (sess *Session) DropTable() error {
	s := sess.Raw(fmt.Sprintf("DROP TABLE IF EXISTS %s", sess.RefTable().Name))
	if s == nil {
		log.ERROR("sess is nil...")
		return errors.New("session is nil")
	}
	_, err := s.Exec()
	return err
}

func (sess *Session) HasTable() bool {
	sqlstring, values := sess.dialect.TableExistsSQL(sess.RefTable().Name)
	row := sess.Raw(sqlstring, values...).QueryRow()
	var tmp string
	_ = row.Scan(&tmp)
	return tmp == sess.RefTable().Name
}

func (sess *Session) Raw(sql string, values ...interface{}) *Session {
	sess.sql.WriteString(sql)
	sess.sql.WriteString(" ")
	sess.sqlVars = append(sess.sqlVars, values...)
	return sess
}

func (sess *Session) Exec() (result sql.Result, err error) {
	defer sess.Clean()
	log.INFO(sess.sql.String(), sess.sqlVars)
	if result, err = sess.DB().Exec(sess.sql.String(), sess.sqlVars...); err != nil {
		log.ERROR(err)
	}
	return
}

func (sess *Session) Clean() {
	sess.sql.Reset()
	sess.sqlVars = sess.sqlVars[:0]
}

func (sess *Session) DB() *sql.DB {
	return sess.db
}

func (sess *Session) QueryRow() *sql.Row {
	defer sess.Clean()
	log.INFO(sess.sql.String(), sess.sqlVars)
	return sess.DB().QueryRow(sess.sql.String(), sess.sqlVars...)
}

func (sess *Session) QueryRows() (rows *sql.Rows, err error) {
	defer sess.Clean()
	log.INFO(sess.sql.String(), sess.sqlVars)
	if rows, err = sess.DB().Query(sess.sql.String(), sess.sqlVars...); err != nil {
		log.ERROR(err)
	}
	return
}
