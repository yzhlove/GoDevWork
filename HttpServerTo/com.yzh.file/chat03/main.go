package main

import (
	"fmt"
	"hash/crc32"
	"strconv"
)

var (
	slot    = "*#06#"
	code    = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	codelen = uint64(len(code))
)

func main() {

	var number uint64 = 123456789
	n := ToEncodeNumber(number)
	fmt.Println("toNumber => ", n)
	c := ToEncodeString(n)
	fmt.Println("code => ", c)

	c2 := ToDecodeString(c)
	fmt.Println("c2 ==> ", c2)

	ToDecodeNumber(c2)

}

func ToEncodeNumber(num uint64) uint64 {
	str := strconv.FormatUint(num, 10) + slot
	checkCode := crc32.ChecksumIEEE([]byte(str)) % (1 << 8)
	var baseNumber uint64
	for i := 0; i < 5; i++ {
		baseNumber <<= uint64(i) * 11
		baseNumber |= num & (1<<11 - 1)
		num >>= 10
	}
	return baseNumber | (1 << 63) | (uint64(checkCode) << 55)
}

func ToEncodeString(number uint64) string {
	newCode := make([]rune, 0, 11)
	for number >= codelen {
		newCode = append(newCode, code[number%codelen])
		number /= codelen
	}
	newCode = append(newCode, code[number])
	return string(newCode)
}

func ToDecodeNumber(number uint64) uint64 {
	var baseNumber uint64
	for i := 0; i < 5; i++ {
		baseNumber <<= uint64(i) * 10
		fmt.Printf("%b \n", number)
		fmt.Printf("nnn ==> %b \n", uint64(1<<10-1))
		fmt.Println(" result =>  ", number&uint64(1<<10-1))
		baseNumber |= number & uint64(1<<10-1)
		fmt.Printf("%b \n", baseNumber)
		number >>= 11
	}
	fmt.Println("baseNumber => ", baseNumber)
	return 0
}

func ToDecodeString(code string) (number uint64) {
	runes := []rune(code)
	post := func(char rune) uint64 {
		if char < 58 {
			return uint64(char - 48)
		} else if char < 91 {
			return uint64(char - 65 + 10)
		}
		return uint64(char - 97 + 36)
	}
	for i := len(runes); i > 0; i-- {
		number = number*codelen + post(runes[i-1])
	}
	return
}
