package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

//错误示例

var file = "/Users/yostar/workSpace/gowork/src/GoDevWork/Minpa.jpg"
var out = "./Minpa_back.png"

func main() {

	ch := make(chan struct{})

	go func() {
		s := &server{address: ":1234"}
		fmt.Println(s.httpLis())
	}()

	go func() {
		Go()
		ch <- struct{}{}
	}()

	<-ch

}

type server struct {
	address string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open(file)
	if err != nil {
		statError(w, http.StatusInternalServerError, fmt.Sprintf("open file error:%v", err))
		return
	}
	defer f.Close()

	conn, bfw, err := w.(http.Hijacker).Hijack()
	if err != nil {
		statError(w, http.StatusInternalServerError, fmt.Sprintf("hijack error:%v", err))
		return
	}
	defer conn.Close()

	data := make([]byte, 1024)

	for {
		n, err := f.Read(data)
		if err != nil {
			if err != io.EOF {
				statError(w, http.StatusInternalServerError, fmt.Sprintf("read file error:%v", err))
			}
			return
		}
		conn.Write(data[:n])
		bfw.Flush()
	}

}

func statError(w http.ResponseWriter, code int, errInfo string) {
	w.WriteHeader(code)
	w.Write([]byte(errInfo))
}

func (s *server) httpLis() error {
	http.Handle("/down", s)
	return http.ListenAndServe(s.address, nil)
}

func Go() {

	f, err := os.Create(out)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	resp, err := http.Get("http://localhost:1234/down")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data := make([]byte, 1024)

	if resp.StatusCode == http.StatusInternalServerError {
		n, _ := resp.Body.Read(data)
		fmt.Println("read error message:", string(data[:n]))
		panic("down file error")
	}

	bfw := bufio.NewReader(resp.Body)

	for {
		n, err := bfw.Read(data)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic("read buffer error")
		}
		f.Write(data[:n])
	}

}
