package main

import (
	"github.com/unknwon/goconfig"
)

func main() {
	c, err := goconfig.LoadConfigFile("./goconfig.ini")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		//c.SetSectionComments("test", strconv.Itoa(i))
		//c.SetValue("pwd", "-", strconv.Itoa(i))
		//c.SetKeyComments("abc", "ttt", strconv.Itoa(i))
		//c.SetKeyComments("test", "t_"+strconv.Itoa(i+1), strconv.Itoa(i))
	}

	//for i := 10; i < 20; i++ {
	//	c.SetSectionComments("pwd", strconv.Itoa(i))
	//}

	goconfig.SaveConfigFile(c, "./goconfig.ini")

}
