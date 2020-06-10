package zlog

import "os"

var StanderLogger = NewLogger(os.Stderr, "", BIT_DEFAULT)

func init() {
	StanderLogger.callDepth = 3
}

func Flags() uint8 {
	return StanderLogger.Flags()
}

func ResetFlag(flag uint8) {
	StanderLogger.ResetFlags(flag)
}

func AddFlag(flag uint8) {
	StanderLogger.AddFlag(flag)
}

func SetPrefix(prefix string) {
	StanderLogger.SetPrefix(prefix)
}

func SetLoggerFile(logfile string) error {
	return StanderLogger.SetLogFile(logfile)
}

func CloseDebug() {
	StanderLogger.closeDebug()
}

func OpenDebug() {
	StanderLogger.openDebug()
}

func Debugf(format string, a ...interface{}) {
	StanderLogger.Debugf(format, a...)
}

func Debug(a ...interface{}) {
	StanderLogger.Debug(a...)
}

func Infof(format string, a ...interface{}) {
	StanderLogger.Infof(format, a...)
}

func Info(a ...interface{}) {
	StanderLogger.Info(a...)
}

func Warnf(format string, a ...interface{}) {
	StanderLogger.Warnf(format, a...)
}

func Warn(a ...interface{}) {
	StanderLogger.Warn(a...)
}

func Errorf(format string, a ...interface{}) {
	StanderLogger.Errorf(format, a...)
}

func Error(a ...interface{}) {
	StanderLogger.Error(a...)
}

func Panicf(format string, a ...interface{}) {
	StanderLogger.Panicf(format, a...)
}

func Panic(a ...interface{}) {
	StanderLogger.Panic(a...)
}

func Fatalf(format string, a ...interface{}) {
	StanderLogger.Fatalf(format, a...)
}

func Fatal(a ...interface{}) {
	StanderLogger.Fatal(a...)
}

func Stack(a ...interface{}) {
	StanderLogger.Stack(a...)
}
