package userloggers

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const msgLen = 4

type LogRecord struct {
	Operator  string
	TimeStamp string
	Event     string
	Status    string
}

type LogInfo struct {
	Operator  string
	Event     string
	TimeStamp string
	Status    bool
}

func showFileList() []string {
	logName := getNewLogFile()
	var fileNames []string
	err := filepath.Walk(filepath.Dir(logName), func(s string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}
		if !info.IsDir() {
			//检查文件以及后缀名
			if strings.Contains(s, "user_logger") && path.Ext(s) == ".log" {
				fileNames = append(fileNames, s)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("showFileListErr: ", err)
		return nil
	}
	return fileNames
}

func readLogFile(path string) (records []LogRecord, err error) {
	var f *os.File
	if f, err = os.Open(path); err != nil {
		return
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	var line string

	for {
		if line, err = reader.ReadString('\n'); err != nil {
			if io.EOF == err {
				err = nil
			}
			break
		}
		//去除换行符
		line := strings.Trim(line, "\n")
		if str := strings.Split(line, "\t"); len(str) == msgLen {
			records = append(records, LogRecord{
				Operator: str[0], TimeStamp: str[1], Event: str[2], Status: str[3],
			})
		}
	}
	return
}

func formatLogFile(records []LogRecord) []LogInfo {
	logs := make([]LogInfo, 0, len(records))
	format := "2006-01-02 15:04:05"
	for _, record := range records {
		if ts, err := strconv.ParseInt(record.TimeStamp, 10, 64); err != nil {
			continue
		} else {
			var status bool
			if record.Status == "1" {
				status = true
			}
			logs = append(logs, LogInfo{
				Operator: record.Operator, Event: record.Event, TimeStamp: time.Unix(ts, 0).Format(format), Status: status,
			})
		}
	}
	return logs
}
