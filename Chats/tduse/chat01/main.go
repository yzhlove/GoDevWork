package main

import (
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
)

func main() {

	dir := "dbToTied"
	fd, err := db.OpenDB(dir)
	if err != nil {
		panic(err)
	}
	if err := fd.Create("User"); err != nil {
		panic(err)
	}
	if err := fd.Create("Student"); err != nil {
		panic(err)
	}

	for _, dbName := range fd.AllCols() {
		fmt.Println(dbName)
	}

}
