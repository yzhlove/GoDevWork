package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, _ := sql.Open("sqlite3", "./gee.db")
	defer func() {
		db.Close()
	}()
	db.Exec("DROP TABLE IF EXISTS User;")
	db.Exec("CREATE TABLE User (Name text);")
	result, err := db.Exec("INSERT INTO User(`Name`) values(?),(?);", "Tom", "Sam")
	if err == nil {
		affected, _ := result.RowsAffected()
		log.Println(affected)
	}
	rows, err := db.Query("select * from User")
	if err != nil {
		panic(err)
	}
	var name string
	for rows.Next() {
		if err := rows.Scan(&name); err != nil {
			panic(err)
		}
		log.Println("name => ", name)
	}
	rows.Close()
}
