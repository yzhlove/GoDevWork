package main

import "fmt"

func main() {

	arr := []string{"hello", "world", "very", "good", "nice"}

	//for k, v := range arr {
	//	if v == "nice" {
	//		arr = append(arr[:k], arr[k+1:]...)
	//	}
	//}

	for i := 0; i < len(arr); i++ {
		if arr[i] == "nice" {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}

	fmt.Println(arr)



}
