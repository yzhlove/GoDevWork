package main

import (
	"fmt"
	"time"
)

func main() {

	dc := &DecoderCoder{
		Address: "https://oapi.dingtalk.com/robot/send?access_token=c5180d84fa90f7d852d69a9734e5718b8818c9fe7ba2d2a13ca0b3067994b2bf",
		Secret:  []byte("SEC5a4160fc8a3fbfa0986ed2ac8d88bffa39f514df9fc9c68b0079b2faaf73f260"),
	}
	dc.Parse(time.Now().UnixNano())

	doc := &MarkDoc{MsgType: "markdown", MarkData: MarkData{
		Tag:     "天仙",
		Content: msg,
	}}

	res, err := dingTalk(dc, doc)
	if err != nil {
		panic(err)
	}

	fmt.Println("ding => ", res)

}

var msg = "#### 杭州天气  \n> 9度，西北风1级，空气良89，相对温度73%\n> ![screenshot](http://image.qianye88.com/pic/fb803f9e848476306bac335c1073e0a0)\n> ###### 10点20分发布 [天气](https://www.dingtalk.com) \n"
