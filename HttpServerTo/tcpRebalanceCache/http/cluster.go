package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type clusterH struct {
	*Server
}

func (s *clusterH) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if data, err := json.Marshal(s.Members()); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		w.Write(data)
	}
}

func (s *Server) clusterHandle() http.Handler {
	return &clusterH{s}
}
