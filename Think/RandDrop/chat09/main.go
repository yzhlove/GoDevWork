package main

import "fmt"

func main() {
	setup([]int{4, 8, 12, 16})
}

func setup(numbers []int) {

	n, avg := len(numbers), 10

	lessQ := make([]int, 0, n)
	moreQ := make([]int, 0, n)

	for i := 0; i < n; i++ {
		if numbers[i] > avg {
			moreQ = append(moreQ, numbers[i])
		} else {
			lessQ = append(lessQ, numbers[i])
		}
	}

	less, more := 0, 0

	for more < len(moreQ) {
		if c := lessQ[less] + moreQ[more]; c >= avg {
			last := c - avg
			moreQ[more] = last
			if last <= avg {
				lessQ = append(lessQ, last)
				more++
			}
			less++
		} else {
			lessQ = append(lessQ, moreQ[more])
			more++
		}
	}

	fmt.Println("lessQ => ", lessQ, " moreQ => ", moreQ)

}
