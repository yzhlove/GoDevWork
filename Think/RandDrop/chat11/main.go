package main

import "fmt"

func main() {
	//setup([]int{4, 8, 12, 16})
	setup([]int{10, 10, 10, 10})
}

type ele struct {
	ws  int //自身占比
	idx int //别人的下标
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
	fmt.Println("less=>", less, " more=>", more)
	less, more = less-1, more+1
	eles := make([]ele, n)
	fmt.Println("less=>", less, " more=>", more)
	for more < n {
		a, b := Queue[less], Queue[more]
		eles[less] = ele{ws: Queue[less], idx: more}
		s := a + b - avg
		Queue[more] = s
		if s > avg {
			less--
		} else {
			less = more
			more++
		}
	}

	fmt.Println("less=>", less, " more=>", more)

	for ; less < more; less++ {
		eles[less] = ele{idx: less}
	}

	fmt.Println("Queue=>", Queue)
	fmt.Println("Eles=>", eles)

}
