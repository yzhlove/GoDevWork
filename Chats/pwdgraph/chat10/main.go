package main

import (
	"crypto/rc4"
	"encoding/binary"
	"fmt"
	"hash/crc32"
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
	fmt.Println("str ==> ", str)
}

func getRc4Encode32(number uint32) (uint32, error) {
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
	return uint16(crc32.Checksum([]byte(s), crc32.IEEETable) % (1 << 16))
}

func getEncodeString(id, number uint32) string {
	//16位ID	16位校验码	32位随机数
	var code uint64
	rc4Number, err := getRc4Encode32(number)
	if err != nil {
		panic(err)
	}
	verify := getVerifyCode16(rc4Number)

	code |= uint64(id) << (32 + 16)
	code |= uint64(verify) << 32
	code |= uint64(number)

	str := strings.Builder{}
	var status uint64 = 0x1F
	for i := 0; i < 12; i++ {
		str.WriteString(Dict[code&status])
		code >>= 5
	}
	str.WriteString(Dict[code&status])
	return str.String()
}

func getDecodeString(code string) {

}
