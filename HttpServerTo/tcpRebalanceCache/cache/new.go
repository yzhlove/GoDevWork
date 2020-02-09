package cache

import (
	"log"
	"strconv"
)

const (
	MEMORY = iota + 1
)

func New(typ int) Cache {
	if typ == MEMORY {
		log.Println("[cache read to server ok.]")
		return newInMemory()
	}
	panic("unknown cache type : " + strconv.Itoa(typ))
	return nil
}
