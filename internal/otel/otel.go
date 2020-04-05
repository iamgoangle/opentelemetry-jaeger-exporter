package otel

import (
	"context"
	"fmt"
	"log"
	"runtime"

	"go.opentelemetry.io/otel/api/trace"
	"google.golang.org/grpc/codes"

	"go.opentelemetry.io/otel/api/core"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/key"
	"go.opentelemetry.io/otel/exporter/trace/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

//go:generate mockgen -destination=./mock/otel.go -package=mock github.com/iamgoangle/opentelemetry-jaeger-exporter/internal/otel Tracer

type Config struct {
	Service        string
	ThriftEndpoint string
}

type Tracer interface {
	TracerStart(ctx context.Context, name string) (context.Context, trace.Span)

	StartSpanWithContext(ctx context.Context, name string, fn func(ctx context.Context) error) error

	SetStringAttribute(ctx context.Context, k, v string)

	SetIntAttribute(ctx context.Context, k string, v int)

	SetJaegerStatusOK(ctx context.Context)

	SetJaegerStatusCanceled(ctx context.Context)

	SetJaegerStatusInternal(ctx context.Context)

	PrintSpanContext(ctx context.Context)

	TraceID(ctx context.Context) string

	SpanID(ctx context.Context) string
}

type tracing struct{}

// InitTracer creates a new trace provider instance and registers it as global trace provider.
func InitTracer(c *Config) func() {
	exporter, err := jaeger.NewExporter(
		jaeger.WithCollectorEndpoint(c.ThriftEndpoint),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: c.Service,
			Tags: []core.KeyValue{
				key.String("exporter", "jaeger"),
				key.Float64("float", 312.23),
			},
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	// For demoing purposes, always sample. In a production application, you should
	// configure this to a trace.ProbabilitySampler set at the desired
	// probability.
	tp, err := sdktrace.NewProvider(
		sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
		sdktrace.WithSyncer(exporter))
	if err != nil {
		log.Fatal(err)
	}

	global.SetTraceProvider(tp)

	return func() {
		exporter.Flush()
	}
}

func NewTracer() Tracer {
	return &tracing{}
}

// TracerStart trace start wrapper
// https://github.com/open-telemetry/opentelemetry-go/blob/master/example/basic/main.go
func (t *tracing) TracerStart(ctx context.Context, name string) (context.Context, trace.Span) {
	tr := global.TraceProvider().Tracer(name)

	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	callerDetails := fmt.Sprintf("%s - %s#%d", frame.Function, frame.File, frame.Line)

	// fnName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()

	// trace.SpanFromContext(ctx).AddEvent(ctx, "Nice operation!", key.New("bogons").Int(100))
	// fileKey := key.New("file")
	// fnKey := key.New("function")
	// trace.SpanFromContext(ctx).SetAttributes(fileKey.String(callerDetails))
	// trace.SpanFromContext(ctx).SetAttributes(fnKey.String(fnName))

	t.SetStringAttribute(ctx, "file", callerDetails)

	return tr.Start(ctx, name)
}

func (t *tracing) StartSpanWithContext(ctx context.Context, name string, fn func(ctx context.Context) error) error {
	tr := global.TraceProvider().Tracer(name)
	err := tr.WithSpan(ctx, name, fn)
	if err != nil {
		return err
	}

	return nil
}

func (t *tracing) SetStringAttribute(ctx context.Context, k, v string) {
	keyName := key.New(k)
	trace.SpanFromContext(ctx).SetAttributes(keyName.String(v))
}

func (t *tracing) SetIntAttribute(ctx context.Context, k string, v int) {
	keyName := key.New(k)
	trace.SpanFromContext(ctx).SetAttributes(keyName.Int(v))
}

func (t *tracing) SetJaegerStatusOK(ctx context.Context) {
	trace.SpanFromContext(ctx).SetStatus(codes.OK)
}

func (t *tracing) SetJaegerStatusCanceled(ctx context.Context) {
	trace.SpanFromContext(ctx).SetStatus(codes.Canceled)
}

func (t *tracing) SetJaegerStatusInternal(ctx context.Context) {
	trace.SpanFromContext(ctx).SetStatus(codes.Internal)
}

func (t *tracing) PrintSpanContext(ctx context.Context) {
	fmt.Printf("%+v", trace.SpanFromContext(ctx).SpanContext())
}

func (t *tracing) TraceID(ctx context.Context) string {
	return trace.SpanFromContext(ctx).SpanContext().TraceIDString()
}

func (t *tracing) SpanID(ctx context.Context) string {
	return trace.SpanFromContext(ctx).SpanContext().SpanIDString()
}
