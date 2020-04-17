package main

import "fmt"

func main() {

	testMap([]string{"hello", "world"})

}

func testMap(parts []string) (partMap map[string]string) {
	partMap = map[string]string{}
	for index, value := range parts {
		partMap[fmt.Sprintf("%d", index)] = value
	}
	return
}
