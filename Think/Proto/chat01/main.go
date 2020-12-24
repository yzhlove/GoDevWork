package main

import "fmt"

//go:generate msgp -io=false -tests=false
type Content struct {
	Id    uint32
	Reply *Content
}

type Wechat struct {
	Hid      uint32
	Contents []*Content
}

func main() {

	a := &Content{Id: 123, Reply: nil}
	b := &Content{Id: 456, Reply: a}
	c := &Content{Id: 789, Reply: b}
	d := &Content{Id: 100, Reply: c}

	w := &Wechat{Hid: 12138, Contents: []*Content{a, b, c, d}}
	result, err := w.MarshalMsg(nil)
	if err != nil {
		panic(err)
	}

	wechat := &Wechat{}
	if _, err = wechat.UnmarshalMsg(result); err != nil {
		panic(err)
	}

	fmt.Println("Hid ==> ", wechat.Hid)
	for _, content := range wechat.Contents {
		fmt.Println("content ==> ", content, " ptr => ", content.Reply)
	}

}
