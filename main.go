package main

import (
	"context"

	"go.opentelemetry.io/otel/api/global"

	"github.com/iamgoangle/opentelemetry-jaeger-exporter/internal/otel"

	"github.com/iamgoangle/opentelemetry-jaeger-exporter/api"
)

func main() {
	initTrace := otel.InitTracer(&otel.Config{
		Service:        "my-app",
		ThriftEndpoint: "http://localhost:14268/api/traces",
	})
	defer initTrace()

	ctx := context.Background()

	tr := global.TraceProvider().Tracer("main")
	ctx, span := tr.Start(ctx, "main-span")
	defer span.End()

	tracer := otel.NewTracer()

	apiTest := &api.TestHandler{
		Tracer: tracer,
	}
	apiTest.Handler(ctx)
}
