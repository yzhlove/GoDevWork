package transport

import (
	"context"
	"encoding/json"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	ept "go-kit-two/endpoint"
	"go-kit-two/service"
	"go-kit-two/utils"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func NewHttpHandler(endpoint ept.EndpointServer, log *zap.Logger) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter) {
			log.Warn(fmt.Sprint(ctx.Value(service.ContextReqUUID)), zap.Error(err))
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
		}),
	}
	httptransport.ServerErrorHandler(NewZapLogErrorHandler(log))
	httptransport.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
		UUID := utils.GetUUID()
		log.Debug("add uuid", zap.String("UUID", UUID))
		ctx = context.WithValue(ctx, service.ContextReqUUID, UUID)
		return ctx
	})
	m := http.NewServeMux()
	m.Handle("/sum", httptransport.NewServer(
		endpoint.AddEndpoint,
		decoderHttpAddRequest,
		encoderHttpGenericResponse,
		options...))
	return m
}

func decoderHttpAddRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var in service.Add
	var err error
	if in.A, err = strconv.Atoi(r.FormValue("a")); err != nil {
		return in, err
	}
	if in.B, err = strconv.Atoi(r.FormValue("b")); err != nil {
		return in, err
	}
	utils.GetLog().Debug(fmt.Sprint(ctx.Value(service.ContextReqUUID)), zap.Any("decoder", in))
	return in, nil
}

func encoderHttpGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	utils.GetLog().Debug(fmt.Sprint(ctx.Value(service.ContextReqUUID)), zap.Any("[encoder Generic] (request) over", response))
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorWrapper struct {
	Error string `json:"errors"`
}
