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

	_Queue := make([]int, n)
	copy(_Queue, numbers)

	_Vector := make([]int, n)
	less, more := 0, n-1
	for i, r := range _Queue {
		if r > avg {
			_Vector[more] = i
			more--
		} else {
			_Vector[less] = i
			less++
		}
	}

	less, more = less-1, more+1
	eles := make([]ele, n)

	for less >= 0 && more < n {
		eles[_Vector[less]] = ele{w: _Queue[_Vector[less]], x: _Vector[more]}
		_Queue[_Vector[more]] = _Queue[_Vector[less]] + _Queue[_Vector[more]] - avg
		if _Queue[_Vector[more]] > avg {
			less--
		} else {
			_Vector[less] = _Vector[more]
			more++
		}
	}

	for ; less >= 0; less-- {
		eles[_Vector[less]] = ele{w: avg}
	}

	for ; more < n; more++ {
		eles[_Vector[more]] = ele{w: avg}
	}
	
	fmt.Println("Queue => ", _Queue, " Elements => ", eles)
}
