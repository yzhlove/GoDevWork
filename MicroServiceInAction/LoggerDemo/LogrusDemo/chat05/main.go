package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func init() {
	log.Formatter = &logrus.JSONFormatter{}
	f, _ := os.Create("gin.log")
	log.Out = f
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = log.Out
	log.Level = logrus.InfoLevel
}

func main() {
	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {
		log.WithFields(logrus.Fields{
			"router": "hello",
			"size":   10,
			"params": context.Params,
		}).Warn("router by hello")
		context.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	r.Run(":1234")
}
