package helper

import (
	"hash/crc32"
	"strconv"
)

const (
	slot = "*#@123456"
	code = 0xFFF
)

var (
	_code    = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	_codelen = uint64(len(_code))
)

func Encode(id int, num uint64) uint64 {
	verifyCode := crc32.ChecksumIEEE([]byte(strconv.FormatUint(num, 10) + slot))
	return (uint64(id) << 54) | (num << 12) | uint64(verifyCode%code)
}

func Decode(num uint64) (id uint32, ok bool) {
	id = uint32(num >> 54)
	checkCode := uint32(num & code)
	source := (num >> 12) & (1<<42 - 1)
	verifyCode := crc32.ChecksumIEEE([]byte(strconv.FormatUint(source, 10) + slot))
	ok = verifyCode%code == checkCode
	return
}

func IdToCode(id uint64) string {
	var code []rune
	for id >= _codelen {
		idx := id % _codelen
		id /= _codelen
		code = append(code, _code[idx])
	}
	code = append(code, _code[id])
	return string(code)
}

func CodeToID(code string) (id uint64) {
	runes := []rune(code)
	l := len(runes)
	for l > 0 {
		id = id*_codelen + rune2Idx(runes[l-1])
		l -= 1
	}
	return
}

func rune2Idx(r rune) uint64 {
	if r < 58 { //0-9
		return uint64(r - 48)
	}
	if r < 91 { //A-Z
		return uint64(r - 65 + 10)
	}
	return uint64(r - 97 + 36)
}
