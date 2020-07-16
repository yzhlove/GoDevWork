package tool

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Options struct {
	LogFileDir    string //文件保存地方
	LogPrefix     string //日志前缀
	ErrorFileName string
	WarnFileName  string
	InfoFileName  string
	DebugFileName string
	Level         zapcore.Level //日志等级
	MaxSize       int           //日志文件大小
	MaxBackups    int           //最多存在多少个切片文件
	MaxAge        int           //保存的最大天数
	Development   bool          //是否是开发这模式
	zap.Config
}

type ModeOption func(option *Options)

var (
	l                              *Logger
	sp                             = string(filepath.Separator)
	errWs, warnWs, infoWs, debugWs zapcore.WriteSyncer
	debugConsoleWs                 = zapcore.Lock(os.Stdout)
	errorConsoleWs                 = zapcore.Lock(os.Stderr)
)

type Logger struct {
	*zap.Logger
	sync.RWMutex
	Opts   *Options `json:"opts"`
	zapCfg zap.Config
	inited bool
}

func NewLogger(opts ...ModeOption) *zap.Logger {
	l = &Logger{}
	l.Lock()
	defer l.Unlock()
	if l.inited {
		l.Info("[NewLogger] logger Inited")
		return nil
	}
	l.Opts = GetDefaultOpts()
	if l.Opts.LogFileDir == "" {
		l.Opts.LogFileDir, _ = filepath.Abs(filepath.Dir(filepath.Join(".")))
		l.Opts.LogFileDir += sp + "logs" + sp
	}
	if l.Opts.Development {
		l.zapCfg = zap.NewDevelopmentConfig()
		l.zapCfg.EncoderConfig.EncodeTime = timeEncoder
	} else {
		l.zapCfg = zap.NewProductionConfig()
		l.zapCfg.EncoderConfig.EncodeTime = timeUnixNano
	}

	if len(l.Opts.OutputPaths) == 0 {
		l.zapCfg.OutputPaths = []string{"stdout"}
	}
	if len(l.Opts.ErrorOutputPaths) == 0 {
		l.zapCfg.ErrorOutputPaths = []string{"stderr"}
	}

	for _, fn := range opts {
		fn(l.Opts)
	}

	l.zapCfg.Level.SetLevel(l.Opts.Level)
	l.init()
	l.inited = true
	l.Info("[NewLogger] success")

	return l.Logger
}

func (l *Logger) init() {
	l.setSyncers()
	var err error
	if l.Logger, err = l.zapCfg.Build(l.cores()); err != nil {
		panic(err)
	}
	defer l.Logger.Sync()
}

func (l *Logger) setSyncers() {
	f := func(fileName string) zapcore.WriteSyncer {
		return zapcore.AddSync(&lumberjack.Logger{
			Filename:   l.Opts.LogFileDir + sp + l.Opts.LogPrefix + "-" + fileName,
			MaxSize:    l.Opts.MaxSize,
			MaxAge:     l.Opts.MaxAge,
			MaxBackups: l.Opts.MaxBackups,
			Compress:   true,
			LocalTime:  true,
		})
	}
	errWs = f(l.Opts.ErrorFileName)
	warnWs = f(l.Opts.WarnFileName)
	infoWs = f(l.Opts.InfoFileName)
	debugWs = f(l.Opts.DebugFileName)
}

func (l *Logger) cores() zap.Option {
	fileEncoder := zapcore.NewJSONEncoder(l.zapCfg.EncoderConfig)
	//consoleEncoder := zapcore.NewConsoleEncoder(l.zapCfg.EncoderConfig)
	encoderCfg := zap.NewDevelopmentEncoderConfig()
	encoderCfg.EncodeTime = timeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderCfg)

	errPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.ErrorLevel && zapcore.ErrorLevel-l.zapCfg.Level.Level() > -1
	})

	warnPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.WarnLevel && zapcore.WarnLevel-l.zapCfg.Level.Level() > -1
	})

	infoPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.InfoLevel && zap.InfoLevel-l.zapCfg.Level.Level() > -1
	})

	debugPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.DebugLevel && zap.DebugLevel-l.zapCfg.Level.Level() > -1
	})

	cores := []zapcore.Core{
		zapcore.NewCore(fileEncoder, errWs, errPriority),
		zapcore.NewCore(fileEncoder, warnWs, warnPriority),
		zapcore.NewCore(fileEncoder, infoWs, infoPriority),
		zapcore.NewCore(fileEncoder, debugWs, debugPriority),
	}

	if l.Opts.Development {
		cores = append(cores, []zapcore.Core{
			zapcore.NewCore(consoleEncoder, errorConsoleWs, errPriority),
			zapcore.NewCore(consoleEncoder, debugConsoleWs, warnPriority),
			zapcore.NewCore(consoleEncoder, debugConsoleWs, infoPriority),
			zapcore.NewCore(consoleEncoder, debugConsoleWs, debugPriority),
		}...)
	}

	return zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.NewTee(cores...)
	})
}

func SetMaxSize(size int) ModeOption {
	return func(option *Options) {
		option.MaxSize = size
	}
}

func SetMaxBackups(backup int) ModeOption {
	return func(option *Options) {
		option.MaxBackups = backup
	}
}

func SetMaxAge(age int) ModeOption {
	return func(option *Options) {
		option.MaxAge = age
	}
}

func SetLogFileDir(logFileDir string) ModeOption {
	return func(option *Options) {
		option.LogFileDir = logFileDir
	}
}

func SetLogPrefix(prefix string) ModeOption {
	return func(option *Options) {
		option.LogPrefix = prefix
	}
}

func SetLevel(level zapcore.Level) ModeOption {
	return func(option *Options) {
		option.Level = level
	}
}

func SetErrorFileName(fileName string) ModeOption {
	return func(option *Options) {
		option.ErrorFileName = fileName
	}
}

func SetWarnFileName(fileName string) ModeOption {
	return func(option *Options) {
		option.WarnFileName = fileName
	}
}

func SetInfoFileName(fileName string) ModeOption {
	return func(option *Options) {
		option.InfoFileName = fileName
	}
}

func SetDebugFileName(fileName string) ModeOption {
	return func(option *Options) {
		option.DebugFileName = fileName
	}
}

func SetDevelopment(development bool) ModeOption {
	return func(option *Options) {
		option.Development = development
	}
}

func GetDefaultOpts() *Options {
	return &Options{
		LogFileDir:    "",
		LogPrefix:     "app_log",
		ErrorFileName: "error.log",
		WarnFileName:  "warn.log",
		InfoFileName:  "info.log",
		DebugFileName: "debug.log",
		Level:         zapcore.DebugLevel,
		MaxSize:       100,
		MaxBackups:    60,
		MaxAge:        30,
	}
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func timeUnixNano(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt64(t.UnixNano() / 1e6)
}
