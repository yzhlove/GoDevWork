package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/hello", Hello)
	log.Println(http.ListenAndServe(":1234", nil))

}

func Hello(w http.ResponseWriter, r *http.Request) {

	if hack, ok := w.(http.Hijacker); ok {
		if conn, buffer, err := hack.Hijack(); err != nil {
			log.Println("Hijack err:", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		} else {
			go func(conn net.Conn, buffer io.ReadWriter) {
				//defer conn.Close()
				var sb strings.Builder

				//var buf = make([]byte, 1024)
				//if n, err := conn.Read(buf); err != nil {
				sb.WriteString("no read message to http client")
				//} else {
				//	str := string(buf[:n])
				//	sb.WriteString(str)
				//	log.Println(conn.LocalAddr(), " read msg:", str)
				//}
				conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
				if n, err := conn.Write([]byte(strings.ToUpper(sb.String()))); err != nil {
					log.Println("buffer write msg error:", err)
				} else {
					log.Println(" n = ", n)
				}
				conn.Write([]byte(" ababbasbdbasbdbas "))
				//conn.Close()
			}(conn, buffer)
		}
	}
}
