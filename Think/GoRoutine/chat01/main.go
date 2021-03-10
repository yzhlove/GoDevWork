package main

import (
	"errors"
	"fmt"
	"time"
)

//select default

func ReadWithSelect(ch chan int) (x int, err error) {
	select {
	case x = <-ch:
	default:
		err = errors.New("read error")
	}
	return
}

func WriteWithSelect(ch chan int, value int) (err error) {
	select {
	case ch <- value:
	default:
		err = errors.New("write error")
	}
	return
}

func ReadNoBufferWitchSelect() {
	bufCh := make(chan int)

	if x, err := ReadWithSelect(bufCh); err != nil {
		fmt.Println("read error => ", err)
	} else {
		fmt.Println("read value => ", x)
	}
}

func ReadBufferWithSelect() {
	bufCh := make(chan int, 1)

	bufCh <- 100

	if x, err := ReadWithSelect(bufCh); err != nil {
		fmt.Println("read error => ", err)
	} else {
		fmt.Println("read value => ", x)
	}
}

func WriteNoBufferWithSelect() {
	bufCh := make(chan int)

	if err := WriteWithSelect(bufCh, 1000); err != nil {
		fmt.Println("write error => ", err)
	} else {
		fmt.Println("write success")
	}

}

func WriteBufferWithSelect() {
	bufCh := make(chan int, 1)

	if err := WriteWithSelect(bufCh, 10000); err != nil {
		fmt.Println("write error => ", err)
	} else {
		fmt.Println("write success")
	}
}

func ReadTimeoutWithSelect(ch chan int) (x int, err error) {

	timer := time.NewTimer(500 * time.Millisecond)

	select {
	case x = <-ch:
	case <-timer.C:
		err = errors.New("read time out")
	}

	return
}

func WriteTimeoutWithSelect(ch chan int, x int) (err error) {

	timer := time.NewTimer(500 * time.Millisecond)

	select {
	case ch <- x:
	case <-timer.C:
		err = errors.New("write time out")
	}

	return
}

func main() {

	ReadBufferWithSelect()
	ReadNoBufferWitchSelect()
	WriteBufferWithSelect()
	WriteNoBufferWithSelect()
	
}
