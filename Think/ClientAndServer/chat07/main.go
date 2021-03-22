package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		s := server{address: ":1234"}
		fmt.Println(s.httpListen())
	}()

	time.Sleep(time.Second)

	go func() {
		defer wg.Done()
		httpDown()
	}()

	wg.Wait()

}

type server struct{ address string }

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var path = "/Users/yostar/workSpace/gowork/src/GoDevWork/1.jpg"
	f, err := os.Open(path)
	if err != nil {
		log.Println("open file error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	flusher, ok := w.(http.Flusher)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var reader = bufio.NewReader(f)
	var buffer = make([]byte, 1024)
	var cnt int

	for {
		if n, err := reader.Read(buffer); err != nil {
			log.Println("reader error:", err)
			break
		} else {
			cnt += n
			w.Write(buffer[:n])
			flusher.Flush()
		}
	}
	fmt.Println("send data size:", cnt)
}

func (s *server) httpListen() error {
	http.Handle("/down", s)
	return http.ListenAndServe(s.address, nil)
}

func httpDown() {

	f, err := os.Create("./Minpa.jpg")
	if err != nil {
		log.Println("create file error:", err)
		return
	}

	resp, err := http.Get("http://127.0.0.1:1234/down")
	if err != nil {
		log.Println("get error:", err)
		return
	}
	var reader = bufio.NewReader(resp.Body)
	var buffer = make([]byte, 1024)
	var cnt int

	for {
		if n, err := reader.Read(buffer); err != nil {
			log.Println("reader error:", err)
			break
		} else {
			cnt += n
			f.Write(buffer[:n])
		}
	}
	resp.Body.Close()
	f.Sync()
	f.Close()

	fmt.Println("receive data size:", cnt)
}
