package data

type condition struct {
	Name   string
	Params []uint32
}

type achieveRecord struct {
	Id         uint32
	Conditions []condition
}

type Achieve struct {
	data map[uint32]*achieveRecord
}

func (a *Achieve) Get(id uint32) *achieveRecord {
	return a.data[id]
}

func (a *Achieve) All() []*achieveRecord {
	records := make([]*achieveRecord, len(a.data))
	index := 0
	for _, record := range a.data {
		records[index] = record
		index++
	}
	return records
}

func Construct() *Achieve {
	return &Achieve{data: make(map[uint32]*achieveRecord)}
}
