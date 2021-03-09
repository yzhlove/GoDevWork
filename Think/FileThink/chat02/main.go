package main

import (
	"fmt"
	"strings"
)

func main() {

	paths := []string{
		"./tables",
		"/a/b/c/tables",
		"./tables/a.xlsx",
		"/a/b/c/tables/d.xlsx",
		"./tables/",
		"/a/b/c/tables/",
	}

	for _, path := range paths {
		if idx := strings.Index(path, "tables"); idx != -1 {
			target := path[:idx]
			value := path[:idx-1]
			fmt.Println(path, " -> ", target, " -> ", value)
		}
	}

}
