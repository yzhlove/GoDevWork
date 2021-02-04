package main

import "fmt"

func main() {
	setup([]int{10, 10, 10, 10})
	setup([]int{15, 8, 9, 8})
	setup([]int{5, 8, 10, 17})
	setup([]int{2, 18, 10, 10})
	setup([]int{5, 15, 4, 16})
}

type ele struct {
	w int
	x int
}

func setup(numbers []int) {

	n, avg := len(numbers), 10

	Queue := make([]int, n)
	less, more := 0, n-1

	eles := make([]ele, n)

	for _, r := range numbers {
		if r > avg {
			Queue[more] = r
			more--
		} else {
			Queue[less] = r
			less++
		}
	}

	if less == n {
		for i := 0; i < less; i++ {
			eles[i] = ele{w: avg}
		}
		fmt.Println("Elements => ", eles)
		return
	}

	less, more = less-1, more+1
	tx := make([]bool, n)

	for more < n {
		if !tx[less] {
			eles[less] = ele{w: Queue[less], x: more}
			Queue[more] = Queue[less] + Queue[more] - avg
			tx[less] = true
			if Queue[more] <= avg {
				less = more
				more++
				continue
			}
		}
		less--
	}
	eles[more-1] = ele{w: avg}

	fmt.Println("Elements => ", eles)
}
