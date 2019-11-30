package log

import (
	"testing"
	"time"
)

func Test_WriteLog(t *testing.T) {

	if err := UserLogInit(); err != nil {
		t.Error(err)
		return
	}

	for i := 0; i <= 20; i++ {
		WriteLog("ts", time.Now().Unix(), "operator", "yzhlove", "event", "test write", "counter", i)
		time.Sleep(time.Second)
	}

}
