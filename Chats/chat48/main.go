package main

import "fmt"

func main() {

	arr := make([]string, 0, 4)
	fmt.Printf("len %d cap %d \n", len(arr), cap(arr))

	arr = append(arr, "hello")
	arr = append(arr, "world")
	arr = append(arr, "yes")
	arr = append(arr, "no")

	fmt.Printf("len %d cap %d \n", len(arr), cap(arr))

	for i, v := range arr {
		if v == "no" || v == "world" {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}

	fmt.Printf("len %d cap %d \n", len(arr), cap(arr))

}
