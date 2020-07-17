package transport

import (
	"context"
	"encoding/json"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	ept "go-kit-three/endpoint"
	"go-kit-three/service"
	"go-kit-three/utils"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func NewHttpHandler(point ept.Server, log *zap.Logger) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter) {
			log.Warn(fmt.Sprint(ctx.Value(service.ContextReqUUID)), zap.Error(err))
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(errorWrap{Error: err.Error()})
		}),
		httptransport.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
			UID := utils.GetUID()
			log.Debug("GetUID", zap.String("UID", UID))
			ctx = context.WithValue(ctx, service.ContextReqUUID, UID)
			ctx = context.WithValue(ctx, utils.JwtContextKey, request.Header.Get("X-Token"))
			log.Debug("GetTokenString", zap.String("X-Token", request.Header.Get("X-Token")))
			return ctx
		}),
	}
	h := http.NewServeMux()
	h.Handle("/sum", httptransport.NewServer(
		point.Add,
		decoderHttpAddRequest,
		encoderHttpGenericResponse,
		options...,
	))
	h.Handle("/login", httptransport.NewServer(
		point.Login,
		decoderHttpLoginRequest,
		encoderHttpGenericResponse,
		options...,
	))
	return h
}

func decoderHttpLoginRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var login service.Login
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		return nil, err
	}
	utils.GetLog().Debug(fmt.Sprint(service.ContextReqUUID), zap.Any("decoder login", login))
	return login, nil
}

func decoderHttpAddRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var in service.Add
	var err error
	if in.A, err = strconv.Atoi(r.FormValue("a")); err != nil {
		return nil, err
	}
	if in.B, err = strconv.Atoi(r.FormValue("b")); err != nil {
		return nil, err
	}
	utils.GetLog().Debug(fmt.Sprint(ctx.Value(service.ContextReqUUID)), zap.Any("decoder add", in))
	return in, err
}

func encoderHttpGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	utils.GetLog().Debug(fmt.Sprint(ctx.Value(service.ContextReqUUID)),
		zap.Any("encoderHttpResponse", response))
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorWrap struct {
	Error string `json:"errors"`
}
