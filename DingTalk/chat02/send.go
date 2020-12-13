package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type DecoderCoder struct {
	Timestamp  int64
	Address    string
	Secret     []byte
	DecoderStr string
}

func (d *DecoderCoder) Parse(ts int64) {
	d.Timestamp = ts / 1e6
	data := strconv.FormatInt(d.Timestamp, 10) + "\n" + string(d.Secret)
	mac := hmac.New(sha256.New, d.Secret)
	mac.Write([]byte(data))
	d.DecoderStr = url.QueryEscape(base64.StdEncoding.EncodeToString(mac.Sum(nil)))
}

func (d *DecoderCoder) Format() string {
	return fmt.Sprintf("%s&timestamp=%d&sign=%s", d.Address, d.Timestamp, d.DecoderStr)
}

func dingTalk(d *DecoderCoder, s Sender) (string, error) {
	resp, err := http.Post(d.Format(),
		"application/json;charset=UTF-8",
		bytes.NewBuffer(s.Send()))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
