package cache

type Stat struct {
	Count     uint32 `json:"count"`
	KeySize   uint32 `json:"key_size"`
	ValueSize uint32 `json:"value_size"`
}

func (s *Stat) add(key string, value []byte) {
	s.Count += 1
	s.KeySize += uint32(len(key))
	s.ValueSize += uint32(len(value))
}

func (s *Stat) del(key string, value []byte) {
	s.Count -= 1
	s.KeySize -= uint32(len(key))
	s.ValueSize -= uint32(len(value))
}
