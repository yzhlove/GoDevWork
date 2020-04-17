package helper

import (
	"WorkSpace/GoDevWork/GiftServerTwo/config"
	"hash/crc32"
	"strconv"
)

const (
	check = 0xFFF
	slot  = "*#06#"
)

var (
	code   = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	length = uint64(len(code))
)

func GetBucketTop(code string) int {
	return int(crc32.ChecksumIEEE([]byte(code))) % config.BucketMax
}

func ToEncodeStr(number uint64) string {
	newCode := make([]rune, 0, 11)
	for number >= length {
		newCode = append(newCode, code[number%length])
		number /= length
	}
	newCode = append(newCode, code[number])
	return string(newCode)
}

func ToDecodeStr(code string) (number uint64) {
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
		number = number*length + post(runes[i-1])
	}
	return
}

func encodeId(id uint32) uint32 {
	return id | uint32(1<<11)
}

func decodeId(id uint32) uint32 {
	return id & uint32(1<<11-1)
}

func ToEncodeNumber(id uint32, number int64) uint64 {
	id = encodeId(id)
	verifyCode := crc32.ChecksumIEEE([]byte(strconv.FormatInt(number, 10) + slot))
	return (uint64(id) << 54) | uint64(number<<12) | uint64(verifyCode%check)
}

func ToDecodeNumber(number uint64) (id uint32, ok bool) {
	id = decodeId(uint32(number >> 54))
	checkCode := uint32(number & check)
	source := (number >> 12) & (1<<42 - 1)
	verifyCode := crc32.ChecksumIEEE([]byte(strconv.FormatUint(source, 10) + slot))
	ok = verifyCode%check == checkCode
	return
}
