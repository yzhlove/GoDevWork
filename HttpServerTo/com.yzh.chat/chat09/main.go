package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(exponent(2, 3))
	fmt.Println(1 << 3)
	fmt.Println(uint64(math.Pow(2, 3)))

	fmt.Println(exponent(2, 42))
	fmt.Println(1 << 42)
	fmt.Println(uint64(math.Pow(2, 42)))

	fmt.Println(4 &^ 12)
	fmt.Println(12 &^ 4)

	//16 8 4 2 1

	fmt.Println(1<<5 - 1)

}

func exponent(a, n int64) int64 {
	result := int64(1)
	for i := n; i > 0; i >>= 1 {
		if i&1 != 0 {
			result *= a
		}
		a *= a
	}
	return result
}
