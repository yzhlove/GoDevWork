package helper

import (
	"context"
	"math/rand"
)

func GenerateCode(ctx context.Context, number uint32) chan int64 {

	codeChan := make(chan int64, 256)
	status := make(map[int64]struct{}, 128)

	go func() {
		defer close(codeChan)
		var count uint32
		for count < number {
			num := rand.Int63() >> 22
			if _, ok := status[num]; ok {
				continue
			}
			status[num] = struct{}{}
			select {
			case <-ctx.Done():
				return
			default:
				codeChan <- num
				count++
			}
		}
	}()
	return codeChan
}
