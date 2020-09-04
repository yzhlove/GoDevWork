package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	services := make(map[string][]chan string, 4)

	go func() {
		for {
			for k, v := range services {
				for _, ch := range v {
					select {
					case r := <-ch:
						log.Println(" k = ", k, " v = ", r)
					default:
					}
				}
			}
			fmt.Println("====================================")
			time.Sleep(time.Second)
		}
	}()
	tags := []string{"A", "B", "C", "D"}
	for i := 0; i < 4; i++ {
		ch := make(chan string, 1)
		services[tags[i]] = append(services[tags[i]], ch)
	}

	for k, v := range services {
		for i, ch := range v {
			msg := fmt.Sprintf("%s--index:%d", k, i)
			select {
			case ch <- msg:
			default:
			}
		}
	}

	time.Sleep(5 * time.Second)

}
