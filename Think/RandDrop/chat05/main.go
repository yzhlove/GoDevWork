package main

import (
	"fmt"
	"github.com/ljfuyuan/aliasrand"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	rands = []uint64{5, 10, 10, 15}
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
