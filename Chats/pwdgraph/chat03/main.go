package main

import (
	"fmt"
	"math/rand"
)

//
const DICT_SIZE = 32
const AWARD_CODE_BIT = 5
const AWARD_CODE_NUM = 7

var AwardCodeDict = []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J',
	'K', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'U',
	'V', 'W', 'X', 'Y', 'Z',
	'1', '2', '3', '4', '5', '6', '7', '8', '9'}

type UINT64 uint64

func main() {
	var awardId = 55
	for i := 0; i < 10; i++ {
		codeVal := GenerateAwardCodeVaule(UINT64(awardId))
		strs := DecodeAwardCodeValue(codeVal)
		fmt.Print(codeVal)
		fmt.Print(" ", string(strs))
		fmt.Print(" ", GetAwardCodeValue(strs))
		fmt.Print(" ", GetAwardID(strs))
		fmt.Println()
	}
}

func GenerateAwardCodeVaule(awardId UINT64) UINT64 {
	var codeVal UINT64
	for i := 0; i < AWARD_CODE_NUM; i++ {
		key := UINT64(rand.Int63() % DICT_SIZE)
		codeVal |= key << (AWARD_CODE_BIT * i)
	}
	codeVal |= awardId << (AWARD_CODE_BIT * AWARD_CODE_NUM)
	return codeVal
}

func DecodeAwardCodeValue(codeVal UINT64) []rune {
	strs := make([]rune, 0, 16)
	for ; codeVal > 0; {
		key := codeVal & 0x1F
		strs = append(strs, AwardCodeDict[key])
		codeVal = codeVal >> AWARD_CODE_BIT
	}
	return strs
}

func GetAwardID(strs []rune) UINT64 {
	if strs == nil || len(strs) <= AWARD_CODE_NUM {
		return 0
	}
	var awardId UINT64
	for i := AWARD_CODE_NUM; i < len(strs); i++ {
		val := UINT64(strs[i])
		awardId |= val << (AWARD_CODE_BIT * (i - AWARD_CODE_NUM))
	}
	return awardId
}

func GetAwardCodeValue(strs []rune) UINT64 {
	if strs == nil {
		return 0
	}
	var codeVal UINT64
	for i := len(strs) - 1; i >= 0; i-- {
		key := UINT64(strs[i])
		codeVal |= key << (AWARD_CODE_BIT * i)
	}
	return codeVal
}
