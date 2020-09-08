package main

import (
	"fmt"
	"github.com/mind1949/googletrans"
	"golang.org/x/text/language"
)

func main() {
	params := googletrans.TranslateParams{
		Src:  "auto",
		Dest: language.SimplifiedChinese.String(),
		//Text: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. ",
		Text: "\"Hello golang\",\n\t\t\t\t\"Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.\",\n\t\t\t\t\"The Go programming language is an open source project to make programmers more productive.\",\n",
	}
	translated, err := googletrans.Translate(params)
	if err != nil {
		panic(err)
	}
	fmt.Printf("text: %q \npronunciation: %q", translated.Text, translated.Pronunciation)
}
