package base

type LogMessage []interface{}
type LogCondMessage string

type LogBase interface {
	GetLogName() string
	Init() error
}

type LogWriter interface {
	LogBase
	Write(message LogMessage) (err error)
}

type LogReader interface {
	LogBase
	Read(cond LogCondMessage) (interface{}, error)
}
