package main

import "fmt"

func main() {
	//setup([]int{10, 10, 10, 10})
	//setup([]int{15, 8, 9, 8})
	//setup([]int{5, 8, 10, 17})
	//setup([]int{2, 18, 10, 10})
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
	for _, r := range numbers {
		if r > avg {
			Queue[more] = r
			more--
		} else {
			Queue[less] = r
			less++
		}
	}

	eles := make([]ele, n)

	if less == n {
		for i := 0; i < less; i++ {
			eles[i] = ele{w: 10}
		}

		fmt.Println("eles =>", eles)

		return
	}

	less, more = less-1, more+1
	for more < n {
		if Queue[less] >= avg {
			eles[less] = ele{w: Queue[less]}
		} else {
			eles[less] = ele{w: Queue[less], x: more}
		}
		s := Queue[less] + Queue[more] - avg
		Queue[less], Queue[more] = avg, s
		if s > avg {
			less--
		} else {
			less = more
			more++
		}
	}
	eles[more-1] = ele{w: 10}

	fmt.Println("eles =>", eles)
}
