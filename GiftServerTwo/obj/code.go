package obj

import (
	"strings"
	"time"
)

//go:generate msgp -io=false -tests=false
type Item struct {
	Id  uint32
	Num int32
}

type Code struct {
	Id           uint32
	FixCode      string
	Num          uint32
	StartTime    int64
	EndTime      int64
	TimesPerCode uint16
	TimesPerUser uint16
	ZoneIds      []uint32
	Items        []Item
	Used         uint32
}

//msgp:ignore Entity
type Entity struct {
	AutoId uint32
	Infos  map[uint32]*Code
	Fixed  map[string]uint32
}

func (c *Code) ZoneCheck(zone uint32) bool {
	if len(c.ZoneIds) == 0 {
		return true
	}
	for _, z := range c.ZoneIds {
		if z == zone {
			return true
		}
	}
	return false
}

func (c *Code) Expired() bool {
	t := time.Now().Unix()
	if c.StartTime != 0 && t < c.StartTime {
		return false
	}
	if c.EndTime != 0 && t >= c.EndTime {
		return false
	}
	return true
}

func (e *Entity) Update(code *Code) {
	e.AutoId++
	e.Infos[code.Id] = code
	if fixed := strings.TrimSpace(code.FixCode); len(fixed) > 0 {
		e.Fixed[code.FixCode] = code.Id
	}
}
