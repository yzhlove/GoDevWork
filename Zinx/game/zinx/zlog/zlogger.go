package zlog

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
	"time"
)

const LOG_MAX_BUF = 1024 * 1024

const (
	BIT_DATA = 1 << iota
	BIT_TIME
	BIT_MICRO_SECONDS
	BIT_LOG_FILE
	BIT_SHORT_FILE
	BIT_LEVEL
	BIT_STD_FLAG = BIT_DATA | BIT_TIME
	BIT_DEFAULT  = BIT_LEVEL | BIT_SHORT_FILE | BIT_STD_FLAG
)

const (
	LOG_DEBUG = iota
	LOG_INFO
	LOG_WARN
	LOG_ERROR
	LOG_PANIC
	LOG_FATAL
)

var levels = []string{
	"\033[32m [DEBUG] \033[0m",
	"\033[36m [INFO ] \033[0m",
	"\033[33m [WARN ] \033[0m",
	"\033[31m [ERROR] \033[0m",
	"\033[35m [PANIC] \033[0m",
	"\033[35m [FATAL] \033[0m"}

type Logger struct {
	sync.Mutex
	prefix    string
	flag      uint8
	out       io.Writer
	buf       bytes.Buffer
	file      *os.File
	isDebug   bool
	callDepth int
}

func NewLogger(out io.Writer, prefix string, flag uint8) *Logger {
	l := &Logger{out: out, prefix: prefix, flag: flag, callDepth: 2}
	runtime.SetFinalizer(l, cleanLogger)
	return l
}

func cleanLogger(log *Logger) {
	log.closeFile()
}

func (l *Logger) check(flag uint8) bool {
	if l.flag&flag != 0 {
		return true
	}
	return false
}

func (l *Logger) formatHead(buf *bytes.Buffer, t time.Time, file string, line, level int) {

	if l.prefix != "" {
		buf.WriteByte('[')
		buf.WriteString(l.prefix)
		buf.WriteByte(']')
	}

	if l.check(BIT_DATA | BIT_TIME | BIT_MICRO_SECONDS) {
		if l.check(BIT_DATA) {
			y, m, d := t.Date()
			stuff(buf, y, 4)
			buf.WriteByte('-')
			stuff(buf, int(m), 2)
			buf.WriteByte('-')
			stuff(buf, d, 2)
			buf.WriteByte(' ')
		}
		if l.check(BIT_TIME | BIT_MICRO_SECONDS) {
			h, m, s := t.Clock()
			stuff(buf, h, 2)
			buf.WriteByte(':')
			stuff(buf, m, 2)
			buf.WriteByte(':')
			stuff(buf, s, 2)
			if l.check(BIT_MICRO_SECONDS) {
				buf.WriteByte('.')
				stuff(buf, t.Nanosecond()/1e3, 6)
			}
			buf.WriteByte(' ')
		}
		if l.check(BIT_LEVEL) {
			buf.WriteString(levels[level])
		}
		if l.check(BIT_SHORT_FILE | BIT_LOG_FILE) {
			if l.check(BIT_SHORT_FILE) {
				file = path.Base(file)
			}
			buf.WriteString(file)
			buf.WriteByte(':')
			stuff(buf, line, -1)
			buf.WriteString(": ")
		}
	}
}

func (l *Logger) output(level int, s string) error {
	now := time.Now()
	var file string
	var line int
	l.Lock()
	defer l.Unlock()
	if l.check(BIT_SHORT_FILE | BIT_LOG_FILE) {
		l.Unlock()
		var ok bool
		if _, file, line, ok = runtime.Caller(l.callDepth); !ok {
			file = "unknown-file"
		}
		l.Lock()
	}
	l.buf.Reset()
	l.formatHead(&l.buf, now, file, line, level)
	l.buf.WriteString(s)
	if len(s) > 0 && s[len(s)-1] != '\n' {
		l.buf.WriteByte('\n')
	}
	//writer IO
	_, err := l.out.Write(l.buf.Bytes())
	return err
}

