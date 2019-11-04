package main

import (
	"fmt"
	"math"
)

func main() {

	var a int32 = math.MaxInt32
	var b int32 = a + a + 1
	var c int32 = a + a + 2
	var d int32 = a + a + 3
	fmt.Printf("b = %v c = %v d = %v ", b, c, d)
}
