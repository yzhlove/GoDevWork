package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	jsonFile = "./loveJson.txt"
	textFile = "./loveText.txt"
)

func main() {

	fmt.Println("[服务器已经启动，正在监听本地 (1314) 端口] -v- ")
	http.HandleFunc("/love", loveHandle)
	http.HandleFunc("/like", likeHandle)
	if err := http.ListenAndServe(":1314", nil); err != nil {
		fmt.Println("Error:", err)
	}

}

func setDomain(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func loveHandle(resp http.ResponseWriter, _ *http.Request) {

	setDomain(resp)

	jsonData := make(map[string]string, 16)

	if file, err := os.Open(jsonFile); err != nil {
		_, _ = io.WriteString(resp, "Read Json File Error:"+err.Error())
	} else {
		defer file.Close()
		buffer := bufio.NewReader(file)
		var count int
		for {
			line, err := buffer.ReadString('\n')
			if err != nil || err == io.EOF {
				break
			}
			count++
			str := strings.Split(line[0:len(line)-1], ":")
			if str[0] != "" && str[1] != "" {
				jsonData[strings.Trim(str[0], " ")] = strings.Trim(str[1], " ")
			}
		}
		if len(jsonData) == 0 {
			_, _ = io.WriteString(resp, "Error: file invalid !")
			return
		}
		jsonData["counter"] = strconv.Itoa(count)
		if jsonStr, err := json.Marshal(jsonData); err != nil {
			_, _ = io.WriteString(resp, "Error: json marshal err "+err.Error())
		} else {
			_, _ = io.WriteString(resp, string(jsonStr))
		}
		return
	}

}

func likeHandle(resp http.ResponseWriter, _ *http.Request) {

	setDomain(resp)

	if data, err := ioutil.ReadFile(textFile); err != nil {
		_, _ = io.WriteString(resp, "Error: read text file err "+err.Error())
	} else {
		_, _ = io.WriteString(resp, string(data))
	}
	return
}
