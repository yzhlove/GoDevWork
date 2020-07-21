package main

import "math"

const EPSILON = 0.00000000000001 // 最小误差值

func P2C(p float64) float64 {
	var low, last float64
	var high = p
	var middle float64
	for {
		middle = (low + high) / 2
		pr := C2P(middle)
		if math.Abs(pr-last) <= EPSILON {
			break
		}
		if pr > p {
			high = middle
		} else {
			low = middle
		}
		last = pr
	}
	return middle
}

func C2P(c float64) float64 {
	var expectation, succeed, current float64
	times := math.Ceil(1 / c)
	for i := 0; i <= int(times); i++ {
		current = (1 - succeed) * math.Min(1, c*float64(i))
		succeed += current
		expectation += current * float64(i)
	}
	return 1 / expectation
}
