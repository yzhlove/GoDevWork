package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var Head = []byte("HTTP/1.1 200 OK\r\n\r\n")

func main() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/loop", Loop)
	log.Println(http.ListenAndServe(":1234", nil))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	if jack, ok := w.(http.Hijacker); ok {
		if conn, buffer, err := jack.Hijack(); err == nil {
			go Event(conn, buffer)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("get http connection error"))
}

func Event(conn net.Conn, buffer io.ReadWriter) {
	defer conn.Close()
	sb := strings.Builder{}
	sb.Write(Head)
	sb.WriteString("server to client message:\n")
	sb.WriteString("not read client message!")

	if _, err := conn.Write([]byte(sb.String())); err != nil {
		log.Println("buffer writer message error: ", err)
	}
}

func Loop(w http.ResponseWriter, r *http.Request) {
	if hijacker, ok := w.(http.Hijacker); ok {
		if conn, buffer, err := hijacker.Hijack(); err == nil {
			go LoopEvent(conn, buffer)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("hijack error"))
}

func LoopEvent(conn net.Conn, buffer io.ReadWriter) {
	defer conn.Close()
	var chQueue = make(chan string, 1)

	go func() {
		defer close(chQueue)
		var count int
		for range time.NewTicker(time.Second).C {
			count++
			if count > 10 {
				break
			}
			msg := "server to client message : " + strconv.Itoa(count) + "\n"
			chQueue <- msg
		}
	}()

	for msg := range chQueue {
		log.Println("send message:", msg)
		if _, err := conn.Write(append(Head, []byte(msg)...)); err != nil {
			log.Println("send message error:", err)
		}
	}
}
