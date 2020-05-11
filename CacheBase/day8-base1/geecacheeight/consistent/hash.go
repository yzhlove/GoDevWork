package consistent

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type H func(data []byte) uint32

type Map struct {
	hashFunc H
	replicas int
	keys     []int
	hashMap  map[int]string
}

func NewConsistent(replicas int, fn H) *Map {
	m := &Map{
		hashFunc: fn,
		replicas: replicas,
		hashMap:  make(map[int]string),
	}
	if m.hashFunc == nil {
		m.hashFunc = crc32.ChecksumIEEE
	}
	return m
}

func (m *Map) Set(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			bucket := int(m.hashFunc([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, bucket)
			m.hashMap[bucket] = key
		}
	}
	sort.Ints(m.keys)
}

func (m *Map) Get(key string) string {
	if len(m.keys) != 0 {
		bucket := int(m.hashFunc([]byte(key)))
		length := len(m.keys)
		idx := sort.Search(length, func(i int) bool {
			return m.keys[i] >= bucket
		})
		return m.hashMap[m.keys[idx%length]]
	}
	return ""
}
