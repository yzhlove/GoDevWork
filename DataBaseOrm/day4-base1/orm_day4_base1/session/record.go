package session

import (
	"errors"
	"orm_day4_base1/clause"
	"orm_day4_base1/log"
	"reflect"
)

// object user1 user2
func (sess *Session) Insert(values ...interface{}) (int64, error) {

	objects := make([]interface{}, 0, len(values))
	for _, value := range values {
		tb := sess.Model(value).RefTb()
		sess.clause.Set(clause.INSERT, tb.Name, tb.FieldNames)
		objects = append(objects, tb.Build(value))
	}
	sess.clause.Set(clause.VALUES, objects...)
	sql, vars := sess.clause.Build(clause.INSERT, clause.VALUES)
	result, err := sess.BuildSQL(sql, vars...).Exec()
	if err != nil {
		log.Error("insert err:", err)
		return 0, err
	}
	return result.RowsAffected()
}

//object list []Users
func (sess *Session) Find(values interface{}) error {

	destSlice := reflect.Indirect(reflect.ValueOf(values))
	destType := destSlice.Type().Elem()
	tb := sess.Model(reflect.New(destType).Elem().Interface()).RefTb()

	sess.clause.Set(clause.SELECT, tb.Name, tb.FieldNames)
	sql, vars := sess.clause.Build(clause.SELECT, clause.WHERE, clause.ORDER_BY, clause.LIMIT)
	rows, err := sess.BuildSQL(sql, vars...).QueryRows()
	if err != nil {
		log.Error("find err:", err)
		return err
	}
	for rows.Next() {
		dest := reflect.New(destType).Elem()
		var object []interface{}
		for _, field := range tb.FieldNames {
			object = append(object, dest.FieldByName(field).Addr().Interface())
		}
		if err := rows.Scan(object...); err != nil {
			log.Error("find scan err:", err)
			return err
		}
		destSlice.Set(reflect.Append(destSlice, dest))
	}
	return rows.Close()
}

func (sess *Session) First(value interface{}) error {
	dest := reflect.Indirect(reflect.ValueOf(value))
	destSlice := reflect.New(reflect.SliceOf(dest.Type())).Elem()
	if err := sess.Limit(1).Find(destSlice.Addr().Interface()); err != nil {
		log.Error("first err:", err)
		return err
	}
	if destSlice.Len() == 0 {
		return errors.New("not found")
	}
	dest.Set(destSlice.Index(0))
	return nil
}

func (sess *Session) Limit(num int) *Session {
	sess.clause.Set(clause.LIMIT, num)
	return sess
}

func (sess *Session) Where(desc string, args ...interface{}) *Session {
	var vars []interface{}
	sess.clause.Set(clause.WHERE, append(append(vars, desc), args...)...)
	return sess
}

func (sess *Session) OrderBy(desc string) *Session {
	sess.clause.Set(clause.ORDER_BY, desc)
	return sess
}

func (sess *Session) Update(kv ...interface{}) (int64, error) {
	m, ok := kv[0].(map[string]interface{})
	if !ok {
		m = make(map[string]interface{})
		for i := 0; i < len(kv); i += 2 {
			if key, ok := kv[i].(string); ok {
				m[key] = kv[i+1]
			}
		}
	}
	sess.clause.Set(clause.UPDATE, sess.RefTb().Name, m)
	sql, vars := sess.clause.Build(clause.UPDATE, clause.WHERE)
	result, err := sess.BuildSQL(sql, vars...).Exec()
	if err != nil {
		log.Error("update err:", err)
		return 0, err
	}
	return result.RowsAffected()
}

func (sess *Session) Delete() (int64, error) {
	sess.clause.Set(clause.DELETE, sess.RefTb().Name)
	sql, vars := sess.clause.Build(clause.DELETE, clause.WHERE)
	result, err := sess.BuildSQL(sql, vars...).Exec()
	if err != nil {
		log.Error("delete err:", err)
		return 0, err
	}
	return result.RowsAffected()
}

func (sess *Session) Count() (int64, error) {
	sess.clause.Set(clause.COUNT, sess.RefTb().Name)
	sql, vars := sess.clause.Build(clause.COUNT, clause.WHERE)
	row := sess.BuildSQL(sql, vars...).QueryRow()
	var result int64
	if err := row.Scan(&result); err != nil {
		log.Error("count err:", err)
		return 0, err
	}
	return result, nil
}
