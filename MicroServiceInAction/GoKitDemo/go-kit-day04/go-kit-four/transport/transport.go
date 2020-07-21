package transport

import (
	"context"
	"encoding/json"
	"fmt"
	tps "github.com/go-kit/kit/transport/http"
	ept "go-kit-four/endpoint"
	"go-kit-four/service"
	"go-kit-four/utils"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type errorWrap struct {
	Error string `json:"errors"`
}

func GetHttpHandle(e ept.Endpoint, log *zap.Logger) http.Handler {
	options := []tps.ServerOption{
		tps.ServerErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter) {
			log.Warn(fmt.Sprint(ctx.Value(service.ContextUID)),
				zap.Error(err))
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(errorWrap{Error: err.Error()})
		}),
		tps.ServerBefore(func(ctx context.Context, request *http.Request) context.Context {
			UID := utils.GetUID()
			log.Debug("set UID", zap.String("UID", UID))
			ctx = context.WithValue(ctx, service.ContextUID, UID)
			ctx = context.WithValue(ctx, utils.JwtContextKey, request.Header.Get("X-Token"))
			log.Debug("set Token", zap.String("TOKEN", request.Header.Get("X-Token")))
			return ctx
		}),
	}

	h := http.NewServeMux()
	h.Handle("/sum", tps.NewServer(e.Add, decodeHttpAdd, encodeHttpResp, options...))
	h.Handle("/login", tps.NewServer(e.Login, decodeHttpLogin, encodeHttpResp, options...))
	return h
}

func decodeHttpAdd(ctx context.Context, r *http.Request) (interface{}, error) {
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

func decodeHttpLogin(ctx context.Context, r *http.Request) (interface{}, error) {
	var login service.Login
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		return nil, err
	}
	utils.GetLog().Debug(fmt.Sprint(ctx.Value(service.ContextUID)),
		zap.Any("func-->", "transport.decodeHttpLogin"),
		zap.Any("login", login))
	return login, nil
}

func encodeHttpResp(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	utils.GetLog().Debug(fmt.Sprint(ctx.Value(service.ContextUID)),
		zap.Any("func-->", "transport.encodeHttpResp"),
		zap.Any("resp", resp))
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(resp)
}
