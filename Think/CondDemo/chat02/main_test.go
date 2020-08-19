package chat02

import (
	"fmt"
	"sync"
	"testing"
)

func TestNewSpinLocker(t *testing.T) {

	l := NewSpinLocker()
	var wg sync.WaitGroup
	var count int
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			l.Lock()
			count++
			t.Log("count ==> ", count)
			l.Unlock()
		}()
	}
	wg.Wait()
	t.Log("count ===> ", count)
}

func BenchmarkNewSpinLocker(b *testing.B) {
	benchmarkFunc(NewSpinLocker(), b.N)
}

func BenchmarkNewMutexLock(b *testing.B) {
	benchmarkFunc(&sync.Mutex{}, b.N)
}

func BenchmarkNewRWMutex(b *testing.B) {
	benchmarkFunc(&sync.RWMutex{}, b.N)
}

func benchmarkFunc(locker sync.Locker, count int) {
	var index int
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			locker.Lock()
			index++
			locker.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("count = ", index)
}
