package main

import (
	"fmt"
	"github.com/unknwon/goconfig"
)

func main() {

	g, err := goconfig.LoadConfigFile("./goconfig.ini")
	if err != nil {
		panic(err)
	}

	fmt.Println(g.GetKeyList("test"))
	fmt.Println(g.GetSection("test"))

	//for i := 0; i < 5; i++ {
	//	g.SetKeyComments("pwd", "-", "10")
	//}

	//g.SetValue("pwd", "-", "14")
	//for i := 0; i < 5; i++ {
	//	g.SetValue("pwd", "-", "15-"+strconv.Itoa(i))
	//}

	g.DeleteSection("pwd")

	_ = goconfig.SaveConfigFile(g, "./goconfig.ini")

}
