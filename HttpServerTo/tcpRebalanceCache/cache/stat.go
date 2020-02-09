package cache

type Stat struct {
	Count     uint32 `json:"count"`
	KeySize   uint32 `json:"key_size"`
	ValueSize uint32 `json:"value_size"`
}

func (s *Stat) add(k string, v []byte) {
	s.Count += 1
	s.KeySize += uint32(len(k))
	s.ValueSize += uint32(len(v))
}

func (s *Stat) del(k string, v []byte) {
	s.Count -= 1
	s.KeySize -= uint32(len(k))
	s.ValueSize -= uint32(len(v))
}
