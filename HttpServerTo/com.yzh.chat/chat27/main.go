package main

import "fmt"

func main() {

	u := User{"lcmm", 20}
	fmt.Println(u)
	u.Change()
	fmt.Println(u)

}

type User struct {
	name string
	age  int
}

func (u *User) Change() {
	u.name = "yzh"
	u.age = 120
}
