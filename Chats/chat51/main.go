package main

import (
	"fmt"
	"sort"
)

type INTS []int

func (i INTS) Len() int {
	return len(i)
}

func (it INTS) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}

func (it INTS) Less(i, j int) bool {
	return it[i] >= it[j]
}

func main() {

	is := []int{1, 3, 5, 7, 8, 6, 4, 2, 9, 1, 12}
	sort.Sort(INTS(is))

	fmt.Println(is)

	vaule := 24

	index := sort.Search(len(is), func(i int) bool {
		return is[i] >= vaule
	})
	fmt.Println("index = ", index, " length = ", len(is))
}
