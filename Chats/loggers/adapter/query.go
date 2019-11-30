package adapter

type QueryCond string

type LogQuery interface {
	Query(cond QueryCond) (interface{}, error)
	GetLogName() string
	Init() error
}
