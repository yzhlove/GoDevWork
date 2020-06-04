package main

import (
	"errors"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	MaxQueueNumber = 5
	PoolSize       = 2
)

var GenID int32

type DBConn struct {
	ID int32
}

func create() (io.Closer, error) {
	id := atomic.AddInt32(&GenID, 1)
	return &DBConn{ID: id}, nil
}

func dbQuery(query int, pool *Pool) {
	if conn, err := pool.Acquire(); err != nil {
		log.Println("acquire conn err:", err)
		return
	} else {
		defer pool.Release(conn)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		log.Println("db conn :", query)
	}
}

func (db *DBConn) Close() error {
	log.Println("db closed")
	return nil
}

func main() {

	var wg sync.WaitGroup
	wg.Add(MaxQueueNumber)

	p, err := NewPool(PoolSize, create)
	if err != nil {
		panic(err)
	}
	for i := 0; i < MaxQueueNumber; i++ {
		go func(query int) {
			dbQuery(query, p.(*Pool))
			wg.Done()
		}(i)
	}
	wg.Wait()
	log.Println("close pool")
	p.Close()
}

type PoolFunc func() (closer io.Closer, err error)

type Pool struct {
	mutex   sync.Mutex
	res     chan io.Closer
	factory PoolFunc
	closed  bool
}

func NewPool(size uint, fn PoolFunc) (io.Closer, error) {
	if size == 0 {
		return nil, errors.New("size is no zero")
	}
	return &Pool{
		factory: fn,
		res:     make(chan io.Closer, size),
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.res:
		if ok {
			return r, nil
		}
		return nil, errors.New("pool is closed")
	default:
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.closed {
		r.Close()
		return
	}
	select {
	case p.res <- r:
	default:
		r.Close()
	}
}

func (p *Pool) Close() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if !p.closed {
		p.closed = true
		close(p.res)
		for r := range p.res {
			if err := r.Close(); err != nil {
				return err
			}
		}
	}
	return nil
}
