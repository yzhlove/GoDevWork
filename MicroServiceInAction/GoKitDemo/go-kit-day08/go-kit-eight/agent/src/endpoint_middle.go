package src

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/time/rate"
)

func NewGoRateMid(l *rate.Limiter) endpoint.Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !l.Allow() {
				return nil, errors.New("limiter request")
			}
			return e(ctx, request)
		}
	}
}
