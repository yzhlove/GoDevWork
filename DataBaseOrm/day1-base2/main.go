package main

import (
	"orm_day1_base2"
	"orm_day1_base2/log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, err := orm_day1_base2.NewEngine("sqlite3", "gee.db")
	if err != nil {
		log.Error(err)
		return
	}
	defer engine.Close()
	sess := engine.NewSession()
	sess.Raw("drop table if exists User;").Exec()
	sess.Raw("create table  User(Name text);").Exec()
	sess.Raw("create table User(Name text);").Exec()
	result, err := sess.Raw("insert into User(`Name`) values(?),(?)", "tom", "sam").Exec()
	if err != nil {
		log.Error(err)
		return
	}
	count, _ := result.RowsAffected()
	log.Info("count -> ", count)
}
