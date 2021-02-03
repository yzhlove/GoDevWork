package main

import "fmt"

func main() {
	setup([]int{4, 8, 12, 16})
}

func setup(numbers []int) {

	n, avg := len(numbers), 10

	Queue := make([]int, n)

	less, more := 0, n-1

	for i := 0; i < n; i++ {
		if numbers[i] > avg {
			Queue[more] = numbers[i]
			more--
		} else {
			Queue[less] = numbers[i]
			less++
		}
	}

	fmt.Println("Queue => ", Queue)
	less--
	more++
	for less >= 0 && more < n {
		a, b := Queue[less], Queue[more]
		if c := a + b - avg; c > avg {
			Queue[more] = c
			less--
		} else {
			Queue[more] = c
			less = more
			more++
		}
	}

	fmt.Println("Queue => ", Queue)

}
