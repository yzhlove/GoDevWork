package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	a := &Alias{eles: setup([]int{5, 10, 10, 15})}

	n := 1000000
	aMap := make(map[int]int, n)

	for i := 0; i < n; i++ {
		aMap[a.pick()]++
	}

	tShow(aMap, n)

}

type ele struct {
	w int
	x int
}

type Alias struct {
	eles []ele
}

func (a *Alias) pick() int {
	r := int(rand.Int31())
	x := r % 4
	w := r % 10
	if w < a.eles[x].w {
		return x
	}
	return a.eles[x].x
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

func setup(numbers []int) []ele {
	n, avg := len(numbers), 10
	eles := make([]ele, n)

	queue := make([]int, n)
	less, more := 0, n-1

	for _, r := range numbers {
		if r > avg {
			queue[more] = r
			more--
		} else {
			queue[less] = r
			less++
		}
	}

	if less == n {
		for ; less > 0; less-- {
			eles[less-1] = ele{w: queue[less-1]}
		}
		return eles
	}

	less, more = less-1, more+1
	x := make([]bool, n)

	for more < n {
		if !x[less] {
			eles[less] = ele{w: queue[less], x: more}
			queue[more] = queue[less] + queue[more] - avg
			x[less] = true
			if queue[more] <= avg {
				less, more = more, more+1
				continue
			}
		}
		less--
	}
	eles[more-1] = ele{w: avg}
	return eles
}
