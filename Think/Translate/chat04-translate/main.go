package main

import (
	"encoding/json"
	"github.com/mind1949/googletrans"
	"golang.org/x/text/language"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
)

var head = "HTTP/1.1 200 OK\r\nContent-Type: application/json;charset=UTF-8\r\n\r\n"

func main() {
	s := &server{}
	s.cache.New = func() interface{} {
		return new(googletrans.TranslateParams)
	}
	s.packetQueue = make(chan *packet, 512)
	go s.translateService()
	log.Println("⌘ -服务器启动,在本地[1234]端口监听!")
	if err := http.ListenAndServe(":1234", s); err != nil {
		panic(err)
	}
}

type server struct {
	packetQueue chan *packet
	cache       sync.Pool
}

type packet struct {
	source string
	conn   net.Conn
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	result := r.FormValue("translate")
	log.Println("==> data:", result)

	if jack, ok := w.(http.Hijacker); ok {
		if conn, _, err := jack.Hijack(); err == nil {
			s.packetQueue <- &packet{conn: conn, source: result}
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("消息处理失败，请稍后重试~"))
}

func (s *server) translateService() {
	for pt := range s.packetQueue {
		if pt == nil || len(strings.TrimSpace(pt.source)) == 0 || pt.conn == nil {
			log.Println("source is empty!")
			continue
		}
		if res, ok := s.cache.Get().(*googletrans.TranslateParams); ok {
			if result, err := translate(res, pt.source); err != nil {
				send(pt.conn, err.Error())
			} else {
				send(pt.conn, result)
				s.cache.Put(res)
			}
		}
	}
}

func send(conn net.Conn, msg string) {
	defer conn.Close()
	var sb strings.Builder
	sb.WriteString(head)
	d := toJson(msg)
	sb.Write(d)
	log.Println("服务器返回的消息:↓\n" + string(d))
	if _, err := conn.Write([]byte(sb.String())); err != nil {
		log.Println("send message error:", err)
	}
}

func toJson(text string) []byte {
	result, _ := json.Marshal(map[string]string{"text": text})
	return result
}

func translate(ts *googletrans.TranslateParams, text string) (string, error) {
	ts.Src = "auto"
	ts.Dest = language.SimplifiedChinese.String()
	ts.Text = text

	if translated, err := googletrans.Translate(*ts); err != nil {
		return "", err
	} else {
		return translated.Text, nil
	}
}
