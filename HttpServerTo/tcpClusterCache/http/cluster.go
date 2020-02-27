package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type clusterHandler struct {
	*Server
}

func (h *clusterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if data, err := json.Marshal(h.Members()); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		w.Write(data)
	}
}

func (s *Server) clusterHandle() http.Handler {
	return &clusterHandler{s}
}