package src

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"golang.org/x/time/rate"
)

func NewGoRateMiddle(limit *rate.Limiter) endpoint.Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !limit.Allow() {
				return nil, errors.New("rate limit allow")
			}
			return e(ctx, request)
		}
	}
}

func NewTracerEndpointMiddle(tracer opentracing.Tracer) endpoint.Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			span, ctx := opentracing.StartSpanFromContextWithTracer(ctx, tracer, "endpoint", opentracing.Tag{
				Key:   string(ext.Component),
				Value: "NewTracerEndpointMiddle",
			})
			defer span.Finish()
			return e(ctx, request)
		}
	}
}
