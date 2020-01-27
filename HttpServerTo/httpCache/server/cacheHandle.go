package server

import (
	"errors"
	"log"
	"net/http"
	"strings"
)

type cacheHandler struct {
	*Server
}

func respErr(w http.ResponseWriter, err error) {
	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
}

func getUrls(r *http.Request) (key, value string, err error) {
	path := r.URL.EscapedPath()
	params := make([]string, 0, 2)
	for _, v := range strings.Split(path, "/") {
		params = append(params, strings.Trim(v, " "))
	}
	if len(params) > 3 || len(params) < 1 {
		err = errors.New("invalid url")
		return
	}
	key = params[1]
	if len(params) == 3 {
		value = params[2]
	}
	log.Println("params ==> ", params)
	return
}

func (h *cacheHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key, value, err := getUrls(r)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	switch r.Method {
	case http.MethodPut:
		if len(value) != 0 {
			if err := h.Set(key, []byte(value)); err != nil {
				respErr(w, err)
				return
			}
		}
	case http.MethodGet:
		if value, err := h.Get(key); err != nil {
			respErr(w, err)
			return
		} else {
			if len(value) > 0 {
				_, _ = w.Write(value)
			} else {
				w.WriteHeader(http.StatusNotFound)
				return
			}
		}
	case http.MethodDelete:
		if err := h.Del(key); err != nil {
			respErr(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *Server) cacheHand() http.Handler {
	return &cacheHandler{s}
}
