package main

import (
	"fmt"
	"github.com/HouzuoGuo/tiedot/db"
)

func main() {

	fd, err := db.OpenDB("dbToTied")
	if err != nil {
		panic(err)
	}

	for _, name := range fd.AllCols() {
		fmt.Println(name)
	}

	teacher := fd.ForceUse("Teacher")
	tid, err := teacher.Insert(map[string]interface{}{
		"name": "lcm",
		"url":  "golang org",
		"age":  88,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("tid => ", tid)

	doc, err := teacher.Read(tid)
	if err != nil {
		panic(err)
	}

	fmt.Println(doc)

	//for _, tids := range teacher.AllIndexes() {
	//	fmt.Println("tids => ", tids)
	//}

	//teacher.ForEachDoc(func(id int, doc []byte) (moveOn bool) {
	//	var f map[string]ziface{}
	//	if json.Unmarshal(doc, &f) != nil {
	//		panic(err)
	//	}
	//	fmt.Println("id => ", id)
	//	fmt.Println("doc => ", f)
	//	return true
	//})

	//if err := teacher.Index([]string{"name"}); err != nil {
	//	panic(err)
	//}

	//var query ziface{}
	//if json.Unmarshal([]byte(`[{"eq":"lcm","in":["name"]}]`), &query) != nil {
	//	panic(err)
	//}

	//var search = make([]map[string]ziface{}, 1)
	//search[0] = make(map[string]ziface{})
	//search[0]["eq"] = "lcm"
	//search[0]["in"] = []string{"name"}
	//
	//fmt.Println("query ziface => ", search)
	//
	//queryResult := make(map[int]struct{})
	//if err := db.EvalQuery(search, teacher, &queryResult); err != nil {
	//	panic(err)
	//}
	//
	//for id := range queryResult {
	//	fmt.Println("id => ", id)
	//
	//	if _doc, err := teacher.Read(id); err != nil {
	//		panic(err)
	//	} else {
	//		fmt.Println("_doc => ", _doc)
	//	}
	//
	//}

}
