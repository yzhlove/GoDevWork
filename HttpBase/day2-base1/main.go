package main

import (
	"gee"
	"log"
	"net/http"
)

func main() {

	r := gee.New()
	r.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>HELLO WORLD</h1>")
	})
	r.GET("/hello", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "hello %s,at %s \n", ctx.Query("name"), ctx.Path)
	})
	r.POST("/login", func(ctx *gee.Context) {
		ctx.JSON(http.StatusOK, gee.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})
	log.Fatal(r.Run(":1234"))
}
