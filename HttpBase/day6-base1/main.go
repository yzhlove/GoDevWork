package main

import (
	"fmt"
	"geesix"
	"html/template"
	"log"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func formatAsDate(t time.Time) string {
	y, m, d := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", y, m, d)
}

func main() {
	gin := geesix.NewEngine()
	gin.Use(geesix.Logger())
	gin.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	gin.LoadHTMLGlob("./templates/*")
	gin.Static("/assets", "/Users/yurisa/Develop/GoWork/src/WorkSpace/GoDevWork/HttpBase/day6-base1/static")
	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	gin.GET("/", func(c *geesix.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	gin.GET("/students", func(c *geesix.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", geesix.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})
	gin.GET("/date", func(c *geesix.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", geesix.H{
			"title": "gee",
			"now":   time.Date(2020, 10, 18, 0, 0, 0, 0, time.UTC),
		})
	})
	log.Println(gin.Run(":9999"))
}
