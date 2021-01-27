package main

import (
	"fmt"
	"math/rand"
	"time"
)

var rands = []int{1, 2, 3, 4, 5}

func main() {

	rand.Seed(time.Now().UnixNano())

	var rMap = make(map[int]map[int]int, len(rands))
	run := 100000
	for i := 0; i < run; i++ {
		drop()
		for i, v := range rands {
			if rMap[i+1] == nil {
				rMap[i+1] = make(map[int]int, len(rands))
			}
			rMap[i+1][v]++
		}
	}

	for i := 1; i <= len(rMap); i++ {
		if rmp, ok := rMap[i]; ok {
			fmt.Printf("location: %d 🤣 \n", i)
			for j := 1; j <= len(rmp); j++ {
				if c, ok := rmp[j]; ok {
					fmt.Printf("\tnumber:%d\tcount:%d\tr:%.7f \n", j, c, float64(c)/float64(run))
				}
			}
			fmt.Println()
		}
	}

}

func drop() {
	for i := len(rands) - 1; i >= 0; i-- {
		s := rand.Int() % (i + 1)
		rands[s], rands[i] = rands[i], rands[s]
	}
}
