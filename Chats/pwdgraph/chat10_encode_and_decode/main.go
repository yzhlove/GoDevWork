package main

import (
	"crypto/rc4"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"math"
	"strconv"
	"strings"
)

var (
	key  = []byte("*#123456789#*")
	slot = "*#06#"
	Dict = [...]string{
		"A", "B", "C", "D", "E", "F", "G", "H",
		"J", "K", "L", "M", "N", "P", "Q", "R",
		"S", "T", "U", "V", "W", "X", "Y", "Z",
		"2", "3", "4", "5", "6", "7", "8", "9",
	}
)

func main() {

	var number uint32 = 223344
	var id uint32 = 1000
	str := getEncodeString(id, number)
	getDecodeString(str)

}

func getRC4By32(number uint32) (uint32, error) {
	bytesBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(bytesBuf, number)
	if cipher, err := rc4.NewCipher(key); err != nil {
		return 0, err
	} else {
		cipher.XORKeyStream(bytesBuf, bytesBuf)
	}
	return binary.BigEndian.Uint32(bytesBuf), nil
}

func getVerifyCode16(number uint32) uint16 {
	s := strconv.FormatUint(uint64(number), 10) + slot
	fmt.Println("get verify code ==> ", s)
	return uint16(crc32.Checksum([]byte(s), crc32.IEEETable) % (1 << 16))
}

func checkCode16(number uint64) bool {
	verify := uint32(number>>32) & uint32(0x0000FFFF)
	number &= uint64(math.MaxUint32)
	s := strconv.FormatUint(number, 10) + slot
	if crc32.Checksum([]byte(s), crc32.IEEETable)%(1<<16) == verify {
		return true
	}
	return false
}

func getEncodeString(id, number uint32) string {
	//16位ID	16位校验码	32位随机数
	var code uint64
	rc4Number, err := getRC4By32(number)
	if err != nil {
		panic(err)
	}
	fmt.Println("rc4 number ==> ", rc4Number)
	verify := getVerifyCode16(rc4Number)
	fmt.Println(" get encode verify ==>  ", verify)
	code |= uint64(id) << (32 + 16)
	code |= uint64(verify) << 32
	code |= uint64(rc4Number)
	fmt.Println(" code ===> ", code)
	str := strings.Builder{}
	var status uint64 = 0x1F
	for i := 0; i < 12; i++ {
		str.WriteString(Dict[code&status])
		code >>= 5
	}
	str.WriteString(Dict[code&status])
	return str.String()
}

func getDict(b uint8) int {
	for i, value := range Dict {
		if string(rune(b)) == value {
			return i
		}
	}
	return -1
}

func getDecodeString(code string) {

	var number uint64
	for i := 0; i < len(code); i++ {
		status := uint64(getDict(code[i]))
		number |= status << (5 * i)
	}
	fmt.Println(" number ==> ", number)
	fmt.Println(checkCode16(number))
	fmt.Println(getRC4By32(uint32(number & uint64(math.MaxUint32))))
	fmt.Println("id ==> ", number>>48)

}
