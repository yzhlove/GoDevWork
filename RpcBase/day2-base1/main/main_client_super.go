package main

import (
	day2_base1 "day2-base1-example"
	"fmt"
	"log"
)

func main() {
	testSuperMain()
}

func testSuperMain() {

	address := make(chan string, 1)
	go startRpcSvc(address)
	startSuperClient(<-address)
}

func startSuperClient(address string) {
	
	rpc, err := day2_base1.NewRpcClient("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	var reply string
	if resp := rpc.Invoke("Foo.Sum", "Abc", &reply); resp != nil {
		fmt.Println("resp => ", resp)
	}
	rpc.Close()
	fmt.Println(rpc.Run())
}
