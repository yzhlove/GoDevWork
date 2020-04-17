package main

import (
	"geefour"
	"log"
	"net/http"
)

func main() {

	gee := geefour.New()
	gee.GET("/", indexHandle)
	v1 := gee.Group("v1")
	v1.GET("/", v1IndexHandle)
	v2 := gee.Group("v2")
	v2.GET("/hello", helloHandle)
	log.Println(gee.Run(":1234"))
}

func indexHandle(ctx *geefour.Context) {
	ctx.JSON(http.StatusOK, geefour.H{"status": 0, "message": "hello world"})
}

func v1IndexHandle(ctx *geefour.Context) {
	ctx.JSON(http.StatusOK, geefour.H{"version": 1, "message": "hello"})
}

func helloHandle(ctx *geefour.Context) {
	ctx.JSON(http.StatusOK, geefour.H{"path": "/hello", "version": 2, "message": "hi"})
}
