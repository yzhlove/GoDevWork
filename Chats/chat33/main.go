package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {

	userList := []User{
		{"yzh", 16}, {"lcm", 19}, {"xjj", 21},
	}

	for _, u := range userList {
		fmt.Println("Main User => ", Show(&u))
	}

}

func Show(u *User) *User {
	user := new(User)
	user.Name = u.Name
	user.Age = u.Age
	fmt.Println("Show User => ", user)
	return user
}
