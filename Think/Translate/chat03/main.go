package main

import (
	"context"
	"fmt"
	"github.com/mind1949/googletrans"
	"golang.org/x/text/language"
)

func main() {
	params := func() <-chan googletrans.TranslateParams {
		params := make(chan googletrans.TranslateParams)
		go func() {
			defer close(params)
			texts := []string{
				"Hello golang",
				"Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.",
				"The Go programming language is an open source project to make programmers more productive.",
			}
			for i := 0; i < len(texts); i++ {
				params <- googletrans.TranslateParams{
					Src:  "auto",
					Dest: language.SimplifiedChinese.String(),
					Text: texts[i],
				}
			}
		}()
		return params
	}()

	for transResult := range googletrans.BulkTranslate(context.Background(), params) {
		if transResult.Err != nil {
			panic(transResult.Err)
		}
		fmt.Printf("text: %q, pronunciation: %q\n", transResult.Text, transResult.Pronunciation)
	}
}
