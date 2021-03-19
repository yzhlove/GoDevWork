package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	//for i := 0; i < 1000; i++ {
	//	test1()
	//}

	test2()

}

func test1() {

	ctx, cancel := context.WithTimeout(context.Background(), 0)
	defer cancel()
	select {
	case <-ctx.Done():
		fmt.Println("done.")
	default:
		fmt.Println("default")
	}

}

func test2() {
	
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)

	go func() {
		fmt.Println("start ...")
		time.Sleep(6 * time.Second)
		fmt.Println("end")
		cancel()
	}()
	<-ctx.Done()
	fmt.Println("err -> ", ctx.Err())
}
