package utils

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}
