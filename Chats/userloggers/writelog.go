package userloggers

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type LogMessage struct {
	Timestamp int64  //操作时间
	Operator  string //操作人
	Event     string //事件
	Status    uint8  //操作状态
}

type WriterLogger interface {
	Write(msgList *LogMessage)
}

type UserLogger struct {
	ChanMessage chan *LogMessage
	Logger      WriterLogger
}

type FileLoggerWriter struct {
	LastTimestamp int64 //上一次写入日志的时间
	File          *os.File
}

func getNewLogFile() string {
	_ = time.Now()
	return ""
	//return fmt.Sprintf("%suser_logger_%d_%d_%d.zlog", config.UserLoggerPath, tm.Year(), tm.Month(), tm.Day())
}

func getNewLoggerWrite() (fw *FileLoggerWriter, err error) {
	fw = new(FileLoggerWriter)
	fw.LastTimestamp = time.Now().Unix()
	fw.File, err = os.OpenFile(getNewLogFile(), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	return
}

func (fw *FileLoggerWriter) Write(msg *LogMessage) {
	//如果不是同一天
	if time.Unix(fw.LastTimestamp, 0).YearDay() != time.Unix(msg.Timestamp, 0).YearDay() {
		//先释放文件资源
		if fw.File != nil {
			{
				defer fw.File.Close()
			}
		}
		if f, err := getNewLoggerWrite(); err != nil {
			fmt.Println("open user zlog err:", err)
			return
		} else {
			fw.File = f.File
		}
	}
	var build strings.Builder
	build.WriteString(msg.Operator + "\t")
	build.WriteString(strconv.FormatInt(msg.Timestamp, 10) + "\t")
	build.WriteString(msg.Event + "\t")
	build.WriteString(strconv.Itoa(int(msg.Status)) + "\n")
	//更新上一次的写入时间
	fw.LastTimestamp = msg.Timestamp
	//写入文件
	_, _ = fw.File.WriteString(build.String())

}

func newUserLogger(size int, logger WriterLogger) *UserLogger {
	if size == 0 {
		size = 128
	}
	return &UserLogger{
		ChanMessage: make(chan *LogMessage, size),
		Logger:      logger,
	}
}

func (logger *UserLogger) start() {

	go func() {
		for {
			select {
			case msg := <-logger.ChanMessage:
				logger.Logger.Write(msg)
			}
		}
	}()

}

//写入用户日志到通道
func WriteUserLog(msg *LogMessage) {
	if userLogger != nil {
		userLogger.ChanMessage <- msg
	}
}

var userLogger *UserLogger

func Init() error {

	if fw, err := getNewLoggerWrite(); err != nil {
		return err
	} else {
		userLogger = newUserLogger(128, fw)
		userLogger.start()
	}
	return nil
}
