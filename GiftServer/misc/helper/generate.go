package helper

import (
	"math/rand"
	"time"
)

const (
	maxNumber = 1 << 42
)

var seed = time.Now().Unix()

func getEngine() *rand.Rand {
	return rand.New(rand.NewSource(seed))
}

func GenerateCode(id, count uint32) (chan int64, error) {
	e := getEngine()
	generate := make(chan int64, 256)
	status := make(map[int64]struct{}, 128)
	go func() {
		var index uint32
		for {
			if index >= count {
				break
			}
			number := e.Int63() % maxNumber
			if _, ok := status[number]; ok {
				continue
			}
			status[number] = struct{}{}
			generate <- number
			index++
		}
	}()
	return generate, nil
}
