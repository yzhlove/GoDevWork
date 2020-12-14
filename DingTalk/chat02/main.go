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

var picture = "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1607952372355&di=1bda1c9514c751090bb2ea2c8e77be8b&imgtype=0&src=http%3A%2F%2Fpic.rmb.bdstatic.com%2F62302488585a816a8938b328951d4982.jpeg"
var msg = "#### 这里是上海  \n> 换个视角看降雪 祖国山河美成了一副水墨画\n> ![screenshot](" + picture + ")\n> ###### [天气](https://www.dingtalk.com) \n"
