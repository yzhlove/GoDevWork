package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var Dict = [...]string{
	"A", "B", "C", "D", "E", "F", "G", "H",
	"J", "K", "L", "M", "N", "P", "Q", "R",
	"S", "T", "U", "V", "W", "X", "Y", "Z",
	"2", "3", "4", "5", "6", "7", "8", "9",
}

const (
	DictSize = 32
	BitSize  = 5
	MoveSize = 7
)

func main() {

	var reward uint64 = 55

	for i := 0; i < 10; i++ {
		number := getEncodeNumber(reward)
		code := getEncodeString(number)
		fmt.Printf("%v %v %v %v\n", number, code, getRewardByInt(number), getRewardByString(code))
	}

}

func getEncodeNumber(reward uint64) uint64 {
	var code uint64
	for i := 0; i < MoveSize; i++ {
		code |= uint64(rand.Intn(DictSize)) << (BitSize * i)
	}
	code |= reward << (BitSize * MoveSize)
	return code
}

func getEncodeString(code uint64) string {
	str := strings.Builder{}
	status := uint64(0x1F)
	for code > 0 {
		str.WriteString(Dict[code&status])
		code >>= BitSize
	}
	return str.String()
}

func getRewardByInt(number uint64) uint64 {
	return number >> (BitSize * MoveSize)
}

func getIndex(str rune) int {
	for i, v := range Dict {
		if v == string(str) {
			return i
		}
	}
	return -1
}

func getRewardByString(code string) uint64 {
	var value uint64
	for i := len(code) - 1; i >= 0; i-- {
		if index := getIndex(rune(code[i])); index != -1 {
			value |= uint64(index) << (BitSize * i)
		}
	}
	return value
}
