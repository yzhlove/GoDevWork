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

type BackupOption struct {
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
}

type LoggerOptions struct {
	BackupOption
	zap.Config
	LogFileDir    string
	LogPrefix     string
	ErrorFilePath string
	WarnFilePath  string
	InfoFilePath  string
	DebugFilePath string
	Level         zapcore.Level
	IsDevelopment bool
}

type ModeOption func(opt *LoggerOptions)

type Logger struct {
	*zap.Logger
	sync.RWMutex
	Opts   *LoggerOptions `json:"opts"`
	zapCfg zap.Config
	init   bool
}

var (
	_logger                        *Logger
	sp                             = string(filepath.Separator) //分割符
	errWs, warnWs, infoWs, debugWs zapcore.WriteSyncer
	debugConsoleWs                 = zapcore.Lock(os.Stdout)
	errorConsoleWs                 = zapcore.Lock(os.Stderr)
)

func NewLogger(opts ...ModeOption) *zap.Logger {
	_logger = &Logger{}
	_logger.Lock()
	defer _logger.Unlock()

	if _logger.init {
		_logger.Info("[NewLogger] logger init")
		return nil
	}

	_logger.Opts = defaultLoggerOptions()
	if _logger.Opts.LogFileDir == "" {
		_logger.Opts.LogFileDir, _ = filepath.Abs(filepath.Dir(filepath.Join(".")))
		_logger.Opts.LogFileDir += sp + "logs" + sp
	}

	if _logger.Opts.IsDevelopment {
		_logger.zapCfg = zap.NewDevelopmentConfig()
		_logger.zapCfg.EncoderConfig.EncodeTime = timeEncoder
	} else {
		_logger.zapCfg = zap.NewProductionConfig()
		_logger.zapCfg.EncoderConfig.EncodeTime = timeUnixNano
	}

	if len(_logger.zapCfg.OutputPaths) == 0 {
		_logger.zapCfg.OutputPaths = []string{"stdout"}
	}
	if len(_logger.zapCfg.OutputPaths) == 0 {
		_logger.zapCfg.OutputPaths = []string{"stderr"}
	}

	for _, fn := range opts {
		fn(_logger.Opts)
	}

	_logger.zapCfg.Level.SetLevel(_logger.Opts.Level)
	_logger.initWs()
	_logger.init = true
	_logger.Info("[NewLogger] succeed!")
	return _logger.Logger
}

func (l *Logger) initWs() {
	l.setSynced()
	var err error
	if l.Logger, err = l.zapCfg.Build(l.cores()); err != nil {
		panic(err)
	}
	defer l.Logger.Sync()
}

func (l *Logger) setSynced() {
	name := l.Opts.LogFileDir + sp + l.Opts.LogPrefix + "-"
	set := func(fn string) zapcore.WriteSyncer {
		return zapcore.AddSync(&lumberjack.Logger{
			Filename:   name + fn,
			MaxSize:    l.Opts.MaxSize,
			MaxBackups: l.Opts.MaxBackups,
			MaxAge:     l.Opts.MaxAge,
			Compress:   l.Opts.Compress,
			LocalTime:  true,
		})
	}
	errWs = set(l.Opts.ErrorFilePath)
	warnWs = set(l.Opts.WarnFilePath)
	infoWs = set(l.Opts.InfoFilePath)
	debugWs = set(l.Opts.DebugFilePath)
}

func (l *Logger) cores() zap.Option {
	fileEncoder := zapcore.NewJSONEncoder(l.zapCfg.EncoderConfig)
	cfg := zap.NewDevelopmentEncoderConfig()
	cfg.EncodeTime = timeEncoder
	cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(cfg)

	fn := func(level zapcore.Level) bool {
		return true
	}
	cores := []zapcore.Core{
		zapcore.NewCore(fileEncoder, errWs, zap.LevelEnablerFunc(fn)),
		zapcore.NewCore(fileEncoder, warnWs, zap.LevelEnablerFunc(fn)),
		zapcore.NewCore(fileEncoder, infoWs, zap.LevelEnablerFunc(fn)),
		zapcore.NewCore(fileEncoder, debugWs, zap.LevelEnablerFunc(fn)),
	}

	if l.Opts.IsDevelopment {
		cores = append(cores, []zapcore.Core{
			zapcore.NewCore(consoleEncoder, errorConsoleWs, zap.LevelEnablerFunc(fn)),
			zapcore.NewCore(consoleEncoder, debugConsoleWs, zap.LevelEnablerFunc(fn)),
		}...)
	}

	return zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.NewTee(cores...)
	})
}

func SetMaxSize(size int) ModeOption {
	return func(opt *LoggerOptions) {
		opt.MaxSize = size
	}
}

func SetMaxBackups(backup int) ModeOption {
	return func(opt *LoggerOptions) {
		opt.MaxBackups = backup
	}
}

func SetMaxAge(age int) ModeOption {
	return func(opt *LoggerOptions) {
		opt.MaxAge = age
	}
}

func SetCompress(compress bool) ModeOption {
	return func(opt *LoggerOptions) {
		opt.Compress = compress
	}
}

func SetLogFileDir(dir string) ModeOption {
	return func(opt *LoggerOptions) {
		opt.LogFileDir = dir
	}
}

func SetLogPrefix(prefix string) ModeOption {
	return func(opt *LoggerOptions) {
		opt.LogPrefix = prefix
	}
}

func SetLevel(level zapcore.Level) ModeOption {
	return func(opt *LoggerOptions) {
		opt.Level = level
	}
}

func SetIsDevelopment(isDev bool) ModeOption {
	return func(opt *LoggerOptions) {
		opt.IsDevelopment = isDev
	}
}

func SetErrorFilePath(pth string) ModeOption {
	return func(opt *LoggerOptions) {
		opt.ErrorFilePath = pth
	}
}

func SetWarnFilePath(pth string) ModeOption {
	return func(opt *LoggerOptions) {
		opt.WarnFilePath = pth
	}
}

func SetInfoFilePath(pth string) ModeOption {
	return func(opt *LoggerOptions) {
		opt.InfoFilePath = pth
	}
}

func SetDebugFilePath(pth string) ModeOption {
	return func(opt *LoggerOptions) {
		opt.DebugFilePath = pth
	}
}

func defaultLoggerOptions() *LoggerOptions {
	return &LoggerOptions{
		LogFileDir:    "",
		LogPrefix:     "app_log",
		ErrorFilePath: "error.log",
		WarnFilePath:  "warn.log",
		InfoFilePath:  "info.log",
		DebugFilePath: "debug.log",
		Level:         zapcore.DebugLevel,
		BackupOption: BackupOption{
			MaxSize:    100,
			MaxBackups: 60,
			MaxAge:     30,
			Compress:   true,
		},
	}
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func timeUnixNano(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt64(t.UnixNano() / 1e6)
}
