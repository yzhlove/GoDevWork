package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"go.uber.org/atomic"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	//TestTimeOut()
	//MaxConcurrent()
	//MaxConcurrentAsync()
	//MaxConcurrentAsyncDo()
	//MaxRequest()
	MaxErrorCount()
}

func init() {
	hystrix.Configure(map[string]hystrix.CommandConfig{
		"time-out-cfg":       {Timeout: 1000},             //超时时间 ,1000ms
		"max-concurrent-cfg": {MaxConcurrentRequests: 2},  //最大并发数，超过并发返回错误
		"max-requests-cfg":   {RequestVolumeThreshold: 4}, //请求数量阀值，用这些数量的请求来计算阀值
		"max-errors-cfg":     {ErrorPercentThreshold: 25}, //错误数量阀值，达到阀值，启动熔断
		"sleeps-restore-cfg": {SleepWindow: 1000},         //熔断尝试恢复时间
	})
}

func TestTimeOut() {

	if err := hystrix.Do("time-out-cfg", func() error {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		select {
		case <-ctx.Done():
			return errors.New("take api test time out ")
			//default:
			//	return nil
		}
	}, func(err error) error {
		return errors.New("call err:" + err.Error())
	}); err != nil {
		log.Println("run err:", err)
	}
}

func MaxConcurrent() {

	for i := 0; i < 100; i++ {
		if err := hystrix.Do("max-concurrent-cfg", func() error {
			log.Println("succeed")
			return nil
		}, func(err error) error {
			log.Println("run err:", err)
			return err
		}); err != nil {
			log.Println("why ? err:", err)
		}
	}

}

func MaxConcurrentAsync() {

	errCh := make(chan error, 16)
	succeedCh := make(chan struct{}, 16)

	go func() {
		var index, count int

		for {
			select {
			case err := <-errCh:
				index++
				if err != nil {
					log.Println("Err:", err)
				}
				fmt.Println("error count = ", index)
			case <-succeedCh:
				count++
				fmt.Println("succeed count = ", count)
			}
		}
	}()

	for i := 0; i < 100; i++ {
		returnErrCh := hystrix.Go("max-concurrent-cfg", func() error {
			succeedCh <- struct{}{}
			return nil
		}, func(err error) error {
			return err
		})
		go func(e chan error) {
			select {
			case err := <-e:
				if err != nil {
					errCh <- err
				}
			}
		}(returnErrCh)
		time.Sleep(10 * time.Millisecond)
	}

	time.Sleep(5 * time.Second)

}

func MaxConcurrentAsyncDo() {

	errChan := make(chan error, 16)

	for i := 0; i < 100; i++ {
		go func() {
			if err := hystrix.Do("max-concurrent-cfg", func() error {
				return nil
			}, func(err error) error {
				return err
			}); err != nil {
				errChan <- err
			}
		}()
	}

	go func() {

		var count int

		for err := range errChan {
			if err != nil {
				count++
				log.Println("count = ", count, " Err:", err)
			}
		}
	}()

	time.Sleep(5 * time.Second)
}

func MaxRequest() {

	errChan := make(chan error, 16)

	go func() {
		for err := range errChan {
			log.Println("err --> ", err)
		}
	}()

	for i := 0; i < 20; i++ {
		go func(i int) {
			if err := hystrix.Do("max-requests-cfg", func() error {
				log.Println("succeed i = ", i)
				return nil
			}, func(err error) error {
				log.Println("error i = ", i)
				return errors.New("index:" + strconv.Itoa(i) + ">>" + err.Error())
			}); err != nil {
				errChan <- err
			}
		}(i)
	}

	time.Sleep(5 * time.Second)

}

func MaxErrorCount() {

	var succeed_count, failed_count atomic.Int32

	for i := 0; i < 100; i++ {
		go func(i int) {
			hystrix.Do("max-errors-cfg", func() error {
				if rand.Int()&1 == 1 {
					return errors.New("return error")
				}
				succeed_count.Add(1)
				return nil
			}, func(err error) error {
				failed_count.Add(1)
				return err
			})
		}(i)
	}

	time.Sleep(time.Second)
	log.Printf("succed %v failed %v \n", succeed_count.Load(), failed_count.Load())
}


