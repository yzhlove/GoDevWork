package main

import (
	"github.com/unknwon/goconfig"
)

//goconfig

const dir = "Chats/auth/test07/"

func main() {

	f, err := goconfig.LoadConfigFile(dir + "test.ini")
	if err != nil {
		panic(err)
	}

	//f.SetValue("test", "lcm", "234567")
	//
	//v, err := f.GetValue("test", "yzh")
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(v)

	f.DeleteSection("test")
	//f.DeleteKey("test", "lcm")


	save(f)

}

func save(c *goconfig.ConfigFile) {
	_ = goconfig.SaveConfigFile(c, dir+"test.ini")
}
