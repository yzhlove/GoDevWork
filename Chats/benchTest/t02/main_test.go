package test

import (
	"os"
	"testing"
)

func Test_OutCsv(t *testing.T) {

	f, err := os.OpenFile("test.csv", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}

	_, _ = f.WriteString("姓名,姓别,生日\n")
	_, _ = f.WriteString("余子涵,男,1996-12-24\n")
	_, _ = f.WriteString("向金晶,女,1996-05-28\n")
	_, _ = f.WriteString("徐宥嘉,女,1996-07-13\n")

	f.Close()

}
