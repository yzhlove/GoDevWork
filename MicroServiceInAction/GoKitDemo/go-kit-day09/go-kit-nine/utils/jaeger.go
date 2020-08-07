package utils

import (
	"context"
	"errors"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"strings"
)

//////////////////////////////////////////////////////////////////
// Jaeger
//////////////////////////////////////////////////////////////////

func NewJaegerTracer(serviceName string) (tracer opentracing.Tracer, closer io.Closer, err error) {

	cfg := &jaegerConfig.Configuration{
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  "const", //固定采样
			Param: 1,       //全采样
		},
		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},
		ServiceName: serviceName,
	}

	if tracer, closer, err = cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger)); err != nil {
		return
	}
	opentracing.SetGlobalTracer(tracer)
	return
}

var TracingComponentTag = opentracing.Tag{Key: string(ext.Component), Value: "grpc"}

type MetaDataReaderWriter struct {
	metadata.MD
}

type ForeachFunc func(k, v string) error

func (m MetaDataReaderWriter) Foreach(h ForeachFunc) error {
	for k, vars := range m.MD {
		for _, v := range vars {
			if err := h(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m MetaDataReaderWriter) Set(k, v string) {
	key := strings.ToLower(k)
	m.MD[key] = append(m.MD[key], v)
}

func JaegerClientInterceptor(tracer opentracing.Tracer) grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context, method string,
		req, reply interface{}, cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

		var parentCtx opentracing.SpanContext
		//判断ctx里面又没有span信息，没有就生成一个
		if parent := opentracing.SpanFromContext(ctx); parent != nil {
			parentCtx = parent.Context()
		}

		span := tracer.StartSpan(method, opentracing.ChildOf(parentCtx), TracingComponentTag, ext.SpanKindRPCClient)
		defer span.Finish()
		//从context中获取metadata
		meta, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			meta = metadata.New(nil)
		} else {
			//对metadata进行修改需要用到拷贝的副本
			meta = meta.Copy()
		}
		writer := MetaDataReaderWriter{meta}
		if err := tracer.Inject(span.Context(), opentracing.TextMap, writer); err != nil {
			log.Printf("inject to metadata err:%v ", err)
		}
		//创建一个新的context并将metadata带上
		ctx = metadata.NewOutgoingContext(ctx, meta)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func JaegerServerInterceptor(tracer opentracing.Tracer) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context, req interface{},
		info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

		meta, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			meta = metadata.New(nil)
		}

		spanCtx, err := tracer.Extract(opentracing.TextMap, MetaDataReaderWriter{meta})
		if err != nil && errors.Is(err, opentracing.ErrSpanContextNotFound) {
			log.Printf("extract from metadata err %v ", err)
		}

		span := tracer.StartSpan(info.FullMethod, ext.RPCServerOption(spanCtx), TracingComponentTag, ext.SpanKindRPCServer)
		defer span.Finish()
		ctx = opentracing.ContextWithSpan(ctx, span)
		return handler(ctx, req)
	}
}
