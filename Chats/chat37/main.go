package main

import "fmt"

func main() {

	strs := []string{"hello", "world"}

	for _, str := range strs {
		fmt.Println(strs)
		strs = append(strs, str)
	}

	fmt.Println(strs)
	fmt.Println()
	temps := []string{"location", "japan"}
	for i := 0; i < len(temps); i++ {
		fmt.Println("length => ", len(temps))
		if len(temps) >= len(strs)<<1 {
			break
		}
		temps = append(temps, temps[i])
	}

	fmt.Println(temps)
}
