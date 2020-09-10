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
		Text: "1. IndexCountPerInstance: The number of indices that will be used in this draw call that defines one instance. This need not be every index in the index buffer; that is, you can draw a contiguous subset of indices.",
	}
	translated, err := googletrans.Translate(params)
	if err != nil {
		panic(err)
	}
	fmt.Printf("text: %q \npronunciation: %q", translated.Text, translated.Pronunciation)
}
