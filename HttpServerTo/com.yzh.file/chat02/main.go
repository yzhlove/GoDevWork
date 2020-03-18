package main

import "fmt"

var (
	_code    = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	_codelen = uint64(len(_code))
)

func main() {

	var number uint64 = 1234

	str := Encode(number)
	//str := IdToCode(number)
	fmt.Println("str => ", str)

	n := Decode(str)
	//n := CodeToID(str)
	fmt.Println("number => ", n)

}

func Encode(number uint64) string {
	code := make([]rune, 0, 11)
	for number >= _codelen {
		code = append(code, _code[number%_codelen])
		number /= _codelen
	}
	code = append(code, _code[number])
	return string(code)
}

func Decode(code string) (number uint64) {
	runes := []rune(code)
	post := func(c rune) uint64 {
		switch {
		case c < 58:
			return uint64(c - 48)
		case c < 91:
			return uint64(c - 65 + 10)
		default:
			return uint64(c - 97 + 36)
		}
	}
	for length := len(runes); length > 0; length-- {
		number = number*_codelen + post(runes[length-1])
	}
	return
}
