package main

import "fmt"

func main() {
	setup([]int{4, 8, 12, 16})
}

type ele struct {
	w int
	x int
}

func setup(numbers []int) {

	n, avg := len(numbers), 10
	Arr := make([]int, len(numbers))

	less, more := 0, n-1

	for i, r := range numbers {
		if r > avg {
			Arr[more] = i
			more--
		} else {
			Arr[less] = i
			less++
		}
	}

	Queue := make([]int, n)
	eles := make([]ele, n)
	copy(Queue, numbers)

	fmt.Println(Queue)
	fmt.Println(Arr)
	fmt.Println("[", Queue[Arr[0]], Queue[Arr[1]], Queue[Arr[2]], Queue[Arr[3]], "]")
	fmt.Println()

	less, more = less-1, more+1
	fmt.Println("less => ", less, Arr[less], Queue[Arr[less]], " more => ", more, Arr[more], Queue[Arr[more]])

	for less > 0 && more < n {
		min, max := Arr[less], Arr[more]
		eles[less] = ele{w: Queue[min], x: max}

		Queue[max] = Queue[min] + Queue[max] - avg
		if Queue[max] > avg {
			less--
		} else {
			Arr[less] = more
			more++
		}
	}

	fmt.Println("Queue => ", Queue, " Arr => ", Arr)

}
