package main

import (
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"math/rand"
)

const DataBase = "TestTieDB_01"
const Table = "yzh"

func main() {

	//for i := 0; i < 100; i++ {
	//	create()
	//}

	//counter()
	//foreach()

	//page()

}

func page() {

	td, err := db.OpenDB(DataBase)
	if err != nil {
		panic(err)
	}

	col := td.Use(Table)
	col.ForEachDocInPage(10, 20, func(id int, doc []byte) bool {
		fmt.Println("id => ", id, " doc => ", string(doc))
		return true
	})

}

func foreach() {
	td, err := db.OpenDB(DataBase)
	if err != nil {
		panic(err)
	}
	col := td.Use(Table)
	counter := 0
	col.ForEachDoc(func(id int, doc []byte) (moveOn bool) {
		counter++
		fmt.Println("counter => ", counter, "id => ", id, " doc => ", string(doc))
		return true
	})
	fmt.Println("counter => ", counter, " c => ", col.ApproxDocCount())
}

func counter() {

	td, err := db.OpenDB(DataBase)
	if err != nil {
		panic(err)
	}

	fmt.Println(td.Use(Table).ApproxDocCount())

}

func create() {
	td, err := db.OpenDB(DataBase)
	if err != nil {
		panic(err)
	}

	col := td.Use(Table)
	doc := map[string]interface{}{
		"name":   "yzhlove",
		"number": rand.Intn(10000) + 100,
	}
	if _, err = col.Insert(doc); err != nil {
		panic(err)
	}

}
