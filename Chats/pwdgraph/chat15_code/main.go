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

//兑换码加密解密

var (
	key  = []byte("*#1008611#*") //rc4
	slot = "*#06#"               //颜值
	dict = [...]rune{ //字典
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
		'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R',
		'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'2', '3', '4', '5', '6', '7', '8', '9',
	}
)

func main() {

	var number uint32 = 223344
	var id uint32 = 1000

	code, err := generateEncodeCode(generateRedisCode(id, number))
	if err != nil {
		panic(err)
	}
	fmt.Println("code => ", code)

}

//rc4加密解密	(兑换码服务器实现)
func generateRC4Code(number uint32) (uint32, error) {
	buffers := make([]byte, 4)
	binary.BigEndian.PutUint32(buffers, number)
	if cipher, err := rc4.NewCipher(key); err != nil {
		return 0, err
	} else {
		cipher.XORKeyStream(buffers, buffers)
	}
	return binary.BigEndian.Uint32(buffers), nil
}

//生成校验码	(兑换码服务器实现)
func generateVerifyCode(number uint32) uint16 {
	str := strconv.FormatUint(uint64(number), 10) + slot
	return uint16(crc32.ChecksumIEEE([]byte(str)) % (1 << 16))
}

//验证校验码是否正确	(兑换码服务器实现)
func checkVerifyCode(number uint64) bool {
	//获取校验码
	verify := uint32(number>>32) & uint32(0x0000FFFF)
	number &= uint64(math.MaxUint32)
	str := strconv.FormatUint(number, 10) + slot
	return crc32.ChecksumIEEE([]byte(str))%(1<<16) == verify
}

//验证兑换码的长度是否合法	(game服务器实现)
func checkVerifyCodeLength(code string) bool {
	return len(code) == 13
}

//生成带批次的随机码 (Redis存储)
func generateRedisCode(id, number uint32) uint64 {
	return (uint64(id) << 32) | uint64(number)
}

//生成兑换码	(兑换码服务器实现)
func generateEncodeCode(number uint64) (string, error) {
	//兑换码组成:	16位ID	16位校验码	32位随机数
	var code uint64
	//rc4code
	rc4code, err := generateRC4Code(uint32(number))
	if err != nil {
		return "", err
	}
	//校验码
	verifyCode := generateVerifyCode(rc4code)
	//组装uint64
	code |= (number & (uint64(math.MaxUint32) << 32)) << 16
	code |= uint64(verifyCode) << 32
	code |= uint64(rc4code)
	//根据字典生成兑换码
	str := strings.Builder{}
	for i := 0; i < 12; i++ {
		str.WriteRune(dict[code&uint64(0x1F)])
		code >>= 5
	}
	str.WriteRune(dict[code&uint64(0x1F)])
	return str.String(), nil
}


