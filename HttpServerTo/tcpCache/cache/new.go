package cache

import (
	"log"
	"strconv"
)

const (
	MEMORY = iota
)

func New(tpy int) (c Cache) {
	if tpy == MEMORY {
		c = newInMemoryCache()
	} else {
		panic("unknown cache type  " + strconv.Itoa(tpy))
	}
	log.Println(tpy, " read to server.")
	return
}
