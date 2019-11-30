package userloggers

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"testing"
	"time"
)

func Test_AppendFile(t *testing.T) {

	path := "test.log"
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		t.Error(err)
		return
	}
	_, _ = f.WriteString("Hello World\n")
	t.Log("Ok.")
}

func Test_WriteLog(t *testing.T) {

	Init()

	for i := 0; i < 10; i++ {
		WriteUserLog(&LogMessage{
			Timestamp: time.Now().Unix(),
			Operator:  "我和我的祖国",
			Event:     strconv.Itoa(i),
			Status:    0,
		})
	}

	time.Sleep(time.Second)

}

func BenchmarkInit(b *testing.B) {

	Init()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		WriteUserLog(&LogMessage{
			Timestamp: time.Now().Unix(),
			Operator:  "yzhlove",
			Event:     strconv.Itoa(i),
			Status:    0,
		})
	}
	b.StopTimer()

	time.Sleep(10 * time.Second)
}

func Test_GetFileDir(t *testing.T) {
	s := "/a/b/c/"
	t.Log("path => ", s)
	t.Log("dir => ", filepath.Dir(s))
}

func Test_GetFileExt(t *testing.T) {

	log := "/user/loggers/user_logger_2019_11_25.log"

	t.Log(path.Base(log))
	t.Log(path.Ext(log))

}

func Test_ShowFileList(t *testing.T) {

	names := showFileList()
	t.Log(names)
}

func Test_ReadFile(t *testing.T) {

	files := showFileList()
	msgList, err := readLogFile(files[0])
	if err != nil {
		t.Error(err)
		return
	}
	for _, msg := range msgList {
		fmt.Println(msg)
	}

	logs := formatLogFile(msgList)
	for _, log := range logs {
		fmt.Println(log)
	}

}

func Test_CrossTime(t *testing.T) {

	if err := Init(); err != nil {
		t.Error(err)
		return
	}

	for i := 0; i <= 18; i++ {
		WriteUserLog(&LogMessage{
			Timestamp: time.Now().Unix(),
			Operator:  "yzhlove",
			Event:     "test",
			Status:    0,
		})
		time.Sleep(10 * time.Second)
	}
}
