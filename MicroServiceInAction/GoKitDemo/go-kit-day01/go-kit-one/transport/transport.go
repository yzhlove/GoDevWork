package transport

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	ept "go-kit-one/endpoint"
	"go-kit-one/service"
	"net/http"
	"strconv"
)

func decodeHttpAddRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var (
		in  service.Add
		err error
	)
	if in.A, err = strconv.Atoi(r.FormValue("a")); err != nil {
		return nil, err
	}
	if in.B, err = strconv.Atoi(r.FormValue("b")); err != nil {
		return nil, err
	}
	return in, nil
}

func encodeHttpGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
		errorEncode(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func errorEncode(ctx context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}

func NewHttpHandle(point ept.EndpointServer) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncode),
	}
	m := http.NewServeMux()
	m.Handle("/sum", httptransport.NewServer(
		point.AddEndpoint,
		decodeHttpAddRequest,
		encodeHttpGenericResponse,
		options...))
	return m
}

type errorWrapper struct {
	Error string `json:"errors"`
}
