package main

import (
	"crypto/md5"
	"fmt"
)

//摘要算法

func main() {
	generate("www.hao123.com")
}

// 1111 1111 1111 1111 1111 1111 1111 1100
// F    F    F    F    F    F    F    3

func generate(url string) string {

	table := []string{
		"A", "B", "C", "D", "E", "F", "G", "H",
		"J", "K", "L", "M", "N", "P", "Q", "R",
		"S", "T", "U", "V", "W", "X", "Y", "Z",
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
	}

	s := md5.Sum([]byte(url))

	t := uint32(0x3FFFFFFF)
	tt := uint32(0x1F)

	for _, b := range s {
		fmt.Print(uint32(b)&t, "\t")

		n := uint32(b) & t

		for i := 0; i < 5; i++ {
			//fmt.Print(n&tt, "-")
			fmt.Print(table[n&tt])
			n >>= 6
		}

		fmt.Println()
	}

	return ""
}
