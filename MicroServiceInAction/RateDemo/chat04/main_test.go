package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const timeout = 5 * time.Second

func TestLimiterListener(t *testing.T) {
	const max = 5
	const attempts = 2048
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	l = LimiterListener(l, max)

	var open int32
	go http.Serve(l, http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if n := atomic.AddInt32(&open, 1); n > max {
			t.Errorf("%d open connections ,want <= %d", n, max)
		}
		defer atomic.AddInt32(&open, -1)
		time.Sleep(10 * time.Millisecond)
		fmt.Fprint(writer, "some body")
	}))

	var wg sync.WaitGroup
	var filed int32

	for i := 0; i < attempts; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c := http.Client{Timeout: 3 * time.Second}
			r, err := c.Get("http://" + l.Addr().String())
			if err != nil {
				t.Log(err)
				atomic.AddInt32(&filed, 1)
				return
			}
			defer r.Body.Close()
			io.Copy(ioutil.Discard, r.Body)
		}()
	}
	wg.Wait()
	t.Log("field ==> ", filed)
}

var errFake = errors.New("fake error from errorListener")

type errorListener struct {
	net.Listener
}

func (errorListener) Accept() (net.Conn, error) {
	return nil, errFake
}

func TestLimiterListenerError(t *testing.T) {
	errChan := make(chan error, 1)
	go func() {
		defer close(errChan)
		const n = 2
		listener := LimiterListener(errorListener{}, n)
		for i := 0; i < n+1; i++ {
			if _, err := listener.Accept(); err != errFake {
				errChan <- fmt.Errorf("Accept error= %v ", err)
				return
			}
		}
	}()

	select {
	case err := <-errChan:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(timeout):
		t.Fatal("timeout. deadlock?")
	}
}

func TestLimiterListenerClose(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Close()

	listener = LimiterListener(listener, 1)

	errChan := make(chan error)
	go func() {
		defer close(errChan)
		c, err := net.DialTimeout("tcp", listener.Addr().String(), timeout)
		if err != nil {
			errChan <- err
			return
		}
		c.Close()
	}()

	c, err := listener.Accept()
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	err = <-errChan
	if err != nil {
		t.Fatalf("DialTimeout: %v ", err)
	}

	acceptDone := make(chan struct{})
	go func() {
		c, err := listener.Accept()
		if err == nil {
			c.Close()
			t.Errorf("Unexcepted successful Accept()")
		}
		close(acceptDone)
	}()

	time.Sleep(10 * time.Millisecond)
	listener.Close()

	select {
	case <-acceptDone:
		t.Log("accept done.")
	case <-time.After(timeout):
		t.Fatalf("Accept() still blocking")
	}
}
