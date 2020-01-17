package main

import (
	"fmt"
	"math/rand"
	"sync"
)

//随机数生成引擎

func main() {

	engine := getEngine(123456)
	//generateNumber(engine)
	generateConcurrencyNumber(engine)

	/*
			number ==>  3550330175404308690
			number ==>  5603678272308080255
			number ==>  4470361554766394783
			number ==>  8531737684161497558
			number ==>  1220792738559528428
			number ==>  1660522245406902871
			number ==>  8197558958854837280
			number ==>  6152978149681228875
			number ==>  1205949487052661650
			number ==>  8599742499877181643

		5603678272308080255
		3550330175404308690
		4470361554766394783
		8531737684161497558
		1220792738559528428
		1660522245406902871
		8197558958854837280
		6152978149681228875
		1205949487052661650
		8599742499877181643

	*/

}

//随机引擎
func getEngine(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}

func generateNumber(engine *rand.Rand) {

	count := 0

	for count < 10 {
		number := engine.Int63()
		fmt.Println("number ==> ", number)
		count++
	}

}

func generateConcurrencyNumber(engine *rand.Rand) {

	var wg sync.WaitGroup
	msgChan := make(chan int64)
	go func() {
		for msg := range msgChan {
			fmt.Println(msg)
		}
	}()

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			msgChan <- engine.Int63()
		}()
	}
	wg.Wait()
	close(msgChan)
}