func stuff(buf *bytes.Buffer, i int, width int) {
	if i == 0 && width <= 1 {
		buf.WriteByte('0')
		return
	}
	var b [32]byte
	bp := len(b)
	for ; i > 0 || width > 0; i /= 10 {
		bp--
		width--
		b[bp] = byte(i%10) + '0' //int + int
	}
	for bp < len(b) {
		buf.WriteByte(b[bp])
		bp++
	}
}

func (l *Logger) Stack(a ...interface{}) {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprint(a...))
	sb.WriteByte('\n')
	buf := make([]byte, LOG_MAX_BUF)
	n := runtime.Stack(buf, true)
	sb.Write(buf[:n])
	sb.WriteByte('\n')
	l.output(LOG_ERROR, sb.String())
}

func (l *Logger) Flags() uint8 {
	l.Lock()
	defer l.Unlock()
	return l.flag
}

func (l *Logger) ResetFlags(flag uint8) {
	l.Lock()
	defer l.Unlock()
	l.flag = flag
}

func (l *Logger) AddFlag(flag uint8) {
	l.Lock()
	defer l.Unlock()
	l.flag |= flag
}

func (l *Logger) SetPrefix(prefix string) {
	l.Lock()
	defer l.Unlock()
	l.prefix = prefix
}

func (l *Logger) SetLogFile(logfile string) (err error) {
	var file *os.File
	if err = mkdirLog(path.Dir(logfile)); err != nil {
		return
	}
	if l.checkFileExists(logfile) {
		file, err = os.OpenFile(logfile, os.O_APPEND|os.O_RDWR, 0644)
	} else {
		file, err = os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	}
	if err != nil {
		return err
	}
	l.Lock()
	defer l.Unlock()
	l.closeFile()
	l.file, l.out = file, file
	return nil
}

func (l *Logger) closeFile() {
	if l.file != nil {
		l.file.Close()
		l.file = nil
		l.out = os.Stderr
	}
}

func (l *Logger) checkFileExists(logFile string) bool {
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		return false
	}
	return true
}

func mkdirLog(dir string) error {
	if _, err := os.Stat(dir); err != nil {
		if !os.IsExist(err) {
			if err = os.Mkdir(dir, 0755); err != nil {
				if os.IsPermission(err) {
					return err
				}
			}
		}
	}
	return nil
}

func (l *Logger) closeDebug() {
	l.isDebug = true
}

func (l *Logger) openDebug() {
	l.isDebug = false
}

func (l *Logger) Debugf(format string, a ...interface{}) {
	if !l.isDebug {
		l.output(LOG_DEBUG, fmt.Sprintf(format, a...))
	}
}

func (l *Logger) Debug(a ...interface{}) {
	if !l.isDebug {
		l.output(LOG_DEBUG, fmt.Sprint(a...))
	}
}

func (l *Logger) Infof(format string, a ...interface{}) {
	l.output(LOG_INFO, fmt.Sprintf(format, a...))
}

func (l *Logger) Info(a ...interface{}) {
	l.output(LOG_INFO, fmt.Sprint(a...))
}

func (l *Logger) Warnf(format string, a ...interface{}) {
	l.output(LOG_WARN, fmt.Sprintf(format, a...))
}

func (l *Logger) Warn(a ...interface{}) {
	l.output(LOG_WARN, fmt.Sprint(a...))
}

func (l *Logger) Errorf(format string, a ...interface{}) {
	l.output(LOG_ERROR, fmt.Sprintf(format, a...))
}

func (l *Logger) Error(a ...interface{}) {
	l.output(LOG_ERROR, fmt.Sprint(a...))
}

func (l *Logger) Fatalf(format string, a ...interface{}) {
	l.output(LOG_FATAL, fmt.Sprintf(format, a...))
	os.Exit(1)
}

func (l *Logger) Fatal(a ...interface{}) {
	l.output(LOG_FATAL, fmt.Sprint(a...))
	os.Exit(1)
}

func (l *Logger) Panicf(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	l.output(LOG_PANIC, s)
	panic(s)
}

func (l *Logger) Panic(a ...interface{}) {
	s := fmt.Sprint(a...)
	l.output(LOG_PANIC, s)
	panic(s)
}
