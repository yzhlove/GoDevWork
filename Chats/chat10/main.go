package main

import (
	"fmt"
	"sort"
)

type Item struct {
	ID  int
	Num int
}

type List []Item

func (q List) Len() int {
	return len(q)
}

func (q List) Less(i, j int) bool {
	return q[i].Num >= q[j].Num
}

func (q List) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func main() {

	var list List
	list = append(list, Item{ID: 1, Num: 1})
	list = append(list, Item{ID: 12, Num: 12})
	list = append(list, Item{ID: 2, Num: 2})
	list = append(list, Item{ID: 17, Num: 17})
	list = append(list, Item{ID: 23, Num: 23})
	list = append(list, Item{ID: 4, Num: 4})
	list = append(list, Item{ID: 24, Num: 24})
	list = append(list, Item{ID: 45, Num: 45})
	list = append(list, Item{ID: 33, Num: 33})
	list = append(list, Item{ID: 65, Num: 65})
	list = append(list, Item{ID: 27, Num: 27})

	sort.Sort(list)

	tag := 32

	fmt.Println(list)

	index := sort.Search(len(list), func(i int) bool {
		return list[i].ID <= tag
	})

	if index != len(list) {
		fmt.Printf("index = %v length = %v value:%v \n", index, len(list), list[index])
	} else {
		fmt.Println("not found")
	}

}
