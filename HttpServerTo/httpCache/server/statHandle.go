package server

import (
	"encoding/json"
	"net/http"
)

type statusHandler struct {
	*Server
}

func (h *statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if body, err := json.Marshal(h.GetStat()); err != nil {
		respErr(w, err)
		return
	} else {
		_, _ = w.Write(body)
	}
}

func (s *Server) statusHand() http.Handler {
	return &statusHandler{s}
}
