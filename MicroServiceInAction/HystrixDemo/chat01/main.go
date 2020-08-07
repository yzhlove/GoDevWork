package main

import (
	"errors"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"log"
	"net/http"
)

func main() {

	out := make(chan struct{}, 1)

	c := hystrix.Go("get_baidu", func() error {
		if _, err := http.Get("http://www.baidu.com"); err != nil {
			return err
		}
		out <- struct{}{}
		return nil
	}, func(err error) error {
		fmt.Println("get an err ,handle it")
		return errors.New("handle error")
	})

	select {
	case <-out:
		log.Println("succeed.")
	case err := <-c:
		if err != nil {
			log.Printf("hystrix err: %v \n", err)
		}
	}
}
