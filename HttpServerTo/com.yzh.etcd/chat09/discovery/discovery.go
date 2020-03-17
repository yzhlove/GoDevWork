package discovery

const (
	PUT = iota + 1
	DELETE
)

type Discovery interface {
	Init() error
	Register(key, value string)
	Watcher() chan Event
}

type Event struct {
	Action int
	Key    string
	Addr   string
}
