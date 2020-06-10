package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

func Test_Client(t *testing.T) {

	/*	conn, err := net.Dial("tcp", ":1234")
		if err != nil {
			panic(err)
		}

		reader := bufio.NewReader(conn)
		for {
			line, isPrefix, err := reader.ReadLine()
			if err != nil {
				zlog.Println(err)
				return
			}
			zlog.Println("line =>", string(line), " isPrefix => ", isPrefix)
			time.Sleep(time.Second * 5)
		}*/

}

func Test_Client2(t *testing.T) {

	var wg sync.WaitGroup
	for i := 0; i < 255; i++ {
		wg.Add(1)
		go testClient(i+1, &wg)
	}
	wg.Wait()
}

func testClient(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		log.Println(err)
		return
	}
	reader := bufio.NewReader(conn)
	for {
		if line, _, err := reader.ReadLine(); err != nil {
			log.Println(err)
		} else {
			log.Println(fmt.Sprintf("[%d] result:%v", id, string(line)))
		}
		time.Sleep(5 * time.Second)
	}
}
