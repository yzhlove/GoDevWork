package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"text/template"
)

var types = []string{"packet_type", "name", "payload", "desc"}
var names = []string{"_req", "_ack", "_notify"}
var checks = []string{"_req", "_succeed_ack", "_failed_ack"}

//读取的格式
type Reader struct {
	packetType string
	name       string
	payload    string
	desc       string
}

//写出的格式
type Writer struct {
	PkgName string
	Packets []*packet
	Acks    map[string][2]string
}

//template需要的数据
type packet struct {
	Code string
	Name string
	Desc string
}

//配置文件结构
type api struct {
	file      string
	maxPacket int
	minPacket int
	tmpl      string
	pkgName   string
}

//配置文件
var _api *api

var api_tmpl_path = "/Users/yostar/workSpace/gowork/src/GoDevWork/GoNet/gonet-tools/proto_scripts/templates/server/api.tmpl"
var api_txt_path = "/Users/yostar/workSpace/gowork/src/GoDevWork/GoNet/gonet-tools/proto_scripts/api/api.txt"
var package_name = "micro_agent"

func init() {
	file := flag.String("f", api_txt_path, "input api.txt")
	min := flag.Int("min", 0, "min proto Code")
	max := flag.Int("max", 10000, "max proto Code")
	tmpl := flag.String("tmpl", api_tmpl_path, "input template")
	pkg := flag.String("pkg", package_name, "import package Name")
	flag.Parse()
	_api = &api{file: *file, minPacket: *min, maxPacket: *max, tmpl: *tmpl, pkgName: *pkg}
}

func main() {
	f, err := os.Open(_api.file)
	if err != nil {
		syntax(err, 0)
	}
	defer f.Close()

	sig := make(chan struct{})
	writer := &Writer{PkgName: _api.pkgName}
	reader := make(chan *Reader, 128)
	writerText(reader, writer, sig)
	readerText(f, reader)
	<-sig

	funcMap := template.FuncMap{
		"isReq": func(name string) bool {
			if strings.HasSuffix(name, "_req") {
				return true
			}
			return false
		},
		"toUpper": func(name string) string {
			if res := strings.Split(name, "_"); len(res) > 0 {
				var str string
				for _, r := range res {
					str += strings.Title(r)
				}
				return str
			}
			return strings.Title(name)
		},
	}
	//*必须要与parseFile的名字一致，如果有多个文件，只需要与其中一个文件相同即可
	tmpl, err := template.New("api.tmpl").Funcs(funcMap).ParseFiles(_api.tmpl)
	if err != nil {
		syntax(err, 0)
	}
	if err = tmpl.Execute(os.Stdout, writer); err != nil {
		syntax(err, 0)
	}
}

//读取
func readerText(in io.Reader, reader chan *Reader) {
	defer close(reader)
	packetMap := make(map[string]string, 4)
	var _succeed bool
	index := 0
	for scan, no := bufio.NewScanner(in), 1; scan.Scan(); no++ {

		line := strings.TrimSpace(scan.Text())
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			if _succeed && index == 4 {
				reader <- &Reader{
					packetType: packetMap["packet_type"],
					name:       packetMap["name"],
					payload:    packetMap["payload"],
					desc:       packetMap["desc"],
				}
			}
			index = 0
			_succeed = false
			continue
		}

		res := strings.Split(line, ":")
		if len(res) != 2 {
			syntax(errors.New("format error"), no)
		}

		pt := strings.TrimSpace(res[0])
		if pt != types[index] {
			syntax(errors.New(types[index]+" not found by packet_type:"+pt), no)
		}

		switch pt {
		case "packet_type":
			if code, err := strconv.Atoi(res[1]); err != nil {
				syntax(errors.New("code must number"), no)
			} else if code < _api.minPacket || code > _api.maxPacket {
				syntax(errors.New("code out of range"), no)
			}
		case "name":
			var ok bool
			for _, ck := range names {
				if strings.HasSuffix(res[1], ck) {
					ok = true
					break
				}
			}
			if !ok {
				syntax(errors.New("name suffix not found"), no)
			}
		}

		packetMap[pt] = strings.TrimSpace(res[1])
		index++
		if index == 4 {
			_succeed = true
		}
	}
}

func writerText(in chan *Reader, w *Writer, sig chan struct{}) {
	go func() {
		hashMap := make(map[string][]string)
		for reader := range in {
			w.Packets = append(w.Packets, &packet{Code: reader.packetType,
				Name: reader.name, Desc: reader.desc})
			for i, ck := range checks {
				if !strings.HasSuffix(reader.name, ck) {
					continue
				}
				t := reader.name[:strings.LastIndex(reader.name, ck)]
				if _, ok := hashMap[t]; !ok {
					hashMap[t] = make([]string, 3)
				}
				hashMap[t][i] = reader.packetType
				break
			}
		}
		w.Acks = make(map[string][2]string, len(hashMap))
		for _, value := range hashMap {
			w.Acks[value[0]] = [2]string{value[1], value[2]}
		}
		close(sig)
	}()
}

func syntax(err error, no int) {
	panic(fmt.Sprintf("\033[1;31m ♠︎ ︎no:%d err:%s \033[0m", no, err))
}
