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
	"time"
)

type DecoderCoder struct {
	Address string
	Secret  []byte
}

func (d *DecoderCoder) Format() string {
	ts := time.Now().UnixNano() / 1e6
	data := strconv.FormatInt(ts, 10) + "\n" + string(d.Secret)
	mac := hmac.New(sha256.New, d.Secret)
	mac.Write([]byte(data))
	secret := url.QueryEscape(base64.StdEncoding.EncodeToString(mac.Sum(nil)))
	return fmt.Sprintf("%s&timestamp=%d&sign=%s", d.Address, ts, secret)
}

type DingTask struct {
	data *DecoderCoder
	send Sender
}

func (d *DingTask) task() error {
	resp, err := http.Post(d.data.Format(),
		"application/json;charset=UTF-8",
		bytes.NewBuffer(d.send.Send()))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Print("ding task result:")
	fmt.Println("↓ ", string(data))
	return nil
}
