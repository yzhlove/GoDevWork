package main

import (
	"fmt"
	"gee"
	"log"
	"net/http"
)

func main() {

	TestMain()

	fmt.Println("------------------------------------------")

	gin := gee.New()
	gin.GET("/", index)
	gin.GET("/hello", hello)
	gin.GET("/hello/:name", hello)
	gin.GET("/assets/*filepath", assert)
	log.Print(gin.Run(":1234"))

}

func index(ctx *gee.Context) {
	ctx.HTML(http.StatusOK, "<h1>Hello World</h1>")
}

func hello(ctx *gee.Context) {
	ctx.String(http.StatusOK, "hello %s ,path %s \n", ctx.Query("name"), ctx.Path)
}

func assert(ctx *gee.Context) {
	ctx.JSON(http.StatusOK, gee.H{
		"filepath": ctx.Param("filepath"),
	})
}

func TestMain() {
	var paths = []string{
		"/",
		"/hello/:name",
		"hello/a/b",
		"/hi/:name",
		"/assets/*filepath",
	}

	r := gee.NewRouter()
	for _, path := range paths {
		r.AddRouter("GET", path, nil)
	}

	for _, path := range paths {
		fmt.Println(gee.ParseUrl(path))
	}
	fmt.Println("===================GGGGGGEEEEEE======================")
	fmt.Println(r.GetRouter("GET", "assets/css/geektutu.css"))
	fmt.Println("=========================================")
	r.Show()

}
