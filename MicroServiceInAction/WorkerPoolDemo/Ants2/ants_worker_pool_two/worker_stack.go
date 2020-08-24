package ants_worker_pool_two

import (
	"sort"
	"time"
)

type Stack struct {
	entries       []*GoWorker
	expireEntries []*GoWorker
	size          int
}

func (s *Stack) Len() int {
	return len(s.expireEntries)
}

func (s *Stack) Swap(i, j int) {
	s.expireEntries[i], s.expireEntries[j] = s.expireEntries[j], s.expireEntries[i]
}

func (s *Stack) Less(i, j int) bool {
	if s.expireEntries[i].expire.Before(s.expireEntries[j].expire) {
		return true
	}
	return false
}

func NewStackQueue(size int) *Stack {
	return &Stack{
		entries: make([]*GoWorker, 0, size),
		size:    size,
	}
}

func (s *Stack) len() int {
	return len(s.entries)
}

func (s *Stack) empty() bool {
	return s.len() == 0
}

func (s *Stack) push(w *GoWorker) error {
	s.entries = append(s.entries, w)
	return nil
}

func (s *Stack) pop() *GoWorker {
	if s.empty() {
		return nil
	}
	w := s.entries[s.len()-1]
	s.entries = s.entries[:s.len()-1]
	return w
}

func (s *Stack) checkExpire(t time.Duration) []*GoWorker {
	if s.empty() {
		return nil
	}
	timeout := time.Now().Add(-t)
	sort.Sort(s)
	index := sort.Search(s.len(), func(i int) bool {
		return timeout.Before(s.entries[i].expire)
	})
	s.expireEntries = s.expireEntries[:0]
	if index != s.len() {
		s.expireEntries = append(s.expireEntries, s.entries[:index]...)
		s.entries = s.entries[:copy(s.entries, s.entries[index+1:])]
	}
	return s.expireEntries
}

func (s *Stack) reset() {
	for i := 0; i < s.len(); i++ {
		ClosedChan(s.entries[i].taskChan)
	}
	s.entries = s.entries[:0]
}
