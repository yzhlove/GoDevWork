package app

import "sync"

type MutexType struct {
	mutex sync.RWMutex
	stat  bool
}

func (mt *MutexType) Get() bool {
	mt.mutex.RLock()
	defer mt.mutex.RUnlock()
	return mt.stat
}

func (mt *MutexType) Lock() {
	mt.mutex.RLock()
	mt.stat = true
}

func (mt *MutexType) Unlock() {
	mt.stat = false
	mt.mutex.Unlock()
}

type MutexTypeList struct {
	sync.Mutex
	queue map[string]*MutexType
}

func (mts *MutexTypeList) Get(code string) *MutexType {
	mts.Lock()
	defer mts.Unlock()

	if m, ok := mts.queue[code]; ok {
		return m
	}
	mt := &MutexType{}
	mts.queue[code] = mt
	return mt
}

func NewMutexTypeList() *MutexTypeList {
	return &MutexTypeList{queue: make(map[string]*MutexType)}
}
