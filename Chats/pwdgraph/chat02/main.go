package main

import (
	"fmt"
	"hash/crc32"
)

//crc32

var solt = []byte("*#06#")
var src = []byte("yuzihan")

func main() {

	fmt.Println(toIEEE())
	fmt.Println(toKoopman())
	fmt.Println(toCastagnoli())
	fmt.Println(toDefault())
	fmt.Println(toSolt())
}

func toIEEE() uint32 {
	t := crc32.MakeTable(crc32.IEEE)
	return crc32.Checksum(src, t)
}

func toKoopman() uint32 {
	t := crc32.MakeTable(crc32.Koopman)
	return crc32.Checksum(src, t)
}

func toCastagnoli() uint32 {
	t := crc32.MakeTable(crc32.Castagnoli)
	return crc32.Checksum(src, t)
}

func toDefault() uint32 {
	return crc32.ChecksumIEEE(src)
}

func toSolt() uint32 {
	//加盐
	return crc32.ChecksumIEEE(append(src, solt...))
}
