package main

import "github.com/sirupsen/logrus"

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.SetReportCaller(true) //将函数名添加位字段
	log.Formatter = &logrus.JSONFormatter{}
}

func main() {
	testLog()
}

func testLog() {
	log.Info("test function out")
}
