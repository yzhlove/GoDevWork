package main

import (
	"encoding/json"
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
	"math/rand"
	"time"
)

const DataBase = "TestTieDB_03"
const Table = "yzh"

func main() {

	//for i := 1; i <= 10; i++ {
	//	create2(i)
	//}

	//counter()
	//foreach()

	//page()

	fmt.Println("===================================================")
	//search()
	//search2()

	//indexs()

	//create3()
	showall()
	search3()
}

func create3() {

	td, err := db.OpenDB(DataBase)
	if err != nil {
		panic(err)
	}

	col := td.ForceUse(Table)
	col.Index([]string{"ts"})

	for i := 0; i < 20; i++ {
		doc := map[string]interface{}{
			"name": "lcm",
			"ts":   time.Now().Unix(),
		}
		_, _ = col.Insert(doc)
		time.Sleep(time.Second)
	}

}

func showall() {

	td, err := db.OpenDB(DataBase)
	if err != nil {
		panic(err)
	}

	col := td.Use(Table)

	col.ForEachDoc(func(id int, doc []byte) (moveOn bool) {
		fmt.Println("id => ", id, " doc => ", string(doc))
		return true
	})

}

func search3() {

	td, err := db.OpenDB(DataBase)
	if err != nil {
		panic(err)
	}
	col := td.Use(Table)

	s := map[string]interface{}{
		"int-form": int(1575438563),
		"int-to":   int(time.Now().Unix()),
		"in":       []interface{}{"ts"},
	}

	ids := make(map[int]struct{})
	_ = db.EvalQuery(s, col, &ids)

	fmt.Println(len(ids))

}

func indexs() {
	td, err := db.OpenDB(DataBase)
	if err != nil {
		panic(err)
	}

	col := td.Use(Table)
	fmt.Println(col.AllIndexes())
}

func search2() {
	td, err := db.OpenDB(DataBase)
	if err != nil {
		panic(err)
	}

	col := td.Use(Table)

	ids := make(map[int]struct{})

	//s1 := map[string]interface{}{
	//	"int-from": 1,
	//	"int-to":   2,
	//	"in":       []interface{}{"number"},
	//	"limit":    20,
	//}

	var s []interface{}

	s = append(s, map[string]interface{}{
		"int-from": 1,
		"int-to":   1,
		"in":       []interface{}{"number"},
	})

	s = append(s, map[string]interface{}{
		"eq": "lcmlove",
		"in": []interface{}{"name"},
	})

	s3 := map[string]interface{}{
		"n": s,
	}

	//var s2 []interface{}
	//s2 = append(s2, map[string]interface{}{
	//	"eq": "lcmlove",
	//	"in": []interface{}{"name"},
	//})
	//
	//s2 = append(s2, map[string]interface{}{
	//	"int-from": 1,
	//	"int-to":   2,
	//	"in":       []interface{}{"number"},
	//	"limit":    20,
	//})

	col.Index([]string{"number", "name"})

	data, _ := json.Marshal(s3)
	fmt.Println("json => ", string(data))

	if err := db.EvalQuery(s3, col, &ids); err != nil {
		panic(err)
	}

	fmt.Println(len(ids))

	for id := range ids {
		if doc, err := col.Read(id); err != nil {
			panic(err)
		} else {
			fmt.Println("id => ", id, " doc => ", doc)
		}
	}

}

func search() {
	td, err := db.OpenDB(DataBase)
	if err != nil {
		panic(err)
	}

	col := td.Use(Table)

	ids := make(map[int]struct{})

	var c interface{}
	if err := json.Unmarshal([]byte(`{"n":[{"in":["number"],"int-from":1,"int-to":2},{"eq":"lcmlove","in":["name"]}]}`), &c); err != nil {
		panic(err)
	}

	fmt.Printf("%T %v \n", c, c)

	d := map[string]interface{}{
		"int-from": 2,
		"int-to":   1,
		"in":       []interface{}{"number"},
		"limit":    20,
	}

	r, _ := json.Marshal(d)
	fmt.Println("marshal => ", string(r))

	if err := db.EvalQuery(c, col, &ids); err != nil {
		panic(err)
	}

	fmt.Println(len(ids))

	for id := range ids {
		if doc, err := col.Read(id); err != nil {
			panic(err)
		} else {
			fmt.Println("id => ", id, " doc => ", doc)
		}
	}

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

func create2(i int) {
	td, err := db.OpenDB(DataBase)
	if err != nil {
		panic(err)
	}

	col := td.ForceUse(Table)
	doc := map[string]interface{}{
		"name":   "lcmlove",
		"number": i,
	}
	if _, err = col.Insert(doc); err != nil {
		panic(err)
	}

}
