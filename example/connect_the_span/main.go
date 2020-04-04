package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"google.golang.org/grpc/codes"

	"go.opentelemetry.io/otel/api/global"

	"github.com/iamgoangle/opentelemetry-jaeger-exporter/internal/otel"
)

type TestHandler struct {
	Tracer otel.Tracer
}

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

	test := &TestHandler{
		Tracer: tracer,
	}
	test.handler(ctx)
}

func (t *TestHandler) handler(ctx context.Context) {
	spanName := "handler"
	thisCtx, span := t.Tracer.TracerStart(ctx, spanName)
	defer span.End()

	time.Sleep(10 * time.Second)

	t.Tracer.SetIntAttribute(thisCtx, "http.code", http.StatusOK)
	span.SetStatus(codes.OK)

	fmt.Println(t.Tracer.TraceID(thisCtx))
	fmt.Println(t.Tracer.SpanID(thisCtx))

	t.service(thisCtx)
}

func (t *TestHandler) service(ctx context.Context) {
	spanName := "service"
	thisCtx, span := t.Tracer.TracerStart(ctx, spanName)
	defer span.End()

	time.Sleep(10 * time.Second)

	t.repository(thisCtx)
}

func (t *TestHandler) repository(ctx context.Context) {
	spanName := "repository"
	_, span := t.Tracer.TracerStart(ctx, spanName)
	defer span.End()

	time.Sleep(10 * time.Second)
}