package http

import (
	"bytes"
	"net/http"
)

type rebalancedH struct {
	*Server
}

func (s *rebalancedH) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	go s.rebalance()
}

func (s *rebalancedH) rebalance() {
	scan := s.NewScanner()
	defer scan.Close()
	c := http.Client{}
	for scan.Scan() {
		key := scan.Key()
		addr, ok := s.ShouldProcess(key)
		if !ok {
			url := "http://" + addr + ":1234/cache/" + key
			r, _ := http.NewRequest(http.MethodPut, url, bytes.NewReader(scan.Value()))
			c.Do(r)
			s.Del(key)
		}
	}
}

func (s *Server) rebalancedHandle() http.Handler {
	return &rebalancedH{s}
}
