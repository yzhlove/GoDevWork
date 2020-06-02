package core

import (
	"sync"
	"time"
)

///////////////////////////////////////////////////////////
// snowflake 雪花算法
// 1 - 41 - 10 - 12
// 1:最高位为0，保证生成的唯一是正数
// 41:毫秒级的时间戳
// 10:机器id，最多支持(1024)
// 12:id序列号的最大值(4096)
///////////////////////////////////////////////////////////

const (
	NumberBits uint8 = 12
	IDBits     uint8 = 10
	MaxNumber  int64 = -1 ^ (-1 << NumberBits) //4095
	Epoch      int64 = 1525705533000
	//MaxID      int64 = -1 ^ (-1 << IDBits)     //1023
)

var (
	WorkerID   int64 = 1
	_snowflake *snowflake
	once       sync.Once
)

type snowflake struct {
	sync.Mutex
	workerID  int64
	timestamp int64
	index     int64
}

func (s *snowflake) generate() int64 {
	s.Lock()
	defer s.Unlock()
	now := time.Now().UnixNano() / 1e6
	if s.timestamp == now {
		s.index++
		if s.index >= MaxNumber {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.index, s.timestamp = 0, now
	}
	return (now-Epoch)<<(NumberBits+IDBits) | s.workerID<<NumberBits | s.index
}

func Get() int64 {
	once.Do(func() {
		_snowflake = &snowflake{workerID: WorkerID}
	})
	return _snowflake.generate()
}
