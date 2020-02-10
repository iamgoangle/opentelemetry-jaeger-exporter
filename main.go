package main

import (
	"context"
	"log"

	"go.opentelemetry.io/exporter/trace/jaeger"
	"go.opentelemetry.io/otel/api/core"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/key"
	"go.opentelemetry.io/otel/api/trace"
)

// initTracer creates a new trace provider instance and registers it as global trace provider.
func initTracer() func() {
	// Create Jaeger Exporter
	exporter, err := jaeger.NewExporter(
		jaeger.WithCollectorEndpoint("http://localhost:14268/api/traces"),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: "trace-demo",
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

	tr := global.TraceProvider().Tracer("component-main")
	ctx, span := tr.Start(ctx, "foo")
	bar(ctx)
	span.End()
}

func sayHello() {
	ctx := context.Background()
	tracer := trace.GlobalTracer()

	ctx, trace := tracer.Start(ctx, "say-hello")

	trace.End()
}
