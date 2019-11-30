package adapter

type LogMessage []interface{}

type LogWriter interface {
	Write(message LogMessage) error
	GetLogName() string
	Init() error
}

var adapters []LogWriter

func Register(adapter ...LogWriter) {
	for _, inter := range adapter {
		adapters = append(adapters, inter)
	}
}

//获取所有的适配器
func GetAdapter() []LogWriter {
	return adapters
}
