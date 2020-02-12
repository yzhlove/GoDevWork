package main

import (
	"fmt"
	"net/rpc"
	"time"
)

const KVServiceName = "path/to/pkg.KVStorageService"

func main() {

	cli, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	doClientWork(cli)
}

func doClientWork(cli *rpc.Client) {
	go func() {
		var keyChanged string
		if err := cli.Call(KVServiceName+".Watch", 3, &keyChanged); err != nil {
			panic(err)
		}
		fmt.Println("watch key => ", keyChanged)
	}()
	if err := cli.Call(KVServiceName+".Set", [2]string{"love", "qinqin"}, nil); err != nil {
		panic(err)
	}
	time.Sleep(time.Second * 5)
}
