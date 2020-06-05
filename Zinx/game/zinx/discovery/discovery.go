package discovery

type EventType int

const (
	EventPut EventType = 1 //新增
	EventDel EventType = 2 //删除
)

type RegisterImp interface {
	Register(key, value string)
}

type DiscoverImp interface {
	Watcher() <-chan Event
}

type RegisterAndDiscoverImp interface {
	RegisterImp
	DiscoverImp
}

type Event struct {
	Action    EventType
	Key, Addr string
}
