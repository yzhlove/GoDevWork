package main

func main() {

	dc := &DecoderCoder{
		Address: "https://oapi.dingtalk.com/robot/send?access_token=c5180d84fa90f7d852d69a9734e5718b8818c9fe7ba2d2a13ca0b3067994b2bf",
		Secret:  []byte("SEC5a4160fc8a3fbfa0986ed2ac8d88bffa39f514df9fc9c68b0079b2faaf73f260"),
	}

	doc := &MarkDoc{MsgType: "markdown", MarkData: MarkData{
		Tag:     "光阴的故事",
		Content: msg,
	}}

	tQueue.set(&DingTask{data: dc, send: doc})
	tQueue.run()
}

var picture = "https://cdn.hk01.com/di/media/images/3764026/org/087c4ba059bc16fe7d25ce3c0cc62e2b.jpg/_qom99EsDMiOMWiJBquJKS2rEfxmkBXDCR5WsgkeVrI?v=w1920"
var msg = "#### 这里是上海  \n> 换个视角看降雪 祖国山河美成了一副水墨画\n> ![screenshot](" + picture + ")\n> ###### [天气](https://www.dingtalk.com) \n"
