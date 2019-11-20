package main

import "fmt"

var (
	_code   = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	_length = uint64(len(_code))
)

func main() {

	//var uid uint64 = 1193826606532005917
	//fmt.Println(IdToCode(98))

	fmt.Println(CodeToId("a1"))

}

func IdToCode(id uint64) string {
	fmt.Println("length => ", _length)
	var code []rune
	for id >= _length {
		fmt.Printf("id -> %d \n", id)
		idx := id % _length
		fmt.Printf("idx -> %d \n", idx)
		id /= _length
		fmt.Printf("id -> %d \n", id)
		code = append(code, _code[idx])
		fmt.Println()
	}
	code = append(code, _code[id])
	return string(code)
}

func CodeToId(code string) (id uint64) {
	runes := []rune(code)
	l := len(runes)
	for l > 0 {
		fmt.Printf("l = %d id = %d length = %d rune = %c %d \n", l, id, _length, runes[l-1], runes[l-1])
		id = id*_length + runeIdx(runes[l-1])
		fmt.Printf("newid => %d\n", id)
		fmt.Println()
		l -= 1
	}
	return
}

func runeIdx(r rune) uint64 {
	if r < 58 { //0-9
		return uint64(r - 48)
	}
	if r < 91 { //A-Z
		return uint64(r - 65 + 10)
	}
	return uint64(r - 97 + 36)
}
