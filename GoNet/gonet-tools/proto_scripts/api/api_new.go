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

var checkTypes = []string{"packet_type", "name", "payload", "desc"}
var checkNames = []string{"_req", "_ack", "_notify"}

//读取的格式
type ReaderString struct {
	packetType string
	name       string
	payload    string
	desc       string
}

//写出的格式
type WriterString struct {
	PkgName string
	Packets []*packet
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
	writer := &WriterString{PkgName: _api.pkgName}
	reader := make(chan *ReaderString, 128)
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
func readerText(in io.Reader, reader chan *ReaderString) {

	packetMap := make(map[string]string, 4)
	var _succeed bool
	index := 0
	for scan, no := bufio.NewScanner(in), 1; scan.Scan(); no++ {

		line := strings.TrimSpace(scan.Text())
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			if _succeed && index == 4 {
				reader <- &ReaderString{
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
		if pt != checkTypes[index] {
			syntax(errors.New(checkTypes[index]+" not found by packet_type:"+pt), no)
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
			for _, ck := range checkNames {
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
	close(reader)
}

func writerText(in chan *ReaderString, w *WriterString, sig chan struct{}) {
	go func() {
		for reader := range in {
			w.Packets = append(w.Packets, &packet{Code: reader.packetType,
				Name: reader.name, Desc: reader.desc})
		}
		close(sig)
	}()
}

func syntax(err error, no int) {
	if err != nil {
		var show string
		if no == 0 {
			show = fmt.Sprintf("\033[1;31m ♠︎ err:%s \033[0m", err)
		} else {
			show = fmt.Sprintf("\033[1;31m ♠︎ ︎no:%d err:%s \033[0m", no, err)
		}
		panic(show)
	}
}
