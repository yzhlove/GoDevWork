package session

import (
	"orm_day3_base1/clause"
	"reflect"
)

func (sess *Session) Insert(values ...interface{}) (int64, error) {

	structValues := make([]interface{}, 0, len(values))
	for _, value := range values {
		table := sess.Model(value).RefTable()
		sess.clause.Set(clause.INSERT, table.Name, table.FieldNames)
		structValues = append(structValues, table.RecordValues(value))
	}
	sess.clause.Set(clause.VALUES, structValues...)
	sql, vars := sess.clause.Build(clause.INSERT, clause.VALUES)
	result, err := sess.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (sess *Session) Find(values interface{}) error {
	destSlice := reflect.Indirect(reflect.ValueOf(values))
	destType := destSlice.Type().Elem()
	table := sess.Model(reflect.New(destType).Elem().Interface()).RefTable()

	sess.clause.Set(clause.SELECT, table.Name, table.FieldNames)
	sql, vars := sess.clause.Build(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT)
	rows, err := sess.Raw(sql, vars...).QueryRows()
	if err != nil {
		return err
	}
	for rows.Next() {
		dest := reflect.New(destType).Elem()
		var values []interface{}
		for _, name := range table.FieldNames {
			values = append(values, dest.FieldByName(name).Addr().Interface())
		}
		if err := rows.Scan(values...); err != nil {
			return err
		}
		destSlice.Set(reflect.Append(destSlice, dest))
	}
	return rows.Close()
}
