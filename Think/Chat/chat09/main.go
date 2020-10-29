package main

import (
	"fmt"
	"os/user"
)

// os/user

func main() {

	if u, err := user.Current(); err == nil {
		fmt.Println(u.Name)
		fmt.Println(u.Uid)
		fmt.Println(u.Gid)
		fmt.Println(u.HomeDir)
		fmt.Println(u.Username)
		fmt.Println(u.GroupIds())
	}

}
