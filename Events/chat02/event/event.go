package event

type Event interface {
	ID() uint32
}

const (
	eventUnknown = iota
	eventLogin
	eventAchieve
)

type Unknown struct{}

func (Unknown) ID() uint32 {
	return eventUnknown
}

type Login struct{}

func (Login) ID() uint32 {
	return eventLogin
}

type Achieve struct {
	Id uint32
}

func (Achieve) ID() uint32 {
	return eventAchieve
}
