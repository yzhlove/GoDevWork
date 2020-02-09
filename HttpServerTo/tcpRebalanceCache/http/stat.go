package http

import (
	"encoding/json"
	"net/http"
)

type statusH struct {
	*Server
}

func (h *statusH) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) statusHandle() http.Handler {
	return &statusH{s}
}
