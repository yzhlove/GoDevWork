package main

import (
	"fmt"
	"math/rand"
)

var (
	weights int32 = 10000
	rands         = []int32{2500, 2500, 2500, 2500}
	rands2        = []int32{500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500, 500}
)

func main() {
	count := 1000_0000
	aMap := make(map[int]int, 4)
	bMap := make(map[int]int, 4)
	cMap := make(map[int]int, 20)
	for i := 0; i < count; i++ {
		if a, b := drop(); a == -1 || b == -1 {
			panic("drop err")
		} else {
			aMap[a]++
			bMap[b]++
		}
		if c := drop3(); c == -1 {
			panic("drop3 error")
		} else {
			cMap[c]++
		}
	}
	tShow(aMap, count)
	tShow(bMap, count)
	tShow(cMap, count)
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

func drop() (i, i2 int) {
	w := weights
	var status bool
	for i, rd := range rands {
		number := rand.Int31n(w)
		if !status {
			i2 = drop2(number)
			status = true
		}
		if rd > number {
			return i, i2
		}
		w -= rd
	}
	return -1, -1
}

func drop2(number int32) int {
	for i, rd := range rands {
		if number < rd {
			return i
		}
		number -= rd
	}
	return -1
}

func drop3() int {
	t := weights
	for i, rd := range rands2 {
		if rd > rand.Int31n(t) {
			return i
		}
		t -= rd
	}
	return -1
}
