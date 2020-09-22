package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io"
	"os"
	"strings"
)

//映射表结构
var _types map[string]map[string]struct {
	T string `json:"t"`
	R string `json:"r"`
	W string `json:"w"`
}

type Member struct {
	Typ        string
	Val        string
	ReaderFunc string
	WriterFunc string
}

type Class struct {
	Name    string
	Members []Member
	Desc    string
}

type Writer struct {
	PkgName string
	Classes []Class
}

type proto struct {
	file    string
	tmpl    string
	pkgName string
	trans   string
}

var _proto *proto

var proto_txt_path = "/Users/yostar/workSpace/gowork/src/GoDevWork/GoNet/gonet-tools/proto_scripts/proto/proto.txt"
var proto_tmpl_path = "/Users/yostar/workSpace/gowork/src/GoDevWork/GoNet/gonet-tools/proto_scripts/templates/server/proto_new.tmpl"
var proto_types = "/Users/yostar/workSpace/gowork/src/GoDevWork/GoNet/gonet-tools/proto_scripts/proto/transform.json"
var package_name = "micro_agent"

func init() {
	file := flag.String("file", proto_txt_path, "input proto.txt")
	tmpl := flag.String("tmpl", proto_tmpl_path, "input proto.tmpl")
	pkgName := flag.String("pkg_name", package_name, "packet name")
	tps := flag.String("trans", proto_types, "type transform file")
	flag.Parse()
	_proto = &proto{*file, *tmpl, *pkgName, *tps}
}

func main() {

	f, err := os.Open(_proto.file)
	if err != nil {
		syntax(err, 0)
	}
	defer f.Close()

	load()
	reader := make(chan *Class, 128)
	writer := &Writer{PkgName: _proto.pkgName}
	sig := make(chan struct{})
	writerText(reader, writer, sig)
	readerText(f, reader)
	<-sig

	tmpl, err := template.New("proto_new.tmpl").ParseFiles(_proto.tmpl)
	if err != nil {
		syntax(err, 0)
	}
	if err = tmpl.Execute(os.Stdout, writer); err != nil {
		syntax(err, 0)
	}
}

func load() {
	f, err := os.Open(_proto.trans)
	if err != nil {
		syntax(err, 0)
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(&_types); err != nil {
		syntax(err, 0)
	}
}

func writerText(in chan *Class, w *Writer, sig chan struct{}) {
	go func() {
		for reader := range in {
			class := Class{Name: toFirstUpper(reader.Name)}
			class.Desc = "// " + class.Name + " " + reader.Desc
			class.Members = make([]Member, 0, len(reader.Members))
			for _, c := range reader.Members {
				member := Member{Val: toFirstUpper(c.Val)}
				if _, ok := _types[c.Typ]; ok {
					member.Typ = _types[c.Typ]["go"].T
					member.ReaderFunc = _types[c.Typ]["go"].R
					member.WriterFunc = _types[c.Typ]["go"].W
				} else {
					syntax(errors.New("not found type:"+c.Typ), 0)
				}
				class.Members = append(class.Members, member)
			}
			w.Classes = append(w.Classes, class)
		}
		close(sig)
	}()
}

func readerText(in io.Reader, ch chan *Class) {
	defer close(ch)
	reader := Class{}
	var _succeed bool
	for scan, no := bufio.NewScanner(in), 0; scan.Scan(); no++ {
		str := strings.TrimSpace(scan.Text())
		if strings.HasPrefix(str, "#") {
			reader.Desc = strings.TrimLeft(str, "#")
		} else if strings.HasSuffix(str, "=") {
			if str == "===" && _succeed {
				ch <- reader.cp()
				reader.Members = reader.Members[:0]
				_succeed = false
			} else {
				reader.Name = strings.TrimRight(str, "=")
				_succeed = true
			}
		} else {
			if len(str) > 0 {
				if res := strings.Split(str, " "); len(res) != 2 {
					syntax(errors.New("format error"), no)
				} else {
					reader.Members = append(reader.Members, Member{Val: res[0], Typ: res[1]})
				}
			}
		}
	}
}

func (r *Class) cp() *Class {
	reader := &Class{Name: r.Name, Desc: r.Desc}
	reader.Members = make([]Member, 0, len(r.Members))
	for _, c := range r.Members {
		reader.Members = append(reader.Members, Member{Typ: c.Typ, Val: c.Val})
	}
	return reader
}

func toFirstUpper(str string) string {
	if res := strings.Split(str, "_"); len(res) != 0 {
		var rule string
		for _, t := range res {
			rule += strings.Title(t)
		}
		return rule
	}
	return strings.Title(str)
}

func syntax(err error, no int) {
	panic(fmt.Sprintf("\033[1;31m ♠︎ ︎no:%d err:%s \033[0m", no, err))
}
