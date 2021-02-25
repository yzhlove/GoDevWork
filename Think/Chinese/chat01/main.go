package main

import (
	"fmt"
	"regexp"
	"unicode"
)

func main() {

	str := "hello ,疯狂外星人"

	fmt.Println(IsChineseChar(str))

}

func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Han, r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}
