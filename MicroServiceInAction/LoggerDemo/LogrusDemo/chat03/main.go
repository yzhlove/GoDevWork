package main

import "github.com/sirupsen/logrus"

//日志分级

func main() {

	log := logrus.New()
	log.SetLevel(logrus.TraceLevel)
	log.Trace("trace")
	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
	log.Error("error")
	log.Fatal("fatal")
	log.Panic("panic")

}
