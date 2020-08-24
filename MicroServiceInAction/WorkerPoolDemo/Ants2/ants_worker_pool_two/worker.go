package ants_worker_pool_two

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

type GoWorker struct {
	p        *Pool
	taskChan chan func()
	expire   time.Time
}

func (w *GoWorker) run() {
	w.p.Inc()
	go func() {
		defer func() {
			w.p.Dec()
			w.p.cache.Put(w)
			if x := recover(); x != nil {
				trace(fmt.Sprint(x))
			}
		}()
		for fn := range w.taskChan {
			if fn == nil {
				panic("running func is nil")
			}
			fn()
			if ok := w.p.Put(w); !ok {
				panic("put worker is false")
			}
		}
	}()
}

func ClosedChan(task chan func()) {
	task <- nil
}

func trace(head string) string {
	var stack [32]uintptr
	i := runtime.Callers(3, stack[:])
	sb := strings.Builder{}
	sb.WriteString("{" + head + "} trace back:")
	for _, pc := range stack[:i] {
		f := runtime.FuncForPC(pc)
		file, line := f.FileLine(pc)
		sb.WriteString(fmt.Sprintf("\n  \033[32m↓\033[0m [\033[31m%s\033[0m] %s:%d", f.Name(), file, line))
	}
	return sb.String()
}
