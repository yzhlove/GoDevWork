package main

import (
	"geefive"
	"log"
	"net/http"
	"time"
)

func onlyForV2() geefive.HandleFunc {
	return func(ctx *geefive.Context) {
		t := time.Now()
		ctx.Fail(http.StatusInternalServerError, "Server Busy")
		log.Printf("[%d] %s in %v for group v2", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
	}
}

func main() {

	gin := geefive.NewEngine()
	gin.Use(geefive.Logger())
	gin.GET("/", func(ctx *geefive.Context) {
		ctx.BinaryValue(http.StatusOK, []byte("Hello World"))
	})
	v1 := gin.NewGroup("v1")
	v1.GET("/hello", func(ctx *geefive.Context) {
		ctx.BinaryValue(http.StatusOK, []byte("v1-hello"))
	})
	v2 := gin.NewGroup("v2")
	v2.Use(onlyForV2())
	v2.GET("/hello", func(ctx *geefive.Context) {
		ctx.BinaryValue(http.StatusOK, []byte("v2-hello"))
	})
	log.Println(gin.Run(":1234"))
}
