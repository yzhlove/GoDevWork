package main

import (
	"fmt"
	"github.com/moby/moby/pkg/pubsub"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {

	p := pubsub.NewPublisher(100*time.Millisecond, 10)
	golang := Register(p, "golang:")
	docker := Register(p, "docker:")
	all := p.Subscribe()
	exit := make(chan bool)

	go func() {
		strs := [...]string{"hi:", "golang:", "docker:"}
		for i := 0; i < 10; i++ {
			msg := strs[rand.Intn(3)] + " send to message by " + strconv.Itoa(i)
			p.Publish(msg)
			time.Sleep(time.Second)
		}
		close(exit)
	}()

	go func() {
		for s := range golang {
			fmt.Println("[GOLANG] ", s)
		}
	}()
	go func() {
		for s := range docker {
			fmt.Println("[DOCKER] ", s)
		}
	}()
	go func() {
		for s := range all {
			fmt.Println("[ALL] ", s)
		}
	}()
	<-exit
}

func Register(p *pubsub.Publisher, tag string) <-chan interface{} {
	return p.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, tag) {
				return true
			}
		}
		return false
	})
}
