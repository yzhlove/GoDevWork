package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const address = "aaa"

var secret = []byte("bbb")

func main() {

	t := time.Now().UnixNano() / 1e6
	res := genSecret(t)
	set := fmt.Sprintf("%s&timestamp=%d&sign=%s", address, t, res)
	tText := `data := strconv.FormatInt(t, 10) + "\n" + string(secret)
	//sha256
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(data))
	//base64
	result := base64.StdEncoding.EncodeToString(h.Sum(nil))
	//urlEncoder
	res := url.QueryEscape(result)`
	result, err := http.Post(set, "application/json;charset=UTF-8", bytes.NewBuffer(createMessage(tText)))
	if err != nil {
		panic(err)
	}

	if data, err := ioutil.ReadAll(result.Body); err != nil {
		panic(err)
	} else {
		defer result.Body.Close()
		fmt.Println("result => ", string(data))
	}

}

// genSecret 生成钉钉签名
func genSecret(t int64) string {

	data := strconv.FormatInt(t, 10) + "\n" + string(secret)
	//sha256
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(data))
	//base64
	result := base64.StdEncoding.EncodeToString(h.Sum(nil))
	//urlEncoder
	res := url.QueryEscape(result)

	return res
}

type Text struct {
	Content string `json:"content"`
}

type Msg struct {
	MsgType string `json:"msgtype"`
	Text    Text   `json:"text"`
}

func createMessage(str string) []byte {
	m := &Msg{MsgType: "text", Text: Text{Content: str}}
	r, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return r
}
