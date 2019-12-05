package userlog

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func init() {
	//config.UserLoggerPath = ""
}

func TestWriteLog(t *testing.T) {

	if err := LogWriterInit(); err != nil {
		t.Error(err)
		return
	}

	//写入日志
	for i := 0; i < 10; i++ {
		writeMessage("succeed", "ok", "index", i)
	}
	time.Sleep(5 * time.Second)
}

func writeMessage(msg ...interface{}) {
	tm := time.Now()
	msg = append(msg, "year", tm.Year(), "day", tm.YearDay(), "ts", tm.Unix(), "operator", "yzh", "event", "test-write")
	WriteLog(msg)
}

func show(messages []interface{}) {
	for i, msg := range messages {
		fmt.Printf("[%d - %T] %v \n", i, msg, msg)
	}
}

func Test_ReaderAll(t *testing.T) {

	if err := LogReaderInit(); err != nil {
		t.Error(err)
		return
	}
	type search struct {
		Event string `json:"event"`
	}

	s, err := json.Marshal(search{
		Event: "all",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("cond => ", string(s))
	//查询
	if data, err := QueryResult(string(s)); err != nil {
		t.Error(err)
		return
	} else {
		//显示查询结果
		fmt.Println(data)

	}
	t.Log("ok.")
}

func Test_ReaderCond(t *testing.T) {

	if err := LogReaderInit(); err != nil {
		t.Error(err)
		return
	}

	type C struct {
		Start int64 `json:"start"`
		End   int64 `json:"end"`
		Limit int   `json:"limit"`
	}

	type search struct {
		Event string `json:"event"`
		Cond  C      `json:"cond"`
	}

	s, err := json.Marshal(search{
		Event: "time",
		Cond: C{
			Start: 1575424800,
			Limit: 20,
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("cond => ", string(s))

	if data, err := QueryResult(string(s)); err != nil {
		t.Error(err)
		return
	} else {
		fmt.Println(data)
	}
	t.Log("ok.")
}

func Test_GenerateConds(t *testing.T) {
	type C struct {
		Operator string `json:"operator"` //操作人员
		Mints    int64  `json:"start"`    //开始时间
		Maxts    int64  `json:"end"`      //结束时间（default:当前时间）
	}

	type search struct {
		Event string `json:"event"`
		Cond  C      `json:"cond"`
	}

	s, err := json.Marshal(search{
		Event: "",
		Cond: C{
			Operator: "yzhtest",
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(string(s))
}

func Test_ReaderComplex(t *testing.T) {

	if err := LogReaderInit(); err != nil {
		t.Error(err)
		return
	}

	type C struct {
		Operator string `json:"operator"` //操作人员
		Mints    int64  `json:"start"`    //开始时间
		Maxts    int64  `json:"end"`      //结束时间（default:当前时间）
		Event    string `json:"event"`    //事件
		Limit    int    `json:"limit"`    //查询条数（default:100）
	}

	type search struct {
		Event string `json:"event"`
		Cond  C      `json:"cond"`
	}

	s, err := json.Marshal(search{
		Event: "",
		Cond: C{
			Mints:    1575437894,
			Operator: "yzh",
		},
	})

	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("json cond => ", string(s))

	data, err := QueryResult(string(s))
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(data)
	t.Log("ok.")
}
