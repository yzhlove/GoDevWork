package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

//encoding使用

func main() {

	//bytesRead()
	//bytesWrite()
	//getUint()
	putUint()
}

func bytesRead() {

	var number uint64
	b := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	bytesBuf := bytes.NewBuffer(b)
	if err := binary.Read(bytesBuf, binary.BigEndian, &number); err != nil {
		panic(err)
	}
	fmt.Println("number => ", number)
	fmt.Println("number ==> ", binary.BigEndian.Uint64(b))
}

func bytesWrite() {
	bytesBuf := new(bytes.Buffer)
	var number uint64 = 72623859790382856
	if err := binary.Write(bytesBuf, binary.BigEndian, number); err != nil {
		panic(err)
	}
	fmt.Println("buf => ", bytesBuf.Bytes())

	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, number)
	fmt.Println("buf ==> ", buf)

}

func getUint() {

	//buf := []byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8}
	//num, ret := binary.Uvarint(buf)
	//fmt.Println(num, "-", ret)

	sbuf := []byte{}
	buf := []byte{144, 192, 192, 132, 136, 140, 144, 122}
	bbuf := []byte{144, 192, 192, 129, 132, 136, 140, 144, 192, 192, 1, 1}

	num, ret := binary.Uvarint(sbuf)
	fmt.Println(num, ret)

	num, ret = binary.Uvarint(buf)
	fmt.Println(num, ret)

	num, ret = binary.Uvarint(bbuf)
	fmt.Println(num, ret)

}

func putUint() {
	i16 := 1234
	i64 := -1234567890
	sbuf := make([]byte, 4)
	buf := make([]byte, 10)

	ret := binary.PutVarint(buf, int64(i16))
	fmt.Println(ret, len(strconv.Itoa(i16)), sbuf)

	ret = binary.PutVarint(buf, int64(i64))
	fmt.Println(ret, len(strconv.Itoa(i64)), buf)
}
