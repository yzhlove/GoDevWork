package main

import (
	"fmt"
	"github.com/mind1949/googletrans"
)

func main() {
	detected, err := googletrans.Detect("hello googletrans")
	if err != nil {
		panic(err)
	}

	format := "language: %q, confidence: %0.2f\n"
	fmt.Printf(format, detected.Lang, detected.Confidence)
}
