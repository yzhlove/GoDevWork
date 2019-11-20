package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {

	userList := make([]User, 0, 10)

	for i := 0; i < 10; i++ {
		user := User{
			Name: fmt.Sprintf("name:%d", i),
			Age:  i + 1,
		}
		userList = append(userList, user)
	}
	fmt.Println("userList => ", userList)
}
