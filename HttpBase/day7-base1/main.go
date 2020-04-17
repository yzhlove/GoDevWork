package main

import (
	"geeseven"
	"log"
	"net/http"
)

func main() {

	gin := geeseven.NewDefaultEngine()
	gin.GET("/", func(ctx *geeseven.Context) {
		ctx.Text(http.StatusOK, []byte("Hello World"))
	})
	gin.GET("/panic", func(ctx *geeseven.Context) {
		name := []string{"test"}
		ctx.String(http.StatusOK, name[100])
	})
	log.Println(gin.Run(":1234"))
}
