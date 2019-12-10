package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type AErr struct {
	err  error
	code int
}

func (e *AErr) Error() string {
	return e.err.Error() + " code:" + strconv.Itoa(e.code)
}

func (e *AErr) Unwrap() error {
	return e.err
}

func main() {

	e := errors.New("this is apple")
	ae := &AErr{err: e, code: -1}

	//err := fmt.Errorf("%w", ae)
	err := ae

	var aerr *AErr
	if errors.Is(err, ae) {
		fmt.Println("exist ok.")
	} else {
		fmt.Println("not exist .")
	}

	if errors.As(err, &aerr) {
		fmt.Println("Aerr -> ", aerr.Error())
	} else {
		fmt.Println("aerr not ok.")
	}

	for e := errors.Unwrap(err); e != nil; e = errors.Unwrap(e) {
		fmt.Println("unwrap -> ", e)
		time.Sleep(time.Second)
	}

}
