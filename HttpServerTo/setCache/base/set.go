package base

import (
	"strings"
	"time"
)

type SetData struct {
	ch chan string
	fi FileWriteInterface
}

func NewSetData(size int, fi FileWriteInterface) *SetData {
	return &SetData{make(chan string, size), fi}
}

func (s *SetData) GetChanString() chan string {
	return s.ch
}

func (s *SetData) Write() {
	defer s.fi.Stop()
	count := 0
	for str := range s.ch {
		s.fi.Writer(str + " ")
		if count != 0 && count%20 == 0 {
			s.fi.Writer("\n")
		}
		count++
	}
}

func (s *SetData) BufferWrite() {
	defer s.fi.Stop()
	sb := strings.Builder{}
	t := time.NewTimer(time.Second)
	count := 0
	for {
		select {
		case str, ok := <-s.ch:
			if ok {
				if count != 0 && count%20 == 0 {
					s.fi.Writer(sb.String())
					sb.Reset()
					//timer
					if !t.Stop() {
						<-t.C
					}
					t.Reset(time.Second)
				}
				sb.WriteString(str + " ")
				count++
			} else {
				return
			}
		case <-t.C:
			if sb.Len() != 0 {
				s.fi.Writer(sb.String())
				sb.Reset()
				t.Reset(time.Second)
			}
		}
	}
}
