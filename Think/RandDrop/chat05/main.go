package main

import (
	"fmt"
	"github.com/ljfuyuan/aliasrand"
)

var (
	weights int32 = 10000
	rands         = []uint64{2500, 2500, 2500, 2500}
)

func main() {
	count := 10000000
	aMap := make(map[int]int, 4)
	as, err := aliasrand.NewWeight(rands)
	if err != nil {
		panic(err)
	}

	for i := 0; i < count; i++ {
		aMap[as.Pick()]++
	}

	tShow(aMap, count)

}

func tShow(mp map[int]int, count int) {
	for i := 0; i < len(mp); i++ {
		if num, ok := mp[i]; ok {
			id := i
			fmt.Printf("© %d\t %d\t✈︎\tcount:%d \trand:%0.7f \n", id+1, num, count, float64(num)/float64(count))
		}
	}
	fmt.Println()
}
