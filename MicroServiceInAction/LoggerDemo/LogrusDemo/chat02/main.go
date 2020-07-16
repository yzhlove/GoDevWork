package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func main() {

	log.Out = os.Stdout
	f, err := os.OpenFile("logrus_test_chat02.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.Out = f
	log.Info("set log file succeed.")
	log.WithFields(logrus.Fields{"animal": "cat", "size": 120}).Info("wahahahahahah")

}
