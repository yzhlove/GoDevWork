package base

import (
	"fmt"
	"strings"
)

type DataSource struct {
	ch chan string
}

func NewDataSource(ch chan string) *DataSource {
	return &DataSource{ch: ch}
}

func (s DataSource) Stop() {
	close(s.ch)
}

func (s DataSource) GenerateSource(size int) {
	str := strings.Repeat("a", 5)
	go func() {
		defer s.Stop()
		for i := 0; i < size; i++ {
			s.ch <- str
		}
	}()
	fmt.Println("data source generate ok.")
}
