package consistenthash

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
)

///////////////////////////////////////////////
// 一致性哈希
///////////////////////////////////////////////

type H func(data []byte) uint32

type Map struct {
	hash     H
	replicas int
	keys     []int
	hashMap  map[int]string
}

func NewConsistentHash(replicas int, fn H) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

func (m *Map) Set(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			bucket := int(m.hash([]byte(strconv.Itoa(i) + key)))
			fmt.Println("set --> ", bucket)
			m.keys = append(m.keys, bucket)
			m.hashMap[bucket] = key
		}
	}
	sort.Ints(m.keys)
}

func (m *Map) Get(key string) string {
	if len(m.keys) != 0 {
		bucket := int(m.hash([]byte(key)))
		idx := sort.Search(len(m.keys), func(i int) bool {
			return m.keys[i] >= bucket
		})
		return m.hashMap[m.keys[idx%len(m.keys)]]
	}
	return ""
}
