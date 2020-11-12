package main

import (
	"errors"
	"fmt"
)

var _NotErr = &InsultError{err: "what are you doing", target: 12138}

func main() {

	var notErr *InsultError
	if err := Get(); err != nil {
		if errors.Is(err, _NotErr) {
			fmt.Println("err is ok")
			if errors.As(err, &notErr) {
				fmt.Println("err target => ", notErr.target, notErr.err)
			} else {
				fmt.Println("error as err")
			}
		} else {
			fmt.Println("errors is err")
		}
	} else {
		fmt.Println("err is nil")
	}

}

type InsultError struct {
	err    string
	target uint32
}

func (ie *InsultError) Error() string {
	return ie.err
}

func Get() error {
	return _NotErr
}
