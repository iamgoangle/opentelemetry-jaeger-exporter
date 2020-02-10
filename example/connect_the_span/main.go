package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/otel/api/core"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/key"

	"go.opentelemetry.io/otel/exporter/trace/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

const (
	jaegerThriftEndpoint = "http://localhost:14268/api/traces"

	serviceName = "opentelemetry-jaeger-exporter"
)

// initTracer creates a new trace provider instance and registers it as global trace provider.
func initTracer() func() {
	// Create Jaeger Exporter
	exporter, err := jaeger.NewExporter(
		jaeger.WithCollectorEndpoint(jaegerThriftEndpoint),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: serviceName,
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

func main() {
	fn := initTracer()
	defer fn()

	ctx := context.Background()

	tr := global.TraceProvider().Tracer("main")
	ctx, span := tr.Start(ctx, "main-span")

	sayHello(ctx, "Golf")
	go asyncJob(ctx)
	sayGoodBye(ctx, "Golf")

	span.End()
}

func sayHello(ctx context.Context, name string) {
	trace := global.TraceProvider().Tracer("sayHello")
	spanName := fmt.Sprintf("Hello, %s", name)

	ctx, span := trace.Start(ctx, spanName)
	defer span.End()

	time.Sleep(10 * time.Second)

	sayProfileByName(ctx, name)
}

func sayProfileByName(ctx context.Context, name string) {
	trace := global.TraceProvider().Tracer("sayProfileByName")
	spanName := fmt.Sprintf("Profile, %s", name)

	_, span := trace.Start(ctx, spanName)
	defer span.End()

	time.Sleep(10 * time.Second)
}

func asyncJob(ctx context.Context) {
	trace := global.TraceProvider().Tracer("asyncJob")
	spanName := fmt.Sprintf("Async Job")

	_, span := trace.Start(ctx, spanName)
	defer span.End()

	time.Sleep(2 * time.Second)
}

func sayGoodBye(ctx context.Context, name string) {
	trace := global.TraceProvider().Tracer("sayGoodBye")
	spanName := fmt.Sprintf("Bye, %s", name)

	_, span := trace.Start(ctx, spanName)
	defer span.End()

	time.Sleep(10 * time.Second)
}
