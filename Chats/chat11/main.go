package main

import (
	"fmt"
	"sort"
)

type book struct {
	id   int
	rank int
}

type books []book

func (b books) Len() int {
	return len(b)
}
func (b books) Less(i, j int) bool {
	return (b[i].id < b[j].id) && (b[i].rank < b[j].rank)
}

func (b books) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {

	var bks books
	bks = append(bks, book{id: 5, rank: 7})
	bks = append(bks, book{id: 2, rank: 3})
	bks = append(bks, book{id: 7, rank: 9})
	bks = append(bks, book{id: 3, rank: 5})
	bks = append(bks, book{id: 16, rank: 20})
	bks = append(bks, book{id: 9, rank: 14})


	sort.Sort(bks)

	fmt.Println(bks)

	tag := 15

	index := sort.Search(len(bks), func(i int) bool {
		return tag < bks[i].rank
	})

	if index != len(bks) {
		fmt.Printf("i = %v value = %v \n", index, bks[index])
	} else {
		fmt.Println("not found ")
	}

}
