package main

import (
	"errors"
	"fmt"
)

func main() {

	e := errors.New("this is err")
	w := fmt.Errorf("new err %w", e)
	fmt.Println(errors.Unwrap(w))

}
