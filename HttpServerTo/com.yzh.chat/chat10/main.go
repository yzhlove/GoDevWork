package main

import "fmt"

func main() {

	type User struct {
		Name string
		Age  uint32
	}

	statusA := make(map[User]uint32, 8)
	statusB := make(map[*User]uint32, 8)

	a := User{Name: "abc", Age: 16}
	b := User{Name: "abc", Age: 16}

	statusA[a] = 1
	statusA[b] = 2

	fmt.Println("status A length => ", len(statusA))

	statusB[&a] = 1
	statusB[&b] = 2

	fmt.Println("status B length => ", len(statusB))

	var tv map[uint32]struct{}

	if tv == nil {
		fmt.Println("waht are you doing ???")
	}

}
