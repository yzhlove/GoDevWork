package main

import (
	"fmt"
	"hash/crc32"
	"strconv"
)

var (
	slot    = "*#06#" //颜值
	code    = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	codelen = uint64(len(code))
)

func main() {

	var source uint64 = 123456789
	//补位
	newSource := ToEncodeNumber(source)
	fmt.Printf("source %v \tnewSource %v \n", source, newSource)
	//编码
	encodeStr := ToEncodeString(newSource)
	fmt.Printf("newSource %v \tencodeStr %v \n", newSource, encodeStr)
	//解码
	decodeNumber := ToDecodeString(encodeStr)
	fmt.Printf("encodeStr %v \tdecodeNumber %v \n", encodeStr, decodeNumber)
	//还原source
	toSource := ToDecodeNumber(decodeNumber)
	fmt.Printf("toSource %v \n", toSource)

	/*
		output:
		source 123456789        newSource 18217445214902027264
		newSource 18217445214902027264  encodeStr 6PNiL1j5khL
		encodeStr 6PNiL1j5khL   decodeNumber 18217445214902027264
		toSource 123456789
	*/
}

func ToEncodeNumber(num uint64) uint64 {
	str := strconv.FormatUint(num, 10) + slot
	checkCode := crc32.ChecksumIEEE([]byte(str)) % (1 << 8)
	var baseNumber uint64
	for i := 0; i < 5; i++ {
		baseNumber <<= 11
		baseNumber |= num&(1<<10-1) | (1 << 10)
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

func ToDecodeNumber(number uint64) uint64 {
	var baseNumber uint64
	for i := 0; i < 5; i++ {
		baseNumber <<= 10
		baseNumber |= number & uint64(1<<10-1)
		number >>= 11
	}
	checkCode := number & (1<<8 - 1)
	str := strconv.FormatUint(baseNumber, 10) + slot
	getCheckCode := crc32.ChecksumIEEE([]byte(str)) % (1 << 8)
	if checkCode == uint64(getCheckCode) {
		return baseNumber
	}
	return 0
}
