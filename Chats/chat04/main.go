package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Person struct {
	*User
	Number string
}

func Transform(user *User) {

	fmt.Printf("user 2=> %p \n", user)

	u := &User{Name: "yzh", Age: 20}
	user = u
	fmt.Printf("user 3=> %p \n", user)

}

func main() {

	person := &Person{User: &User{}}

	fmt.Printf("user 1=> %p \n", person.User)

	Transform(person.User)

	fmt.Printf("user 4=> %p \n", person.User)

	fmt.Println("User => ", person.User)

}
