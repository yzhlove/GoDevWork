package src

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/time/rate"
)

func GolangRateMiddle(limit *rate.Limiter) endpoint.Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !limit.Allow() {
				return nil, errors.New("limit request")
			}
			return e(ctx, request)
		}
	}
}
